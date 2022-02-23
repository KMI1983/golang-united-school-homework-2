package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import (
	"math"
)

type intCustomType int

const (
	SidesCircle   intCustomType = 0
	SidesTriangle               = 3
	SidesSquare                 = 4
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {

	switch sidesNum {
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	case SidesTriangle:
		return (sideLen * sideLen * math.Sqrt(3)) / 4
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
