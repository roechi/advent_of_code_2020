package main

import (
	"../utils"
	"fmt"
	"golang.org/x/tools/container/intsets"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day13/puzzle_input.txt")

	fmt.Println("Solution 1: ", SolvePart1(lines))
	fmt.Println("Solution 2: ", SolvePart2(lines[1]))
}

func SolvePart1(lines []string) int {
	earliestDep, _ := strconv.Atoi(lines[0])

	ids := make(map[int]int)

	rawIds := strings.Split(lines[1], ",")
	for _, busIdString := range rawIds {
		if busIdString != "x" {
			id, _ := strconv.Atoi(busIdString)

			timeDiff := earliestDep % id
			if timeDiff == 0 {
				ids[id] = earliestDep
			} else {
				ids[id] = earliestDep - timeDiff + id
			}
		}
	}

	earliestId := intsets.MaxInt
	earliestPossibleDeparture := intsets.MaxInt

	for id, dep := range ids {
		if dep < earliestPossibleDeparture {
			earliestPossibleDeparture = dep
			earliestId = id
		}
	}

	fmt.Println(ids)

	return (earliestPossibleDeparture - earliestDep) * earliestId
}

func SolvePart2(schedule string) int {
	var buses []int
	for _, id := range strings.Split(schedule, ",") {
		if id == "x" {
			id = "1"
		}
		num, _ := strconv.Atoi(id)
		buses = append(buses, num)
	}

	timestamp := 1

	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {
			if (timestamp+offset)%buses[offset] != 0 {
				valid = false
				break
			}

			timeToSkipIfNoMatch *= buses[offset]
		}
		if valid {
			return timestamp
		}

		timestamp += timeToSkipIfNoMatch
	}
}
