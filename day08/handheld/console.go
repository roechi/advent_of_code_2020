package handheld

import (
	"errors"
	"fmt"
)

type Console struct {
	acc  int
	code []string
	pnt  int
	vpos []int
	seen map[int]struct{}
}

func NewConsole(code []string) *Console {
	return &Console{acc: 0, code: code, pnt: 0, vpos: make([]int, 0), seen: map[int]struct{}{}}
}

func (c *Console) RunProgram() (int, error) {

	for c.pnt < len(c.code) {
		if _, ok := c.seen[c.pnt]; ok {
			return c.acc, errors.New("encountered infinite loop")
		}
		c.seen[c.pnt] = struct{}{}

		c.RunNextInstruction()
	}

	return c.acc, nil
}

func (c *Console) RunNextInstruction() {

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

	c.vpos = append(c.vpos, c.pnt)
}

func (c *Console) GetPnt() int {
	return c.pnt
}

func (c *Console) GetAcc() int {
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
