package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/01"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := -1

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolvePartOne() = %d; want %d", result, expected)
	}
}
