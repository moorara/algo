package heap

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

type fibonacciNode[K, V any] struct {
	key        K
	val        V
	order      int                  // order of the tree rooted at this node
	prev, next *fibonacciNode[K, V] // siblings of this node
	child      *fibonacciNode[K, V] // child of this node
}

// fibonacci implements a Fibonacci heap.
type fibonacci[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]
}

// NewFibonacci creates a new Fibonacci heap that can be used as a priority queue.
//
// ...
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewFibonacci[K, V any](cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) Heap[K, V] {
	return &fibonacci[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
	}
}

// Size returns the number of items on the heap.
func (h *fibonacci[K, V]) Size() int {
	// TODO:
	return 0
}

// IsEmpty returns true if the heap is empty.
func (h *fibonacci[K, V]) IsEmpty() bool {
	// TODO:
	return false
}

// Insert adds a new key-value pair to the heap.
func (h *fibonacci[K, V]) Insert(key K, val V) {
	// TODO:
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *fibonacci[K, V]) Delete() (K, V, bool) {
	// TODO:
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *fibonacci[K, V]) Peek() (K, V, bool) {
	// TODO:
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// ContainsKey returns true if the given key is on the heap.
func (h *fibonacci[K, V]) ContainsKey(key K) bool {
	// TODO:
	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *fibonacci[K, V]) ContainsValue(val V) bool {
	// TODO:
	return false
}

// Graphviz returns a visualization of ... in Graphviz format.
func (h *fibonacci[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Fibonacci Heap", "", "", "", graphviz.ShapeOval)

	// TODO:

	return graph.DotCode()
}
