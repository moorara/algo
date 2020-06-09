package heap

type maxHeap struct {
	n      int           // No. of items on heap
	keys   []interface{} // binary heap of keys (priorities) using 1-based indexing
	values []interface{} // binary heap of values using 1-based indexing
	cmpKey CompareFunc
	cmpVal CompareFunc
}

// NewMaxHeap creates a new maximum heap (priority queue).
// The heap (priority queue) will be expanded or shrunk when needed.
//
// initialCapacity is the initial capacity of heap (priority queue).
// cmpKey and cmpVal are comparator functions for keys and values respectively.
func NewMaxHeap(initialCapacity int, cmpKey, cmpVal CompareFunc) Heap {
	return &maxHeap{
		n:      0,
		keys:   make([]interface{}, initialCapacity+1),
		values: make([]interface{}, initialCapacity+1),
		cmpKey: cmpKey,
		cmpVal: cmpVal,
	}
}

func (h *maxHeap) resize(newSize int) {
	newKeys := make([]interface{}, newSize)
	newValues := make([]interface{}, newSize)

	copy(newKeys, h.keys)
	copy(newValues, h.values)

	h.keys = newKeys
	h.values = newValues
}

// Size returns the number of items on heap.
func (h *maxHeap) Size() int {
	return h.n
}

// IsEmpty returns true if heap is empty.
func (h *maxHeap) IsEmpty() bool {
	return h.n == 0
}

// Insert adds a new key-value pair to heap.
func (h *maxHeap) Insert(key, value interface{}) {
	if h.n == len(h.keys)-1 {
		h.resize(len(h.keys) * 2)
	}

	h.n++
	h.keys[h.n] = key
	h.values[h.n] = value

	// swim/promotion
	// exchange k with its parent (k/2) until heap is restored.
	var k int
	for k = h.n; k > 1 && h.cmpKey(h.keys[k/2], key) < 0; k /= 2 {
		h.keys[k] = h.keys[k/2]
		h.values[k] = h.values[k/2]
	}

	h.keys[k] = key
	h.values[k] = value
}

// Delete removes the maximum key with its value from heap.
// If heap is empty, the last return value will be false.
func (h *maxHeap) Delete() (interface{}, interface{}, bool) {
	if h.n == 0 {
		return nil, nil, false
	}

	maxKey := h.keys[1]
	maxValue := h.values[1]
	key := h.keys[h.n]
	value := h.values[h.n]

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
		h.values[k] = h.values[j]
	}

	h.keys[k] = key
	h.values[k] = value

	// remove stale references to help with garbage collection
	h.keys[h.n+1] = nil
	h.values[h.n+1] = nil

	if h.n < len(h.keys)/4 {
		h.resize(len(h.keys) / 2)
	}

	return maxKey, maxValue, true
}

// Peek returns the the maximum key with its value on heap without removing it from heap.
// If heap is empty, the last return value will be false.
func (h *maxHeap) Peek() (interface{}, interface{}, bool) {
	if h.n == 0 {
		return nil, nil, false
	}

	return h.keys[1], h.values[1], true
}

// ContainsKey returns true if a given key is on heap.
func (h *maxHeap) ContainsKey(key interface{}) bool {
	for i := 1; i <= h.n; i++ {
		if h.cmpKey(h.keys[i], key) == 0 {
			return true
		}
	}

	return false
}

// ContainsValue returns true if a given value is on heap.
func (h *maxHeap) ContainsValue(value interface{}) bool {
	for i := 1; i <= h.n; i++ {
		if h.cmpVal(h.values[i], value) == 0 {
			return true
		}
	}

	return false
}
