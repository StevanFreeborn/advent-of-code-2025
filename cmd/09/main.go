package main

import (
	"math"
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

// TODO: Refactor below so more understandable

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

			rectMinRow := int(math.Min(float64(start.Row()), float64(end.Row())))
			rectMaxRow := int(math.Max(float64(start.Row()), float64(end.Row())))
			rectMinCol := int(math.Min(float64(start.Column()), float64(end.Column())))
			rectMaxCol := int(math.Max(float64(start.Column()), float64(end.Column())))

			width := int(math.Abs(float64(rectMaxCol-rectMinCol)) + 1)
			height := int(math.Abs(float64(rectMaxRow-rectMinRow)) + 1)

			a := width * height

			if a <= largestArea {
				continue
			}

			centerRow := int((rectMaxRow + rectMinRow) / 2)
			centerCol := int((rectMaxCol + rectMinCol) / 2)
			center := position.From(centerRow, centerCol)

			centerIsInside := false

			for i := range numberOfPositions {
				edgeStart := positions[i]
				edgeEndIndex := (i + 1) % numberOfPositions
				edgeEnd := positions[edgeEndIndex]

				isCenterVerticallyBetweenEdge := (edgeStart.Row() <= center.Row()) && (center.Row() <= edgeEnd.Row())

				edgeTotalHeight := int(math.Abs(float64(edgeEnd.Row() - edgeStart.Row())))
				edgeTotalWidth := int(math.Abs(float64(edgeEnd.Column() - edgeStart.Column())))
				centerPercentageOfEdgeHeight := int(float64(center.Row()-edgeStart.Row()) / float64(edgeTotalHeight))
				edgeColumnAtCenterRow := edgeStart.Column() + centerPercentageOfEdgeHeight*edgeTotalWidth
				isCenterHorizontallyToTheLeft := center.Row() <= edgeColumnAtCenterRow

				if isCenterVerticallyBetweenEdge && isCenterHorizontallyToTheLeft {
					centerIsInside = !centerIsInside
				}
			}

			if centerIsInside == false {
				continue
			}

			isIntersectedByEdge := false

			for i := range numberOfPositions {
				currPos := positions[i]
				nextPosIndex := (i + 1) % numberOfPositions
				nextPos := positions[nextPosIndex]

				if currPos.Column() == nextPos.Column() {
					col := currPos.Column()

					if col > rectMinCol && col < rectMaxCol {
						edgeMinRow := math.Min(float64(currPos.Row()), float64(nextPos.Row()))
						edgeMaxRow := math.Max(float64(currPos.Row()), float64(nextPos.Row()))

						overlapStart := int(math.Max(edgeMinRow, float64(rectMinRow)))
						overlapEnd := int(math.Min(edgeMaxRow, float64(rectMaxRow)))

						if overlapEnd > overlapStart {
							isIntersectedByEdge = true
							break
						}
					}
				} else {
					row := currPos.Row()

					if row > rectMinRow && row < rectMaxRow {
						edgeMinCol := math.Min(float64(currPos.Column()), float64(nextPos.Column()))
						edgeMaxCol := math.Max(float64(currPos.Column()), float64(nextPos.Column()))

						overlapStart := int(math.Max(edgeMinCol, float64(rectMinCol)))
						overlapEnd := int(math.Min(edgeMaxCol, float64(rectMaxCol)))

						if overlapEnd > overlapStart {
							isIntersectedByEdge = true
							break
						}
					}
				}
			}

			if isIntersectedByEdge {
				continue
			}

			largestArea = a
		}
	}

	return largestArea
}
