package main

import (
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const SeparatorCharacter = ","
const RangeSeparatorCharacter = "-"

type Range struct {
	Start int64
	End   int64
}

func SolvePartOne(filePath string) int64 {
	input := file.ReadAllText(filePath)
	strRanges := strings.SplitSeq(input, SeparatorCharacter)

	rngs := []Range{}

	for r := range strRanges {
		parts := strings.Split(strings.TrimSpace(r), "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end, _ := strconv.ParseInt(parts[1], 10, 64)
		rng := Range{
			Start: start,
			End:   end,
		}
		rngs = append(rngs, rng)
	}

	total := int64(0)

	for _, rng := range rngs {
		for id := rng.Start; id <= rng.End; id++ {
			idStr := strconv.FormatInt(id, 10)
			idLength := len(idStr)

			if idLength%2 != 0 {
				continue
			}

			midpoint := idLength / 2

			firstHalf := idStr[:midpoint]
			secondHalf := idStr[midpoint:]

			if firstHalf != secondHalf {
				continue
			}

			total += id
		}
	}

	return total
}
