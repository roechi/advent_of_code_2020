package emulator

import (
	"errors"
	"fmt"
)

type Emulator struct {
	acc  int
	code []string
	pnt  int
	seen map[int]struct{}
}

func NewConsole(code []string) *Emulator {
	return &Emulator{acc: 0, code: code, pnt: 0, seen: map[int]struct{}{}}
}

func (c *Emulator) RunProgram() (int, error) {

	for c.pnt < len(c.code) {
		_, ok := c.seen[c.pnt]
		if ok {
			return c.acc, errors.New("encountered infinite loop")
		}
		c.seen[c.pnt] = struct{}{}

		c.RunNextInstruction()
	}

	return c.acc, nil
}

func (c *Emulator) RunNextInstruction() {

	var instrCode string
	var val int

	fmt.Sscanf(c.code[c.pnt], "%3s %d", &instrCode, &val)

	switch instrCode {
	case "nop":
		if val+c.pnt >= len(c.code) {
			fmt.Println("#### Change nop to jmp at: ", c.pnt)
		}
		c.pnt++
	case "acc":
		c.acc += val
		c.pnt++
	case "jmp":
		c.pnt += val
	}
}

func (c *Emulator) GetPnt() int {
	return c.pnt
}

func (c *Emulator) GetAcc() int {
	return c.acc
}

func contains(s []int, v int) bool {
	for _, v := range s {
		if v == v {
			return true
		}
	}

	return false
}
