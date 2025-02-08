package heap

import (
	"fmt"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

// binary implements an indexed binary heap tree.
type indexedBinary[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	n    int                       // current number of items on heap
	heap []int                     // binary heap of indices using 1-based indexing
	pos  []int                     // map of indices to positions on heap
	kvs  []*generic.KeyValue[K, V] // map of indices to key-values
}

// NewIndexedBinary creates a new indexed binary heap that can be used as a priority queue.
//
// An indexed heap (priority queue) associates an index with each key-value pair.
// It allows changing the key (priority) of an index, deleting by index, and looking up by index.
// The size of an indexed binary heap is fixed.
//
// Parameters:
//
//   - cap is the maximum number of items on the heap.
//   - cmpKey is a function for comparing two keys.
//   - eqVal is a function for checking the equality of two values.
func NewIndexedBinary[K, V any](cap int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) IndexedHeap[K, V] {
	pos := make([]int, cap)
	for i := range pos {
		pos[i] = -1
	}

	return &indexedBinary[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		n:      0,
		heap:   make([]int, cap+1),
		pos:    pos,
		kvs:    make([]*generic.KeyValue[K, V], cap),
	}
}

// nolint: unused
// This method verifies the integrity of an indexed binary heap.
func (h *indexedBinary[K, V]) verify() bool {
	// Verify the heap is a complete tree.
	for i := 1; i <= h.n; i++ {
		if j := h.heap[i]; h.pos[j] == -1 || h.kvs[j] == nil {
			return false
		}
	}

	// Verify the heap property (heap order).
	for k := 1; k <= h.n; k++ {
		if l := 2 * k; l <= h.n {
			if h.compare(k, l) > 0 {
				return false
			}
		}

		if r := 2*k + 1; r <= h.n {
			if h.compare(k, r) > 0 {
				return false
			}
		}
	}

	return true
}

// compare compares two keys on the heap by their positions.
func (h *indexedBinary[K, V]) compare(a, b int) int {
	i, j := h.heap[a], h.heap[b]
	return h.cmpKey(h.kvs[i].Key, h.kvs[j].Key)
}

// swap exchanges the elements at indices i and j in the heap array
// and updates their corresponding positions in the pos map.
//
// This method is defined on the indexedBinomial struct to prevent name clashes
// with other similar implementations in this package.
func (h *indexedBinary[K, V]) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
	h.pos[h.heap[i]], h.pos[h.heap[j]] = i, j
}

// Promotion operation in heap (a.k.a. swim).
// Exchange child k with its parent k/2 until the heap is restored.
func (h *indexedBinary[K, V]) promote(k int) {
	for ; k > 1 && h.compare(k/2, k) > 0; k /= 2 {
		h.swap(k, k/2)
	}
}

// Demotion operation in heap (a.k.a. sink).
// Exchange parent k with its smallest/largest child j until the heap is restored.
func (h *indexedBinary[K, V]) demote(k int) {
	for j := 2 * k; j <= h.n; k, j = j, 2*j {
		if j < h.n && h.compare(j+1, j) < 0 {
			j++
		}
		if h.compare(k, j) < 0 {
			break
		}
		h.swap(k, j)
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
func (h *indexedBinary[K, V]) Insert(i int, key K, val V) bool {
	// ContainsIndex validates the index too.
	if h.ContainsIndex(i) {
		return false
	}

	h.n++
	h.heap[h.n] = i
	h.pos[i] = h.n
	h.kvs[i] = &generic.KeyValue[K, V]{
		Key: key,
		Val: val,
	}

	h.promote(h.n)

	return true
}

// ChangeKey changes the key associated with an index.
func (h *indexedBinary[K, V]) ChangeKey(i int, key K) bool {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		return false
	}

	h.kvs[i].Key = key
	h.promote(h.pos[i])
	h.demote(h.pos[i])

	return true
}

// Delete removes the extremum (minimum or maximum) key with its value and index on the heap.
// If the heap is empty, the second return value will be false.
func (h *indexedBinary[K, V]) Delete() (int, K, V, bool) {
	if h.n == 0 {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	i := h.heap[1]
	ext := h.kvs[i]
	h.swap(1, h.n)
	h.n--

	h.demote(1)

	// Delete index and remove stale reference to help with garbage collection.
	h.pos[i] = -1
	h.kvs[i] = nil

	return i, ext.Key, ext.Val, true
}

// DeleteIndex removes a key-value pair and its associated index from the heap.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinary[K, V]) DeleteIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	k := h.pos[i]
	kv := h.kvs[i]
	h.swap(k, h.n)
	h.n--

	h.promote(k)
	h.demote(k)

	// Delete index and remove stale reference to help with garbage collection.
	h.pos[i] = -1
	h.kvs[i] = nil

	return kv.Key, kv.Val, true
}

// DeleteAll deletes all keys with their values and indices on the heap, leaving it empty.
func (h *indexedBinary[K, V]) DeleteAll() {
	h.n = 0
	h.heap = make([]int, len(h.heap))
	h.pos = make([]int, len(h.pos))
	h.kvs = make([]*generic.KeyValue[K, V], len(h.kvs))

	for i := range h.pos {
		h.pos[i] = -1
	}
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
	return i, h.kvs[i].Key, h.kvs[i].Val, true
}

// PeekIndex returns a key-value pair on the heap by its associated index without removing it.
// If the index is not valid or not on the heap, the second return value will be false.
func (h *indexedBinary[K, V]) PeekIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index too.
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.kvs[i].Key, h.kvs[i].Val, true
}

// ContainsIndex returns true if a given index is on the heap.
func (h *indexedBinary[K, V]) ContainsIndex(i int) bool {
	return 0 <= i && i < len(h.kvs) && h.pos[i] != -1
}

// ContainsKey returns true if a given key is on the heap.
func (h *indexedBinary[K, V]) ContainsKey(key K) bool {
	for i := 0; i < h.n; i++ {
		if h.kvs[i] != nil && h.cmpKey(h.kvs[i].Key, key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if a given value is on the heap.
func (h *indexedBinary[K, V]) ContainsValue(val V) bool {
	for i := 0; i < h.n; i++ {
		if h.kvs[i] != nil && h.eqVal(h.kvs[i].Val, val) {
			return true
		}
	}

	return false
}

// DOT generates a DOT representation of the heap.
func (h *indexedBinary[K, V]) DOT() string {
	graph := dot.NewGraph(true, true, false, "Indexed Binary Heap", "", "", "", dot.ShapeMrecord)

	for k := 1; k <= h.n; k++ {
		i := h.heap[k]
		name := fmt.Sprintf("%d", i)

		rec := dot.NewRecord(
			dot.NewComplexField(
				dot.NewRecord(
					dot.NewSimpleField("", fmt.Sprintf("%v", i)),
					dot.NewComplexField(
						dot.NewRecord(
							dot.NewSimpleField("", fmt.Sprintf("%v", h.kvs[i].Key)),
							dot.NewSimpleField("", fmt.Sprintf("%v", h.kvs[i].Val)),
						),
					),
				),
			),
		)

		graph.AddNode(dot.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if l := 2 * k; l <= h.n {
			left := fmt.Sprintf("%d", h.heap[l])
			graph.AddEdge(dot.NewEdge(name, left, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if r := 2*k + 1; r <= h.n {
			right := fmt.Sprintf("%d", h.heap[r])
			graph.AddEdge(dot.NewEdge(name, right, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}
	}

	return graph.DOT()
}
