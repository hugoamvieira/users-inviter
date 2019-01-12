package data

import (
	"errors"
	"sync"
)

var (
	// ErrEmptyQueue is returned when there's no items left in the queue
	ErrEmptyQueue = errors.New("The queue is empty")
)

// Queue implements a thread-safe and type-safe queue using an array of []byte and a mutex
// It's a bit heavy on copies (1 per dequeue) and it trusts append to handle array expansions, but
// this could be changed.
// Decided to do it this way because, while it isn't right now, I wanted to make
// this program very easy to make concurrent.
type Queue struct {
	arr [][]byte
	mu  sync.Mutex
}

// NewQueue returns a new queue (which uses a slice as a list as an underlying data structure)
func NewQueue() *Queue {
	return &Queue{
		arr: make([][]byte, 0),
	}
}

// Enqueue puts an element (of type []byte) into the queue
func (q *Queue) Enqueue(el []byte) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.arr = append(q.arr, el)
}

// Dequeue gets the first element and returns it (removing it from the queue)
func (q *Queue) Dequeue() ([]byte, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.arr) == 0 {
		return nil, ErrEmptyQueue
	}

	el := q.arr[0]

	// Delete q.arr[0] from the list. Doing it like this so that the internal
	// slice structure is able to stop referencing the value and the GC can then
	// collect it.
	copy(q.arr[0:], q.arr[1:])
	q.arr[len(q.arr)-1] = nil
	q.arr = q.arr[:len(q.arr)-1]

	return el, nil
}
