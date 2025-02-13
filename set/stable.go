// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"iter"

	"github.com/moorara/algo/generic"
)

// stable is an implementation of the Set interface that preserves the order in which members are added.
// This is useful for cases where a deterministic ordering of set members is required.
type stable[T any] struct {
	members []T
	equal   generic.EqualFunc[T]
	format  StringFormat[T]
}

// NewStable creates a new set that preserves the order in which members are added.
// This is useful for cases where a deterministic ordering of set members is required.
func NewStable[T any](equal generic.EqualFunc[T], vals ...T) Set[T] {
	s := &stable[T]{
		members: make([]T, 0),
		equal:   equal,
		format:  defaultStringFormat[T],
	}

	s.Add(vals...)

	return s
}

// NewStableWithFormat creates a new stable set with a custom format for String method.
func NewStableWithFormat[T any](equal generic.EqualFunc[T], format StringFormat[T], vals ...T) Set[T] {
	s := &stable[T]{
		members: make([]T, 0),
		equal:   equal,
		format:  format,
	}

	s.Add(vals...)

	return s
}

func (s *stable[T]) find(v T) int {
	for i, m := range s.members {
		if s.equal(m, v) {
			return i
		}
	}

	return -1
}

func (s *stable[T]) String() string {
	return s.format(s.members)
}

func (s *stable[T]) Equal(rhs Set[T]) bool {
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

func (s *stable[T]) Clone() Set[T] {
	t := &stable[T]{
		members: make([]T, len(s.members)),
		equal:   s.equal,
		format:  s.format,
	}

	copy(t.members, s.members)

	return t
}

func (s *stable[T]) CloneEmpty() Set[T] {
	return &stable[T]{
		members: make([]T, 0),
		equal:   s.equal,
		format:  s.format,
	}
}

func (s *stable[T]) Size() int {
	return len(s.members)
}

func (s *stable[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *stable[T]) Add(vals ...T) {
	for _, v := range vals {
		if !s.Contains(v) {
			s.members = append(s.members, v)
		}
	}
}

func (s *stable[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.find(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

func (s *stable[T]) RemoveAll() {
	s.members = make([]T, 0)
}

func (s *stable[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}

	return true
}

func (s *stable[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, m := range s.members {
			if !yield(m) {
				return
			}
		}
	}
}

func (s *stable[T]) AnyMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *stable[T]) AllMatch(p generic.Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *stable[T]) FirstMatch(p generic.Predicate1[T]) (T, bool) {
	for _, m := range s.members {
		if p(m) {
			return m, true
		}
	}

	var zeroT T
	return zeroT, false
}

func (s *stable[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	matched := s.CloneEmpty()

	for _, m := range s.members {
		if p(m) {
			matched.Add(m)
		}
	}

	return matched
}

func (s *stable[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
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

func (s *stable[T]) IsSubset(superset Set[T]) bool {
	for m := range s.All() {
		if !superset.Contains(m) {
			return false
		}
	}

	return true
}

func (s *stable[T]) IsSuperset(subset Set[T]) bool {
	for m := range subset.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *stable[T]) Union(sets ...Set[T]) Set[T] {
	t := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			t.Add(m)
		}
	}

	return t
}

func (s *stable[T]) Intersection(sets ...Set[T]) Set[T] {
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

func (s *stable[T]) Difference(sets ...Set[T]) Set[T] {
	t := s.Clone()

	for _, set := range sets {
		for m := range set.All() {
			t.Remove(m)
		}
	}

	return t
}
