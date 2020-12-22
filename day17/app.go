package main

import (
	"../utils"
	. "./simulation"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./day17/puzzle_input.txt")

	simulation := NewSimulation(lines)
	for i := 0; i < 6; i++ {
		simulation.Cycle()
	}

	fmt.Println("Solution: ", simulation.CountActive())
}
