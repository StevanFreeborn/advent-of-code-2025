package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

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

			width := int(math.Abs(float64(end.Column()-start.Column())) + 1)
			height := int(math.Abs(float64(end.Row()-start.Row())) + 1)

			a := width * height

			if a > largestArea {
				largestArea = a
			}
		}
	}

	return largestArea
}

// TODO: How would we construct all the positions
// in the polygon?

func SolvePartTwo(filePath string) int {
	positions := []position.Position{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := position.From(y, x)
		positions = append(positions, p)
	}

	rowPositions := []int{}
	columnPositions := []int{}

	for _, positions := range positions {
		rowPositions = append(rowPositions, positions.Row())
		columnPositions = append(columnPositions, positions.Column())
	}

	minRow := slices.Min(rowPositions)
	maxRow := slices.Max(rowPositions)
	minCol := slices.Min(columnPositions)
	maxCol := slices.Max(columnPositions)

	numberOfPositions := len(positions)
	largestArea := 0

	for i := range numberOfPositions {
		for j := i + 1; j < numberOfPositions; j++ {
			start := positions[i]
			end := positions[j]

			width := int(math.Abs(float64(end.Column()-start.Column())) + 1)
			height := int(math.Abs(float64(end.Row()-start.Row())) + 1)

			a := width * height

			if a <= largestArea {
				continue
			}

			rectMinRow := math.Min(float64(start.Row()), float64(end.Row()))
			rectMaxRow := math.Max(float64(start.Row()), float64(end.Row()))
			rectMinCol := math.Min(float64(start.Column()), float64(end.Column()))
			rectMaxCol := math.Max(float64(start.Column()), float64(end.Column()))

			// is center of rect inside our polygon?
			centerRow := int((rectMaxRow + rectMinRow) / 2)
			centerCol := int((rectMaxCol + rectMinCol) / 2)
			center := position.From(centerRow, centerCol)

			centerIsInside := false
			j := numberOfPositions - 1

			for i := 0; i < numberOfPositions; i++ {
				edgeStart := positions[i]
				edgeEnd := positions[j]

				edgeCoversCenterRow := (edgeStart.Row() > center.Row()) != (edgeEnd.Row() > center.Row())

				edgeHorizontalDistance := edgeEnd.Column() - edgeStart.Column()
				edgeVeritcalDistance := edgeEnd.Row() - edgeStart.Row()
				centerVerticalOffset := center.Row() - edgeStart.Row()
				verticalDistanceCompleted := centerVerticalOffset / edgeVeritcalDistance
				xIntersection := center.Column() + edgeHorizontalDistance*verticalDistanceCompleted
				edgeIsToRightOfCenter := center.Column() < xIntersection

				if edgeCoversCenterRow && edgeIsToRightOfCenter {
					centerIsInside = !centerIsInside
				}

				j = i
			}

			if centerIsInside == false {
				continue
			}

			// TODO: We need to figure this out
			// if edges of polygon intersect with
			// edges of rectangle

			largestArea = a
		}
	}

	return largestArea
}
