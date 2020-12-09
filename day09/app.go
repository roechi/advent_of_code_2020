package main

import (
	"../utils"
	. "./encoding"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := utils.ReadLines("./day09/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	faulty := FindFirstFaulty(ints, 25)

	fmt.Println("First faulty: ", ints[faulty])

	set := FindContiguousSet(ints[:faulty], ints[faulty])

	sort.Ints(set)
	fmt.Println("Set: ", set)
	fmt.Println("Solution: ", set[0]+set[len(set)-1])
}
