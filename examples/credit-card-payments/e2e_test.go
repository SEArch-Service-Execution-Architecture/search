package creditcardpayments_test

import (
	"context"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SEArch-Service-Execution-Architecture/search/examples/credit-card-payments/creditcard"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/SEArch-Service-Execution-Architecture/search/internal/broker"
	"github.com/SEArch-Service-Execution-Architecture/search/internal/middleware"
	"github.com/SEArch-Service-Execution-Architecture/search/internal/testenv"
)

// TestCreditCardPayments runs the whole example in-process: a broker using the real Python
// bisimulation, one middleware per participant, the Python backend (Srv) as a subprocess, and
// the Go payments service (PPS) and client (ClientApp).
func TestCreditCardPayments(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping end-to-end test in short mode")
	}
	testenv.RequirePython(t, broker.PythonExecutable, "cfsm_bisimulation", "grpclib", "betterproto", "simplejson")

	brokerAddr := startBroker(t)
	backendMiddleware, _ := startMiddleware(t, brokerAddr)
	_, ppsStub := startMiddleware(t, brokerAddr)
	_, clientStub := startMiddleware(t, brokerAddr)
	startBackend(t, backendMiddleware)
	startPPS(t, ppsStub)

	contract, err := os.ReadFile("contract.fsa")
	require.NoError(t, err)

	// The second purchase reuses the compatibility results cached by the broker.
	for i := range 2 {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		order := creditcard.Order{
			Items:           map[string]int{"1984": 2, "The Hobbit": 1},
			ShippingAddress: "742 Evergreen Terrace",
		}
		purchase, err := creditcard.StartPurchase(ctx, clientStub, contract, order)
		require.NoError(t, err, "purchase %d", i)
		// The backend charges 100 per book plus 10 for shipping.
		require.Equal(t, 310.0, purchase.TotalAmount)

		receipt, err := purchase.Pay(creditcard.Card{Number: "4111111111111111", ExpirationDate: "12/30", SecurityCode: "123"})
		require.NoError(t, err, "purchase %d", i)
		require.Equal(t, "1234567890", receipt.Nonce)
		require.True(t, receipt.Approved)

		resp, err := clientStub.CloseChannel(ctx, &pb.CloseChannelRequest{ChannelId: purchase.ChannelID})
		require.NoError(t, err)
		require.Equal(t, pb.CloseChannelResponse_RESULT_CLOSED, resp.GetResult())
	}
}

func startBroker(t *testing.T) string {
	t.Helper()
	b := broker.NewBrokerServer(filepath.Join(t.TempDir(), "broker.db"))
	b.SetCompatFunc(broker.BisimilarityAlgorithm)
	addr := make(chan string, 1)
	go b.StartServer("localhost:0", false, "", "", addr)
	t.Cleanup(b.Stop)
	return <-addr
}

// startMiddleware returns the private address of a new middleware and a client for it.
func startMiddleware(t *testing.T, brokerAddr string) (string, pb.PrivateMiddlewareServiceClient) {
	t.Helper()
	mw := middleware.NewMiddlewareServer(brokerAddr)
	var wg sync.WaitGroup
	mw.StartMiddlewareServer(&wg, "", "localhost:0", "localhost:0", false, "", "", nil)
	t.Cleanup(func() {
		mw.Stop()
		wg.Wait()
	})
	conn, err := grpc.NewClient(mw.PrivateURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })
	return mw.PrivateURL, pb.NewPrivateMiddlewareServiceClient(conn)
}

// startBackend runs the Python backend against the middleware at middlewareAddr and waits
// until it has registered.
func startBackend(t *testing.T, middlewareAddr string) {
	t.Helper()
	host, port, err := net.SplitHostPort(middlewareAddr)
	require.NoError(t, err)
	registeredFile := filepath.Join(t.TempDir(), "registered")
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, broker.PythonExecutable, "main.py",
		"--middleware-host", host, "--middleware-port", port, "--registered-file", registeredFile)
	cmd.Dir = "backend"
	out := &testLogWriter{t: t, prefix: "[backend] "}
	cmd.Stdout, cmd.Stderr = out, out
	cmd.WaitDelay = 5 * time.Second
	require.NoError(t, cmd.Start())
	exited := make(chan struct{})
	go func() {
		_ = cmd.Wait()
		close(exited)
	}()
	t.Cleanup(func() {
		cancel()
		<-exited
	})

	deadline := time.After(30 * time.Second)
	for {
		if _, err := os.Stat(registeredFile); err == nil {
			return
		}
		select {
		case <-exited:
			t.Fatalf("backend exited before registering")
		case <-deadline:
			t.Fatalf("backend did not register within 30s")
		case <-time.After(50 * time.Millisecond):
		}
	}
}

// startPPS runs the payments service and waits until it has registered.
func startPPS(t *testing.T, stub pb.PrivateMiddlewareServiceClient) {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	registered := make(chan struct{})
	done := make(chan error, 1)
	logger := log.New(&testLogWriter{t: t, prefix: "[PPS] "}, "", 0)
	go func() {
		done <- creditcard.ServePPS(ctx, stub, logger, func(string) { close(registered) })
	}()
	t.Cleanup(func() {
		cancel()
		<-done
	})
	select {
	case <-registered:
	case err := <-done:
		t.Fatalf("payments service exited before registering: %v", err)
	case <-time.After(30 * time.Second):
		t.Fatalf("payments service did not register within 30s")
	}
}

// testLogWriter sends the output of the example's services to the test log.
type testLogWriter struct {
	t      *testing.T
	prefix string
}

func (w *testLogWriter) Write(p []byte) (int, error) {
	w.t.Log(w.prefix + strings.TrimRight(string(p), "\n"))
	return len(p), nil
}
