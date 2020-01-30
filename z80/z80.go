package z80

import (
	"fmt"
	"os"
)

// Word macro for uint16
type Word uint16

func byteToHigh(w Word, b Word) Word { return ((b << 8) | (w & 0x00FF)) }
func byteToLow(w Word, b Word) Word  { return ((b >> 8) | (w & 0xFF00)) }

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
	clockCycles[0x76] = 4
	clockCycles[0x03] = 4
	clockCycles[0x80] = 4
}

// LoadTest loads test instructions
func LoadTest() {
	writeByte(0x0100, 0x3C)
	writeByte(0x0101, 0x3C)
	writeByte(0x0102, 0x3C)
	writeByte(0x0103, 0x80)
	writeByte(0x0104, 0x76)
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
		break

	// HALT
	case 0x76:
		halted = true
		break

	// INC A    Example for increments a particular byte
	case 0x3C:
		AF = hightoHigh(AF, AF+0x0100)
		//ignore setting flags for now
		break

	// INC BC    Example for increments a word
	case 0x03:
		BC++
		//ignore setting flags for now
		break

	// ADD A,B   Add  byte to a byte  result in A
	case 0x80:
		AF = byteToHigh(AF, (AF>>8)+(BC>>8))
		//ignore setting flags for now
		break

	default:
		fmt.Printf("Instruction %02x not valid (at %04x)\n\n", instruction, PC-1)
		halt()
		break
	}
}
