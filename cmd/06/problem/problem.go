// Package problem provides models and methods for representing a problem.
package problem

type Problem interface {
	Solve() int
}

type problem struct {
	operator string
	operands []int
}

func From(operator string, operands []int) Problem {
	return problem{
		operator: operator,
		operands: operands,
	}
}

func (p problem) Solve() int {
	result := 0

	switch p.operator {
	case "+":
		for _, operand := range p.operands {
			result += operand
		}
	case "*":
		result = 1

		for _, operand := range p.operands {
			result *= operand
		}
	}

	return result
}
