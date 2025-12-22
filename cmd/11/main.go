package main

import (
	"fmt"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
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

type node struct {
	value               string
	numberOfParentPaths int
}

func SolvePartTwo(filePath string) int {
	adjacencyList := map[string][]string{}

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, ": ")
		from := parts[0]
		toList := strings.Split(parts[1], " ")
		adjacencyList[from] = toList
	}

	stack := stack.New[node]()
	stack.Push(node{
		value: SVR_NODE,
	})
	visisted := map[node]int{}
	pathCount := 0

	for stack.IsEmpty() == false {
		current, _ := stack.Pop()

		_, ok := visisted[current]

		if ok {
			visisted[current]++
		} else {
			visisted[current] = 1
		}

		if current.value == OUT_NODE {
			pathCount++
			continue
		}

		neighbors := adjacencyList[current.value]

		for _, n := range neighbors {
			stack.Push(node{
				value: n,
			})
		}
	}

	fmt.Println(visisted)

	return pathCount
}
