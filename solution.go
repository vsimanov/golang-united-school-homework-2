package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type Sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle Sides = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch {
	case sidesNum == 0:
		return math.Pi * math.Pow(sideLen, 2)
	case sidesNum == 3:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case sidesNum == 4:
		return math.Pow(sideLen, 2)
	default:
		return 0
	}

}
