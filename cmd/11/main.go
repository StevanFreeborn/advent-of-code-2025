package main

import (
	"fmt"
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

type node struct {
	value   string
	fftSeen bool
	dacSeen bool
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
	pathCount := 0

	for stack.IsEmpty() == false {
		current, _ := stack.Pop()

		if current.value == FFT_NODE {
			current.fftSeen = true
		}

		if current.value == DAC_NODE {
			current.dacSeen = true
		}

		if current.value == OUT_NODE && current.dacSeen && current.fftSeen {
			pathCount++
			continue
		}

		neighbors := adjacencyList[current.value]

		for _, n := range neighbors {
			stack.Push(node{
				value:   n,
				fftSeen: current.fftSeen,
				dacSeen: current.dacSeen,
			})
		}
	}

	return pathCount
}

func SolvePartTwoAgain(filePath string) int {
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

	pathsCount := map[string]int{}
	pathsCount[SVR_NODE] = 1

	for _, item := range processOrder {
		if pathsCount[item] == 0 {
			continue
		}

		for _, n := range adjacencyList[item] {
			pathsCount[n] += pathsCount[item]
		}
	}

	// TODO: Could I process the graph
	// in smaller pieces and then multiply
	// those path to figure out what the
	// total number of paths are.
	fmt.Println(pathsCount[OUT_NODE])
	fmt.Println(pathsCount[DAC_NODE])
	fmt.Println(pathsCount[FFT_NODE])

	return 0
}
