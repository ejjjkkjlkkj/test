package zero

import "errors"

// SemanticState is the deterministic state owned by one ZERO identity.
type SemanticState struct {
	Identity string
	Status   string
	Sequence uint64
}

// Transition applies an already-authorized semantic result.
// It rejects contradictory or unknown results instead of guessing a state.
func Transition(state SemanticState, result Result) (SemanticState, error) {
	if state.Identity == "" {
		return SemanticState{}, errors.New("missing state identity")
	}
	if result.SourceID != state.Identity {
		return state, errors.New("result identity mismatch")
	}
	switch result.Status {
	case "SUCCESS", "DEFERRED", "FAILED":
		if result.Next == "" {
			return state, errors.New("missing next state")
		}
		return SemanticState{Identity: state.Identity, Status: result.Next, Sequence: state.Sequence + 1}, nil
	case "UNKNOWN", "CONTRADICTION":
		return state, errors.New("non-transitioning result")
	default:
		return state, errors.New("unknown result status")
	}
}

// EngineReceipt is the complete deterministic bootstrap output of one semantic operation.
type EngineReceipt struct {
	Canonical     string
	Authorization Authorization
	Result        Result
	Previous      SemanticState
	Next          SemanticState
	Transitioned  bool
}

// ExecuteRecord performs decode/validate, authorization, semantic execution and state transition.
// It is intentionally a pure reference pipeline around the existing deterministic primitives.
func ExecuteRecord(line, actor, operation string, state SemanticState) (EngineReceipt, error) {
	record, err := Decode(line)
	if err != nil {
		return EngineReceipt{}, err
	}
	canonical, err := Encode(record)
	if err != nil {
		return EngineReceipt{}, err
	}
	req := ExecutionRequest{
		Operation: Instruction(operation),
		Actor: actor,
		Target: record.Target,
	}
	decision := Authorize(req)
	if !decision.Allowed {
		return EngineReceipt{Canonical: canonical, Authorization: decision}, errors.New(decision.Reason)
	}
	result := Execute(record, operation)
	if state.Identity == "" {
		state = SemanticState{Identity: record.Identity, Status: result.Previous}
	}
	if result.Status == "CONTRADICTION" {
		return EngineReceipt{
			Canonical: canonical, Authorization: decision, Result: result,
			Previous: state, Next: state, Transitioned: false,
		}, nil
	}
	next, err := Transition(state, result)
	if err != nil {
		return EngineReceipt{
			Canonical: canonical, Authorization: decision, Result: result, Previous: state,
		}, err
	}
	return EngineReceipt{
		Canonical: canonical, Authorization: decision, Result: result,
		Previous: state, Next: next, Transitioned: true,
	}, nil
}
