package symboltable

import (
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/math"
)

const (
	lpMinM          = 32    // Minimum number of entries in the hash table (must be a power of 2 for efficient hashing)
	lpMinLoadFactor = 0.125 // Minimum load factor before resizing (shrinking)
	lpMaxLoadFactor = 0.50  // Maximum load factor before resizing (expanding)
)

// linearHashTable is a hash table with linear probing for conflict resolution.
type linearHashTable[K, V any] struct {
	entries []*generic.KeyValue[K, V]
	m       int     // The total number of entries in the hash table
	n       int     // The number of key-values stored in the hash table
	minLF   float32 // The minimum load factor before resizing (shrinking) the hash table
	maxLF   float32 // The maximum load factor before resizing (expanding) the hash table

	hashKey hash.HashFunc[K]
	eqKey   generic.EqualFunc[K]
	eqVal   generic.EqualFunc[V]
}

// NewLinearHashTable creates a new hash table with linear probing for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// This hash table implements open addressing with linear probing, where collisions are resolved
// by checking subsequent indices in a linear fashion (i+1, i+2, i+3, ...) until an empty slot is found.
func NewLinearHashTable[K, V any](hashKey hash.HashFunc[K], eqKey generic.EqualFunc[K], eqVal generic.EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap < lpMinM {
		opts.InitialCap = lpMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = lpMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = lpMaxLoadFactor
	}

	if M := opts.InitialCap; !math.IsPowerOf2(M) {
		panic("The hash table capacity must be a power of 2")
	}

	return &linearHashTable[K, V]{
		entries: make([]*generic.KeyValue[K, V], opts.InitialCap),
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
func (t *linearHashTable[K, V]) verify() bool {
	if lf := t.loadFactor(); lf > t.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, e := range t.entries {
		if e != nil {
			if _, ok := t.Get(e.Key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In linear probing, the load factor ranges between 0 and 1.
func (t *linearHashTable[K, V]) loadFactor() float32 {
	return float32(t.n) / float32(t.m)
}

// probe returns a function that generates the next index in a linear probing sequence.
// The sequence starts at h and increments linearly: h, h+1, h+2, h+3, ...
func (t *linearHashTable[K, V]) probe(key K) func() int {
	h := t.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	// M must be a power of 2
	M := uint64(t.m)
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
func (t *linearHashTable[K, V]) resize(m int) {
	// Ensure the minimum table size
	if m < lpMinM {
		return
	}

	new := &linearHashTable[K, V]{
		entries: make([]*generic.KeyValue[K, V], m),
		m:       m,
		n:       0,
		minLF:   t.minLF,
		maxLF:   t.maxLF,
		hashKey: t.hashKey,
		eqKey:   t.eqKey,
		eqVal:   t.eqVal,
	}

	for _, e := range t.entries {
		if e != nil {
			new.Put(e.Key, e.Val)
		}
	}

	t.entries = new.entries
	t.m = new.m
	t.n = new.n
}

// Size returns the number of key-values in the hash table.
func (t *linearHashTable[K, V]) Size() int {
	return t.n
}

// IsEmpty returns true if the hash table is empty.
func (t *linearHashTable[K, V]) IsEmpty() bool {
	return t.n == 0
}

// Put adds a new key-value to the hash table.
func (t *linearHashTable[K, V]) Put(key K, val V) {
	if t.loadFactor() >= t.maxLF {
		t.resize(2 * t.m)
	}

	var i int
	next := t.probe(key)
	for i = next(); t.entries[i] != nil; i = next() {
		if t.eqKey(t.entries[i].Key, key) {
			t.entries[i].Val = val
			return
		}
	}

	t.entries[i] = &generic.KeyValue[K, V]{
		Key: key,
		Val: val,
	}

	t.n++
}

// Get returns the value of a given key in the hash table.
func (t *linearHashTable[K, V]) Get(key K) (V, bool) {
	next := t.probe(key)
	for i := next(); t.entries[i] != nil; i = next() {
		if t.eqKey(t.entries[i].Key, key) {
			return t.entries[i].Val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete deletes a key-value from the hash table.
func (t *linearHashTable[K, V]) Delete(key K) (V, bool) {
	next := t.probe(key)
	i := next()
	for t.entries[i] != nil && !t.eqKey(t.entries[i].Key, key) {
		i = next()
	}

	// Key not found
	if t.entries[i] == nil {
		var zeroV V
		return zeroV, false
	}

	// Remove the entry from the hash table
	val := t.entries[i].Val
	t.entries[i] = nil
	t.n--

	// Re-hash all subsequent entries to maintain the probe sequence
	for i = next(); t.entries[i] != nil; i = next() {
		key, val := t.entries[i].Key, t.entries[i].Val
		t.entries[i] = nil
		t.n--
		t.Put(key, val)
	}

	if t.loadFactor() <= t.minLF {
		t.resize(t.m / 2)
	}

	return val, true
}

// DeleteAll deletes all key-values from the hash table, leaving it empty.
func (t *linearHashTable[K, V]) DeleteAll() {
	t.entries = make([]*generic.KeyValue[K, V], t.m)
	t.n = 0
}

// String returns a string representation of the hash table.
func (t *linearHashTable[K, V]) String() string {
	pairs := make([]string, t.Size())
	i := 0

	for key, val := range t.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two hash tables have the same key-values.
func (t *linearHashTable[K, V]) Equal(rhs SymbolTable[K, V]) bool {
	tt, ok := rhs.(*linearHashTable[K, V])
	if !ok {
		return false
	}

	return t.AllMatch(func(key K, val V) bool { // t ⊂ tt
		v, ok := tt.Get(key)
		return ok && t.eqVal(val, v)
	}) && tt.AllMatch(func(key K, val V) bool { // tt ⊂ t
		v, ok := t.Get(key)
		return ok && t.eqVal(val, v)
	})
}

// All returns an iterator sequence containing all the key-values in the hash table.
func (t *linearHashTable[K, V]) All() iter.Seq2[K, V] {
	// Create a list of indices representing the entries.
	indices := make([]int, len(t.entries))
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
			if e := t.entries[i]; e != nil {
				if !yield(e.Key, e.Val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (t *linearHashTable[K, V]) AnyMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (t *linearHashTable[K, V]) AllMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// FirstMatch returns the first key-value in the hash table that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *linearHashTable[K, V]) FirstMatch(p generic.Predicate2[K, V]) (K, V, bool) {
	for key, val := range t.All() {
		if p(key, val) {
			return key, val, true
		}
	}

	var zeroK K
	var zeroV V
	return zeroK, zeroV, false
}

// SelectMatch selects a subset of key-values from the hash table that satisfy the given predicate.
// It returns a new hash table containing the matching key-values, of the same type as the original hash table.
func (t *linearHashTable[K, V]) SelectMatch(p generic.Predicate2[K, V]) generic.Collection2[K, V] {
	new := NewLinearHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
		MinLoadFactor: t.minLF,
		MaxLoadFactor: t.maxLF,
	})

	for key, val := range t.All() {
		if p(key, val) {
			new.Put(key, val)
		}
	}

	return new
}

// PartitionMatch partitions the key-values in the hash table
// into two separate hash tables based on the provided predicate.
// The first hash table contains the key-values that satisfy the predicate (matched key-values),
// while the second hash table contains those that do not satisfy the predicate (unmatched key-values).
// Both hash tables are of the same type as the original hash table.
func (t *linearHashTable[K, V]) PartitionMatch(p generic.Predicate2[K, V]) (generic.Collection2[K, V], generic.Collection2[K, V]) {
	matched := NewLinearHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
		MinLoadFactor: t.minLF,
		MaxLoadFactor: t.maxLF,
	})

	unmatched := NewLinearHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
		MinLoadFactor: t.minLF,
		MaxLoadFactor: t.maxLF,
	})

	for key, val := range t.All() {
		if p(key, val) {
			matched.Put(key, val)
		} else {
			unmatched.Put(key, val)
		}
	}

	return matched, unmatched
}

// print displays the current state of the hash table in the terminal,
// including its parameters and a detailed table of indices, key-values, and hash function calculations.
//
// This method is intended for debugging and troubleshooting purposes.
/* func (t *linearHashTable[K, V]) print() {
	header := fmt.Sprintf("M: %d    N: %d    Min LF: %.2f    Max LF: %.2f    Load Factor: %.2f",
		t.m, t.n, t.minLF, t.maxLF, t.loadFactor())

	fmt.Printf("┌─────────────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  %-89s  │\n", header)
	fmt.Printf("├─────┬──────────────────────────────┬───────────────────────┬────────┬─────┬─────┬─────┬─────┤\n")
	fmt.Printf("│Index│          Key-Value           │       hash(key)       │ h(key) │ h+1 │ h+2 │ h+3 │ h+4 │\n")
	fmt.Printf("├─────┼──────────────────────────────┼───────────────────────┼────────┼─────┼─────┼─────┼─────┤\n")

	for i, kv := range t.entries {
		if kv == nil {
			fmt.Printf("│ %-3d │                              │                       │        │     │     │     │     │\n", i)
		} else {
			pair := fmt.Sprintf("%v:%v", kv.Key, kv.Val)

			h := t.hashKey(kv.Key)
			h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
			hash := fmt.Sprintf("%-20d", h)

			next := t.probe(kv.Key)

			i0 := fmt.Sprintf("%-4d", next())
			i1 := fmt.Sprintf("%-2d", next())
			i2 := fmt.Sprintf("%-2d", next())
			i3 := fmt.Sprintf("%-2d", next())
			i4 := fmt.Sprintf("%-2d", next())

			fmt.Printf("│ %-3d │ %-28s │  %s │   %s │  %s │  %s │  %s │  %s │\n", i, pair, hash, i0, i1, i2, i3, i4)
		}

		if i < len(t.entries)-1 {
			fmt.Println("├─────┼──────────────────────────────┼───────────────────────┼────────┼─────┼─────┼─────┼─────┤")
		} else {
			fmt.Println("└─────┴──────────────────────────────┴───────────────────────┴────────┴─────┴─────┴─────┴─────┘")
		}
	}
} */
