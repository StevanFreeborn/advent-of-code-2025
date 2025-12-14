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

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 40

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolveTwoOne() = %d; want %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := 73_007_003_089_792

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("SolveTwoOne() = %d; want %d", result, expected)
	}
}
