// Package heap implements heap (a.k.a. priority queue) data structures.
package heap

// The CompareFunc type is a function for comparing two values of the same type.
type CompareFunc func(interface{}, interface{}) int

// Heap represents a heap (priority queue) abstract data type.
type Heap interface {
	Size() int
	IsEmpty() bool
	Insert(interface{}, interface{})
	Delete() (interface{}, interface{})
	Peek() (interface{}, interface{})
	ContainsKey(interface{}) bool
	ContainsValue(interface{}) bool
}
