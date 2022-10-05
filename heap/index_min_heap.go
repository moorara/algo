package heap

import "github.com/moorara/algo/generic"

type indexMinHeap[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	cap  int   // maximum number of items on heap
	n    int   // current number of items on heap
	heap []int // binary heap of indices using 1-based indexing
	pos  []int // map of indices to positions on heap
	keys []K   // map of indices to keys (priorities)
	vals []V   // map of indices to values
}

// NewIndexMinHeap creates a new indexed minimum heap (priority queue).
// An indexed minimum heap (priority queue) associates an index to each key-value pair.
// It allows changing the key (priority) of an index, deleting by index, and looking up by index.
// The size of an indexed heap (priority queue) is fixed.
//
// cap is the maximum number of items on heap (priority queue).
// cmpKey is a function for comparing and ordering keys.
// eqVal is a function for checking equality of values.
func NewIndexMinHeap[K, V any](cap int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) IndexHeap[K, V] {
	pos := make([]int, cap)
	for i := range pos {
		pos[i] = -1
	}

	return &indexMinHeap[K, V]{
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

func (h *indexMinHeap[K, V]) validateIndex(i int) {
	if i < 0 || i >= h.cap {
		panic("index is out of range")
	}
}

// compare compares two positions (nodes) in heap.
func (h *indexMinHeap[K, V]) compare(a, b int) int {
	i, j := h.heap[a], h.heap[b]
	return h.cmpKey(h.keys[i], h.keys[j])
}

// Promotion operation in heap (a.k.a. swim).
// Exchange k with its parent (k/2) until heap is restored.
func (h *indexMinHeap[K, V]) promote(k int) {
	for k > 1 && h.compare(k/2, k) > 0 {
		h.heap[k], h.heap[k/2] = h.heap[k/2], h.heap[k]
		h.pos[h.heap[k]] = k
		h.pos[h.heap[k/2]] = k / 2
		k /= 2
	}
}

// Demotion operation in heap (a.k.a. sink).
// Exchange k with its smallest child (j) until heap is restored.
func (h *indexMinHeap[K, V]) demote(k int) {
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

// Size returns the number of items on heap.
func (h *indexMinHeap[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if heap is empty.
func (h *indexMinHeap[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to heap with an associated index.
func (h *indexMinHeap[K, V]) Insert(i int, key K, val V) {
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
func (h *indexMinHeap[K, V]) ChangeKey(i int, key K) {
	// ContainsIndex validates the index as well
	if !h.ContainsIndex(i) {
		panic("index is not on heap")
	}

	h.keys[i] = key
	h.promote(h.pos[i])
	h.demote(h.pos[i])
}

// Delete removes the minimum key with its value and index from heap.
// If heap is empty, the last return value will be false.
func (h *indexMinHeap[K, V]) Delete() (int, K, V, bool) {
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

// DeleteIndex removes a key-value pair and its associated index from heap.
// If the index is not valid or not on heap, the last return value will be false.
func (h *indexMinHeap[K, V]) DeleteIndex(i int) (K, V, bool) {
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

// Peek returns the the minimum key with its value and index on heap without removing it from heap.
// If heap is empty, the last return value will be false.
func (h *indexMinHeap[K, V]) Peek() (int, K, V, bool) {
	if h.n == 0 {
		var zeroK K
		var zeroV V
		return -1, zeroK, zeroV, false
	}

	i := h.heap[1]
	return i, h.keys[i], h.vals[i], true
}

// PeekIndex returns a key-value pair on heap by its associated index without removing it from heap.
// If the index is not valid or not on heap, the last return value will be false.
func (h *indexMinHeap[K, V]) PeekIndex(i int) (K, V, bool) {
	// ContainsIndex validates the index as well
	if !h.ContainsIndex(i) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.keys[i], h.vals[i], true
}

// ContainsIndex returns true if a given index is on heap.
func (h *indexMinHeap[K, V]) ContainsIndex(i int) bool {
	h.validateIndex(i)
	return h.pos[i] != -1
}

// ContainsKey returns true if a given key is on heap.
func (h *indexMinHeap[K, V]) ContainsKey(key K) bool {
	for i := 0; i < h.n; i++ {
		if h.cmpKey(h.keys[i], key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if a given value is on heap.
func (h *indexMinHeap[K, V]) ContainsValue(val V) bool {
	for i := 0; i < h.n; i++ {
		if h.eqVal(h.vals[i], val) {
			return true
		}
	}

	return false
}
