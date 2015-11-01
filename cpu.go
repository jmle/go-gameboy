package main

type Cpu struct {
	// Registers
	pc, sp           int // Program counter, Stack pointer
	a                int // Accumulator, Flag
	b, c, d, e, h, l int // Auxiliary registers
	p                FlagReg

	// Memory
	m Memory
}

type FlagReg struct {
	z, n, h, c int
}

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

func (cpu *Cpu) loadImm(reg *int, val int) {
	*reg = val
}
