package main

import (
	"../utils"
	. "./encoding"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	lines := utils.ReadLines("./day09/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	startFaulty := time.Now()
	faulty := FindFirstFaulty(ints, 25)
	endFaulty := time.Since(startFaulty)

	fmt.Println("First faulty: ", ints[faulty])
	fmt.Println("Time: ", endFaulty)

	startSet := time.Now()
	set := FindContiguousSet(ints[:faulty], ints[faulty])
	endSet := time.Since(startSet)
	sort.Ints(set)
	fmt.Println("Set: ", set)
	fmt.Println("Solution: ", set[0]+set[len(set)-1])
	fmt.Println("Time: ", endSet)
}
