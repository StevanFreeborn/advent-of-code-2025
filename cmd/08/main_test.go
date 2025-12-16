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
	expected := 62186

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
	expected := 62186

	result := solution.SolvePartOneAgain("INPUT.txt", 1000)

	if result != expected {
		t.Errorf("SolvePartOneAgain returned %d, expected %d", result, expected)
	}
}
