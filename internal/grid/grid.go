package grid

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

type Grid interface {
	NumberOfRows() int
	NumberOfColumns() int
	InBounds(position.Position) bool
	GetValueAt(position.Position) string
	GetSameNeighborsOf(position.Position, []move.Move) []position.Position
	Positions() map[position.Position]string
}

type grid struct {
	numberOfRows    int
	numberOfColumns int
	positions       map[position.Position]string
}

func From(input []string) Grid {
	numberOfRows := len(input)
	numberOfColumns := len(input[0])
	positions := map[position.Position]string{}

	for row := range numberOfRows {
		for column := range numberOfColumns {
			value := string(input[row][column])
			pos := position.From(row, column)
			positions[pos] = value
		}
	}

	return grid{
		numberOfRows:    numberOfRows,
		numberOfColumns: numberOfColumns,
		positions:       positions,
	}
}

func (g grid) NumberOfRows() int {
	return g.numberOfRows
}

func (g grid) NumberOfColumns() int {
	return g.numberOfColumns
}

func (g grid) InBounds(pos position.Position) bool {
	_, ok := g.positions[pos]
	return ok
}

func (g grid) GetValueAt(pos position.Position) string {
	return g.positions[pos]
}

func (g grid) GetSameNeighborsOf(pos position.Position, moves []move.Move) []position.Position {
	similarNeighbors := []position.Position{}
	originalValue := g.GetValueAt(pos)

	for _, m := range moves {
		postionRow := pos.Row()
		positionColumn := pos.Column()

		neighborRow := postionRow + m.NumberOfRows()
		neighborColumn := positionColumn + m.NumberOfColumns()
		neighborPos := position.From(neighborRow, neighborColumn)

		if g.InBounds(neighborPos) == false {
			continue
		}

		neighborValue := g.GetValueAt(neighborPos)

		if neighborValue != originalValue {
			continue
		}

		similarNeighbors = append(similarNeighbors, neighborPos)
	}

	return similarNeighbors
}

func (g grid) Positions() map[position.Position]string {
	positions := map[position.Position]string{}

	for pos, value := range g.positions {
		newPosition := position.From(pos.Row(), pos.Column())
		positions[newPosition] = value
	}

	return positions
}
