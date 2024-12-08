package heap

import (
	"fmt"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
)

// binary implements a Binary heap.
type binary[K, V any] struct {
	cmpKey CompareFunc[K]
	eqVal  EqualFunc[V]

	n    int // number of items on heap
	keys []K // binary heap of keys (priorities) using 1-based indexing
	vals []V // binary heap of values using 1-based indexing
}

// NewBinary creates a new Binary heap that can be used as a priority queue.
// The heap size will be automatically increased or decreased as needed.
//
// size is the initial size of the heap (priority queue).
// cmpKey is a function for comparing two keys.
// eqVal is a function for checking the equality of two values.
func NewBinary[K, V any](size int, cmpKey CompareFunc[K], eqVal EqualFunc[V]) Heap[K, V] {
	return &binary[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,

		n:    0,
		keys: make([]K, size+1),
		vals: make([]V, size+1),
	}
}

func (h *binary[K, V]) resize(size int) {
	newKeys := make([]K, size)
	newVals := make([]V, size)

	copy(newKeys, h.keys)
	copy(newVals, h.vals)

	h.keys = newKeys
	h.vals = newVals
}

// Size returns the number of items on the heap.
func (h *binary[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if the heap is empty.
func (h *binary[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to the heap.
func (h *binary[K, V]) Insert(key K, val V) {
	if h.n == len(h.keys)-1 {
		h.resize(len(h.keys) * 2)
	}

	h.n++

	// swim/promotion
	// exchange k with its parent (k/2) until heap is restored.
	var k int
	for k = h.n; k > 1 && h.cmpKey(h.keys[k/2], key) > 0; k /= 2 {
		h.keys[k] = h.keys[k/2]
		h.vals[k] = h.vals[k/2]
	}

	h.keys[k] = key
	h.vals[k] = val
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *binary[K, V]) Delete() (K, V, bool) {
	var zeroK K
	var zeroV V

	if h.n == 0 {
		return zeroK, zeroV, false
	}

	extKey := h.keys[1]
	extVal := h.vals[1]
	key := h.keys[h.n]
	val := h.vals[h.n]

	h.n--

	// sink/demotion
	// exchange k with its smallest/largest child (j) until heap is restored.
	var k, j int
	for k, j = 1, 2; j <= h.n; k, j = j, 2*j {
		if j < h.n && h.cmpKey(h.keys[j], h.keys[j+1]) > 0 {
			j++
		}
		if h.cmpKey(key, h.keys[j]) <= 0 {
			break
		}
		h.keys[k] = h.keys[j]
		h.vals[k] = h.vals[j]
	}

	h.keys[k] = key
	h.vals[k] = val

	// remove stale references to help with garbage collection
	h.keys[h.n+1] = zeroK
	h.vals[h.n+1] = zeroV

	if h.n < len(h.keys)/4 {
		h.resize(len(h.keys) / 2)
	}

	return extKey, extVal, true
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *binary[K, V]) Peek() (K, V, bool) {
	if h.n == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.keys[1], h.vals[1], true
}

// ContainsKey returns true if the given key is on the heap.
func (h *binary[K, V]) ContainsKey(key K) bool {
	for k := 1; k <= h.n; k++ {
		if h.cmpKey(h.keys[k], key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *binary[K, V]) ContainsValue(val V) bool {
	for k := 1; k <= h.n; k++ {
		if h.eqVal(h.vals[k], val) {
			return true
		}
	}

	return false
}

// Graphviz returns a visualization of the heap in Graphviz format.
func (h *binary[K, V]) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "Binary Heap", "", "", "", graphviz.ShapeMrecord)

	for k := 1; k <= h.n; k++ {
		name := fmt.Sprintf("%d", k)

		rec := graphviz.NewRecord(
			graphviz.NewSimpleField("", fmt.Sprintf("%v", h.keys[k])),
			graphviz.NewSimpleField("", fmt.Sprintf("%v", h.vals[k])),
		)

		graph.AddNode(graphviz.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if l := 2 * k; l <= h.n {
			left := fmt.Sprintf("%d", l)
			graph.AddEdge(graphviz.NewEdge(name, left, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if r := 2*k + 1; r <= h.n {
			right := fmt.Sprintf("%d", r)
			graph.AddEdge(graphviz.NewEdge(name, right, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
