package list

import . "github.com/moorara/algo/generic"

// SoftQueue represents the abstract data type for a queue with soft deletion.
type SoftQueue[T any] interface {
	Size() int
	IsEmpty() bool
	Enqueue(T) int
	Dequeue() (T, int)
	Peek() (T, int)
	Contains(T) int
	Values() []T
}

type softQueue[T any] struct {
	equal EqualFunc[T]

	front int
	rear  int
	list  []T
}

// NewSoftQueue creates a new array-list queue with soft deletion.
// Deleted entries remain in the queue and are searchable.
func NewSoftQueue[T any](equal EqualFunc[T]) SoftQueue[T] {
	return &softQueue[T]{
		equal: equal,

		front: 0,
		rear:  -1,
		list:  make([]T, 0),
	}
}

// Size returns the number of values in the queue.
func (q *softQueue[T]) Size() int {
	return q.rear - q.front + 1
}

// IsEmpty returns true if the queue is empty.
func (q *softQueue[T]) IsEmpty() bool {
	return q.front > q.rear
}

// Enqueue inserts a new value to the queue.
func (q *softQueue[T]) Enqueue(val T) int {
	q.list = append(q.list, val)

	if len(q.list) == 1 {
		q.front, q.rear = 0, 0
	} else {
		q.rear++
	}

	return q.rear
}

// Dequeue deletes a value from the queue.
// The deletion is soft and the entries remain in the queue.
// They are searchable using the Contains method.
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

// Peek returns the next value in queue without deleting it from the queue.
func (q *softQueue[T]) Peek() (T, int) {
	if q.IsEmpty() {
		var zero T
		return zero, -1
	}

	return q.list[q.front], q.front
}

// Contains returns true if a given value is either in the queue or deleted in the past.
func (q *softQueue[T]) Contains(val T) int {
	for i, v := range q.list {
		if q.equal(v, val) {
			return i
		}
	}

	return -1
}

// Values returns the list of all values in the queue including the deleted ones.
func (q *softQueue[T]) Values() []T {
	vals := make([]T, len(q.list))
	copy(vals, q.list)

	return vals
}
