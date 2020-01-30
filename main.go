package main

import (
	"github.com/zatkins-dev/hw1-atkins/z80"
)

func main() {
	z80.InitMemory()
	z80.LoadTest()
	
	z80.CPUPower = true
	for {
		z80.CPUStep()
		
	}
}
