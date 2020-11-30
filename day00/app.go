package main

import (
	"./fuelcalculator"
	"fmt"
	"strconv"
	"../utils"
)

func main() {
	lines := utils.ReadLines("./day00/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	totalWeight := fuelcalulator.CalcFuelForModules(ints)
	fmt.Println("Total fuel needed: ", totalWeight)

	totalWeightAdjusted := fuelcalulator.CalcFuelForModulesRecursively(ints)
	fmt.Println("The adjusted amount of total fuel needed is: ", totalWeightAdjusted)
}
