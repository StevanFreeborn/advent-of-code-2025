package main

import (
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/09/edge"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/09/rectangle"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

func SolvePartOne(filePath string) int {
	positions := []position.Position{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := position.From(y, x)
		positions = append(positions, p)
	}

	numberOfPositions := len(positions)
	largestArea := 0

	for i := range numberOfPositions {
		for j := i + 1; j < numberOfPositions; j++ {
			start := positions[i]
			end := positions[j]
			rect := rectangle.From(start, end)

			if rect.Area() > largestArea {
				largestArea = rect.Area()
			}
		}
	}

	return largestArea
}

func SolvePartTwo(filePath string) int {
	positions := []position.Position{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := position.From(y, x)
		positions = append(positions, p)
	}

	numberOfPositions := len(positions)
	largestArea := 0

	for i := range numberOfPositions {
		for j := i + 1; j < numberOfPositions; j++ {
			start := positions[i]
			end := positions[j]
			rect := rectangle.From(start, end)

			if rect.Area() <= largestArea {
				continue
			}

			centerIsInside := false

			for i := range numberOfPositions {
				edgeStart := positions[i]
				edgeEndIndex := (i + 1) % numberOfPositions
				edgeEnd := positions[edgeEndIndex]
				edge := edge.From(edgeStart, edgeEnd)

				if edge.VerticallyContains(rect.Center()) && edge.IntersectsHorizontalLineAt(rect.Center()) {
					centerIsInside = !centerIsInside
				}
			}

			if centerIsInside == false {
				continue
			}

			isIntersectedByEdge := false

			for i := range numberOfPositions {
				edgeStart := positions[i]
				edgeEndIndex := (i + 1) % numberOfPositions
				edgeEnd := positions[edgeEndIndex]
				edge := edge.From(edgeStart, edgeEnd)

				if edge.Intersects(rect) {
					isIntersectedByEdge = true
					break
				}
			}

			if isIntersectedByEdge {
				continue
			}

			largestArea = rect.Area()
		}
	}

	return largestArea
}
