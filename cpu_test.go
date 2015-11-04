package main

import (
	"testing"
)

type Memory struct {
	memory [1 << 16]int
}

func (m *Memory) Read(addr int) int {
	return m.memory[addr]
}

func (m *Memory) Write(addr, val int) {
	m.memory[addr] = val
}

func testDecodeWithSize1(t *testing.T) {
	cpu := Cpu{m: Memory{}}

	expectedInstr := instructionSet[0x00]

	instr := cpu.decode(0)

	if instr != expectedInstr {
		t.Errorf("Expected %+v, got %+v\n", expectedInstr, instr)
	}
}

func testDecodeWithSize2(t *testing.T) {
	cpu := Cpu{m: Memory{}}
	cpu.m.Write(0, 2)

	expectedInstr := instructionSet[0x06]
	expectedInstr.operands[0] = 2

	instr := cpu.decode(0x06)

	if instr != expectedInstr {
		t.Errorf("Expected %+v, got %+v\n", expectedInstr, instr)
	}

	if cpu.pc != 1 {
		t.Errorf("Expected %+v, got %+v\n", 1, cpu.pc)
	}
}

func testDecodeWithSize3(t *testing.T) {
	cpu := Cpu{m: Memory{}}
	cpu.m.Write(0, 2)
	cpu.m.Write(1, 4)

	expectedInstr := instructionSet[0x08]
	expectedInstr.operands[0] = 2
	expectedInstr.operands[1] = 4

	instr := cpu.decode(0x08)

	if instr != expectedInstr {
		t.Errorf("Expected %+v, got %+v\n", expectedInstr, instr)
	}

	if cpu.pc != 2 {
		t.Errorf("Expected %+v, got %+v\n", 2, cpu.pc)
	}
}

//func testLdRegWithImm(t *testing.T) {
//	var mem Memory
//	cpu := Cpu{
//		m: mem,
//	}
//
//	cpu.loadImm(&cpu.b, 35)
//
//	t.Log("testLoadImm")
//
//	if cpu.b != 35 {
//		t.Errorf("Expected %+v, got %+v\n", 35, cpu.b)
//	}
//}

//func testLdRegWithReg(t *testing.T) {
//	var mem Memory
//	cpu := Cpu{
//		m: mem,
//		d: 5,
//		b: 3,
//	}
//
//	cpu.loadReg(&cpu.b, &cpu.d)
//
//	t.Log("testLoadReg")
//
//	if cpu.b != cpu.d {
//		t.Errorf("Expected %+v, got %+v\n", cpu.b, cpu.d)
//	}
//}

//func testLdAcc(t *testing.T) {
//	var mem Memory
//	cpu := Cpu{
//		m: mem,
//		d: 5,
//		b: 3,
//	}
//
//	cpu.loadReg(&cpu.b, &cpu.d)
//
//	t.Log("testLoadReg")
//
//	if cpu.b != cpu.d {
//		t.Errorf("Expected %+v, got %+v\n", cpu.b, cpu.d)
//	}
//}
