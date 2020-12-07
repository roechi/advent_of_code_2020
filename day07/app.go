package main

import (
	"../utils"
	. "./bags"
)

func main() {
	lines := utils.ReadLines("./day07/puzzle_input.txt")

	ParseRules(lines)
}
