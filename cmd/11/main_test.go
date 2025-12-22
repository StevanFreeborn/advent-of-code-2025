package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/11"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 5

	result := solution.SolvePartOne("EXAMPLE_ONE.txt")

	if result != expected {
		t.Errorf("SolvePartOne(EXAMPLE_ONE.txt) = %d; want %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 599

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartOne(INPUT.txt) = %d; want %d", result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 2

	result := solution.SolvePartTwo("EXAMPLE_TWO.txt")

	if result != expected {
		t.Errorf("SolvePartTwo(EXAMPLE_TWO.txt) = %d; want %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := 393_474_305_030_400

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartTwo(INPUT.txt) = %d; want %d", result, expected)
	}
}
