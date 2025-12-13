package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/queue"
)

const StartCharacter = "S"
const SplitterCharacter = "^"

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	grid := grid.From(input)

	var startPosition position.Position

	for row := range grid.NumberOfRows() - 1 {
		for col := range grid.NumberOfColumns() - 1 {
			current := position.From(row, col)
			v := grid.GetValueAt(current)

			if v == StartCharacter {
				startPosition = current
				break
			}
		}
	}

	if startPosition == nil {
		panic("unable to find start location")
	}

	beamStart := startPosition.Move(move.Down)
	beamsQueue := queue.New[position.Position]()
	beamsQueue.Enqueue(beamStart)

	visited := map[position.Position]bool{}

	total := 0

	for beamsQueue.IsEmpty() == false {
		currentBeam, _ := beamsQueue.Dequeue()

		if visited[currentBeam] {
			continue
		}

		if grid.InBounds(currentBeam) == false {
			continue
		}

		visited[currentBeam] = true

		value := grid.GetValueAt(currentBeam)

		if value == SplitterCharacter {
			total++

			rightBeam := currentBeam.Move(move.Right)
			leftBeam := currentBeam.Move(move.Left)

			beamsQueue.Enqueue(rightBeam)
			beamsQueue.Enqueue(leftBeam)
			continue
		}

		nextBeam := currentBeam.Move(move.Down)
		beamsQueue.Enqueue(nextBeam)
	}

	return total
}
