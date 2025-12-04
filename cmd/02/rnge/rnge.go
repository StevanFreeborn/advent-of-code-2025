// Package rnge provides a simple structure to represent a range with a start and end value.
package rnge

import (
	"iter"
	"strconv"
	"strings"
)

const RangeSeparatorCharacter = "-"

type Range interface {
	InvalidIdsWithEqualHalves() iter.Seq[int64]
	InvalidIdsWithTwoOrMoreSeq() iter.Seq[int64]
	Start() int64
	End() int64
}

type rnge struct {
	start int64
	end   int64
}

func From(rngeStr string) Range {
	parts := strings.Split(strings.TrimSpace(rngeStr), RangeSeparatorCharacter)
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)

	return rnge{
		start: start,
		end:   end,
	}
}

func (r rnge) Start() int64 {
	return r.start
}

func (r rnge) End() int64 {
	return r.end
}

func (r rnge) InvalidIdsWithEqualHalves() iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for id := r.start; id <= r.end; id++ {
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

			ok := yield(id)

			if !ok {
				return
			}
		}
	}
}

func (r rnge) InvalidIdsWithTwoOrMoreSeq() iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for id := r.start; id <= r.end; id++ {
			idStr := strconv.FormatInt(id, 10)
			idLength := len(idStr)

			isValid := true

			for pLen := 1; pLen <= idLength/2; pLen++ {
				v := idLength % pLen
				repeatedPatterDoesNotFit := v != 0

				if repeatedPatterDoesNotFit {
					continue
				}

				pattern := idStr[:pLen]
				numOfRepeats := idLength / pLen
				repeatedPattern := strings.Repeat(pattern, numOfRepeats)

				if repeatedPattern != idStr {
					continue
				}

				isValid = false
			}

			if isValid {
				continue
			}

			ok := yield(id)

			if !ok {
				return
			}
		}
	}
}
