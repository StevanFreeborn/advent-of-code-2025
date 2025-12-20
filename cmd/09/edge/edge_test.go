package edge_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/09/edge"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

func TestFrom(t *testing.T) {
	start := position.From(1, 2)
	end := position.From(3, 4)

	e := edge.From(start, end)

	if e.Start() != start {
		t.Errorf("Expected start position %v, got %v", start, e.Start())
	}

	if e.End() != end {
		t.Errorf("Expected end position %v, got %v", end, e.End())
	}
}

func TestVerticallyContains(t *testing.T) {
	start := position.From(1, 2)
	end := position.From(3, 4)
	e := edge.From(start, end)

	tests := []struct {
		p        position.Position
		expected bool
	}{
		{position.From(2, 3), true},
		{position.From(1, 2), true},
		{position.From(3, 4), true},
		{position.From(0, 0), false},
		{position.From(4, 5), false},
	}

	for _, test := range tests {
		result := e.VerticallyContains(test.p)

		if result != test.expected {
			t.Errorf("VerticallyContains(%v) = %v; want %v", test.p, result, test.expected)
		}
	}
}

func TestIntersectsHorizontalLineAt(t *testing.T) {
	start := position.From(1, 2)
	end := position.From(3, 4)
	e := edge.From(start, end)

	tests := []struct {
		p        position.Position
		expected bool
	}{
		{position.From(2, 3), true},
		{position.From(1, 2), true},
		{position.From(3, 4), true},
	}

	for _, test := range tests {
		result := e.IntersectsHorizontalLineAt(test.p)

		if result != test.expected {
			t.Errorf("IntersectsHorizontalLineAt(%v) = %v; want %v", test.p, result, test.expected)
		}
	}
}
