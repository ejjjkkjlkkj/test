package zero

import (
	"errors"
	"strings"
)

// MapADMWS12Record maps a capability/evidence state without upgrading proof.
func MapADMWS12Record(kind, identity, state, value string, sequence uint64) (Record, error) {
	if strings.TrimSpace(identity) == "" {
		return Record{}, errors.New("empty identity")
	}
	recordType := ""
	switch kind {
	case "Capability":
		recordType = "CAPABILITY"
		switch state {
		case "UNKNOWN", "AVAILABLE", "UNAVAILABLE", "UNSUPPORTED", "FAILED":
		default:
			return Record{}, errors.New("invalid capability state")
		}
	case "Evidence":
		recordType = "OBSERVATION"
		switch state {
		case "OBSERVED", "ABSENT", "UNSUPPORTED", "FAILED", "UNKNOWN":
		default:
			return Record{}, errors.New("invalid evidence state")
		}
	default:
		return Record{}, errors.New("unknown source kind")
	}
	payload := strings.Join([]string{
		"source_project=ADMWS12",
		"source_kind=" + escape(kind),
		"source_state=" + escape(state),
		"value=" + escape(value),
	}, "|")
	r := Record{
		Version: "1", Type: recordType, Identity: identity, Sequence: sequence,
		Time: "UNKNOWN", Source: "ADMWS12", Target: "ZERO",
		Payload: payload, Proof: "UNPROVEN",
	}
	if err := Validate(r); err != nil {
		return Record{}, err
	}
	return r, nil
}
