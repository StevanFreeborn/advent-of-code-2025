package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/10/machine"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	total := 0

	for line := range file.ReadLines(filePath) {
		machine := machine.From(line)
		buttonsPressed := machine.ConfigureLights()
		total += buttonsPressed
	}

	return total
}

func SolvePartTwo(filePath string) int {
	total := 0

	for line := range file.ReadLines(filePath) {
		machine := machine.From(line)
		buttonsPressed := machine.ConfigureJoltages()
		total += buttonsPressed
	}

	return total
}
