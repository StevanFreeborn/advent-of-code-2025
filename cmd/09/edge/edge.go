// Package edge provides a model and methods for representing and manipulating edges.
package edge

import (
	"math"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/09/rectangle"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/position"
)

type Edge interface {
	Start() position.Position
	End() position.Position
	VerticallyContains(position.Position) bool
	IntersectsHorizontalLineAt(position.Position) bool
	Intersects(rect rectangle.Rectangle) bool
}

type edge struct {
	start       position.Position
	end         position.Position
	totalHeight int
	totalWidth  int
	minRow      float64
	maxRow      float64
	minColumn   float64
	maxColumn   float64
}

func From(start position.Position, end position.Position) Edge {
	totalHeight := int(math.Abs(float64(end.Row() - start.Row())))
	totalWidth := int(math.Abs(float64(end.Column() - start.Column())))

	minRow := math.Min(float64(start.Row()), float64(end.Row()))
	maxRow := math.Max(float64(start.Row()), float64(end.Row()))
	minColumn := math.Min(float64(start.Column()), float64(end.Column()))
	maxColumn := math.Max(float64(start.Column()), float64(end.Column()))

	return edge{
		start:       start,
		end:         end,
		totalHeight: totalHeight,
		totalWidth:  totalWidth,
		minRow:      minRow,
		maxRow:      maxRow,
		minColumn:   minColumn,
		maxColumn:   maxColumn,
	}
}

func (e edge) Start() position.Position {
	return e.start
}

func (e edge) End() position.Position {
	return e.end
}

func (e edge) VerticallyContains(p position.Position) bool {
	return (e.start.Row() <= p.Row()) && (p.Row() <= e.end.Row())
}

func (e edge) IntersectsHorizontalLineAt(p position.Position) bool {
	centerPercentageOfEdgeHeight := int(float64(p.Row()-e.start.Row()) / float64(e.totalHeight))
	edgeColumnAtCenterRow := e.start.Column() + centerPercentageOfEdgeHeight*e.totalWidth
	return p.Row() <= edgeColumnAtCenterRow
}

func (e edge) Intersects(rect rectangle.Rectangle) bool {
	isVerticalEdge := e.start.Column() == e.end.Column()

	if isVerticalEdge {
		isRectBetweenEdgeColumns := e.start.Column() > rect.MinColumn() && e.start.Column() < rect.MaxColumn()

		if isRectBetweenEdgeColumns {
			overlapStart := int(math.Max(e.minRow, float64(rect.MinRow())))
			overlapEnd := int(math.Min(e.maxRow, float64(rect.MaxRow())))

			if overlapEnd > overlapStart {
				return true
			}
		}
	} else {
		isRectBetweenEdgeRows := e.start.Row() > rect.MinRow() && e.start.Row() < rect.MaxRow()

		if isRectBetweenEdgeRows {
			overlapStart := int(math.Max(e.minColumn, float64(rect.MinColumn())))
			overlapEnd := int(math.Min(e.maxColumn, float64(rect.MaxColumn())))

			if overlapEnd > overlapStart {
				return true
			}
		}
	}

	return false
}
