// Package transport builds the gRPC transport credentials shared by the broker and the middleware.
package transport

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc/credentials"
)

// ClientTLSCredentials returns credentials for dialing TLS servers. Server certificates are
// verified against the CA certificate(s) in caFile, or against the system roots if caFile is empty.
func ClientTLSCredentials(caFile string) (credentials.TransportCredentials, error) {
	if caFile == "" {
		return credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12}), nil
	}
	creds, err := credentials.NewClientTLSFromFile(caFile, "")
	if err != nil {
		return nil, fmt.Errorf("failed to load CA file %s: %w", caFile, err)
	}
	return creds, nil
}

// ServerTLSCredentials returns credentials for a TLS server. Both files are required.
func ServerTLSCredentials(certFile, keyFile string) (credentials.TransportCredentials, error) {
	if certFile == "" || keyFile == "" {
		return nil, fmt.Errorf("TLS requires both a certificate file and a key file")
	}
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load TLS certificate: %w", err)
	}
	return creds, nil
}
