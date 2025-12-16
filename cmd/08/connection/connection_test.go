package connection_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/connection"
)

func TestFrom(t *testing.T) {
	start := box.From("1,2,3")
	end := box.From("4,5,6")
	expectedDistance := start.DistanceFrom(end)

	conn := connection.From(start, end)

	if conn.Start() != start {
		t.Errorf("Connection Start() = %v; want %v", conn.Start(), start)
	}

	if conn.End() != end {
		t.Errorf("Connection End() = %v; want %v", conn.End(), end)
	}

	if conn.Distance() != expectedDistance {
		t.Errorf("Connection Distance() = %v; want %v", conn.Distance(), expectedDistance)
	}
}
