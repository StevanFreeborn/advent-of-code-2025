// Package rectangle provides a model and methods for representing and manipulating rectangles.
package rectangle

import (
	"math"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

type Rectangle interface {
	Corners() []position.Position
	MaxRow() int
	MaxColumn() int
	MinRow() int
	MinColumn() int
	Area() int
	Center() position.Position
}

type rectangle struct {
	corners   []position.Position
	maxRow    int
	maxColumn int
	minRow    int
	minColumn int
	area      int
	center    position.Position
}

func From(cornerOne position.Position, cornerTwo position.Position) Rectangle {
	thirdCorner := position.From(cornerOne.Row(), cornerTwo.Column())
	fourthCorner := position.From(cornerTwo.Row(), cornerOne.Column())

	corners := []position.Position{cornerOne, cornerTwo, thirdCorner, fourthCorner}
	maxRow := math.Max(float64(cornerOne.Row()), float64(cornerTwo.Row()))
	maxColumn := math.Max(float64(cornerOne.Column()), float64(cornerTwo.Column()))
	minRow := math.Min(float64(cornerOne.Row()), float64(cornerTwo.Row()))
	minColumn := math.Min(float64(cornerOne.Column()), float64(cornerTwo.Column()))

	area := int((maxRow - minRow + 1) * (maxColumn - minColumn + 1))

	centerRow := (maxRow + minRow) / 2
	centerColumn := (maxColumn + minColumn) / 2
	center := position.From(int(centerRow), int(centerColumn))

	return rectangle{
		corners:   corners,
		maxRow:    int(maxRow),
		maxColumn: int(maxColumn),
		minRow:    int(minRow),
		minColumn: int(minColumn),
		area:      area,
		center:    center,
	}
}

func (r rectangle) Corners() []position.Position {
	return r.corners
}

func (r rectangle) MaxRow() int {
	return r.maxRow
}

func (r rectangle) MinRow() int {
	return r.minRow
}

func (r rectangle) MaxColumn() int {
	return r.maxColumn
}

func (r rectangle) MinColumn() int {
	return r.minColumn
}

func (r rectangle) Area() int {
	return r.area
}

func (r rectangle) Center() position.Position {
	return r.center
}
