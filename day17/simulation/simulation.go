package simulation

type Point struct {
	x, y, z, w int
}

func NewPoint(x int, y int, z int, w int) *Point {
	return &Point{x: x, y: y, z: z, w: w}
}

func NewPointZero() *Point {
	return NewPoint(0, 0, 0, 0)
}

type Simulation struct {
	points map[Point]int
	cycle  int
}

func NewSimulation(lines []string) *Simulation {
	m := make(map[Point]int)

	for y, line := range lines {
		for x, c := range line {
			point := NewPoint(x, y, 0, 0)
			if c == '.' {
				m[*point] = 0
			} else if c == '#' {
				m[*point] = 1
			}
		}
	}

	return &Simulation{points: m, cycle: 0}
}

func (s *Simulation) Cycle() {
	copyWithNeighbors := addNeighbors(s.points)
	newMap := make(map[Point]int)

	for p, v := range copyWithNeighbors {
		activeNeighbors := countActiveNeighbors(p, copyWithNeighbors)
		if v == 1 {
			if activeNeighbors != 2 && activeNeighbors != 3 {
				v = 0
			}
		} else {
			if activeNeighbors == 3 {
				v = 1
			}
		}
		newMap[p] = v
	}

	s.points = newMap
	s.cycle++
}

func (s *Simulation) CountActive() int {
	active := 0
	for _, v := range s.points {
		if v == 1 {
			active++
		}
	}
	return active
}

func addNeighbors(points map[Point]int) map[Point]int {
	newMap := make(map[Point]int)

	for point, _ := range points {
		for _, x := range []int{-1, 0, 1} {
			for _, y := range []int{-1, 0, 1} {
				for _, z := range []int{-1, 0, 1} {
					for _, w := range []int{-1, 0, 1} {
						newPoint := NewPoint(point.x+x, point.y+y, point.z+z, point.w+w)
						if v, ok := points[*newPoint]; ok {
							newMap[*newPoint] = v
						} else {
							newMap[*newPoint] = 0
						}
					}
				}
			}
		}
	}
	return newMap
}

func countActiveNeighbors(point Point, points map[Point]int) int {
	activeNeighbors := 0
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				for _, w := range []int{-1, 0, 1} {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						newPoint := NewPoint(point.x+x, point.y+y, point.z+z, point.w+w)
						if v, ok := points[*newPoint]; ok {
							if v == 1 {
								activeNeighbors++
							}
						}
					}
				}
			}
		}
	}

	return activeNeighbors
}
