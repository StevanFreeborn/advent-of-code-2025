// Package machine defines a model for a machine and methods to manipulate it.
package machine

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/10/button"
)

type Machine interface {
	ConfigureLights() int
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
