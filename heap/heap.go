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
	DeleteAll()
	Peek() (K, V, bool)
	ContainsKey(K) bool
	ContainsValue(V) bool
	Graphviz() string
}

// IndexedHeap represents an indexed heap (priority queue) abstract data type.
type IndexedHeap[K, V any] interface {
	Size() int
	IsEmpty() bool
	Insert(int, K, V)
	ChangeKey(int, K)
	Delete() (int, K, V, bool)
	DeleteIndex(int) (K, V, bool)
	DeleteAll()
	Peek() (int, K, V, bool)
	PeekIndex(int) (K, V, bool)
	ContainsIndex(int) bool
	ContainsKey(K) bool
	ContainsValue(V) bool
	Graphviz() string
}

// MergeableHeap represents a mergeable heap (priority queue) abstract data type.
type MergeableHeap[K, V any] interface {
	Heap[K, V]
	Merge(MergeableHeap[K, V])
}

// IndexedMergeableHeap represents a mergeable indexed heap (priority queue) abstract data type.
type IndexedMergeableHeap[K, V any] interface {
	IndexedHeap[K, V]
	Merge(IndexedMergeableHeap[K, V])
}
