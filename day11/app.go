package main

import (
	"../utils"
	. "./seating"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./day11/puzzle_input.txt")

	fmt.Println("Solution 1: ", SolvePart1(lines))
	fmt.Println("Solution 2: ", SolvePart2(lines))
}

func SolvePart1(lines []string) int {
	seating := NewSeating(lines)

	rounds := 0
	for {
		changed := seating.ApplyRound()
		rounds++
		if !changed {
			break
		}
	}

	return seating.CountOccs()
}

func SolvePart2(lines []string) int {
	seating := NewSeating(lines)

	rounds := 0
	for {
		changed := seating.ApplyRound2()
		rounds++
		if !changed {
			break
		}
	}

	return seating.CountOccs()
}
