package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

type machine struct {
	switches []bool
	buttons  []button
}

func NewMachine(numberOfLights int, buttonsAsStrings []string) machine {
	buttons := []button{}

	for _, bs := range buttonsAsStrings {
		b := NewButton(bs)
		buttons = append(buttons, b)
	}

	switches := []bool{}

	for range numberOfLights {
		switches = append(switches, false)
	}

	return machine{
		buttons:  buttons,
		switches: switches,
	}
}

func (m machine) String() string {
	var str strings.Builder

	str.WriteString("[MACHINE]\n")

	switchStr := fmt.Sprintf("[SWITCHES]: %v", m.switches)
	str.WriteString(switchStr + "\n")

	str.WriteString("[BUTTONS]:\n")

	for _, b := range m.buttons {
		str.WriteString(b.String() + "\n")
	}

	str.WriteString("\n")

	return str.String()
}

type button struct {
	switchesControlled []int
}

func (b button) String() string {
	return fmt.Sprintf("%v", b.switchesControlled)
}

func NewButton(line string) button {
	buttonRegex := regexp.MustCompile(`\d+`)
	matches := buttonRegex.FindAllString(line, -1)

	switchesControlled := []int{}

	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		switchesControlled = append(switchesControlled, num)
	}

	return button{
		switchesControlled: switchesControlled,
	}
}

func (m machine) Configure(desiredState []bool) int {
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
		switchesCopy := []bool{}

		for _, v := range m.switches {
			switchesCopy = append(switchesCopy, v)
		}

		for bi, bs := range currentCombination {
			if bs == false {
				continue
			}

			currentPresses++
			switchesToToggle := m.buttons[bi].switchesControlled

			for _, switchToToggle := range switchesToToggle {
				switchesCopy[switchToToggle] = !switchesCopy[switchToToggle]
			}
		}

		if slices.Equal(switchesCopy, desiredState) == false {
			continue
		}

		if currentPresses < minPresses {
			minPresses = currentPresses
		}
	}

	return minPresses
}

func SolvePartOne(filePath string) int {
	total := 0

	for line := range file.ReadLines(filePath) {
		parts := strings.Split(line, " ")
		lightDiagramPart := parts[0]
		buttonsPart := parts[1 : len(parts)-1]
		// NOTE: Parse joltage here

		lightDiagram := []bool{}

		for _, c := range lightDiagramPart {
			if c == '.' {
				lightDiagram = append(lightDiagram, false)
			}

			if c == '#' {
				lightDiagram = append(lightDiagram, true)
			}
		}

		numberOfLights := len(lightDiagram)

		machine := NewMachine(numberOfLights, buttonsPart)
		buttonsPressed := machine.Configure(lightDiagram)
		total += buttonsPressed
	}

	return total
}
