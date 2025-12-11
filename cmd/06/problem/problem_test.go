package problem_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/06/problem"
)

func TestSolve(t *testing.T) {
	expected := 6
	operator := "+"
	operands := []int{1, 2, 3}

	result := problem.From(operator, operands).Solve()

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
