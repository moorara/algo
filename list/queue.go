package list

import "github.com/moorara/algo/generic"

// Queue represents a queue abstract data type.
type Queue[T any] interface {
	Size() int
	IsEmpty() bool
	Enqueue(T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Contains(T) bool
}

type arrayQueue[T any] struct {
	nodeSize int
	equal    generic.EqualFunc[T]

	listSize   int
	frontIndex int
	rearIndex  int
	frontNode  *arrayNode[T]
	rearNode   *arrayNode[T]
}

// NewQueue creates a new array-list queue.
func NewQueue[T any](nodeSize int, equal generic.EqualFunc[T]) Queue[T] {
	return &arrayQueue[T]{
		nodeSize: nodeSize,
		equal:    equal,

		listSize:   0,
		frontIndex: -1,
		rearIndex:  -1,
		frontNode:  nil,
		rearNode:   nil,
	}
}

// Size returns the number of values in queue.
func (q *arrayQueue[T]) Size() int {
	return q.listSize
}

// IsEmpty returns true if queue is empty.
func (q *arrayQueue[T]) IsEmpty() bool {
	return q.listSize == 0
}

// Enqueue adds a new value to queue.
func (q *arrayQueue[T]) Enqueue(val T) {
	q.listSize++
	q.rearIndex++

	if q.frontNode == nil {
		q.frontIndex = 0
		q.frontNode = newArrayNode[T](q.nodeSize, nil)
		q.rearNode = q.frontNode
	} else if q.rearIndex == q.nodeSize {
		q.rearNode.next = newArrayNode[T](q.nodeSize, nil)
		q.rearNode = q.rearNode.next
		q.rearIndex = 0
	}

	q.rearNode.block[q.rearIndex] = val
}

// Dequeue removes a value from queue.
func (q *arrayQueue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	val := q.frontNode.block[q.frontIndex]
	q.frontIndex++
	q.listSize--

	if q.frontIndex == q.nodeSize {
		q.frontNode = q.frontNode.next
		q.frontIndex = 0
	}

	return val, true
}

// Peek returns the next value in queue without removing it from queue.
func (q *arrayQueue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.frontNode.block[q.frontIndex], true
}

// Contains returns true if a given value is already in queue.
func (q *arrayQueue[T]) Contains(val T) bool {
	n := q.frontNode
	i := q.frontIndex

	for n != nil && (n != q.rearNode || i <= q.rearIndex) {
		if q.equal(n.block[i], val) {
			return true
		}

		i++

		if i == q.nodeSize {
			n = n.next
			i = 0
		}
	}

	return false
}
