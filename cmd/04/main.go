package main

import (
	"fmt"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

const PaperRollCharacter = "@"

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	grid := grid.From(input)

	total := 0

	for row := range grid.NumberOfRows() {
		for column := range grid.NumberOfColumns() {
			currentPosition := position.From(row, column)
			value := grid.GetValueAt(currentPosition)

			if value != PaperRollCharacter {
				continue
			}

			sameNeighbors := grid.GetSameNeighborsOf(currentPosition, move.AllDirections)

			if row == 0 && column == 3 {
				fmt.Println(sameNeighbors)
			}

			if len(sameNeighbors) < 4 {
				total++
			}
		}
	}

	return total
}
