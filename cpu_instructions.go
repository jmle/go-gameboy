package main

type Instruction struct {
	name     string
	size     int
	cycles   int
	operands [2]int
}

// Instruction set info extracted from http://www.pastraiser.com/cpu/gameboy/gameboy_opcodes.html
var instructionSet = map[int]Instruction{
	0x0:  Instruction{name: "NOP", size: 1, cycles: 4},
	0x1:  Instruction{name: "LD BC,d16", size: 3, cycles: 12},
	0x2:  Instruction{name: "LD (BC),A", size: 1, cycles: 8},
	0x3:  Instruction{name: "INC BC", size: 1, cycles: 8},
	0x4:  Instruction{name: "INC B", size: 1, cycles: 4},
	0x5:  Instruction{name: "DEC B", size: 1, cycles: 4},
	0x6:  Instruction{name: "LD B,d8", size: 2, cycles: 8},
	0x7:  Instruction{name: "RLCA", size: 1, cycles: 4},
	0x8:  Instruction{name: "LD (a16),SP", size: 3, cycles: 20},
	0x9:  Instruction{name: "ADD HL,BC", size: 1, cycles: 8},
	0xa:  Instruction{name: "LD A,(BC)", size: 1, cycles: 8},
	0xb:  Instruction{name: "DEC BC", size: 1, cycles: 8},
	0xc:  Instruction{name: "INC C", size: 1, cycles: 4},
	0xd:  Instruction{name: "DEC C", size: 1, cycles: 4},
	0xe:  Instruction{name: "LD C,d8", size: 2, cycles: 8},
	0xf:  Instruction{name: "RRCA", size: 1, cycles: 4},
	0x10: Instruction{name: "STOP 0", size: 2, cycles: 4},
	0x11: Instruction{name: "LD DE,d16", size: 3, cycles: 12},
	0x12: Instruction{name: "LD (DE),A", size: 1, cycles: 8},
	0x13: Instruction{name: "INC DE", size: 1, cycles: 8},
	0x14: Instruction{name: "INC D", size: 1, cycles: 4},
	0x15: Instruction{name: "DEC D", size: 1, cycles: 4},
	0x16: Instruction{name: "LD D,d8", size: 2, cycles: 8},
	0x17: Instruction{name: "RLA", size: 1, cycles: 4},
	0x18: Instruction{name: "JR r8", size: 2, cycles: 12},
	0x19: Instruction{name: "ADD HL,DE", size: 1, cycles: 8},
	0x1a: Instruction{name: "LD A,(DE)", size: 1, cycles: 8},
	0x1b: Instruction{name: "DEC DE", size: 1, cycles: 8},
	0x1c: Instruction{name: "INC E", size: 1, cycles: 4},
	0x1d: Instruction{name: "DEC E", size: 1, cycles: 4},
	0x1e: Instruction{name: "LD E,d8", size: 2, cycles: 8},
	0x1f: Instruction{name: "RRA", size: 1, cycles: 4},
	0x20: Instruction{name: "JR NZ,r8", size: 2, cycles: 12 / 8},
	0x21: Instruction{name: "LD HL,d16", size: 3, cycles: 12},
	0x22: Instruction{name: "LD (HL+),A", size: 1, cycles: 8},
	0x23: Instruction{name: "INC HL", size: 1, cycles: 8},
	0x24: Instruction{name: "INC H", size: 1, cycles: 4},
	0x25: Instruction{name: "DEC H", size: 1, cycles: 4},
	0x26: Instruction{name: "LD H,d8", size: 2, cycles: 8},
	0x27: Instruction{name: "DAA", size: 1, cycles: 4},
	0x28: Instruction{name: "JR Z,r8", size: 2, cycles: 12 / 8},
	0x29: Instruction{name: "ADD HL,HL", size: 1, cycles: 8},
	0x2a: Instruction{name: "LD A,(HL+)", size: 1, cycles: 8},
	0x2b: Instruction{name: "DEC HL", size: 1, cycles: 8},
	0x2c: Instruction{name: "INC L", size: 1, cycles: 4},
	0x2d: Instruction{name: "DEC L", size: 1, cycles: 4},
	0x2e: Instruction{name: "LD L,d8", size: 2, cycles: 8},
	0x2f: Instruction{name: "CPL", size: 1, cycles: 4},
	0x30: Instruction{name: "JR NC,r8", size: 2, cycles: 12 / 8},
	0x31: Instruction{name: "LD SP,d16", size: 3, cycles: 12},
	0x32: Instruction{name: "LD (HL-),A", size: 1, cycles: 8},
	0x33: Instruction{name: "INC SP", size: 1, cycles: 8},
	0x34: Instruction{name: "INC (HL)", size: 1, cycles: 12},
	0x35: Instruction{name: "DEC (HL)", size: 1, cycles: 12},
	0x36: Instruction{name: "LD (HL),d8", size: 2, cycles: 12},
	0x37: Instruction{name: "SCF", size: 1, cycles: 4},
	0x38: Instruction{name: "JR C,r8", size: 2, cycles: 12 / 8},
	0x39: Instruction{name: "ADD HL,SP", size: 1, cycles: 8},
	0x3a: Instruction{name: "LD A,(HL-)", size: 1, cycles: 8},
	0x3b: Instruction{name: "DEC SP", size: 1, cycles: 8},
	0x3c: Instruction{name: "INC A", size: 1, cycles: 4},
	0x3d: Instruction{name: "DEC A", size: 1, cycles: 4},
	0x3e: Instruction{name: "LD A,d8", size: 2, cycles: 8},
	0x3f: Instruction{name: "CCF", size: 1, cycles: 4},
	0x40: Instruction{name: "LD B,B", size: 1, cycles: 4},
	0x41: Instruction{name: "LD B,C", size: 1, cycles: 4},
	0x42: Instruction{name: "LD B,D", size: 1, cycles: 4},
	0x43: Instruction{name: "LD B,E", size: 1, cycles: 4},
	0x44: Instruction{name: "LD B,H", size: 1, cycles: 4},
	0x45: Instruction{name: "LD B,L", size: 1, cycles: 4},
	0x46: Instruction{name: "LD B,(HL)", size: 1, cycles: 8},
	0x47: Instruction{name: "LD B,A", size: 1, cycles: 4},
	0x48: Instruction{name: "LD C,B", size: 1, cycles: 4},
	0x49: Instruction{name: "LD C,C", size: 1, cycles: 4},
	0x4a: Instruction{name: "LD C,D", size: 1, cycles: 4},
	0x4b: Instruction{name: "LD C,E", size: 1, cycles: 4},
	0x4c: Instruction{name: "LD C,H", size: 1, cycles: 4},
	0x4d: Instruction{name: "LD C,L", size: 1, cycles: 4},
	0x4e: Instruction{name: "LD C,(HL)", size: 1, cycles: 8},
	0x4f: Instruction{name: "LD C,A", size: 1, cycles: 4},
	0x50: Instruction{name: "LD D,B", size: 1, cycles: 4},
	0x51: Instruction{name: "LD D,C", size: 1, cycles: 4},
	0x52: Instruction{name: "LD D,D", size: 1, cycles: 4},
	0x53: Instruction{name: "LD D,E", size: 1, cycles: 4},
	0x54: Instruction{name: "LD D,H", size: 1, cycles: 4},
	0x55: Instruction{name: "LD D,L", size: 1, cycles: 4},
	0x56: Instruction{name: "LD D,(HL)", size: 1, cycles: 8},
	0x57: Instruction{name: "LD D,A", size: 1, cycles: 4},
	0x58: Instruction{name: "LD E,B", size: 1, cycles: 4},
	0x59: Instruction{name: "LD E,C", size: 1, cycles: 4},
	0x5a: Instruction{name: "LD E,D", size: 1, cycles: 4},
	0x5b: Instruction{name: "LD E,E", size: 1, cycles: 4},
	0x5c: Instruction{name: "LD E,H", size: 1, cycles: 4},
	0x5d: Instruction{name: "LD E,L", size: 1, cycles: 4},
	0x5e: Instruction{name: "LD E,(HL)", size: 1, cycles: 8},
	0x5f: Instruction{name: "LD E,A", size: 1, cycles: 4},
	0x60: Instruction{name: "LD H,B", size: 1, cycles: 4},
	0x61: Instruction{name: "LD H,C", size: 1, cycles: 4},
	0x62: Instruction{name: "LD H,D", size: 1, cycles: 4},
	0x63: Instruction{name: "LD H,E", size: 1, cycles: 4},
	0x64: Instruction{name: "LD H,H", size: 1, cycles: 4},
	0x65: Instruction{name: "LD H,L", size: 1, cycles: 4},
	0x66: Instruction{name: "LD H,(HL)", size: 1, cycles: 8},
	0x67: Instruction{name: "LD H,A", size: 1, cycles: 4},
	0x68: Instruction{name: "LD L,B", size: 1, cycles: 4},
	0x69: Instruction{name: "LD L,C", size: 1, cycles: 4},
	0x6a: Instruction{name: "LD L,D", size: 1, cycles: 4},
	0x6b: Instruction{name: "LD L,E", size: 1, cycles: 4},
	0x6c: Instruction{name: "LD L,H", size: 1, cycles: 4},
	0x6d: Instruction{name: "LD L,L", size: 1, cycles: 4},
	0x6e: Instruction{name: "LD L,(HL)", size: 1, cycles: 8},
	0x6f: Instruction{name: "LD L,A", size: 1, cycles: 4},
	0x70: Instruction{name: "LD (HL),B", size: 1, cycles: 8},
	0x71: Instruction{name: "LD (HL),C", size: 1, cycles: 8},
	0x72: Instruction{name: "LD (HL),D", size: 1, cycles: 8},
	0x73: Instruction{name: "LD (HL),E", size: 1, cycles: 8},
	0x74: Instruction{name: "LD (HL),H", size: 1, cycles: 8},
	0x75: Instruction{name: "LD (HL),L", size: 1, cycles: 8},
	0x76: Instruction{name: "HALT", size: 1, cycles: 4},
	0x77: Instruction{name: "LD (HL),A", size: 1, cycles: 8},
	0x78: Instruction{name: "LD A,B", size: 1, cycles: 4},
	0x79: Instruction{name: "LD A,C", size: 1, cycles: 4},
	0x7a: Instruction{name: "LD A,D", size: 1, cycles: 4},
	0x7b: Instruction{name: "LD A,E", size: 1, cycles: 4},
	0x7c: Instruction{name: "LD A,H", size: 1, cycles: 4},
	0x7d: Instruction{name: "LD A,L", size: 1, cycles: 4},
	0x7e: Instruction{name: "LD A,(HL)", size: 1, cycles: 8},
	0x7f: Instruction{name: "LD A,A", size: 1, cycles: 4},
	0x80: Instruction{name: "ADD A,B", size: 1, cycles: 4},
	0x81: Instruction{name: "ADD A,C", size: 1, cycles: 4},
	0x82: Instruction{name: "ADD A,D", size: 1, cycles: 4},
	0x83: Instruction{name: "ADD A,E", size: 1, cycles: 4},
	0x84: Instruction{name: "ADD A,H", size: 1, cycles: 4},
	0x85: Instruction{name: "ADD A,L", size: 1, cycles: 4},
	0x86: Instruction{name: "ADD A,(HL)", size: 1, cycles: 8},
	0x87: Instruction{name: "ADD A,A", size: 1, cycles: 4},
	0x88: Instruction{name: "ADC A,B", size: 1, cycles: 4},
	0x89: Instruction{name: "ADC A,C", size: 1, cycles: 4},
	0x8a: Instruction{name: "ADC A,D", size: 1, cycles: 4},
	0x8b: Instruction{name: "ADC A,E", size: 1, cycles: 4},
	0x8c: Instruction{name: "ADC A,H", size: 1, cycles: 4},
	0x8d: Instruction{name: "ADC A,L", size: 1, cycles: 4},
	0x8e: Instruction{name: "ADC A,(HL)", size: 1, cycles: 8},
	0x8f: Instruction{name: "ADC A,A", size: 1, cycles: 4},
	0x90: Instruction{name: "SUB B", size: 1, cycles: 4},
	0x91: Instruction{name: "SUB C", size: 1, cycles: 4},
	0x92: Instruction{name: "SUB D", size: 1, cycles: 4},
	0x93: Instruction{name: "SUB E", size: 1, cycles: 4},
	0x94: Instruction{name: "SUB H", size: 1, cycles: 4},
	0x95: Instruction{name: "SUB L", size: 1, cycles: 4},
	0x96: Instruction{name: "SUB (HL)", size: 1, cycles: 8},
	0x97: Instruction{name: "SUB A", size: 1, cycles: 4},
	0x98: Instruction{name: "SBC A,B", size: 1, cycles: 4},
	0x99: Instruction{name: "SBC A,C", size: 1, cycles: 4},
	0x9a: Instruction{name: "SBC A,D", size: 1, cycles: 4},
	0x9b: Instruction{name: "SBC A,E", size: 1, cycles: 4},
	0x9c: Instruction{name: "SBC A,H", size: 1, cycles: 4},
	0x9d: Instruction{name: "SBC A,L", size: 1, cycles: 4},
	0x9e: Instruction{name: "SBC A,(HL)", size: 1, cycles: 8},
	0x9f: Instruction{name: "SBC A,A", size: 1, cycles: 4},
	0xa0: Instruction{name: "AND B", size: 1, cycles: 4},
	0xa1: Instruction{name: "AND C", size: 1, cycles: 4},
	0xa2: Instruction{name: "AND D", size: 1, cycles: 4},
	0xa3: Instruction{name: "AND E", size: 1, cycles: 4},
	0xa4: Instruction{name: "AND H", size: 1, cycles: 4},
	0xa5: Instruction{name: "AND L", size: 1, cycles: 4},
	0xa6: Instruction{name: "AND (HL)", size: 1, cycles: 8},
	0xa7: Instruction{name: "AND A", size: 1, cycles: 4},
	0xa8: Instruction{name: "XOR B", size: 1, cycles: 4},
	0xa9: Instruction{name: "XOR C", size: 1, cycles: 4},
	0xaa: Instruction{name: "XOR D", size: 1, cycles: 4},
	0xab: Instruction{name: "XOR E", size: 1, cycles: 4},
	0xac: Instruction{name: "XOR H", size: 1, cycles: 4},
	0xad: Instruction{name: "XOR L", size: 1, cycles: 4},
	0xae: Instruction{name: "XOR (HL)", size: 1, cycles: 8},
	0xaf: Instruction{name: "XOR A", size: 1, cycles: 4},
	0xb0: Instruction{name: "OR B", size: 1, cycles: 4},
	0xb1: Instruction{name: "OR C", size: 1, cycles: 4},
	0xb2: Instruction{name: "OR D", size: 1, cycles: 4},
	0xb3: Instruction{name: "OR E", size: 1, cycles: 4},
	0xb4: Instruction{name: "OR H", size: 1, cycles: 4},
	0xb5: Instruction{name: "OR L", size: 1, cycles: 4},
	0xb6: Instruction{name: "OR (HL)", size: 1, cycles: 8},
	0xb7: Instruction{name: "OR A", size: 1, cycles: 4},
	0xb8: Instruction{name: "CP B", size: 1, cycles: 4},
	0xb9: Instruction{name: "CP C", size: 1, cycles: 4},
	0xba: Instruction{name: "CP D", size: 1, cycles: 4},
	0xbb: Instruction{name: "CP E", size: 1, cycles: 4},
	0xbc: Instruction{name: "CP H", size: 1, cycles: 4},
	0xbd: Instruction{name: "CP L", size: 1, cycles: 4},
	0xbe: Instruction{name: "CP (HL)", size: 1, cycles: 8},
	0xbf: Instruction{name: "CP A", size: 1, cycles: 4},
	0xc0: Instruction{name: "RET NZ", size: 1, cycles: 20 / 8},
	0xc1: Instruction{name: "POP BC", size: 1, cycles: 12},
	0xc2: Instruction{name: "JP NZ,a16", size: 3, cycles: 16 / 12},
	0xc3: Instruction{name: "JP a16", size: 3, cycles: 16},
	0xc4: Instruction{name: "CALL NZ,a16", size: 3, cycles: 24 / 12},
	0xc5: Instruction{name: "PUSH BC", size: 1, cycles: 16},
	0xc6: Instruction{name: "ADD A,d8", size: 2, cycles: 8},
	0xc7: Instruction{name: "RST 00H", size: 1, cycles: 16},
	0xc8: Instruction{name: "RET Z", size: 1, cycles: 20 / 8},
	0xc9: Instruction{name: "RET", size: 1, cycles: 16},
	0xca: Instruction{name: "JP Z,a16", size: 3, cycles: 16 / 12},
	0xcb: Instruction{name: "PREFIX CB", size: 1, cycles: 4},
	0xcc: Instruction{name: "CALL Z,a16", size: 3, cycles: 24 / 12},
	0xcd: Instruction{name: "CALL a16", size: 3, cycles: 24},
	0xce: Instruction{name: "ADC A,d8", size: 2, cycles: 8},
	0xcf: Instruction{name: "RST 08H", size: 1, cycles: 16},
	0xd0: Instruction{name: "RET NC", size: 1, cycles: 20 / 8},
	0xd1: Instruction{name: "POP DE", size: 1, cycles: 12},
	0xd2: Instruction{name: "JP NC,a16", size: 3, cycles: 16 / 12},
	0xd4: Instruction{name: "CALL NC,a16", size: 3, cycles: 24 / 12},
	0xd5: Instruction{name: "PUSH DE", size: 1, cycles: 16},
	0xd6: Instruction{name: "SUB d8", size: 2, cycles: 8},
	0xd7: Instruction{name: "RST 10H", size: 1, cycles: 16},
	0xd8: Instruction{name: "RET C", size: 1, cycles: 20 / 8},
	0xd9: Instruction{name: "RETI", size: 1, cycles: 16},
	0xda: Instruction{name: "JP C,a16", size: 3, cycles: 16 / 12},
	0xdc: Instruction{name: "CALL C,a16", size: 3, cycles: 24 / 12},
	0xde: Instruction{name: "SBC A,d8", size: 2, cycles: 8},
	0xdf: Instruction{name: "RST 18H", size: 1, cycles: 16},
	0xe0: Instruction{name: "LDH (a8),A", size: 2, cycles: 12},
	0xe1: Instruction{name: "POP HL", size: 1, cycles: 12},
	0xe2: Instruction{name: "LD (C),A", size: 2, cycles: 8},
	0xe5: Instruction{name: "PUSH HL", size: 1, cycles: 16},
	0xe6: Instruction{name: "AND d8", size: 2, cycles: 8},
	0xe7: Instruction{name: "RST 20H", size: 1, cycles: 16},
	0xe8: Instruction{name: "ADD SP,r8", size: 2, cycles: 16},
	0xe9: Instruction{name: "JP (HL)", size: 1, cycles: 4},
	0xea: Instruction{name: "LD (a16),A", size: 3, cycles: 16},
	0xee: Instruction{name: "XOR d8", size: 2, cycles: 8},
	0xef: Instruction{name: "RST 28H", size: 1, cycles: 16},
	0xf0: Instruction{name: "LDH A,(a8)", size: 2, cycles: 12},
	0xf1: Instruction{name: "POP AF", size: 1, cycles: 12},
	0xf2: Instruction{name: "LD A,(C)", size: 2, cycles: 8},
	0xf3: Instruction{name: "DI", size: 1, cycles: 4},
	0xf5: Instruction{name: "PUSH AF", size: 1, cycles: 16},
	0xf6: Instruction{name: "OR d8", size: 2, cycles: 8},
	0xf7: Instruction{name: "RST 30H", size: 1, cycles: 16},
	0xf8: Instruction{name: "LD HL,SP+r8", size: 2, cycles: 12},
	0xf9: Instruction{name: "LD SP,HL", size: 1, cycles: 8},
	0xfa: Instruction{name: "LD A,(a16)", size: 3, cycles: 16},
	0xfb: Instruction{name: "EI", size: 1, cycles: 4},
	0xfe: Instruction{name: "CP d8", size: 2, cycles: 8},
	0xff: Instruction{name: "RST 38H", size: 1, cycles: 16},
}
