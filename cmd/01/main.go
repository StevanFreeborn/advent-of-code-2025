package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/dial"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/instruction"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	lines := file.ReadLines(filePath)
	dial := dial.New()
	total := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		instruction := instruction.FromLine(line)
		dial.Turn(instruction)

		if dial.Value() == 0 {
			total++
		}
	}

	return total
}

func SolvePartTwo(filePath string) int {
	lines := file.ReadLines(filePath)
	dial := dial.New()

	for _, line := range lines {
		if line == "" {
			continue
		}

		instruction := instruction.FromLine(line)
		dial.Turn(instruction)
	}

	return dial.ZeroCount()
}
