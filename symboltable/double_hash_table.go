package symboltable

import (
	"fmt"
	"iter"
	"strings"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

const (
	dhMinM          = 32    // Minimum number of entries in the hash table (must be at least 4 and a power of 2 for efficient hashing)
	dhMinLoadFactor = 0.125 // Minimum load factor before resizing (shrinking)
	dhMaxLoadFactor = 0.50  // Maximum load factor before resizing (expanding)
)

// isPrime determines whether or not a given integer n is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Check for prime numbers less 100 directly
	if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 || n == 23 || n == 29 || n == 31 || n == 37 ||
		n == 41 || n == 43 || n == 47 || n == 53 || n == 59 || n == 61 || n == 67 || n == 71 || n == 73 || n == 79 || n == 83 || n == 89 || n == 97 {
		return true
	} else if n <= 100 {
		return false
	}

	// Check if n is prime using trial division
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// gcd computes the greatest common divisor of two numbers.
// It implements the Euclidean algorithm.
func gcd(a, b uint64) uint64 {
	// Ensure a ≥ b
	a, b = max(a, b), min(a, b)

	/*
	 * Let the quotient be q and the remainder be r, so that a = b × q + r
	 * Replace a with b and b with r
	 * Repeat this until the remainder r becomes 0
	 * The GCD is the last non-zero remainder
	 */

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// doubleHashTable is a hash table with double hashing for conflict resolution.
type doubleHashTable[K, V any] struct {
	entries []*hashTableEntry[K, V]
	m       int     // Total number of entries in the hash table
	p       int     // Largest prime number less than m, used in secondary hashing
	n       int     // Number of key-values stored in the hash table
	minLF   float32 // Minimum load factor before resizing (shrinking) the hash table
	maxLF   float32 // Maximum load factor before resizing (expanding) the hash table

	hashKey HashFunc[K]
	eqKey   EqualFunc[K]
	eqVal   EqualFunc[V]
}

// NewDoubleHashTable creates a new hash table with double hashing for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// This hash table implements open addressing with double hashing, where collisions are resolved
// by applying a second hash function to determine the step size for probing.
// The indices are computed as h₁, h₁ + h₂, h₁ + 2h₂, h₁ + 3h₂, ...,
// where h₁ is the primary hash and h₂ is the secondary hash.
func NewDoubleHashTable[K, V any](hashKey HashFunc[K], eqKey EqualFunc[K], eqVal EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap == 0 {
		opts.InitialCap = dhMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = dhMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = dhMaxLoadFactor
	}

	opts.verify()

	// Find the biggest prime smaller than M
	var p int
	for p = opts.InitialCap; p >= 2; p-- {
		if isPrime(p) {
			break
		}
	}

	return &doubleHashTable[K, V]{
		entries: make([]*hashTableEntry[K, V], opts.InitialCap),
		m:       opts.InitialCap,
		p:       p,
		n:       0,
		minLF:   opts.MinLoadFactor,
		maxLF:   opts.MaxLoadFactor,
		hashKey: hashKey,
		eqKey:   eqKey,
		eqVal:   eqVal,
	}
}

// nolint: unused
func (ht *doubleHashTable[K, V]) verify() bool {
	if lf := ht.loadFactor(); lf > ht.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, e := range ht.entries {
		if e != nil && !e.deleted {
			if _, ok := ht.Get(e.key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In double hashing, the load factor ranges between 0 and 1.
func (ht *doubleHashTable[K, V]) loadFactor() float32 {
	return float32(ht.n) / float32(ht.m)
}

// probe returns a function that generates the next index in a double hashing sequence.
// The sequence starts at h₁ and increments by h₂ on each step: h₁, h₁+h₂, h₁+2h₂, h₁+3h₂, ...
func (ht *doubleHashTable[K, V]) probe(key K) func() int {
	h := ht.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	// M must be a power of 2
	M, P := uint64(ht.m), uint64(ht.p)
	h1 := h & (M - 1) // [0, M-1]

	var h2, i, next uint64

	return func() int {
		if i == 0 {
			next = h1
		} else {
			if i == 1 {
				// The secondary hash value, h₂(k), must satisfy the following properties:
				//   1. h₂(k) ≠ 0: Ensures progress in probing and prevents infinite loops.
				//   2. h₂(k) and m are coprime: Guarantees the entire table is cycled through for any key.
				//   3. Independence from the primary hash, h₁(k): Avoids overlapping collision paths.
				//   4. Efficient computation: Suitable for high-performance hash table operations.
				h2 = P - (h % P) // [1, p] (p < m)

				// The hash table size (M) and the secondary hash must be relatively prime.
				// If they share a common divisor greater than 1, the sequence generated by (h1 + i*h2) % M
				// will only cover a subset of the indices in [0, M-1] and an infinite loop occurs.
				if gcd(M, h2) != 1 {
					h2++
				}
			}

			next = (h1 + i*h2) % M
		}

		i++
		return int(next)
	}
}

// resize adjusts the hash table to a new size and re-hashes all keys.
func (ht *doubleHashTable[K, V]) resize(m int) {
	newHT := NewDoubleHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		InitialCap:    m,
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	}).(*doubleHashTable[K, V])

	for key, val := range ht.All() {
		newHT.Put(key, val)
	}

	ht.entries = newHT.entries
	ht.m = newHT.m
	ht.p = newHT.p
	ht.n = newHT.n
}

// Size returns the number of key-values in the hash table.
func (ht *doubleHashTable[K, V]) Size() int {
	return ht.n
}

// IsEmpty returns true if the hash table is empty.
func (ht *doubleHashTable[K, V]) IsEmpty() bool {
	return ht.n == 0
}

// Put adds a new key-value to the hash table.
func (ht *doubleHashTable[K, V]) Put(key K, val V) {
	if ht.loadFactor() >= ht.maxLF {
		ht.resize(2 * ht.m)
	}

	var i int
	next := ht.probe(key)
	for i = next(); ht.entries[i] != nil; i = next() {
		if ht.eqKey(ht.entries[i].key, key) {
			ht.entries[i].val = val
			ht.entries[i].deleted = false
			return
		}
	}

	ht.entries[i] = &hashTableEntry[K, V]{
		key:     key,
		val:     val,
		deleted: false,
	}

	ht.n++
}

// Get returns the value of a given key in the hash table.
func (ht *doubleHashTable[K, V]) Get(key K) (V, bool) {
	next := ht.probe(key)
	for i := next(); ht.entries[i] != nil; i = next() {
		if !ht.entries[i].deleted && ht.eqKey(ht.entries[i].key, key) {
			return ht.entries[i].val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value from the hash table.
func (ht *doubleHashTable[K, V]) Delete(key K) (V, bool) {
	next := ht.probe(key)
	i := next()
	for ht.entries[i] != nil && !ht.eqKey(ht.entries[i].key, key) {
		i = next()
	}

	// Key not found
	if ht.entries[i] == nil || ht.entries[i].deleted {
		var zeroV V
		return zeroV, false
	}

	// Remove the entry from the hash table
	val := ht.entries[i].val
	ht.entries[i].deleted = true
	ht.n--

	// During resizing, soft-deleted keys are removed, and remaining active keys are rehashed
	if ht.m > scMinM && ht.loadFactor() <= ht.minLF {
		ht.resize(ht.m / 2)
	}

	return val, true
}

// String returns a string representation of the hash table.
func (ht *doubleHashTable[K, V]) String() string {
	pairs := make([]string, ht.Size())
	i := 0

	for key, val := range ht.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equals determines whether or not two hash tables have the same key-values.
func (ht *doubleHashTable[K, V]) Equals(rhs SymbolTable[K, V]) bool {
	ht2, ok := rhs.(*doubleHashTable[K, V])
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
func (ht *doubleHashTable[K, V]) All() iter.Seq2[K, V] {
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
			if e := ht.entries[i]; e != nil && !e.deleted {
				if !yield(e.key, e.val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (ht *doubleHashTable[K, V]) AnyMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (ht *doubleHashTable[K, V]) AllMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// SelectMatch selects a subset of key-values from the hash table that satisfy the given predicate.
// It returns a new hash table containing the matching key-values, of the same type as the original hash table.
func (ht *doubleHashTable[K, V]) SelectMatch(p Predicate2[K, V]) Collection2[K, V] {
	newHT := NewDoubleHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	})

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
/* func (ht *doubleHashTable[K, V]) print() {
	reset := "\033[00m"
	red := "\033[31m"

	header := fmt.Sprintf("M: %d    N: %d    P: %d    Min LF: %.2f    Max LF: %.2f    Load Factor: %.2f",
		ht.m, ht.n, ht.p, ht.minLF, ht.maxLF, ht.loadFactor())

	fmt.Printf("┌───────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  %-111s  │\n", header)
	fmt.Printf("├─────┬──────────────────────────────┬───────────────────────┬─────────┬─────────┬───────┬────────┬────────┬────────┤\n")
	fmt.Printf("│Index│          Key-Value           │       hash(key)       │ h₁(key) │ h₂(key) │ h₁+h₂ │ h₁+2h₂ │ h₁+3h₂ │ h₁+4h₂ │\n")
	fmt.Printf("├─────┼──────────────────────────────┼───────────────────────┼─────────┼─────────┼───────┼────────┼────────┼────────┤\n")

	for i, kv := range ht.entries {
		if kv == nil {
			fmt.Printf("│ %-3d │                              │                       │         │         │       │        │        │        │\n", i)
		} else {
			color := reset
			if kv.deleted {
				color = red
			}

			pair := fmt.Sprintf("%s%v:%v%s", color, kv.key, kv.val, reset)

			h := ht.hashKey(kv.key)
			h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)
			hash := fmt.Sprintf("%s%-20d%s", color, h, reset)

			next := ht.probe(kv.key)
			h1 := fmt.Sprintf("%s%-5d%s", color, next(), reset)

			P := uint64(ht.p)
			h2 := fmt.Sprintf("%s%-5d%s", color, P-(h%P), reset)

			i1 := fmt.Sprintf("%s%-4d%s", color, next(), reset)
			i2 := fmt.Sprintf("%s%-4d%s", color, next(), reset)
			i3 := fmt.Sprintf("%s%-4d%s", color, next(), reset)
			i4 := fmt.Sprintf("%s%-4d%s", color, next(), reset)

			fmt.Printf("│ %-3d │ %-38s │  %s │   %s │   %s │  %s │   %s │   %s │   %s │\n", i, pair, hash, h1, h2, i1, i2, i3, i4)
		}

		if i < len(ht.entries)-1 {
			fmt.Printf("├─────┼──────────────────────────────┼───────────────────────┼─────────┼─────────┼───────┼────────┼────────┼────────┤\n")
		} else {
			fmt.Printf("└─────┴──────────────────────────────┴───────────────────────┴─────────┴─────────┴───────┴────────┴────────┴────────┘\n")
		}
	}
} */
