package direction_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/direction"
)

func TestDirectionConstants(t *testing.T) {
	if direction.Left != "L" {
		t.Errorf("Expected Left to be 'L', got '%s'", direction.Left)
	}

	if direction.Right != "R" {
		t.Errorf("Expected Right to be 'R', got '%s'", direction.Right)
	}
}
