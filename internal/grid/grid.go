package grid

import (
	"fmt"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

type Grid interface {
	NumberOfRows() int
	NumberOfColumns() int
	InBounds(position.Position) bool
	GetValueAt(position.Position) string
	GetSameNeighborsOf(position.Position, []move.Move) []position.Position
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
	if pos.Row() < 0 || pos.Row() >= g.numberOfRows {
		return false
	}

	if pos.Column() < 0 || pos.Column() >= g.numberOfColumns {
		return false
	}

	return true
}

func (g grid) GetValueAt(pos position.Position) string {
	return g.positions[pos]
}

func (g grid) GetSameNeighborsOf(pos position.Position, moves []move.Move) []position.Position {
	similarNeighbors := []position.Position{}
	originalValue := g.GetValueAt(pos)

	fmt.Println("Moves to check:", moves)
	for _, m := range moves {
		postionRow := pos.Row()
		positionColumn := pos.Column()

		fmt.Println("Checking move:", m, "from position:", pos)
		fmt.Println("Move rows:", m.NumberOfRows(), "columns:", m.NumberOfColumns())
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
