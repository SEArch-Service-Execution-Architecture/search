package broker

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/SEArch-Service-Execution-Architecture/search/contract"
	"github.com/SEArch-Service-Execution-Architecture/search/ent"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/compatibilityresult"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const clientServerGlobalFSA = `
.outputs Client
.state graph
0 Srv ! hi 1
.marking 0
.end

.outputs Srv
.state graph
0 Client ? hi 1
.marking 0
.end
`

func newTestBroker(t *testing.T) *brokerServer {
	t.Helper()
	b := NewBrokerServer(fmt.Sprintf("%s/t.db", t.TempDir()))
	t.Cleanup(b.Stop)
	return b
}

func mustGlobalContract(t *testing.T, fsa, initiator string) contract.GlobalContract {
	t.Helper()
	c, err := contract.ConvertPBGlobalContract(&pb.GlobalContract{
		Contract:      []byte(fsa),
		Format:        pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA,
		InitiatorName: initiator,
	})
	require.NoError(t, err)
	return c
}

// registerTestProvider saves a provider whose contract only differs from others by its message
// name, so each distinct msg yields a distinct provider contract.
func registerTestProvider(t *testing.T, b *brokerServer, msg, providerURL string) *ent.RegisteredProvider {
	t.Helper()
	return registerTestProviderFSA(t, b, fmt.Sprintf(".outputs\n.state graph\n0 Client ? %s 1\n.marking 0\n.end\n", msg), providerURL)
}

func registerTestProviderFSA(t *testing.T, b *brokerServer, fsa, providerURL string) *ent.RegisteredProvider {
	t.Helper()
	c, err := contract.ConvertPBLocalContract(&pb.LocalContract{
		Contract: []byte(fsa),
		Format:   pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA,
	})
	require.NoError(t, err)
	rc, err := b.getOrSaveContract(context.Background(), c)
	require.NoError(t, err)
	u, err := url.Parse(providerURL)
	require.NoError(t, err)
	p, err := b.saveProvider(context.Background(), uuid.New(), u, rc)
	require.NoError(t, err)
	return p
}

// saveRequirementProjection stores the projection of req for participant, as BrokerChannel does.
func saveRequirementProjection(t *testing.T, b *brokerServer, req contract.GlobalContract, participant string) *ent.RegisteredContract {
	t.Helper()
	projection, err := req.GetProjection(participant)
	require.NoError(t, err)
	rc, err := b.getOrSaveContract(context.Background(), projection)
	require.NoError(t, err)
	return rc
}

// waitOrFail runs f and fails the test if it does not return within timeout.
func waitOrFail(t *testing.T, timeout time.Duration, what string, f func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		f()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Fatalf("timed out after %s waiting for %s", timeout, what)
	}
}

func TestGetBestCandidate_MultipleCompatibleProvidersDoNotLeakWorkers(t *testing.T) {
	b := newTestBroker(t)
	req := mustGlobalContract(t, clientServerGlobalFSA, "Client")
	saveRequirementProjection(t, b, req, "Srv")
	registerTestProvider(t, b, "hi", "p1.example.org:1")
	registerTestProvider(t, b, "hello", "p2.example.org:1")
	registerTestProvider(t, b, "hey", "p3.example.org:1")
	b.SetCompatFunc(func(ctx context.Context, req, prov contract.LocalContract) (bool, map[string]string, error) {
		return true, map[string]string{"Client": "Client"}, nil
	})

	result, err := b.getBestCandidate(context.Background(), req, "Srv")

	require.NoError(t, err)
	require.NotNil(t, result)
	// All background checks must finish; previously every compatible provider after the first
	// blocked forever on a full channel, so Stop() hung.
	waitOrFail(t, 5*time.Second, "background compatibility checks", b.compatChecksWaitGroup.Wait)
}

func TestGetBestCandidate_BackgroundChecksSurviveRequestCancellation(t *testing.T) {
	b := newTestBroker(t)
	req := mustGlobalContract(t, clientServerGlobalFSA, "Client")
	reqRC := saveRequirementProjection(t, b, req, "Srv")
	registerTestProvider(t, b, "hi", "p1.example.org:1")
	registerTestProvider(t, b, "hello", "p2.example.org:1")
	registerTestProvider(t, b, "hey", "p3.example.org:1")
	var calls atomic.Int32
	b.SetCompatFunc(func(ctx context.Context, req, prov contract.LocalContract) (bool, map[string]string, error) {
		if calls.Add(1) == 1 {
			return true, map[string]string{"Client": "Client"}, nil
		}
		time.Sleep(200 * time.Millisecond)
		return false, nil, ctx.Err()
	})
	ctx, cancel := context.WithCancel(context.Background())

	_, err := b.getBestCandidate(ctx, req, "Srv")
	cancel() // The RPC that triggered the checks has finished.

	require.NoError(t, err)
	waitOrFail(t, 5*time.Second, "background compatibility checks", b.compatChecksWaitGroup.Wait)
	saved, err := b.dbClient.CompatibilityResult.Query().
		Where(compatibilityresult.RequirementContractID(reqRC.ID)).
		Count(context.Background())
	require.NoError(t, err)
	require.Equal(t, 3, saved, "all compatibility results should be saved even after the request context is cancelled")
}

func TestGetBestCandidate_ChecksEachProviderContractOnce(t *testing.T) {
	b := newTestBroker(t)
	req := mustGlobalContract(t, clientServerGlobalFSA, "Client")
	saveRequirementProjection(t, b, req, "Srv")
	// Two providers share the same contract.
	registerTestProvider(t, b, "hi", "p1.example.org:1")
	registerTestProvider(t, b, "hi", "p2.example.org:1")
	registerTestProvider(t, b, "hello", "p3.example.org:1")
	var mu sync.Mutex
	checked := make(map[string]int)
	b.SetCompatFunc(func(ctx context.Context, req, prov contract.LocalContract) (bool, map[string]string, error) {
		mu.Lock()
		defer mu.Unlock()
		checked[prov.GetContractID()]++
		return false, nil, nil
	})

	_, err := b.getBestCandidate(context.Background(), req, "Srv")
	require.Error(t, err)
	b.compatChecksWaitGroup.Wait()

	require.Len(t, checked, 2)
	for id, n := range checked {
		require.Equal(t, 1, n, "provider contract %s checked %d times", id, n)
	}

	// A second request must reuse the cached (negative) results instead of recomputing them.
	_, err = b.getBestCandidate(context.Background(), req, "Srv")
	require.Error(t, err)
	b.compatChecksWaitGroup.Wait()
	for id, n := range checked {
		require.Equal(t, 1, n, "provider contract %s checked %d times", id, n)
	}
}

func TestGetBestCandidates_ParticipantNotInContract(t *testing.T) {
	b := newTestBroker(t)
	req := mustGlobalContract(t, clientServerGlobalFSA, "Client")

	_, err := b.getBestCandidates(context.Background(), req, []string{"Nope"}, nil)

	require.Equal(t, codes.InvalidArgument, status.Code(err), "got %v", err)
}

func TestGetBestCandidates_ReportsFailedParticipants(t *testing.T) {
	b := newTestBroker(t)
	req := mustGlobalContract(t, clientServerGlobalFSA, "Client")
	saveRequirementProjection(t, b, req, "Srv")

	_, err := b.getBestCandidates(context.Background(), req, []string{"Srv"}, nil)

	st := status.Convert(err)
	require.Equal(t, "CHANNEL_BROKERAGE_FAILED", st.Message())
	require.NotEmpty(t, st.Details())
}

func TestBrokerChannel_UnreachableProviderReturnsError(t *testing.T) {
	b := newTestBroker(t)
	// Nothing listens on port 1, so InitChannel fails.
	registerTestProvider(t, b, "hi", "localhost:1")
	b.SetCompatFunc(func(ctx context.Context, req, prov contract.LocalContract) (bool, map[string]string, error) {
		return true, map[string]string{"Client": "Client"}, nil
	})

	_, err := b.BrokerChannel(context.Background(), &pb.BrokerChannelRequest{
		Contract: &pb.GlobalContract{
			Contract:      []byte(clientServerGlobalFSA),
			Format:        pb.GlobalContractFormat_GLOBAL_CONTRACT_FORMAT_FSA,
			InitiatorName: "Client",
		},
		PresetParticipants: map[string]*pb.RemoteParticipant{
			"Client": {Url: "localhost:1", AppId: uuid.NewString()},
		},
	})

	// Previously this called log.Fatalf and killed the whole broker.
	require.Equal(t, codes.Unavailable, status.Code(err), "got %v", err)
	require.ErrorContains(t, err, "Srv")
}
