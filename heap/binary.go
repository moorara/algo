package heap

import (
	"fmt"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

// binary implements a binary heap tree.
type binary[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	n    int                       // number of items on heap
	heap []*generic.KeyValue[K, V] // binary heap of key-values using 1-based indexing
}

// NewBinary creates a new binary heap that can be used as a priority queue.
// The heap size will be automatically increased or decreased as needed.
//
// Parameters:
//
//   - size is the initial size of the heap (priority queue).
//   - cmpKey is a function for comparing two keys.
//   - eqVal is a function for checking the equality of two values.
func NewBinary[K, V any](size int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) Heap[K, V] {
	return &binary[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,
		n:      0,
		heap:   make([]*generic.KeyValue[K, V], size+1),
	}
}

// nolint: unused
// This method verifies the integrity of a binary heap.
func (h *binary[K, V]) verify() bool {
	if h.heap[0] != nil {
		return false
	}

	// Verify the heap is a complete tree.
	for i := 1; i <= h.n; i++ {
		if h.heap[i] == nil {
			return false
		}
	}

	// Verify the deleted items are dereferenced.
	for i := h.n + 1; i < len(h.heap); i++ {
		if h.heap[i] != nil {
			return false
		}
	}

	// Verify the heap property (heap order).
	for k := 1; k <= h.n; k++ {
		if l := 2 * k; l <= h.n {
			if h.cmpKey(h.heap[k].Key, h.heap[l].Key) > 0 {
				return false
			}
		}

		if r := 2*k + 1; r <= h.n {
			if h.cmpKey(h.heap[k].Key, h.heap[r].Key) > 0 {
				return false
			}
		}
	}

	return true
}

func (h *binary[K, V]) resize(size int) {
	newH := make([]*generic.KeyValue[K, V], size)
	copy(newH, h.heap)
	h.heap = newH
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
	if h.n == len(h.heap)-1 {
		h.resize(len(h.heap) * 2)
	}

	h.n++

	// Swim/Promotion
	// Exchange child k with its parent k/2 until the heap is restored.
	var k int
	for k = h.n; k > 1 && h.cmpKey(h.heap[k/2].Key, key) > 0; k /= 2 {
		h.heap[k] = h.heap[k/2]
	}

	h.heap[k] = &generic.KeyValue[K, V]{
		Key: key,
		Val: val,
	}
}

// Delete removes the extremum (minimum or maximum) key with its value on the heap.
// If the heap is empty, the second return value will be false.
func (h *binary[K, V]) Delete() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	ext := h.heap[1] // extremum key-value
	kv := h.heap[h.n]
	h.n--

	// Sink/Demotion
	// Exchange parent k with its smallest/largest child j until the heap is restored.
	var k, j int
	for k, j = 1, 2; j <= h.n; k, j = j, 2*j {
		if j < h.n && h.cmpKey(h.heap[j+1].Key, h.heap[j].Key) < 0 {
			j++
		}
		if h.cmpKey(kv.Key, h.heap[j].Key) < 0 {
			break
		}
		h.heap[k] = h.heap[j]
	}

	h.heap[k] = kv

	// Remove stale reference to help with garbage collection.
	h.heap[h.n+1] = nil

	if h.n < len(h.heap)/4 {
		h.resize(len(h.heap) / 2)
	}

	return ext.Key, ext.Val, true
}

// DeleteAll deletes all keys with their values on the heap, leaving it empty.
func (h *binary[K, V]) DeleteAll() {
	h.n = 0
	h.heap = make([]*generic.KeyValue[K, V], len(h.heap))
}

// Peek returns the extremum (minimum or maximum) key with its value on the heap without removing it.
// If the heap is empty, the second return value will be false.
func (h *binary[K, V]) Peek() (K, V, bool) {
	if h.IsEmpty() {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.heap[1].Key, h.heap[1].Val, true
}

// ContainsKey returns true if the given key is on the heap.
func (h *binary[K, V]) ContainsKey(key K) bool {
	for k := 1; k <= h.n; k++ {
		if h.cmpKey(h.heap[k].Key, key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if the given value is on the heap.
func (h *binary[K, V]) ContainsValue(val V) bool {
	for k := 1; k <= h.n; k++ {
		if h.eqVal(h.heap[k].Val, val) {
			return true
		}
	}

	return false
}

// DOT generates a DOT representation of the heap.
func (h *binary[K, V]) DOT() string {
	graph := dot.NewGraph(true, true, false, "Binary Heap", "", "", "", dot.ShapeMrecord)

	for k := 1; k <= h.n; k++ {
		name := fmt.Sprintf("%d", k)

		rec := dot.NewRecord(
			dot.NewSimpleField("", fmt.Sprintf("%v", h.heap[k].Key)),
			dot.NewSimpleField("", fmt.Sprintf("%v", h.heap[k].Val)),
		)

		graph.AddNode(dot.NewNode(name, "", rec.Label(), "", "", "", "", ""))

		if l := 2 * k; l <= h.n {
			left := fmt.Sprintf("%d", l)
			graph.AddEdge(dot.NewEdge(name, left, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		if r := 2*k + 1; r <= h.n {
			right := fmt.Sprintf("%d", r)
			graph.AddEdge(dot.NewEdge(name, right, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}
	}

	return graph.DOT()
}
