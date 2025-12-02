// Package dial provides a model for representing a dial
package dial

import (
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/direction"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/instruction"
)

type Dial interface {
	ZeroCount() int
	Value() int
	Turn(instruction instruction.Instruction)
}

type dial struct {
	zeroCount int
	value     int
}

func New() Dial {
	return &dial{
		zeroCount: 0,
		value:     50,
	}
}

func From(value int) Dial {
	return &dial{
		zeroCount: 0,
		value:     value,
	}
}

func (d *dial) ZeroCount() int {
	return d.zeroCount
}

func (d *dial) Value() int {
	return d.value
}

func (d *dial) Turn(instruction instruction.Instruction) {
	if instruction.Direction() == direction.Left {
		for range instruction.Distance() {
			d.decreaseOneClick()

			if d.value == 0 {
				d.zeroCount++
			}

			if d.value < 0 {
				d.value += 100
			}
		}

		return
	}

	if instruction.Direction() == direction.Right {
		for range instruction.Distance() {
			d.increaseOneClick()

			if d.value >= 100 {
				d.value -= 100
			}

			if d.value == 0 {
				d.zeroCount++
			}
		}

		return
	}
}

func (d *dial) decreaseOneClick() {
	d.value--
}

func (d *dial) increaseOneClick() {
	d.value++
}
