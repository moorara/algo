// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"iter"

	"github.com/moorara/algo/generic"
)

// sorted is an implementation of the Set interface that always keeps its members sorted.
// This is useful for cases where a sorted ordering of set members is required.
type sorted[T any] struct {
	members []T
	compare generic.CompareFunc[T]
	format  FormatFunc[T]
}

// NewSorted creates a new set that maintains its members in sorted order.
// This is useful for cases where a sorted ordering of set members is required.
func NewSorted[T any](compare generic.CompareFunc[T], vals ...T) Set[T] {
	s := &sorted[T]{
		members: make([]T, 0),
		compare: compare,
		format:  defaultFormat[T],
	}

	s.Add(vals...)

	return s
}

// NewSortedWithFormat creates a new sorted set with a custom format for String method.
func NewSortedWithFormat[T any](compare generic.CompareFunc[T], format FormatFunc[T], vals ...T) Set[T] {
	s := &sorted[T]{
		members: make([]T, 0),
		compare: compare,
		format:  format,
	}

	s.Add(vals...)

	return s
}

// find implements a binary search.
func (s *sorted[T]) find(v T) int {
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

func (s *sorted[T]) String() string {
	return s.format(s.members)
}

func (s *sorted[T]) Equal(rhs Set[T]) bool {
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

func (s *sorted[T]) Clone() Set[T] {
	t := &sorted[T]{
		members: make([]T, len(s.members)),
		compare: s.compare,
		format:  s.format,
	}

	copy(t.members, s.members)

	return t
}

func (s *sorted[T]) CloneEmpty() Set[T] {
	return &sorted[T]{
		members: make([]T, 0),
		compare: s.compare,
		format:  s.format,
	}
}

func (s *sorted[T]) Size() int {
	return len(s.members)
}

func (s *sorted[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *sorted[T]) Add(vals ...T) {
	for _, v := range vals {
		s.add(v)
	}
}

func (s *sorted[T]) add(val T) {
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

func (s *sorted[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.find(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

func (s *sorted[T]) RemoveAll() {
	s.members = make([]T, 0)
}

func (s *sorted[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}

	return true
}

func (s *sorted[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, m := range s.members {
			if !yield(m) {
				return
			}
		}
	}
}

func (s *sorted[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *sorted[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *sorted[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if p(m) {
			return m, true
		}
	}

	var zeroT T
	return zeroT, false
}

func (s *sorted[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		}
	}

	return matched
}

func (s *sorted[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
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

func (s *sorted[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *sorted[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *sorted[T]) Union(sets ...Set[T]) Set[T] {
	t := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			t.Add(m)
		}
	}

	return t
}

func (s *sorted[T]) Intersection(sets ...Set[T]) Set[T] {
	t := s.CloneEmpty()

	for _, m := range s.members {
		isInAll := generic.AllMatch(sets, func(set Set[T]) bool {
			return set.Contains(m)
		})

		if isInAll {
			t.Add(m)
		}
	}

	return t
}

func (s *sorted[T]) Difference(sets ...Set[T]) Set[T] {
	t := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			t.Remove(m)
		}
	}

	return t
}
