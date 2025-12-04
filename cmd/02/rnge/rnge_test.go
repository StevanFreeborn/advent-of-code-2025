package rnge_test

import (
	"slices"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/02/rnge"
)

func TestFrom(t *testing.T) {
	expectedStart := int64(1000)
	expectedEnd := int64(2000)

	r := rnge.From("1000-2000")

	if r.Start() != expectedStart {
		t.Errorf("expected start %d, got %d", expectedStart, r.Start())
	}

	if r.End() != expectedEnd {
		t.Errorf("expected end %d, got %d", expectedEnd, r.End())
	}
}

func TestStart(t *testing.T) {
	expectedStart := int64(500)

	r := rnge.From("500-1500")

	if r.Start() != expectedStart {
		t.Errorf("expected start %d, got %d", expectedStart, r.Start())
	}
}

func TestEnd(t *testing.T) {
	expectedEnd := int64(1500)

	r := rnge.From("500-1500")

	if r.End() != expectedEnd {
		t.Errorf("expected end %d, got %d", expectedEnd, r.End())
	}
}

func TestInvalidIdsWithEqualValues(t *testing.T) {
	r := rnge.From("1000-1100")
	expectedInvalidIds := []int64{1010}
	resultInvalidIds := []int64{}

	for id := range r.InvalidIdsWithEqualHalves() {
		resultInvalidIds = append(resultInvalidIds, id)
	}

	if slices.Equal(resultInvalidIds, expectedInvalidIds) == false {
		t.Errorf("expected invalid IDs %v, got %v", expectedInvalidIds, resultInvalidIds)
	}
}

type TestCase struct {
	rnge               rnge.Range
	expectedInvalidIds []int64
}

func TestInvalidIdsWithTwoOrMoreSeq(t *testing.T) {
	testCases := []TestCase{
		{
			rnge:               rnge.From("11-22"),
			expectedInvalidIds: []int64{11, 22},
		},
		{
			rnge:               rnge.From("95-115"),
			expectedInvalidIds: []int64{99, 111},
		},
		{
			rnge:               rnge.From("998-1012"),
			expectedInvalidIds: []int64{999, 1010},
		},
		{
			rnge:               rnge.From("1188511880-1188511890"),
			expectedInvalidIds: []int64{1188511885},
		},
		{
			rnge:               rnge.From("222220-222224"),
			expectedInvalidIds: []int64{222222},
		},
	}

	for _, c := range testCases {
		resultInvalidIds := []int64{}

		for id := range c.rnge.InvalidIdsWithTwoOrMoreSeq() {
			resultInvalidIds = append(resultInvalidIds, id)
		}

		if slices.Equal(resultInvalidIds, c.expectedInvalidIds) == false {
			t.Errorf("expected invalid IDs %v, got %v", c.expectedInvalidIds, resultInvalidIds)
		}
	}

}
