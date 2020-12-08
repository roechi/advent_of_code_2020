package main

import (
	"../utils"
	. "./emulator"
	"fmt"
	"strings"
)

func main() {
	programCode := utils.ReadLines("./day08/puzzle_input.txt")

	for i, instr := range programCode {
		manipulatedCode := make([]string, len(programCode))
		copy(manipulatedCode, programCode)
		manipulatedCode[i] = strings.NewReplacer("jmp", "nop", "nop", "jmp").Replace(instr)

		console := NewConsole(manipulatedCode)

		acc, err := console.RunProgram()
		if err == nil {
			fmt.Println("Success: ", acc)
			break
		} else {
			fmt.Println("Infinite Loop detected")
		}
	}
}
