package heap

type maxHeap struct {
	last   int
	keys   []interface{}
	values []interface{}
	cmpKey CompareFunc
	cmpVal CompareFunc
}

// NewMaxHeap creates a new max-heap (priority queue).
func NewMaxHeap(initialSize int, cmpKey, cmpVal CompareFunc) Heap {
	return &maxHeap{
		last:   0,
		keys:   make([]interface{}, initialSize),
		values: make([]interface{}, initialSize),
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
	return h.last
}

// IsEmpty returns true if heap is empty.
func (h *maxHeap) IsEmpty() bool {
	return h.last == 0
}

// Insert adds a new key-value pair to heap.
func (h *maxHeap) Insert(key, value interface{}) {
	if h.last == len(h.keys)-1 {
		h.resize(len(h.keys) * 2)
	}

	h.last++
	var i int

	for i = h.last; true; i /= 2 {
		if i == 1 || h.cmpKey(key, h.keys[i/2]) <= 0 {
			break
		}
		h.keys[i] = h.keys[i/2]
		h.values[i] = h.values[i/2]
	}

	h.keys[i] = key
	h.values[i] = value
}

// Delete removes a key-value pair from heap.
func (h *maxHeap) Delete() (interface{}, interface{}) {
	if h.last == 0 {
		return nil, nil
	}

	maxKey := h.keys[1]
	maxValue := h.values[1]
	lastKey := h.keys[h.last]
	lastValue := h.values[h.last]

	h.last--
	var i, j int

	for i, j = 1, 2; j <= h.last; i, j = j, j*2 {
		if j < h.last && h.cmpKey(h.keys[j], h.keys[j+1]) < 0 {
			j++
		}
		if h.cmpKey(lastKey, h.keys[j]) >= 0 {
			break
		}
		h.keys[i] = h.keys[j]
		h.values[i] = h.values[j]
	}

	h.keys[i] = lastKey
	h.values[i] = lastValue

	if h.last < len(h.keys)/4 {
		h.resize(len(h.keys) / 2)
	}

	return maxKey, maxValue
}

// Peek returns the next key-value pair on heap without removing it from heap.
func (h *maxHeap) Peek() (interface{}, interface{}) {
	if h.last == 0 {
		return nil, nil
	}

	return h.keys[1], h.values[1]
}

// Contains returns true if a given key is already on heap.
func (h *maxHeap) ContainsKey(key interface{}) bool {
	for i := 1; i <= h.last; i++ {
		if h.cmpKey(h.keys[i], key) == 0 {
			return true
		}
	}

	return false
}

// Contains returns true if a given value is already on heap.
func (h *maxHeap) ContainsValue(value interface{}) bool {
	for i := 1; i <= h.last; i++ {
		if h.cmpVal(h.values[i], value) == 0 {
			return true
		}
	}

	return false
}
