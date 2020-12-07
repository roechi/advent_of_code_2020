package customsanswers

import (
	. "../../utils/stringset"
)

func CountUniqueAnswers(groupAnswers []string) int {
	uniqueAnswers := NewStringSet()
	for _, p := range groupAnswers {
		for _, ans := range p {
			uniqueAnswers.Add(string(ans))
		}
	}
	return uniqueAnswers.Length()
}

func CountAllYes(groupAnswers []string) int {
	var inter *StringSet

	for _, ans := range groupAnswers {
		set := NewStringSet()
		for _, r := range ans {
			set.Add(string(r))
		}
		if inter == nil {
			inter = set
		} else {
			inter = inter.Intersection(set)
		}
	}

	return inter.Length()
}
