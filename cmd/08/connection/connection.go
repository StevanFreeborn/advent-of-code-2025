// Package connection provides a model and methods for connections between junction boxes.
package connection

import "github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"

type Connection interface {
	Start() box.Box
	End() box.Box
	Distance() float64
}

type connection struct {
	start    box.Box
	end      box.Box
	distance float64
}

func (c connection) Start() box.Box {
	return c.start
}

func (c connection) End() box.Box {
	return c.end
}

func (c connection) Distance() float64 {
	return c.distance
}

func From(start box.Box, end box.Box) Connection {
	return connection{
		start:    start,
		end:      end,
		distance: start.DistanceFrom(end),
	}
}
