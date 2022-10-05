package heap

import "github.com/moorara/algo/generic"

type maxHeap[K, V any] struct {
	cmpKey generic.CompareFunc[K]
	eqVal  generic.EqualFunc[V]

	n    int // no. of items on heap
	keys []K // binary heap of keys (priorities) using 1-based indexing
	vals []V // binary heap of values using 1-based indexing
}

// NewMaxHeap creates a new maximum heap (priority queue).
// The heap (priority queue) will be expanded or shrunk when needed.
//
// size is the initial size of heap (priority queue).
// cmpKey is a function for comparing and ordering keys.
// eqVal is a function for checking equality of values.
func NewMaxHeap[K, V any](size int, cmpKey generic.CompareFunc[K], eqVal generic.EqualFunc[V]) Heap[K, V] {
	return &maxHeap[K, V]{
		cmpKey: cmpKey,
		eqVal:  eqVal,

		n:    0,
		keys: make([]K, size+1),
		vals: make([]V, size+1),
	}
}

func (h *maxHeap[K, V]) resize(size int) {
	newKeys := make([]K, size)
	newVals := make([]V, size)

	copy(newKeys, h.keys)
	copy(newVals, h.vals)

	h.keys = newKeys
	h.vals = newVals
}

// Size returns the number of items on heap.
func (h *maxHeap[K, V]) Size() int {
	return h.n
}

// IsEmpty returns true if heap is empty.
func (h *maxHeap[K, V]) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to heap.
func (h *maxHeap[K, V]) Insert(key K, val V) {
	if h.n == len(h.keys)-1 {
		h.resize(len(h.keys) * 2)
	}

	h.n++
	h.keys[h.n] = key
	h.vals[h.n] = val

	// swim/promotion
	// exchange k with its parent (k/2) until heap is restored.
	var k int
	for k = h.n; k > 1 && h.cmpKey(h.keys[k/2], key) < 0; k /= 2 {
		h.keys[k] = h.keys[k/2]
		h.vals[k] = h.vals[k/2]
	}

	h.keys[k] = key
	h.vals[k] = val
}

// Delete removes the maximum key with its value from heap.
// If heap is empty, the last return value will be false.
func (h *maxHeap[K, V]) Delete() (K, V, bool) {
	var zeroK K
	var zeroV V

	if h.n == 0 {
		return zeroK, zeroV, false
	}

	maxKey := h.keys[1]
	maxVal := h.vals[1]
	key := h.keys[h.n]
	val := h.vals[h.n]

	h.n--

	// sink/demotion
	// exchange k with its largest child (j) until heap is restored.
	var k, j int
	for k, j = 1, 2; j <= h.n; k, j = j, 2*j {
		if j < h.n && h.cmpKey(h.keys[j], h.keys[j+1]) < 0 {
			j++
		}
		if h.cmpKey(key, h.keys[j]) >= 0 {
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

	return maxKey, maxVal, true
}

// Peek returns the the maximum key with its value on heap without removing it from heap.
// If heap is empty, the last return value will be false.
func (h *maxHeap[K, V]) Peek() (K, V, bool) {
	if h.n == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	return h.keys[1], h.vals[1], true
}

// ContainsKey returns true if a given key is on heap.
func (h *maxHeap[K, V]) ContainsKey(key K) bool {
	for i := 1; i <= h.n; i++ {
		if h.cmpKey(h.keys[i], key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if a given value is on heap.
func (h *maxHeap[K, V]) ContainsValue(val V) bool {
	for i := 1; i <= h.n; i++ {
		if h.eqVal(h.vals[i], val) {
			return true
		}
	}

	return false
}
