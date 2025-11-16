package set

import (
	"iter"

	"github.com/moorara/algo/generic"
)

// stableSet is an implementation of the Set interface that preserves the order in which members are added.
// This is useful for cases where a deterministic ordering of set members is required.
type stableSet[T any] struct {
	members []T
	equal   generic.EqualFunc[T]
	format  FormatFunc[T]
}

// NewStableSet creates a new set that preserves the order in which members are added.
// This is useful for cases where a deterministic ordering of set members is required.
func NewStableSet[T any](equal generic.EqualFunc[T], vals ...T) Set[T] {
	return NewStableSetWithFormat(equal, defaultFormat[T], vals...)
}

// NewStableSetWithFormat creates a new stable set with a custom format for String method.
func NewStableSetWithFormat[T any](equal generic.EqualFunc[T], format FormatFunc[T], vals ...T) Set[T] {
	s := &stableSet[T]{
		members: make([]T, 0),
		equal:   equal,
		format:  format,
	}

	s.Add(vals...)

	return s
}

func (s *stableSet[T]) find(v T) int {
	for i, m := range s.members {
		if s.equal(m, v) {
			return i
		}
	}

	return -1
}

func (s *stableSet[T]) String() string {
	return s.format(s.members)
}

func (s *stableSet[T]) Equal(rhs Set[T]) bool {
	if s.Size() != rhs.Size() {
		return false
	}

	for _, m := range s.members {
		if !rhs.Contains(m) {
			return false
		}
	}

	return true
}

func (s *stableSet[T]) Clone() Set[T] {
	ss := &stableSet[T]{
		members: make([]T, len(s.members)),
		equal:   s.equal,
		format:  s.format,
	}

	copy(ss.members, s.members)

	return ss
}

func (s *stableSet[T]) CloneEmpty() Set[T] {
	return &stableSet[T]{
		members: make([]T, 0),
		equal:   s.equal,
		format:  s.format,
	}
}

func (s *stableSet[T]) Size() int {
	return len(s.members)
}

func (s *stableSet[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *stableSet[T]) Add(vals ...T) {
	for _, v := range vals {
		if s.find(v) == -1 {
			s.members = append(s.members, v)
		}
	}
}

func (s *stableSet[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.find(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

func (s *stableSet[T]) RemoveAll() {
	s.members = make([]T, 0)
}

func (s *stableSet[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}

	return true
}

func (s *stableSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, m := range s.members {
			if !yield(m) {
				return
			}
		}
	}
}

func (s *stableSet[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *stableSet[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *stableSet[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if p(m) {
			return m, true
		}
	}

	var zero T
	return zero, false
}

func (s *stableSet[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		}
	}

	return matched
}

func (s *stableSet[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
	matched := s.CloneEmpty()
	unmatched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		} else {
			unmatched.Add(m)
		}
	}

	return matched, unmatched
}

func (s *stableSet[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *stableSet[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *stableSet[T]) Union(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Add(m)
		}
	}

	return ss
}

func (s *stableSet[T]) Intersection(sets ...Set[T]) Set[T] {
	ss := s.CloneEmpty()

	for _, m := range s.members {
		isInAll := generic.AllMatch(sets, func(set Set[T]) bool {
			return set.Contains(m)
		})

		if isInAll {
			ss.Add(m)
		}
	}

	return ss
}

func (s *stableSet[T]) Difference(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Remove(m)
		}
	}

	return ss
}

func stablePowerset[T any](s Set[T]) Set[Set[T]] {
	// The power set
	P := NewStableSet(eqStableSet[T])

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
	for subset := range stablePowerset(tail).All() {
		P.Add(subset)             // Add every subset to the power set
		P.Add(head.Union(subset)) // Prepend s[0] to every subset and then add it to the power set
	}

	return P
}

func stablePartitions[T any](s Set[T]) Set[Set[Set[T]]] {
	// The set of all partitions
	Ps := NewStableSet(eqStablePartition[T])

	// Recursion end condition
	if s.Size() == 0 {
		// Single empty partition
		P := NewStableSet(eqStableSet[T])
		Ps.Add(P)
		return Ps
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range stablePartitions(tail).All() {
		Pmembers := generic.Collect1(P.All())

		// Prepend s[0] to the current partition
		Q := NewStableSet(eqStableSet[T])
		Q.Add(head.Clone())
		Q.Add(Pmembers...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the current partition
		for i := range Pmembers {
			Q := NewStableSet(eqStableSet[T])
			Q.Add(Pmembers[0:i]...)
			Q.Add(head.Union(Pmembers[i]))
			Q.Add(Pmembers[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}

func eqStableSet[T any](a, b Set[T]) bool {
	return a.Equal(b)
}

func eqStablePartition[T any](a, b Set[Set[T]]) bool {
	return a.Equal(b)
}
