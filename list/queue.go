package list

import . "github.com/moorara/algo/generic"

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
	equal    EqualFunc[T]

	listSize   int
	frontIndex int
	rearIndex  int
	frontNode  *arrayNode[T]
	rearNode   *arrayNode[T]
}

// NewQueue creates a new array-list queue.
func NewQueue[T any](nodeSize int, equal EqualFunc[T]) Queue[T] {
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

// Size returns the number of values in the queue.
func (q *arrayQueue[T]) Size() int {
	return q.listSize
}

// IsEmpty returns true if the queue is empty.
func (q *arrayQueue[T]) IsEmpty() bool {
	return q.listSize == 0
}

// Enqueue adds a new value to the queue.
func (q *arrayQueue[T]) Enqueue(val T) {
	q.listSize++
	q.rearIndex++

	if q.frontNode == nil {
		q.frontNode, q.frontIndex = newArrayNode[T](q.nodeSize, nil), 0
		q.rearNode = q.frontNode
	} else if q.rearIndex == q.nodeSize {
		q.rearNode.next = newArrayNode[T](q.nodeSize, nil)
		q.rearNode, q.rearIndex = q.rearNode.next, 0
	}

	q.rearNode.block[q.rearIndex] = val
}

// Dequeue removes a value from the queue.
func (q *arrayQueue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	val := q.frontNode.block[q.frontIndex]
	q.frontIndex++
	q.listSize--

	if q.frontIndex == q.nodeSize {
		q.frontNode, q.frontIndex = q.frontNode.next, 0
	}

	return val, true
}

// Peek returns the next value in queue without removing it from the queue.
func (q *arrayQueue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.frontNode.block[q.frontIndex], true
}

// Contains returns true if a given value is already in the queue.
func (q *arrayQueue[T]) Contains(val T) bool {
	n, i := q.frontNode, q.frontIndex

	for n != nil && (n != q.rearNode || i <= q.rearIndex) {
		if q.equal(n.block[i], val) {
			return true
		}

		if i++; i == q.nodeSize {
			n = n.next
			i = 0
		}
	}

	return false
}
