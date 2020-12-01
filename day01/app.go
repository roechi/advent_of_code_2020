package main

import (
	"../utils"
	. "./expenses"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadLines("./day01/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	expenses := CalcExpenses(ints)
	fmt.Println("First result: ", expenses)

	expenses2 := CalcExpenses2(ints)

	fmt.Println("Second result: ", expenses2)
}
