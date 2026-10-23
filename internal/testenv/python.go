// Package testenv has helpers for tests that depend on the environment they run in.
package testenv

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// RequirePython skips the test when python cannot import all of the given modules. Set
// SEARCH_REQUIRE_PYTHON=1 (as CI does) to fail instead of skipping.
func RequirePython(t testing.TB, python string, modules ...string) {
	t.Helper()
	out, err := exec.Command(python, "-c", "import "+strings.Join(modules, ", ")).CombinedOutput()
	if err == nil {
		return
	}
	if os.Getenv("SEARCH_REQUIRE_PYTHON") != "" {
		t.Fatalf("Python modules %v are required but not importable: %v\n%s", modules, err, out)
	}
	t.Skipf("skipping: Python modules %v not importable: %v", modules, err)
}
