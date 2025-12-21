package main

import (
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/stack"
)

const START_NODE = "you"
const END_NODE = "out"

func SolvePartOne(filePath string) int {
	adjacencyList := map[string][]string{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ": ")
		from := parts[0]
		toList := strings.Split(parts[1], " ")
		adjacencyList[from] = toList
	}

	stack := stack.New[string]()
	stack.Push(START_NODE)
	pathCount := 0

	for stack.IsEmpty() == false {
		current, _ := stack.Pop()

		if current == END_NODE {
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
