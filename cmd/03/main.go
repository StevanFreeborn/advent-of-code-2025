package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/03/bank"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func Solve(filePath string, numOfCells int) int64 {
	input := file.ReadAllLines(filePath)
	total := int64(0)

	for _, line := range input {
		bank := bank.From(line)
		total += bank.Joltage(numOfCells)
	}

	return total
}
