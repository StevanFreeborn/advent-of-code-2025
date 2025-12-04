package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/02"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := int64(1_227_775_554)

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := int64(18_952_700_150)

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := int64(4_174_379_265)

	result := solution.SolvePartTwo("EXAMPLE.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := int64(28858486244)

	result := solution.SolvePartTwo("INPUT.txt")

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
