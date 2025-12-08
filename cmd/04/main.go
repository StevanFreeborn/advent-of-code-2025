package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

const PaperRollCharacter = "@"
const EmptySpaceCharacter = "."

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

func SolvePartTwo(filePath string) int {
	input := file.ReadAllLines(filePath)
	g := grid.From(input)

	total := 0

	for {
		positionsToRemove := []position.Position{}

		for position := range g.Positions() {
			value := g.GetValueAt(position)

			if value != PaperRollCharacter {
				continue
			}

			sameNeighbors := g.GetSameNeighborsOf(position, move.AllDirections)

			if len(sameNeighbors) < 4 {
				positionsToRemove = append(positionsToRemove, position)
				total++
			}
		}

		if len(positionsToRemove) == 0 {
			break
		}

		g.SetValuesAt(positionsToRemove, EmptySpaceCharacter)
	}

	return total
}
