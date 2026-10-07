package zero

import "testing"

func TestTransitionDeterministicAndIdentityBound(t *testing.T) {
	state := SemanticState{Identity: "id", Status: "CREATED"}
	result := Result{Status: "SUCCESS", Next: "OBSERVED", SourceID: "id"}
	first, err := Transition(state, result)
	if err != nil { t.Fatal(err) }
	second, err := Transition(state, result)
	if err != nil { t.Fatal(err) }
	if first != second { t.Fatalf("non-deterministic transition: %+v != %+v", first, second) }
	if first.Status != "OBSERVED" || first.Sequence != 1 { t.Fatalf("unexpected state: %+v", first) }
}

func TestTransitionRejectsIdentityMismatchAndNonTransition(t *testing.T) {
	state := SemanticState{Identity: "id", Status: "CREATED"}
	if _, err := Transition(state, Result{Status: "SUCCESS", Next: "OBSERVED", SourceID: "other"}); err == nil {
		t.Fatal("identity mismatch accepted")
	}
	if _, err := Transition(state, Result{Status: "UNKNOWN", Next: "CREATED", SourceID: "id"}); err == nil {
		t.Fatal("unknown result transitioned state")
	}
}

func TestExecuteRecordProducesCompleteReceipt(t *testing.T) {
	line := "RECORD|1|OBSERVE|id|7|2026-10-07T12:00:00Z|source|target|payload|proof"
	receipt, err := ExecuteRecord(line, "actor", "OBSERVE", SemanticState{})
	if err != nil { t.Fatal(err) }
	if !receipt.Authorization.Allowed { t.Fatal("record was not authorized") }
	if receipt.Result.Status != "SUCCESS" { t.Fatalf("unexpected result: %+v", receipt.Result) }
	if receipt.Previous.Identity != "id" || receipt.Next.Status != "OBSERVED" || receipt.Next.Sequence != 1 {
		t.Fatalf("unexpected receipt states: %+v -> %+v", receipt.Previous, receipt.Next)
	}
	if receipt.Canonical != line { t.Fatalf("canonical representation changed: %q", receipt.Canonical) }
}

func TestExecuteRecordRejectsBeforeSemanticExecution(t *testing.T) {
	line := "RECORD|1|OBSERVE|id|7|2026-10-07T12:00:00Z|source|target|payload|proof"
	state := SemanticState{Identity: "id", Status: "CREATED"}
	receipt, err := ExecuteRecord(line, "", "OBSERVE", state)
	if err == nil { t.Fatal("missing actor accepted") }
	if receipt.Authorization.Allowed { t.Fatal("unauthorized receipt marked allowed") }
	if receipt.Result.Status != "" || receipt.Next != (SemanticState{}) {
		t.Fatalf("semantic execution occurred after rejection: %+v", receipt)
	}
}
