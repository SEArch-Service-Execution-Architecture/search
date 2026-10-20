package middleware

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/sourcegraph/conc/pool"

	"github.com/SEArch-Service-Execution-Architecture/search/contract"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/SEArch-Service-Execution-Architecture/search/internal/searcherrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"

	"github.com/vishalkuo/bimap"
)

var (
	bufferSize = 100 // number of messages that can be buffered in each channel.
)

type MiddlewareServer struct {
	pb.UnimplementedPublicMiddlewareServiceServer
	pb.UnimplementedPrivateMiddlewareServiceServer

	// Local Service Providers. key: appID, value: RegisterAppServer (connection is kept open
	// with each local app so as to notify new channels)
	registeredApps map[string]registeredApp
	providersLock  *sync.Mutex

	// channels already brokered. key: ChannelID
	brokeredChannels map[string]*SEARCHChannel
	// channels registered by local apps that have not yet been brokered. key: LocalID
	unBrokeredChannels map[string]*SEARCHChannel
	// channels being brokered. key: LocalID
	brokeringChannels map[string]*SEARCHChannel
	// channles for which brokering failed. key: LocalID, value: list of participants for which the broker was unable to find matches.
	brokeringFailedChannels map[string][]string
	// mapping of channels' LocalID <--> ID (global). This only makes sense for already brokered channels.
	localChannels *bimap.BiMap[string, string]
	channelLock   *sync.RWMutex // protects all previous maps/bimaps for channels

	brokerAddr string // maybe replace with net.url.URL type?

	// pointers to gRPC servers
	publicServer  *grpc.Server
	privateServer *grpc.Server
	PublicURL     string
	PrivateURL    string

	// logger with prefix
	logger *log.Logger
}

// NewMiddlewareServer is MiddlewareServer's constructor
func NewMiddlewareServer(brokerAddr string) *MiddlewareServer {
	var s MiddlewareServer
	s.localChannels = bimap.NewBiMap[string, string]() // mapping between local channelID and global channelID. When initiator not local, they are equal
	s.registeredApps = make(map[string]registeredApp)
	s.providersLock = new(sync.Mutex)

	s.brokeredChannels = make(map[string]*SEARCHChannel)   // channels already brokered (locally initiated or not)
	s.unBrokeredChannels = make(map[string]*SEARCHChannel) // channels locally registered but not yet brokered
	s.brokeringChannels = make(map[string]*SEARCHChannel)  // channels being brokered
	s.brokeringFailedChannels = make(map[string][]string)  // channels for which brokering failed
	s.channelLock = new(sync.RWMutex)

	s.brokerAddr = brokerAddr
	s.logger = log.New(os.Stderr, "[MIDDLEWARE] - ", log.LstdFlags|log.Lmsgprefix)
	return &s
}

// Struct to represent each Service Provider that is registered in the middleware.
// We notify to the local Service Provider that a new channel is being initiatied for it using the NotifyChan.
type registeredApp struct {
	Contract   contract.LocalContract
	NotifyChan chan initChannelNotification // Each message in the channel is a Channel ID (extracted from pb.InitChannelNotification).
}

// We use this struct internally to avoid copying the protobuf structure with its lock.
type initChannelNotification struct {
	ChannelID string
}

// SEARCHChannel represents what we call a "channel" in SEARCH
type SEARCHChannel struct {
	LocalID    uuid.UUID // the identifier the local app uses to identify the channel
	ID         uuid.UUID // channel identifier assigned by the broker
	Contract   contract.Contract
	ContractPB *pb.GlobalContract // We only save this when this middleware is representing the Service Client.
	AppID      uuid.UUID

	// localParticipant string                           // Name of the local participant in the contract.
	addresses          map[string]*pb.RemoteParticipant // map participant names to remote URLs and AppIDs, indexed by participant name
	participants       map[string]string                // participant names indexed by AppID
	presetParticipants map[string]*pb.RemoteParticipant

	// buffers for incoming/outgoing messages from/to each participant
	Outgoing map[string]chan *pb.AppMessage
	Incoming map[string]chan *pb.AppMessage

	sendersPool           *pool.ContextPool  // pool of workers for sending outbound messages.
	sendersPoolCancelFunc context.CancelFunc // cancel function for the pool of workers that send outbound messages.

	brokerageDone chan struct{} // closed when the brokerage process finishes (nil until it starts).
	brokerageErr  error         // result of the brokerage process. Only valid after brokerageDone is closed.

	// pointer to middleware
	mw *MiddlewareServer
}

// newSEARCHChannel returns a pointer to a new channel created for the MiddlewareServer.
// localParticipant is the name of the local app in the contract.
func (mw *MiddlewareServer) newSEARCHChannel(c contract.Contract, pbContract *pb.GlobalContract, presetParticipants map[string]*pb.RemoteParticipant) (*SEARCHChannel, error) {
	// TODO: review ContractPB. Instead of keeping a reference to the protobuf, maybe keep only a copy
	// of the contract bytes?
	var r SEARCHChannel
	r.mw = mw
	r.LocalID = uuid.New()
	r.Contract = c
	r.ContractPB = pbContract
	r.addresses = make(map[string]*pb.RemoteParticipant)
	r.participants = make(map[string]string)
	r.presetParticipants = presetParticipants

	r.Outgoing = make(map[string]chan *pb.AppMessage)
	r.Incoming = make(map[string]chan *pb.AppMessage)
	for _, p := range c.GetRemoteParticipantNames() {
		r.Outgoing[p] = make(chan *pb.AppMessage, bufferSize)
		r.Incoming[p] = make(chan *pb.AppMessage, bufferSize)
	}

	return &r, nil
}

func (s *MiddlewareServer) connectBroker() (pb.BrokerServiceClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // TODO: use tls
	conn, err := grpc.NewClient(s.brokerAddr, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create connection to broker addr %s: %w", s.brokerAddr, err)
	}
	return pb.NewBrokerServiceClient(conn), conn, nil
}

// broker connects to the broker, sends the contract, waits for the result and saves it in the channel.
// It must be called with the channel in mw.brokeringChannels, and it moves the channel to either
// mw.brokeredChannels or mw.brokeringFailedChannels.
func (r *SEARCHChannel) broker(ctx context.Context) error {
	r.mw.logger.Printf("Requesting brokerage of contract: '%v'", r.Contract)
	localID := r.LocalID.String()
	presetParticipants := make(map[string]*pb.RemoteParticipant)
	for k, v := range r.presetParticipants {
		presetParticipants[k] = v
	}
	presetParticipants[r.ContractPB.InitiatorName] = &pb.RemoteParticipant{
		Url:   r.mw.PublicURL,
		AppId: localID, // we use the channel's LocalID as AppID for Service Clients.
	}
	var brokerresult *pb.BrokerChannelResponse
	client, conn, err := r.mw.connectBroker()
	if err == nil {
		defer conn.Close()
		brokerresult, err = client.BrokerChannel(ctx, &pb.BrokerChannelRequest{
			Contract:           r.ContractPB,
			PresetParticipants: presetParticipants,
		})
	}
	var channelID uuid.UUID
	if err == nil {
		channelID, err = uuid.Parse(brokerresult.GetChannelId())
		if err != nil {
			err = fmt.Errorf("broker returned an invalid channel ID %q: %w", brokerresult.GetChannelId(), err)
		}
	}

	r.mw.channelLock.Lock()
	defer r.mw.channelLock.Unlock()
	delete(r.mw.brokeringChannels, localID)
	if err != nil {
		r.mw.logger.Printf("brokering failed for channel with LocalID %s. broker err: %v", localID, err)
		failedParticipants := []string{}
		if st, ok := status.FromError(err); ok {
			for _, detail := range st.Details() {
				if info, ok := detail.(*errdetails.ErrorInfo); ok && info.GetMetadata()["failed_participants"] != "" {
					failedParticipants = strings.Split(info.GetMetadata()["failed_participants"], ",")
					break
				}
			}
		}
		r.mw.brokeringFailedChannels[localID] = failedParticipants
		return searcherrors.BrokerageFailedError(failedParticipants)
	}

	r.addresses = brokerresult.GetParticipants()
	for k, v := range r.addresses {
		r.participants[v.AppId] = k
	}
	r.ID = channelID
	r.mw.brokeredChannels[r.ID.String()] = r
	r.mw.localChannels.Insert(localID, r.ID.String())

	r.startSenderRoutines()
	return nil
}

// RegisterApp is invoked by Service Providers on the private interface in order to register their service with
// the broker. This function keeps the connection open between the middleware and the local Service Provider
// to notify it of new channels being initiated for it. When the provider disconnects, it is unregistered from
// this middleware.
func (s *MiddlewareServer) RegisterApp(req *pb.RegisterAppRequest, stream pb.PrivateMiddlewareService_RegisterAppServer) error {
	provContract, err := contract.ConvertPBLocalContract(req.GetProviderContract())
	if err != nil {
		s.logger.Printf("Error parsing protobuf contract. %v", err)
		return status.Errorf(codes.InvalidArgument, "invalid provider contract: %v", err)
	}
	client, conn, err := s.connectBroker()
	if err != nil {
		s.logger.Printf("Error connecting to broker: %v", err)
		return status.Error(codes.Unavailable, "error connecting to broker")
	}
	defer conn.Close()

	notifyInitChannel := make(chan initChannelNotification, bufferSize)
	appID, err := func() (string, error) {
		// Hold the lock until the app is saved: the broker may send an InitChannel for this app
		// as soon as RegisterProvider returns.
		s.providersLock.Lock()
		defer s.providersLock.Unlock()
		s.logger.Printf("Sending provider contract to broker...")
		res, err := client.RegisterProvider(stream.Context(), &pb.RegisterProviderRequest{
			Contract: req.GetProviderContract(),
			Url:      s.PublicURL,
		})
		if err != nil {
			return "", err
		}
		s.registeredApps[res.GetAppId()] = registeredApp{
			Contract:   provContract,
			NotifyChan: notifyInitChannel,
		}
		return res.GetAppId(), nil
	}()
	if err != nil {
		s.logger.Printf("Error registering provider with broker: %v", err)
		return status.Error(codes.Unavailable, "error registering provider with broker")
	}
	defer func() {
		s.providersLock.Lock()
		delete(s.registeredApps, appID)
		s.providersLock.Unlock()
		s.logger.Printf("Provider with appID %s disconnected, unregistered it from this middleware", appID)
	}()

	// Send ACK to local app via the stream.
	s.logger.Printf("Sending ACK to provider to confirm successful registration with appID %s", appID)
	ack := pb.RegisterAppResponse{
		AckOrNew: &pb.RegisterAppResponse_AppId{
			AppId: appID}}
	if err := stream.Send(&ack); err != nil {
		return err
	}

	for {
		select {
		case <-stream.Context().Done():
			return status.FromContextError(stream.Context().Err()).Err()
		case newChan := <-notifyInitChannel:
			msg := pb.RegisterAppResponse{
				AckOrNew: &pb.RegisterAppResponse_Notification{
					Notification: &pb.InitChannelNotification{
						ChannelId: newChan.ChannelID,
					},
				},
			}
			if err := stream.Send(&msg); err != nil {
				return err
			}
		}
	}
}

// Invoked by local Service Client with a requirements contract.
func (s *MiddlewareServer) RegisterChannel(ctx context.Context, in *pb.RegisterChannelRequest) (*pb.RegisterChannelResponse, error) {
	// TODO: create a monitor for this channel.
	requirementsContract, err := contract.ConvertPBGlobalContract(in.GetRequirementsContract())
	if err != nil {
		s.logger.Printf("Error parsing protobuf contract. %v", err)
		return nil, err
	}
	c, err := s.newSEARCHChannel(requirementsContract, in.GetRequirementsContract(), in.GetPresetParticipants())
	if err != nil {
		// TODO: newSEARCHChannel may fail because of other kinds of errors in the future.
		err := status.Error(codes.InvalidArgument, "invalid contract")
		return nil, err
	}
	c.AppID = c.LocalID
	s.channelLock.Lock()
	s.unBrokeredChannels[c.LocalID.String()] = c // this has to be changed when brokering
	s.channelLock.Unlock()
	return &pb.RegisterChannelResponse{ChannelId: c.LocalID.String()}, nil
}

// routine that actually sends messages to remote particpant on a channel. When ctx is cancelled, it
// sends any messages still buffered for the participant before closing the stream.
func (r *SEARCHChannel) sender(ctx context.Context, participant string) error {
	r.mw.logger.Printf("Started sender routine for channel %s, participant %s\n", r.ID, participant)
	outgoing := r.Outgoing[participant]
	// wait for first message
	var msg *pb.AppMessage
	select {
	case msg = <-outgoing:
	case <-ctx.Done():
		select {
		case msg = <-outgoing:
			// A message was enqueued right before the channel was closed: deliver it.
		default:
			r.mw.logger.Printf("Context cancelled (cause: %s). Closing sender routine for channel %s, participant %s. No connection established to remote Service Provider.", context.Cause(ctx), r.ID, participant)
			return nil
		}
	}
	r.mw.logger.Printf("Received first outbound message to send on channel %s for participant %s. Opening connection to remote Service Provider...", r.ID, participant)
	// connect and save stream in r.outStreams
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // TODO: use tls
	provconn, err := grpc.NewClient(r.addresses[participant].GetUrl(), opts...)
	if err != nil {
		return fmt.Errorf("fail to dial: %w", err)
	}
	defer provconn.Close()
	provClient := pb.NewPublicMiddlewareServiceClient(provconn)

	provClientctx, provClientCancel := context.WithCancel(context.Background())
	defer provClientCancel()
	stream, err := provClient.MessageExchange(provClientctx)
	if err != nil {
		return fmt.Errorf("failed to establish MessageExchange: %w", err)
	}
	r.mw.logger.Printf("Established connection to remote Service Provider for channel %s, participant %s\n", r.ID, participant)
	send := func(msg *pb.AppMessage) error {
		messageWithHeaders := pb.MessageExchangeRequest{
			SenderId:    r.AppID.String(),
			ChannelId:   r.ID.String(),
			RecipientId: r.addresses[participant].AppId,
			Content:     msg}
		if err := stream.Send(&messageWithHeaders); err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}
		r.mw.logger.Printf("Sent message to remote Service Provider for channel %s, participant %s\n", r.ID, participant)
		return nil
	}
	for {
		if err := send(msg); err != nil {
			return err
		}
		select {
		case msg = <-outgoing:
			r.mw.logger.Printf("Received outbound message to send on channel %s for participant %s\n", r.ID, participant)
			continue
		case <-ctx.Done():
		}
		r.mw.logger.Printf("Context cancelled (cause: %s). Flushing and closing sender routine for channel %s, participant %s\n", context.Cause(ctx), r.ID, participant)
		for flushing := true; flushing; {
			select {
			case msg = <-outgoing:
				if err := send(msg); err != nil {
					return err
				}
			default:
				flushing = false
			}
		}
		reply, err := stream.CloseAndRecv()
		if err != nil {
			r.mw.logger.Printf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			return err
		}
		if reply.Result != pb.MessageExchangeResponse_RESULT_OK {
			r.mw.logger.Printf("Error in MessageExchange when attempting to close stream: %s", reply.ErrorMessage)
			return fmt.Errorf("error in MessageExchange when attempting to close stream: %s", reply.ErrorMessage)
		}
		r.mw.logger.Printf("Closed sender stream for channel %s, participant %s\n", r.ID, participant)
		return nil
	}
}

// getChannelForUsage returns the channel with the given LocalID. If the channel has not been brokered yet,
// it starts the brokerage. If blockUntilBrokered is true, it waits (until ctx is done) for the brokerage to
// finish.
func (s *MiddlewareServer) getChannelForUsage(ctx context.Context, localID string, blockUntilBrokered bool) (*SEARCHChannel, error) {
	s.channelLock.Lock()
	if failedParticipants, ok := s.brokeringFailedChannels[localID]; ok {
		s.channelLock.Unlock()
		return nil, searcherrors.BrokerageFailedError(failedParticipants)
	}
	if c, ok := s.unBrokeredChannels[localID]; ok {
		// First use of the channel: start brokering it.
		delete(s.unBrokeredChannels, localID)
		s.brokeringChannels[localID] = c
		c.brokerageDone = make(chan struct{})
		go func() {
			// The brokerage must outlive the request that triggered it.
			c.brokerageErr = c.broker(context.Background())
			close(c.brokerageDone)
		}()
	}
	if c, ok := s.brokeringChannels[localID]; ok {
		s.channelLock.Unlock()
		if !blockUntilBrokered {
			return c, nil
		}
		s.logger.Printf("waiting for brokerage to finish for channel %s\n", localID)
		select {
		case <-c.brokerageDone:
		case <-ctx.Done():
			return nil, status.FromContextError(ctx.Err()).Err()
		}
		if c.brokerageErr != nil {
			return nil, c.brokerageErr
		}
		return c, nil
	}
	defer s.channelLock.Unlock()
	channelID, ok := s.localChannels.Get(localID)
	if !ok {
		s.logger.Printf("getChannelForUsage invoked on channel ID %s: there's no localChannel with that ID.", localID)
		return nil, status.Errorf(codes.NotFound, "channel %s not found", localID)
	}
	return s.brokeredChannels[channelID], nil
}

// simple (g)rpc the local app uses when sending a message to a remote participant on an already registered channel
func (s *MiddlewareServer) AppSend(ctx context.Context, req *pb.AppSendRequest) (*pb.AppSendResponse, error) {
	c, err := s.getChannelForUsage(ctx, req.GetChannelId(), false)
	if err != nil {
		return nil, err
	}
	outgoing, ok := c.Outgoing[req.GetRecipient()]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "%q is not a participant of channel %s", req.GetRecipient(), req.GetChannelId())
	}
	// Enqueue message in outgoing buffer. This blocks while the buffer is full.
	select {
	case outgoing <- req.GetMessage():
	case <-ctx.Done():
		return nil, status.FromContextError(ctx.Err()).Err()
	}
	s.logger.Printf("Enqueued message of type %s in outgoing buffer for channel %s, participant %s\n", req.GetMessage().GetType(), req.GetChannelId(), req.GetRecipient())
	return &pb.AppSendResponse{Result: pb.AppSendResponse_RESULT_OK}, nil
}

func (s *MiddlewareServer) AppRecv(ctx context.Context, req *pb.AppRecvRequest) (*pb.AppRecvResponse, error) {
	c, err := s.getChannelForUsage(ctx, req.GetChannelId(), true)
	if err != nil {
		return nil, err
	}
	incoming, ok := c.Incoming[req.GetParticipant()]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "%q is not a participant of channel %s", req.GetParticipant(), req.GetChannelId())
	}
	select {
	case msg := <-incoming:
		return &pb.AppRecvResponse{
			ChannelId: req.GetChannelId(),
			Sender:    req.GetParticipant(),
			Message:   msg,
		}, nil
	case <-ctx.Done():
		return nil, status.FromContextError(ctx.Err()).Err()
	}
}

// A local app (Service Client or Service Provider) invokes this when it wants to close a channel, using the
// channel ID it knows (the LocalID for Service Clients). Closing fails if:
// 1. There are still inbound messages that the app should receive (by running AppRecv). The channel stays usable.
// 2. The channelID is not registered in the middleware, or its brokerage failed.
// 3. Outbound messages could not be delivered.
// If the channel is being brokered, CloseChannel waits for the brokerage to finish. Outbound messages that are
// already buffered are delivered before the channel is closed. A channel that was registered but never used is
// simply discarded.
func (s *MiddlewareServer) CloseChannel(ctx context.Context, req *pb.CloseChannelRequest) (*pb.CloseChannelResponse, error) {
	channelID := req.GetChannelId()
	s.logger.Printf("CloseChannel. ChannelID: %s", channelID)
	s.channelLock.Lock()
	if _, ok := s.unBrokeredChannels[channelID]; ok {
		delete(s.unBrokeredChannels, channelID)
		s.channelLock.Unlock()
		return &pb.CloseChannelResponse{Result: pb.CloseChannelResponse_RESULT_CLOSED}, nil
	}
	s.channelLock.Unlock()

	c, err := s.getChannelForUsage(ctx, channelID, true)
	if err != nil {
		s.channelLock.Lock()
		delete(s.brokeringFailedChannels, channelID)
		s.channelLock.Unlock()
		return nil, err
	}

	// Check that there are no inbound messages in the buffers of the channel.
	participantsWithIncoming := make([]string, 0)
	for participant, chMsg := range c.Incoming {
		if len(chMsg) > 0 {
			participantsWithIncoming = append(participantsWithIncoming, participant)
		}
	}
	if len(participantsWithIncoming) > 0 {
		slices.Sort(participantsWithIncoming)
		return &pb.CloseChannelResponse{
			Result:                         pb.CloseChannelResponse_RESULT_PENDING_INBOUND,
			ErrorMessage:                   "There are still inbound messages to receive!",
			ParticipantsWithPendingInbound: participantsWithIncoming,
		}, nil
	}

	// Stop the sender goroutines. They deliver the messages still in the outbound buffers before exiting.
	if c.sendersPoolCancelFunc != nil {
		c.sendersPoolCancelFunc()
	}
	var sendErr error
	if c.sendersPool != nil {
		sendErr = c.sendersPool.Wait()
	}

	s.channelLock.Lock()
	delete(s.brokeredChannels, c.ID.String())
	s.localChannels.Delete(c.LocalID.String())
	s.channelLock.Unlock()

	if sendErr != nil {
		return nil, status.Errorf(codes.Unavailable, "channel closed, but failed to deliver outbound messages: %v", sendErr)
	}
	return &pb.CloseChannelResponse{
		Result: pb.CloseChannelResponse_RESULT_CLOSED,
	}, nil
}

// When the middleware receives a message in its public interface, it must enqueue it so that
// the corresponding local app can receive it. Each stream carries messages from a single sender
// on a single channel.
func (s *MiddlewareServer) MessageExchange(stream pb.PublicMiddlewareService_MessageExchangeServer) error {
	s.logger.Print("Received MessageExchange...")
	in, err := stream.Recv()
	if err == io.EOF {
		s.logger.Print("Received EOF in MessageExchange without receiving any message...")
		return err
	}
	if err != nil {
		s.logger.Printf("Error in MessageExchange when attempting to recv from stream: %s", err)
		return err
	}
	channelID, senderID := in.GetChannelId(), in.GetSenderId()
	s.channelLock.RLock()
	c, ok := s.brokeredChannels[channelID]
	s.channelLock.RUnlock()
	if !ok {
		// TODO: attempt to get the channel from s.brokeringChannels (this can happen when some provider starts sending messages to the Service Client
		// before we have finished processing the BrokerChannelResponse)
		s.logger.Printf("Received MessageExchange with ChannelID %s but it is not a brokered Channel in this middleware.", channelID)
		return status.Errorf(codes.NotFound, "channel %s is not a brokered channel in this middleware", channelID)
	}
	// TODO: we could be hosting two different apps from the same channel; brokeredChannels can only hold one.
	if in.GetRecipientId() != c.AppID.String() {
		return status.Errorf(codes.InvalidArgument, "recipient %s is not the app of channel %s in this middleware", in.GetRecipientId(), channelID)
	}
	participantName, ok := c.participants[senderID]
	if !ok {
		return status.Errorf(codes.InvalidArgument, "sender %s is not a participant of channel %s", senderID, channelID)
	}
	incoming, ok := c.Incoming[participantName]
	if !ok {
		return status.Errorf(codes.InvalidArgument, "sender %s is not a remote participant of channel %s", senderID, channelID)
	}

	for {
		s.logger.Printf("Received message from %s of type %s", participantName, in.GetContent().GetType())
		select {
		case incoming <- in.GetContent():
		case <-stream.Context().Done():
			return status.FromContextError(stream.Context().Err()).Err()
		}
		s.logger.Printf("Enqueued message of type %s in incoming buffer for channel %s, participant %s", in.GetContent().GetType(), channelID, participantName)

		s.logger.Printf("Waiting for next message from %s...", participantName)
		in, err = stream.Recv()
		if err == io.EOF {
			s.logger.Print("Received EOF in MessageExchange...")
			err = stream.SendAndClose(&pb.MessageExchangeResponse{Result: pb.MessageExchangeResponse_RESULT_OK})
			if err != nil {
				s.logger.Printf("Error when closing MessageExchange after receiving EOF: %s", err)
				return err
			}
			s.logger.Print("Successful close of MessageExchange after receiving EOF. Exiting MessageExchange...")
			return nil
		}
		if err != nil {
			s.logger.Printf("Error attempting to receive msg from %s in MessageExchange: %s", participantName, err)
			return err
		}
		if in.GetChannelId() != channelID || in.GetSenderId() != senderID {
			return status.Error(codes.InvalidArgument, "all messages in a MessageExchange stream must have the same channel and sender")
		}
	}
}

// rpc invoked by broker when initializing channel for a Service Provider that this middleware represents. This is received
// when the broker has matched one of our registered (provider) apps with a requirements contract.
func (s *MiddlewareServer) InitChannel(ctx context.Context, icr *pb.InitChannelRequest) (*pb.InitChannelResponse, error) {
	s.logger.Printf("Received InitChannel. ChannelID: %s. AppID: %s", icr.GetChannelId(), icr.GetAppId())
	s.logger.Printf("InitChannel mapping received:%v", icr.GetParticipants())
	appID, err := uuid.Parse(icr.GetAppId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid AppID %q", icr.GetAppId())
	}
	channelID, err := uuid.Parse(icr.GetChannelId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid ChannelID %q", icr.GetChannelId())
	}
	s.channelLock.Lock()
	defer s.channelLock.Unlock()
	s.providersLock.Lock()
	defer s.providersLock.Unlock()
	regapp, ok := s.registeredApps[icr.GetAppId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "received InitChannel for unknown AppID %s", icr.GetAppId())
	}
	r, err := s.newSEARCHChannel(regapp.Contract, nil, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create channel: %v", err)
	}
	r.AppID = appID
	r.LocalID = channelID // in non locally initiated channels, LocalID == ID
	r.ID = channelID
	r.addresses = icr.GetParticipants()
	for k, v := range r.addresses {
		r.participants[v.AppId] = k
	}
	s.brokeredChannels[icr.GetChannelId()] = r
	s.localChannels.Insert(r.LocalID.String(), r.ID.String())

	return &pb.InitChannelResponse{Result: pb.InitChannelResponse_RESULT_ACK}, nil
}

// StartChannel is fired by the broker and sent to all participants (initiator included) to signal that
// all participants are ready and communication can start.
func (s *MiddlewareServer) StartChannel(ctx context.Context, req *pb.StartChannelRequest) (*pb.StartChannelResponse, error) {
	s.logger.Printf("StartChannel. ChannelID: %s. AppID: %s", req.ChannelId, req.AppId)
	s.channelLock.RLock()
	c, ok := s.brokeredChannels[req.ChannelId]
	s.channelLock.RUnlock()
	if !ok {
		s.logger.Printf("ERROR: Received StartChannel on non brokered ChannelID: %s", req.ChannelId)
		return nil, status.Error(codes.InvalidArgument, "invalid ChannelID")
	}

	// notify local provider app
	s.providersLock.Lock()
	regapp, ok := s.registeredApps[req.AppId]
	if !ok {
		s.providersLock.Unlock()
		return nil, status.Error(codes.InvalidArgument, "invalid AppID")
	}
	s.providersLock.Unlock()
	select {
	case regapp.NotifyChan <- initChannelNotification{ChannelID: req.ChannelId}:
	case <-ctx.Done():
		return nil, status.FromContextError(ctx.Err()).Err()
	}

	c.startSenderRoutines()

	return &pb.StartChannelResponse{Result: pb.StartChannelResponse_RESULT_ACK}, nil
}

func (c *SEARCHChannel) startSenderRoutines() {
	sendersCtx, cancel := context.WithCancel(context.Background())
	c.sendersPoolCancelFunc = cancel
	c.sendersPool = pool.New().WithContext(sendersCtx).WithCancelOnError()
	for _, p := range c.Contract.GetRemoteParticipantNames() {
		thisP := p
		c.sendersPool.Go(func(cont context.Context) error {
			return c.sender(cont, thisP)
		})
	}
}

// StartServer starts gRPC middleware server
func (s *MiddlewareServer) StartMiddlewareServer(wg *sync.WaitGroup, publicURL string, publicAddr string, privateAddr string, tls bool, certFile string, keyFile string, notifyListening chan string) {
	lisPub, err := net.Listen("tcp", publicAddr)
	if err != nil {
		s.logger.Fatalf("failed to listen on public interface: %v", err)
	}
	if publicURL == "" {
		s.PublicURL = fmt.Sprintf("dns:%s", lisPub.Addr().String())
	} else {
		s.PublicURL = publicURL
	}
	s.logger = log.New(os.Stderr, fmt.Sprintf("[MIDDLEWARE] %s - ", s.PublicURL), log.LstdFlags|log.Lmsgprefix|log.Lshortfile)

	lisPriv, err := net.Listen("tcp", privateAddr)
	if err != nil {
		s.logger.Fatalf("failed to listen on private interface: %v", err)
	}
	s.PrivateURL = lisPriv.Addr().String()
	if notifyListening != nil {
		notifyListening <- s.PrivateURL
	}

	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			certFile = testdata.Path("server1.pem")
		}
		if keyFile == "" {
			keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			s.logger.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	publicGrpcServer := grpc.NewServer(opts...)
	privateGrpcServer := grpc.NewServer(opts...)

	s.publicServer = publicGrpcServer
	s.privateServer = privateGrpcServer

	// Register GRPC healthcheck service on both servers.
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(publicGrpcServer, healthServer)
	healthpb.RegisterHealthServer(privateGrpcServer, healthServer)

	pb.RegisterPublicMiddlewareServiceServer(publicGrpcServer, s)
	pb.RegisterPrivateMiddlewareServiceServer(privateGrpcServer, s)

	wg.Add(2)

	go func() {
		defer wg.Done()
		// s.logger.Println("Waiting for SIGINT or SIGTERM on public server.")
		if err := publicGrpcServer.Serve(lisPub); err != nil {
			s.logger.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		// s.logger.Println("Waiting for SIGINT or SIGTERM on private server.")
		if err := privateGrpcServer.Serve(lisPriv); err != nil {
			s.logger.Fatalf("failed to serve: %v", err)
		}
	}()

}

func (s *MiddlewareServer) Stop() {
	// TODO: flush all buffers before stopping!
	if s.publicServer != nil {
		s.publicServer.Stop()
	}
	if s.privateServer != nil {
		s.privateServer.Stop()
	}
}
