package move_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
)

func TestNumberOfRows(t *testing.T) {
	tests := []struct {
		move     move.Move
		expected int
	}{
		{move.Up, -1},
		{move.Down, 1},
		{move.Right, 0},
		{move.Left, 0},
		{move.UpRight, -1},
		{move.UpLeft, -1},
		{move.DownRight, 1},
		{move.DownLeft, 1},
	}

	for _, test := range tests {
		if test.move.NumberOfRows() != test.expected {
			t.Errorf("expected number of rows for move to be %d, got %d", test.expected, test.move.NumberOfRows())
		}
	}
}

func TestNumberOfColumns(t *testing.T) {
	tests := []struct {
		move     move.Move
		expected int
	}{
		{move.Up, 0},
		{move.Down, 0},
		{move.Right, 1},
		{move.Left, -1},
		{move.UpRight, 1},
		{move.UpLeft, -1},
		{move.DownRight, 1},
		{move.DownLeft, -1},
	}

	for _, test := range tests {
		if test.move.NumberOfColumns() != test.expected {
			t.Errorf("expected number of columns for move to be %d, got %d", test.expected, test.move.NumberOfColumns())
		}
	}
}

func TestAllDirectionsLength(t *testing.T) {
	expectedLength := 8

	if len(move.AllDirections) != expectedLength {
		t.Errorf("expected AllDirections length to be %d, got %d", expectedLength, len(move.AllDirections))
	}
}

func TestDefinedMoves(t *testing.T) {
	definedMoves := []struct {
		name            string
		move            move.Move
		expectedRows    int
		expectedColumns int
	}{
		{"Up", move.Up, -1, 0},
		{"Down", move.Down, 1, 0},
		{"Right", move.Right, 0, 1},
		{"Left", move.Left, 0, -1},
		{"UpRight", move.UpRight, -1, 1},
		{"UpLeft", move.UpLeft, -1, -1},
		{"DownRight", move.DownRight, 1, 1},
		{"DownLeft", move.DownLeft, 1, -1},
	}

	for _, dm := range definedMoves {
		if dm.move == nil {
			t.Errorf("expected move %s to be defined, got nil", dm.name)
		}

		if dm.move.NumberOfRows() != dm.expectedRows {
			t.Errorf("expected move %s to have %d rows, got %d", dm.name, dm.expectedRows, dm.move.NumberOfRows())
		}

		if dm.move.NumberOfColumns() != dm.expectedColumns {
			t.Errorf("expected move %s to have %d columns, got %d", dm.name, dm.expectedColumns, dm.move.NumberOfColumns())
		}
	}
}
