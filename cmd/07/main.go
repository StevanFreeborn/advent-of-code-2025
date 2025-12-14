package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/queue"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/stack"
)

const StartCharacter = "S"
const SplitterCharacter = "^"

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	grid := grid.From(input)

	var startPosition position.Position

	for p, v := range grid.Positions() {
		if v == StartCharacter {
			startPosition = p
			break
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

func SolvePartTwo(filePath string) int {
	input := file.ReadAllLines(filePath)
	grid := grid.From(input)

	var startPosition position.Position

	for p, v := range grid.Positions() {
		if v == StartCharacter {
			startPosition = p
			break
		}
	}

	if startPosition == nil {
		panic("unable to find start location")
	}

	beamStartPosition := startPosition.Move(move.Down)
	beamsStack := stack.New[position.Position]()
	beamsStack.Push(beamStartPosition)

	positionsToTimelinesMap := map[position.Position]int{}

	for beamsStack.IsEmpty() == false {
		currentBeam, _ := beamsStack.Peek()

		_, hasBeenProcessed := positionsToTimelinesMap[currentBeam]

		if hasBeenProcessed {
			beamsStack.Pop()
			continue
		}

		value := grid.GetValueAt(currentBeam)
		nextTimelinePositions := []position.Position{}

		if value == SplitterCharacter {
			rightBeam := currentBeam.Move(move.Right)
			leftBeam := currentBeam.Move(move.Left)
			nextTimelinePositions = append(nextTimelinePositions, rightBeam, leftBeam)
		} else {
			nextBeam := currentBeam.Move(move.Down)
			nextTimelinePositions = append(nextTimelinePositions, nextBeam)
		}

		numberOfTimelinesForPosition := 0
		finalizePosition := true

		for _, newTimelinePosition := range nextTimelinePositions {
			if grid.InBounds(newTimelinePosition) == false {
				numberOfTimelinesForPosition++
				continue
			}

			v, existing := positionsToTimelinesMap[newTimelinePosition]

			if existing {
				numberOfTimelinesForPosition += v
				continue
			}

			beamsStack.Push(newTimelinePosition)
			finalizePosition = false
		}

		if finalizePosition == false {
			continue
		}

		beamsStack.Pop()
		positionsToTimelinesMap[currentBeam] = numberOfTimelinesForPosition
	}

	return positionsToTimelinesMap[beamStartPosition]
}
