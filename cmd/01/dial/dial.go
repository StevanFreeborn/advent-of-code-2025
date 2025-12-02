// Package dial provides a model for representing a dial
package dial

import (
	"fmt"

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
	logPrefix := fmt.Sprintf("[TURN %s %d]", instruction.Direction(), instruction.Distance())

	if d.value < 0 || d.value > 99 {
		panic("NEVER SHALL YOU PASS!!!")
	}

	// d.value is always 0 to 99

	if instruction.Direction() == direction.Left {
		if d.value == 0 {
			d.zeroCount--
		}

		d.value -= instruction.Distance()

		// that here we are at 0

		if d.value == 0 {
			fmt.Printf("%s zeroCount++\n", logPrefix)
			d.zeroCount++
		}

		for d.value < 0 {
			// WE NEVER GET HERE
			// UNLESS d.value is -1 or less
			currentValue := d.value
			d.zeroCount++
			d.value += 100
			fmt.Printf("%s zeroCount++ => Correcting %d to %d (%d)\n", logPrefix, currentValue, d.value, d.zeroCount)
		}

		return
	}

	if instruction.Direction() == direction.Right {
		d.value += instruction.Distance()

		// that here we at 100 or the equilavent
		// of 0. this is counted

		for d.value > 99 {
			// WE NEVER GET HERE
			// UNLESS d.value is 100 or more
			currentValue := d.value
			d.zeroCount++
			d.value -= 100
			fmt.Printf("%s zeroCount++ => Correcting %d to %d (%d)\n", logPrefix, currentValue, d.value, d.zeroCount)
		}

		return
	}
}
