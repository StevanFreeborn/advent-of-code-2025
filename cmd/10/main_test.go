package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/10"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 7

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 527

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := 33

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := 19810

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
