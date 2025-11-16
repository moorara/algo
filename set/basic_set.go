package set

import (
	"iter"

	"github.com/moorara/algo/generic"
)

// set is an implementation of the Set interface.
// It does not maintain any specific order for its members.
type set[T any] struct {
	members []T
	equal   generic.EqualFunc[T]
	format  FormatFunc[T]
}

// New creates a new set that does not maintain any specific order for its members.
func New[T any](equal generic.EqualFunc[T], vals ...T) Set[T] {
	return NewWithFormat(equal, defaultFormat[T], vals...)
}

// NewWithFormat creates a new set with a custom format for String method.
func NewWithFormat[T any](equal generic.EqualFunc[T], format FormatFunc[T], vals ...T) Set[T] {
	s := &set[T]{
		members: make([]T, 0),
		equal:   equal,
		format:  format,
	}

	s.Add(vals...)

	return s
}

func (s *set[T]) find(v T) int {
	for i, m := range s.members {
		if s.equal(m, v) {
			return i
		}
	}

	return -1
}

func (s *set[T]) String() string {
	return s.format(s.members)
}

func (s *set[T]) Equal(rhs Set[T]) bool {
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

func (s *set[T]) Clone() Set[T] {
	ss := &set[T]{
		members: make([]T, len(s.members)),
		equal:   s.equal,
		format:  s.format,
	}

	copy(ss.members, s.members)

	return ss
}

func (s *set[T]) CloneEmpty() Set[T] {
	return &set[T]{
		members: make([]T, 0),
		equal:   s.equal,
		format:  s.format,
	}
}

func (s *set[T]) Size() int {
	return len(s.members)
}

func (s *set[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *set[T]) Add(vals ...T) {
	for _, v := range vals {
		if s.find(v) == -1 {
			s.members = append(s.members, v)
		}
	}
}

func (s *set[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.find(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

func (s *set[T]) RemoveAll() {
	s.members = make([]T, 0)
}

func (s *set[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}

	return true
}

func (s *set[T]) All() iter.Seq[T] {
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
			if !yield(s.members[i]) {
				return
			}
		}
	}
}

func (s *set[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *set[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if p(m) {
			return m, true
		}
	}

	var zero T
	return zero, false
}

func (s *set[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		}
	}

	return matched
}

func (s *set[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
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

func (s *set[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) Union(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Add(m)
		}
	}

	return ss
}

func (s *set[T]) Intersection(sets ...Set[T]) Set[T] {
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

func (s *set[T]) Difference(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Remove(m)
		}
	}

	return ss
}

func powerset[T any](s Set[T]) Set[Set[T]] {
	// The power set
	P := New(eqSet[T])

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
	for subset := range powerset(tail).All() {
		P.Add(subset)             // Add every subset to the power set
		P.Add(head.Union(subset)) // Prepend s[0] to every subset and then add it to the power set
	}

	return P
}

func partitions[T any](s Set[T]) Set[Set[Set[T]]] {
	// The set of all partitions
	Ps := New(eqPartition[T])

	// Recursion end condition
	if s.Size() == 0 {
		// Single empty partition
		P := New(eqSet[T])
		Ps.Add(P)
		return Ps
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range partitions(tail).All() {
		Pmembers := generic.Collect1(P.All())

		// Prepend s[0] to the current partition
		Q := New(eqSet[T])
		Q.Add(head.Clone())
		Q.Add(Pmembers...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the current partition
		for i := range Pmembers {
			Q := New(eqSet[T])
			Q.Add(Pmembers[0:i]...)
			Q.Add(head.Union(Pmembers[i]))
			Q.Add(Pmembers[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}

func eqSet[T any](a, b Set[T]) bool {
	return a.Equal(b)
}

func eqPartition[T any](a, b Set[Set[T]]) bool {
	return a.Equal(b)
}
