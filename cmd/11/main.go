package main

import (
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/queue"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/stack"
)

const YOU_NODE = "you"
const OUT_NODE = "out"
const SVR_NODE = "svr"
const DAC_NODE = "dac"
const FFT_NODE = "fft"

func SolvePartOne(filePath string) int {
	adjacencyList := map[string][]string{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ": ")
		from := parts[0]
		toList := strings.Split(parts[1], " ")
		adjacencyList[from] = toList
	}

	stack := stack.New[string]()
	stack.Push(YOU_NODE)
	pathCount := 0

	for stack.IsEmpty() == false {
		current, _ := stack.Pop()

		if current == OUT_NODE {
			pathCount++
			continue
		}

		neighbors := adjacencyList[current]

		for _, n := range neighbors {
			stack.Push(n)
		}
	}

	return pathCount
}

func SolvePartTwo(filePath string) int {
	adjacencyList := map[string][]string{}
	inDegrees := map[string]int{}
	allNodes := map[string]bool{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ": ")
		from := parts[0]
		toList := strings.Split(parts[1], " ")

		allNodes[from] = true

		_, existingInDegrees := inDegrees[from]

		if existingInDegrees == false {
			inDegrees[from] = 0
		}

		for _, to := range toList {
			allNodes[to] = true
			adjacencyList[from] = append(adjacencyList[from], to)
			inDegrees[to]++
		}
	}

	queue := queue.New[string]()

	for node := range allNodes {
		if inDegrees[node] == 0 {
			queue.Enqueue(node)
		}
	}

	processOrder := []string{}

	for queue.IsEmpty() == false {
		current, _ := queue.Dequeue()
		processOrder = append(processOrder, current)

		for _, neighbor := range adjacencyList[current] {
			inDegrees[neighbor]--

			if inDegrees[neighbor] == 0 {
				queue.Enqueue(neighbor)
			}
		}
	}

	pathA := countPaths(SVR_NODE, FFT_NODE, processOrder, adjacencyList) * countPaths(FFT_NODE, DAC_NODE, processOrder, adjacencyList) * countPaths(DAC_NODE, OUT_NODE, processOrder, adjacencyList)
	pathB := countPaths(SVR_NODE, DAC_NODE, processOrder, adjacencyList) * countPaths(DAC_NODE, FFT_NODE, processOrder, adjacencyList) * countPaths(FFT_NODE, OUT_NODE, processOrder, adjacencyList)
	return pathA + pathB
}

func countPaths(start string, end string, processOrder []string, adjacencyList map[string][]string) int {
	pathsCount := map[string]int{}
	pathsCount[start] = 1

	for _, item := range processOrder {
		if pathsCount[item] == 0 {
			continue
		}

		for _, n := range adjacencyList[item] {
			pathsCount[n] += pathsCount[item]
		}
	}

	return pathsCount[end]
}
