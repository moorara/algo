package list

import "github.com/moorara/algo/generic"

// SoftQueue represents the abstract data type for a queue with soft deletion.
type SoftQueue[T any] interface {
	// Size returns the number of values in the queue.
	Size() int

	// IsEmpty returns true if the queue is empty.
	IsEmpty() bool

	// Enqueue inserts a new value to the queue.
	// It returns the index of the newly enqueued value in the queue.
	Enqueue(T) int

	// Dequeue deletes a value from the queue.
	// The deletion is soft and the entries remain in the queue.
	// Deleted entries are searchable using the Contains method.
	// The second return value is the index of the dequeued value in the queue.
	Dequeue() (T, int)

	// Peek returns the next value in queue without deleting it from the queue.
	// The second return value is the index of the peeked value in the queue.
	Peek() (T, int)

	// Contains returns true if a given value is either in the queue or deleted in the past.
	// If the value is found, its index in the queue is returned; otherwise, -1 is returned.
	Contains(T) int

	// Values returns the list of all values in the queue including the deleted ones.
	Values() []T
}

type softQueue[T any] struct {
	equal generic.EqualFunc[T]

	front int
	rear  int
	list  []T
}

// NewSoftQueue creates a new array-list queue with soft deletion.
// Deleted entries remain in the queue and are searchable.
func NewSoftQueue[T any](equal generic.EqualFunc[T]) SoftQueue[T] {
	return &softQueue[T]{
		equal: equal,

		front: 0,
		rear:  -1,
		list:  make([]T, 0),
	}
}

func (q *softQueue[T]) Size() int {
	return q.rear - q.front + 1
}

func (q *softQueue[T]) IsEmpty() bool {
	return q.front > q.rear
}

func (q *softQueue[T]) Enqueue(val T) int {
	q.list = append(q.list, val)

	if len(q.list) == 1 {
		q.front, q.rear = 0, 0
	} else {
		q.rear++
	}

	return q.rear
}

func (q *softQueue[T]) Dequeue() (T, int) {
	if q.IsEmpty() {
		var zero T
		return zero, -1
	}

	val := q.list[q.front]
	i := q.front
	q.front++

	return val, i
}

func (q *softQueue[T]) Peek() (T, int) {
	if q.IsEmpty() {
		var zero T
		return zero, -1
	}

	return q.list[q.front], q.front
}

func (q *softQueue[T]) Contains(val T) int {
	for i, v := range q.list {
		if q.equal(v, val) {
			return i
		}
	}

	return -1
}

func (q *softQueue[T]) Values() []T {
	vals := make([]T, len(q.list))
	copy(vals, q.list)

	return vals
}
