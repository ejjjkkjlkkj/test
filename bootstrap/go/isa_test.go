package zero

import (
	"bytes"
	"testing"
)

func TestMachineReadWriteBounds(t *testing.T) {
	m, err := NewMachine(8)
	if err != nil { t.Fatal(err) }
	want := []byte{1, 2, 3}
	if err := m.Write(2, want); err != nil { t.Fatal(err) }
	got, err := m.Read(2, 3)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(got, want) { t.Fatalf("read mismatch: %v != %v", got, want) }
	if _, err := m.Read(7, 2); err == nil { t.Fatal("out-of-bounds read accepted") }
	if err := m.Write(7, []byte{1, 2}); err == nil { t.Fatal("out-of-bounds write accepted") }
}

func TestMachineDeterministicExecution(t *testing.T) {
	m, err := NewMachine(8)
	if err != nil { t.Fatal(err) }
	r, err := m.Execute(OpWrite, 0, 0, []byte{42})
	if err != nil { t.Fatal(err) }
	if r.PCBefore != 0 || r.PCAfter != 1 || !r.Changed { t.Fatalf("bad write result: %#v", r) }
	r, err = m.Execute(OpRead, 0, 1, nil)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(r.Value, []byte{42}) || r.PCBefore != 1 || r.PCAfter != 2 { t.Fatalf("bad read result: %#v", r) }
}

func TestMachineRejectsUnknownInstruction(t *testing.T) {
	m, err := NewMachine(1)
	if err != nil { t.Fatal(err) }
	if _, err := m.Execute(Instruction("RUN_ARBITRARY_PAYLOAD"), 0, 0, nil); err == nil {
		t.Fatal("unknown instruction accepted")
	}
}
