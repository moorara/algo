package symboltable

import (
	"fmt"
	"iter"
	"strings"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

const (
	scMinM          = 4    // Minimum number of buckets in the hash table (must be at least 4 and a power of 2 for efficient hashing)
	scMinLoadFactor = 2.0  // Minimum load factor before resizing (shrinking)
	scMaxLoadFactor = 10.0 // Maximum load factor before resizing (expanding)
)

type chainNode[K, V any] struct {
	key  K
	val  V
	next *chainNode[K, V]
}

// chainHashTable is a hash table with separate chaining for conflict resolution.
type chainHashTable[K, V any] struct {
	buckets []*chainNode[K, V]
	m       int     // The total number of buckets in the hash table
	n       int     // The number of key-values stored in the hash table
	minLF   float32 // The minimum load factor before resizing (shrinking) the hash table
	maxLF   float32 // The maximum load factor before resizing (expanding) the hash table

	hashKey HashFunc[K]
	eqKey   EqualFunc[K]
	eqVal   EqualFunc[V]
}

// NewChainHashTable creates a new hash table with separate chaining for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// It resolves hash collisions, where multiple keys hash to the same bucket,
// by maintaining a linked list of all key-values that hash to the same bucket.
// Each bucket contains a chain of elements.
func NewChainHashTable[K, V any](hashKey HashFunc[K], eqKey EqualFunc[K], eqVal EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap == 0 {
		opts.InitialCap = scMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = scMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = scMaxLoadFactor
	}

	opts.verify()

	return &chainHashTable[K, V]{
		buckets: make([]*chainNode[K, V], opts.InitialCap),
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
func (ht *chainHashTable[K, V]) verify() bool {
	if lf := ht.loadFactor(); lf > ht.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, first := range ht.buckets {
		for x := first; x != nil; x = x.next {
			if _, ok := ht.Get(x.key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In separate chaining, the load factor can exceed 1.
func (ht *chainHashTable[K, V]) loadFactor() float32 {
	return float32(ht.n) / float32(ht.m)
}

// hash compute the hash for a key and returns an index in [0, M-1] range.
func (ht *chainHashTable[K, V]) hash(key K) int {
	h := ht.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	// M must be a power of 2
	M := uint64(ht.m)
	h1 := h & (M - 1) // [0, M-1]

	return int(h1)
}

// resize adjusts the hash table to a new size and re-hashes all keys.
func (ht *chainHashTable[K, V]) resize(m int) {
	newHT := NewChainHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		InitialCap:    m,
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	}).(*chainHashTable[K, V])

	for key, val := range ht.All() {
		newHT.Put(key, val)
	}

	ht.buckets = newHT.buckets
	ht.m = newHT.m
	ht.n = newHT.n
}

// Size returns the number of key-values in the hash table.
func (ht *chainHashTable[K, V]) Size() int {
	return ht.n
}

// IsEmpty returns true if the hash table is empty.
func (ht *chainHashTable[K, V]) IsEmpty() bool {
	return ht.n == 0
}

// Put adds a new key-value to the hash table.
func (ht *chainHashTable[K, V]) Put(key K, val V) {
	if ht.loadFactor() >= ht.maxLF {
		ht.resize(2 * ht.m)
	}

	i := ht.hash(key)
	for x := ht.buckets[i]; x != nil; x = x.next {
		if ht.eqKey(x.key, key) {
			x.val = val
			return
		}
	}

	// Add a new node at the beginning of the chain
	ht.buckets[i] = &chainNode[K, V]{
		key:  key,
		val:  val,
		next: ht.buckets[i],
	}

	ht.n++
}

// Get returns the value of a given key in the hash table.
func (ht *chainHashTable[K, V]) Get(key K) (V, bool) {
	i := ht.hash(key)
	for x := ht.buckets[i]; x != nil; x = x.next {
		if ht.eqKey(x.key, key) {
			return x.val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete removes a key-value from the hash table.
func (ht *chainHashTable[K, V]) Delete(key K) (val V, ok bool) {
	i := ht.hash(key)
	ht.buckets[i], val, ok = ht._delete(ht.buckets[i], key)

	if ht.m > scMinM && ht.loadFactor() <= ht.minLF {
		ht.resize(ht.m / 2)
	}

	return val, ok
}

func (ht *chainHashTable[K, V]) _delete(n *chainNode[K, V], key K) (*chainNode[K, V], V, bool) {
	if n == nil {
		var zeroV V
		return nil, zeroV, false
	}

	if ht.eqKey(n.key, key) {
		ht.n--
		return n.next, n.val, true
	}

	var val V
	var ok bool
	n.next, val, ok = ht._delete(n.next, key)

	return n, val, ok
}

// String returns a string representation of the hash table.
func (ht *chainHashTable[K, V]) String() string {
	pairs := make([]string, ht.Size())
	i := 0

	for key, val := range ht.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equals determines whether or not two hash tables have the same key-values.
func (ht *chainHashTable[K, V]) Equals(rhs SymbolTable[K, V]) bool {
	ht2, ok := rhs.(*chainHashTable[K, V])
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
func (ht *chainHashTable[K, V]) All() iter.Seq2[K, V] {
	// Create a list of indices representing the buckets.
	indices := make([]int, len(ht.buckets))
	for i := range indices {
		indices[i] = i
	}

	// Shuffle the indices list to randomize the order in which buckets are traversed.
	// This ensures that the traversal order is non-deterministic, reflecting the unordered nature of hash table.
	r.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	return func(yield func(K, V) bool) {
		for _, i := range indices {
			for x := ht.buckets[i]; x != nil; x = x.next {
				if !yield(x.key, x.val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (ht *chainHashTable[K, V]) AnyMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (ht *chainHashTable[K, V]) AllMatch(p Predicate2[K, V]) bool {
	for key, val := range ht.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// SelectMatch selects a subset of key-values from the hash table that satisfy the given predicate.
// It returns a new hash table containing the matching key-values, of the same type as the original hash table.
func (ht *chainHashTable[K, V]) SelectMatch(p Predicate2[K, V]) Collection2[K, V] {
	newHT := NewChainHashTable[K, V](ht.hashKey, ht.eqKey, ht.eqVal, HashOpts{
		MinLoadFactor: ht.minLF,
		MaxLoadFactor: ht.maxLF,
	}).(*chainHashTable[K, V])

	for key, val := range ht.All() {
		if p(key, val) {
			newHT.Put(key, val)
		}
	}

	return newHT
}
