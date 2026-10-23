package main

import (
	"context"
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SEArch-Service-Execution-Architecture/search/examples/credit-card-payments/creditcard"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
)

var middlewareURL = flag.String("middleware-url", "middleware-payments-service:11000", "The URL for the middleware")

func main() {
	flag.Parse()
	var logger = log.New(os.Stderr, "[PPS] - ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
	conn, err := grpc.NewClient(*middlewareURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatalf("Error connecting to middleware URL %s", *middlewareURL)
	}
	defer conn.Close()

	err = creditcard.ServePPS(context.Background(), pb.NewPrivateMiddlewareServiceClient(conn), logger, func(string) {
		// Create an empty file for the Docker Compose healthcheck.
		f, err := os.Create("/tmp/registered")
		if err != nil {
			logger.Printf("Error creating /tmp/registered: %v", err)
			return
		}
		f.Close()
	})
	if err != nil {
		logger.Fatal(err)
	}
}
