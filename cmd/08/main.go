package main

import (
	"slices"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/estimator"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/queue"
)

func SolvePartTwo(filePath string) int {
	boxes := []box.Box{}

	for line := range file.ReadLines(filePath) {
		boxes = append(boxes, box.From(line))
	}

	estimator := estimator.From(boxes)
	circuitsMap := estimator.CreateMap()

	numberOfSets := len(boxes)

	for _, connection := range estimator.PossibleConnections() {
		startRoot := circuitsMap.FindRootBoxFor(connection.Start())
		endRoot := circuitsMap.FindRootBoxFor(connection.End())

		if startRoot == endRoot {
			continue
		}

		startRootValue, _ := circuitsMap.GetValueFor(startRoot)
		endRootValue, _ := circuitsMap.GetValueFor(endRoot)

		if startRootValue.Size() < endRootValue.Size() {
			startRootValue.UpdateParent(endRoot)
			endRootValue.IncreaseSize(startRootValue.Size())
		} else {
			endRootValue.UpdateParent(startRoot)
			startRootValue.IncreaseSize(endRootValue.Size())
		}

		circuitsMap.UpdateValueFor(startRoot, startRootValue)
		circuitsMap.UpdateValueFor(endRoot, endRootValue)

		numberOfSets--

		if numberOfSets == 1 {
			return connection.Start().X() * connection.End().X()
		}
	}

	return -1
}

func SolvePartOne(filePath string, numOfConnections int) int {
	boxes := []box.Box{}

	for line := range file.ReadLines(filePath) {
		boxes = append(boxes, box.From(line))
	}

	estimator := estimator.From(boxes)
	circuitsMap := estimator.CreateMap()

	for _, connection := range estimator.PossibleConnections()[:numOfConnections] {
		startRoot := circuitsMap.FindRootBoxFor(connection.Start())
		endRoot := circuitsMap.FindRootBoxFor(connection.End())

		if startRoot == endRoot {
			continue
		}

		startRootValue, _ := circuitsMap.GetValueFor(startRoot)
		endRootValue, _ := circuitsMap.GetValueFor(endRoot)

		if startRootValue.Size() < endRootValue.Size() {
			startRootValue.UpdateParent(endRoot)
			endRootValue.IncreaseSize(startRootValue.Size())
		} else {
			endRootValue.UpdateParent(startRoot)
			startRootValue.IncreaseSize(endRootValue.Size())
		}

		circuitsMap.UpdateValueFor(startRoot, startRootValue)
		circuitsMap.UpdateValueFor(endRoot, endRootValue)
	}

	circuitSizes := []int{}

	for root, circuit := range circuitsMap.Entries() {
		if circuit.Parent() != root {
			continue
		}

		circuitSizes = append(circuitSizes, circuit.Size())
	}

	slices.SortFunc(circuitSizes, func(a int, b int) int {
		return b - a
	})

	if len(circuitSizes) < 3 {
		panic("not enough circuits")
	}

	result := 1

	for i := range 3 {
		result *= circuitSizes[i]
	}

	return result
}

func SolvePartOneAgain(filePath string, numOfConnections int) int {
	boxes := []box.Box{}

	for line := range file.ReadLines(filePath) {
		boxes = append(boxes, box.From(line))
	}

	estimator := estimator.From(boxes)

	neighbors := map[box.Box][]box.Box{}

	for _, b := range boxes {
		neighbors[b] = []box.Box{}
	}

	for _, connection := range estimator.PossibleConnections()[:numOfConnections] {
		startNeighbors := neighbors[connection.Start()]
		endNeighbors := neighbors[connection.End()]

		startNeighbors = append(startNeighbors, connection.End())
		endNeighbors = append(endNeighbors, connection.Start())

		neighbors[connection.Start()] = startNeighbors
		neighbors[connection.End()] = endNeighbors
	}

	circuits := [][]box.Box{}
	visited := map[box.Box]bool{}

	for _, b := range boxes {
		if visited[b] {
			continue
		}

		currentCircuit := []box.Box{}
		boxesQueue := queue.New[box.Box]()
		boxesQueue.Enqueue(b)
		visited[b] = true

		for boxesQueue.IsEmpty() == false {
			current, _ := boxesQueue.Dequeue()
			currentCircuit = append(currentCircuit, current)
			currentNeighbors := neighbors[current]

			for _, n := range currentNeighbors {
				if visited[n] {
					continue
				}

				visited[n] = true
				boxesQueue.Enqueue(n)
			}
		}

		circuits = append(circuits, currentCircuit)
	}

	circuitSizes := []int{}

	for _, c := range circuits {
		circuitSizes = append(circuitSizes, len(c))
	}

	slices.SortFunc(circuitSizes, func(a int, b int) int {
		return b - a
	})

	if len(circuitSizes) < 3 {
		panic("not enough circuits")
	}

	result := 1

	for i := range 3 {
		result *= circuitSizes[i]
	}

	return result
}
