package main

import (
	. "./navigation"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRotLeft(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "L90")

	assert.Equal(t, ferry.XDir, 0)
	assert.Equal(t, ferry.YDir, 1)

	ApplyNavigationCommand(ferry, "L90")

	assert.Equal(t, ferry.XDir, -1)
	assert.Equal(t, ferry.YDir, 0)

	ApplyNavigationCommand(ferry, "L90")

	assert.Equal(t, ferry.XDir, 0)
	assert.Equal(t, ferry.YDir, -1)
}

func TestRoRight(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "R90")

	assert.Equal(t, ferry.XDir, 0)
	assert.Equal(t, ferry.YDir, -1)

	ApplyNavigationCommand(ferry, "R90")

	assert.Equal(t, ferry.XDir, -1)
	assert.Equal(t, ferry.YDir, 0)

	ApplyNavigationCommand(ferry, "R90")

	assert.Equal(t, ferry.XDir, 0)
	assert.Equal(t, ferry.YDir, 1)
}

func TestGoNorth(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "N10")

	assert.Equal(t, ferry.XDir, 1)
	assert.Equal(t, ferry.YDir, 0)
	assert.Equal(t, ferry.XPos, 0)
	assert.Equal(t, ferry.YPos, 10)
}

func TestGoSouth(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "S10")

	assert.Equal(t, ferry.XDir, 1)
	assert.Equal(t, ferry.YDir, 0)
	assert.Equal(t, ferry.XPos, 0)
	assert.Equal(t, ferry.YPos, -10)
}

func TestGoEast(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "E10")

	assert.Equal(t, ferry.XDir, 1)
	assert.Equal(t, ferry.YDir, 0)
	assert.Equal(t, ferry.XPos, 10)
	assert.Equal(t, ferry.YPos, 0)
}

func TestGoWest(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "W10")

	assert.Equal(t, ferry.XDir, 1)
	assert.Equal(t, ferry.YDir, 0)
	assert.Equal(t, ferry.XPos, -10)
	assert.Equal(t, ferry.YPos, 0)
}

func TestGoForward(t *testing.T) {
	ferry := NewFerry(0, 0, 1, 0)

	ApplyNavigationCommand(ferry, "F10")

	assert.Equal(t, ferry.XDir, 1)
	assert.Equal(t, ferry.YDir, 0)
	assert.Equal(t, ferry.XPos, 10)
	assert.Equal(t, ferry.YPos, 0)
}
