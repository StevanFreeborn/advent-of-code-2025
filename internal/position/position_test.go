package position_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
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

func TestMove(t *testing.T) {
	expectedPos := position.From(3, 4)
	startPos := position.From(2, 4)

	result := startPos.Move(move.Down)

	if result != expectedPos {
		t.Errorf("expected position to be %+v, got %+v", expectedPos, result)
	}
}
