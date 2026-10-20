package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SEArch-Service-Execution-Architecture/search/internal/broker"
)

var (
	tls          = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile     = flag.String("cert_file", "", "The TLS cert file")
	keyFile      = flag.String("key_file", "", "The TLS key file")
	caFile       = flag.String("ca_file", "", "The CA root cert file used to verify middlewares when using TLS (defaults to the system roots)")
	host         = flag.String("host", "", "The host on which the broker listens (defaults to all interfaces)")
	port         = flag.Int("port", 10000, "Port number on which the broker listens")
	databaseFile = flag.String("database_file", "SEARCHbroker.db", "The path for the broker database")
	python       = flag.String("python", broker.PythonExecutable, "Python interpreter with the cfsm-bisimulation package installed")
)

func main() {
	flag.Parse()
	b := broker.NewBrokerServer(*databaseFile)
	if *tls {
		if err := b.SetClientTLS(*caFile); err != nil {
			log.Fatalf("failed to set up TLS: %v", err)
		}
	}

	// Set the bisimilarity algorithm as the default algorithm for the broker.
	broker.PythonExecutable = *python
	b.SetCompatFunc(broker.BisimilarityAlgorithm)
	address := fmt.Sprintf("%s:%d", *host, *port)
	b.StartServer(address, *tls, *certFile, *keyFile, nil)
}
