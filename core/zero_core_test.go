package core

import "testing"

func TestTruthInvariant(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Fatal("1+1 must equal 2")
	}
	if VerifyAddition(1, 1, 3) {
		t.Fatal("1+1=3 must be rejected")
	}
	if err := RejectWrongAddition(1, 1, 3); err == nil {
		t.Fatal("wrong arithmetic claim was accepted")
	}
}

func TestHardwareIndependentTruth(t *testing.T) {
	profiles := []string{"PC", "IPHONE", "ANDROID", "EMBEDDED", "AIRBORNE", "SPACECRAFT", "REMOTE"}
	for _, profile := range profiles {
		if Add(1, 1) != 2 {
			t.Fatalf("%s changed formal truth", profile)
		}
	}
}

func TestAccessibilityPreservesSemantics(t *testing.T) {
	want := SemanticResult{Identity: "ARITHMETIC:1+1", Operation: "ADD", Value: "2", Proof: ProofTested}
	for _, modality := range Modalities {
		got, err := Project(want, modality)
		if err != nil {
			t.Fatalf("%s: %v", modality, err)
		}
		if got != want {
			t.Fatalf("%s changed semantic result: %#v", modality, got)
		}
	}
}

func TestProofPromotion(t *testing.T) {
	for _, pair := range [][2]Proof{
		{ProofSimulated, ProofHardware},
		{ProofQEMU, ProofHardware},
		{ProofHardware, ProofRF},
		{ProofRF, ProofSatellite},
	} {
		if IsProofPromotionAllowed(pair[0], pair[1]) {
			t.Fatalf("illegal proof promotion %s -> %s accepted", pair[0], pair[1])
		}
	}
}

func TestRecordRoundTrip(t *testing.T) {
	want := Record{
		Version: "1", Type: "OBJECT", Identity: "TEST|1",
		Sequence: 1, Time: "UNKNOWN", Source: "ZERO",
		Target: "TEST", Payload: "meaning=accessible|machine",
		Proof: "TESTED",
	}
	got, err := DecodeRecord(EncodeRecord(want))
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("round trip mismatch: %#v != %#v", got, want)
	}
}

func TestInvalidEscapeRejected(t *testing.T) {
	if _, err := DecodeRecord("RECORD|1|OBJECT|x|1|UNKNOWN|ZERO|TEST|bad\\q|TESTED"); err == nil {
		t.Fatal("invalid escape accepted")
	}
}
