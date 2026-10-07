package core

import "testing"

func TestTruthInvariant(t *testing.T) {
    if Add(1,1)!=2 { t.Fatal("1+1 must equal 2") }
    if VerifyAddition(1,1,3) { t.Fatal("1+1=3 must be rejected") }
    if err:=RejectWrongAddition(1,1,3); err==nil { t.Fatal("wrong arithmetic claim was accepted") }
    truth:=CanonicalTruth()
    if !truth.Immutable || truth.Result!=2 || truth.Expression!="1+1" || OnePlusOne()!=2 { t.Fatal("canonical truth invariant failed") }
    if !VerifyCanonicalTruth(2) || VerifyCanonicalTruth(3) { t.Fatal("canonical truth verifier failed") }
}
func TestHardwareIndependentTruth(t *testing.T) {
    for _,profile:=range []string{"PC","IPHONE","ANDROID","EMBEDDED","AIRBORNE","SPACECRAFT","REMOTE","VIRTUALIZED"} {
        if OnePlusOne()!=2 || Add(1,1)!=2 { t.Fatalf("%s changed formal truth",profile) }
    }
}
func TestAccessibilityPreservesSemantics(t *testing.T) {
    want:=SemanticResult{Identity:"ARITHMETIC:1+1",Operation:"ADD",Value:"2",Proof:ProofTested}
    for _,modality:=range Modalities() { got,err:=Project(want,modality); if err!=nil { t.Fatalf("%s: %v",modality,err) }; if got!=want { t.Fatalf("%s changed semantic result: %#v",modality,got) } }
}
func TestProofPromotionRejectsNonAdjacentEvidence(t *testing.T) {
    for _,pair:=range [][2]Proof{{ProofUnknown,ProofTested},{ProofDefined,ProofHardware},{ProofSimulated,ProofHardware},{ProofQEMU,ProofRF},{ProofHardware,ProofSatellite}} {
        if IsProofPromotionAllowed(pair[0],pair[1]) { t.Fatalf("unsupported proof jump %s -> %s accepted",pair[0],pair[1]) }
    }
}
func TestRecordRoundTrip(t *testing.T) {
    want:=Record{Version:"1",Type:"OBJECT",Identity:"TEST|1",Sequence:1,Time:"UNKNOWN",Source:"ZERO",Target:"TEST",Payload:"meaning=accessible|machine",Proof:"TESTED"}
    got,err:=DecodeRecord(EncodeRecord(want)); if err!=nil { t.Fatal(err) }; if got!=want { t.Fatalf("round trip mismatch: %#v != %#v",got,want) }
}
func TestInvalidEscapeRejected(t *testing.T) {
    if _,err:=DecodeRecord("RECORD|1|OBJECT|x|1|UNKNOWN|ZERO|TEST|bad\\q|TESTED"); err==nil { t.Fatal("invalid escape accepted") }
}
func TestUnsupportedAccessibilityIsRejected(t *testing.T) {
    if _,err:=Project(SemanticResult{Identity:"TEST",Value:"2"},"UNKNOWN_MODALITY"); err==nil { t.Fatal("unsupported accessibility modality accepted") }
}
func TestModalitiesCannotMutateCoreRegistry(t *testing.T) {
    got:=Modalities(); got[0]="CORRUPTED"
    if IsModalitySupported("CORRUPTED") || !IsModalitySupported("VOICE") { t.Fatal("caller mutated core modality registry") }
}
func TestProofPromotionRequiresAdjacentEvidence(t *testing.T) {
    allowed:=[][2]Proof{{ProofUnknown,ProofDefined},{ProofDefined,ProofImplemented},{ProofImplemented,ProofTested},{ProofTested,ProofSimulated},{ProofSimulated,ProofQEMU},{ProofQEMU,ProofHardware},{ProofHardware,ProofRF},{ProofRF,ProofSatellite}}
    for _,pair:=range allowed { if !IsProofPromotionAllowed(pair[0],pair[1]) { t.Fatalf("valid adjacent promotion %s -> %s rejected",pair[0],pair[1]) } }
    for _,pair:=range [][2]Proof{{ProofUnknown,ProofTested},{ProofDefined,ProofHardware},{ProofSimulated,ProofHardware},{ProofQEMU,ProofRF},{ProofHardware,ProofSatellite}} {
        if IsProofPromotionAllowed(pair[0],pair[1]) { t.Fatalf("unsupported proof jump %s -> %s accepted",pair[0],pair[1]) }
    }
}
