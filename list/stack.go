package list

import . "github.com/moorara/algo/generic"

// Stack represents a stack abstract data type.
type Stack[T any] interface {
	Size() int
	IsEmpty() bool
	Push(T)
	Pop() (T, bool)
	Peek() (T, bool)
	Contains(T) bool
}

type arrayStack[T any] struct {
	nodeSize int
	equal    EqualFunc[T]

	listSize int
	topIndex int
	topNode  *arrayNode[T]
}

// NewStack creates a new array-list stack.
func NewStack[T any](nodeSize int, equal EqualFunc[T]) Stack[T] {
	return &arrayStack[T]{
		nodeSize: nodeSize,
		equal:    equal,

		listSize: 0,
		topIndex: -1,
		topNode:  nil,
	}
}

// Size returns the number of values on the stack.
func (s *arrayStack[T]) Size() int {
	return s.listSize
}

// IsEmpty returns true if the stack is empty.
func (s *arrayStack[T]) IsEmpty() bool {
	return s.listSize == 0
}

// Enqueue adds a new value to the stack.
func (s *arrayStack[T]) Push(val T) {
	s.listSize++
	s.topIndex++

	if s.topNode == nil {
		s.topNode = newArrayNode[T](s.nodeSize, nil)
	} else if s.topIndex == s.nodeSize {
		s.topNode = newArrayNode[T](s.nodeSize, s.topNode)
		s.topIndex = 0
	}

	s.topNode.block[s.topIndex] = val
}

// Dequeue removes a value from the stack.
func (s *arrayStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	val := s.topNode.block[s.topIndex]
	s.topIndex--
	s.listSize--

	if s.topIndex == -1 {
		s.topNode = s.topNode.next
		if s.topNode != nil {
			s.topIndex = s.nodeSize - 1
		}
	}

	return val, true
}

// Peek returns the next value on stack without removing it from the stack.
func (s *arrayStack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.topNode.block[s.topIndex], true
}

// Contains returns true if a given value is already on the stack.
func (s *arrayStack[T]) Contains(val T) bool {
	n := s.topNode
	i := s.topIndex

	for n != nil {
		if s.equal(n.block[i], val) {
			return true
		}

		if i--; i < 0 {
			n = n.next
			i = s.nodeSize - 1
		}
	}

	return false
}
