// Package estimator provides model and methods
// for answering questions about cable needs
// for connecting junction boxes.
package estimator

import (
	"slices"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/circuitMap"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/connection"
)

type Estimator interface {
	PossibleConnections() []connection.Connection
	CreateMap() circuitMap.CircuitMap
}

type estimator struct {
	possibleConnections []connection.Connection
	boxes               []box.Box
}

func From(boxes []box.Box) Estimator {
	return estimator{
		boxes:               boxes,
		possibleConnections: createAllPossibleConnections(boxes),
	}
}

func (e estimator) PossibleConnections() []connection.Connection {
	return e.possibleConnections
}

func (e estimator) CreateMap() circuitMap.CircuitMap {
	return circuitMap.From(e.boxes)
}

func createAllPossibleConnections(boxes []box.Box) []connection.Connection {
	connections := []connection.Connection{}
	numOfBoxes := len(boxes)

	for i := range numOfBoxes {
		for j := i + 1; j < numOfBoxes; j++ {
			start := boxes[i]
			end := boxes[j]
			conn := connection.From(start, end)
			connections = append(connections, conn)
		}
	}

	slices.SortFunc(connections, func(a connection.Connection, b connection.Connection) int {
		return int(a.Distance() - b.Distance())
	})

	return connections
}
