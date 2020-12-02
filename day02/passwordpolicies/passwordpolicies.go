package passwordpolicies

import "strings"

func ApplyFirstPolicy(minTimes int, maxTimes int, char string, password string) bool {
	count := strings.Count(password, char)
	return minTimes <= count && count <= maxTimes
}

func ApplySecondPolicy(firstPos int, secondPos int, char string, password string) bool {
	firstMatch := string(password[firstPos-1]) == char
	secondMatch := string(password[secondPos-1]) == char

	return firstMatch != secondMatch
}
