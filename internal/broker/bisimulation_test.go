package broker

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/SEArch-Service-Execution-Architecture/search/contract"
	pb "github.com/SEArch-Service-Execution-Architecture/search/gen/go/search/v1"
	"github.com/stretchr/testify/require"
)

// useFakePython points PythonExecutable at a shell script with the given body for the
// duration of the test.
func useFakePython(t *testing.T, body string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "fakepython")
	err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0o755)
	require.NoError(t, err)
	orig := PythonExecutable
	PythonExecutable = path
	t.Cleanup(func() { PythonExecutable = orig })
}

func clientServerContracts(t *testing.T) (contract.LocalContract, contract.LocalContract) {
	t.Helper()
	parse := func(fsa string) contract.LocalContract {
		c, err := contract.ConvertPBLocalContract(&pb.LocalContract{
			Contract: []byte(fsa),
			Format:   pb.LocalContractFormat_LOCAL_CONTRACT_FORMAT_FSA,
		})
		require.NoError(t, err)
		return c
	}
	req := parse(".outputs\n.state graph\n0 Client ? hi 1\n.marking 0\n.end\n")
	prov := parse(".outputs\n.state graph\n0 Client ? hi 1\n.marking 0\n.end\n")
	return req, prov
}

func TestBisimilarityAlgorithmIgnoresStderr(t *testing.T) {
	useFakePython(t, `cat >/dev/null
echo "SomeWarning: this is not JSON" >&2
echo '{"relation": true, "participant_matches": {"Self": "Self", "Client": "Client"}}'`)
	req, prov := clientServerContracts(t)

	ok, mapping, err := BisimilarityAlgorithm(context.Background(), req, prov)

	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, map[string]string{"Client": "Client"}, mapping)
}

func TestBisimilarityAlgorithmReportsStderrOnFailure(t *testing.T) {
	useFakePython(t, `cat >/dev/null
echo "ModuleNotFoundError: No module named 'cfsm_bisimulation'" >&2
exit 1`)
	req, prov := clientServerContracts(t)

	_, _, err := BisimilarityAlgorithm(context.Background(), req, prov)

	require.ErrorContains(t, err, "No module named 'cfsm_bisimulation'")
}

func TestBisimilarityAlgorithmRespectsContext(t *testing.T) {
	useFakePython(t, `sleep 10`)
	req, prov := clientServerContracts(t)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, _, err := BisimilarityAlgorithm(ctx, req, prov)

	require.Error(t, err)
	require.True(t, errors.Is(err, context.DeadlineExceeded), "expected deadline exceeded, got %v", err)
	require.Less(t, time.Since(start), 5*time.Second)
}
