package main

import (
	"github.com/zatkins-dev/z80emulator/z80"
)

func main() {
	z80.InitMemory()
	// z80.LoadTest()
	z80.LoadTest2()
	// z80.LoadFirst()

	z80.CPUPower = true
	for {
		z80.CPUStep()
	}
}
