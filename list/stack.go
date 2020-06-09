package list

import "github.com/moorara/algo/compare"

// Stack represents a stack abstract data type.
type Stack interface {
	Size() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	Contains(interface{}, compare.Func) bool
}

type arrayStack struct {
	listSize  int
	nodeSize  int
	nodeIndex int
	topNode   *arrayNode
}

// NewStack creates a new array-list stack.
func NewStack(nodeSize int) Stack {
	return &arrayStack{
		listSize:  0,
		nodeSize:  nodeSize,
		nodeIndex: -1,
		topNode:   nil,
	}
}

// Size returns the number of items on stack.
func (s *arrayStack) Size() int {
	return s.listSize
}

// IsEmpty returns true if stack is empty.
func (s *arrayStack) IsEmpty() bool {
	return s.listSize == 0
}

// Enqueue adds a new item to stack.
func (s *arrayStack) Push(item interface{}) {
	s.listSize++
	s.nodeIndex++

	if s.topNode == nil {
		s.topNode = newArrayNode(s.nodeSize, nil)
	} else {
		if s.nodeIndex == s.nodeSize {
			s.nodeIndex = 0
			s.topNode = newArrayNode(s.nodeSize, s.topNode)
		}
	}

	s.topNode.block[s.nodeIndex] = item
}

// Dequeue removes an item from stack.
func (s *arrayStack) Pop() interface{} {
	if s.listSize == 0 {
		return nil
	}

	item := s.topNode.block[s.nodeIndex]
	s.nodeIndex--
	s.listSize--

	if s.nodeIndex == -1 {
		s.topNode = s.topNode.next
		if s.topNode != nil {
			s.nodeIndex = s.nodeSize - 1
		}
	}

	return item
}

// Peek returns the next item on stack without removing it from stack.
func (s *arrayStack) Peek() interface{} {
	if s.listSize == 0 {
		return nil
	}

	return s.topNode.block[s.nodeIndex]
}

// Contains returns true if a given item is already on stack.
func (s *arrayStack) Contains(item interface{}, cmp compare.Func) bool {
	n := s.topNode
	i := s.nodeIndex

	for n != nil {
		if cmp(n.block[i], item) == 0 {
			return true
		}

		i--
		if i < 0 {
			n = n.next
			i = s.nodeSize - 1
		}
	}

	return false
}
