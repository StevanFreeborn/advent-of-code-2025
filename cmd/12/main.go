package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/shape"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const COLON = ":"
const X = "x"

func SolvePartOne(filePath string) int {
	twoNewLineRegex := regexp.MustCompile(`\r?\n\r?\n`)
	newLineRegex := regexp.MustCompile(`\r?\n`)

	input := file.ReadAllText(filePath)
	sections := twoNewLineRegex.Split(strings.TrimSpace(input), -1)

	total := 0
	shapes := map[int][]shape.Shape{}
	regions := []string{}

	for _, section := range sections {
		lines := newLineRegex.Split(strings.TrimSpace(section), -1)
		header := lines[0]

		if strings.Contains(header, COLON) && strings.Contains(header, X) == false {
			shape := shape.From(lines)
			shapes[shape.Id()] = shape.GenerateVariants()
			continue
		}

		for _, line := range lines {
			if strings.Contains(line, COLON) {
				regions = append(regions, line)
			}
		}
	}

	for _, region := range regions {
		parts := strings.Split(region, COLON)
		dims := strings.Split(parts[0], X)
		width, _ := strconv.Atoi(dims[0])
		height, _ := strconv.Atoi(dims[1])

		countsStr := strings.Fields(parts[1])
		requiredShapes := []int{}

		for id, s := range countsStr {
			count, _ := strconv.Atoi(s)

			for range count {
				requiredShapes = append(requiredShapes, id)
			}
		}

		if canFit(width, height, requiredShapes, shapes) {
			total++
		}
	}

	return total
}

type item struct {
	id       int
	area     int
	variants []shape.Shape
}

func canFit(width int, height int, requiredShapes []int, shapes map[int][]shape.Shape) bool {
	totalArea := 0

	itemsToPlace := make([]item, 0, len(requiredShapes))

	for _, id := range requiredShapes {
		shapeVariants := shapes[id]
		area := shapeVariants[0].Area()
		totalArea += area
		itemsToPlace = append(itemsToPlace, item{id: id, area: area, variants: shapeVariants})
	}

	if totalArea > width*height {
		return false
	}

	sort.Slice(itemsToPlace, func(i, j int) bool {
		return itemsToPlace[i].area > itemsToPlace[j].area
	})

	grid := make([][]bool, height)

	for i := range grid {
		grid[i] = make([]bool, width)
	}

	return checkFit(0, itemsToPlace, grid, width, height)
}

func checkFit(index int, itemsToPlace []item, grid [][]bool, width int, height int) bool {
	if index == len(itemsToPlace) {
		return true
	}

	itemToPlace := itemsToPlace[index]

	for _, variant := range itemToPlace.variants {
		maxRow := variant.MaxRow()
		maxColumn := variant.MaxColumn()

		lastRow := height - maxRow
		lastColumn := width - maxColumn

		for r := range lastRow {
			for c := range lastColumn {
				if canPlace(grid, r, c, variant) {
					place(grid, r, c, variant)

					if checkFit(index+1, itemsToPlace, grid, width, height) {
						return true
					}

					unplace(grid, r, c, variant)
				}
			}
		}
	}

	return false
}

func canPlace(grid [][]bool, r, c int, variant shape.Shape) bool {
	for _, position := range variant.Positions() {
		absoluteRow := r + position.Row()
		absoluteColumn := c + position.Column()
		isOccupied := grid[absoluteRow][absoluteColumn]

		if isOccupied {
			return false
		}
	}

	return true
}

func place(grid [][]bool, r, c int, variant shape.Shape) {
	for _, position := range variant.Positions() {
		absoluteRow := r + position.Row()
		absoluteColumn := c + position.Column()
		grid[absoluteRow][absoluteColumn] = true
	}
}

func unplace(grid [][]bool, r, c int, variant shape.Shape) {
	for _, position := range variant.Positions() {
		absoluteRow := r + position.Row()
		absoluteColumn := c + position.Column()
		grid[absoluteRow][absoluteColumn] = false
	}
}
