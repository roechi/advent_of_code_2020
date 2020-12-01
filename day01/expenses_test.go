package main

import (
	. "./expenses"
	"github.com/magiconair/properties/assert"
	"testing"
)


func Test(t *testing.T) {
	ints := []int{976, 1044}
	want := 1018944

	got := CalcExpenses(ints)

	assert.Equal(t, want, got)
}

func TestExpenses2(t *testing.T) {
	ints := []int{979, 366, 675}

	want := 241861950

	got := CalcExpenses2(ints)

	assert.Equal(t, want, got)
}

