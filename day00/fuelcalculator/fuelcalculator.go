package fuelcalulator

import "math"

func CalcFuelForModule(weight int) int {
	return int(math.Floor(float64(weight / 3)) - 2)
}

func CalcFuelForModules(moduleWeights []int) int {
	totalFuel := 0

	for _, w := range moduleWeights {
		totalFuel += CalcFuelForModule(w)
	}

	return totalFuel
}

func CalcFuelForModuleRecursively(weight int) int {
	fuelForThisModule := CalcFuelForModule(weight)

	if (fuelForThisModule <= 0) {
		return 0
	} else {
		return fuelForThisModule + CalcFuelForModuleRecursively(fuelForThisModule)
	}
}

func CalcFuelForModulesRecursively(moduleWeights []int) int {
	totalFuel := 0

	for _, w := range moduleWeights {
		totalFuel += CalcFuelForModuleRecursively(w)
	}
	return totalFuel
}