package heap

import (
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

// TODO:
type indexedBinomialNode[K, V any] struct {
	key            K
	val            V
	index          int // index associated with the key
	order          int // order of the tree rooted at this node
	parent         *indexedBinomialNode[K, V]
	child, sibling *indexedBinomialNode[K, V]
}

// indexedBinomial implements an indexed binomial heap tree.
type indexedBinomial[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	cap   int                        // maximum number of items on heap
	n     int                        // current number of items on heap
	root  *indexedBinomialNode[K, V] // root of the root list
	nodes *indexedBinomialNode[K, V] // list of indexed nodes
}

// NewIndexedBinomial creates a new indexed binomial heap that can be used as a priority queue.
//
// TODO:
//
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewIndexedBinomial[K, V any](cap int, cmpKey CompareFunc[K], eqVal EqualFunc[V]) IndexedMergeableHeap[K, V] {
	return &indexedBinomial[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		cap:    cap,
		n:      0,
	}
}

// merge merges another heap with the current heap.
func (h *indexedBinomial[K, V]) Merge(H IndexedMergeableHeap[K, V]) {
	// TODO:
}

// Size returns the number of items on the heap.
func (h *indexedBinomial[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *indexedBinomial[K, V]) IsEmpty() bool {
	return h.root == nil
}

// Insert adds a new key-value pair to the heap.
func (h *indexedBinomial[K, V]) Insert(i int, key K, val V) {
	// TODO:
}

// ChangeKey changes the key associated with an index.
func (h *indexedBinomial[K, V]) ChangeKey(i int, key K) {
	// TODO:
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedBinomial[K, V]) Delete() (int, K, V, bool) {
	var zeroK K
	var zeroV V
	return -1, zeroK, zeroV, false
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinomial[K, V]) DeleteIndex(i int) (K, V, bool) {
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *indexedBinomial[K, V]) Peek() (int, K, V, bool) {
	var zeroK K
	var zeroV V
	return -1, zeroK, zeroV, false
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinomial[K, V]) PeekIndex(i int) (K, V, bool) {
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedBinomial[K, V]) ContainsIndex(i int) bool {
	// TODO:
	return false
}

// ContainsKey returns true if the given key is on the heap.
func (h *indexedBinomial[K, V]) ContainsKey(key K) bool {
	// TODO:
	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *indexedBinomial[K, V]) ContainsValue(val V) bool {
	// TODO:
	return false
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *indexedBinomial[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Indexed Binomial Heap", "", "", "", graphviz.ShapeMrecord)

	// TODO:

	return graph.DotCode()
}
