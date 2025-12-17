package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/08"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 40

	result := solution.SolvePartOne("EXAMPLE.txt", 10)

	if result != expected {
		t.Errorf("SolvePartOne returned %d, expected %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 62_186

	result := solution.SolvePartOne("INPUT.txt", 1000)

	if result != expected {
		t.Errorf("SolvePartOne returned %d, expected %d", result, expected)
	}
}

func TestSolvePartOneAgainWithExampleInput(t *testing.T) {
	expected := 40

	result := solution.SolvePartOneAgain("EXAMPLE.txt", 10)

	if result != expected {
		t.Errorf("SolvePartOneAgain returned %d, expected %d", result, expected)
	}
}

func TestSolvePartOneAgainWithInput(t *testing.T) {
	expected := 62_186

	result := solution.SolvePartOneAgain("INPUT.txt", 1000)

	if result != expected {
		t.Errorf("SolvePartOneAgain returned %d, expected %d", result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 25_272

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolvePartTwo returned %d, expected %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := 8_420_405_530

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartTwo returned %d, expected %d", result, expected)
	}
}
