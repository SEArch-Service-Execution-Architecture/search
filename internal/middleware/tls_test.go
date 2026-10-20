package middleware

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/SEArch-Service-Execution-Architecture/search/contract"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/SEArch-Service-Execution-Architecture/search/internal/broker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// writeSelfSignedCert writes a self-signed certificate for localhost to dir and returns the
// paths of the certificate and its key. The certificate doubles as the CA for clients.
func writeSelfSignedCert(t *testing.T, dir string) (certFile, keyFile string) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		DNSNames:              []string{"localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("::1")},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	require.NoError(t, err)
	keyDER, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)
	certFile = filepath.Join(dir, "cert.pem")
	keyFile = filepath.Join(dir, "key.pem")
	require.NoError(t, os.WriteFile(certFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}), 0o600))
	require.NoError(t, os.WriteFile(keyFile, pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}), 0o600))
	return certFile, keyFile
}

func startTLSMiddleware(t *testing.T, brokerAddr, certFile, keyFile string) (*MiddlewareServer, pb.PrivateMiddlewareServiceClient) {
	t.Helper()
	mw := NewMiddlewareServer(brokerAddr)
	require.NoError(t, mw.SetClientTLS(certFile))
	var wg sync.WaitGroup
	mw.StartMiddlewareServer(&wg, "", "localhost:0", "localhost:0", true, certFile, keyFile, nil)
	t.Cleanup(func() {
		mw.Stop()
		wg.Wait()
	})
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	require.NoError(t, err)
	conn, err := grpc.NewClient(mw.PrivateURL, grpc.WithTransportCredentials(creds))
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })
	return mw, pb.NewPrivateMiddlewareServiceClient(conn)
}

// With -tls, every hop must use TLS: middleware -> broker, broker -> middleware and
// middleware -> middleware.
func TestTLSEndToEnd(t *testing.T) {
	certFile, keyFile := writeSelfSignedCert(t, t.TempDir())

	bs := broker.NewBrokerServer(fmt.Sprintf("%s/t.db", t.TempDir()))
	require.NoError(t, bs.SetClientTLS(certFile))
	bs.SetCompatFunc(func(ctx context.Context, req, prov contract.LocalContract) (bool, map[string]string, error) {
		return true, map[string]string{"Client": "Client"}, nil
	})
	brokerAddr := make(chan string, 1)
	go bs.StartServer("localhost:0", true, certFile, keyFile, brokerAddr)
	addr := <-brokerAddr
	t.Cleanup(bs.Stop)

	_, providerClient := startTLSMiddleware(t, addr, certFile, keyFile)
	_, clientClient := startTLSMiddleware(t, addr, certFile, keyFile)

	// Provider registers with the broker through its middleware.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	regStream, err := providerClient.RegisterApp(ctx, &pb.RegisterAppRequest{
		ProviderContract: &pb.LocalContract{Contract: []byte(".outputs Srv\n.state graph\n0 Client ? hi 1\n.marking 0\n.end\n"), Format: pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA},
	})
	require.NoError(t, err)
	ack, err := regStream.Recv()
	require.NoError(t, err)
	require.NotEmpty(t, ack.GetAppId())

	// Service Client sends a message, which triggers the brokerage.
	reg, err := clientClient.RegisterChannel(ctx, &pb.RegisterChannelRequest{
		RequirementsContract: &pb.GlobalContract{
			Contract:      []byte(".outputs Client\n.state graph\n0 Srv ! hi 1\n.marking 0\n.end\n\n.outputs Srv\n.state graph\n0 Client ? hi 1\n.marking 0\n.end\n"),
			Format:        pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA,
			InitiatorName: "Client",
		},
	})
	require.NoError(t, err)
	_, err = clientClient.AppSend(ctx, &pb.AppSendRequest{ChannelId: reg.GetChannelId(), Recipient: "Srv", Message: &pb.AppMessage{Type: "hi", Body: []byte("hello over TLS")}})
	require.NoError(t, err)

	// Provider is notified of the new channel and receives the message.
	notification, err := regStream.Recv()
	require.NoError(t, err)
	channelID := notification.GetNotification().GetChannelId()
	require.NotEmpty(t, channelID)
	recv, err := providerClient.AppRecv(ctx, &pb.AppRecvRequest{ChannelId: channelID, Participant: "Client"})
	require.NoError(t, err)
	require.Equal(t, "hello over TLS", string(recv.GetMessage().GetBody()))
}
