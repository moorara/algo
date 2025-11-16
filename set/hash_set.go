package set

import (
	"iter"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/math"
)

const (
	minM          = 31    // Minimum capacity of the hash set (must be a prime number)
	minLoadFactor = 0.125 // Minimum load factor before resizing (shrinking)
	maxLoadFactor = 0.50  // Maximum load factor before resizing (expanding)
)

// HashSetOpts represents configuration options for a hash set.
type HashSetOpts struct {
	// The initial capacity of the hash set.
	// It must be a prime number for better distribution.
	InitialCap int
	// The minimum load factor before resizing (shrinking) the hash set.
	MinLoadFactor float32
	// The maximum load factor before resizing (expanding) the hash set.
	MaxLoadFactor float32
}

// hashSet is an implementation of the Set interface backed by a hash data structure for faster operations.
// It is a trade-off between time and space (memory usage),
// suited for large sets with frequent operations where performance is critical.
type hashSet[T any] struct {
	members []*hashSetMember[T]
	m       int     // The capacity of the hash set
	n       int     // The current number of members stored in the hash set
	minLF   float32 // The minimum load factor before resizing (shrinking) the hash set
	maxLF   float32 // The maximum load factor before resizing (expanding) the hash set
	hash    hash.HashFunc[T]
	equal   generic.EqualFunc[T]
	format  FormatFunc[T]
}

// hashSetMember represents a member in the hash set.
type hashSetMember[T any] struct {
	val     T
	deleted bool
}

// NewHashSet creates a new set backed by a hash data structure for faster operations.
// It is a trade-off between time and space (memory usage),
// suited for large sets with frequent operations where performance is critical.
func NewHashSet[T any](hash hash.HashFunc[T], equal generic.EqualFunc[T], opts HashSetOpts, vals ...T) Set[T] {
	return NewHashSetWithFormat(hash, equal, defaultFormat[T], opts, vals...)
}

// NewHashSetWithFormat creates a new hash set with a custom format for String method.
func NewHashSetWithFormat[T any](hash hash.HashFunc[T], equal generic.EqualFunc[T], format FormatFunc[T], opts HashSetOpts, vals ...T) Set[T] {
	if opts.InitialCap < minM {
		opts.InitialCap = minM
	}

	if opts.MinLoadFactor == 0 {
		opts.MinLoadFactor = minLoadFactor
	}

	if opts.MaxLoadFactor == 0 {
		opts.MaxLoadFactor = maxLoadFactor
	}

	if M := opts.InitialCap; !math.IsPrime(M) {
		panic("The hash set capacity must be a prime number")
	}

	s := &hashSet[T]{
		members: make([]*hashSetMember[T], opts.InitialCap),
		m:       opts.InitialCap,
		n:       0,
		minLF:   opts.MinLoadFactor,
		maxLF:   opts.MaxLoadFactor,
		hash:    hash,
		equal:   equal,
		format:  format,
	}

	s.Add(vals...)

	return s
}

// loadFactor calculates the current load factor of the hash set.
// The load factor is between 0 and 1.
func (s *hashSet[T]) loadFactor() float32 {
	return float32(s.n) / float32(s.m)
}

// probe returns a function that generates the next index in a quadratic probing sequence.
// The sequence starts at h and increments quadratically: h, h+1², h+2², h+3², ...
func (s *hashSet[T]) probe(val T) func() int {
	h := s.hash(val)
	h ^= (h >> 20) ^ (h >> 12) ^ (h >> 7) ^ (h >> 4)

	M := uint64(s.m)
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

// resize adjusts the hash set to a new size and re-hashes all members.
func (s *hashSet[T]) resize(m int) {
	// Ensure the minimum set size
	if m < minM {
		return
	}

	// Ensure the hash set size remains prime
	m = math.SmallestPrimeLargerThan(m)

	new := &hashSet[T]{
		members: make([]*hashSetMember[T], m),
		m:       m,
		n:       0,
		minLF:   s.minLF,
		maxLF:   s.maxLF,
		hash:    s.hash,
		equal:   s.equal,
		format:  s.format,
	}

	for _, m := range s.members {
		if m != nil && !m.deleted {
			new.Add(m.val)
		}
	}

	// Replace the current set with the resized set
	s.members = new.members
	s.m = new.m
	s.n = new.n
}

func (s *hashSet[T]) String() string {
	members := make([]T, 0, s.n)
	for _, m := range s.members {
		if m != nil && !m.deleted {
			members = append(members, m.val)
		}
	}

	return s.format(members)
}

func (s *hashSet[T]) Equal(rhs Set[T]) bool {
	if s.Size() != rhs.Size() {
		return false
	}

	for _, m := range s.members {
		if m != nil && !m.deleted {
			if !rhs.Contains(m.val) {
				return false
			}
		}
	}

	return true
}

func (s *hashSet[T]) Clone() Set[T] {
	ss := &hashSet[T]{
		members: make([]*hashSetMember[T], s.m),
		m:       s.m,
		n:       s.n,
		minLF:   s.minLF,
		maxLF:   s.maxLF,
		hash:    s.hash,
		equal:   s.equal,
		format:  s.format,
	}

	for i, m := range s.members {
		if m != nil {
			ss.members[i] = &hashSetMember[T]{
				val:     m.val,
				deleted: m.deleted,
			}
		}
	}

	return ss
}

func (s *hashSet[T]) CloneEmpty() Set[T] {
	ss := &hashSet[T]{
		members: make([]*hashSetMember[T], s.m),
		m:       s.m,
		n:       0,
		minLF:   s.minLF,
		maxLF:   s.maxLF,
		hash:    s.hash,
		equal:   s.equal,
		format:  s.format,
	}

	return ss
}

func (s *hashSet[T]) Size() int {
	return s.n
}

func (s *hashSet[T]) IsEmpty() bool {
	return s.n == 0
}

func (s *hashSet[T]) Add(vals ...T) {
	for _, v := range vals {
		s.add(v)
	}
}

func (s *hashSet[T]) add(val T) {
	if s.loadFactor() >= s.maxLF {
		s.resize(2 * s.m)
	}

	var i int
	next := s.probe(val)
	for i = next(); s.members[i] != nil; i = next() {
		if s.equal(s.members[i].val, val) {
			if s.members[i].deleted {
				s.members[i].deleted = false
				s.n++
			}
			return
		}
	}

	s.members[i] = &hashSetMember[T]{
		val:     val,
		deleted: false,
	}

	s.n++
}

func (s *hashSet[T]) Remove(vals ...T) {
	for _, v := range vals {
		s.remove(v)
	}
}

func (s *hashSet[T]) remove(val T) {
	next := s.probe(val)
	i := next()
	for s.members[i] != nil && !s.equal(s.members[i].val, val) {
		i = next()
	}

	// Member not found
	if s.members[i] == nil || s.members[i].deleted {
		return
	}

	// Remove the entry from the hash set
	s.members[i].deleted = true
	s.n--

	// During resizing, soft-deleted members are removed, and remaining active members are rehashed
	if s.loadFactor() <= s.minLF {
		s.resize(s.m / 2)
	}
}

func (s *hashSet[T]) RemoveAll() {
	s.members = make([]*hashSetMember[T], s.m)
	s.n = 0
}

func (s *hashSet[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if !s.contains(v) {
			return false
		}
	}

	return true
}

func (s *hashSet[T]) contains(val T) bool {
	next := s.probe(val)
	for i := next(); s.members[i] != nil; i = next() {
		if !s.members[i].deleted && s.equal(s.members[i].val, val) {
			return true
		}
	}

	return false
}

func (s *hashSet[T]) All() iter.Seq[T] {
	// Create a list of indices representing the members.
	indices := make([]int, len(s.members))
	for i := range indices {
		indices[i] = i
	}

	// Shuffle the indices list to randomize the order in which members are traversed.
	// This ensures that the traversal order is non-deterministic, reflecting the unordered nature of set.
	r.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	return func(yield func(T) bool) {
		for _, i := range indices {
			if m := s.members[i]; m != nil && !m.deleted {
				if !yield(m.val) {
					return
				}
			}
		}
	}
}

func (s *hashSet[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if m != nil && !m.deleted {
			if p(m.val) {
				return true
			}
		}
	}

	return false
}

func (s *hashSet[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if m != nil && !m.deleted {
			if !p(m.val) {
				return false
			}
		}
	}

	return true
}

func (s *hashSet[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if m != nil && !m.deleted {
			if p(m.val) {
				return m.val, true
			}
		}
	}

	var zero T
	return zero, false
}

func (s *hashSet[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if m != nil && !m.deleted {
			if p(m.val) {
				matched.Add(m.val)
			}
		}
	}

	return matched
}

func (s *hashSet[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
	matched := s.CloneEmpty()
	unmatched := s.CloneEmpty()

	for _, m := range s.members {
		if m != nil && !m.deleted {
			if p(m.val) {
				matched.Add(m.val)
			} else {
				unmatched.Add(m.val)
			}
		}
	}

	return matched, unmatched
}

func (s *hashSet[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *hashSet[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *hashSet[T]) Union(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Add(m)
		}
	}

	return ss
}

func (s *hashSet[T]) Intersection(sets ...Set[T]) Set[T] {
	ss := s.CloneEmpty()

	for _, m := range s.members {
		if m != nil && !m.deleted {
			isInAll := generic.AllMatch(sets, func(set Set[T]) bool {
				return set.Contains(m.val)
			})

			if isInAll {
				ss.Add(m.val)
			}
		}
	}

	return ss
}

func (s *hashSet[T]) Difference(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Remove(m)
		}
	}

	return ss
}

func hashPowerset[T any](s Set[T]) Set[Set[T]] {
	// The power set
	P := NewHashSet(hashHashSet[T], eqHashSet[T], HashSetOpts{})

	// Recursion end condition
	if s.Size() == 0 {
		// Single empty set
		es := s.CloneEmpty()
		P.Add(es)
		return P
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every subset of s[1:]
	for subset := range hashPowerset(tail).All() {
		P.Add(subset)             // Add every subset to the power set
		P.Add(head.Union(subset)) // Prepend s[0] to every subset and then add it to the power set
	}

	return P
}

func hashPartitions[T any](s Set[T]) Set[Set[Set[T]]] {
	// The set of all partitions
	Ps := NewHashSet(hashHashPartition[T], eqHashPartition[T], HashSetOpts{})

	// Recursion end condition
	if s.Size() == 0 {
		// Single empty partition
		P := NewHashSet(hashHashSet[T], eqHashSet[T], HashSetOpts{})
		Ps.Add(P)
		return Ps
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range hashPartitions(tail).All() {
		Pmembers := generic.Collect1(P.All())

		// Prepend s[0] to the current partition
		Q := NewHashSet(hashHashSet[T], eqHashSet[T], HashSetOpts{})
		Q.Add(head.Clone())
		Q.Add(Pmembers...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the current partition
		for i := range Pmembers {
			Q := NewHashSet(hashHashSet[T], eqHashSet[T], HashSetOpts{})
			Q.Add(Pmembers[0:i]...)
			Q.Add(head.Union(Pmembers[i]))
			Q.Add(Pmembers[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}

func hashHashSet[T any](s Set[T]) uint64 {
	ss := s.(*hashSet[T])

	// Use XOR which is commutative and associative, so insertion order does not matter.
	var h uint64
	for _, m := range ss.members {
		if m != nil && !m.deleted {
			h ^= ss.hash(m.val)
		}
	}

	return h
}

func eqHashSet[T any](a, b Set[T]) bool {
	return a.Equal(b)
}

func hashHashPartition[T any](P Set[Set[T]]) uint64 {
	PP := P.(*hashSet[Set[T]])

	// Use XOR which is commutative and associative, so insertion order does not matter.
	var h uint64
	for _, m := range PP.members {
		if m != nil && !m.deleted {
			h ^= PP.hash(m.val)
		}
	}

	return h
}

func eqHashPartition[T any](a, b Set[Set[T]]) bool {
	return a.Equal(b)
}
