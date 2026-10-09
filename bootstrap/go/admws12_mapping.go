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

 
// DecodeADMWS12Payload decodes the adapter's fixed-order payload and validates
// its source kind/state pair. It does not promote the source claim to proof.
func DecodeADMWS12Payload(payload string) (kind, state, value string, err error) {
	fields, err := splitEscaped(payload)
	if err != nil {
		return "", "", "", err
	}
	if len(fields) != 4 || fields[0] != "source_project=ADMWS12" {
		return "", "", "", errors.New("invalid ADMWS12 payload shape")
	}
	kindKey, kindValue, ok := strings.Cut(fields[1], "=")
	if !ok || kindKey != "source_kind" {
		return "", "", "", errors.New("invalid ADMWS12 kind field")
	}
	stateKey, stateValue, ok := strings.Cut(fields[2], "=")
	if !ok || stateKey != "source_state" {
		return "", "", "", errors.New("invalid ADMWS12 state field")
	}
	valueKey, valueValue, ok := strings.Cut(fields[3], "=")
	if !ok || valueKey != "value" {
		return "", "", "", errors.New("invalid ADMWS12 value field")
	}
	if _, err := MapADMWS12Record(kindValue, "payload-check", stateValue, valueValue, 0); err != nil {
		return "", "", "", err
	}
	return kindValue, stateValue, valueValue, nil
}
