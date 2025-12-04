package main

import (
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/02/rnge"
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const SeparatorCharacter = ","

func SolvePartOne(filePath string) int64 {
	input := file.ReadAllText(filePath)
	strRanges := strings.SplitSeq(input, SeparatorCharacter)

	total := int64(0)

	for r := range strRanges {
		rng := rnge.From(r)

		for id := range rng.InvalidIds() {
			total += id
		}
	}

	return total
}
