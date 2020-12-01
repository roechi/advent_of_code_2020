package expenses


func CalcExpenses(items []int) int {
	for _, i := range items {
		remainder := 2020 - i
		if contains(items, remainder) {
			return i * remainder
		}
	}
	return 0
}

func CalcExpenses2(items []int) int {
	for x, i := range items {
		for y, j := range items {
			remainder := 2020 - i - j
			if x != y && contains(items, remainder) {
				return i * j * remainder
			}
		}
	}
	return 0
}

func contains(s []int, c int) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}
	return false
}