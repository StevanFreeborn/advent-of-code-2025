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

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := int64(14)

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("SolvePartTwo with example input: expected %d, got %d", expected, result)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := int64(342_018_167_474_526)

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("SolvePartTwo with input: expected %d, got %d", expected, result)
	}
}
