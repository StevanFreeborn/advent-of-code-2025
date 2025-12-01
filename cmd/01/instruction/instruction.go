// Package instruction provides a model for representing dial instructions
package instruction

import "strconv"

type Instruction interface {
	Direction() string
	Distance() int
}

type instruction struct {
	direction string
	distance  int
}

func FromLine(line string) Instruction {
	dir := string(line[0])
	dist := line[1:]
	distAsInt, _ := strconv.Atoi(dist)

	return instruction{
		direction: dir,
		distance:  distAsInt,
	}
}

func FromParts(direction string, distance int) Instruction {
	return instruction{
		direction: direction,
		distance:  distance,
	}
}

func (i instruction) Direction() string {
	return i.direction
}

func (i instruction) Distance() int {
	return i.distance
}
