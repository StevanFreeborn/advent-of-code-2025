package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/09"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 50

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 4_758_121_828

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 24

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := -1

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
