package middleware

import (
	"context"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// fakeBroker lets tests control the broker's answers.
type fakeBroker struct {
	pb.UnimplementedBrokerServiceServer
	registerProvider func(context.Context, *pb.RegisterProviderRequest) (*pb.RegisterProviderResponse, error)
	brokerChannel    func(context.Context, *pb.BrokerChannelRequest) (*pb.BrokerChannelResponse, error)
	registerCalls    atomic.Int32
}

func (f *fakeBroker) RegisterProvider(ctx context.Context, req *pb.RegisterProviderRequest) (*pb.RegisterProviderResponse, error) {
	f.registerCalls.Add(1)
	if f.registerProvider != nil {
		return f.registerProvider(ctx, req)
	}
	return &pb.RegisterProviderResponse{AppId: uuid.NewString()}, nil
}

func (f *fakeBroker) BrokerChannel(ctx context.Context, req *pb.BrokerChannelRequest) (*pb.BrokerChannelResponse, error) {
	if f.brokerChannel != nil {
		return f.brokerChannel(ctx, req)
	}
	return nil, status.Error(codes.Unimplemented, "not configured")
}

// brokerSucceedsWith returns a BrokerChannel implementation that assigns the given remote
// participants (in addition to the preset ones) after an optional delay.
func brokerSucceedsWith(delay time.Duration, brokered map[string]*pb.RemoteParticipant) func(context.Context, *pb.BrokerChannelRequest) (*pb.BrokerChannelResponse, error) {
	return func(ctx context.Context, req *pb.BrokerChannelRequest) (*pb.BrokerChannelResponse, error) {
		time.Sleep(delay)
		participants := make(map[string]*pb.RemoteParticipant)
		for k, v := range req.GetPresetParticipants() {
			participants[k] = v
		}
		for k, v := range brokered {
			participants[k] = v
		}
		return &pb.BrokerChannelResponse{ChannelId: uuid.NewString(), Participants: participants}, nil
	}
}

func startFakeBroker(t *testing.T, fb *fakeBroker) string {
	t.Helper()
	lis, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	srv := grpc.NewServer()
	pb.RegisterBrokerServiceServer(srv, fb)
	go func() { _ = srv.Serve(lis) }()
	t.Cleanup(srv.Stop)
	return lis.Addr().String()
}

func startTestMiddleware(t *testing.T, brokerAddr string) (*MiddlewareServer, pb.PrivateMiddlewareServiceClient) {
	t.Helper()
	mw := NewMiddlewareServer(brokerAddr)
	var wg sync.WaitGroup
	mw.StartMiddlewareServer(&wg, "", "localhost:0", "localhost:0", false, "", "", nil)
	t.Cleanup(func() {
		mw.Stop()
		wg.Wait()
	})
	conn, err := grpc.NewClient(mw.PrivateURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })
	return mw, pb.NewPrivateMiddlewareServiceClient(conn)
}

func publicClient(t *testing.T, mw *MiddlewareServer) pb.PublicMiddlewareServiceClient {
	t.Helper()
	conn, err := grpc.NewClient(mw.PublicURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })
	return pb.NewPublicMiddlewareServiceClient(conn)
}

// waitOrFail runs f and fails the test if it does not return within timeout.
func waitOrFail(t *testing.T, timeout time.Duration, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Fatalf("timed out after %s waiting for %s", timeout, what)
	}
}

// Provider contract with two remote participants, A and B.
const twoParticipantsProviderFSA = `.outputs
.state graph
0 A ? m 1
1 B ? m 2
.marking 0
.end
`

// Global contract for a Service Client "Client" talking to "Srv".
const clientServerGlobalFSA = `.outputs Client
.state graph
0 Srv ! hi 1
1 Srv ? bye 2
.marking 0
.end

.outputs Srv
.state graph
0 Client ? hi 1
1 Client ! bye 2
.marking 0
.end
`

// registerTestApp registers a provider app through RegisterApp and returns its AppID. The
// registration stream stays open until the test ends (or cancel is called).
func registerTestApp(t *testing.T, client pb.PrivateMiddlewareServiceClient, fsa string) (appID string, cancel context.CancelFunc) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	stream, err := client.RegisterApp(ctx, &pb.RegisterAppRequest{
		ProviderContract: &pb.LocalContract{Contract: []byte(fsa), Format: pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA},
	})
	require.NoError(t, err)
	ack, err := stream.Recv()
	require.NoError(t, err)
	require.NotEmpty(t, ack.GetAppId())
	return ack.GetAppId(), cancel
}

// initTestChannel creates a provider-side channel for appID, as if the broker had sent InitChannel.
func initTestChannel(t *testing.T, mw *MiddlewareServer, appID string, participants map[string]*pb.RemoteParticipant) string {
	t.Helper()
	channelID := uuid.NewString()
	res, err := mw.InitChannel(context.Background(), &pb.InitChannelRequest{
		ChannelId:    channelID,
		AppId:        appID,
		Participants: participants,
	})
	require.NoError(t, err)
	require.Equal(t, pb.InitChannelResponse_RESULT_ACK, res.GetResult())
	return channelID
}

func remoteParticipants(names ...string) map[string]*pb.RemoteParticipant {
	res := make(map[string]*pb.RemoteParticipant)
	for _, n := range names {
		res[n] = &pb.RemoteParticipant{Url: "localhost:1", AppId: uuid.NewString()}
	}
	return res
}

// M1: a failed broker registration must not leave providersLock locked.
func TestRegisterAppBrokerFailureDoesNotDeadlock(t *testing.T) {
	fb := &fakeBroker{registerProvider: func(context.Context, *pb.RegisterProviderRequest) (*pb.RegisterProviderResponse, error) {
		return nil, status.Error(codes.Internal, "broker is having a bad day")
	}}
	_, client := startTestMiddleware(t, startFakeBroker(t, fb))

	for i := range 2 {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		stream, err := client.RegisterApp(ctx, &pb.RegisterAppRequest{
			ProviderContract: &pb.LocalContract{Contract: []byte(twoParticipantsProviderFSA), Format: pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA},
		})
		require.NoError(t, err)
		_, err = stream.Recv()
		cancel()
		require.Error(t, err)
		require.NotEqual(t, codes.DeadlineExceeded, status.Code(err), "RegisterApp attempt %d hung", i+1)
	}
}

// M1: an invalid contract must be rejected before it is sent to the broker.
func TestRegisterAppRejectsInvalidContract(t *testing.T) {
	fb := &fakeBroker{}
	_, client := startTestMiddleware(t, startFakeBroker(t, fb))

	stream, err := client.RegisterApp(context.Background(), &pb.RegisterAppRequest{
		ProviderContract: &pb.LocalContract{Contract: []byte("not a contract"), Format: pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA},
	})
	require.NoError(t, err)
	_, err = stream.Recv()

	require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
	require.Zero(t, fb.registerCalls.Load(), "invalid contract was sent to the broker")
}

// M2: when the provider app disconnects, the middleware must stop serving its registration.
func TestRegisterAppUnregistersOnDisconnect(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, cancel := registerTestApp(t, client, twoParticipantsProviderFSA)

	cancel()

	require.Eventually(t, func() bool {
		mw.providersLock.Lock()
		defer mw.providersLock.Unlock()
		_, ok := mw.registeredApps[appID]
		return !ok
	}, 3*time.Second, 10*time.Millisecond, "app still registered after disconnecting")
}

// M3: malformed IDs coming from the network must not panic.
func TestInitChannelInvalidChannelID(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)

	_, err := mw.InitChannel(context.Background(), &pb.InitChannelRequest{
		ChannelId:    "not-a-uuid",
		AppId:        appID,
		Participants: remoteParticipants("A", "B"),
	})

	require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
}

// M4: unexpected error details from the broker must not panic.
func TestBrokerageWithUnexpectedErrorDetails(t *testing.T) {
	fb := &fakeBroker{brokerChannel: func(context.Context, *pb.BrokerChannelRequest) (*pb.BrokerChannelResponse, error) {
		st, err := status.New(codes.NotFound, "nope").WithDetails(&errdetails.DebugInfo{Detail: "surprise"})
		if err != nil {
			return nil, err
		}
		return nil, st.Err()
	}}
	mw, client := startTestMiddleware(t, startFakeBroker(t, fb))
	reg, err := client.RegisterChannel(context.Background(), &pb.RegisterChannelRequest{
		RequirementsContract: &pb.GlobalContract{Contract: []byte(clientServerGlobalFSA), Format: pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA, InitiatorName: "Client"},
	})
	require.NoError(t, err)

	var recvErr error
	waitOrFail(t, 5*time.Second, "AppRecv", func() {
		_, recvErr = mw.AppRecv(context.Background(), &pb.AppRecvRequest{ChannelId: reg.GetChannelId(), Participant: "Srv"})
	})

	require.Error(t, recvErr)
}

// M5: sending to a participant that is not in the contract must fail instead of blocking.
func TestAppSendUnknownRecipient(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	channelID := initTestChannel(t, mw, appID, remoteParticipants("A", "B"))

	var err error
	waitOrFail(t, 3*time.Second, "AppSend", func() {
		_, err = mw.AppSend(context.Background(), &pb.AppSendRequest{ChannelId: channelID, Recipient: "Nobody", Message: &pb.AppMessage{Type: "m"}})
	})

	require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
}

// M5: receiving from a participant that is not in the contract must fail instead of blocking.
func TestAppRecvUnknownParticipant(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	channelID := initTestChannel(t, mw, appID, remoteParticipants("A", "B"))

	var err error
	waitOrFail(t, 3*time.Second, "AppRecv", func() {
		_, err = mw.AppRecv(context.Background(), &pb.AppRecvRequest{ChannelId: channelID, Participant: "Nobody"})
	})

	require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
}

// M5: AppRecv must give up when the caller's context ends.
func TestAppRecvRespectsContext(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	channelID := initTestChannel(t, mw, appID, remoteParticipants("A", "B"))
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	var err error
	waitOrFail(t, 3*time.Second, "AppRecv", func() {
		_, err = mw.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: channelID, Participant: "A"})
	})

	require.Equal(t, codes.DeadlineExceeded, status.Code(err), "got %v", err)
}

// M5: messages from unknown senders, or for another app, must be rejected instead of blocking.
func TestMessageExchangeRejectsUnknownSenderAndRecipient(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	participants := remoteParticipants("A", "B")
	channelID := initTestChannel(t, mw, appID, participants)
	public := publicClient(t, mw)

	cases := map[string]*pb.MessageExchangeRequest{
		"unknown sender":    {ChannelId: channelID, SenderId: uuid.NewString(), RecipientId: appID, Content: &pb.AppMessage{Type: "m"}},
		"unknown recipient": {ChannelId: channelID, SenderId: participants["A"].AppId, RecipientId: uuid.NewString(), Content: &pb.AppMessage{Type: "m"}},
	}
	for name, msg := range cases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			stream, err := public.MessageExchange(ctx)
			require.NoError(t, err)
			require.NoError(t, stream.Send(msg))
			_, err = stream.CloseAndRecv()
			require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
		})
	}
}

// brokeredClientChannel registers a Service Client channel and completes its brokerage. It
// returns the channel's LocalID (the only ID the client app knows).
func brokeredClientChannel(t *testing.T, client pb.PrivateMiddlewareServiceClient) string {
	t.Helper()
	reg, err := client.RegisterChannel(context.Background(), &pb.RegisterChannelRequest{
		RequirementsContract: &pb.GlobalContract{Contract: []byte(clientServerGlobalFSA), Format: pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA, InitiatorName: "Client"},
	})
	require.NoError(t, err)
	// AppRecv blocks until the brokerage finishes; nobody sends anything, so it times out.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err = client.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: reg.GetChannelId(), Participant: "Srv"})
	require.Equal(t, codes.DeadlineExceeded, status.Code(err), "got %v", err)
	return reg.GetChannelId()
}

// M6: a Service Client closes its channel using the LocalID it got from RegisterChannel.
func TestCloseChannelByLocalID(t *testing.T) {
	fb := &fakeBroker{brokerChannel: brokerSucceedsWith(0, remoteParticipants("Srv"))}
	_, client := startTestMiddleware(t, startFakeBroker(t, fb))
	localID := brokeredClientChannel(t, client)

	res, err := client.CloseChannel(context.Background(), &pb.CloseChannelRequest{ChannelId: localID})

	require.NoError(t, err)
	require.Equal(t, pb.CloseChannelResponse_RESULT_CLOSED, res.GetResult())
}

// M7: a closed channel must reject further use instead of panicking or silently dropping data.
func TestAppSendAfterCloseChannel(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	channelID := initTestChannel(t, mw, appID, remoteParticipants("A", "B"))
	_, err := mw.StartChannel(context.Background(), &pb.StartChannelRequest{ChannelId: channelID, AppId: appID})
	require.NoError(t, err)

	res, err := client.CloseChannel(context.Background(), &pb.CloseChannelRequest{ChannelId: channelID})
	require.NoError(t, err)
	require.Equal(t, pb.CloseChannelResponse_RESULT_CLOSED, res.GetResult())

	_, err = client.AppSend(context.Background(), &pb.AppSendRequest{ChannelId: channelID, Recipient: "A", Message: &pb.AppMessage{Type: "m"}})
	require.Equal(t, codes.NotFound, status.Code(err), "got %v", err)
}

// M7: a CloseChannel that reports pending inbound messages must leave the channel usable.
func TestCloseChannelWithPendingInboundKeepsChannelUsable(t *testing.T) {
	mw, client := startTestMiddleware(t, startFakeBroker(t, &fakeBroker{}))
	appID, _ := registerTestApp(t, client, twoParticipantsProviderFSA)
	participants := remoteParticipants("A", "B")
	channelID := initTestChannel(t, mw, appID, participants)
	// B sends a message that the app has not received yet.
	stream, err := publicClient(t, mw).MessageExchange(context.Background())
	require.NoError(t, err)
	require.NoError(t, stream.Send(&pb.MessageExchangeRequest{ChannelId: channelID, SenderId: participants["B"].AppId, RecipientId: appID, Content: &pb.AppMessage{Type: "m"}}))
	_, err = stream.CloseAndRecv()
	require.NoError(t, err)

	res, err := client.CloseChannel(context.Background(), &pb.CloseChannelRequest{ChannelId: channelID})
	require.NoError(t, err)
	require.Equal(t, pb.CloseChannelResponse_RESULT_PENDING_INBOUND, res.GetResult())
	require.Equal(t, []string{"B"}, res.GetParticipantsWithPendingInbound())

	// Receiving from A must still wait for a real message (not return a nil one, nor hang forever).
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	waitOrFail(t, 3*time.Second, "AppRecv from A", func() {
		_, err = mw.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: channelID, Participant: "A"})
	})
	require.Equal(t, codes.DeadlineExceeded, status.Code(err), "got %v", err)
	recv, err := mw.AppRecv(context.Background(), &pb.AppRecvRequest{ChannelId: channelID, Participant: "B"})
	require.NoError(t, err)
	require.Equal(t, "m", recv.GetMessage().GetType())
}

// M7: messages enqueued before CloseChannel must all be delivered.
func TestCloseChannelFlushesOutboundMessages(t *testing.T) {
	const numMessages = 50
	brokerAddr := startFakeBroker(t, &fakeBroker{})
	senderMw, senderClient := startTestMiddleware(t, brokerAddr)
	receiverMw, receiverClient := startTestMiddleware(t, brokerAddr)
	senderAppID, _ := registerTestApp(t, senderClient, ".outputs\n.state graph\n0 R ! m 0\n.marking 0\n.end\n")
	receiverAppID, _ := registerTestApp(t, receiverClient, ".outputs\n.state graph\n0 S ? m 0\n.marking 0\n.end\n")
	channelID := uuid.NewString()
	_, err := senderMw.InitChannel(context.Background(), &pb.InitChannelRequest{
		ChannelId: channelID, AppId: senderAppID,
		Participants: map[string]*pb.RemoteParticipant{"R": {Url: receiverMw.PublicURL, AppId: receiverAppID}},
	})
	require.NoError(t, err)
	_, err = receiverMw.InitChannel(context.Background(), &pb.InitChannelRequest{
		ChannelId: channelID, AppId: receiverAppID,
		Participants: map[string]*pb.RemoteParticipant{"S": {Url: senderMw.PublicURL, AppId: senderAppID}},
	})
	require.NoError(t, err)
	_, err = senderMw.StartChannel(context.Background(), &pb.StartChannelRequest{ChannelId: channelID, AppId: senderAppID})
	require.NoError(t, err)

	for i := range numMessages {
		_, err := senderClient.AppSend(context.Background(), &pb.AppSendRequest{ChannelId: channelID, Recipient: "R", Message: &pb.AppMessage{Type: "m", Body: fmt.Appendf(nil, "%d", i)}})
		require.NoError(t, err)
	}
	res, err := senderClient.CloseChannel(context.Background(), &pb.CloseChannelRequest{ChannelId: channelID})
	require.NoError(t, err)
	require.Equal(t, pb.CloseChannelResponse_RESULT_CLOSED, res.GetResult())

	for i := range numMessages {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		recv, err := receiverClient.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: channelID, Participant: "S"})
		cancel()
		require.NoError(t, err, "message %d was not delivered", i)
		require.Equal(t, fmt.Sprintf("%d", i), string(recv.GetMessage().GetBody()))
	}
}

// M8: concurrent blocking calls while a channel is being brokered must all wait for the brokerage.
func TestConcurrentCallsDuringBrokerage(t *testing.T) {
	fb := &fakeBroker{brokerChannel: brokerSucceedsWith(300*time.Millisecond, remoteParticipants("Srv"))}
	_, client := startTestMiddleware(t, startFakeBroker(t, fb))
	reg, err := client.RegisterChannel(context.Background(), &pb.RegisterChannelRequest{
		RequirementsContract: &pb.GlobalContract{Contract: []byte(clientServerGlobalFSA), Format: pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA, InitiatorName: "Client"},
	})
	require.NoError(t, err)

	var wg sync.WaitGroup
	errs := make([]error, 3)
	for i := range errs {
		wg.Go(func() {
			// Nothing is ever sent, so after a successful brokerage each call times out.
			ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
			defer cancel()
			_, errs[i] = client.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: reg.GetChannelId(), Participant: "Srv"})
		})
		time.Sleep(20 * time.Millisecond)
	}
	wg.Wait()

	for i, err := range errs {
		require.Equal(t, codes.DeadlineExceeded, status.Code(err), "call %d got %v", i, err)
	}
}
