package heap

import (
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/symboltable"
)

// TODO:
type indexedFibonacciNode[K, V any] struct {
	key           K
	val           V
	index         int                         // index associated with the key
	order         int                         // order of the tree rooted at this node
	mark          bool                        // indicates if this node already lost a child
	prev, next    *indexedFibonacciNode[K, V] // siblings of this node
	parent, child *indexedFibonacciNode[K, V] // parent and child of this node
}

// fibonacci implements a Fibonacci heap tree.
type indexedFibonacci[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	cap   int                                                      // maximum number of items on heap
	n     int                                                      // current number of items on heap
	root  *indexedFibonacciNode[K, V]                              // root of the circular root list
	ext   *indexedFibonacciNode[K, V]                              // extremum (minimum or maximum) node of the root list
	nodes *indexedFibonacciNode[K, V]                              // list of indexed nodes
	table symboltable.SymbolTable[int, indexedFibonacciNode[K, V]] // used for the consolidate operation
}

// NewIndexedFibonacci creates a new indexed Fibonacci heap that can be used as a priority queue.
//
// TODO:
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewIndexedFibonacci[K, V any](cap int, cmpKey CompareFunc[K], eqVal EqualFunc[V]) IndexedMergeableHeap[K, V] {
	return &indexedFibonacci[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		cap:    cap,
		n:      0,
	}
}

// Merge merges another heap with the current heap.
func (h *indexedFibonacci[K, V]) Merge(H IndexedMergeableHeap[K, V]) {
	// TODO:
}

// Size returns the number of items on the heap.
func (h *indexedFibonacci[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *indexedFibonacci[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to the heap.
func (h *indexedFibonacci[K, V]) Insert(i int, key K, val V) {
	// TODO:
}

// ChangeKey changes the key associated with an index.
func (h *indexedFibonacci[K, V]) ChangeKey(i int, key K) {
	// TODO:
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedFibonacci[K, V]) Delete() (int, K, V, bool) {
	var zeroK K
	var zeroV V
	return -1, zeroK, zeroV, false
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedFibonacci[K, V]) DeleteIndex(i int) (K, V, bool) {
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// DeleteAll deletes all keys with their values and indices on the heap, leaving it empty.
func (h *indexedFibonacci[K, V]) DeleteAll() {
	// TODO:
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *indexedFibonacci[K, V]) Peek() (int, K, V, bool) {
	var zeroK K
	var zeroV V
	return -1, zeroK, zeroV, false
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedFibonacci[K, V]) PeekIndex(i int) (K, V, bool) {
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedFibonacci[K, V]) ContainsIndex(i int) bool {
	// TODO:
	return false
}

// ContainsKey returns true if the given key is on the heap.
func (h *indexedFibonacci[K, V]) ContainsKey(key K) bool {
	// TODO:
	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *indexedFibonacci[K, V]) ContainsValue(val V) bool {
	// TODO:
	return false
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *indexedFibonacci[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Indexed Fibonacci Heap", "", "", "", graphviz.ShapeMrecord)

	// TODO:

	return graph.DotCode()
}
