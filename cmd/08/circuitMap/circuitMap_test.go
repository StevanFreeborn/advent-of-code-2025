package circuitMap_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/circuitMap"
)

func TestFrom(t *testing.T) {
	boxes := []box.Box{
		box.From("0,0,0"),
		box.From("1,1,1"),
	}
	expectedNumOfEntries := len(boxes)

	result := circuitMap.From(boxes)

	if result.Length() != expectedNumOfEntries {
		t.Errorf("CreateMap() returned map with %d entries; expected %d", result.Length(), expectedNumOfEntries)
	}

	for _, b := range boxes {
		_, exists := result.GetValueFor(b)

		if !exists {
			t.Errorf("CreateMap() result missing entry for box %v", b)
		}
	}
}

func TestGetAndUpdateValueFor(t *testing.T) {
	b := box.From("2,2,2")
	cMap := circuitMap.From([]box.Box{b})

	value, exists := cMap.GetValueFor(b)

	if !exists {
		t.Fatalf("GetValueFor() did not find value for box %v", b)
	}

	if value.Parent() != b {
		t.Errorf("GetValueFor() returned value with Parent() = %v; expected %v", value.Parent(), b)
	}

	newParent := box.From("3,3,3")

	value.UpdateParent(newParent)
	cMap.UpdateValueFor(b, value)

	updatedValue, exists := cMap.GetValueFor(b)

	if !exists {
		t.Fatalf("GetValueFor() after UpdateValueFor() did not find value for box %v", b)
	}

	if updatedValue.Parent() != newParent {
		t.Errorf("After UpdateValueFor(), GetValueFor() returned value with Parent() = %v; expected %v", updatedValue.Parent(), newParent)
	}
}

func TestFindRootBoxFor(t *testing.T) {
	b1 := box.From("0,0,0")
	b2 := box.From("1,1,1")
	b3 := box.From("2,2,2")

	cMap := circuitMap.From([]box.Box{b1, b2, b3})

	// Manually create connections
	value2, _ := cMap.GetValueFor(b2)
	value2.UpdateParent(b1)
	cMap.UpdateValueFor(b2, value2)

	value3, _ := cMap.GetValueFor(b3)
	value3.UpdateParent(b2)
	cMap.UpdateValueFor(b3, value3)

	root := cMap.FindRootBoxFor(b3)

	if root != b1 {
		t.Errorf("FindRootBoxFor(%v) = %v; expected %v", b3, root, b1)
	}
}

func TestEntries(t *testing.T) {
	b1 := box.From("0,0,0")
	b2 := box.From("1,1,1")

	cMap := circuitMap.From([]box.Box{b1, b2})
	entries := cMap.Entries()

	if len(entries) != 2 {
		t.Errorf("Entries() returned %d entries; expected 2", len(entries))
	}

	for _, b := range []box.Box{b1, b2} {
		if _, exists := entries[b]; !exists {
			t.Errorf("Entries() missing entry for box %v", b)
		}
	}
}
