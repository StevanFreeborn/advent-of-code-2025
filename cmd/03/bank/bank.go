// Package bank represents a bank of batteries
package bank

import (
	"strconv"
	"strings"
)

type Bank interface {
	Joltage(numOfCellsToTurnOn int) int64
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

func (b bank) Joltage(numOfCellsToTurnOn int) int64 {
	bankLength := len(b.cells)
	cellsOn := []int{}
	removalBudget := bankLength - numOfCellsToTurnOn

	for _, cell := range b.cells {
		for len(cellsOn) > 0 && removalBudget > 0 && cell > cellsOn[len(cellsOn)-1] {
			cellToRemoveIndex := len(cellsOn) - 1
			cellsOn = cellsOn[:cellToRemoveIndex]
			removalBudget--
		}

		cellsOn = append(cellsOn, cell)
	}

	var sb strings.Builder

	for _, n := range cellsOn[:numOfCellsToTurnOn] {
		sb.WriteString(strconv.Itoa(n))
	}

	finalString := sb.String()

	finalJoltage, _ := strconv.ParseInt(finalString, 10, 64)

	return finalJoltage
}

func (b bank) Cells() []int {
	return b.cells
}
