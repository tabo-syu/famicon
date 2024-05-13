package cpu

import "errors"

type addressingMode int

const (
	ImpliedMode addressingMode = iota
	AccumulatorMode
	ImmediateMode
	ZeroPageMode
	ZeroPageXMode
	ZeroPageYMode
	RelativeMode
	AbsoluteMode
	AbsoluteXMode
	AbsoluteYMode
	IndirectMode
	IndirectXMode
	IndirectYMode
	NoneAddressingMode
)

func (cpu *CPU) BRK(mode addressingMode) error {
	return errors.New("BRK called")
}

func (cpu *CPU) STA(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	cpu.memory.Write(address, cpu.registerA)

	return nil
}

func (cpu *CPU) LDA(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.memory.Read(address)

	cpu.registerA = value
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) TAX(mode addressingMode) error {
	cpu.registerX = cpu.registerA
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) INX(mode addressingMode) error {
	cpu.registerX++
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) getOperandAddress(mode addressingMode) uint16 {
	switch mode {
	case ImmediateMode:
		return cpu.programCounter

	case ZeroPageMode:
		return uint16(cpu.memory.Read(cpu.programCounter))

	case ZeroPageXMode:
		position := cpu.memory.Read(cpu.programCounter)
		address := uint16(position + cpu.registerX)

		return address

	case ZeroPageYMode:
		position := cpu.memory.Read(cpu.programCounter)
		address := uint16(position + cpu.registerY)

		return address

	case AbsoluteMode:
		return cpu.memory.ReadUint16(cpu.programCounter)

	case AbsoluteXMode:
		base := cpu.memory.ReadUint16(cpu.programCounter)
		address := base + uint16(cpu.registerX)

		return address

	case AbsoluteYMode:
		base := cpu.memory.ReadUint16(cpu.programCounter)
		address := base + uint16(cpu.registerY)

		return address

	case IndirectXMode:
		base := cpu.memory.Read(cpu.programCounter)
		pointer := base + cpu.registerX
		low := cpu.memory.Read(uint16(pointer))
		high := cpu.memory.Read(uint16(pointer + 1))
		address := uint16(high)<<8 | uint16(low)

		return address

	case IndirectYMode:
		base := cpu.memory.Read(cpu.programCounter)
		low := cpu.memory.Read(uint16(base))
		high := cpu.memory.Read(uint16(base + 1))
		derefBase := uint16(high)<<8 | uint16(low)
		deref := derefBase + uint16(cpu.registerY)

		return deref

	case NoneAddressingMode:
		panic("through `NoneAddressing`")

	default:
		return 0
	}
}

func (cpu *CPU) updateZeroAndNegativeFlags(register uint8) {
	if register == 0 {
		cpu.status.setZ(true)
	} else {
		cpu.status.setZ(false)
	}

	if register&0b0100_0000 != 0 {
		cpu.status.setN(true)
	} else {
		cpu.status.setN(false)
	}
}