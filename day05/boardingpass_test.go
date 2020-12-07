package main

import (
	. "./boardingpass"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCalcSeat(t *testing.T) {
	pass := "BFFFBBFRRR"
	row := 70
	column := 7
	seatID := 567

	seatRow, seatCol := FindSeat(pass)

	assert.Equal(t, seatRow, row)
	assert.Equal(t, seatCol, column)

	gotId := CalculateId(seatRow, seatCol)

	assert.Equal(t, gotId, seatID)
}

func TestCalcSeat1(t *testing.T) {
	pass := "FFFBBBFRRR"
	row := 14
	column := 7
	seatID := 119

	seatRow, seatCol := FindSeat(pass)

	assert.Equal(t, seatRow, row)
	assert.Equal(t, seatCol, column)

	gotId := CalculateId(seatRow, seatCol)

	assert.Equal(t, gotId, seatID)
}

func TestCalcSeat2(t *testing.T) {
	pass := "BBFFBBFRLL"
	row := 102
	column := 4
	seatID := 820

	seatRow, seatCol := FindSeat(pass)

	assert.Equal(t, seatRow, row)
	assert.Equal(t, seatCol, column)

	gotId := CalculateId(seatRow, seatCol)

	assert.Equal(t, gotId, seatID)
}
