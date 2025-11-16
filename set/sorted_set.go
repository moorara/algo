package set

import (
	"iter"

	"github.com/moorara/algo/generic"
)

// sortedSet is an implementation of the Set interface that always keeps its members sorted.
// This is useful for cases where a sorted ordering of set members is required.
type sortedSet[T any] struct {
	members []T
	compare generic.CompareFunc[T]
	format  FormatFunc[T]
}

// NewSortedSet creates a new set that maintains its members in sorted order.
// This is useful for cases where a sorted ordering of set members is required.
func NewSortedSet[T any](compare generic.CompareFunc[T], vals ...T) Set[T] {
	return NewSortedSetWithFormat(compare, defaultFormat[T], vals...)
}

// NewSortedSetWithFormat creates a new sorted set with a custom format for String method.
func NewSortedSetWithFormat[T any](compare generic.CompareFunc[T], format FormatFunc[T], vals ...T) Set[T] {
	s := &sortedSet[T]{
		members: make([]T, 0),
		compare: compare,
		format:  format,
	}

	s.Add(vals...)

	return s
}

// find implements a binary search.
func (s *sortedSet[T]) find(v T) int {
	low, high := 0, len(s.members)-1

	for low <= high {
		mid := (low + high) / 2
		cmp := s.compare(v, s.members[mid])

		if cmp < 0 {
			high = mid - 1 // Search in the left half
		} else if cmp > 0 {
			low = mid + 1 // Search in the right half
		} else {
			return mid // Found target
		}
	}

	return -1
}

func (s *sortedSet[T]) String() string {
	return s.format(s.members)
}

func (s *sortedSet[T]) Equal(rhs Set[T]) bool {
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

func (s *sortedSet[T]) Compare(rhs *sortedSet[T]) int {
	for i := 0; i < len(s.members) && i < len(rhs.members); i++ {
		if c := s.compare(s.members[i], rhs.members[i]); c != 0 {
			return c
		}
	}

	return len(s.members) - len(rhs.members)
}

func (s *sortedSet[T]) Clone() Set[T] {
	ss := &sortedSet[T]{
		members: make([]T, len(s.members)),
		compare: s.compare,
		format:  s.format,
	}

	copy(ss.members, s.members)

	return ss
}

func (s *sortedSet[T]) CloneEmpty() Set[T] {
	return &sortedSet[T]{
		members: make([]T, 0),
		compare: s.compare,
		format:  s.format,
	}
}

func (s *sortedSet[T]) Size() int {
	return len(s.members)
}

func (s *sortedSet[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *sortedSet[T]) Add(vals ...T) {
	for _, v := range vals {
		s.add(v)
	}
}

func (s *sortedSet[T]) add(val T) {
	low, high := 0, len(s.members)-1

	for low <= high {
		mid := (low + high) / 2
		cmp := s.compare(val, s.members[mid])

		if cmp < 0 {
			high = mid - 1 // Search in the left half
		} else if cmp > 0 {
			low = mid + 1 // Search in the right half
		} else {
			return // Member already exists
		}
	}

	// Insert at index low
	s.members = append(s.members[:low], append([]T{val}, s.members[low:]...)...)
}

func (s *sortedSet[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.find(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

func (s *sortedSet[T]) RemoveAll() {
	s.members = make([]T, 0)
}

func (s *sortedSet[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}

	return true
}

func (s *sortedSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, m := range s.members {
			if !yield(m) {
				return
			}
		}
	}
}

func (s *sortedSet[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *sortedSet[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *sortedSet[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if p(m) {
			return m, true
		}
	}

	var zero T
	return zero, false
}

func (s *sortedSet[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		}
	}

	return matched
}

func (s *sortedSet[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
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

func (s *sortedSet[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *sortedSet[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *sortedSet[T]) Union(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Add(m)
		}
	}

	return ss
}

func (s *sortedSet[T]) Intersection(sets ...Set[T]) Set[T] {
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

func (s *sortedSet[T]) Difference(sets ...Set[T]) Set[T] {
	ss := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			ss.Remove(m)
		}
	}

	return ss
}

func sortedPowerset[T any](s Set[T]) Set[Set[T]] {
	// The power set
	P := NewSortedSet(cmpSortedSet[T])

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
	for subset := range sortedPowerset(tail).All() {
		P.Add(subset)             // Add every subset to the power set
		P.Add(head.Union(subset)) // Prepend s[0] to every subset and then add it to the power set
	}

	return P
}

func sortedPartitions[T any](s Set[T]) Set[Set[Set[T]]] {
	// The set of all partitions
	Ps := NewSortedSet(cmpSortedPartition[T])

	// Recursion end condition
	if s.Size() == 0 {
		// Single empty partition
		P := NewSortedSet(cmpSortedSet[T])
		Ps.Add(P)
		return Ps
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range sortedPartitions(tail).All() {
		Pmembers := generic.Collect1(P.All())

		// Prepend s[0] to the current partition
		Q := NewSortedSet(cmpSortedSet[T])
		Q.Add(head.Clone())
		Q.Add(Pmembers...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the current partition
		for i := range Pmembers {
			Q := NewSortedSet(cmpSortedSet[T])
			Q.Add(Pmembers[0:i]...)
			Q.Add(head.Union(Pmembers[i]))
			Q.Add(Pmembers[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}

func cmpSortedSet[T any](a, b Set[T]) int {
	aa, bb := a.(*sortedSet[T]), b.(*sortedSet[T])
	return aa.Compare(bb)
}

func cmpSortedPartition[T any](a, b Set[Set[T]]) int {
	aa, bb := a.(*sortedSet[Set[T]]), b.(*sortedSet[Set[T]])
	return aa.Compare(bb)
}
