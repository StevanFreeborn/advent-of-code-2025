package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/10"
)

func TestSolvePartOne(t *testing.T) {
	expected := -1

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf("got %d but wanted %d", result, expected)
	}
}
