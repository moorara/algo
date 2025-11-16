package symboltable

// HashOpts represents configuration options for a hash table.
type HashOpts struct {
	// The initial capacity of the hash table.
	// For chain and linear hash tables, it must be a power of 2 for efficient hashing.
	// For quadratic and double hash tables, it must be a prime number for better distribution.
	InitialCap int
	// The minimum load factor before resizing (shrinking) the hash table.
	MinLoadFactor float32
	// The maximum load factor before resizing (expanding) the hash table.
	MaxLoadFactor float32
}

// hashTableEntry represents an entry in a non-linear probing hash table that requires soft deletion.
type hashTableEntry[K, V any] struct {
	key     K
	val     V
	deleted bool
}
