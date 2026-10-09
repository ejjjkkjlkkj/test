package zero

import (
	"strings"
	"testing"
)

func TestCanonicalEscaping(t *testing.T) {
	r := Record{
		Version: "1", Type: "EVENT", Identity: "escape", Sequence: 7,
		Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE",
		Payload: "a|b\\c\nd\te — café", Proof: "DEFINED",
	}
	wire, err := Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Decode(wire)
	if err != nil {
		t.Fatal(err)
	}
	if got != r {
		t.Fatalf("escaping changed record: %#v != %#v", got, r)
	}
}

func TestCanonicalInteger(t *testing.T) {
	r := Record{
		Version: "1", Type: "EVENT", Identity: "integer", Sequence: 1,
		Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE",
		Payload: "x", Proof: "DEFINED",
	}
	wire, err := Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	if wire == "" {
		t.Fatal("empty encoding")
	}
	got, err := Decode(wire)
	if err != nil || got.Sequence != 1 {
		t.Fatalf("bad canonical sequence: %v %#v", err, got)
	}
}

func TestDecodeRejectsNonCanonicalSequence(t *testing.T) {
	line := "RECORD|1|EVENT|id|01|UNKNOWN|source|target|payload|proof"
	if _, err := Decode(line); err == nil {
		t.Fatal("sequence with a leading zero was accepted")
	}
}

func TestDecodeRejectsUnknownEscape(t *testing.T) {
	line := "RECORD|1|EVENT|id|1|UNKNOWN|source|target|bad\\qescape|DEFINED"
	if _, err := Decode(line); err == nil {
		t.Fatal("unknown escape sequence was accepted")
	}
}

func TestRoundTripIsCanonicalAndStable(t *testing.T) {
	line := "RECORD|1|EVENT|id|12|UNKNOWN|source|target|payload|DEFINED"
	got, err := RoundTrip(line)
	if err != nil {
		t.Fatal(err)
	}
	if got != line {
		t.Fatalf("round-trip changed canonical bytes: %q != %q", got, line)
	}
}

func TestADMWS12MappingPreservesUnprovenState(t *testing.T) {
	r, err := MapADMWS12Record("Capability", "cpu", "AVAILABLE", "present", 1)
	if err != nil {
		t.Fatal(err)
	}
	if r.Type != "CAPABILITY" || r.Source != "ADMWS12" || r.Target != "ZERO" || r.Proof != "UNPROVEN" {
		t.Fatalf("unexpected mapping: %#v", r)
	}
}

func TestADMWS12MappingRejectsUnknownState(t *testing.T) {
	if _, err := MapADMWS12Record("Capability", "cpu", "OBSERVED", "present", 1); err == nil {
		t.Fatal("invalid capability state was accepted")
	}
}

func TestADMWS12EvidenceStateAndPayloadRoundTrip(t *testing.T) {
	r, err := MapADMWS12Record("Evidence", "device", "ABSENT", "café | \\ value", 2)
	if err != nil {
		t.Fatal(err)
	}
	if r.Type != "OBSERVATION" || r.Proof != "UNPROVEN" {
		t.Fatalf("unexpected evidence mapping: %#v", r)
	}
	wire, err := Encode(r)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Decode(wire)
	if err != nil {
		t.Fatal(err)
	}
	if got != r {
		t.Fatalf("mapping round trip changed record: %#v != %#v", got, r)
	}
	if !strings.Contains(got.Payload, "source_state=ABSENT") {
		t.Fatalf("source state not preserved: %q", got.Payload)
	}
}
