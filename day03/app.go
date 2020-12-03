package main

import (
	"../utils"
	. "./sled"
	"fmt"
	"log"
)

func main() {
	lines := utils.ReadLines("./day03/puzzle_input.txt")

	slopes := make([][]int, 5)
	slopes[0] = []int{1, 1}
	slopes[1] = []int{1, 3}
	slopes[2] = []int{1, 5}
	slopes[3] = []int{1, 7}
	slopes[4] = []int{2, 1}

	yLen := len(lines)
	xLen := len(lines[0])
	matrix := make([][]int, yLen)
	for i := 0; i < yLen; i++ {
		matrix[i] = make([]int, xLen)
	}

	for y, line := range lines {
		for x, char := range line {
			if string(char) == "." {
				matrix[y][x] = 0
			} else if string(char) == "#" {
				matrix[y][x] = 1
			} else {
				log.Fatal("Unknown terrain field: ", char)
			}
		}
	}

	totalObstacles := 1
	for _, slope := range slopes {
		fmt.Println("Current slope: ", slope[0], ", ", slope[1])
		obstacles := ScanSlopeForObstacles(slope[1], slope[0], matrix)
		fmt.Println("Current obstacles: ", obstacles)
		totalObstacles *= obstacles
	}

	log.Println("Obstacles: ", totalObstacles)

}
