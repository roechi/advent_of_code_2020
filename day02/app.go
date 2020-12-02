package main

import (
	"../utils"
	. "./passwordpolicies"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day02/puzzle_input.txt")

	correctPol1 := 0
	correctPol2 := 0
	for _, s := range lines {
		split := strings.Split(s, " ")
		times := strings.Split(split[0], "-")
		char := split[1][0]
		password := split[2]
		firstNum, errFirst := strconv.Atoi(times[0])
		secondNum, errSecond := strconv.Atoi(times[1])

		if errFirst != nil {
			log.Fatal(errFirst)
		}
		if errSecond != nil {
			log.Fatal(errSecond)
		}

		if ApplyFirstPolicy(firstNum, secondNum, string(char), password) {
			correctPol1++
		}
		if ApplySecondPolicy(firstNum, secondNum, string(char), password) {
			correctPol2++
		}
	}

	log.Println("Correct passwords for first policy: ", correctPol1)
	log.Println("Correct passwords for second policy: ", correctPol2)
}
