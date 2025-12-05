// Package bank represents a bank of batteries
package bank

import (
	"fmt"
	"strconv"
)

type Bank interface {
	Joltage() int
	Cells() []int
}

type bank struct {
	cells []int
}

func From(line string) Bank {
	cells := []int{}

	for _, char := range line {
		num, _ := strconv.Atoi(string(char))
		cells = append(cells, num)
	}

	return bank{
		cells: cells,
	}
}

type cell struct {
	index int
	value int
}

func (b bank) Joltage() int {
	// TODO: We still need to
	// find the max
	// however we should then
	// find the next max to the
	// left of the highest
	// and find the next max to
	// the right of the higest
	// we should then see what
	// the large number we can make
	// with these values

	max := cell{
		value: -1,
	}

	for ci, cv := range b.cells {
		if cv > max.value {
			max = cell{
				index: ci,
				value: cv,
			}
		}
	}

	maxLeft := cell{
		value: -1,
	}

	for ci, cv := range b.cells[:max.index] {
		if cv > maxLeft.value {
			maxLeft = cell{
				index: ci,
				value: cv,
			}
		}
	}

	maxRight := cell{
		value: -1,
	}

	for ci, cv := range b.cells[max.index+1:] {
		if cv > maxRight.value {
			maxRight = cell{
				index: ci,
				value: cv,
			}
		}
	}

	maxInTensPlace := fmt.Sprintf("%d%d", max.value, maxRight.value)
	maxInOnesPlace := fmt.Sprintf("%d%d", maxLeft.value, max.value)

	maxInTensPlaceAsNumber, _ := strconv.Atoi(maxInTensPlace)
	maxInOnesPlaceAsNumber, _ := strconv.Atoi(maxInOnesPlace)

	if maxInTensPlaceAsNumber > maxInOnesPlaceAsNumber {
		return maxInTensPlaceAsNumber
	}

	return maxInOnesPlaceAsNumber
}

func (b bank) Cells() []int {
	return b.cells
}
