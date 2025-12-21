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
	matrix := m.createMatrix()

	eliminated := performGaussianElimination(matrix)
	pivots, freeVars := analyzeMatrix(eliminated)

	numVars := len(matrix[0]) - 1
	values := make([]int, numVars)
	bestSolution := Solution{sum: math.MaxInt}

	iterativeSearch(freeVars, pivots, eliminated, values, &bestSolution)

	return bestSolution.sum
}

type Solution struct {
	values []int
	sum    int
	found  bool
}

func (m machine) createMatrix() [][]float64 {
	rows := len(m.desiredJoltages)
	cols := len(m.buttons)
	matrix := make([][]float64, rows)

	for r := range rows {
		matrix[r] = make([]float64, cols+1)

		for i, b := range m.buttons {
			for _, sw := range b.Switches() {
				if sw == r {
					matrix[r][i] = 1
				}
			}
		}

		matrix[r][cols] = float64(m.desiredJoltages[r])
	}

	return matrix
}

func performGaussianElimination(m [][]float64) [][]float64 {
	rows := len(m)
	cols := len(m[0])
	pivotColumn := 0

	mCopy := make([][]float64, rows)

	for i := range rows {
		mCopy[i] = make([]float64, cols)
		copy(mCopy[i], m[i])
	}

	for r1 := range rows {
		if cols <= pivotColumn {
			return mCopy
		}

		currentRow := r1

		for mCopy[currentRow][pivotColumn] == 0 {
			currentRow++

			if rows == currentRow {
				currentRow = r1
				pivotColumn++

				if cols == pivotColumn {
					return mCopy
				}
			}
		}

		mCopy[currentRow], mCopy[r1] = mCopy[r1], mCopy[currentRow]

		pivotValue := mCopy[r1][pivotColumn]

		if pivotValue != 0 {
			for j := range cols {
				mCopy[r1][j] /= pivotValue
			}
		}

		for r2 := range rows {
			if r2 != r1 {
				factor := mCopy[r2][pivotColumn]

				for col := range cols {
					mCopy[r2][col] -= factor * mCopy[r1][col]
				}
			}
		}

		pivotColumn++
	}

	return mCopy
}

func analyzeMatrix(m [][]float64) (map[int]int, []int) {
	pivots := make(map[int]int)

	cols := len(m[0])
	numVars := cols - 1

	isFree := make([]bool, numVars)

	for i := range isFree {
		isFree[i] = true
	}

	rows := len(m)

	for r := range rows {
		for c := 0; c < cols-1; c++ {
			if math.Abs(m[r][c]-1.0) < 1e-9 {
				pivots[c] = r
				isFree[c] = false
				break
			}
		}
	}

	freeVars := []int{}

	for i, free := range isFree {
		if free {
			freeVars = append(freeVars, i)
		}
	}

	return pivots, freeVars
}

func iterativeSearch(freeVars []int, pivots map[int]int, matrix [][]float64, values []int, best *Solution) {
	if len(freeVars) == 0 {
		evaluateSolution(pivots, matrix, values, best)
		return
	}

	counters := make([]int, len(freeVars))
	limit := 250

	for {
		for i, counterVal := range counters {
			values[freeVars[i]] = counterVal
		}

		evaluateSolution(pivots, matrix, values, best)

		idx := len(counters) - 1

		for idx >= 0 {
			counters[idx]++

			if counters[idx] > limit {
				counters[idx] = 0
				idx--
			} else {
				break
			}
		}

		if idx < 0 {
			break
		}
	}
}

func evaluateSolution(pivots map[int]int, m [][]float64, values []int, best *Solution) {
	isValid := true
	currentSum := 0
	cols := len(m[0])

	for col, row := range pivots {
		sum := m[row][cols-1]

		for c := 0; c < cols-1; c++ {
			if c != col {
				coeff := m[row][c]
				sum -= coeff * float64(values[c])
			}
		}

		values[col] = int(math.Round(sum))
	}

	for _, v := range values {
		if v < 0 {
			isValid = false
			break
		}

		currentSum += v
	}

	if isValid {
		if currentSum < best.sum {
			best.sum = currentSum

			best.values = make([]int, len(values))
			copy(best.values, values)

			best.found = true
		}
	}
}
