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
	scMinM          = 4    // Minimum number of buckets in the hash table (must be a power of 2 for efficient hashing)
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

	hashKey hash.HashFunc[K]
	eqKey   generic.EqualFunc[K]
	eqVal   generic.EqualFunc[V]
}

// NewChainHashTable creates a new hash table with separate chaining for conflict resolution.
//
// A hash table is an unordered symbol table providing efficient insertion, deletion, and lookup operations.
// It resolves hash collisions, where multiple keys hash to the same bucket,
// by maintaining a linked list of all key-values that hash to the same bucket.
// Each bucket contains a chain of elements.
func NewChainHashTable[K, V any](hashKey hash.HashFunc[K], eqKey generic.EqualFunc[K], eqVal generic.EqualFunc[V], opts HashOpts) SymbolTable[K, V] {
	if opts.InitialCap < scMinM {
		opts.InitialCap = scMinM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = scMinLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = scMaxLoadFactor
	}

	if M := opts.InitialCap; !math.IsPowerOf2(M) {
		panic("The hash table capacity must be a power of 2")
	}

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
func (t *chainHashTable[K, V]) verify() bool {
	if lf := t.loadFactor(); lf > t.maxLF {
		return false
	}

	// Check that each key in table can be found by Get
	for _, first := range t.buckets {
		for x := first; x != nil; x = x.next {
			if _, ok := t.Get(x.key); !ok {
				return false
			}
		}
	}

	return true
}

// loadFactor calculates the current load factor of the hash table.
// In separate chaining, the load factor can exceed 1.
func (t *chainHashTable[K, V]) loadFactor() float32 {
	return float32(t.n) / float32(t.m)
}

// hash computes the hash for a key and returns an index in [0, M-1] range.
func (t *chainHashTable[K, V]) hash(key K) int {
	h := t.hashKey(key)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	// M must be a power of 2
	M := uint64(t.m)
	h1 := h & (M - 1) // [0, M-1]

	return int(h1)
}

// resize adjusts the hash table to a new size and re-hashes all keys.
func (t *chainHashTable[K, V]) resize(m int) {
	// Ensure the minimum table size
	if m < scMinM {
		return
	}

	new := &chainHashTable[K, V]{
		buckets: make([]*chainNode[K, V], m),
		m:       m,
		n:       0,
		minLF:   t.minLF,
		maxLF:   t.maxLF,
		hashKey: t.hashKey,
		eqKey:   t.eqKey,
		eqVal:   t.eqVal,
	}

	for _, x := range t.buckets {
		for ; x != nil; x = x.next {
			new.Put(x.key, x.val)
		}
	}

	t.buckets = new.buckets
	t.m = new.m
	t.n = new.n
}

// Size returns the number of key-values in the hash table.
func (t *chainHashTable[K, V]) Size() int {
	return t.n
}

// IsEmpty returns true if the hash table is empty.
func (t *chainHashTable[K, V]) IsEmpty() bool {
	return t.n == 0
}

// Put adds a new key-value to the hash table.
func (t *chainHashTable[K, V]) Put(key K, val V) {
	if t.loadFactor() >= t.maxLF {
		t.resize(2 * t.m)
	}

	i := t.hash(key)
	for x := t.buckets[i]; x != nil; x = x.next {
		if t.eqKey(x.key, key) {
			x.val = val
			return
		}
	}

	// Add a new node at the beginning of the chain
	t.buckets[i] = &chainNode[K, V]{
		key:  key,
		val:  val,
		next: t.buckets[i],
	}

	t.n++
}

// Get returns the value of a given key in the hash table.
func (t *chainHashTable[K, V]) Get(key K) (V, bool) {
	i := t.hash(key)
	for x := t.buckets[i]; x != nil; x = x.next {
		if t.eqKey(x.key, key) {
			return x.val, true
		}
	}

	var zeroV V
	return zeroV, false
}

// Delete deletes a key-value from the hash table.
func (t *chainHashTable[K, V]) Delete(key K) (val V, ok bool) {
	i := t.hash(key)
	t.buckets[i], val, ok = t._delete(t.buckets[i], key)

	if t.loadFactor() <= t.minLF {
		t.resize(t.m / 2)
	}

	return val, ok
}

func (t *chainHashTable[K, V]) _delete(n *chainNode[K, V], key K) (*chainNode[K, V], V, bool) {
	if n == nil {
		var zeroV V
		return nil, zeroV, false
	}

	if t.eqKey(n.key, key) {
		t.n--
		return n.next, n.val, true
	}

	var val V
	var ok bool
	n.next, val, ok = t._delete(n.next, key)

	return n, val, ok
}

// DeleteAll deletes all key-values from the hash table, leaving it empty.
func (t *chainHashTable[K, V]) DeleteAll() {
	t.buckets = make([]*chainNode[K, V], t.m)
	t.n = 0
}

// String returns a string representation of the hash table.
func (t *chainHashTable[K, V]) String() string {
	pairs := make([]string, t.Size())
	i := 0

	for key, val := range t.All() {
		pairs[i] = fmt.Sprintf("<%v:%v>", key, val)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(pairs, " "))
}

// Equal determines whether or not two hash tables have the same key-values.
func (t *chainHashTable[K, V]) Equal(rhs SymbolTable[K, V]) bool {
	tt, ok := rhs.(*chainHashTable[K, V])
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
func (t *chainHashTable[K, V]) All() iter.Seq2[K, V] {
	// Create a list of indices representing the buckets.
	indices := make([]int, len(t.buckets))
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
			for x := t.buckets[i]; x != nil; x = x.next {
				if !yield(x.key, x.val) {
					return
				}
			}
		}
	}
}

// AnyMatch returns true if at least one key-value in the hash table satisfies the provided predicate.
func (t *chainHashTable[K, V]) AnyMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if p(key, val) {
			return true
		}
	}
	return false
}

// AllMatch returns true if all key-values in the hash table satisfy the provided predicate.
// If the BST is empty, it returns true.
func (t *chainHashTable[K, V]) AllMatch(p generic.Predicate2[K, V]) bool {
	for key, val := range t.All() {
		if !p(key, val) {
			return false
		}
	}
	return true
}

// FirstMatch returns the first key-value in the hash table that satisfies the given predicate.
// If no match is found, it returns the zero values of K and V, along with false.
func (t *chainHashTable[K, V]) FirstMatch(p generic.Predicate2[K, V]) (K, V, bool) {
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
func (t *chainHashTable[K, V]) SelectMatch(p generic.Predicate2[K, V]) generic.Collection2[K, V] {
	new := NewChainHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
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
func (t *chainHashTable[K, V]) PartitionMatch(p generic.Predicate2[K, V]) (generic.Collection2[K, V], generic.Collection2[K, V]) {
	matched := NewChainHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
		MinLoadFactor: t.minLF,
		MaxLoadFactor: t.maxLF,
	})

	unmatched := NewChainHashTable(t.hashKey, t.eqKey, t.eqVal, HashOpts{
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
