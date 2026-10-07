package zero

import "errors"

type Instruction string

const (
	OpObserve   Instruction = "OBSERVE"
	OpRead      Instruction = "READ"
	OpWrite     Instruction = "WRITE"
	OpStep      Instruction = "STEP"
	OpRequest   Instruction = "REQUEST"
	OpNextEvent Instruction = "NEXT_EVENT"
)

type Machine struct {
	Memory []byte
	PC     uint64
}

type ExecutionResult struct {
	Instruction Instruction
	Value       []byte
	PCBefore    uint64
	PCAfter     uint64
	Changed     bool
}

func NewMachine(size uint64) (*Machine, error) {
	if size == 0 {
		return nil, errors.New("memory size must be non-zero")
	}
	return &Machine{Memory: make([]byte, size)}, nil
}

func (m *Machine) checkRange(address, length uint64) error {
	if address > uint64(len(m.Memory)) || length > uint64(len(m.Memory))-address {
		return errors.New("memory range out of bounds")
	}
	return nil
}

func (m *Machine) Read(address, length uint64) ([]byte, error) {
	if err := m.checkRange(address, length); err != nil {
		return nil, err
	}
	out := make([]byte, length)
	copy(out, m.Memory[address:address+length])
	return out, nil
}

func (m *Machine) Write(address uint64, data []byte) error {
	if err := m.checkRange(address, uint64(len(data))); err != nil {
		return err
	}
	copy(m.Memory[address:address+uint64(len(data))], data)
	return nil
}

func (m *Machine) Execute(op Instruction, address, length uint64, data []byte) (ExecutionResult, error) {
	before := m.PC
	switch op {
	case OpRead:
		value, err := m.Read(address, length)
		if err != nil { return ExecutionResult{}, err }
		m.PC++
		return ExecutionResult{Instruction: op, Value: value, PCBefore: before, PCAfter: m.PC}, nil
	case OpWrite:
		if err := m.Write(address, data); err != nil { return ExecutionResult{}, err }
		m.PC++
		return ExecutionResult{Instruction: op, PCBefore: before, PCAfter: m.PC, Changed: len(data) != 0}, nil
	case OpStep:
		m.PC++
		return ExecutionResult{Instruction: op, PCBefore: before, PCAfter: m.PC}, nil
	case OpObserve, OpRequest, OpNextEvent:
		m.PC++
		return ExecutionResult{Instruction: op, PCBefore: before, PCAfter: m.PC}, nil
	default:
		return ExecutionResult{}, errors.New("unknown instruction")
	}
}
