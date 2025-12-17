package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	points := []struct {
		x int
		y int
	}{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		p := struct {
			x int
			y int
		}{x: x, y: y}
		points = append(points, p)
	}

	numberOfPoints := len(points)
	largestArea := 0

	for i := range numberOfPoints {
		for j := i + 1; j < numberOfPoints; j++ {
			start := points[i]
			end := points[j]

			width := int(math.Abs(float64(end.x-start.x)) + 1)
			height := int(math.Abs(float64(end.y-start.y)) + 1)

			a := width * height

			if a > largestArea {
				largestArea = a
			}
		}
	}

	return largestArea
}
