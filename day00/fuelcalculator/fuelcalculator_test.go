package fuelcalulator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFuelForModule(t *testing.T) {
	given := 14
	want := 2

	got := CalcFuelForModule(given)

	assert.Equal(t, want, got)
}

func TestCalcFuelForModules(t *testing.T) {
	given := []int{14, 12, 1969}
	want := 658

	got := CalcFuelForModules(given)

	assert.Equal(t, want, got)
}

func TestCalcFuelRecursively(t *testing.T) {
	given := 1969
	want := 966

	got := CalcFuelForModuleRecursively(given)

	assert.Equal(t, want, got)
}
