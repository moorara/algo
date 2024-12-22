package heap

import (
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/symboltable"
)

// TODO:
type fibonacciNode[K, V any] struct {
	key        K
	val        V
	order      int                  // order of the tree rooted at this node
	prev, next *fibonacciNode[K, V] // siblings of this node
	child      *fibonacciNode[K, V] // child of this node
}

// fibonacci implements a Fibonacci heap tree.
type fibonacci[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	n     int                                               // number of items on heap
	root  *fibonacciNode[K, V]                              // root of the circular root list
	ext   *fibonacciNode[K, V]                              // extremum (minimum or maximum) node of the root list
	table symboltable.SymbolTable[int, fibonacciNode[K, V]] // used for the consolidate operation
}

// NewFibonacci creates a new Fibonacci heap that can be used as a priority queue.
//
// TODO:
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewFibonacci[K, V any](cmpKey CompareFunc[K], eqVal EqualFunc[V]) MergeableHeap[K, V] {
	return &fibonacci[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
	}
}

// Merge merges another heap with the current heap.
func (h *fibonacci[K, V]) Merge(H MergeableHeap[K, V]) {
	// TODO:
}

// Size returns the number of items on the heap.
func (h *fibonacci[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *fibonacci[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to the heap.
func (h *fibonacci[K, V]) Insert(key K, val V) {
	// TODO:
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *fibonacci[K, V]) Delete() (K, V, bool) {
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *fibonacci[K, V]) Peek() (K, V, bool) {
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

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *fibonacci[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Fibonacci Heap", "", "", "", graphviz.ShapeMrecord)

	// TODO:

	return graph.DotCode()
}
