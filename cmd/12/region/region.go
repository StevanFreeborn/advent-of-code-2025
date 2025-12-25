// Package region provides a model and methods for manipulating a region
package region

import (
	"sort"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/shape"
)

type Region interface {
	CanFit(shapes map[int][]shape.Shape) bool
}

type region struct {
	width          int
	height         int
	requiredShapes []int
}

func From(line string) Region {
	parts := strings.Split(line, ":")
	dims := strings.Split(parts[0], "x")
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

	return region{
		width:          width,
		height:         height,
		requiredShapes: requiredShapes,
	}
}

func (r region) CanFit(shapes map[int][]shape.Shape) bool {
	totalArea := 0

	itemsToPlace := r.createItems(shapes)

	for _, item := range itemsToPlace {
		totalArea += item.area
	}

	if totalArea > r.area() {
		return false
	}

	return r.checkFit(0, itemsToPlace, r.createGrid())
}

type item struct {
	id       int
	area     int
	variants []shape.Shape
}

func (r region) checkFit(index int, itemsToPlace []item, grid [][]bool) bool {
	if index == len(itemsToPlace) {
		return true
	}

	currentItem := itemsToPlace[index]

	for _, variant := range currentItem.variants {
		maxRow := variant.MaxRow()
		maxColumn := variant.MaxColumn()

		lastRow := r.height - maxRow
		lastColumn := r.width - maxColumn

		for rowNumber := range lastRow {
			for columnNumber := range lastColumn {
				if canPlace(grid, rowNumber, columnNumber, variant) {
					place(grid, rowNumber, columnNumber, variant)

					if r.checkFit(index+1, itemsToPlace, grid) {
						return true
					}

					unplace(grid, rowNumber, columnNumber, variant)
				}
			}
		}
	}

	return false
}

func (r region) area() int {
	return r.width * r.height
}

func (r region) createGrid() [][]bool {
	grid := make([][]bool, r.height)

	for i := range grid {
		grid[i] = make([]bool, r.width)
	}

	return grid
}

func (r region) createItems(shapes map[int][]shape.Shape) []item {
	itemsToPlace := make([]item, 0, len(r.requiredShapes))

	for _, id := range r.requiredShapes {
		shapeVariants := shapes[id]
		area := shapeVariants[0].Area()
		itemsToPlace = append(itemsToPlace, item{id: id, area: area, variants: shapeVariants})
	}

	sort.Slice(itemsToPlace, func(i, j int) bool {
		return itemsToPlace[i].area > itemsToPlace[j].area
	})

	return itemsToPlace
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
