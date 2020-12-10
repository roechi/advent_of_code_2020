package main

import (
	"../utils"
	. "./joltage"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day10/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	onejolt, _, threeJolt := AdapterChain(ints)

	reg, _ := regexp.Compile("(11+)")

	joltageIncs := ToJoltageIncrease(ints)

	fmt.Println("JI: ", joltageIncs)

	joltageString, _ := json.Marshal(joltageIncs)
	trimmed := strings.Trim(string(joltageString), "[]")
	cleaned := strings.ReplaceAll(trimmed, ",", "")
	fmt.Println("TR: ", cleaned)

	groups := reg.FindAll([]byte(cleaned), -1)

	result := 1
	for _, g := range groups {
		gstr := string(g)
		gstrwo0 := strings.ReplaceAll(gstr, "0", "")
		gstrclean := strings.ReplaceAll(gstrwo0, "3", "")

		intgroup := make([]int, len(gstrclean))
		for i, r := range gstrclean {
			intgroup[i], _ = strconv.Atoi(string(r))
		}
		fmt.Println(intgroup)
		variations, err := CountVariations(len(intgroup))
		if err == nil {
			result *= variations
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("Solution 1: ", onejolt*threeJolt)

	fmt.Println("Solution 2: ", result)
}
