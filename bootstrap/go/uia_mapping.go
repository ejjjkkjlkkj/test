package zero

import (
	"errors"
	"strings"
)

// MapUIASemanticRecord imports a platform-neutral UIA semantic node into ZERO.
// It intentionally carries no raw control value; callers supply the accessible
// name only, avoiding accidental serialization of password/value content.
func MapUIASemanticRecord(identity, role, nativeRole, name string, states []string, sequence uint64) (Record, error) {
	if strings.TrimSpace(identity) == "" {
		return Record{}, errors.New("empty semantic identity")
	}
	if strings.TrimSpace(role) == "" {
		return Record{}, errors.New("empty semantic role")
	}
	for _, state := range states {
		if strings.TrimSpace(state) == "" || strings.Contains(state, ",") {
			return Record{}, errors.New("invalid semantic state")
		}
	}
	payload := strings.Join([]string{
		"source_project=NVDA-RUST-UIA-STANDALONE",
		"role=" + escape(role),
		"native_role=" + escape(nativeRole),
		"name=" + escape(name),
		"states=" + escape(strings.Join(states, ",")),
	}, "|")
	r := Record{
		Version: "1", Type: "SEMANTIC_NODE", Identity: identity, Sequence: sequence,
		Time: "UNKNOWN", Source: "NVDA-RUST-UIA-STANDALONE", Target: "ZERO",
		Payload: payload, Proof: "UNPROVEN",
	}
	if err := Validate(r); err != nil {
		return Record{}, err
	}
	return r, nil
}

// DecodeUIASemanticPayload decodes a semantic-node payload while preserving
// unknown role and state names for forward compatibility.
func DecodeUIASemanticPayload(payload string) (role, nativeRole, name string, states []string, err error) {
	fields, err := splitEscaped(payload)
	if err != nil {
		return "", "", "", nil, err
	}
	if len(fields) != 5 || fields[0] != "source_project=NVDA-RUST-UIA-STANDALONE" {
		return "", "", "", nil, errors.New("invalid UIA semantic payload shape")
	}
	read := func(field, expected string) (string, error) {
		key, value, ok := strings.Cut(field, "=")
		if !ok || key != expected {
			return "", errors.New("invalid UIA semantic payload field")
		}
		return value, nil
	}
	if role, err = read(fields[1], "role"); err != nil {
		return "", "", "", nil, err
	}
	if nativeRole, err = read(fields[2], "native_role"); err != nil {
		return "", "", "", nil, err
	}
	if name, err = read(fields[3], "name"); err != nil {
		return "", "", "", nil, err
	}
	stateText, err := read(fields[4], "states")
	if err != nil {
		return "", "", "", nil, err
	}
	if strings.TrimSpace(role) == "" {
		return "", "", "", nil, errors.New("empty semantic role")
	}
	if stateText == "" {
		states = []string{}
	} else {
		states = strings.Split(stateText, ",")
		for _, state := range states {
			if strings.TrimSpace(state) == "" {
				return "", "", "", nil, errors.New("invalid semantic state")
			}
		}
	}
	return role, nativeRole, name, states, nil
}
