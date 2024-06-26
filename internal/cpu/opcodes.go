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

func (cpu *CPU) ADC(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	var carry byte
	if cpu.status.c() {
		carry = 1
	}

	aSign := uint16(cpu.registerA) & 0b1000_0000
	vSign := uint16(value) & 0b1000_0000
	mayOverflow := aSign^vSign == 0b0000_0000

	result := uint16(cpu.registerA) + uint16(value) + uint16(carry)

	rSign := result & 0b1000_0000
	isDiffSign := rSign != aSign
	if mayOverflow && isDiffSign {
		cpu.status.setO(true)
	}
	if result > 0xFF {
		cpu.status.setC(true)
	}

	cpu.registerA = byte(result & 0xFF)
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) AND(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	result := cpu.registerA & value
	cpu.registerA = result
	cpu.updateZeroAndNegativeFlags(result)

	return nil
}

func (cpu *CPU) ASL(mode addressingMode) error {
	var (
		value   byte
		address uint16
	)
	if mode == AccumulatorMode {
		value = cpu.registerA
	} else {
		address = cpu.getOperandAddress(mode)
		value = cpu.Bus.ReadMemory(address)
	}

	cpu.status.setC(value&0b1000_0000 != 0)
	value = value << 1

	if mode == AccumulatorMode {
		cpu.registerA = value
	} else {
		cpu.Bus.WriteMemory(address, value)
	}

	cpu.updateZeroFlag(cpu.registerA)
	cpu.updateNegativeFlag(value)

	return nil
}

func (cpu *CPU) BCC(mode addressingMode) error {
	if !cpu.status.c() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BCS(mode addressingMode) error {
	if cpu.status.c() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BEQ(mode addressingMode) error {
	if cpu.status.z() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BIT(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	result := cpu.registerA & value

	isOverflow := (result & byte(0b0100_0000) >> 6) == 1
	cpu.status.setO(isOverflow)

	isNegative := (result & byte(0b1000_0000) >> 7) == 1
	cpu.status.setN(isNegative)

	cpu.updateZeroFlag(result)

	return nil
}

func (cpu *CPU) BMI(mode addressingMode) error {
	if cpu.status.n() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BNE(mode addressingMode) error {
	if !cpu.status.z() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BPL(mode addressingMode) error {
	if !cpu.status.n() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BRK(mode addressingMode) error {
	return errors.New("BRK called")
}

func (cpu *CPU) BVC(mode addressingMode) error {
	if !cpu.status.o() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) BVS(mode addressingMode) error {
	if cpu.status.o() {
		address := cpu.getOperandAddress(mode)
		cpu.ProgramCounter = address
	}

	return nil
}

func (cpu *CPU) CLC(mode addressingMode) error {
	cpu.status.setC(false)

	return nil
}

func (cpu *CPU) CLD(mode addressingMode) error {
	cpu.status.setD(false)

	return nil
}

func (cpu *CPU) CLI(mode addressingMode) error {
	cpu.status.setI(false)

	return nil
}

func (cpu *CPU) CLV(mode addressingMode) error {
	cpu.status.setO(false)

	return nil
}

func (cpu *CPU) CMP(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	value := cpu.Bus.ReadMemory(address)
	if cpu.registerA == value {
		cpu.status.setZ(true)
	}
	if cpu.registerA >= value {
		cpu.status.setC(true)
	}
	if value&0b1000_0000 != 0 {
		cpu.status.setN(true)
	} else {
		cpu.status.setN(false)
	}

	return nil
}

func (cpu *CPU) CPX(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	value := cpu.Bus.ReadMemory(address)
	if cpu.registerX == value {
		cpu.status.setZ(true)
	}
	if cpu.registerX >= value {
		cpu.status.setC(true)
	}
	if value&0b1000_0000 != 0 {
		cpu.status.setN(true)
	} else {
		cpu.status.setN(false)
	}

	return nil
}

func (cpu *CPU) CPY(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	value := cpu.Bus.ReadMemory(address)
	if cpu.registerY == value {
		cpu.status.setZ(true)
	}
	if cpu.registerY >= value {
		cpu.status.setC(true)
	}
	if value&0b1000_0000 != 0 {
		cpu.status.setN(true)
	} else {
		cpu.status.setN(false)
	}

	return nil
}

func (cpu *CPU) DEC(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	value := cpu.Bus.ReadMemory(address)
	value--
	cpu.Bus.WriteMemory(address, value)

	cpu.updateZeroAndNegativeFlags(value)

	return nil
}

func (cpu *CPU) DEX(mode addressingMode) error {
	cpu.registerX--
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) DEY(mode addressingMode) error {
	cpu.registerY--
	cpu.updateZeroAndNegativeFlags(cpu.registerY)

	return nil
}

func (cpu *CPU) EOR(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	result := cpu.registerA ^ value
	cpu.registerA = result
	cpu.updateZeroAndNegativeFlags(result)

	return nil
}

func (cpu *CPU) INC(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	value := cpu.Bus.ReadMemory(address)
	value++
	cpu.Bus.WriteMemory(address, value)

	cpu.updateZeroAndNegativeFlags(value)

	return nil
}

func (cpu *CPU) INX(mode addressingMode) error {
	cpu.registerX++
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) INY(mode addressingMode) error {
	cpu.registerY++
	cpu.updateZeroAndNegativeFlags(cpu.registerY)

	return nil
}

func (cpu *CPU) JMP(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	cpu.ProgramCounter = address - 2

	return nil
}

func (cpu *CPU) JSR(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)

	cpu.pushStackUint16(cpu.ProgramCounter + 2 - 1)
	cpu.ProgramCounter = address - 2

	return nil
}

func (cpu *CPU) LDA(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	cpu.registerA = value
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) LDX(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	cpu.registerX = value
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) LDY(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	cpu.registerY = value
	cpu.updateZeroAndNegativeFlags(cpu.registerY)

	return nil
}

func (cpu *CPU) LSR(mode addressingMode) error {
	var (
		value   byte
		address uint16
	)
	if mode == AccumulatorMode {
		value = cpu.registerA
	} else {
		address = cpu.getOperandAddress(mode)
		value = cpu.Bus.ReadMemory(address)
	}

	cpu.status.setC(value&0b0000_0001 != 0)
	value = value >> 1

	if mode == AccumulatorMode {
		cpu.registerA = value
	} else {
		cpu.Bus.WriteMemory(address, value)
	}

	cpu.updateZeroAndNegativeFlags(value)

	return nil
}

func (cpu *CPU) NOP(mode addressingMode) error {
	return nil
}

func (cpu *CPU) ORA(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	result := cpu.registerA | value
	cpu.registerA = result
	cpu.updateZeroAndNegativeFlags(result)

	return nil
}

func (cpu *CPU) PHA(mode addressingMode) error {
	cpu.pushStack(cpu.registerA)

	return nil
}

func (cpu *CPU) PHP(mode addressingMode) error {
	cpu.pushStack(byte(cpu.status))

	return nil
}

func (cpu *CPU) PLA(mode addressingMode) error {
	cpu.registerA = cpu.popStack()
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) PLP(mode addressingMode) error {
	cpu.status = status(cpu.popStack())

	return nil
}

func (cpu *CPU) ROL(mode addressingMode) error {
	var (
		value   byte
		address uint16
	)
	if mode == AccumulatorMode {
		value = cpu.registerA
	} else {
		address = cpu.getOperandAddress(mode)
		value = cpu.Bus.ReadMemory(address)
	}

	var new0Bit = 0
	if cpu.status.c() {
		new0Bit = 1
	}
	cpu.status.setC(value&0b1000_0000 != 0)
	value = value<<1 | byte(new0Bit)

	if mode == AccumulatorMode {
		cpu.registerA = value
	} else {
		cpu.Bus.WriteMemory(address, value)
	}

	cpu.updateZeroAndNegativeFlags(value)

	return nil
}

func (cpu *CPU) ROR(mode addressingMode) error {
	var (
		value   byte
		address uint16
	)
	if mode == AccumulatorMode {
		value = cpu.registerA
	} else {
		address = cpu.getOperandAddress(mode)
		value = cpu.Bus.ReadMemory(address)
	}

	var new7Bit = 0b0000_0000
	if cpu.status.c() {
		new7Bit = 0b1000_0000
	}
	cpu.status.setC(value&0b0000_0001 != 0)
	value = value>>1 | byte(new7Bit)

	if mode == AccumulatorMode {
		cpu.registerA = value
	} else {
		cpu.Bus.WriteMemory(address, value)
	}

	cpu.updateZeroFlag(cpu.registerA)
	cpu.updateNegativeFlag(value)

	return nil
}

func (cpu *CPU) RTI(mode addressingMode) error {
	cpu.status = status(cpu.popStack())
	cpu.ProgramCounter = cpu.popStackUint16() + 1

	return nil
}

func (cpu *CPU) RTS(mode addressingMode) error {
	cpu.ProgramCounter = cpu.popStackUint16() + 1

	return nil
}

func (cpu *CPU) SBC(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	value := cpu.Bus.ReadMemory(address)

	var carry byte
	if cpu.status.c() {
		carry = 1
	}

	aSign := int16(cpu.registerA) & 0b1000_0000
	vSign := int16(value) & 0b1000_0000
	mayOverflow := aSign^vSign == 0b1000_0000

	result := int16(cpu.registerA) - int16(value) - int16(1-carry)

	rSign := result & 0b1000_0000
	isDiffSign := rSign != aSign
	isOverflow := mayOverflow && isDiffSign
	if isOverflow {
		cpu.status.setO(true)
	}
	if result >= 0 {
		cpu.status.setC(true)
	}

	cpu.registerA = byte(result & 0xFF)
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) SEC(mode addressingMode) error {
	cpu.status.setC(true)

	return nil
}

func (cpu *CPU) SED(mode addressingMode) error {
	cpu.status.setD(true)

	return nil
}

func (cpu *CPU) SEI(mode addressingMode) error {
	cpu.status.setI(true)

	return nil
}

func (cpu *CPU) STA(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	cpu.Bus.WriteMemory(address, cpu.registerA)

	return nil
}

func (cpu *CPU) STX(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	cpu.Bus.WriteMemory(address, cpu.registerX)

	return nil
}

func (cpu *CPU) STY(mode addressingMode) error {
	address := cpu.getOperandAddress(mode)
	cpu.Bus.WriteMemory(address, cpu.registerY)

	return nil
}

func (cpu *CPU) TAX(mode addressingMode) error {
	cpu.registerX = cpu.registerA
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) TAY(mode addressingMode) error {
	cpu.registerY = cpu.registerA
	cpu.updateZeroAndNegativeFlags(cpu.registerY)

	return nil
}

func (cpu *CPU) TSX(mode addressingMode) error {
	cpu.registerX = byte(cpu.stackPointer)
	cpu.updateZeroAndNegativeFlags(cpu.registerX)

	return nil
}

func (cpu *CPU) TXA(mode addressingMode) error {
	cpu.registerA = cpu.registerX
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) TXS(mode addressingMode) error {
	cpu.stackPointer = stackPointer(cpu.registerX)

	return nil
}

func (cpu *CPU) TYA(mode addressingMode) error {
	cpu.registerA = cpu.registerY
	cpu.updateZeroAndNegativeFlags(cpu.registerA)

	return nil
}

func (cpu *CPU) getOperandAddress(mode addressingMode) uint16 {
	switch mode {
	case ImmediateMode:
		return cpu.ProgramCounter

	case ZeroPageMode:
		return uint16(cpu.Bus.ReadMemory(cpu.ProgramCounter))

	case ZeroPageXMode:
		position := cpu.Bus.ReadMemory(cpu.ProgramCounter)
		address := uint16(position + cpu.registerX)

		return address

	case ZeroPageYMode:
		position := cpu.Bus.ReadMemory(cpu.ProgramCounter)
		address := uint16(position + cpu.registerY)

		return address

	case RelativeMode:
		operand := int8(cpu.Bus.ReadMemory(cpu.ProgramCounter))
		address := uint16(int32(cpu.ProgramCounter) + int32(operand))

		return address

	case AbsoluteMode:
		return cpu.Bus.ReadMemoryUint16(cpu.ProgramCounter)

	case AbsoluteXMode:
		base := cpu.Bus.ReadMemoryUint16(cpu.ProgramCounter)
		address := base + uint16(cpu.registerX)

		return address

	case AbsoluteYMode:
		base := cpu.Bus.ReadMemoryUint16(cpu.ProgramCounter)
		address := base + uint16(cpu.registerY)

		return address

	case IndirectMode:
		base := cpu.Bus.ReadMemoryUint16(cpu.ProgramCounter)
		address := cpu.Bus.ReadMemoryUint16(base)

		return address

	case IndirectXMode:
		base := cpu.Bus.ReadMemory(cpu.ProgramCounter)
		pointer := base + cpu.registerX
		low := cpu.Bus.ReadMemory(uint16(pointer))
		high := cpu.Bus.ReadMemory(uint16(pointer + 1))
		address := uint16(high)<<8 | uint16(low)

		return address

	case IndirectYMode:
		base := cpu.Bus.ReadMemory(cpu.ProgramCounter)
		low := cpu.Bus.ReadMemory(uint16(base))
		high := cpu.Bus.ReadMemory(uint16(base + 1))
		derefBase := uint16(high)<<8 | uint16(low)
		deref := derefBase + uint16(cpu.registerY)

		return deref

	case NoneAddressingMode:
		panic("through `NoneAddressing`")

	default:
		return 0
	}
}

func (cpu *CPU) pushStack(value byte) {
	cpu.Bus.WriteMemory(cpu.stackPointer.toAddress(), value)
	cpu.stackPointer--
}

func (cpu *CPU) popStack() byte {
	cpu.stackPointer++
	value := cpu.Bus.ReadMemory(cpu.stackPointer.toAddress())

	return value
}

func (cpu *CPU) pushStackUint16(value uint16) {
	high := byte(value >> 8)
	low := byte(value & 0x00_FF)

	cpu.pushStack(high)
	cpu.pushStack(low)
}

func (cpu *CPU) popStackUint16() uint16 {
	low := uint16(cpu.popStack())
	high := uint16(cpu.popStack())

	return high<<8 | low
}

func (cpu *CPU) updateZeroFlag(value byte) {
	if value == 0 {
		cpu.status.setZ(true)
	} else {
		cpu.status.setZ(false)
	}
}

func (cpu *CPU) updateNegativeFlag(value byte) {
	if value&0b1000_0000 != 0 {
		cpu.status.setN(true)
	} else {
		cpu.status.setN(false)
	}
}

func (cpu *CPU) updateZeroAndNegativeFlags(value byte) {
	cpu.updateZeroFlag(value)
	cpu.updateNegativeFlag(value)
}
