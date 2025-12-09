package main

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const RangeSeparatorCharacter = "-"

type rnge struct {
	start int64
	end   int64
}

func NewRange(rngeStr string) rnge {
	parts := strings.Split(strings.TrimSpace(rngeStr), RangeSeparatorCharacter)
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)

	return rnge{
		start: start,
		end:   end,
	}
}

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)

	parseIds := false

	ids := []int64{}
	ranges := []rnge{}

	for _, line := range input {
		if line == "" {
			parseIds = true
		}

		if parseIds {
			id, _ := strconv.ParseInt(line, 10, 64)
			ids = append(ids, id)
			continue
		}

		rnge := NewRange(line)
		ranges = append(ranges, rnge)
	}

	slices.SortFunc(ranges, func(a rnge, b rnge) int {
		return cmp.Compare(a.start, b.start)
	})

	mergedRanges := []rnge{
		ranges[0],
	}

	for _, currentRange := range ranges {
		lastMergedRange := mergedRanges[len(mergedRanges)-1]

		if currentRange.start > lastMergedRange.end {
			mergedRanges = append(mergedRanges, currentRange)
			continue
		}

		lastMergedRange.end = int64(math.Max(float64(lastMergedRange.end), float64(currentRange.end)))
		mergedRanges[len(mergedRanges)-1] = lastMergedRange
	}

	total := 0

	for _, id := range ids {
		for _, rnge := range ranges {
			if id >= rnge.start && id <= rnge.end {
				total++
				break
			}
		}

	}

	return total
}
