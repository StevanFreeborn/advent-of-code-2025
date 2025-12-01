// Package dial provides a model for representing a dial
package dial

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/direction"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/instruction"
)

type Dial interface {
	Value() int
	Turn(instruction instruction.Instruction)
}

type dial struct {
	value int
}

func New() Dial {
	return &dial{
		value: 50,
	}
}

func (d *dial) Value() int {
	return d.value
}

func (d *dial) Turn(instruction instruction.Instruction) {
	if instruction.Direction() == direction.Left {
		d.value -= instruction.Distance()

		for d.value < 0 {
			d.value += 100
		}

		return
	}

	if instruction.Direction() == direction.Right {
		d.value += instruction.Distance()

		for d.value > 99 {
			d.value -= 100
		}

		return
	}
}
