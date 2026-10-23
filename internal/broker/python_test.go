package broker

import (
	"testing"

	"github.com/SEArch-Service-Execution-Architecture/search/internal/testenv"
)

// requirePythonBisimulation skips the test when the cfsm_bisimulation Python package is not
// importable. Set SEARCH_REQUIRE_PYTHON=1 (as CI does) to fail instead of skipping.
func requirePythonBisimulation(t *testing.T) {
	t.Helper()
	testenv.RequirePython(t, PythonExecutable, "cfsm_bisimulation")
}
