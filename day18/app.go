package main

import (
	"../utils"
	. "./specialmath"
	"fmt"
)

func main() {
	lines := utils.ReadLines("./day18/puzzle_input.txt")
	sum := 0

	for _, line := range lines {
		sum += SolveWithParenthesis(line)
	}

	fmt.Println("Solution 1: ", sum)

	sum2 := 0

	for _, line := range lines {
		sum2 += SolveWithParenthesisSpecial(line)
	}

	fmt.Println("Solution 2: ", sum2)
}
