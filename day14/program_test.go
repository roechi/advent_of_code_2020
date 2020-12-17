package main

import (
	. "./docking"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleRun(t *testing.T) {
	lines := []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	program := NewProgram()

	for _, l := range lines {
		program.ApplyCommand(l)
	}

	mem := program.GetMem()

	assert.Equal(t, mem[8], 64)
	assert.Equal(t, mem[7], 101)
}
func TestSimpleRun2(t *testing.T) {
	lines := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	program := NewProgram()

	for _, l := range lines {
		program.ApplyCommand2(l)
	}

	sum := 0

	for _, val := range program.GetMem() {
		sum += val
	}

	assert.Equal(t, 208, sum)
}
