package main

import (
	"regexp"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/region"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/shape"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const COLON = ":"
const X = "x"

func SolvePartOne(filePath string) int {
	twoNewLineRegex := regexp.MustCompile(`\r?\n\r?\n`)
	newLineRegex := regexp.MustCompile(`\r?\n`)

	input := file.ReadAllText(filePath)
	sections := twoNewLineRegex.Split(strings.TrimSpace(input), -1)

	total := 0
	shapes := map[int][]shape.Shape{}
	regions := []region.Region{}

	for _, section := range sections {
		lines := newLineRegex.Split(strings.TrimSpace(section), -1)
		header := lines[0]

		if strings.Contains(header, COLON) && strings.Contains(header, X) == false {
			shape := shape.From(lines)
			shapes[shape.Id()] = shape.GenerateVariants()
			continue
		}

		for _, line := range lines {
			if strings.Contains(line, COLON) {
				regions = append(regions, region.From(line))
			}
		}
	}

	for _, region := range regions {
		if region.CanFit(shapes) {
			total++
		}
	}

	return total
}
