package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/03/bank"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	input := file.ReadLines(filePath)
	total := 0

	for _, line := range input {
		bank := bank.From(line)
		total += bank.Joltage()
	}

	return total
}
