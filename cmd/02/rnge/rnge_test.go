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

func TestInvalidIds(t *testing.T) {
	r := rnge.From("1000-1100")
	expectedInvalidIds := []int64{1010}
	resultInvalidIds := []int64{}

	for id := range r.InvalidIds() {
		resultInvalidIds = append(resultInvalidIds, id)
	}

	if slices.Equal(resultInvalidIds, expectedInvalidIds) == false {
		t.Errorf("expected invalid IDs %v, got %v", expectedInvalidIds, resultInvalidIds)
	}
}
