// Package circuitMap provides a data structure to manage and track connected junction boxes.
package circuitMap

import "github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"

type CircuitMap interface {
	Length() int
	GetValueFor(b box.Box) (CircuitMapValue, bool)
	UpdateValueFor(b box.Box, value CircuitMapValue)
	FindRootBoxFor(b box.Box) box.Box
	Entries() map[box.Box]CircuitMapValue
}

type CircuitMapValue interface {
	Parent() box.Box
	Size() int
	UpdateParent(newParent box.Box)
	IncreaseSize(by int)
}

type circuitMap map[box.Box]CircuitMapValue

func From(boxes []box.Box) CircuitMap {
	circuitsMap := circuitMap{}

	for _, b := range boxes {
		circuitsMap[b] = &circuitMapValue{
			parent: b,
			size:   1,
		}
	}

	return circuitsMap
}

func (c circuitMap) Length() int {
	return len(c)
}

func (c circuitMap) GetValueFor(b box.Box) (CircuitMapValue, bool) {
	v, exists := c[b]
	return v, exists
}

func (c circuitMap) UpdateValueFor(b box.Box, value CircuitMapValue) {
	c[b] = value
}

func (c circuitMap) FindRootBoxFor(b box.Box) box.Box {
	root := b

	for c[root].Parent() != root {
		root = c[root].Parent()
	}

	current := b

	for current != root {
		item := c[current]
		next := item.Parent()
		item.UpdateParent(root)
		c[current] = item
		current = next
	}

	return root
}

func (c circuitMap) Entries() map[box.Box]CircuitMapValue {
	return c
}

type circuitMapValue struct {
	parent box.Box
	size   int
}

func (c *circuitMapValue) Parent() box.Box {
	return c.parent
}

func (c *circuitMapValue) Size() int {
	return c.size
}

func (c *circuitMapValue) UpdateParent(newParent box.Box) {
	c.parent = newParent
}

func (c *circuitMapValue) IncreaseSize(by int) {
	c.size += by
}
