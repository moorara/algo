// Package heap implements heap data structures.
//
// Heaps are also known as priority queues.
package heap

// Heap represents a heap (priority queue) abstract data type.
type Heap[K, V any] interface {
	Size() int
	IsEmpty() bool
	Insert(K, V)
	Delete() (K, V, bool)
	Peek() (K, V, bool)
	ContainsKey(K) bool
	ContainsValue(V) bool
}

// IndexHeap represents an indexed heap (priority queue) abstract data type.
type IndexHeap[K, V any] interface {
	Size() int
	IsEmpty() bool
	Insert(int, K, V)
	ChangeKey(int, K)
	Delete() (int, K, V, bool)
	DeleteIndex(int) (K, V, bool)
	Peek() (int, K, V, bool)
	PeekIndex(int) (K, V, bool)
	ContainsIndex(int) bool
	ContainsKey(K) bool
	ContainsValue(V) bool
}
