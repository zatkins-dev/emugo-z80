package z80

import (
	"fmt"
	"os"
)

// Word macro for uint16
type Word uint16

func byteToHigh(w Word, b Word) Word { return ((b << 8) | (w & 0x00FF)) }
func byteToLow(w Word, b Word) Word  { return (b | (w & 0xFF00)) }

func hightoHigh(dst Word, src Word) Word { return ((src & 0xFF00) | (dst & 0x00FF)) }
func lowtoHigh(dst Word, src Word) Word  { return ((src << 8) | (dst & 0x00FF)) }
func hightoLow(dst Word, src Word) Word  { return ((src >> 8) | (dst & 0xFF00)) }
func lowtoLow(dst Word, src Word) Word   { return ((src & 0x00FF) | (dst & 0xFF00)) }

var memory [0xFFFF]byte

// AF accumulator/flags register
var AF Word = 0x01B0

// BC register
var BC Word = 0x0804

// DE register
var DE Word = 0x0201

// HL register
var HL Word = 0x0000

// SP stack pointer register
var SP Word = 0xFFFE

// PC program counter register
var PC Word = 0x0100

func printRegisters() {
	fmt.Printf("PC: %04x, AF: %04x, BC: %04x, DE: %04x, HL: %04x, SP: %04x\n", PC, AF, BC, DE, HL, SP)
}

var clockCycles []byte = make([]byte, 0xFF)
var instructions4cc []byte = []byte{0x04, 0x05, 0x0C, 0x0D, 0x14, 0x15, 0x1C, 0x1D, 0x24, 0x25, 0x2C, 0x2D, 0x3C, 0x3D, 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA7, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAF, 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97}
var instructionsCount []uint32 = make([]uint32, 0xFF)
var instruction byte = 0x00
var halted bool = false

func writeByte(loc uint, data byte) { memory[loc] = data }
func readByte(loc uint) byte        { return memory[loc] }

func fetch() byte {
	result := readByte(uint(PC))
	PC++
	return result
}
func readWord() Word {
	low := Word(readByte(uint(PC)))
	PC++
	high := Word(readByte(uint(PC)))
	PC++
	return low | (high << 8)
}

// Halt to program
func halt() {
	fmt.Printf("Total Clock Cycles: %d\n", TotalClockCycles)
	for i := 0; i < len(instructionsCount); i++ {
		if instructionsCount[i] > 0 {
			fmt.Printf("Instruction 0x%02x count is  %04x\n", i, instructionsCount[i])
		}
	}
	fmt.Printf("Halting now.\n")
	os.Exit(0)
}

// CPUPower Power state of the CPU
var CPUPower bool = false

// TotalClockCycles running count of CPU cycles
var TotalClockCycles uint64 = 0

// InitMemory sets up memory
func InitMemory() {
	for i := 0; i < len(clockCycles); i++ {
		clockCycles[i] = 1
	}
	clockCycles[0x3C] = 4
	clockCycles[0x76] = 1
	clockCycles[0x03] = 4
	clockCycles[0x80] = 4
	for ind := range instructions4cc {
		clockCycles[instructions4cc[ind]] = 4
	}
}

// LoadTest loads test instructions
func LoadTest() {
	writeByte(0x0100, 0x3C)
	writeByte(0x0101, 0x3C)
	writeByte(0x0102, 0x3C)
	writeByte(0x0103, 0x80)
	writeByte(0x0104, 0x76)
}

// LoadTest2 loads test instructions
func LoadTest2() {
	BC = 0x3331
	writeByte(0x0100, 0x3C)
	writeByte(0x0101, 0x3C)
	writeByte(0x0102, 0x3C)
	writeByte(0x0103, 0x3C)
	writeByte(0x0104, 0xA0)
	writeByte(0x0105, 0xA1)
	writeByte(0x0106, 0x76)
}

// LoadFirst loads first hw assignment instructions
func LoadFirst() {
	writeByte(0x0100, 0x3C)
	writeByte(0x0101, 0x3C)
	writeByte(0x0102, 0x3C)
	writeByte(0x0103, 0x80)
	writeByte(0x0104, 0xA0)
	writeByte(0x0105, 0xB3)
	writeByte(0x0106, 0x1C)
	writeByte(0x0107, 0xB3)
	writeByte(0x0108, 0xA1)
	writeByte(0x0109, 0x0C)
	writeByte(0x010A, 0xA1)
	writeByte(0x010B, 0xAA)
	writeByte(0x010C, 0x15)
	writeByte(0x010D, 0xAA)
	writeByte(0x010E, 0x76)
}

// CPUStep executes next instruction on the stack
func CPUStep() {
	if !CPUPower {
		return
	}
	printRegisters()
	fmt.Printf("TotalClockCycles at %08d\n", TotalClockCycles)

	// Check if halted
	if halted {
		halt()
	}

	//fetch
	instruction = fetch()
	fmt.Printf("instruction = %04x\n ", instruction)

	//decode
	TotalClockCycles += uint64(clockCycles[instruction])
	instructionsCount[instruction]++
	switch instruction {
	// NOP
	case 0x0:
	// HALT
	case 0x76:
		halted = true
	/* INC/DEC */
	// INC B
	case 0x04:
		BC = hightoHigh(BC, BC+0x0100)
	// DEC B
	case 0x05:
		BC = hightoHigh(BC, BC-0x0100)
	// INC C
	case 0x0C:
		BC = lowtoLow(BC, BC+0x0001)
	// DEC C
	case 0x0D:
		BC = lowtoLow(BC, BC-0x0001)
	// INC D
	case 0x14:
		DE = hightoHigh(DE, DE+0x0100)
	// DEC D
	case 0x15:
		DE = hightoHigh(DE, DE-0x0100)
	// INC E
	case 0x1C:
		DE = lowtoLow(DE, DE+0x0001)
	// DEC E
	case 0x1D:
		DE = lowtoLow(DE, DE-0x0001)
	// INC H
	case 0x24:
		HL = hightoHigh(HL, HL+0x0100)
	// DEC H
	case 0x25:
		HL = hightoHigh(HL, HL-0x0100)
	// INC L
	case 0x2C:
		HL = lowtoLow(HL, HL+0x0001)
	// DEC L
	case 0x2D:
		HL = lowtoLow(HL, HL-0x0001)
	// INC A
	case 0x3C:
		AF = hightoHigh(AF, AF+0x0100)
	// DEC A
	case 0x3D:
		AF = hightoHigh(AF, AF-0x0100)
	// INC BC    Example for increments a word
	case 0x03:
		BC++
		//ignore setting flags for now

	/* AND */
	// AND A,B
	case 0xA0:
		AF = byteToHigh(AF, (AF>>8)&(BC>>8))
	// AND A,C
	case 0xA1:
		AF = byteToHigh(AF, (AF>>8)&(BC))
	// AND A,D
	case 0xA2:
		AF = byteToHigh(AF, (AF>>8)&(DE>>8))
	// AND A,E
	case 0xA3:
		AF = byteToHigh(AF, (AF>>8)&(DE))
	// AND A,H
	case 0xA4:
		AF = byteToHigh(AF, (AF>>8)&(HL>>8))
	// AND A,L
	case 0xA5:
		AF = byteToHigh(AF, (AF>>8)&(HL))
	// AND A,A
	case 0xA7:
		AF = byteToHigh(AF, (AF>>8)&(AF>>8))

	/* OR */
	// OR A,B
	case 0xB0:
		AF = byteToHigh(AF, (AF>>8)|(BC>>8))
	// OR A,C
	case 0xB1:
		AF = byteToHigh(AF, (AF>>8)|(BC))
	// OR A,D
	case 0xB2:
		AF = byteToHigh(AF, (AF>>8)|(DE>>8))
	// OR A,E
	case 0xB3:
		AF = byteToHigh(AF, (AF>>8)|(DE))
	// OR A,H
	case 0xB4:
		AF = byteToHigh(AF, (AF>>8)|(HL>>8))
	// OR A,L
	case 0xB5:
		AF = byteToHigh(AF, (AF>>8)|(HL))
	// OR A, A pretty sure this is a NOP
	case 0xB7:
		AF = byteToHigh(AF, (AF>>8)|(AF>>8))

	/* XOR */
	// XOR A,B
	case 0xA8:
		AF = byteToHigh(AF, (AF>>8)^(BC>>8))
	// XOR A,C
	case 0xA9:
		AF = byteToHigh(AF, (AF>>8)^(BC))
	// XOR A,D
	case 0xAA:
		AF = byteToHigh(AF, (AF>>8)^(DE>>8))
	// XOR A,E
	case 0xAB:
		AF = byteToHigh(AF, (AF>>8)^(DE))
	// XOR A,H
	case 0xAC:
		AF = byteToHigh(AF, (AF>>8)^(HL>>8))
	// XOR A,L
	case 0xAD:
		AF = byteToHigh(AF, (AF>>8)^(HL))
	// XOR A,A
	case 0xAF:
		AF = byteToHigh(AF, (AF>>8)^(AF>>8))

	/* ADD */
	// ADD A,B
	case 0x80:
		AF = byteToHigh(AF, (AF>>8)+(BC>>8))
	// ADD A,C
	case 0x81:
		AF = byteToHigh(AF, (AF>>8)+(BC))
	// ADD A,D
	case 0x82:
		AF = byteToHigh(AF, (AF>>8)+(DE>>8))
	// ADD A,E
	case 0x83:
		AF = byteToHigh(AF, (AF>>8)+(DE))
	// ADD A,H
	case 0x84:
		AF = byteToHigh(AF, (AF>>8)+(HL>>8))
	// ADD A,L
	case 0x85:
		AF = byteToHigh(AF, (AF>>8)+(HL))
	// ADD A,A
	case 0x87:
		AF = byteToHigh(AF, (AF>>8)+(AF>>8))

	/* SUB */
	// SUB A,B
	case 0x90:
		AF = byteToHigh(AF, (AF>>8)-(BC>>8))
	// SUB A,C
	case 0x91:
		AF = byteToHigh(AF, (AF>>8)-(BC))
	// SUB A,D
	case 0x92:
		AF = byteToHigh(AF, (AF>>8)-(DE>>8))
	// SUB A,E
	case 0x93:
		AF = byteToHigh(AF, (AF>>8)-(DE))
	// SUB A,H
	case 0x94:
		AF = byteToHigh(AF, (AF>>8)-(HL>>8))
	// SUB A,L
	case 0x95:
		AF = byteToHigh(AF, (AF>>8)-(HL))
	// SUB A,A
	case 0x97:
		AF = byteToHigh(AF, (AF>>8)-(AF>>8))

	default:
		fmt.Printf("Instruction %02x not valid (at %04x)\n\n", instruction, PC-1)
		halt()
	}
}
