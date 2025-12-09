package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/05"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 3

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolvePartOne with example input: expected %d, got %d", expected, result)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 643

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartOne with input: expected %d, got %d", expected, result)
	}
}
