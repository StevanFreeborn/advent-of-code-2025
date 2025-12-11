package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/06"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 4277556

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 6169101504608

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
