package main

import (
	"../utils"
	. "./navigation"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./day12/puzzle_input.txt")

	ferry := NewFerry(0, 0, 1, 0)

	for _, line := range lines {
		ApplyNavigationCommand(ferry, line)
	}

	fmt.Println("Solution 1: ", ManhattanDist(ferry))

	ferry2 := NewFerry2(0, 0, 10, 1)

	for _, line := range lines {
		ApplyNavigationCommand2(ferry2, line)
	}

	fmt.Println("Solution 2: ", ManhattanDist2(ferry2))
}
