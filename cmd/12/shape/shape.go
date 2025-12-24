// Package shape provides a model and methods for representing and manipulating shapes
package shape

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

type Shape interface {
	Id() int
	GetKey() string
	GenerateVariants() []Shape
	RotateClockwise() Shape
	FlipVertically() Shape
	String() string
	Area() int
	MaxRow() int
	MaxColumn() int
	Positions() []position.Position
}

type shape struct {
	id        int
	positions []position.Position
}

func From(lines []string) Shape {
	idStr := strings.TrimSuffix(lines[0], ":")
	id, _ := strconv.Atoi(idStr)

	positions := []position.Position{}

	for r, line := range lines[1:] {
		for c, char := range line {
			if char != '#' {
				continue
			}

			position := position.From(r, c)
			positions = append(positions, position)
		}
	}

	return shape{id: id, positions: positions}
}

func (s shape) Positions() []position.Position {
	return s.positions
}

func (s shape) MaxRow() int {
	maxRow := 0

	for _, p := range s.positions {
		if p.Row() > maxRow {
			maxRow = p.Row()
		}
	}

	return maxRow
}

func (s shape) MaxColumn() int {
	maxColumn := 0

	for _, p := range s.positions {
		if p.Column() > maxColumn {
			maxColumn = p.Column()
		}
	}

	return maxColumn
}

func (s shape) Area() int {
	return len(s.positions)
}

func (s shape) GenerateVariants() []Shape {
	shapes := []Shape{}
	var current Shape
	current = s
	numberOfTransformations := 4

	for range numberOfTransformations {
		shapes = append(shapes, current)
		flipped := current.FlipVertically()
		shapes = append(shapes, flipped)
		current = current.RotateClockwise()
	}

	// // Rotate 90 clockwise
	// for range numberOfTransformations {
	// 	current = current.RotateClockwise()
	// 	shapes = append(shapes, current)
	// }
	//
	// current = s
	//
	// // Flip then rotate 90 clockwise
	// for range numberOfTransformations {
	// 	shapes = append(shapes, current)
	// 	flipped := current.FlipVertically()
	// 	shapes = append(shapes, flipped)
	// 	current = flipped.RotateClockwise()
	// }
	//
	// current = s
	//
	// // Rotate 90 clockwise then flip
	// for range numberOfTransformations {
	// 	shapes = append(shapes, current)
	// 	rotated := current.RotateClockwise()
	// 	shapes = append(shapes, rotated)
	// 	current = rotated.FlipVertically()
	// }
	//
	// current = s
	// rotated := current.RotateClockwise().RotateClockwise()
	// flipped := rotated.FlipVertically()
	// shapes = append(shapes, flipped)

	seenShapes := map[string]Shape{}

	for _, s := range shapes {
		seenShapes[s.GetKey()] = s
	}

	return slices.Collect(maps.Values(seenShapes))
}

func (s shape) RotateClockwise() Shape {
	rotatedPositions := []position.Position{}

	for _, p := range s.positions {
		rotatedPosition := position.From(p.Column(), -p.Row())
		rotatedPositions = append(rotatedPositions, rotatedPosition)
	}

	rotatedShape := shape{
		id:        s.id,
		positions: rotatedPositions,
	}

	return normalize(rotatedShape)
}

func (s shape) FlipVertically() Shape {
	flippedPositions := []position.Position{}

	for _, p := range s.positions {
		flippedPosition := position.From(p.Row(), -p.Column())
		flippedPositions = append(flippedPositions, flippedPosition)
	}

	flippedShape := shape{
		id:        s.id,
		positions: flippedPositions,
	}

	return normalize(flippedShape)
}

func (s shape) GetKey() string {
	var sb strings.Builder

	for _, p := range s.positions {
		fmt.Fprintf(&sb, "%d,%d|", p.Row(), p.Column())
	}

	return sb.String()
}

func (s shape) Id() int {
	return s.id
}

func normalize(s shape) Shape {
	if len(s.positions) == 0 {
		return shape{
			id:        s.id,
			positions: s.positions,
		}
	}

	minRow := s.positions[0].Row()
	minColumn := s.positions[0].Column()

	for _, p := range s.positions {
		if p.Row() < minRow {
			minRow = p.Row()
		}

		if p.Column() < minColumn {
			minColumn = p.Column()
		}
	}

	nps := []position.Position{}

	for _, p := range s.positions {
		nr := p.Row() - minRow
		nc := p.Column() - minColumn
		np := position.From(nr, nc)
		nps = append(nps, np)
	}

	return shape{
		id:        s.id,
		positions: nps,
	}
}

func (s shape) String() string {
	if len(s.positions) == 0 {
		return "(EMPTY SHAPE)"
	}

	maxRow := 0
	maxColumn := 0

	for _, p := range s.positions {
		if p.Row() > maxRow {
			maxRow = p.Row()
		}

		if p.Column() > maxColumn {
			maxColumn = p.Column()
		}
	}

	tempGrid := make([][]string, maxRow+1)

	for r := range tempGrid {
		tempGrid[r] = make([]string, maxColumn+1)

		for c := range tempGrid[r] {
			tempGrid[r][c] = "."
		}
	}

	for _, p := range s.positions {
		tempGrid[p.Row()][p.Column()] = "#"
	}

	var shape strings.Builder

	for _, row := range tempGrid {
		var rowString strings.Builder

		for _, str := range row {
			rowString.WriteString(str)
		}

		shape.WriteString(rowString.String() + "\n")
	}

	return shape.String()
}
