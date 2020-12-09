package encoding

func FindFirstFaulty(nums []int, preambleLength int) int {

	for i, val := range nums[preambleLength:] {
		posVal := false
		preamble := nums[i : i+preambleLength]
		for j, check := range preamble {
			remainder := val - check
			valsToCheck := make([]int, 0)
			if j < len(preamble)-1 {
				valsToCheck = append(valsToCheck, preamble[:j]...)
				valsToCheck = append(valsToCheck, preamble[j+1:]...)
			} else {
				valsToCheck = append(valsToCheck, preamble[:j]...)
			}

			if contains(valsToCheck, remainder) {
				posVal = true
				break
			}
		}
		if !posVal {
			return i + preambleLength
		}
	}
	return -1
}

func FindContiguousSet(nums []int, val int) []int {
	i := 2
	jFound := -1
	iFound := -1

	for i < len(nums) {
		j := 0
		for j < len(nums)-i {
			sum := Sum(nums[j : j+i])
			if sum == val {
				jFound = j
				break
			}
			j++
		}
		if jFound > -1 {
			iFound = i
			break
		}
		i++
	}

	if jFound > -1 {
		return nums[jFound : jFound+iFound]
	}

	return nil
}

func contains(s []int, c int) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}
	return false
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
