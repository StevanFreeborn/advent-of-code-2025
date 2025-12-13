package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

const StartCharacter = "S"
const SplitterCharacter = "^"

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	numOfRows := len(input)
	numOfCols := len(input[0])

	startRow := -1
	startCol := -1

	for row := range numOfRows - 1 {
		for col := range numOfCols - 1 {
			v := string(input[row][col])

			if v == StartCharacter {
				startRow = row
				startCol = col
				break
			}
		}
	}

	if startRow < 0 || startCol < 0 {
		panic("unable to find start location")
	}

	beamStart := position.From(startRow+1, startCol)
	beamsQueue := []position.Position{
		beamStart,
	}
	visited := map[position.Position]bool{}

	total := 0

	for len(beamsQueue) > 0 {
		currentBeam := beamsQueue[0]
		beamsQueue = beamsQueue[1:]

		if visited[currentBeam] {
			continue
		}

		if currentBeam.Row() < 0 || currentBeam.Row() >= numOfRows {
			continue
		}

		if currentBeam.Column() < 0 || currentBeam.Column() >= numOfCols {
			continue
		}

		visited[currentBeam] = true

		value := string(input[currentBeam.Row()][currentBeam.Column()])

		if value == SplitterCharacter {
			total++
			rightBeam := position.From(currentBeam.Row(), currentBeam.Column()+1)
			leftBeam := position.From(currentBeam.Row(), currentBeam.Column()-1)
			beamsQueue = append(beamsQueue, rightBeam, leftBeam)
			continue
		}

		nextBeam := position.From(currentBeam.Row()+1, currentBeam.Column())
		beamsQueue = append(beamsQueue, nextBeam)
	}

	return total
}
