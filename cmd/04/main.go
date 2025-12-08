package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
)

const PaperRollCharacter = "@"

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	grid := grid.From(input)

	total := 0

	for position := range grid.Positions() {
		value := grid.GetValueAt(position)

		if value != PaperRollCharacter {
			continue
		}

		sameNeighbors := grid.GetSameNeighborsOf(position, move.AllDirections)

		if len(sameNeighbors) < 4 {
			total++
		}
	}
	return total
}
