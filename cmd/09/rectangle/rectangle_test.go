package rectangle_test

import (
	"slices"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/09/rectangle"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

func TestFrom(t *testing.T) {
	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	expectedCorners := []position.Position{
		cornerOne,
		cornerTwo,
		position.From(1, 4),
		position.From(3, 2),
	}

	result := rectangle.From(cornerOne, cornerTwo)

	if slices.Equal(result.Corners(), expectedCorners) == false {
		t.Errorf("Expected corners %v, but got %v", expectedCorners, result.Corners())
	}
}

func TestMaxRow(t *testing.T) {
	expectedMaxRow := 3

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).MaxRow()

	if result != expectedMaxRow {
		t.Errorf("Expected max row %d, but got %d", expectedMaxRow, result)
	}
}

func TestMinRow(t *testing.T) {
	expectedMinRow := 1

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).MinRow()

	if result != expectedMinRow {
		t.Errorf("Expected min row %d, but got %d", expectedMinRow, result)
	}
}

func TestMaxColumn(t *testing.T) {
	expectedMaxColumn := 4

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).MaxColumn()

	if result != expectedMaxColumn {
		t.Errorf("Expected max column %d, but got %d", expectedMaxColumn, result)
	}
}

func TestMinColumn(t *testing.T) {
	expectedMinColumn := 2

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).MinColumn()

	if result != expectedMinColumn {
		t.Errorf("Expected min column %d, but got %d", expectedMinColumn, result)
	}
}

func TestArea(t *testing.T) {
	expectedArea := 9

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).Area()

	if result != expectedArea {
		t.Errorf("Expected area %d, but got %d", expectedArea, result)
	}
}

func TestCenter(t *testing.T) {
	expectedCenter := position.From(2, 3)

	cornerOne := position.From(1, 2)
	cornerTwo := position.From(3, 4)

	result := rectangle.From(cornerOne, cornerTwo).Center()

	if result != expectedCenter {
		t.Errorf("Expected center %v, but got %v", expectedCenter, result)
	}
}
