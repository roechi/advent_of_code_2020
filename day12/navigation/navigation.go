package navigation

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

type Ferry struct {
	XPos, YPos, XDir, YDir int
}

type Ferry2 struct {
	XPos, YPos, XWay, YWay int
}

func NewFerry(xPos, yPos, xDir, yDir int) *Ferry {
	return &Ferry{XPos: xPos, YPos: yPos, XDir: xDir, YDir: yDir}
}

func NewFerry2(xPos, yPos, xWay, yWay int) *Ferry2 {
	return &Ferry2{XPos: xPos, YPos: yPos, XWay: xWay, YWay: yWay}
}

func ApplyNavigationCommand(ferry *Ferry, command string) {
	var commandKey string
	var commandVal int

	fmt.Sscanf(command, "%1s%d", &commandKey, &commandVal)

	switch commandKey {
	case "F":
		ferry.XPos += ferry.XDir * commandVal
		ferry.YPos += ferry.YDir * commandVal
		break
	case "N":
		ferry.YPos += commandVal
		break
	case "S":
		ferry.YPos -= commandVal
		break
	case "E":
		ferry.XPos += commandVal
		break
	case "W":
		ferry.XPos -= commandVal
		break
	case "L":
		rot := mat.NewDense(2, 2, []float64{0, -1, 1, 0})
		dir := mat.NewVecDense(2, []float64{float64(ferry.XDir), float64(ferry.YDir)})
		i := 0
		for i < commandVal/90 {
			newDir := mat.NewVecDense(2, make([]float64, 2))
			newDir.MulVec(rot, dir)
			dir = newDir
			i++
		}
		ferry.XDir = int(dir.At(0, 0))
		ferry.YDir = int(dir.At(1, 0))
	case "R":
		rot := mat.NewDense(2, 2, []float64{0, -1, 1, 0})
		dir := mat.NewVecDense(2, []float64{float64(ferry.XDir), float64(ferry.YDir)})
		i := 0
		for i < 4-(commandVal/90) {
			newDir := mat.NewVecDense(2, make([]float64, 2))
			newDir.MulVec(rot, dir)
			dir = newDir
			i++
		}
		ferry.XDir = int(dir.At(0, 0))
		ferry.YDir = int(dir.At(1, 0))
	}
}

func ApplyNavigationCommand2(ferry *Ferry2, command string) {
	var commandKey string
	var commandVal int

	fmt.Sscanf(command, "%1s%d", &commandKey, &commandVal)

	switch commandKey {
	case "F":
		ferry.XPos += ferry.XWay * commandVal
		ferry.YPos += ferry.YWay * commandVal
		break
	case "N":
		ferry.YWay += commandVal
		break
	case "S":
		ferry.YWay -= commandVal
		break
	case "E":
		ferry.XWay += commandVal
		break
	case "W":
		ferry.XWay -= commandVal
		break
	case "L":
		rot := mat.NewDense(2, 2, []float64{0, -1, 1, 0})
		dir := mat.NewVecDense(2, []float64{float64(ferry.XWay), float64(ferry.YWay)})
		i := 0
		for i < commandVal/90 {
			newDir := mat.NewVecDense(2, make([]float64, 2))
			newDir.MulVec(rot, dir)
			dir = newDir
			i++
		}
		ferry.XWay = int(dir.At(0, 0))
		ferry.YWay = int(dir.At(1, 0))
	case "R":
		rot := mat.NewDense(2, 2, []float64{0, -1, 1, 0})
		dir := mat.NewVecDense(2, []float64{float64(ferry.XWay), float64(ferry.YWay)})
		i := 0
		for i < 4-(commandVal/90) {
			newDir := mat.NewVecDense(2, make([]float64, 2))
			newDir.MulVec(rot, dir)
			dir = newDir
			i++
		}
		ferry.XWay = int(dir.At(0, 0))
		ferry.YWay = int(dir.At(1, 0))
	}
}

func ManhattanDist(ferry *Ferry) float64 {
	return math.Abs(float64(ferry.XPos)) + math.Abs(float64(ferry.YPos))
}

func ManhattanDist2(ferry *Ferry2) float64 {
	return math.Abs(float64(ferry.XPos)) + math.Abs(float64(ferry.YPos))
}
