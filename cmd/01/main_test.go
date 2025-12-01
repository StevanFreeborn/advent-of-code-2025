package main_test

import (
	"testing"

	solution "github.com/StevanFreeborn/advent-of-code-2025/cmd/01"
)

func TestSolvePartOneWithExampleInput(t *testing.T) {
	expected := 3

	result := solution.SolvePartOne("EXAMPLE.txt")

	if result != expected {
		t.Errorf(`SolvePartOne("EXAMPLE.txt") = %d; want %d`, result, expected)
	}
}

func TestSolvePartOneWithInput(t *testing.T) {
	expected := 1011

	result := solution.SolvePartOne("INPUT.txt")

	if result != expected {
		t.Errorf(`SolvePartOne("INPUT.txt") = %d; want %d`, result, expected)
	}
}
