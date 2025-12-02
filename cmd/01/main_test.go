package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/01"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 3

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf(`SolvePartOne("EXAMPLE.txt") = %d; want %d`, result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 1011

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf(`SolvePartOne("INPUT.txt") = %d; want %d`, result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 6

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf(`SolvePartTwo("EXAMPLE.txt") = %d; want %d`, result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := -1
	tooLow := 5849
	tooHigh := 5967

	result := solution.SolvePartTwo("INPUT.txt")

	if result <= tooLow || result >= tooHigh {
		t.Errorf(`SolvePartTwo("INPUT.txt") = %d; want something between %d and %d`, result, tooLow, tooHigh)
	}

	if result != expected {
		t.Errorf(`SolvePartTwo("INPUT.txt") = %d; want %d`, result, expected)
	}
}
