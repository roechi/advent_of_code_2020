package main

import (
	. "./specialmath"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSolvePrecedence(t *testing.T) {
	line := "1 + 2 * 3 + 4 * 5 + 6"
	want := 71
	got := Solve(line)

	assert.Equal(t, got, want)
}

func TestSolveWithParenthesis(t *testing.T) {
	line := "1 + (2 * 3) + (4 * (5 + 6))"
	want := 51
	got := SolveWithParenthesis(line)

	assert.Equal(t, got, want)
}

func TestSolveWithParenthesisSpecial(t *testing.T) {
	line := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	want := 23340
	got := SolveWithParenthesisSpecial(line)

	assert.Equal(t, got, want)
}
