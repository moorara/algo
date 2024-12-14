package symboltable

// HashOpts represents configuration options for a hash table.
type HashOpts struct {
	// The initial capacity of the hash table (must be a power of 2 for efficient hashing).
	InitialCap int
	// The minimum load factor before resizing (shrinking) the hash table.
	MinLoadFactor float32
	// The maximum load factor before resizing (expanding) the hash table.
	MaxLoadFactor float32
}

func (h *HashOpts) verify() {
	m := h.InitialCap
	isPowOf2 := m&(m-1) == 0
	if m < 4 || !isPowOf2 {
		panic("The hash table capacity must be at least 4 and a power of 2 for efficient hashing.")
	}
}

// hashTableEntry represents an entry in a non-linear probing hash table that requires soft deletion.
type hashTableEntry[K, V any] struct {
	key     K
	val     V
	deleted bool
}
