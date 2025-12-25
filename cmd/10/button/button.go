// Package button implements a button that can control multiple switches.
package button

import (
	"fmt"
	"regexp"
	"strconv"
)

type Button interface {
	Switches() []int
}

type button struct {
	switches []int
}

func (b button) String() string {
	return fmt.Sprintf("%v", b.switches)
}

func From(line string) Button {
	buttonRegex := regexp.MustCompile(`\d+`)
	matches := buttonRegex.FindAllString(line, -1)

	switches := []int{}

	for _, m := range matches {
		num, _ := strconv.Atoi(m)
		switches = append(switches, num)
	}

	return button{
		switches: switches,
	}
}

func (b button) Switches() []int {
	return b.switches
}
