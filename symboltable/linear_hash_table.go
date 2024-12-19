package symboltable

import (
	"fmt"
	"iter"
	"strings"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

const (
	lpMinM          = 32    // Minimum number of entries in the hash table (must be at least 4 and a power of 2 for efficient hashing)
	lpMinLoadFactor = 0.125 // Minimum load factor before resizing (shrinking)
	lpMaxLoadFactor = 0.50  // Maximum load factor before resizing (expanding)
)

// linearHashTable is a hash table with linear probing for conflict resolution.
type linearHashTable[K, V any] struct {
	entries []*KeyValue[K, V]
	m       int     // The total number of entries in the hash table
	n       int     // The number of key-values stored in the hash table
	minLF   float32 // The minimum load factor before resizing (shrinking) the hash table
	maxLF   float32 // The maximum load factor before resizing (expanding) the hash table

	hashKey HashFunc[K]
	eqKey   EqualFunc[K]
	eqVal   EqualFunc[V]
}

// NewLinearHashTable creates a new hash table with linear probing for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// This hash table implements open addressing with linear probing, where collisions are resolved
// by checking subsequent indices in a linear fashion (i+1, i+2, i+3, ...) until an empty slot is found.
func NewLinearHashTable[K, V any](hashKey HashFunc[K], eqKey EqualFunc[K], eqVal EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap == 0 {
		opts.InitialCap = lpMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = lpMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = lpMaxLoadFactor
	}

	opts.verify()

	return &linearHashTable[K, V]{
		entries: make([]*KeyValue[K, V], opts.InitialCap),
		m:       opts.InitialCap,
		n:       0,
		minLF:   opts.MinLoadFactor,
		maxLF:   opts.MaxLoadFactor,
		hashKey: hashKey,
		eqKey:   eqKey,
		eqVal:   eqVal,
	}
}

// nolint: unused
func (ht *linearHashTable[K, V]) verify() bool {
	if lf := ht.loadFactor(); lf > ht.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, e := range ht.entries {
		if e != nil {
			if _, ok := ht.Get(e.Key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In linear probing, the load factor ranges between 0 and 1.
func (ht *linearHashTable[K, V]) loadFactor() float32 {
	return float32(ht.n) / float32(ht.m)
}

// probe returns a function that generates the next index in a linear probing sequence.
// The sequence starts at h and increments linearly: h, h+1, h+2, h+3, ...
func (ht *linearHashTable[K, V]) probe(key K) func() int {
	h := ht.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	// M must be a power of 2
	M := uint64(ht.m)
	h1 := h & (M - 1) // [0, M-1]

	var i, next uint64

	return func() int {
		if i == 0 {
			next = h1
		} else {
			next = (h1 + i) % M
		}

		i++
		return int(next)
	}
}

// resize adjusts the hash table to a new size and re-hashes all keys.
func (ht *linearHashTable[K, V]) resize(m int) {
	newHT := NewLinearHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		InitialCap:    m,
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	}).(*linearHashTable[K, V])

	for key, val := range ht.All() {
		newHT.Put(key, val)
	}

	ht.entries = newHT.entries
	ht.m = newHT.m
	ht.n = newHT.n
}

// Size returns the number of key-values in the hash table.
func (ht *linearHashTable[K, V]) Size() int {
	return ht.n
}

// IsEmpty returns true if the hash table is empty.
func (ht *linearHashTable[K, V]) IsEmpty() bool {
	return ht.n == 0
}

// Put adds a new key-value to the hash table.
func (ht *linearHashTable[K, V]) Put(key K, val V) {
	if ht.loadFactor() >= ht.maxLF {
		ht.resize(2 * ht.m)
	}

	var i int
	next := ht.probe(key)
	for i = next(); ht.entries[i] != nil; i = next() {
		if ht.eqKey(ht.entries[i].Key, key) {
			ht.entries[i].Val = val
			return
		}
	}

	ht.entries[i] = &KeyValue[K, V]{
		Key: key,
		Val: val,
	}

	ht.n++
}

// Get returns the value of a given key in the hash table.
func (ht *linearHashTable[K, V]) Get(key K) (V, bool) {
	next := ht.probe(key)
	for i := next(); ht.entries[i] != nil; i = next() {
		if ht.eqKey(ht.entries[i].Key, key) {
			return ht.entries[i].Val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value from the hash table.
func (ht *linearHashTable[K, V]) Delete(key K) (V, bool) {
	next := ht.probe(key)
	i := next()
	for ht.entries[i] != nil && !ht.eqKey(ht.entries[i].Key, key) {
		i = next()
	}

	// Key not found
	if ht.entries[i] == nil {
		var zeroV V
		return zeroV, false
	}

	// Remove the entry from the hash table
	val := ht.entries[i].Val
	ht.entries[i] = nil
	ht.n--

	// Re-hash all subsequent entries to maintain the probe sequence
	for i = next(); ht.entries[i] != nil; i = next() {
		key, val := ht.entries[i].Key, ht.entries[i].Val
		ht.entries[i] = nil
		ht.n--
		ht.Put(key, val)
	}

	if ht.m > scMinM && ht.loadFactor() <= ht.minLF {
		ht.resize(ht.m / 2)
	}

	return val, true
}

// String returns a string representation of the hash table.
func (ht *linearHashTable[K, V]) String() string {
	pairs := make([]string, ht.Size())
	i := 0

	for key, val := range ht.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equals determines whether or not two hash tables have the same key-values.
func (ht *linearHashTable[K, V]) Equals(rhs SymbolTable[K, V]) bool {
	ht2, ok := rhs.(*linearHashTable[K, V])
	if !ok {
		return false
	}

	return ht.AllMatch(func(key K, val V) bool { // ht ⊂ ht2
		v, ok := ht2.Get(key)
		return ok && ht.eqVal(val, v)
	}) && ht2.AllMatch(func(key K, val V) bool { // ht2 ⊂ ht
		v, ok := ht.Get(key)
		return ok && ht.eqVal(val, v)
	})
}

// All returns an iterator sequence containing all the key-values in the hash table.
func (ht *linearHashTable[K, V]) All() iter.Seq2[K, V] {
	// Create a list of indices representing the entries.
	indices := make([]int, len(ht.entries))
	for i := range indices {
		indices[i] = i
	}

	// Shuffle the indices list to randomize the order in which entries are traversed.
	// This ensures that the traversal order is non-deterministic, reflecting the unordered nature of hash table.
	r.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	return func(yield func(K, V) bool) {
		for _, i := range indices {
			if e := ht.entries[i]; e != nil {
				if !yield(e.Key, e.Val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (ht *linearHashTable[K, V]) AnyMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (ht *linearHashTable[K, V]) AllMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// SelectMatch selects a subset of key-values from the hash table that satisfy the given predicate.
// It returns a new hash table containing the matching key-values, of the same type as the original hash table.
func (ht *linearHashTable[K, V]) SelectMatch(p Predicate2[K, V]) Collection2[K, V] {
	newHT := NewLinearHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	}).(*linearHashTable[K, V])

	for key, val := range ht.All() {
		if p(key, val) {
			newHT.Put(key, val)
		}
	}

	return newHT
}

// print displays the current state of the hash table in the terminal,
// including its parameters and a detailed table of indices, key-values, and hash function calculations.
//
// This method is intended for debugging and troubleshooting purposes.
/* func (ht *linearHashTable[K, V]) print() {
	header := fmt.Sprintf("M: %d    N: %d    Min LF: %.2f    Max LF: %.2f    Load Factor: %.2f",
		ht.m, ht.n, ht.minLF, ht.maxLF, ht.loadFactor())

	fmt.Printf("┌─────────────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  %-89s  │\n", header)
	fmt.Printf("├─────┬──────────────────────────────┬───────────────────────┬────────┬─────┬─────┬─────┬─────┤\n")
	fmt.Printf("│Index│          Key-Value           │       hash(key)       │ h(key) │ h+1 │ h+2 │ h+3 │ h+4 │\n")
	fmt.Printf("├─────┼──────────────────────────────┼───────────────────────┼────────┼─────┼─────┼─────┼─────┤\n")

	for i, kv := range ht.entries {
		if kv == nil {
			fmt.Printf("│ %-3d │                              │                       │        │     │     │     │     │\n", i)
		} else {
			pair := fmt.Sprintf("%v:%v", kv.Key, kv.Val)

			h := ht.hashKey(kv.Key)
			h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
			hash := fmt.Sprintf("%-20d", h)

			next := ht.probe(kv.Key)

			i0 := fmt.Sprintf("%-4d", next())
			i1 := fmt.Sprintf("%-2d", next())
			i2 := fmt.Sprintf("%-2d", next())
			i3 := fmt.Sprintf("%-2d", next())
			i4 := fmt.Sprintf("%-2d", next())

			fmt.Printf("│ %-3d │ %-28s │  %s │   %s │  %s │  %s │  %s │  %s │\n", i, pair, hash, i0, i1, i2, i3, i4)
		}

		if i < len(ht.entries)-1 {
			fmt.Println("├─────┼──────────────────────────────┼───────────────────────┼────────┼─────┼─────┼─────┼─────┤")
		} else {
			fmt.Println("└─────┴──────────────────────────────┴───────────────────────┴────────┴─────┴─────┴─────┴─────┘")
		}
	}
} */
