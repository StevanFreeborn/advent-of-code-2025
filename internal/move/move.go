// Package move provides types and functions to represent and manipulate moves on a 2D grid.
package move

type Move interface {
	NumberOfRows() int
	NumberOfColumns() int
}

type move struct {
	numberOfRows    int
	numberOfColumns int
}

func (m move) NumberOfRows() int {
	return m.numberOfRows
}

func (m move) NumberOfColumns() int {
	return m.numberOfColumns
}

var (
	Up       = move{numberOfRows: -1, numberOfColumns: 0}
	Down     = move{numberOfRows: 1, numberOfColumns: 0}
	Right    = move{numberOfRows: 0, numberOfColumns: 1}
	Left     = move{numberOfRows: 0, numberOfColumns: -1}
	cardinal = []Move{Up, Down, Right, Left}

	UpRight   = move{numberOfRows: -1, numberOfColumns: 1}
	UpLeft    = move{numberOfRows: -1, numberOfColumns: -1}
	DownRight = move{numberOfRows: 1, numberOfColumns: 1}
	DownLeft  = move{numberOfRows: 1, numberOfColumns: -1}
	ordinal   = []Move{UpRight, UpLeft, DownRight, DownLeft}

	AllDirections = append(cardinal, ordinal...)
)
