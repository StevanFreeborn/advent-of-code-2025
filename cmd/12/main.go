package main

import (
	"fmt"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/shape"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

// TODO: Make the splitting
// safe for LF and CRLF
const NEWLINE = "\r\n"
const COLON = ":"
const X = "x"

func SolvePartOne(filePath string) int {
	input := file.ReadAllText(filePath)
	sections := strings.Split(strings.TrimSpace(input), NEWLINE+NEWLINE)

	shapes := []shape.Shape{}
	regions := []string{}

	for _, section := range sections {
		lines := strings.Split(strings.TrimSpace(section), NEWLINE)
		header := lines[0]

		if strings.Contains(header, COLON) && strings.Contains(header, X) == false {
			shapes = append(shapes, shape.From(lines))
			continue
		}

		for _, line := range lines {
			if strings.Contains(line, COLON) {
				regions = append(regions, line)
			}
		}
	}

	// TODO: We are currently not generating
	// all the expected variants. Need to debug this
	// WRITE AUTOMATED UNIT TESTS DUMMY
	// WHEN YOU READ THIS IN THE FUTURE YOU ARE GOING
	// TO WANT TO IGNORE IT...DON'T!
	// - Past Stevan
	for _, s := range shapes {
		if s.Id() != 1 && s.Id() != 2 && s.Id() != 0 {
			continue
		}

		fmt.Println("SHAPE ID", s.Id())

		for _, v := range s.GenerateVariants() {
			fmt.Println(v)
		}
	}

	return 0
}
