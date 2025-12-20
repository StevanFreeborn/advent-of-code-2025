// Package machine defines a model for a machine and methods to manipulate it.
package machine

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/10/button"
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
	combinations := [][]bool{}
	minPresses := math.MaxInt
	numberOfButtons := len(m.buttons)
	numberOfCombinations := int(math.Pow(2, float64(numberOfButtons)))
	currentCombination := make([]bool, numberOfButtons)

	for range numberOfCombinations {
		temp := make([]bool, numberOfButtons)
		copy(temp, currentCombination)

		combinations = append(combinations, temp)

		for j := range numberOfButtons {
			if currentCombination[j] == false {
				currentCombination[j] = true
				break
			} else {
				currentCombination[j] = false
			}
		}
	}

	for _, currentCombination := range combinations {
		currentPresses := 0
		initialLightState := make([]bool, len(m.desiredLightState))

		for bi, bs := range currentCombination {
			if bs == false {
				continue
			}

			currentPresses++
			switchesToToggle := m.buttons[bi].Switches()

			for _, switchToToggle := range switchesToToggle {
				initialLightState[switchToToggle] = !initialLightState[switchToToggle]
			}
		}

		if slices.Equal(initialLightState, m.desiredLightState) == false {
			continue
		}

		if currentPresses < minPresses {
			minPresses = currentPresses
		}
	}

	return minPresses
}

func (m machine) ConfigureJoltages() int {
	minPresses := 0
	// counters := make([]int, len(m.joltageSettings))

	// given the target joltage of a counter
	// how many times can I press a particular button
	// before making one of the counters that the button
	// affects invalid
	// 3,5,4,7
	// 0 (3)
	// 1 (1,3)
	// 2 (2)
	// 3 (2,3)
	// 4 (0,2)
	// 5 (0,1)

	// (0n * 1) + (1n * 0) + (1n * 1) + (3n * 1) = 7
	// 2n + 3n + 4n = 4

	// TODO: I need to use Gausian elimination
	// to solve this
	// TODO: Or matrix method maybe?

	rows := len(m.desiredJoltages)
	cols := len(m.buttons)
	grid := make([][]int, rows)

	for r := range rows {
		grid[r] = make([]int, cols+1)

		for i, b := range m.buttons {
			for _, sw := range b.Switches() {
				if sw == r {
					grid[r][i] = 1
				}
			}
		}

		grid[r][cols] = m.desiredJoltages[r]

		fmt.Println(grid[r])
	}

	fmt.Println()

	return minPresses
}
