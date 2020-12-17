package seating

type Seating struct {
	area [][]rune
}

func NewSeating(lines []string) *Seating {
	area := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		area[i] = make([]rune, len(lines[0]))
	}

	for y, line := range lines {
		for x, rune := range line {
			area[y][x] = rune
		}
	}

	return &Seating{area: area}
}

func (s *Seating) ApplyRound() bool {
	stateChanged := false
	carea := make([][]rune, len(s.area))
	for i := 0; i < len(s.area); i++ {
		carea[i] = make([]rune, len(s.area[0]))
	}
	for i, l := range s.area {
		copy(carea[i], l)
	}

	for y, row := range s.area {
		for x, pl := range row {
			if s.area[y][x] != '.' {
				occs := 0
				if y > 0 {
					if s.area[y-1][x] == '#' {
						occs++
					}
					if x > 0 {
						if s.area[y-1][x-1] == '#' {
							occs++
						}
					}
					if x < len(row)-1 {
						if s.area[y-1][x+1] == '#' {
							occs++
						}
					}
				}
				if x < len(row)-1 {
					if s.area[y][x+1] == '#' {
						occs++
					}
				}
				if y < len(s.area)-1 {
					if s.area[y+1][x] == '#' {
						occs++
					}
					if x > 0 {
						if s.area[y+1][x-1] == '#' {
							occs++
						}
					}
					if x < len(row)-1 {
						if s.area[y+1][x+1] == '#' {
							occs++
						}
					}
				}
				if x > 0 {
					if s.area[y][x-1] == '#' {
						occs++
					}
				}

				if pl == 'L' && occs == 0 {
					carea[y][x] = '#'
					stateChanged = true
				} else if pl == '#' && occs >= 4 {
					carea[y][x] = 'L'
					stateChanged = true
				}
			}
		}
	}

	s.area = carea

	return stateChanged
}

func (s *Seating) ApplyRound2() bool {
	stateChanged := false
	carea := make([][]rune, len(s.area))
	for i := 0; i < len(s.area); i++ {
		carea[i] = make([]rune, len(s.area[0]))
	}
	for i, l := range s.area {
		copy(carea[i], l)
	}

	for y, row := range s.area {
		for x, pl := range row {
			if s.area[y][x] != '.' {
				occs := 0

				if checkDirectionOcc(y, x, -1, -1, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, -1, 0, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, -1, 1, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, 0, 1, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, 1, 1, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, 1, 0, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, 1, -1, s.area) {
					occs++
				}
				if checkDirectionOcc(y, x, 0, -1, s.area) {
					occs++
				}

				if pl == 'L' && occs == 0 {
					carea[y][x] = '#'
					stateChanged = true
				} else if pl == '#' && occs >= 5 {
					carea[y][x] = 'L'
					stateChanged = true
				}
			}
		}
	}

	s.area = carea

	return stateChanged
}

func checkDirectionOcc(posY, posX, dirY, dirX int, area [][]rune) bool {
	occd := false
	for {
		posY += dirY
		posX += dirX

		if posX < 0 || posY < 0 || posX >= len(area[0]) || posY >= len(area) {
			break
		} else {
			if area[posY][posX] == '#' {
				occd = true
				break
			} else if area[posY][posX] == 'L' {
				break
			}
		}
	}
	return occd
}

func (s *Seating) CountOccs() int {
	occs := 0
	for _, row := range s.area {
		for _, rune := range row {
			if rune == '#' {
				occs++
			}
		}
	}
	return occs
}

func (s *Seating) GetArea() [][]string {
	area := make([][]string, len(s.area))
	for i := 0; i < len(s.area); i++ {
		area[i] = make([]string, len(s.area[0]))
	}

	for y, line := range s.area {
		for x, rune := range line {
			area[y][x] = string(rune)
		}
	}

	return area
}
