// Package heap implements heap data structures.
//
// Heaps are also known as priority queues.
package heap

// The CompareFunc type is a function for comparing two values of the same type.
type CompareFunc func(interface{}, interface{}) int

// Heap represents a heap (priority queue) abstract data type.
type Heap interface {
	Size() int
	IsEmpty() bool
	Insert(interface{}, interface{})
	Delete() (interface{}, interface{}, bool)
	Peek() (interface{}, interface{}, bool)
	ContainsKey(interface{}) bool
	ContainsValue(interface{}) bool
}

// IndexHeap TODO:
type IndexHeap interface {
	Size() int
	IsEmpty() bool
	Insert(int, interface{}, interface{})
	ChangeKey(int, interface{})
	Delete() (int, interface{}, interface{}, bool)
	DeleteIndex(int) (interface{}, interface{}, bool)
	Peek() (int, interface{}, interface{}, bool)
	PeekIndex(int) (interface{}, interface{}, bool)
	ContainsIndex(int) bool
	ContainsKey(interface{}) bool
	ContainsValue(interface{}) bool
}
