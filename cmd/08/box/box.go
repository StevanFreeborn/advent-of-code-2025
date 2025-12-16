// Package box provides a model and methods for junction boxes.
package box

import (
	"math"
	"strconv"
	"strings"
)

type Box interface {
	Y() int
	X() int
	Z() int
	DistanceFrom(neighbor Box) float64
}

type box struct {
	x int
	y int
	z int
}

func (b box) X() int {
	return b.x
}

func (b box) Y() int {
	return b.y
}

func (b box) Z() int {
	return b.z
}

func (b box) DistanceFrom(neighbor Box) float64 {
	xDiff := math.Pow(float64(b.x-neighbor.X()), 2)
	yDiff := math.Pow(float64(b.y-neighbor.Y()), 2)
	zDiff := math.Pow(float64(b.z-neighbor.Z()), 2)

	return math.Sqrt(xDiff + yDiff + zDiff)
}

func From(str string) Box {
	parts := strings.Split(str, ",")

	if len(parts) != 3 {
		panic("invalid junction box")
	}

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return box{
		x: x,
		y: y,
		z: z,
	}
}
