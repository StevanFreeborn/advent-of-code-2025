// Package stack provides a thread-safe LIFO stack implementation.
package stack

import "sync"

type Stack[T any] interface {
	Push(item T)
	Peek() (T, bool)
	Pop() (T, bool)
	IsEmpty() bool
	Size() int
}

type stack[T any] struct {
	items []T
	lock  sync.Mutex
}

func New[T any]() Stack[T] {
	return &stack[T]{items: []T{}}
}

func (s *stack[T]) Push(item T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, item)
}

func (s *stack[T]) Pop() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.items)

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastItemIndex := length - 1
	item := s.items[lastItemIndex]
	s.items = s.items[:lastItemIndex]
	return item, true
}

func (s *stack[T]) Peek() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.items)

	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastItemIndex := length - 1
	item := s.items[lastItemIndex]
	return item, true
}

func (s *stack[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items) == 0
}

func (s *stack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}
