package main

import (
	"../utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day15/puzzle_input.txt")

	numberstrings := strings.Split(lines[0], ",")
	ints := make([]int, len(numberstrings))

	for i, str := range numberstrings {
		ints[i], _ = strconv.Atoi(str)
	}

	mem := make(map[int]int)

	t := 0
	lastNum := 0
	for _, num := range ints[:len(ints)-1] {
		mem[num] = t
		t++
		lastNum = num
	}
	lastNum = ints[len(ints)-1]

	for t < 30000000-1 {
		lastTimeSpoken, known := mem[lastNum]
		mem[lastNum] = t
		if known {
			lastNum = t - lastTimeSpoken
		} else {
			lastNum = 0
		}

		t++
	}

	fmt.Println(lastNum)
}
