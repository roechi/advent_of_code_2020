package main

import (
	"../utils"
	. "./customsanswers"
	"log"
)

func main() {
	lines := utils.ReadLines("./day06/puzzle_input.txt")

	i := 0
	yesCount := 0
	allMatchCount := 0
	for j, l := range lines {
		if l == "" {
			group := lines[i:j]
			yesCount += CountUniqueAnswers(group)
			allMatchCount += CountAllYes(group)
			i = j + 1
		}
	}

	log.Println("Result Part 1: ", yesCount)
	log.Println("Result Part 2: ", allMatchCount)
}
