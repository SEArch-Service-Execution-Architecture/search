package broker

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	pb "github.com/clpombo/search/api"
	"google.golang.org/grpc"
)

func init() {
	go StartServer(3333, false, "", "", "")
}

func TestBrokerChannel_Request(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", "localhost", 3333), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewBrokerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c := pb.Contract{
		Contract: "hola",
		RemoteParticipants: []string{"self", "p1"},
	}
	brokerresult, err := client.BrokerChannel(ctx, &pb.BrokerChannelRequest{Contract: &c})
	if err != nil {
		t.Error("Received error from broker.")
	}
	if brokerresult.Result != pb.BrokerChannelResponse_ACK {
		t.Error("Non ACK return code from broker")
	}
}
