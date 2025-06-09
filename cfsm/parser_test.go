package cfsm

import (
	"fmt"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldFSAParse(t *testing.T) {
	const exampleFSAContent = `
.outputs A
.state graph
q0 1 ! hello q1
q0 1 ! world q1
.marking q0
.end

.outputs
.state graph
q0 0 ? hello q1
q0 0 ? world q1
.marking q0
.end`
	fs := fstest.MapFS{
		"example.fsa": {Data: []byte(exampleFSAContent)},
	}

	exampleFile, err := fs.Open("example.fsa")
	if err != nil {
		t.Error("Failed reading example.fsa.")
	}
	defer exampleFile.Close()

	sys, err := ParseSystemCFSMsFSA(exampleFile)
	if err != nil {
		t.Errorf("Error parsing FSA: %s", err)
	}

	if len(sys.CFSMs) != 2 {
		t.Errorf("Parsed FSA has %d CFSMs. Expected 2.", len(sys.CFSMs))
	}
	firstCFSM := sys.CFSMs[0]
	if firstCFSM.Name != "A" {
		t.Errorf("Expected fist CFSM to have Name 'A' but instead had %s", firstCFSM.Name)
	}
	if len(firstCFSM.States()) != 2 {
		t.Errorf("Expected 2 states in first CFSM. Found %d", len(firstCFSM.States()))
	}
	if firstCFSM.Start.Label != "q0" {
		t.Errorf("Expected start state of first CFSM to be 'q0'. Found %s instead", firstCFSM.Start.Label)
	}

}

func TestPingPongFSAParser(t *testing.T) {
	const exampleFSAContent = `
.outputs Ping
.state graph
0 1 ! ping 5
2 1 ? bye 1
3 1 ! bye 2
3 1 ! finished 2
4 1 ! *<1 0
4 1 ! >*1 3
5 1 ? pong 4
.marking 0
.end

.outputs Pong
.state graph
0 0 ? ping 5
2 0 ! bye 1
3 0 ? bye 2
3 0 ? finished 2
4 0 ? *<1 0
4 0 ? >*1 3
5 0 ! pong 4
.marking 0
.end
`
	fs := fstest.MapFS{
		"example.fsa": {Data: []byte(exampleFSAContent)},
	}

	exampleFile, err := fs.Open("example.fsa")
	if err != nil {
		t.Error("Failed reading example.fsa.")
	}
	defer exampleFile.Close()

	sys, err := ParseSystemCFSMsFSA(exampleFile)
	if err != nil {
		t.Errorf("Error parsing FSA: %s", err)
	}

	if len(sys.CFSMs) != 2 {
		t.Errorf("Parsed FSA has %d CFSMs. Expected 2.", len(sys.CFSMs))
	}

	require.ElementsMatch(t, []string{"Ping", "Pong"}, sys.GetAllMachineNames())

	firstCFSM := sys.CFSMs[0]
	if firstCFSM.Name != "Ping" {
		t.Errorf("Expected first CFSM to have Name 'Ping' but instead had %s", firstCFSM.Name)
	}
	if len(firstCFSM.States()) != 6 {
		t.Errorf("Expected 6 states in first CFSM. Found %d", len(firstCFSM.States()))
	}

}

func TestPingPongFSALocalContractParser(t *testing.T) {
	const exampleFSAContent = `
.outputs Ping
.state graph
0 1 ! ping 5
2 1 ? bye 1
3 1 ! bye 2
3 1 ! finished 2
4 1 ! *<1 0
4 1 ! >*1 3
5 1 ? pong 4
.marking 0
.end
`
	fs := fstest.MapFS{
		"example.fsa": {Data: []byte(exampleFSAContent)},
	}

	exampleFile, err := fs.Open("example.fsa")
	if err != nil {
		t.Error("Failed reading example.fsa.")
	}
	defer exampleFile.Close()

	machine, err := ParseSingleCFSMFSA(exampleFile)
	if err != nil {
		t.Errorf("Error parsing FSA: %s", err)
	}

	if machine.Name != "Ping" {
		t.Errorf("Expected CFSM to have Name 'Ping' but instead had %s", machine.Name)
	}
	if len(machine.States()) != 6 {
		t.Errorf("Expected 6 states in first CFSM. Found %d", len(machine.States()))
	}

}

func TestCFSMSerialization(t *testing.T) {
	const pingPongFSA = `
.outputs Ping
.state graph
0 Pong ! ping 1
1 Pong ? pong 0
0 Pong ! bye 2
0 Pong ! finished 3
2 Pong ? bye 3
.marking 0
.end

.outputs Pong
.state graph
0 Ping ? ping 1
1 Ping ! pong 0
0 Ping ? bye 2
2 Ping ! bye 3
0 Ping ? finished 3
.marking 0
.end
`
	const cfsmStringPingPong = `
-- Machine #0
.outputs Ping
.state graph
q00 Pong ! ping q01
q00 Pong ! bye q02
q00 Pong ! finished q03
q01 Pong ? pong q00
q02 Pong ? bye q03
.marking q00
.end

-- Machine #1
.outputs Pong
.state graph
q10 Ping ? ping q11
q10 Ping ? bye q12
q10 Ping ? finished q13
q11 Ping ! pong q10
q12 Ping ! bye q13
.marking q10
.end
`
	pingPongSystem, err := ParseSystemCFSMsFSA(strings.NewReader(pingPongFSA))
	if err != nil {
		t.Errorf("Error parsing FSA: %s", err)
	}

	require.Equal(t, cfsmStringPingPong, pingPongSystem.String())
}

func TestSingleCFSMSerialization(t *testing.T) {
	const cfsmPongFSA = `.outputs Pong
	.state graph
	0 Ping ? ping 1
	1 Ping ! pong 0
	0 Ping ? bye 2
	2 Ping ! bye 3
	0 Ping ? finished 3
	.marking 0
	.end`
	const normalizedPongString = `
-- Machine #0
.outputs Pong
.state graph
q00 Ping ? ping q01
q00 Ping ? bye q02
q00 Ping ? finished q03
q01 Ping ! pong q00
q02 Ping ! bye q03
.marking q00
.end
`
	pongCFSM, err := ParseSingleCFSMFSA(strings.NewReader(cfsmPongFSA))
	if err != nil {
		t.Errorf("Error parsing FSA: %s", err)
	}

	require.Equal(t, normalizedPongString, pongCFSM.String())
}

func TestParseFSATransitionStructure(t *testing.T) {
	const exampleFSAContent = `
.outputs Alice
.state graph
start Bob ! hello waiting
waiting Bob ? reply end
start Bob ! goodbye end
.marking start
.end

.outputs Bob
.state graph
idle Alice ? hello responding
responding Alice ! reply idle
idle Alice ? goodbye idle
.marking idle
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(exampleFSAContent))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	// Test Alice CFSM
	alice := sys.CFSMs[0]
	require.Equal(t, "Alice", alice.Name)
	require.Equal(t, "start", alice.Start.Label)

	states := alice.States()
	require.Len(t, states, 3)

	// Find states by label
	var startState, waitingState, endState *State
	for _, state := range states {
		switch state.Label {
		case "start":
			startState = state
		case "waiting":
			waitingState = state
		case "end":
			endState = state
		}
	}
	require.NotNil(t, startState)
	require.NotNil(t, waitingState)
	require.NotNil(t, endState)

	// Check start state transitions
	startTransitions := startState.Transitions()
	require.Len(t, startTransitions, 2)

	// Verify hello and goodbye transitions
	var helloTransition, goodbyeTransition Transition
	for _, tr := range startTransitions {
		if tr.Message() == "hello" {
			helloTransition = tr
		} else if tr.Message() == "goodbye" {
			goodbyeTransition = tr
		}
	}
	require.NotNil(t, helloTransition)
	require.NotNil(t, goodbyeTransition)

	// Check hello transition details
	require.True(t, helloTransition.IsSend())
	require.Equal(t, "hello", helloTransition.Message())
	require.Equal(t, "Bob", helloTransition.NameOfOtherCFSM())
	require.Equal(t, waitingState, helloTransition.State())

	// Check goodbye transition details
	require.True(t, goodbyeTransition.IsSend())
	require.Equal(t, "goodbye", goodbyeTransition.Message())
	require.Equal(t, "Bob", goodbyeTransition.NameOfOtherCFSM())
	require.Equal(t, endState, goodbyeTransition.State())

	// Check waiting state has one receive transition
	waitingTransitions := waitingState.Transitions()
	require.Len(t, waitingTransitions, 1)
	replyRecv := waitingTransitions[0]
	require.False(t, replyRecv.IsSend())
	require.Equal(t, "reply", replyRecv.Message())
	require.Equal(t, "Bob", replyRecv.NameOfOtherCFSM())
	require.Equal(t, endState, replyRecv.State())

	// Test Bob CFSM
	bob := sys.CFSMs[1]
	require.Equal(t, "Bob", bob.Name)
	require.Equal(t, "idle", bob.Start.Label)

	bobStates := bob.States()
	require.Len(t, bobStates, 2)

	var idleState, respondingState *State
	for _, state := range bobStates {
		switch state.Label {
		case "idle":
			idleState = state
		case "responding":
			respondingState = state
		}
	}
	require.NotNil(t, idleState)
	require.NotNil(t, respondingState)

	// Check idle state transitions
	idleTransitions := idleState.Transitions()
	require.Len(t, idleTransitions, 2)

	var helloRecv, goodbyeRecv Transition
	for _, tr := range idleTransitions {
		if tr.Message() == "hello" {
			helloRecv = tr
		} else if tr.Message() == "goodbye" {
			goodbyeRecv = tr
		}
	}
	require.NotNil(t, helloRecv)
	require.NotNil(t, goodbyeRecv)

	// Check hello receive transition
	require.False(t, helloRecv.IsSend())
	require.Equal(t, "hello", helloRecv.Message())
	require.Equal(t, "Alice", helloRecv.NameOfOtherCFSM())
	require.Equal(t, respondingState, helloRecv.State())

	// Check goodbye receive transition (self-loop)
	require.False(t, goodbyeRecv.IsSend())
	require.Equal(t, "goodbye", goodbyeRecv.Message())
	require.Equal(t, "Alice", goodbyeRecv.NameOfOtherCFSM())
	require.Equal(t, idleState, goodbyeRecv.State())

	// Check responding state has one send transition
	respondingTransitions := respondingState.Transitions()
	require.Len(t, respondingTransitions, 1)
	replySend := respondingTransitions[0]
	require.True(t, replySend.IsSend())
	require.Equal(t, "reply", replySend.Message())
	require.Equal(t, "Alice", replySend.NameOfOtherCFSM())
	require.Equal(t, idleState, replySend.State())
}

func TestParseFSAComplexTransitions(t *testing.T) {
	const complexFSAContent = `
.outputs Server
.state graph
s0 Client ! welcome s1
s1 Client ? login s2
s1 Client ? register s3
s2 Client ! loginSuccess s4
s2 Client ! loginFailed s0
s3 Client ! registerSuccess s1
s3 Client ! registerFailed s0
s4 Client ? logout s0
s4 Client ? request s5
s5 Client ! response s4
.marking s0
.end

.outputs Client
.state graph
c0 Server ? welcome c1
c1 Server ! login c2
c1 Server ! register c3
c2 Server ? loginSuccess c4
c2 Server ? loginFailed c0
c3 Server ? registerSuccess c1
c3 Server ? registerFailed c0
c4 Server ! logout c0
c4 Server ! request c5
c5 Server ? response c4
.marking c0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(complexFSAContent))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	server := sys.CFSMs[0]
	client := sys.CFSMs[1]

	require.Equal(t, "Server", server.Name)
	require.Equal(t, "Client", client.Name)

	// Test Server states
	serverStates := server.States()
	require.Len(t, serverStates, 6)

	// Find specific states
	var s0, s1, s2 *State
	for _, state := range serverStates {
		switch state.Label {
		case "s0":
			s0 = state
		case "s1":
			s1 = state
		case "s2":
			s2 = state
		}
	}
	require.NotNil(t, s0)
	require.NotNil(t, s1)
	require.NotNil(t, s2)

	// Test s0 transitions (should have 1 send)
	s0Transitions := s0.Transitions()
	require.Len(t, s0Transitions, 1)
	welcomeTr := s0Transitions[0]
	require.True(t, welcomeTr.IsSend())
	require.Equal(t, "welcome", welcomeTr.Message())
	require.Equal(t, "Client", welcomeTr.NameOfOtherCFSM())

	// Test s1 transitions (should have 2 receives)
	s1Transitions := s1.Transitions()
	require.Len(t, s1Transitions, 2)
	for _, tr := range s1Transitions {
		require.False(t, tr.IsSend())
		require.Equal(t, "Client", tr.NameOfOtherCFSM())
		require.Contains(t, []string{"login", "register"}, tr.Message())
	}

	// Test s2 transitions (should have 2 sends)
	s2Transitions := s2.Transitions()
	require.Len(t, s2Transitions, 2)
	for _, tr := range s2Transitions {
		require.True(t, tr.IsSend())
		require.Equal(t, "Client", tr.NameOfOtherCFSM())
		require.Contains(t, []string{"loginSuccess", "loginFailed"}, tr.Message())
	}

	// Verify state connections are correct
	require.Equal(t, server.Start, s0)

	// Count total transitions across all states
	totalServerTransitions := 0
	for _, state := range serverStates {
		totalServerTransitions += len(state.Transitions())
	}
	require.Equal(t, 10, totalServerTransitions)

	// Test Client states
	clientStates := client.States()
	require.Len(t, clientStates, 6)

	totalClientTransitions := 0
	for _, state := range clientStates {
		totalClientTransitions += len(state.Transitions())
	}
	require.Equal(t, 10, totalClientTransitions)
}

func TestParseFSASelfLoops(t *testing.T) {
	const selfLoopFSAContent = `
.outputs Worker
.state graph
idle Boss ! ready working
working Boss ? task working
working Boss ? stop idle
.marking idle
.end

.outputs Boss
.state graph
waiting Worker ? ready assigning
assigning Worker ! task waiting
assigning Worker ! stop waiting
.marking waiting
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(selfLoopFSAContent))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	worker := sys.CFSMs[0]
	require.Equal(t, "Worker", worker.Name)

	workerStates := worker.States()
	require.Len(t, workerStates, 2)

	var idleState, workingState *State
	for _, state := range workerStates {
		switch state.Label {
		case "idle":
			idleState = state
		case "working":
			workingState = state
		}
	}
	require.NotNil(t, idleState)
	require.NotNil(t, workingState)

	// Test working state has a self-loop
	workingTransitions := workingState.Transitions()
	require.Len(t, workingTransitions, 2)

	var taskRecv, stopRecv Transition
	for _, tr := range workingTransitions {
		if tr.Message() == "task" {
			taskRecv = tr
		} else if tr.Message() == "stop" {
			stopRecv = tr
		}
	}
	require.NotNil(t, taskRecv)
	require.NotNil(t, stopRecv)

	// Verify self-loop
	require.Equal(t, workingState, taskRecv.State())
	require.Equal(t, idleState, stopRecv.State())
}

func TestParseFSANumericStateNames(t *testing.T) {
	const numericStatesFSAContent = `
.outputs Machine0
.state graph
0 Machine1 ! start 1
1 Machine1 ? ack 2
2 Machine1 ! data 3
3 Machine1 ? done 0
.marking 0
.end

.outputs Machine1
.state graph
0 Machine0 ? start 1
1 Machine0 ! ack 2
2 Machine0 ? data 3
3 Machine0 ! done 0
.marking 0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(numericStatesFSAContent))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	m0 := sys.CFSMs[0]
	require.Equal(t, "Machine0", m0.Name)
	require.Equal(t, "0", m0.Start.Label)

	m0States := m0.States()
	require.Len(t, m0States, 4)

	// Verify all states have correct labels
	expectedLabels := []string{"0", "1", "2", "3"}
	actualLabels := make([]string, len(m0States))
	for i, state := range m0States {
		actualLabels[i] = state.Label
	}
	require.ElementsMatch(t, expectedLabels, actualLabels)

	// Each state should have exactly one transition
	for _, state := range m0States {
		require.Len(t, state.Transitions(), 1)
	}
}

func TestParseFSAErrorCases(t *testing.T) {
	testCases := []struct {
		name        string
		fsaContent  string
		expectError bool
		errorMsg    string
	}{
		{
			name: "missing .outputs",
			fsaContent: `
.state graph
q0 m1 ! msg q1
.marking q0
.end`,
			expectError: true,
			errorMsg:    "Expected .outputs",
		},
		{
			name: "missing .state graph",
			fsaContent: `
.outputs Machine
q0 m1 ! msg q1
.marking q0
.end`,
			expectError: true,
			errorMsg:    "expected '.state graph'",
		},
		{
			name: "missing .marking",
			fsaContent: `
.outputs Machine
.state graph
q0 m1 ! msg q1
.end`,
			expectError: true,
			errorMsg:    "expected transition, got .end",
		},
		{
			name: "missing .end",
			fsaContent: `
.outputs Machine
.state graph
q0 m1 ! msg q1
.marking q0`,
			expectError: true,
			errorMsg:    "unexpected EOF",
		},
		{
			name: "invalid transition format",
			fsaContent: `
.outputs Machine
.state graph
q0 m1 invalid q1
.marking q0
.end`,
			expectError: true,
			errorMsg:    "expected transition",
		},
		{
			name: "invalid action symbol",
			fsaContent: `
.outputs Machine
.state graph
q0 m1 @ msg q1
.marking q0
.end`,
			expectError: true,
			errorMsg:    "expected transition, got q0 m1 @ msg q1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseSystemCFSMsFSA(strings.NewReader(tc.fsaContent))
			if tc.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParseFSANamedVsNumberedMachines(t *testing.T) {
	const namedMachinesFSA = `
.outputs Alice
.state graph
q0 Bob ! hello q1
.marking q0
.end

.outputs Bob
.state graph
q0 Alice ? hello q1
.marking q0
.end`

	const numberedMachinesFSA = `
.outputs 0
.state graph
q0 1 ! hello q1
.marking q0
.end

.outputs 1
.state graph
q0 0 ? hello q1
.marking q0
.end`

	// Test named machines
	namedSys, err := ParseSystemCFSMsFSA(strings.NewReader(namedMachinesFSA))
	require.NoError(t, err)
	require.Len(t, namedSys.CFSMs, 2)
	require.Equal(t, "Alice", namedSys.CFSMs[0].Name)
	require.Equal(t, "Bob", namedSys.CFSMs[1].Name)

	// Test numbered machines
	numberedSys, err := ParseSystemCFSMsFSA(strings.NewReader(numberedMachinesFSA))
	require.NoError(t, err)
	require.Len(t, numberedSys.CFSMs, 2)
	require.Equal(t, "0", numberedSys.CFSMs[0].Name)
	require.Equal(t, "1", numberedSys.CFSMs[1].Name)

	// Verify transitions reference correct machines
	aliceTransition := namedSys.CFSMs[0].Start.Transitions()[0]
	require.Equal(t, "Bob", aliceTransition.NameOfOtherCFSM())

	zeroTransition := numberedSys.CFSMs[0].Start.Transitions()[0]
	require.Equal(t, "1", zeroTransition.NameOfOtherCFSM())
}

func TestParseFSAUnnamedMachines(t *testing.T) {
	const unnamedMachinesFSA = `
.outputs
.state graph
q0 1 ! hello q1
.marking q0
.end

.outputs
.state graph
q0 0 ? hello q1
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(unnamedMachinesFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	// Unnamed machines should get default numeric names
	require.Equal(t, "0", sys.CFSMs[0].Name)
	require.Equal(t, "1", sys.CFSMs[1].Name)

	// Verify transitions work correctly
	firstTransition := sys.CFSMs[0].Start.Transitions()[0]
	require.Equal(t, "1", firstTransition.NameOfOtherCFSM())
}

func TestParseFSACommentsAndWhitespace(t *testing.T) {
	const fsaWithCommentsAndWhitespace = `
-- This is a comment
.outputs Client

-- Another comment
.state graph

-- Comment before transition
q0 Server ! request q1

-- Multiple comments
-- in a row
q1 Server ? response q0

-- Comment before marking
.marking q0

-- Final comment
.end

-- Machine 2 with whitespace
.outputs   Server

.state graph

q0 Client ? request q1

q1 Client ! response q0


.marking q0
.end

-- End comment`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(fsaWithCommentsAndWhitespace))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	client := sys.CFSMs[0]
	server := sys.CFSMs[1]

	require.Equal(t, "Client", client.Name)
	require.Equal(t, "Server", server.Name)

	// Verify transitions work despite comments and extra whitespace
	clientTransitions := client.Start.Transitions()
	require.Len(t, clientTransitions, 1)
	require.Equal(t, "request", clientTransitions[0].Message())
	require.Equal(t, "Server", clientTransitions[0].NameOfOtherCFSM())
}

func TestParseFSASingleCFSMWithNamedTargets(t *testing.T) {
	const singleCFSMContent = `
.outputs LocalMachine
.state graph
start RemoteService ! connect waiting
waiting RemoteService ? connected active
active RemoteService ! disconnect start
active AnotherService ! ping active
.marking start
.end`

	cfsm, err := ParseSingleCFSMFSA(strings.NewReader(singleCFSMContent))
	require.NoError(t, err)
	require.Equal(t, "LocalMachine", cfsm.Name)
	require.Equal(t, "start", cfsm.Start.Label)

	states := cfsm.States()
	require.Len(t, states, 3)

	// Find start state and verify its transitions
	var startState *State
	for _, state := range states {
		if state.Label == "start" {
			startState = state
			break
		}
	}
	require.NotNil(t, startState)

	startTransitions := startState.Transitions()
	require.Len(t, startTransitions, 1)

	connectTr := startTransitions[0]
	require.True(t, connectTr.IsSend())
	require.Equal(t, "connect", connectTr.Message())
	require.Equal(t, "RemoteService", connectTr.NameOfOtherCFSM())

	// Find active state and verify its transitions to different services
	var activeState *State
	for _, state := range states {
		if state.Label == "active" {
			activeState = state
			break
		}
	}
	require.NotNil(t, activeState)

	activeTransitions := activeState.Transitions()
	require.Len(t, activeTransitions, 2)

	// Count transitions by target service
	serviceCount := make(map[string]int)
	for _, tr := range activeTransitions {
		serviceCount[tr.NameOfOtherCFSM()]++
	}
	require.Equal(t, 1, serviceCount["RemoteService"])
	require.Equal(t, 1, serviceCount["AnotherService"])
}

func TestParseFSAMultipleMessagesAndStates(t *testing.T) {
	const multipleMessagesFSA = `
.outputs Producer
.state graph
start Consumer ! msg1 state1
start Consumer ! msg2 state2
state1 Consumer ? ack1 end
state2 Consumer ? ack2 end
state1 Consumer ! final end
state2 Consumer ! final end
.marking start
.end

.outputs Consumer
.state graph
idle Producer ? msg1 processing1
idle Producer ? msg2 processing2
processing1 Producer ! ack1 done
processing2 Producer ! ack2 done
processing1 Producer ? final idle
processing2 Producer ? final idle
.marking idle
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(multipleMessagesFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	producer := sys.CFSMs[0]
	consumer := sys.CFSMs[1]

	// Test Producer structure
	require.Equal(t, "Producer", producer.Name)
	prodStates := producer.States()
	require.Len(t, prodStates, 4) // start, state1, state2, end

	// Find start state
	startState := producer.Start
	require.Equal(t, "start", startState.Label)

	// Start state should have 2 send transitions
	startTransitions := startState.Transitions()
	require.Len(t, startTransitions, 2)

	// Verify both messages from start state
	messages := make(map[string]bool)
	for _, tr := range startTransitions {
		require.True(t, tr.IsSend())
		require.Equal(t, "Consumer", tr.NameOfOtherCFSM())
		messages[tr.Message()] = true
	}
	require.True(t, messages["msg1"])
	require.True(t, messages["msg2"])

	// Test Consumer structure
	require.Equal(t, "Consumer", consumer.Name)
	consumerStates := consumer.States()
	require.Len(t, consumerStates, 4) // idle, processing1, processing2, done

	// Verify idle state has 2 receive transitions
	idleState := consumer.Start
	require.Equal(t, "idle", idleState.Label)
	idleTransitions := idleState.Transitions()
	require.Len(t, idleTransitions, 2)

	for _, tr := range idleTransitions {
		require.False(t, tr.IsSend())
		require.Equal(t, "Producer", tr.NameOfOtherCFSM())
		require.Contains(t, []string{"msg1", "msg2"}, tr.Message())
	}
}

func TestParseFSAStateReordering(t *testing.T) {
	// Test that states are created correctly even when referenced in different orders
	const reorderedStatesFSA = `
.outputs Machine
.state graph
middle Other ! forward end
start Other ! begin middle
end Other ? done start
.marking start
.end

.outputs Other
.state graph
init Machine ? begin active
init Machine ? forward respond
active Machine ! forward init
respond Machine ! done init
.marking init
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(reorderedStatesFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	machine := sys.CFSMs[0]
	require.Equal(t, "Machine", machine.Name)
	require.Equal(t, "start", machine.Start.Label)

	states := machine.States()
	require.Len(t, states, 3)

	// Verify all expected states exist
	stateLabels := make(map[string]*State)
	for _, state := range states {
		stateLabels[state.Label] = state
	}
	require.Contains(t, stateLabels, "start")
	require.Contains(t, stateLabels, "middle")
	require.Contains(t, stateLabels, "end")

	// Test transition connectivity
	startState := stateLabels["start"]
	middleState := stateLabels["middle"]
	endState := stateLabels["end"]

	// start -> middle
	startTransitions := startState.Transitions()
	require.Len(t, startTransitions, 1)
	require.Equal(t, middleState, startTransitions[0].State())

	// middle -> end
	middleTransitions := middleState.Transitions()
	require.Len(t, middleTransitions, 1)
	require.Equal(t, endState, middleTransitions[0].State())

	// end -> start (cycle)
	endTransitions := endState.Transitions()
	require.Len(t, endTransitions, 1)
	require.Equal(t, startState, endTransitions[0].State())
}

func TestParseFSASpecialCharactersInMessages(t *testing.T) {
	const specialCharsFSA = `
.outputs Sender
.state graph
q0 Receiver ! message_with_underscores q1
q1 Receiver ! message-with-dashes q2
q2 Receiver ! message123 q3
q3 Receiver ! *special* q0
.marking q0
.end

.outputs Receiver
.state graph
q0 Sender ? message_with_underscores q1
q1 Sender ? message-with-dashes q2
q2 Sender ? message123 q3
q3 Sender ? *special* q0
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(specialCharsFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	sender := sys.CFSMs[0]
	require.Equal(t, "Sender", sender.Name)

	// Collect all messages from sender
	messages := make([]string, 0)
	for _, state := range sender.States() {
		for _, tr := range state.Transitions() {
			messages = append(messages, tr.Message())
		}
	}

	expectedMessages := []string{"message_with_underscores", "message-with-dashes", "message123", "*special*"}
	require.ElementsMatch(t, expectedMessages, messages)
}

func TestParseFSAInconsistentMachineNumbering(t *testing.T) {
	// Test error when machine numbers don't match expected sequence
	const inconsistentNumberingFSA = `
.outputs 0
.state graph
q0 1 ! msg q1
.marking q0
.end

.outputs 2
.state graph
q0 0 ? msg q1
.marking q0
.end`

	_, err := ParseSystemCFSMsFSA(strings.NewReader(inconsistentNumberingFSA))
	require.Error(t, err)
	require.Contains(t, err.Error(), "machine number 1 is named 2")
}

func TestParseFSANonExistentCFSMReference(t *testing.T) {
	// Test error when referencing a non-existent CFSM
	const nonExistentCFSMFSA = `
.outputs Machine0
.state graph
q0 NonExistentMachine ! msg q1
.marking q0
.end

.outputs Machine1
.state graph
q0 Machine0 ? msg q1
.marking q0
.end`

	_, err := ParseSystemCFSMsFSA(strings.NewReader(nonExistentCFSMFSA))
	require.Error(t, err)
	require.Contains(t, err.Error(), "non existant CFSM referenced")
}

func TestParseFSAEmptyStatesAndTransitions(t *testing.T) {
	const emptyStateFSA = `
.outputs EmptyMachine
.state graph
.marking q0
.end

.outputs SingleTransition
.state graph
q0 EmptyMachine ! ping q1
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(emptyStateFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	emptyMachine := sys.CFSMs[0]
	require.Equal(t, "EmptyMachine", emptyMachine.Name)
	require.Equal(t, "q0", emptyMachine.Start.Label)

	// EmptyMachine should have one state (the start state) with no transitions
	emptyStates := emptyMachine.States()
	require.Len(t, emptyStates, 1)
	require.Len(t, emptyStates[0].Transitions(), 0)

	singleTransition := sys.CFSMs[1]
	require.Equal(t, "SingleTransition", singleTransition.Name)

	// SingleTransition should have transitions
	singleStates := singleTransition.States()
	require.Len(t, singleStates, 2)
	require.Len(t, singleStates[0].Transitions(), 1)
}

func TestParseFSATransitionOrderPreservation(t *testing.T) {
	const orderedTransitionsFSA = `
.outputs Sequencer
.state graph
start Other ! first middle
start Other ! second end
start Other ! third start
.marking start
.end

.outputs Other
.state graph
q0 Sequencer ? first q1
q0 Sequencer ? second q2
q0 Sequencer ? third q3
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(orderedTransitionsFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	sequencer := sys.CFSMs[0]
	startState := sequencer.Start
	transitions := startState.Transitions()
	require.Len(t, transitions, 3)

	// Verify that transitions maintain their definition order
	expectedOrder := []string{"first", "second", "third"}
	actualOrder := make([]string, len(transitions))
	for i, tr := range transitions {
		actualOrder[i] = tr.Message()
	}
	require.Equal(t, expectedOrder, actualOrder)
}

func TestParseFSALongTransitionChains(t *testing.T) {
	const longChainFSA = `
.outputs ChainMachine
.state graph
s0 Other ! msg0 s1
s1 Other ! msg1 s2
s2 Other ! msg2 s3
s3 Other ! msg3 s4
s4 Other ! msg4 s5
s5 Other ! msg5 s0
.marking s0
.end

.outputs Other
.state graph
q0 ChainMachine ? msg0 q1
q1 ChainMachine ? msg1 q2
q2 ChainMachine ? msg2 q3
q3 ChainMachine ? msg3 q4
q4 ChainMachine ? msg4 q5
q5 ChainMachine ? msg5 q0
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(longChainFSA))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	chainMachine := sys.CFSMs[0]
	require.Equal(t, "ChainMachine", chainMachine.Name)

	states := chainMachine.States()
	require.Len(t, states, 6) // s0 through s5

	// Verify chain connectivity
	stateMap := make(map[string]*State)
	for _, state := range states {
		stateMap[state.Label] = state
	}

	for i := 0; i < 6; i++ {
		currentStateLabel := fmt.Sprintf("s%d", i)
		nextStateLabel := fmt.Sprintf("s%d", (i+1)%6) // s5 wraps to s0

		currentState := stateMap[currentStateLabel]
		require.NotNil(t, currentState)

		transitions := currentState.Transitions()
		require.Len(t, transitions, 1)

		tr := transitions[0]
		require.True(t, tr.IsSend())
		require.Equal(t, fmt.Sprintf("msg%d", i), tr.Message())
		require.Equal(t, "Other", tr.NameOfOtherCFSM())
		require.Equal(t, stateMap[nextStateLabel], tr.State())
	}
}

func TestParseFSAFlexibleWhitespace(t *testing.T) {
	// This test uses multiple spaces between transition elements
	// to verify the parser can handle flexible whitespace
	const fsaWithMultipleSpaces = `
.outputs Client
.state graph
q0   Server   !   request   q1
q1   Server   ?   response   q0
.marking q0
.end

.outputs Server
.state graph
q0   Client   ?   request   q1
q1   Client   !   response   q0
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(fsaWithMultipleSpaces))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	client := sys.CFSMs[0]
	server := sys.CFSMs[1]

	require.Equal(t, "Client", client.Name)
	require.Equal(t, "Server", server.Name)

	// Verify transitions work despite multiple spaces
	clientTransitions := client.Start.Transitions()
	require.Len(t, clientTransitions, 1)
	require.Equal(t, "request", clientTransitions[0].Message())
	require.Equal(t, "Server", clientTransitions[0].NameOfOtherCFSM())
}

func TestParseFSAAcceptingStates(t *testing.T) {
	// Test the new .accepting feature that specifies accepting/terminal states
	const fsaWithAcceptingStates = `
.outputs ImgP
.state graph
0 WebUI ? req 1
0 WebUI ? stop 2
1 WebUI ! img 0
.marking 0
.accepting 2
.end

.outputs WebUI
.state graph
0 ImgP ! req 1
0 ImgP ! stop 3
1 ImgP ? img 0
.marking 0
.accepting 3
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(fsaWithAcceptingStates))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	imgP := sys.CFSMs[0]
	webUI := sys.CFSMs[1]

	require.Equal(t, "ImgP", imgP.Name)
	require.Equal(t, "WebUI", webUI.Name)

	// Test ImgP accepting states
	imgPStates := imgP.States()
	require.Len(t, imgPStates, 3) // states 0, 1, 2

	// Find state with label "2" which should be accepting
	var state2 *State
	for _, state := range imgPStates {
		if state.Label == "2" {
			state2 = state
			break
		}
	}
	require.NotNil(t, state2)

	// Verify ImgP has accepting states and state "2" is accepting
	acceptingStates := imgP.AcceptingStates()
	require.Len(t, acceptingStates, 1)
	require.Contains(t, acceptingStates, state2)
	require.True(t, imgP.IsAcceptingState(state2))

	// Test WebUI accepting states
	webUIStates := webUI.States()
	require.Len(t, webUIStates, 3) // states 0, 1, 3

	// Find state with label "3" which should be accepting
	var state3 *State
	for _, state := range webUIStates {
		if state.Label == "3" {
			state3 = state
			break
		}
	}
	require.NotNil(t, state3)

	// Verify WebUI has accepting states and state "3" is accepting
	webUIAcceptingStates := webUI.AcceptingStates()
	require.Len(t, webUIAcceptingStates, 1)
	require.Contains(t, webUIAcceptingStates, state3)
	require.True(t, webUI.IsAcceptingState(state3))

	// Verify non-accepting states
	for _, state := range imgPStates {
		if state.Label != "2" {
			require.False(t, imgP.IsAcceptingState(state))
		}
	}
	for _, state := range webUIStates {
		if state.Label != "3" {
			require.False(t, webUI.IsAcceptingState(state))
		}
	}
}

func TestParseFSAMultipleAcceptingStates(t *testing.T) {
	// Test multiple accepting states in a single CFSM
	const fsaWithMultipleAcceptingStates = `
.outputs MultiAccept
.state graph
start Service ! init state1
start Service ! cancel end1
state1 Service ? ok end1
state1 Service ? error end2
.marking start
.accepting end1 end2
.end

.outputs Service  
.state graph
idle MultiAccept ? init processing
idle MultiAccept ? cancel idle
processing MultiAccept ! ok idle
processing MultiAccept ! error idle
.marking idle
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(fsaWithMultipleAcceptingStates))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	multiAccept := sys.CFSMs[0]
	service := sys.CFSMs[1]

	// Test MultiAccept has multiple accepting states
	acceptingStates := multiAccept.AcceptingStates()
	require.Len(t, acceptingStates, 2)

	// Find end1 and end2 states
	var end1State, end2State *State
	for _, state := range multiAccept.States() {
		if state.Label == "end1" {
			end1State = state
		} else if state.Label == "end2" {
			end2State = state
		}
	}
	require.NotNil(t, end1State)
	require.NotNil(t, end2State)

	// Verify both are accepting
	require.True(t, multiAccept.IsAcceptingState(end1State))
	require.True(t, multiAccept.IsAcceptingState(end2State))
	require.Contains(t, acceptingStates, end1State)
	require.Contains(t, acceptingStates, end2State)

	// Test Service has no accepting states (no .accepting line)
	serviceAcceptingStates := service.AcceptingStates()
	require.Len(t, serviceAcceptingStates, 0)

	// Verify no states are accepting in Service
	for _, state := range service.States() {
		require.False(t, service.IsAcceptingState(state))
	}
}

func TestParseFSANoAcceptingStates(t *testing.T) {
	// Test CFSM without any .accepting line
	const fsaWithoutAcceptingStates = `
.outputs Simple
.state graph
q0 Other ! msg q1
q1 Other ? ack q0
.marking q0
.end

.outputs Other
.state graph
q0 Simple ? msg q1
q1 Simple ! ack q0
.marking q0
.end`

	sys, err := ParseSystemCFSMsFSA(strings.NewReader(fsaWithoutAcceptingStates))
	require.NoError(t, err)
	require.Len(t, sys.CFSMs, 2)

	simple := sys.CFSMs[0]
	other := sys.CFSMs[1]

	// Both CFSMs should have no accepting states
	require.Len(t, simple.AcceptingStates(), 0)
	require.Len(t, other.AcceptingStates(), 0)

	// No states should be accepting
	for _, state := range simple.States() {
		require.False(t, simple.IsAcceptingState(state))
	}
	for _, state := range other.States() {
		require.False(t, other.IsAcceptingState(state))
	}
}
