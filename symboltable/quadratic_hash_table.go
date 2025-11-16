package symboltable

import (
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/math"
)

/*
 * If i ≥ M:
 *   • The quadratic probing sequence may revisit indices due to periodicity in i² % M.
 *   • This can lead to infinite loops or redundant checks if the table is not full.
 *   • For non-prime M, some slots may never be visited because (h1 + i²) % M might not yield all indices in [0, M-1]
 *
 * To ensure robustness:
 *   • Use a prime number for M to guarantee the sequence can visit all slots.
 *   • Alternatively, handle the case where i ≥ M by resizing the table or signaling an error.
 */

const (
	qpMinM          = 31    // Minimum number of entries in the hash table (must be a prime number)
	qpMinLoadFactor = 0.125 // Minimum load factor before resizing (shrinking)
	qpMaxLoadFactor = 0.50  // Maximum load factor before resizing (expanding)
)

// quadraticHashTable is a hash table with quadratic probing for conflict resolution.
type quadraticHashTable[K, V any] struct {
	entries []*hashTableEntry[K, V]
	m       int     // The total number of entries in the hash table
	n       int     // The number of key-values stored in the hash table
	minLF   float32 // The minimum load factor before resizing (shrinking) the hash table
	maxLF   float32 // The maximum load factor before resizing (expanding) the hash table

	hashKey hash.HashFunc[K]
	eqKey   generic.EqualFunc[K]
	eqVal   generic.EqualFunc[V]
}

// NewQuadraticHashTable creates a new hash table with quadratic probing for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// This hash table implements open addressing with quadratic probing, where collisions are resolved
// by checking subsequent indices using a quadratic function (i+1², i+2², i+3², ...) until an empty slot is found.
func NewQuadraticHashTable[K, V any](hashKey hash.HashFunc[K], eqKey generic.EqualFunc[K], eqVal generic.EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap < qpMinM {
		opts.InitialCap = qpMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = qpMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = qpMaxLoadFactor
	}

	if M := opts.InitialCap; !math.IsPrime(M) {
		panic("The hash table capacity must be a prime number")
	}

	return &quadraticHashTable[K, V]{
		entries: make([]*hashTableEntry[K, V], opts.InitialCap),
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
func (t *quadraticHashTable[K, V]) verify() bool {
	if lf := t.loadFactor(); lf > t.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, e := range t.entries {
		if e != nil && !e.deleted {
			if _, ok := t.Get(e.key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In quadratic probing, the load factor ranges between 0 and 1.
func (t *quadraticHashTable[K, V]) loadFactor() float32 {
	return float32(t.n) / float32(t.m)
}

// probe returns a function that generates the next index in a quadratic probing sequence.
// The sequence starts at h and increments quadratically: h, h+1², h+2², h+3², ...
func (t *quadraticHashTable[K, V]) probe(key K) func() int {
	h := t.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	M := uint64(t.m)
	h1 := h % M // [0, M-1]

	var i, next uint64

	return func() int {
		if i == 0 {
			next = h1
		} else {
			next = (h1 + i*i) % M
		}

		i++
		return int(next)
	}
}

// resize adjusts the hash table to a new size and re-hashes all keys.
func (t *quadraticHashTable[K, V]) resize(m int) {
	// Ensure the minimum table size
	if m < qpMinM {
		return
	}

	// Ensure the table size remains prime
	m = math.SmallestPrimeLargerThan(m)

	new := &quadraticHashTable[K, V]{
		entries: make([]*hashTableEntry[K, V], m),
		m:       m,
		n:       0,
		minLF:   t.minLF,
		maxLF:   t.maxLF,
		hashKey: t.hashKey,
		eqKey:   t.eqKey,
		eqVal:   t.eqVal,
	}

	for _, e := range t.entries {
		if e != nil && !e.deleted {
			new.Put(e.key, e.val)
		}
	}

	t.entries = new.entries
	t.m = new.m
	t.n = new.n
}

// Size returns the number of key-values in the hash table.
func (t *quadraticHashTable[K, V]) Size() int {
	return t.n
}

// IsEmpty returns true if the hash table is empty.
func (t *quadraticHashTable[K, V]) IsEmpty() bool {
	return t.n == 0
}

// Put adds a new key-value to the hash table.
func (t *quadraticHashTable[K, V]) Put(key K, val V) {
	if t.loadFactor() >= t.maxLF {
		t.resize(2 * t.m)
	}

	var i int
	next := t.probe(key)
	for i = next(); t.entries[i] != nil; i = next() {
		if t.eqKey(t.entries[i].key, key) {
			t.entries[i].val = val
			t.entries[i].deleted = false
			return
		}
	}

	t.entries[i] = &hashTableEntry[K, V]{
		key:     key,
		val:     val,
		deleted: false,
	}

	t.n++
}

// Get returns the value of a given key in the hash table.
func (t *quadraticHashTable[K, V]) Get(key K) (V, bool) {
	next := t.probe(key)
	for i := next(); t.entries[i] != nil; i = next() {
		if !t.entries[i].deleted && t.eqKey(t.entries[i].key, key) {
			return t.entries[i].val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete deletes a key-value from the hash table.
func (t *quadraticHashTable[K, V]) Delete(key K) (V, bool) {
	next := t.probe(key)
	i := next()
	for t.entries[i] != nil && !t.eqKey(t.entries[i].key, key) {
		i = next()
	}

	// Key not found
	if t.entries[i] == nil || t.entries[i].deleted {
		var zeroV V
		return zeroV, false
	}

	// Remove the entry from the hash table
	val := t.entries[i].val
	t.entries[i].deleted = true
	t.n--

	// During resizing, soft-deleted keys are removed, and remaining active keys are rehashed
	if t.loadFactor() <= t.minLF {
		t.resize(t.m / 2)
	}

	return val, true
}

// DeleteAll deletes all key-values from the hash table, leaving it empty.
func (t *quadraticHashTable[K, V]) DeleteAll() {
	t.entries = make([]*hashTableEntry[K, V], t.m)
	t.n = 0
}

// String returns a string representation of the hash table.
func (t *quadraticHashTable[K, V]) String() string {
	pairs := make([]string, t.Size())
	i := 0

	for key, val := range t.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two hash tables have the same key-values.
func (t *quadraticHashTable[K, V]) Equal(rhs SymbolTable[K, V]) bool {
	tt, ok := rhs.(*quadraticHashTable[K, V])
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
func (t *quadraticHashTable[K, V]) All() iter.Seq2[K, V] {
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
			if e := t.entries[i]; e != nil && !e.deleted {
				if !yield(e.key, e.val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (t *quadraticHashTable[K, V]) AnyMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (t *quadraticHashTable[K, V]) AllMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// FirstMatch returns the first key-value in the hash table that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *quadraticHashTable[K, V]) FirstMatch(p generic.Predicate2[K, V]) (K, V, bool) {
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
func (t *quadraticHashTable[K, V]) SelectMatch(p generic.Predicate2[K, V]) generic.Collection2[K, V] {
	new := NewQuadraticHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
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
func (t *quadraticHashTable[K, V]) PartitionMatch(p generic.Predicate2[K, V]) (generic.Collection2[K, V], generic.Collection2[K, V]) {
	matched := NewQuadraticHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
		MinLoadFactor: t.minLF,
		MaxLoadFactor: t.maxLF,
	})

	unmatched := NewQuadraticHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
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
/* func (t *quadraticHashTable[K, V]) print() {
	reset := "\033[00m"
	red := "\033[31m"

	header := fmt.Sprintf("M: %d    N: %d    Min LF: %.2f    Max LF: %.2f    Load Factor: %.2f",
		t.m, t.n, t.minLF, t.maxLF, t.loadFactor())

	fmt.Printf("┌─────────────────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  %-93s  │\n", header)
	fmt.Printf("├─────┬──────────────────────────────┬───────────────────────┬────────┬──────┬──────┬──────┬──────┤\n")
	fmt.Printf("│Index│          Key-Value           │       hash(key)       │ h(key) │ h+1² │ h+2² │ h+3² │ h+4² │\n")
	fmt.Printf("├─────┼──────────────────────────────┼───────────────────────┼────────┼──────┼──────┼──────┼──────┤\n")

	for i, kv := range t.entries {
		if kv == nil {
			fmt.Printf("│ %-3d │                              │                       │        │      │      │      │      │\n", i)
		} else {
			color := reset
			if kv.deleted {
				color = red
			}

			pair := fmt.Sprintf("%s%v:%v%s", color, kv.key, kv.val, reset)

			h := t.hashKey(kv.key)
			h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
			hash := fmt.Sprintf("%s%-20d%s", color, h, reset)

			next := t.probe(kv.key)

			i0 := fmt.Sprintf("%s%-4d%s", color, next(), reset)
			i1 := fmt.Sprintf("%s%-3d%s", color, next(), reset)
			i2 := fmt.Sprintf("%s%-3d%s", color, next(), reset)
			i3 := fmt.Sprintf("%s%-3d%s", color, next(), reset)
			i4 := fmt.Sprintf("%s%-3d%s", color, next(), reset)

			fmt.Printf("│ %-3d │ %-38s │  %s │   %s │  %s │  %s │  %s │  %s │\n", i, pair, hash, i0, i1, i2, i3, i4)
		}

		if i < len(t.entries)-1 {
			fmt.Println("├─────┼──────────────────────────────┼───────────────────────┼────────┼──────┼──────┼──────┼──────┤")
		} else {
			fmt.Println("└─────┴──────────────────────────────┴───────────────────────┴────────┴──────┴──────┴──────┴──────┘")
		}
	}
} */
