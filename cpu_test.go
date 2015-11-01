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

func testLdAuxWithImm(t *testing.T) {
	var mem Memory
	cpu := Cpu{
		m: mem,
	}

	cpu.loadImm(&cpu.b, 35)

	t.Log("testLoadImm")

	if cpu.b != 35 {
		t.Errorf("Expected %+v, got %+v\n", 35, cpu.b)
	}
}
