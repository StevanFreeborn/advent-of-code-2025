// Package machine defines a model for a machine and methods to manipulate it.
package machine

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/10/button"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/stack"
)

type Machine interface {
	ConfigureLights() int
	ConfigureJoltages() int
}

type machine struct {
	desiredLightState []bool
	buttons           []button.Button
	desiredJoltages   []int
}

func From(line string) Machine {
	parts := strings.Split(line, " ")
	lastPartIndex := len(parts) - 1

	lightDiagramPart := parts[0]
	desiredLightState := []bool{}

	for _, c := range lightDiagramPart {
		if c == '.' {
			desiredLightState = append(desiredLightState, false)
		}

		if c == '#' {
			desiredLightState = append(desiredLightState, true)
		}
	}

	buttonsPart := parts[1:lastPartIndex]

	buttons := []button.Button{}

	for _, bs := range buttonsPart {
		b := button.From(bs)
		buttons = append(buttons, b)
	}

	joltagesPart := parts[lastPartIndex]
	desiredJoltages := []int{}
	joltageRegex := regexp.MustCompile(`\d+`)
	matches := joltageRegex.FindAllString(joltagesPart, -1)

	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		desiredJoltages = append(desiredJoltages, num)
	}

	return machine{
		desiredLightState: desiredLightState,
		buttons:           buttons,
		desiredJoltages:   desiredJoltages,
	}
}

func (m machine) ConfigureLights() int {
	combs := m.generateCombinations(len(m.desiredLightState))

	minPresses := math.MaxInt
	found := false

	for _, comb := range combs {
		matches := true

		for i, count := range comb.deltas {
			isLightOn := count%2 != 0

			if isLightOn != m.desiredLightState[i] {
				matches = false
				break
			}
		}

		if matches {
			if comb.numPresses < minPresses {
				minPresses = comb.numPresses
				found = true
			}
		}
	}

	if found == false {
		return 0
	}

	return minPresses
}

type combination struct {
	deltas     []int
	numPresses int
}

type searchState struct {
	goal        []int
	currentCost int
	weight      int
}

// If a target is odd I must press a combination of
// buttons that contributes an odd value to the target.
// This means I can pre-compute what all combinations
// of buttons do when pressed exactly once.
// I then can look for a combination that matches the
// odd/even pattern of the target
// When I find a match I can subtract it from the target
// and then divide the target by 2 to get a new target
// I repeat this until I reach a target of all zeros

// i.e. Goal: [13, 7]
// Button A:  [1, 0]
// Button B:  [1, 1]
//
// Combinations:
// 0 presses:  [0, 0]
// 1 press:    [1, 0] (A)
// 1 press:    [1, 1] (B)
// 1 press:    [2, 1] (A, B)
//
// 1st iteration:
// Target: [13, 7] (odd, odd)
// Match:  [1, 1] (B)
// New Target: [(13-1)/2, (7-1)/2] = [6, 3]
// Presses: 1 * weight 1 = 1
//
// Second iteration:
// Target: [6, 3] (even, odd)
// Match:  [2, 1] (A, B)
// New Target: [(6-2)/2, (3-1)/2] = [2, 1]
// Presses: 2 * weight 2 = 4
//
// Third iteration:
// Target: [2, 1] (even, odd)
// Match:  [2, 1] (A, B)
// New Target: [(2-1)/2, (1-0)/2] = [0, 0]
// Presses: 2 * weight 4 = 8
//
// Total presses: 1 + 4 + 8 = 13
func (m machine) ConfigureJoltages() int {
	combinations := m.generateCombinations(len(m.desiredJoltages))

	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].numPresses < combinations[j].numPresses
	})

	stack := stack.New[searchState]()
	stack.Push(searchState{
		goal:        m.desiredJoltages,
		currentCost: 0,
		weight:      1,
	})

	minTotalCost := math.MaxInt
	foundSolution := false

	for stack.IsEmpty() == false {
		curr, _ := stack.Pop()

		if curr.currentCost >= minTotalCost {
			continue
		}

		if isZero(curr.goal) {
			if curr.currentCost < minTotalCost {
				minTotalCost = curr.currentCost
				foundSolution = true
			}

			continue
		}

		for _, combination := range combinations {
			if smallerOrEqual(combination.deltas, curr.goal) == false {
				continue
			}

			if hasSameParity(combination.deltas, curr.goal) == false {
				continue
			}

			nextGoal := make([]int, len(curr.goal))

			for i := 0; i < len(curr.goal); i++ {
				nextGoal[i] = (curr.goal[i] - combination.deltas[i]) / 2
			}

			stepCost := combination.numPresses * curr.weight

			stack.Push(searchState{
				goal:        nextGoal,
				currentCost: curr.currentCost + stepCost,
				weight:      curr.weight * 2,
			})
		}
	}

	if foundSolution == false {
		return 0
	}

	return minTotalCost
}

func (m machine) generateCombinations(size int) []combination {
	res := []combination{{
		deltas:     make([]int, size),
		numPresses: 0,
	}}

	for _, btn := range m.buttons {
		currentCount := len(res)

		for i := range currentCount {
			existing := res[i]

			newDeltas := make([]int, size)
			copy(newDeltas, existing.deltas)

			for _, switchIdx := range btn.Switches() {
				if switchIdx < size {
					newDeltas[switchIdx]++
				}
			}

			res = append(res, combination{
				deltas:     newDeltas,
				numPresses: existing.numPresses + 1,
			})
		}
	}

	return res
}

func isZero(arr []int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func smallerOrEqual(a []int, b []int) bool {
	for i := range a {
		if a[i] > b[i] {
			return false
		}
	}

	return true
}

func hasSameParity(a []int, b []int) bool {
	for i := range a {
		if a[i]%2 != b[i]%2 {
			return false
		}
	}

	return true
}
