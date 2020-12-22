package main

import (
	"../utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day19/puzzle_input.txt")
	var ruleLines []string
	var linesToMatch []string

	mode := 0
	for _, line := range lines {
		if line != "" {
			if mode == 0 {
				ruleLines = append(ruleLines, line)
			} else if mode == 1 {
				linesToMatch = append(linesToMatch, line)
			}
		} else {
			mode = 1
		}
	}

	var rules = make(map[string]string)

	for _, s := range ruleLines {
		rule := strings.Split(s, ": ")
		rules[rule[0]] = rule[1]
	}

	expr := "^" + resolveRegex("0", rules) + "$"
	rgxp := regexp.MustCompile(expr)
	match := 0
	for _, line := range linesToMatch {
		if rgxp.MatchString(line) {
			match++
		}
	}
	fmt.Println("Solution 1: ", match)

	rules["8"] = `"` + resolveRegex("42", rules) + `+"`
	rules["11"] = ""
	for i := 1; i <= 42; i++ {
		rules["11"] += fmt.Sprintf("|%s{%d}%s{%d}", resolveRegex("42", rules), i, resolveRegex("31", rules), i)
	}
	rules["11"] = `"(?:` + rules["11"][1:] + `)"`

	expr2 := "^" + resolveRegex("0", rules) + "$"
	rgxp2 := regexp.MustCompile(expr2)
	match2 := 0
	for _, line := range linesToMatch {
		if rgxp2.MatchString(line) {
			match2++
		}
	}
	fmt.Println("Solution 2: ", match2)
}

func resolveRegex(rule string, rules map[string]string) (re string) {
	if rules[rule][0] == '"' {
		return rules[rule][1 : len(rules[rule])-1]
	}
	for _, s := range strings.Split(rules[rule], " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += resolveRegex(s, rules)
		}
	}
	return "(?:" + re[1:] + ")"
}
