package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Proof string

const (
	ProofUnknown Proof = "UNKNOWN"
	ProofDefined Proof = "DEFINED"
	ProofImplemented Proof = "IMPLEMENTED"
	ProofTested Proof = "TESTED"
	ProofSimulated Proof = "SIMULATED"
	ProofQEMU Proof = "QEMU"
	ProofHardware Proof = "HARDWARE"
	ProofRF Proof = "RF_PROVEN"
	ProofSatellite Proof = "SATELLITE_LINK_PROVEN"
)

type Modality string

var modalities = [...]Modality{"VOICE", "BRAILLE", "KEYBOARD", "DISPLAY", "TOUCH", "POINTER", "NETWORK", "AUTOMATION"}

func IsModalitySupported(modality Modality) bool {
	for _, supported := range modalities {
		if modality == supported {
			return true
		}
	}
	return false
}

func Modalities() []Modality {
	result := make([]Modality, len(modalities))
	copy(result, modalities[:])
	return result
}

type SemanticResult struct {
	Identity, Operation, Value string
	Proof Proof
}

func Add(a, b int64) int64 { return a + b }

func VerifyAddition(a, b, claimed int64) bool { return Add(a, b) == claimed }

func RejectWrongAddition(a, b, claimed int64) error {
	if VerifyAddition(a, b, claimed) {
		return nil
	}
	return fmt.Errorf("INVALID_ARITHMETIC: %d + %d != %d", a, b, claimed)
}

func Project(result SemanticResult, modality Modality) (SemanticResult, error) {
	if modality == "" {
		return SemanticResult{}, errors.New("ACCESSIBILITY.MODALITY.EMPTY")
	}
	if !IsModalitySupported(modality) {
		return SemanticResult{}, fmt.Errorf("ACCESSIBILITY.MODALITY.UNSUPPORTED: %s", modality)
	}
	return result, nil
}

func proofRank(proof Proof) int {
	switch proof {
	case ProofUnknown:
		return 0
	case ProofDefined:
		return 1
	case ProofImplemented:
		return 2
	case ProofTested:
		return 3
	case ProofSimulated:
		return 4
	case ProofQEMU:
		return 5
	case ProofHardware:
		return 6
	case ProofRF:
		return 7
	case ProofSatellite:
		return 8
	default:
		return -1
	}
}

func IsProofPromotionAllowed(from, to Proof) bool {
	fromRank, toRank := proofRank(from), proofRank(to)
	return fromRank >= 0 && toRank == fromRank+1
}

type Record struct {
	Version, Type, Identity string
	Sequence int64
	Time, Source, Target string
	Payload, Proof string
}

func Escape(value string) string {
	var b strings.Builder
	for _, r := range value {
		switch r {
		case '\\':
			b.WriteString("\\\\")
		case '|':
			b.WriteString("\\|")
		case '\n':
			b.WriteString("\\n")
		case '\r':
			b.WriteString("\\r")
		case '\t':
			b.WriteString("\\t")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

func Unescape(value string) (string, error) {
	var b strings.Builder
	escaped := false
	for _, r := range value {
		if escaped {
			switch r {
			case '\\', '|':
				b.WriteRune(r)
			case 'n':
				b.WriteByte('\n')
			case 'r':
				b.WriteByte('\r')
			case 't':
				b.WriteByte('\t')
			default:
				return "", fmt.Errorf("FORMAT.ESCAPE: %q", r)
			}
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}
		b.WriteRune(r)
	}
	if escaped {
		return "", errors.New("FORMAT.ESCAPE: trailing escape")
	}
	return b.String(), nil
}

func EncodeRecord(r Record) string {
	fields := []string{"RECORD", r.Version, r.Type, r.Identity, strconv.FormatInt(r.Sequence, 10), r.Time, r.Source, r.Target, r.Payload, r.Proof}
	for i := range fields {
		fields[i] = Escape(fields[i])
	}
	return strings.Join(fields, "|")
}

func splitRecord(line string) ([]string, error) {
	var fields []string
	var b strings.Builder
	escaped := false
	for _, r := range line {
		if escaped {
			b.WriteRune('\\')
			b.WriteRune(r)
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			continue
		}
		if r == '|' {
			fields = append(fields, b.String())
			b.Reset()
			continue
		}
		b.WriteRune(r)
	}
	if escaped {
		return nil, errors.New("FORMAT.ESCAPE: trailing escape")
	}
	return append(fields, b.String()), nil
}

func DecodeRecord(line string) (Record, error) {
	fields, err := splitRecord(line)
	if err != nil || len(fields) != 10 || fields[0] != "RECORD" {
		return Record{}, errors.New("FORMAT.RECORD")
	}
	values := make([]string, 10)
	for i := 1; i < 10; i++ {
		values[i], err = Unescape(fields[i])
		if err != nil {
			return Record{}, err
		}
	}
	sequence, err := strconv.ParseInt(values[4], 10, 64)
	if err != nil || sequence < 0 {
		return Record{}, errors.New("FORMAT.SEQUENCE")
	}
	return Record{
		Version: values[1], Type: values[2], Identity: values[3],
		Sequence: sequence, Time: values[5], Source: values[6], Target: values[7],
		Payload: values[8], Proof: values[9],
	}, nil
}
