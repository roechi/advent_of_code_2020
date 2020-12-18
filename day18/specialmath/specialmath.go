package specialmath

import (
	"regexp"
	"strconv"
	"strings"
)

func SolveWithParenthesis(line string) int {
	open := strings.LastIndex(line, "(")
	if open != -1 {
		close := open + strings.Index(line[open:], ")")
		subSolution := Solve(line[open+1 : close])
		newLine := line[:open] + strconv.Itoa(subSolution) + line[close+1:]
		return SolveWithParenthesis(newLine)
	} else {
		return Solve(line)
	}
}

func SolveWithParenthesisSpecial(line string) int {
	open := strings.LastIndex(line, "(")
	if open != -1 {
		close := open + strings.Index(line[open:], ")")
		subSolution := SolveSpecial(line[open+1 : close])
		newLine := line[:open] + strconv.Itoa(subSolution) + line[close+1:]
		return SolveWithParenthesisSpecial(newLine)
	} else {
		return SolveSpecial(line)
	}
}

func Solve(line string) int {
	rgxp, _ := regexp.Compile("([0-9]+)|(\\*)|(\\+)")
	groups := rgxp.FindAllString(line, -1)

	currResult := MustAtoi(groups[0])
	for i := 0; i < len(groups)-2; i += 2 {
		if (groups[i+1]) == "+" {
			currResult += MustAtoi(groups[i+2])
		} else if groups[i+1] == "*" {
			currResult *= MustAtoi(groups[i+2])
		}
	}
	return currResult
}

func SolveSpecial(line string) int {
	rgxp, _ := regexp.Compile("([0-9]+)|(\\*)|(\\+)")
	groups := rgxp.FindAllString(line, -1)

	occAdd := Find(groups, "+")
	occMul := Find(groups, "*")
	if occAdd >= 0 {
		sub := groups[occAdd-1] + groups[occAdd] + groups[occAdd+1]
		subSolution := Solve(sub)
		var reduced string
		for i := 0; i < occAdd-1; i++ {
			reduced += groups[i]
		}
		reduced += strconv.Itoa(subSolution)
		for i := occAdd + 2; i < len(groups); i++ {
			reduced += groups[i]
		}
		return SolveSpecial(reduced)
	} else if occMul >= 0 {
		sub := groups[occMul-1] + groups[occMul] + groups[occMul+1]
		subSolution := Solve(sub)
		var reduced string
		for i := 0; i < occMul-1; i++ {
			reduced += groups[i]
		}
		reduced += strconv.Itoa(subSolution)
		for i := occMul + 2; i < len(groups); i++ {
			reduced += groups[i]
		}
		return SolveSpecial(reduced)
	} else {
		return MustAtoi(line)
	}
}

func MustAtoi(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return num
}

func Find(sSlice []string, s string) int {
	for i, val := range sSlice {
		if val == s {
			return i
		}
	}
	return -1
}
