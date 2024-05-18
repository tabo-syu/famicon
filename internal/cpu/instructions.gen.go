// Code generated by cmd/opcode_scraper/main.go; DO NOT EDIT.
package cpu

import "errors"

type opcode func(mode addressingMode)

type instruction struct {
	opcode string
	bytes  uint16
	cycles uint16
	mode   addressingMode
}

func (i instruction) Call(cpu *CPU) error {
	var err error

	switch i.opcode {
	// case "ADC":
	// 	err = cpu.ADC(i.mode)
	case "AND":
		err = cpu.AND(i.mode)
	// case "ASL":
	// 	err = cpu.ASL(i.mode)
	case "BCC":
		err = cpu.BCC(i.mode)
	case "BCS":
		err = cpu.BCS(i.mode)
	case "BEQ":
		err = cpu.BEQ(i.mode)
	// case "BIT":
	// 	err = cpu.BIT(i.mode)
	case "BMI":
		err = cpu.BMI(i.mode)
	case "BNE":
		err = cpu.BNE(i.mode)
	case "BPL":
		err = cpu.BPL(i.mode)
	case "BRK":
		err = cpu.BRK(i.mode)
	// case "BVC":
	// 	err = cpu.BVC(i.mode)
	// case "BVS":
	// 	err = cpu.BVS(i.mode)
	case "CLC":
		err = cpu.CLC(i.mode)
	case "CLD":
		err = cpu.CLD(i.mode)
	case "CLI":
		err = cpu.CLI(i.mode)
	case "CLV":
		err = cpu.CLV(i.mode)
	case "CMP":
		err = cpu.CMP(i.mode)
	case "CPX":
		err = cpu.CPX(i.mode)
	case "CPY":
		err = cpu.CPY(i.mode)
	case "DEC":
		err = cpu.DEC(i.mode)
	case "DEX":
		err = cpu.DEX(i.mode)
	case "DEY":
		err = cpu.DEY(i.mode)
	case "EOR":
		err = cpu.EOR(i.mode)
	case "INC":
		err = cpu.INC(i.mode)
	case "INX":
		err = cpu.INX(i.mode)
	case "INY":
		err = cpu.INY(i.mode)
	case "JMP":
		err = cpu.JMP(i.mode)
	// case "JSR":
	// 	err = cpu.JSR(i.mode)
	case "LDA":
		err = cpu.LDA(i.mode)
	case "LDX":
		err = cpu.LDX(i.mode)
	case "LDY":
		err = cpu.LDY(i.mode)
	// case "LSR":
	// 	err = cpu.LSR(i.mode)
	// case "NOP":
	// 	err = cpu.NOP(i.mode)
	case "ORA":
		err = cpu.ORA(i.mode)
	// case "PHA":
	// 	err = cpu.PHA(i.mode)
	// case "PHP":
	// 	err = cpu.PHP(i.mode)
	// case "PLA":
	// 	err = cpu.PLA(i.mode)
	// case "PLP":
	// 	err = cpu.PLP(i.mode)
	// case "ROL":
	// 	err = cpu.ROL(i.mode)
	// case "ROR":
	// 	err = cpu.ROR(i.mode)
	// case "RTI":
	// 	err = cpu.RTI(i.mode)
	// case "RTS":
	// 	err = cpu.RTS(i.mode)
	// case "SBC":
	// 	err = cpu.SBC(i.mode)
	case "SEC":
		err = cpu.SEC(i.mode)
	case "SED":
		err = cpu.SED(i.mode)
	case "SEI":
		err = cpu.SEI(i.mode)
	case "STA":
		err = cpu.STA(i.mode)
	case "STX":
		err = cpu.STX(i.mode)
	case "STY":
		err = cpu.STY(i.mode)
	case "TAX":
		err = cpu.TAX(i.mode)
	case "TAY":
		err = cpu.TAY(i.mode)
	case "TSX":
		err = cpu.TSX(i.mode)
	case "TXA":
		err = cpu.TXA(i.mode)
	case "TXS":
		err = cpu.TXS(i.mode)
	case "TYA":
		err = cpu.TYA(i.mode)
	default:
		return errors.New("unexpected opcode")
	}

	cpu.programCounter += i.bytes - 1

	return err
}

func newInstruction(opcode string, bytes uint16, cycles uint16, mode addressingMode) instruction {
	return instruction{
		opcode: opcode,
		bytes:  bytes,
		cycles: cycles,
		mode:   mode,
	}
}

func NewInstructions() map[uint8]instruction {
	return map[uint8]instruction{
		// ADC
		0x69: newInstruction("ADC", 2, 2, ImmediateMode),
		0x65: newInstruction("ADC", 2, 3, ZeroPageMode),
		0x75: newInstruction("ADC", 2, 4, ZeroPageXMode),
		0x6D: newInstruction("ADC", 3, 4, AbsoluteMode),
		0x7D: newInstruction("ADC", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0x79: newInstruction("ADC", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0x61: newInstruction("ADC", 2, 6, IndirectXMode),
		0x71: newInstruction("ADC", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// AND
		0x29: newInstruction("AND", 2, 2, ImmediateMode),
		0x25: newInstruction("AND", 2, 3, ZeroPageMode),
		0x35: newInstruction("AND", 2, 4, ZeroPageXMode),
		0x2D: newInstruction("AND", 3, 4, AbsoluteMode),
		0x3D: newInstruction("AND", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0x39: newInstruction("AND", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0x21: newInstruction("AND", 2, 6, IndirectXMode),
		0x31: newInstruction("AND", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// ASL
		0x0A: newInstruction("ASL", 1, 2, AccumulatorMode),
		0x06: newInstruction("ASL", 2, 5, ZeroPageMode),
		0x16: newInstruction("ASL", 2, 6, ZeroPageXMode),
		0x0E: newInstruction("ASL", 3, 6, AbsoluteMode),
		0x1E: newInstruction("ASL", 3, 7, AbsoluteXMode),
		// BCC
		0x90: newInstruction("BCC", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BCS
		0xB0: newInstruction("BCS", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BEQ
		0xF0: newInstruction("BEQ", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BIT
		0x24: newInstruction("BIT", 2, 3, ZeroPageMode),
		0x2C: newInstruction("BIT", 3, 4, AbsoluteMode),
		// BMI
		0x30: newInstruction("BMI", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BNE
		0xD0: newInstruction("BNE", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BPL
		0x10: newInstruction("BPL", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BRK
		0x00: newInstruction("BRK", 1, 7, ImpliedMode),
		// BVC
		0x50: newInstruction("BVC", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// BVS
		0x70: newInstruction("BVS", 2, 2 /*(+1 if branch succeeds, +2 if to a new page)*/, RelativeMode),
		// CLC
		0x18: newInstruction("CLC", 1, 2, ImpliedMode),
		// CLD
		0xD8: newInstruction("CLD", 1, 2, ImpliedMode),
		// CLI
		0x58: newInstruction("CLI", 1, 2, ImpliedMode),
		// CLV
		0xB8: newInstruction("CLV", 1, 2, ImpliedMode),
		// CMP
		0xC9: newInstruction("CMP", 2, 2, ImmediateMode),
		0xC5: newInstruction("CMP", 2, 3, ZeroPageMode),
		0xD5: newInstruction("CMP", 2, 4, ZeroPageXMode),
		0xCD: newInstruction("CMP", 3, 4, AbsoluteMode),
		0xDD: newInstruction("CMP", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0xD9: newInstruction("CMP", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0xC1: newInstruction("CMP", 2, 6, IndirectXMode),
		0xD1: newInstruction("CMP", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// CPX
		0xE0: newInstruction("CPX", 2, 2, ImmediateMode),
		0xE4: newInstruction("CPX", 2, 3, ZeroPageMode),
		0xEC: newInstruction("CPX", 3, 4, AbsoluteMode),
		// CPY
		0xC0: newInstruction("CPY", 2, 2, ImmediateMode),
		0xC4: newInstruction("CPY", 2, 3, ZeroPageMode),
		0xCC: newInstruction("CPY", 3, 4, AbsoluteMode),
		// DEC
		0xC6: newInstruction("DEC", 2, 5, ZeroPageMode),
		0xD6: newInstruction("DEC", 2, 6, ZeroPageXMode),
		0xCE: newInstruction("DEC", 3, 6, AbsoluteMode),
		0xDE: newInstruction("DEC", 3, 7, AbsoluteXMode),
		// DEX
		0xCA: newInstruction("DEX", 1, 2, ImpliedMode),
		// DEY
		0x88: newInstruction("DEY", 1, 2, ImpliedMode),
		// EOR
		0x49: newInstruction("EOR", 2, 2, ImmediateMode),
		0x45: newInstruction("EOR", 2, 3, ZeroPageMode),
		0x55: newInstruction("EOR", 2, 4, ZeroPageXMode),
		0x4D: newInstruction("EOR", 3, 4, AbsoluteMode),
		0x5D: newInstruction("EOR", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0x59: newInstruction("EOR", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0x41: newInstruction("EOR", 2, 6, IndirectXMode),
		0x51: newInstruction("EOR", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// INC
		0xE6: newInstruction("INC", 2, 5, ZeroPageMode),
		0xF6: newInstruction("INC", 2, 6, ZeroPageXMode),
		0xEE: newInstruction("INC", 3, 6, AbsoluteMode),
		0xFE: newInstruction("INC", 3, 7, AbsoluteXMode),
		// INX
		0xE8: newInstruction("INX", 1, 2, ImpliedMode),
		// INY
		0xC8: newInstruction("INY", 1, 2, ImpliedMode),
		// JMP
		0x4C: newInstruction("JMP", 3, 3, AbsoluteMode),
		0x6C: newInstruction("JMP", 3, 5, IndirectMode),
		// JSR
		0x20: newInstruction("JSR", 3, 6, AbsoluteMode),
		// LDA
		0xA9: newInstruction("LDA", 2, 2, ImmediateMode),
		0xA5: newInstruction("LDA", 2, 3, ZeroPageMode),
		0xB5: newInstruction("LDA", 2, 4, ZeroPageXMode),
		0xAD: newInstruction("LDA", 3, 4, AbsoluteMode),
		0xBD: newInstruction("LDA", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0xB9: newInstruction("LDA", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0xA1: newInstruction("LDA", 2, 6, IndirectXMode),
		0xB1: newInstruction("LDA", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// LDX
		0xA2: newInstruction("LDX", 2, 2, ImmediateMode),
		0xA6: newInstruction("LDX", 2, 3, ZeroPageMode),
		0xB6: newInstruction("LDX", 2, 4, ZeroPageYMode),
		0xAE: newInstruction("LDX", 3, 4, AbsoluteMode),
		0xBE: newInstruction("LDX", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		// LDY
		0xA0: newInstruction("LDY", 2, 2, ImmediateMode),
		0xA4: newInstruction("LDY", 2, 3, ZeroPageMode),
		0xB4: newInstruction("LDY", 2, 4, ZeroPageXMode),
		0xAC: newInstruction("LDY", 3, 4, AbsoluteMode),
		0xBC: newInstruction("LDY", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		// LSR
		0x4A: newInstruction("LSR", 1, 2, AccumulatorMode),
		0x46: newInstruction("LSR", 2, 5, ZeroPageMode),
		0x56: newInstruction("LSR", 2, 6, ZeroPageXMode),
		0x4E: newInstruction("LSR", 3, 6, AbsoluteMode),
		0x5E: newInstruction("LSR", 3, 7, AbsoluteXMode),
		// NOP
		0xEA: newInstruction("NOP", 1, 2, ImpliedMode),
		// ORA
		0x09: newInstruction("ORA", 2, 2, ImmediateMode),
		0x05: newInstruction("ORA", 2, 3, ZeroPageMode),
		0x15: newInstruction("ORA", 2, 4, ZeroPageXMode),
		0x0D: newInstruction("ORA", 3, 4, AbsoluteMode),
		0x1D: newInstruction("ORA", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0x19: newInstruction("ORA", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0x01: newInstruction("ORA", 2, 6, IndirectXMode),
		0x11: newInstruction("ORA", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// PHA
		0x48: newInstruction("PHA", 1, 3, ImpliedMode),
		// PHP
		0x08: newInstruction("PHP", 1, 3, ImpliedMode),
		// PLA
		0x68: newInstruction("PLA", 1, 4, ImpliedMode),
		// PLP
		0x28: newInstruction("PLP", 1, 4, ImpliedMode),
		// ROL
		0x2A: newInstruction("ROL", 1, 2, AccumulatorMode),
		0x26: newInstruction("ROL", 2, 5, ZeroPageMode),
		0x36: newInstruction("ROL", 2, 6, ZeroPageXMode),
		0x2E: newInstruction("ROL", 3, 6, AbsoluteMode),
		0x3E: newInstruction("ROL", 3, 7, AbsoluteXMode),
		// ROR
		0x6A: newInstruction("ROR", 1, 2, AccumulatorMode),
		0x66: newInstruction("ROR", 2, 5, ZeroPageMode),
		0x76: newInstruction("ROR", 2, 6, ZeroPageXMode),
		0x6E: newInstruction("ROR", 3, 6, AbsoluteMode),
		0x7E: newInstruction("ROR", 3, 7, AbsoluteXMode),
		// RTI
		0x40: newInstruction("RTI", 1, 6, ImpliedMode),
		// RTS
		0x60: newInstruction("RTS", 1, 6, ImpliedMode),
		// SBC
		0xE9: newInstruction("SBC", 2, 2, ImmediateMode),
		0xE5: newInstruction("SBC", 2, 3, ZeroPageMode),
		0xF5: newInstruction("SBC", 2, 4, ZeroPageXMode),
		0xED: newInstruction("SBC", 3, 4, AbsoluteMode),
		0xFD: newInstruction("SBC", 3, 4 /*(+1 if page crossed)*/, AbsoluteXMode),
		0xF9: newInstruction("SBC", 3, 4 /*(+1 if page crossed)*/, AbsoluteYMode),
		0xE1: newInstruction("SBC", 2, 6, IndirectXMode),
		0xF1: newInstruction("SBC", 2, 5 /*(+1 if page crossed)*/, IndirectYMode),
		// SEC
		0x38: newInstruction("SEC", 1, 2, ImpliedMode),
		// SED
		0xF8: newInstruction("SED", 1, 2, ImpliedMode),
		// SEI
		0x78: newInstruction("SEI", 1, 2, ImpliedMode),
		// STA
		0x85: newInstruction("STA", 2, 3, ZeroPageMode),
		0x95: newInstruction("STA", 2, 4, ZeroPageXMode),
		0x8D: newInstruction("STA", 3, 4, AbsoluteMode),
		0x9D: newInstruction("STA", 3, 5, AbsoluteXMode),
		0x99: newInstruction("STA", 3, 5, AbsoluteYMode),
		0x81: newInstruction("STA", 2, 6, IndirectXMode),
		0x91: newInstruction("STA", 2, 6, IndirectYMode),
		// STX
		0x86: newInstruction("STX", 2, 3, ZeroPageMode),
		0x96: newInstruction("STX", 2, 4, ZeroPageYMode),
		0x8E: newInstruction("STX", 3, 4, AbsoluteMode),
		// STY
		0x84: newInstruction("STY", 2, 3, ZeroPageMode),
		0x94: newInstruction("STY", 2, 4, ZeroPageXMode),
		0x8C: newInstruction("STY", 3, 4, AbsoluteMode),
		// TAX
		0xAA: newInstruction("TAX", 1, 2, ImpliedMode),
		// TAY
		0xA8: newInstruction("TAY", 1, 2, ImpliedMode),
		// TSX
		0xBA: newInstruction("TSX", 1, 2, ImpliedMode),
		// TXA
		0x8A: newInstruction("TXA", 1, 2, ImpliedMode),
		// TXS
		0x9A: newInstruction("TXS", 1, 2, ImpliedMode),
		// TYA
		0x98: newInstruction("TYA", 1, 2, ImpliedMode),
	}
}
