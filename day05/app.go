package main

import (
	"../utils"
	. "./boardingpass"
	"fmt"
	"sort"
)

func main() {
	lines := utils.ReadLines("./day05/puzzle_input.txt")

	ids := make([]int, len(lines))

	for i, l := range lines {
		row, col := FindSeat(l)
		id := CalculateId(row, col)
		ids[i] = id
	}

	m := 0
	for _, e := range ids {
		if e > m {
			m = e
		}
	}

	fmt.Println("Max ID: ", m)

	sort.Ints(ids)

	for i, id := range ids[:len(ids)-2] {
		if ids[i+1]-id == 2 {
			fmt.Println("My seat: ", id+1)
			break
		}
	}

}
