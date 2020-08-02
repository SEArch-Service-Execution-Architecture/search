package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/google/uuid"

	pb "dc.uba.ar/this/search/protobuf"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	"github.com/vishalkuo/bimap"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	caFile     = flag.String("ca_file", "", "The file containing the CA root cert file")
	port       = flag.Int("port", 10000, "The server port")
	brokerAddr = flag.String("broker_addr", "localhost", "The server address in the format of host:port")
	brokerPort = flag.Int("broker_port", 10000, "The port in which the broker is listening")
	bufferSize = flag.Int("buffer_size", 100, "The size of each connection's  buffer")
)

type middlewareServer struct {
	pb.UnimplementedPublicMiddlewareServer
	pb.UnimplementedPrivateMiddlewareServer

	// local provider apps. key: appID, value: RegisterAppServer (connection is kept open
	// with each local app so as to notify new channels)
	registeredApps map[string]pb.PrivateMiddleware_RegisterAppServer

	// channels already brokered. key: ChannelID
	brokeredChannels map[string]SEARCHChannel

	// channels registered by local apps that have not yet been used. key: LocalID
	unBrokeredChannels map[string]*SEARCHChannel

	// mapping of channels' LocalID <--> ID (global)
	localChannels *bimap.BiMap
}

func newMiddlewareServer() *middlewareServer {
	var s middlewareServer
	s.localChannels = bimap.NewBiMap()
	s.registeredApps = make(map[string]pb.PrivateMiddleware_RegisterAppServer)
	s.brokeredChannels = make(map[string]SEARCHChannel)
	s.unBrokeredChannels = make(map[string]*SEARCHChannel)

	return &s
}

type SEARCHChannel struct {
	LocalID  uuid.UUID // the identifier the local app uses to identify the channel
	ID       uuid.UUID // channel identifier assigned by the broker
	Contract pb.Contract

	addresses   map[string]*pb.RemoteParticipant                      // map participant names to remote URLs and AppIDs
	connections map[string]*pb.PublicMiddleware_MessageExchangeServer // map participant names to streams

	// buffers for incoming/outgoing messages from/to each participant
	Outgoing map[string]chan pb.MessageContent
	Incoming map[string]chan pb.MessageContent
}

// connect to the broker, send contract, wait for result and save data in the channel
func (r *SEARCHChannel) broker(mw *middlewareServer) {
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		// creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		// if err != nil {
		// 	log.Fatalf("Failed to create TLS credentials %v", err)
		//}
		//opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *brokerAddr, *brokerPort), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	brokerresult, err := client.BrokerChannel(ctx, &pb.BrokerChannelRequest{Contract: &r.Contract})
	if err != nil {
		log.Fatalf("%v.BrokerChannel(_) = _, %v: ", client, err)
	}

	r.addresses = brokerresult.GetParticipants()
	r.ID = uuid.MustParse(brokerresult.GetChannelId())

	// TODO: use mutex to handle maps
	mw.brokeredChannels[r.ID.String()] = *r
	delete(mw.unBrokeredChannels, r.LocalID.String())
	mw.localChannels.Insert(r.LocalID.String(), r.ID.String())
}

func newSEARCHChannel(contract pb.Contract) *SEARCHChannel {
	var r SEARCHChannel
	r.ID = uuid.New()
	r.Contract = contract
	r.addresses = make(map[string]*pb.RemoteParticipant)
	r.connections = make(map[string]*pb.PublicMiddleware_MessageExchangeServer)

	r.Outgoing = make(map[string]chan pb.MessageContent)
	r.Incoming = make(map[string]chan pb.MessageContent)
	for _, p := range contract.GetRemoteParticipants() {
		r.Outgoing[p] = make(chan pb.MessageContent, *bufferSize)
		r.Incoming[p] = make(chan pb.MessageContent, *bufferSize)
	}

	return &r
}

func (s *middlewareServer) RegisterApp(req *pb.RegisterAppRequest, stream pb.PrivateMiddleware_RegisterAppServer) error {
	// TODO: talk to the Registry to get my app_id
	mockUUID := uuid.New() // TODO: this should be generated by the Registry
	ack := pb.RegisterAppResponse{
		AckOrNew: &pb.RegisterAppResponse_AppId{
			AppId: mockUUID.String()}}
	if err := stream.Send(&ack); err != nil {
		return err
	}
	s.registeredApps[mockUUID.String()] = stream
	return nil
}

func (s *middlewareServer) RegisterChannel(ctx context.Context, in *pb.RegisterChannelRequest) (*pb.RegisterChannelResponse, error) {
	c := newSEARCHChannel(*in.GetRequirementsContract())
	s.unBrokeredChannels[c.LocalID.String()] = c
	return &pb.RegisterChannelResponse{ChannelId: c.LocalID.String()}, nil
}

func (s *middlewareServer) AppSend(ctx context.Context, req *pb.ApplicationMessageOut) (*pb.AppSendResponse, error) {
	localID := req.ChannelId
	c, ok := s.unBrokeredChannels[localID]
	if ok {
		// channel has not been brokered
		go c.broker(s)
	}
	c.Outgoing[req.Recipient] <- *req.Content

	return &pb.AppSendResponse{Result: 0}, nil // TODO: use enum instead of 0
}

// When the middleware receives a message in its public interface, it must enqueue it so that
// the corresponding local app can receive it
func (s *middlewareServer) MessageExchange(stream pb.PublicMiddleware_MessageExchangeServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("Received message from", in.SenderId, ":", string(in.Content.Body))

		// TODO: send to local app replacing sender name with local name

		ack := pb.ApplicationMessageWithHeaders{
			ChannelId:   in.ChannelId,
			RecipientId: in.SenderId,
			SenderId:    "provmwID-44",
			Content:     &pb.MessageContent{Body: []byte("ack")}}

		if err := stream.Send(&ack); err != nil {
			return err
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	pms := newMiddlewareServer()
	pb.RegisterPublicMiddlewareServer(grpcServer, pms)
	grpcServer.Serve(lis)
}