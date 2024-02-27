package cfsm

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

// This function converts a CFSM to a string in the format of the Python Bisimulation library
// https://github.com/diegosenarruzza/bisimulation/
func ConvertCFSMToPythonBisimulationFormat(contract *CFSM) ([]byte, error) {
	// Import statement
	const importStatement = "from cfsm_bisimulation import CommunicatingFiniteStateMachine\n\n"

	// We need all participant names to start with a single uppercase letter and then all lowercase or numbers.
	// We'll have to keep track of all the translations we do here, so we can translate back when we're done.
	participantNameTranslations := make(map[string]string)
	otherCFSMs := make([]string, 0)
	for _, name := range contract.OtherCFSMs() {
		msgRegex := `[A-Z][a-z0-9]*`
		matched, err := regexp.MatchString(msgRegex, name)
		if err != nil {
			return nil, err
		}
		translatedName := name
		if matched {
			participantNameTranslations[name] = name
		} else {
			translatedName = "C" + filterOutNonAlphanumeric(name)
			participantNameTranslations[name] = translatedName
			// TODO: check that the new name is unique.
		}
		otherCFSMs = append(otherCFSMs, translatedName)
	}

	// Diego's format expects also the name for the CFSM we're defining, so we'll have to pick
	// a name that is unused. We'll use 'self' unless it's already used. If it is, we'll generate
	// a random string and use that (verifying it's also unused).

	selfname := "Self"
	restart := true
	for restart {
		restart = false
		for _, name := range otherCFSMs {
			if name == selfname {
				selfname = "Self" + generateRandomString(5)
				restart = true
				break
			}
		}
	}
	allCFSMs := append([]string{selfname}, otherCFSMs...)
	allCFSMsJSON, err := json.Marshal(allCFSMs)
	if err != nil {
		return nil, err
	}
	createCFSM := "cfsm = CommunicatingFiniteStateMachine(" + string(allCFSMsJSON) + ")\n\n"

	// Add states
	addStates := ""
	for _, state := range contract.States() {
		addStates += fmt.Sprintf("cfsm.add_states('%d')\n", state.ID)
	}
	addStates += "\n"

	// Set initial state
	setInitialState := fmt.Sprintf("cfsm.set_as_initial('%d')\n\n", contract.Start.ID)

	// Add transitions
	addTransitions := ""
	messageTranslations := make(map[string]string)
	messageTranslationsInverted := make(map[string]string)
	for _, state := range contract.States() {
		for _, transition := range state.Transitions() {
			// cfsm_bisimulation format expects the message to be a string with the format like this:
			// Sender Receiver [!|?] function([typed parameters]). e.g "Client Server ! f(int x)" or "ServerClient?g()"
			// It actually uses this regex to match the "action_string" paramter:
			// (?P<participant1>[A-Z][a-z0-9]*)\s*(?P<participant2>[A-Z][a-z0-9]*)\s*(?P<action>!|\?)\s*(?P<tag>\w+)\((?P<payload>.*)\)
			// https://github.com/diegosenarruzza/bisimulation/blob/81af48aed977b79e41b2a96c36160354e230f5b2/src/cfsm_bisimulation/models/communicating_system/action_parser.py#L5-L8

			// So we need participant names to start with a single uppercase letter and then all lowercase or numbers.
			// We also need the "tag" to have parentheses after it. We're not going to use the payload (leave empty).
			// We also will not set any conditions, so we'll use TrueFormula for that.

			// We'll have to keep track of all the translations we do here, so we can translate back when we're done.

			translatedMessage, ok := messageTranslations[transition.Message()]
			if !ok {
				translatedMessage = fmt.Sprintf("message%d()", len(messageTranslations))
				messageTranslations[transition.Message()] = translatedMessage
				messageTranslationsInverted[translatedMessage] = transition.Message()
			}

			var transitionTypeMarker string
			if transition.IsSend() {
				transitionTypeMarker = "!"
			} else {
				transitionTypeMarker = "?"
			}
			action := fmt.Sprintf("%s %s %s %s", selfname, participantNameTranslations[transition.NameOfOtherCFSM()], transitionTypeMarker, translatedMessage)
			addTransitions += fmt.Sprintf("cfsm.add_transition_between('%d', '%d', '%s')\n", state.ID, transition.State().ID, action)
		}
	}

	// Combine all code blocks
	code := importStatement + createCFSM + addStates + setInitialState + addTransitions + "\n"

	return []byte(code), nil
}

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9]+`)

func filterOutNonAlphanumeric(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(strings.ToLower(str), "")
}
