package grid_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/grid"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/move"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

func TestFrom(t *testing.T) {
	expectedNumberOfRows := 3
	expectedNumberOfColumns := 3
	expectedPositions := map[position.Position]string{
		position.From(0, 0): "@",
		position.From(0, 1): ".",
		position.From(0, 2): ".",
		position.From(1, 0): ".",
		position.From(1, 1): "@",
		position.From(1, 2): ".",
		position.From(2, 0): ".",
		position.From(2, 1): ".",
		position.From(2, 2): "@",
	}

	input := []string{
		"@..",
		".@.",
		"..@",
	}

	g := grid.From(input)

	if g.NumberOfRows() != expectedNumberOfRows {
		t.Errorf("expected number of rows to be %d, got %d", expectedNumberOfRows, g.NumberOfRows())
	}

	if g.NumberOfColumns() != expectedNumberOfColumns {
		t.Errorf("expected number of columns to be %d, got %d", expectedNumberOfColumns, g.NumberOfColumns())
	}

	for pos, expectedValue := range expectedPositions {
		value := g.GetValueAt(pos)

		if value != expectedValue {
			t.Errorf("expected value at position (%d, %d) to be %s, got %s", pos.Row(), pos.Column(), expectedValue, value)
		}
	}
}

func TestInBounds(t *testing.T) {
	tests := []struct {
		pos      position.Position
		expected bool
	}{
		{position.From(0, 0), true},
		{position.From(1, 1), true},
		{position.From(2, 2), true},
		{position.From(-1, 0), false},
		{position.From(0, -1), false},
		{position.From(3, 0), false},
		{position.From(0, 3), false},
	}

	input := []string{
		"@..",
		".@.",
		"..@",
	}

	g := grid.From(input)

	for _, test := range tests {
		result := g.InBounds(test.pos)

		if result != test.expected {
			t.Errorf("expected InBounds(%v) to be %v, got %v", test.pos, test.expected, result)
		}
	}
}

func TestGetValueAt(t *testing.T) {
	tests := []struct {
		pos      position.Position
		expected string
	}{
		{position.From(0, 0), "@"},
		{position.From(0, 1), "."},
		{position.From(1, 1), "@"},
		{position.From(2, 2), "@"},
	}

	input := []string{
		"@..",
		".@.",
		"..@",
	}

	g := grid.From(input)

	for _, test := range tests {
		result := g.GetValueAt(test.pos)

		if result != test.expected {
			t.Errorf("expected GetValueAt(%v) to be %s, got %s", test.pos, test.expected, result)
		}
	}
}

func TestGetSameNeighborsOf(t *testing.T) {
	input := []string{
		"@..",
		".@.",
		"..@",
	}

	g := grid.From(input)

	tests := []struct {
		pos      position.Position
		expected []position.Position
	}{
		{
			position.From(1, 1),
			[]position.Position{
				position.From(0, 0),
				position.From(2, 2),
			},
		},
		{
			position.From(0, 0),
			[]position.Position{
				position.From(1, 1),
			},
		},
	}

	for _, test := range tests {
		result := g.GetSameNeighborsOf(test.pos, move.AllDirections)

		if len(result) != len(test.expected) {
			t.Errorf("expected GetSameNeighborsOf(%v) to return %d neighbors, got %d", test.pos, len(test.expected), len(result))
			continue
		}

		for i, expectedPos := range test.expected {
			if result[i] != expectedPos {
				t.Errorf("expected neighbor at index %d to be %v, got %v", i, expectedPos, result[i])
			}
		}
	}
}
