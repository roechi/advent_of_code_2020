package joltage

import (
	"errors"
	"fmt"
	"sort"
)

func AdapterChain(ratings []int) (int, int, int) {
	sort.Ints(ratings)
	fmt.Println("Sorted: ", ratings)
	oneJoltDiff := 0
	twoJoltDiff := 0
	threeJoltDiff := 1

	if ratings[0] == 1 {
		oneJoltDiff++
	} else if ratings[0] == 3 {
		threeJoltDiff++
	} else {
		twoJoltDiff++
	}

	joltageDiffs := make([]int, len(ratings))
	joltageDiffs[0] = ratings[0]
	for i, rating := range ratings {
		if i < len(ratings)-1 {
			diff := ratings[i+1] - rating
			if diff == 1 {
				oneJoltDiff++
				joltageDiffs[i+1] = 1
			} else if diff == 3 {
				threeJoltDiff++
				joltageDiffs[i+1] = 3
			} else {
				twoJoltDiff++
				joltageDiffs[i+1] = 3
			}
		}
	}

	fmt.Println("Joltage diffs: ", joltageDiffs)
	return oneJoltDiff, twoJoltDiff, threeJoltDiff
}

func ToJoltageIncrease(ratings []int) []int {
	sort.Ints(ratings)

	joltageIncs := make([]int, len(ratings)+2)

	joltageIncs[0] = 0
	joltageIncs[1] = ratings[0]
	joltageIncs[len(joltageIncs)-1] = 3

	for i, rating := range ratings {
		if i < len(ratings)-1 {
			joltageIncs[i+2] = ratings[i+1] - rating
		}
	}

	return joltageIncs
}

func CountVariations(groupLen int) (int, error) {
	if groupLen == 2 {
		return 2, nil
	} else if groupLen == 3 {
		return 4, nil
	} else if groupLen == 4 {
		return 7, nil
	} else {
		return -1, errors.New("invalid group size " + string(groupLen))
	}
}
