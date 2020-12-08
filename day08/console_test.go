package main

import (
	. "./handheld"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test(t *testing.T) {
	code := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	console := NewConsole(code)

	console.RunNextInstruction()
	console.RunNextInstruction()
	console.RunNextInstruction()

	assert.Equal(t, console.GetPnt(), 6)
	assert.Equal(t, console.GetAcc(), 1)
}
