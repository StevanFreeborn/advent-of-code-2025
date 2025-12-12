package main

import (
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/06/problem"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	numOfRows := len(input)
	numOfCols := len(strings.Fields(input[0]))
	operatorRowNumber := numOfRows - 1
	operators := strings.Fields(input[operatorRowNumber])

	total := 0

	for col := range numOfCols {
		operator := operators[col]
		operands := []int{}

		for row := range operatorRowNumber {
			operand, _ := strconv.Atoi(strings.Fields(input[row])[col])
			operands = append(operands, operand)
		}

		problem := problem.From(operator, operands)
		result := problem.Solve()
		total += result
	}

	return total
}

func SolvePartTwo(filePath string) int {
	input := file.ReadAllLines(filePath)
	numOfRows := len(input)
	operatorRowIndex := numOfRows - 1
	operatorRow := input[operatorRowIndex]
	operatorRowLength := len(operatorRow)

	total := 0
	operands := []int{}

	for operatorIndex := operatorRowLength - 1; operatorIndex >= 0; operatorIndex-- {
		var operandBuilder strings.Builder

		for row := range operatorRowIndex {
			v := string(input[row][operatorIndex])
			operandBuilder.WriteString(v)
		}

		operandString := strings.TrimSpace(operandBuilder.String())
		operand, _ := strconv.Atoi(operandString)
		operands = append(operands, operand)
		operandBuilder.Reset()

		operator := string(operatorRow[operatorIndex])

		if operator != " " {
			total += problem.From(operator, operands).Solve()
			operatorIndex--
			operands = []int{}
		}
	}

	return total
}
