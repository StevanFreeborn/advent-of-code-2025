// Package queue provides a thread-safe FIFO queue implementation.
package queue

import "sync"

type Queue[T any] interface {
	Enqueue(item T)
	Dequeue() (T, bool)
	IsEmpty() bool
	Size() int
}

type queue[T any] struct {
	items []T
	lock  sync.Mutex
}

func New[T any]() Queue[T] {
	return &queue[T]{items: []T{}}
}

func (q *queue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, item)
}

func (q *queue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *queue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items) == 0
}

func (q *queue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return len(q.items)
}
