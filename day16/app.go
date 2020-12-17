package main

import (
	"../utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day16/puzzle_input.txt")
	rules, otherTickets, myTicket := Parse(lines)

	fmt.Println("Solution 1: ", SolvePart1(rules, otherTickets))
	fmt.Println("Solution 2: ", SolvePart2(rules, otherTickets, myTicket))
}

func SolvePart1(rules []Rule, otherTickets [][]int) int {
	_, sumBadValues := getBadTickets(rules, otherTickets)
	return sumBadValues
}

func SolvePart2(rules []Rule, otherTickets [][]int, myTicket []int) int {

	idxs, _ := getBadTickets(rules, otherTickets)
	validOtherTickets := RemoveTickets(idxs, otherTickets)

	fieldToPossibleIdx := map[string]map[int]bool{}
	for _, field := range rules {
		idxs := map[int]bool{}
		for idx := range validOtherTickets[0] {
			idxs[idx] = true
		}
		for _, ticket := range validOtherTickets {
			for idx, i := range ticket {
				idxs[idx] = idxs[idx] && field.IsValid(i)
			}
		}
		fieldToPossibleIdx[field.Text] = idxs
	}

	fieldToIdx := map[string]int{}
	for {
		idx := -1
		for field, possibleIdx := range fieldToPossibleIdx {
			if size(possibleIdx) == 1 {
				idx = toSlice(possibleIdx)[0]
				fieldToIdx[field] = idx
			}
		}
		if idx < 0 {
			break

		}
		for _, possibleIdx := range fieldToPossibleIdx {
			possibleIdx[idx] = false
		}
	}

	o := 1
	for _, field := range rules {
		if strings.HasPrefix(field.Text, "departure ") {
			o *= myTicket[fieldToIdx[field.Text]]
		}
	}
	return o
}

func RemoveTickets(ids []int, otherTickets [][]int) [][]int {
	m := map[int]bool{}
	for _, idx := range ids {
		m[idx] = true
	}
	var cleanedTickets [][]int
	for idx, ticket := range otherTickets {
		if _, exists := m[idx]; exists {
			continue
		}
		cleanedTickets = append(cleanedTickets, ticket)
	}
	return cleanedTickets
}

type Rule struct {
	Text   string
	Limits []int
}

func (f Rule) IsValid(i int) bool {
	return (i >= f.Limits[0] && i <= f.Limits[1]) || (i >= f.Limits[2] && i <= f.Limits[3])
}

func Parse(lines []string) ([]Rule, [][]int, []int) {
	parts := strings.Split(strings.Join(lines, "\n"), "\n\n")

	var rules []Rule
	for _, line := range strings.Split(parts[0], "\n") {
		rgxp, err := regexp.Compile(`([a-z\s]+):|(\d+)`)
		if err != nil {
			log.Fatal(err)
		}
		groups := rgxp.FindAllString(line, -1)
		rules = append(rules, Rule{
			Text:   groups[0],
			Limits: []int{AsInt(groups[1]), AsInt(groups[2]), AsInt(groups[3]), AsInt(groups[4])},
		})
	}

	var myTicket []int
	for _, s := range strings.Split(strings.Split(parts[1], "\n")[1], ",") {
		myTicket = append(myTicket, AsInt(s))
	}

	var otherTickets [][]int
	for _, line := range strings.Split(parts[2], "\n")[1:] {
		var otherTicket []int
		for _, s := range strings.Split(line, ",") {
			otherTicket = append(otherTicket, AsInt(s))
		}
		otherTickets = append(otherTickets, otherTicket)
	}

	return rules, otherTickets, myTicket
}

func getBadTickets(rules []Rule, otherTickets [][]int) (idxs []int, sumBadValues int) {
	idxs, sumBadValues = []int{}, 0
	for idx, ticket := range otherTickets {
		badTicket := false
		for _, i := range ticket {
			validity := map[bool]int{true: 0, false: 0}
			for _, field := range rules {
				validity[field.IsValid(i)]++
			}
			if validity[true] == 0 {
				badTicket = true
				sumBadValues += i
			}
		}
		if badTicket {
			idxs = append(idxs, idx)
		}
	}
	return
}

func size(m map[int]bool) int {
	o := 0
	for _, v := range m {
		if v {
			o++
		}
	}
	return o
}

func toSlice(m map[int]bool) []int {
	o := []int{}
	for k, v := range m {
		if v {
			o = append(o, k)
		}
	}
	return o
}

func AsInt(val string) int {
	atoi, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return atoi
}
