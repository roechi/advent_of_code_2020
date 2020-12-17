package docking

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Program struct {
	mem  map[int]int
	mask []int
}

func NewProgram() *Program {
	return &Program{
		mem:  make(map[int]int),
		mask: make([]int, 36),
	}
}

func (p *Program) ApplyCommand(commandLine string) *Program {
	if strings.Contains(commandLine, "mask") {
		var mask string
		fmt.Sscanf(commandLine, "mask = %s", &mask)
		newMask := make([]int, 36)
		for i, r := range mask {
			if r != 'X' {
				newMask[i], _ = strconv.Atoi(string(r))
			} else {
				newMask[i] = -1
			}
		}
		p.mask = newMask
	} else {
		if strings.Contains(commandLine, "mem") {
			var addr, val int
			fmt.Sscanf(commandLine, "mem[%d] = %d", &addr, &val)
			bits := strconv.FormatInt(int64(val), 2)
			maskedBits := make([]int, 36)
			for i, bit := range p.mask {
				if bit != -1 {

					maskedBits[i] = bit

				} else {
					if i >= len(p.mask)-len(bits) {
						maskedBits[i], _ = strconv.Atoi(string(bits[i+len(bits)-len(maskedBits)]))
					}
				}
			}
			valToWrite := 0
			for i, b := range maskedBits {
				valToWrite += int(math.Pow(2, float64(35-i))) * b
			}
			p.mem[addr] = valToWrite
		}
	}

	return p
}

func (p *Program) ApplyCommand2(commandLine string) *Program {
	if strings.Contains(commandLine, "mask") {
		var mask string
		fmt.Sscanf(commandLine, "mask = %s", &mask)
		newMask := make([]int, 36)
		for i, r := range mask {
			if r != 'X' {
				newMask[i], _ = strconv.Atoi(string(r))
			} else {
				newMask[i] = -1
			}
		}
		p.mask = newMask
	} else {
		if strings.Contains(commandLine, "mem") {
			var addr, val int
			fmt.Sscanf(commandLine, "mem[%d] = %d", &addr, &val)
			bits := strconv.FormatInt(int64(addr), 2)

			maskedBits := make([]int, 36)
			for i, bit := range p.mask {
				if bit == -1 {
					maskedBits[i] = bit
				} else if bit == 1 {
					maskedBits[i] = 1
				} else {
					if i >= len(p.mask)-len(bits) {
						maskedBits[i], _ = strconv.Atoi(string(bits[i+len(bits)-len(maskedBits)]))
					}
				}
			}
			addresses := make([][]int, 0)
			adds := GetPossibleAddresses(0, maskedBits, addresses)
			for _, addr := range adds {
				targetAddr := 0
				for i, b := range addr {
					targetAddr += int(math.Pow(2, float64(35-i))) * b
				}
				p.mem[targetAddr] = val
			}
		}
	}

	return p
}

func GetPossibleAddresses(currBit int, origAddress []int, addresses [][]int) [][]int {
	if currBit == len(origAddress) {
		return append(addresses, origAddress)
	}
	if origAddress[currBit] == -1 {
		add1 := make([]int, len(origAddress))
		copy(add1, origAddress)
		add1[currBit] = 0
		add2 := make([]int, len(origAddress))
		copy(add2, origAddress)
		add2[currBit] = 1

		possibleAddresses1 := GetPossibleAddresses(currBit+1, add1, addresses)

		possibleAddresses2 := GetPossibleAddresses(currBit+1, add2, addresses)

		allPossibleAddresses := make([][]int, 0)
		for _, add := range possibleAddresses1 {
			allPossibleAddresses = append(allPossibleAddresses, add)
		}
		for _, add := range possibleAddresses2 {
			allPossibleAddresses = append(allPossibleAddresses, add)
		}

		return allPossibleAddresses
	} else {
		return GetPossibleAddresses(currBit+1, origAddress, addresses)
	}
}

func (p *Program) GetMem() map[int]int {
	return p.mem
}

func (p *Program) GetMask() []int {
	return p.mask
}
