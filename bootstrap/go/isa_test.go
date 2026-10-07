package zero

import (
	"bytes"
	"testing"
)

func TestMachineRejectsZeroSize(t *testing.T) {
	if _, err := NewMachine(0); err == nil {
		t.Fatal("zero-size machine accepted")
	}
}

func TestMachineReadReturnsCopy(t *testing.T) {
	m, err := NewMachine(4)
	if err != nil { t.Fatal(err) }
	if err := m.Write(0, []byte{1, 2}); err != nil { t.Fatal(err) }
	got, err := m.Read(0, 2)
	if err != nil { t.Fatal(err) }
	got[0] = 99
	again, err := m.Read(0, 2)
	if err != nil { t.Fatal(err) }
	if !bytes.Equal(again, []byte{1, 2}) {
		t.Fatalf("read exposed machine memory: %v", again)
	}
}

func TestMachineRejectedInstructionDoesNotAdvancePC(t *testing.T) {
	m, err := NewMachine(1)
	if err != nil { t.Fatal(err) }
	before := m.PC
	if _, err := m.Execute(Instruction("INVALID"), 0, 0, nil); err == nil {
		t.Fatal("invalid instruction accepted")
	}
	if m.PC != before {
		t.Fatalf("rejected instruction changed PC: %d -> %d", before, m.PC)
	}
}

func TestMachineFailedMemoryOperationDoesNotAdvancePC(t *testing.T) {
	m, err := NewMachine(4)
	if err != nil { t.Fatal(err) }
	before := m.PC
	if _, err := m.Execute(OpRead, 3, 2, nil); err == nil {
		t.Fatal("out-of-bounds read accepted")
	}
	if m.PC != before {
		t.Fatalf("failed operation changed PC: %d -> %d", before, m.PC)
	}
	if _, err := m.Execute(OpWrite, 3, 2, []byte{1, 2}); err == nil {
		t.Fatal("out-of-bounds write accepted")
	}
	if m.PC != before {
		t.Fatalf("failed write changed PC: %d -> %d", before, m.PC)
	}
}


func TestMachinePCOverflowIsDeterministic(t *testing.T) {
	m, err := NewMachine(1)
	if err != nil { t.Fatal(err) }
	m.PC = ^uint64(0)
	result, err := m.Execute(OpStep, 0, 0, nil)
	if err != nil { t.Fatal(err) }
	if result.PCBefore != ^uint64(0) || result.PCAfter != 0 || m.PC != 0 {
		t.Fatalf("unexpected uint64 PC wrap semantics: %+v PC=%d", result, m.PC)
	}
}
