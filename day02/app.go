package main

import (
	"../utils"
	. "./passwordpolicies"
	"fmt"
	"log"
)

func main() {
	lines := utils.ReadLines("./day02/puzzle_input.txt")

	correctPol1 := 0
	correctPol2 := 0
	for _, line := range lines {
		var firstNum, secondNum int
		var char, pass string

		fmt.Sscanf(line, "%d-%d %1s: %s", &firstNum, &secondNum, &char, &pass)

		if ApplyFirstPolicy(firstNum, secondNum, char, pass) {
			correctPol1++
		}
		if ApplySecondPolicy(firstNum, secondNum, char, pass) {
			correctPol2++
		}
	}

	log.Println("Correct passwords for first policy: ", correctPol1)
	log.Println("Correct passwords for second policy: ", correctPol2)
}
