package heap

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

// binary implements an indexed Binary heap.
type indexedBinary[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	cap  int   // maximum number of items on heap
	n    int   // current number of items on heap
	heap []int // binary heap of indices using 1-based indexing
	pos  []int // map of indices to positions on heap
	keys []K   // map of indices to keys (priorities)
	vals []V   // map of indices to values
}

// NewIndexedBinary creates a new indexed Binary heap that can be used as a priority queue.
// An indexed heap (priority queue) associates an index with each key-value pair.
// It allows changing the key (priority) of an index, deleting by index, and looking up by index.
// The size of an indexed Binary heap is fixed.
//
// cap is the maximum number of items on the heap (priority queue).
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewIndexedBinary[K, V any](cap int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) IndexedHeap[K, V] {
	pos := make([]int, cap)
	for i := range pos {
		pos[i] = -1
	}

	return &indexedBinary[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,

		cap:  cap,
		n:    0,
		heap: make([]int, cap+1),
		pos:  pos,
		keys: make([]K, cap),
		vals: make([]V, cap),
	}
}

func (h *indexedBinary[K, V]) validateIndex(i int) {
	if i < 0 || i >= h.cap {
		panic("index is out of range")
	}
}

// compare compares two keys on the heap by their positions.
func (h *indexedBinary[K, V]) compare(a, b int) int {
	i, j := h.heap[a], h.heap[b]
	return h.cmpKey(h.keys[i], h.keys[j])
}

// Promotion operation in heap (a.k.a. swim).
// Exchange k with its parent (k/2) until heap is restored.
func (h *indexedBinary[K, V]) promote(k int) {
	for k > 1 && h.compare(k/2, k) > 0 {
		h.heap[k], h.heap[k/2] = h.heap[k/2], h.heap[k]
		h.pos[h.heap[k]] = k
		h.pos[h.heap[k/2]] = k / 2
		k /= 2
	}
}

// Demotion operation in heap (a.k.a. sink).
// Exchange k with its smallest/largest child (j) until heap is restored.
func (h *indexedBinary[K, V]) demote(k int) {
	for 2*k <= h.n {
		j := 2 * k
		if j < h.n && h.compare(j, j+1) > 0 {
			j++
		}

		if h.compare(k, j) <= 0 {
			break
		}

		h.heap[k], h.heap[j] = h.heap[j], h.heap[k]
		h.pos[h.heap[k]] = k
		h.pos[h.heap[j]] = j

		k = j
	}
}

// Size returns the number of items on the heap.
func (h *indexedBinary[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *indexedBinary[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to the heap with an associated index.
func (h *indexedBinary[K, V]) Insert(i int, key K, val V) {
	// ContainsIndex validates the index as well
	if h.ContainsIndex(i) {
		panic("index already on heap")
	}

	h.n++
	h.heap[h.n] = i
	h.pos[i] = h.n
	h.keys[i] = key
	h.vals[i] = val
	h.promote(h.n)
}

// ChangeKey changes the key associated with an index.
func (h *indexedBinary[K, V]) ChangeKey(i int, key K) {
	// ContainsIndex validates the index as well
	if !h.ContainsIndex(i) {
		panic("index is not on heap")
	}

	h.keys[i] = key
	h.promote(h.pos[i])
	h.demote(h.pos[i])
}

// Delete removes the extremum (minimum or maximum) key with its value and index on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedBinary[K, V]) Delete() (int, K, V, bool) {
	var zeroK K
	var zeroV V

	if h.n == 0 {
		return -1, zeroK, zeroV, false
	}

	i := h.heap[1]
	key := h.keys[i]
	val := h.vals[i]

	h.heap[1], h.heap[h.n] = h.heap[h.n], h.heap[1]
	h.n--
	h.demote(1)

	// delete index
	h.pos[i] = -1

	// remove stale references to help with garbage collection
	h.keys[i] = zeroK
	h.vals[i] = zeroV

	return i, key, val, true
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinary[K, V]) DeleteIndex(i int) (K, V, bool) {
	var zeroK K
	var zeroV V

	// ContainsIndex validates the index as well
	if !h.ContainsIndex(i) {
		return zeroK, zeroV, false
	}

	k := h.pos[i]
	key := h.keys[i]
	val := h.vals[i]

	h.heap[k], h.heap[h.n] = h.heap[h.n], h.heap[k]
	h.n--
	h.promote(k)
	h.demote(k)

	// delete index
	h.pos[i] = -1

	// remove stale references to help with garbage collection
	h.keys[i] = zeroK
	h.vals[i] = zeroV

	return key, val, true
}

// Peek returns the extremum (minimum or maximum) key with its value and index on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *indexedBinary[K, V]) Peek() (int, K, V, bool) {
	if h.n == 0 {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	i := h.heap[1]
	return i, h.keys[i], h.vals[i], true
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinary[K, V]) PeekIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index as well
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.keys[i], h.vals[i], true
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedBinary[K, V]) ContainsIndex(i int) bool {
	h.validateIndex(i)
	return h.pos[i] != -1
}

// ContainsKey returns true if a given key is on the heap.
func (h *indexedBinary[K, V]) ContainsKey(key K) bool {
	for k := 0; k < h.n; k++ {
		if h.cmpKey(h.keys[k], key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if a given value is on the heap.
func (h *indexedBinary[K, V]) ContainsValue(val V) bool {
	for k := 0; k < h.n; k++ {
		if h.eqVal(h.vals[k], val) {
			return true
		}
	}

	return false
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *indexedBinary[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Indexed Binary Heap", "", "", "", graphviz.ShapeMrecord)

	for k := 1; k <= h.n; k++ {
		i := h.heap[k]
		name := fmt.Sprintf("%d", i)

		rec := graphviz.NewRecord(
			graphviz.NewComplexField(
				graphviz.NewRecord(
					graphviz.NewSimpleField("", fmt.Sprintf("%v", i)),
					graphviz.NewComplexField(
						graphviz.NewRecord(
							graphviz.NewSimpleField("", fmt.Sprintf("%v", h.keys[i])),
							graphviz.NewSimpleField("", fmt.Sprintf("%v", h.vals[i])),
						),
					),
				),
			),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if l := 2 * k; l <= h.n {
			left := fmt.Sprintf("%d", h.heap[l])
			graph.AddEdge(graphviz.NewEdge(name, left, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if r := 2*k + 1; r <= h.n {
			right := fmt.Sprintf("%d", h.heap[r])
			graph.AddEdge(graphviz.NewEdge(name, right, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
