package zero

import "errors"

// ExecutionRequest describes the semantic authorization boundary around an instruction.
// It is a bootstrap reference model; it is not an OS or hardware security boundary.
type ExecutionRequest struct {
	Operation Instruction
	Actor     string
	Target    string
}

// Authorization records the deterministic decision made before execution.
type Authorization struct {
	Allowed bool
	Reason  string
}

// SemanticRequest describes authorization for the canonical record execution layer.
type SemanticRequest struct {
	Operation string
	Actor     string
	Target    string
}

// AuthorizeSemantic applies the explicit semantic operation allowlist.
// Unknown operations are rejected instead of being converted into UNKNOWN implicitly.
func AuthorizeSemantic(req SemanticRequest) Authorization {
	if req.Actor == "" {
		return Authorization{Reason: "missing actor"}
	}
	if req.Target == "" {
		return Authorization{Reason: "missing target"}
	}
	switch req.Operation {
	case "OBSERVE", "SET_STATE:DELIVERED", "DELIVERY:DEFERRED", "DELIVERY:FAILED":
		return Authorization{Allowed: true, Reason: "allowed"}
	default:
		return Authorization{Reason: "unknown semantic operation"}
	}
}

// Event is the deterministic observation emitted after an accepted instruction.
type Event struct {
	Operation Instruction
	Actor     string
	Target    string
	PCBefore  uint64
	PCAfter   uint64
	Changed   bool
}

// Authorize applies the bootstrap policy. Empty actor/target and unknown operations are rejected.
func Authorize(req ExecutionRequest) Authorization {
	if req.Actor == "" {
		return Authorization{Reason: "missing actor"}
	}
	if req.Target == "" {
		return Authorization{Reason: "missing target"}
	}
	switch req.Operation {
	case OpObserve, OpRead, OpWrite, OpStep, OpRequest, OpNextEvent:
		return Authorization{Allowed: true, Reason: "allowed"}
	default:
		return Authorization{Reason: "unknown instruction"}
	}
}

// ExecuteAuthorized enforces authorization before mutating or advancing the machine.
func (m *Machine) ExecuteAuthorized(req ExecutionRequest, address, length uint64, data []byte) (ExecutionResult, Event, error) {
	decision := Authorize(req)
	if !decision.Allowed {
		return ExecutionResult{}, Event{}, errors.New(decision.Reason)
	}
	result, err := m.Execute(req.Operation, address, length, data)
	if err != nil {
		return ExecutionResult{}, Event{}, err
	}
	event := Event{
		Operation: req.Operation,
		Actor:     req.Actor,
		Target:    req.Target,
		PCBefore:  result.PCBefore,
		PCAfter:   result.PCAfter,
		Changed:   result.Changed,
	}
	return result, event, nil
}
