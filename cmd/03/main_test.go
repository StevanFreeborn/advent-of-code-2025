package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/03"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 357

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 17113

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
