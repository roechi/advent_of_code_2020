package boardingpass

import "math"

func FindSeat(pass string) (int, int) {
	binary := make([]int, 0)

	for _, r := range Reverse(pass) {
		if r == 'F' || r == 'L' {
			binary = append(binary, 0)
		} else if r == 'B' || r == 'R' {
			binary = append(binary, 1)
		}
	}

	row := calcBinaryVal(binary[3:])
	col := calcBinaryVal(binary[:3])

	return row, col
}

func calcBinaryVal(bits []int) int {
	sum := 0
	for i, b := range bits {
		sum += b * int(math.Pow(2, float64(i)))
	}
	return sum
}

func CalculateId(row, column int) int {
	return (8 * row) + column
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
