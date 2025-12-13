// Package position provides types and functions to represent and manipulate positions and moves on a 2D grid.
package position

import "github.com/StevanFreeborn/advent-of-code-2025/internal/move"

type Position interface {
	Row() int
	Column() int
	Move(move.Move) Position
}

type position struct {
	row    int
	column int
}

func From(row int, column int) Position {
	return position{
		row:    row,
		column: column,
	}
}

func (p position) Row() int {
	return p.row
}

func (p position) Column() int {
	return p.column
}

func (p position) Move(move move.Move) Position {
	return From(p.row+move.NumberOfRows(), p.column+move.NumberOfColumns())
}
