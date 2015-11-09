package main

import "fmt"

type Cpu struct {
	// Registers
	pc, sp           int // Program counter, Stack pointer
	a                int // Accumulator, Flag
	b, c, d, e, h, l int // Auxiliary registers
	p                FlagReg

	// Memory
	m Memory

	// Next instruction to execute
	nextInstr Instruction
}

type FlagReg struct {
	z, n, h, c int
}

const (
	A = iota
	B
	C
	D
	E
	H
	L
	BC
	HL
	DE
	SP
)

func (f *FlagReg) toInt() int {
	return f.z<<7 | f.n<<6 | f.h<<5 | f.c<<4
}

func fromInt(n int) FlagReg {
	f := FlagReg{
		z: (n & BIT_7) >> 7,
		n: (n & BIT_6) >> 6,
		h: (n & BIT_5) >> 5,
		c: (n & BIT_4) >> 4,
	}

	return f
}

type Mem interface {
	Read(addr int) int
	Write(val, addr int)
}

const (
	BIT_0 = 1 << iota
	BIT_1
	BIT_2
	BIT_3
	BIT_4
	BIT_5
	BIT_6
	BIT_7
)

func (cpu *Cpu) tick() {
	// Read current opcode from memory
	// Decode instruction
	// Build Instruction object
	// Place it as current instruction inside CPU
	// Call corresponding method
}

// Fetches the next instruction
func (cpu *Cpu) fetch() int {
	opcode := cpu.m.Read(cpu.pc)
	cpu.pc++

	return opcode
}

// Creates a new instruction from the given opcode,
// containing the operands (if any)
func (cpu *Cpu) decode(opcode int) {
	instr := instructionSet[opcode]

	for i := 0; i < instr.size-1; i++ {
		instr.operands[i] = cpu.m.Read(cpu.pc)
		cpu.pc++
	}

	cpu.nextInstr = instr
}

func (cpu *Cpu) nop() {
	fmt.Println("NOP")
}

//func (cpu *Cpu) ld_n_nn(reg *int, val int) {
//	*reg = val
//}
//
//func (cpu *Cpu) ld_r1_r2(reg1, reg2 *int) {
//	*reg1 = *reg2
//}
