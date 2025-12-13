package queue_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/queue"
)

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := queue.New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expectedValues := []int{1, 2, 3}

	for _, expected := range expectedValues {
		value, ok := q.Dequeue()

		if !ok {
			t.Errorf("expected dequeue to be successful, but it failed")
		}

		if value != expected {
			t.Errorf("expected value to be %d, got %d", expected, value)
		}

	}

	_, ok := q.Dequeue()

	if ok {
		t.Errorf("expected dequeue to fail on empty queue, but it succeeded")
	}

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty, but it is not")
	}
}

func TestQueue_Size(t *testing.T) {
	q := queue.New[string]()

	if q.Size() != 0 {
		t.Errorf("expected size to be 0, got %d", q.Size())
	}

	q.Enqueue("a")
	q.Enqueue("b")

	if q.Size() != 2 {
		t.Errorf("expected size to be 2, got %d", q.Size())
	}

	q.Dequeue()

	if q.Size() != 1 {
		t.Errorf("expected size to be 1, got %d", q.Size())
	}
}

func TestQueue_ConcurrentAccess(t *testing.T) {
	q := queue.New[int]()
	done := make(chan bool)
	numGoroutines := 100
	numItemsPerGoroutine := 100

	for i := range numGoroutines {
		go func(start int) {
			for j := range numItemsPerGoroutine {
				q.Enqueue(start + j)
			}
			done <- true
		}(i * numItemsPerGoroutine)
	}

	for range numGoroutines {
		<-done
	}

	expectedSize := numGoroutines * numItemsPerGoroutine

	if q.Size() != expectedSize {
		t.Errorf("expected size to be %d, got %d", expectedSize, q.Size())
	}
}
