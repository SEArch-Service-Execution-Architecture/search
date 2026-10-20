package broker

import (
	"context"
	"testing"

	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// The participant mapping sent to a provider in InitChannel must be keyed by the participant
// names in the *provider's* contract, using the mapping computed by the real bisimulation check.
func TestGetParticipantMappingWithBisimulationResult(t *testing.T) {
	requirePythonBisimulation(t)
	b := newTestBroker(t)
	b.SetCompatFunc(BisimilarityAlgorithm)

	const global = `
.outputs ClientApp
.state graph
q0 PPS ! CardDetailsWithTotalAmount q1
q1 PPS ? PaymentNonce q2
.marking q0
.end

.outputs Srv
.state graph
q0 PPS ! RequestChargeWithNonce q1
q1 PPS ? ChargeOK q2
q1 PPS ? ChargeFail q3
.marking q0
.end

.outputs PPS
.state graph
q0 ClientApp ? CardDetailsWithTotalAmount q1
q1 ClientApp ! PaymentNonce q2
q2 Srv ? RequestChargeWithNonce q3
q3 Srv ! ChargeOK q4
q3 Srv ! ChargeFail q5
.marking q0
.end
`
	req := mustGlobalContract(t, global, "ClientApp")
	saveRequirementProjection(t, b, req, "PPS")
	// Same behaviour as the PPS projection, but with different participant names.
	provider := registerTestProviderFSA(t, b, `.outputs PaymentService
.state graph
q0 cliente ? CardDetailsWithTotalAmount q1
q1 cliente ! PaymentNonce q2
q2 backend ? RequestChargeWithNonce q3
q3 backend ! ChargeOK q4
q3 backend ! ChargeFail q5
.marking q0
.end
`, "pps.example.org:1")
	candidate, err := b.getBestCandidate(context.Background(), req, "PPS")
	require.NoError(t, err)
	require.Equal(t, provider.ID, candidate.ID)
	b.compatChecksWaitGroup.Wait()

	clientApp := &pb.RemoteParticipant{Url: "client.example.org:1", AppId: uuid.NewString()}
	srv := &pb.RemoteParticipant{Url: "srv.example.org:1", AppId: uuid.NewString()}
	pps := &pb.RemoteParticipant{Url: "pps.example.org:1", AppId: provider.ID.String()}
	initiatorMapping := map[string]*pb.RemoteParticipant{"ClientApp": clientApp, "Srv": srv, "PPS": pps}

	got, err := b.getParticipantMapping(req, initiatorMapping, "PPS", candidate)

	require.NoError(t, err)
	require.Equal(t, map[string]*pb.RemoteParticipant{"cliente": clientApp, "backend": srv}, got)
}
