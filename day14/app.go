package main

import (
	"../utils"
	. "./docking"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./day14/puzzle_input.txt")

	program := NewProgram()

	for _, line := range lines {
		program.ApplyCommand(line)
	}

	mem := program.GetMem()

	sum := 0

	for _, val := range mem {
		sum += val
	}
	fmt.Println("Solution 1: ", sum)

	program2 := NewProgram()
	for _, line := range lines {
		program2.ApplyCommand2(line)
	}

	mem2 := program2.GetMem()

	sum2 := 0

	for _, val := range mem2 {
		sum2 += val
	}
	fmt.Println("Solution 2: ", sum2)
}
