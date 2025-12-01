package instruction_test

import (
	"fmt"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/instruction"
)

func TestFromLine(t *testing.T) {
	expectedDirection := "L"
	expectedDistance := 7
	line := fmt.Sprintf("%s%d", expectedDirection, expectedDistance)

	result := instruction.FromLine(line)

	if result.Direction() != expectedDirection {
		t.Errorf("got direction %s but expected direction %s", result.Direction(), expectedDirection)
	}

	if result.Distance() != expectedDistance {
		t.Errorf("got distance %d but expected distance %d", result.Distance(), expectedDistance)
	}
}

func TestFromParts(t *testing.T) {
	expectedDirection := "L"
	expectedDistance := 7

	result := instruction.FromParts(expectedDirection, expectedDistance)

	if result.Direction() != expectedDirection {
		t.Errorf("got direction %s but expected direction %s", result.Direction(), expectedDirection)
	}

	if result.Distance() != expectedDistance {
		t.Errorf("got distance %d but expected distance %d", result.Distance(), expectedDistance)
	}
}
