package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/03"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := int64(357)

	result := solution.Solve("EXAMPLE.txt", 2)

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := int64(17113)

	result := solution.Solve("INPUT.txt", 2)

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithExampleInput(t *testing.T) {
	expected := int64(3_121_910_778_619)

	result := solution.Solve("EXAMPLE.txt", 12)

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}

func TestSolvePartTwoWithInput(t *testing.T) {
	expected := int64(169_709_990_062_889)

	result := solution.Solve("INPUT.txt", 12)

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
