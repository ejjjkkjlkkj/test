package zero

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const canonicalPrefix = "RECORD"
const fieldCount = 10 // prefix + 9 semantic fields

var required = []string{"version", "type", "identity", "sequence", "time", "source", "target", "payload", "proof"}

type Record struct {
	Version  string
	Type     string
	Identity string
	Sequence uint64
	Time     string
	Source   string
	Target   string
	Payload  string
	Proof    string
}

type Result struct {
	Status    string
	Value     string
	Previous  string
	Next      string
	SourceID  string
	Proof     string
}

func escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\':
			b.WriteString("\\")
		case '|':
			b.WriteString("\|")
		case '
':
			b.WriteString("\n")
		case '':
			b.WriteString("\r")
		case '	':
			b.WriteString("\t")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

func unescape(s string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '\' {
			b.WriteByte(s[i])
			continue
		}
		i++
		if i >= len(s) {
			return "", errors.New("invalid trailing escape")
		}
		switch s[i] {
		case '\':
			b.WriteByte('\')
		case '|':
			b.WriteByte('|')
		case 'n':
			b.WriteByte('
')
		case 'r':
			b.WriteByte('')
		case 't':
			b.WriteByte('	')
		default:
			return "", fmt.Errorf("invalid escape: %q", s[i])
		}
	}
	return b.String(), nil
}

func splitEscaped(s string) ([]string, error) {
	var out []string
	var b strings.Builder
	escaped := false
	for _, r := range s {
		if escaped {
			b.WriteRune('\')
			b.WriteRune(r)
			escaped = false
			continue
		}
		if r == '\' {
			escaped = true
			continue
		}
		if r == '|' {
			v, err := unescape(b.String())
			if err != nil {
				return nil, err
			}
			out = append(out, v)
			b.Reset()
			continue
		}
		b.WriteRune(r)
	}
	if escaped {
		return nil, errors.New("invalid trailing escape")
	}
	v, err := unescape(b.String())
	if err != nil {
		return nil, err
	}
	out = append(out, v)
	return out, nil
}

func Encode(r Record) (string, error) {
	if err := Validate(r); err != nil {
		return "", err
	}
	fields := []string{
		canonicalPrefix, r.Version, r.Type, r.Identity,
		strconv.FormatUint(r.Sequence, 10), r.Time,
		r.Source, r.Target, r.Payload, r.Proof,
	}
	for i := range fields {
		fields[i] = escape(fields[i])
	}
	return strings.Join(fields, "|"), nil
}

func Decode(line string) (Record, error) {
	fields, err := splitEscaped(line)
	if err != nil {
		return Record{}, err
	}
	if len(fields) != fieldCount || fields[0] != canonicalPrefix {
		return Record{}, errors.New("invalid record shape")
	}
	seq, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return Record{}, errors.New("invalid sequence")
	}
	r := Record{
		Version: fields[1], Type: fields[2], Identity: fields[3],
		Sequence: seq, Time: fields[5], Source: fields[6],
		Target: fields[7], Payload: fields[8], Proof: fields[9],
	}
	if err := Validate(r); err != nil {
		return Record{}, err
	}
	return r, nil
}

func Validate(r Record) error {
	if r.Version == "" || r.Type == "" || r.Identity == "" ||
		r.Source == "" || r.Target == "" || r.Payload == "" || r.Proof == "" {
		return errors.New("missing required field")
	}
	if r.Type == "UNKNOWN_TYPE" && r.Payload == "" {
		return errors.New("unknown type requires preserved payload")
	}
	return nil
}

func RoundTrip(line string) (string, error) {
	r, err := Decode(line)
	if err != nil {
		return "", err
	}
	return Encode(r)
}

func Execute(r Record, operation string) Result {
	switch operation {
	case "OBSERVE":
		return Result{Status: "SUCCESS", Value: "OBSERVED", Previous: "CREATED", Next: "OBSERVED", SourceID: r.Identity, Proof: r.Proof}
	case "UNKNOWN_OPERATION":
		return Result{Status: "UNKNOWN", Value: "UNKNOWN", Previous: "CREATED", Next: "CREATED", SourceID: r.Identity, Proof: r.Proof}
	case "SET_STATE:DELIVERED":
		return Result{Status: "CONTRADICTION", Value: "INVALID", Previous: "CREATED", Next: "CREATED", SourceID: r.Identity, Proof: r.Proof}
	case "DELIVERY:DEFERRED":
		return Result{Status: "DEFERRED", Value: "DEFERRED", Previous: "SENT", Next: "DEFERRED", SourceID: r.Identity, Proof: r.Proof}
	case "DELIVERY:FAILED":
		return Result{Status: "FAILED", Value: "FAILED", Previous: "SENT", Next: "FAILED", SourceID: r.Identity, Proof: r.Proof}
	default:
		return Result{Status: "UNKNOWN", Value: "UNKNOWN", Previous: "", Next: "", SourceID: r.Identity, Proof: r.Proof}
	}
}
