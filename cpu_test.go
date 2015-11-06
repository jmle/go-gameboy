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

func TestDecode(t *testing.T) {
	for _, tt := range []struct {
		testName         string
		opcode           int
		expectedInstr    Instruction
		expectedOperands [2]int
	}{
		{
			testName:      "Decode with size 1",
			opcode:        0x00,
			expectedInstr: instructionSet[0x00],
		},
		{
			testName:         "Decode with size 2",
			opcode:           0x06,
			expectedInstr:    instructionSet[0x06],
			expectedOperands: [2]int{2},
		},
		{
			testName:         "Decode with size 3",
			opcode:           0x08,
			expectedInstr:    instructionSet[0x08],
			expectedOperands: [2]int{2, 4},
		},
	} {
		t.Log(tt.testName)

		cpu := Cpu{m: Memory{}}

		// Populate memory with operands
		for i := 0; i < tt.expectedInstr.size-1; i++ {
			cpu.m.Write(i, tt.expectedOperands[i])
		}

		// Populate expected instruction with operands
		tt.expectedInstr.operands = tt.expectedOperands

		cpu.decode(tt.opcode)
		instr := cpu.nextInstr

		if !instructionsEqual(tt.expectedInstr, instr) {
			t.Errorf("Expected %+v, got %+v\n", tt.expectedInstr, instr)
		}

		if cpu.pc != tt.expectedInstr.size-1 {
			t.Errorf("Expected %+v, got %+v\n", tt.expectedInstr.size-1, cpu.pc)
		}
	}
}

func TestNop(t *testing.T) {
	cpu := Cpu{m: Memory{}}

	cpu.decode(0x00)

	instr := cpu.nextInstr
	operation := instr.operation

	operation(&cpu)
}

// Compares Instructions skipping the operation field
func instructionsEqual(instr1, instr2 Instruction) bool {
	cmpName := instr1.name == instr2.name
	cmpSize := instr1.size == instr2.size
	cmpCycles := instr1.cycles == instr2.cycles
	cmpOperands := instr1.operands == instr2.operands

	return cmpName && cmpCycles && cmpSize && cmpOperands

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
