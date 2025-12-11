package main

import (
	"regexp"
	"strconv"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	inputLength := len(input)
	lastLineIndex := inputLength - 1

	operandLines := input[:lastLineIndex]
	operators := input[lastLineIndex]
	listOfOperands := [][]int{}

	for _, line := range operandLines {
		parsedOperands := parseOperands(line)
		listOfOperands = append(listOfOperands, parsedOperands)
	}

	ops := parseOperators(operators)

	total := 0

	for i := range len(ops) {
		operator := ops[i]

		result := 0

		for _, operands := range listOfOperands {
			if result == 0 {
				result = operands[i]
				continue
			}

			if operator == "*" {
				result *= operands[i]
				continue
			}

			if operator == "+" {
				result += operands[i]
				continue
			}
		}

		total += result
	}

	return total
}

func parseOperators(line string) []string {
	operatorsRegex := regexp.MustCompile(`(\*|\+)`)
	matches := operatorsRegex.FindAllStringSubmatch(line, -1)

	operators := []string{}

	for _, match := range matches {
		operators = append(operators, match[1])
	}

	return operators
}

func parseOperands(line string) []int {
	numbersRegex := regexp.MustCompile(`(\d+)`)
	matches := numbersRegex.FindAllStringSubmatch(line, -1)

	operands := []int{}

	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		operands = append(operands, num)
	}

	return operands
}
