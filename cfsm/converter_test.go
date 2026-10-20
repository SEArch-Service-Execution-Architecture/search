package cfsm

import (
	"bytes"
	"reflect"
	"testing"
)

func TestConvertCFSMToPythonBisimulationFormat(t *testing.T) {
	// Create a sample CFSM for testing
	contract, err := ParseSingleCFSMFSA(bytes.NewReader([]byte(`
	.outputs self
	.state graph
	q0 OurServer ! login() q1
	q1 OurServer ? accept() q2
	.marking q0
	.end
	`)))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := []byte(`{{.MachineName}} = CommunicatingFiniteStateMachine(["Self","Courserver"])

{{.MachineName}}.add_states('0')
{{.MachineName}}.add_states('1')
{{.MachineName}}.add_states('2')

{{.MachineName}}.set_as_initial('0')

{{.MachineName}}.add_transition_between('0', '1', 'Self Courserver ! message0()')
{{.MachineName}}.add_transition_between('1', '2', 'Self Courserver ? message1()')

`)

	output, _, _, _, err := ConvertCFSMToPythonBisimulationFormat(contract)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Output does not match expected value.\nExpected: %s\nGot: %s", expectedOutput, output)
	}
}

func TestConvertCFSMToPythonBisimulationFormatUniqueParticipantNames(t *testing.T) {
	// "a-b", "ab" and "a_b" all sanitize to "Cab", and "Cab" is already a valid name that is
	// kept as-is. Each participant must still get its own translated name.
	contract, err := ParseSingleCFSMFSA(bytes.NewReader([]byte(`
	.outputs self
	.state graph
	q0 a-b ! m1 q1
	q1 ab ! m2 q2
	q2 a_b ! m3 q3
	q3 Cab ! m4 q4
	.marking q0
	.end
	`)))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, translations, _, selfname, err := ConvertCFSMToPythonBisimulationFormat(contract)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	originals := []string{"a-b", "ab", "a_b", "Cab"}
	seen := make(map[string]string)
	for _, original := range originals {
		translated, ok := translations.Get(original)
		if !ok {
			t.Fatalf("no translation for participant %q", original)
		}
		if other, dup := seen[translated]; dup {
			t.Errorf("participants %q and %q both translated to %q", other, original, translated)
		}
		if translated == selfname {
			t.Errorf("participant %q translated to the self name %q", original, selfname)
		}
		seen[translated] = original
		back, ok := translations.GetInverse(translated)
		if !ok || back != original {
			t.Errorf("translation of %q does not round-trip: %q -> %q", original, translated, back)
		}
	}
	if translations.Size() != len(originals) {
		t.Errorf("expected %d translations, got %d", len(originals), translations.Size())
	}
}
