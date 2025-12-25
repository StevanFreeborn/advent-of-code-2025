package main

import (
	"fmt"

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

	linesRead := 0
	for line := range file.ReadLines(filePath) {
		linesRead++
		machine := machine.From(line)
		buttonsPressed := machine.ConfigureJoltages()
		if buttonsPressed == 0 {
			fmt.Printf("No solution found for line %d: %s\n", linesRead, line)
		}
		total += buttonsPressed
	}

	return total
}
