package sled

func ScanSlopeForObstacles(xSlope int, ySlope int, terrain [][]int) int {
	yLen := len(terrain)
	xLen := len(terrain[0])

	curX := 0
	curY := 0

	encObstacles := 0
	for {
		if curY == yLen-1 {
			break
		} else {
			curX += xSlope
			curY += ySlope

			if curX >= xLen {
				curX -= xLen
			}
		}
		if curY < yLen {
			if terrain[curY][curX] == 1 {
				encObstacles++
			}
		}
	}
	return encObstacles
}
