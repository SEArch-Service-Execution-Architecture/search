package broker

import (
	"os"
	"os/exec"
	"testing"
)

// requirePythonBisimulation skips the test when the cfsm_bisimulation Python package is not
// importable. Set SEARCH_REQUIRE_PYTHON=1 (as CI does) to fail instead of skipping.
func requirePythonBisimulation(t *testing.T) {
	t.Helper()
	err := exec.Command("python", "-c", "import cfsm_bisimulation").Run()
	if err == nil {
		return
	}
	if os.Getenv("SEARCH_REQUIRE_PYTHON") != "" {
		t.Fatalf("cfsm_bisimulation Python package is required but not importable: %v", err)
	}
	t.Skipf("skipping: cfsm_bisimulation Python package not importable: %v", err)
}
