package position_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

func TestFrom(t *testing.T) {
	row := 3
	column := 5

	pos := position.From(row, column)

	if pos.Row() != row {
		t.Errorf("expected row to be %d, got %d", row, pos.Row())
	}

	if pos.Column() != column {
		t.Errorf("expected column to be %d, got %d", column, pos.Column())
	}
}
