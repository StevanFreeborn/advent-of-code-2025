package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/07"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 21

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolvePartOne() = %d; want %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 1581

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartOne() = %d; want %d", result, expected)
	}
}
