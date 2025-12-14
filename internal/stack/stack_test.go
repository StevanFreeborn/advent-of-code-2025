package stack_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/stack"
)

func TestNew(t *testing.T) {
	s := stack.New[int]()

	if s == nil {
		t.Errorf("Expected new stack to be non-nil")
	}

	if !s.IsEmpty() {
		t.Errorf("Expected new stack to be empty")
	}

	if s.Size() != 0 {
		t.Errorf("Expected new stack size to be 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("Expected stack to be non-empty after push")
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1 after one push, got %d", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)

	item, ok := s.Pop()

	if !ok {
		t.Errorf("Expected Pop to succeed")
	}

	if item != 2 {
		t.Errorf("Expected popped item to be 2, got %d", item)
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1 after one pop, got %d", s.Size())
	}

	item, ok = s.Pop()

	if !ok {
		t.Errorf("Expected Pop to succeed")
	}

	if item != 1 {
		t.Errorf("Expected popped item to be 1, got %d", item)
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all items")
	}

	item, ok = s.Pop()

	if ok {
		t.Errorf("Expected Pop to fail on empty stack")
	}

	var zero int

	if item != zero {
		t.Errorf("Expected popped item to be zero value, got %d", item)
	}
}

func TestPeek(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)

	item, ok := s.Peek()

	if !ok {
		t.Errorf("Expected Peek to succeed")
	}

	if item != 2 {
		t.Errorf("Expected peeked item to be 2, got %d", item)
	}

	if s.Size() != 2 {
		t.Errorf("Expected stack size to remain 2 after Peek, got %d", s.Size())
	}

	s.Pop()

	item, ok = s.Peek()

	if !ok {
		t.Errorf("Expected Peek to succeed")
	}

	if item != 1 {
		t.Errorf("Expected peeked item to be 1, got %d", item)
	}

	s.Pop()

	item, ok = s.Peek()

	if ok {
		t.Errorf("Expected Peek to fail on empty stack")
	}

	var zero int

	if item != zero {
		t.Errorf("Expected peeked item to be zero value, got %d", item)
	}
}

func TestIsEmpty(t *testing.T) {
	s := stack.New[int]()

	if !s.IsEmpty() {
		t.Errorf("Expected new stack to be empty")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("Expected stack to be non-empty after push")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all items")
	}
}

func TestSize(t *testing.T) {
	s := stack.New[int]()

	if s.Size() != 0 {
		t.Errorf("Expected new stack size to be 0, got %d", s.Size())
	}

	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1 after one push, got %d", s.Size())
	}
}

func TestConcurrentAccess(t *testing.T) {
	s := stack.New[int]()
	done := make(chan bool)

	go func() {
		for i := range 1000 {
			s.Push(i)
		}
		done <- true
	}()

	<-done

	if s.Size() != 1000 {
		t.Errorf("Expected stack size to be 1000 after concurrent pushes, got %d", s.Size())
	}

	go func() {
		for range 1000 {
			s.Pop()
		}
		done <- true
	}()

	<-done

	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0 after concurrent pops, got %d", s.Size())
	}
}
