package zero

import (
	"bytes"
	"testing"
)

func TestAuthorizeRejectsMissingIdentity(t *testing.T) {
	cases := []ExecutionRequest{
		{Operation: OpRead, Target: "memory"},
		{Operation: OpRead, Actor: "actor"},
		{Operation: Instruction("INVALID"), Actor: "actor", Target: "memory"},
	}
	for _, tc := range cases {
		if got := Authorize(tc); got.Allowed {
			t.Fatalf("request unexpectedly authorized: %+v", tc)
		}
	}
}

func TestAuthorizeIsDeterministic(t *testing.T) {
	req := ExecutionRequest{Operation: OpWrite, Actor: "actor", Target: "memory"}
	first := Authorize(req)
	second := Authorize(req)
	if first != second {
		t.Fatalf("authorization changed for identical input: %+v != %+v", first, second)
	}
}

func TestUnauthorizedExecutionDoesNotMutateMachine(t *testing.T) {
	m, err := NewMachine(4)
	if err != nil { t.Fatal(err) }
	before, err := m.Read(0, 4)
	if err != nil { t.Fatal(err) }
	pc := m.PC

	_, _, err = m.ExecuteAuthorized(ExecutionRequest{
		Operation: OpWrite,
		Target:    "memory",
	}, 0, 2, []byte{7, 8})
	if err == nil {
		t.Fatal("unauthorized write accepted")
	}
	after, err := m.Read(0, 4)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(before, after) {
		t.Fatalf("unauthorized request mutated memory: %v -> %v", before, after)
	}
	if m.PC != pc {
		t.Fatalf("unauthorized request advanced PC: %d -> %d", pc, m.PC)
	}
}

func TestAuthorizedExecutionEmitsDeterministicEvent(t *testing.T) {
	m, err := NewMachine(4)
	if err != nil { t.Fatal(err) }
	req := ExecutionRequest{Operation: OpWrite, Actor: "actor", Target: "memory"}
	result, event, err := m.ExecuteAuthorized(req, 1, 2, []byte{7, 8})
	if err != nil { t.Fatal(err) }
	want := Event{
		Operation: OpWrite, Actor: "actor", Target: "memory",
		PCBefore: 0, PCAfter: 1, Changed: true,
	}
	if event != want {
		t.Fatalf("unexpected event: %+v != %+v", event, want)
	}
	if result.PCBefore != event.PCBefore || result.PCAfter != event.PCAfter {
		t.Fatal("result/event PC mismatch")
	}
	got, err := m.Read(1, 2)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(got, []byte{7, 8}) {
		t.Fatalf("write did not persist: %v", got)
	}
}
