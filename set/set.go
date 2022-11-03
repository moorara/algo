// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import "github.com/moorara/algo/generic"

// Set represents a set abstract data type.
type Set[T any] interface {
	Add(...T)
	Remove(...T)
	IsEmpty() bool
	Contains(T) bool
	Members() []T
	Cardinality() int
	Union(...Set[T]) Set[T]
	Intersection(...Set[T]) Set[T]
	Difference(...Set[T]) Set[T]
}

type set[T any] struct {
	equal   generic.EqualFunc[T]
	members []T
}

// New creates a new empty set
func New[T any](equal generic.EqualFunc[T]) Set[T] {
	return &set[T]{
		equal:   equal,
		members: make([]T, 0),
	}
}

func (s *set[T]) Add(vals ...T) {
	for _, v := range vals {
		if !s.Contains(v) {
			s.members = append(s.members, v)
		}
	}
}

func (s *set[T]) Remove(vals ...T) {
	for _, v := range vals {
		if i := s.contains(v); i != -1 {
			s.members = append(s.members[:i], s.members[i+1:]...)
		}
	}
}

// IsEmpty determines whether or not the set is an empty set.
func (s *set[T]) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *set[T]) Contains(v T) bool {
	return s.contains(v) != -1
}

func (s *set[T]) contains(v T) int {
	for i, member := range s.members {
		if s.equal(member, v) {
			return i
		}
	}

	return -1
}

func (s *set[T]) Members() []T {
	members := make([]T, len(s.members))
	copy(members, s.members)

	return members
}

// Cardinality returns the number of members of the set.
func (s *set[T]) Cardinality() int {
	return len(s.members)
}

func (s *set[T]) Union(sets ...Set[T]) Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, len(s.members)),
	}

	copy(t.members, s.members)

	for _, set := range sets {
		t.Add(set.Members()...)
	}

	return t
}

func (s *set[T]) Intersection(sets ...Set[T]) Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, 0),
	}

	for _, m := range s.Members() {
		isInAll := true
		for _, set := range sets {
			if !set.Contains(m) {
				isInAll = false
				break
			}
		}

		if isInAll {
			t.members = append(t.members, m)
		}
	}

	return t
}

func (s *set[T]) Difference(sets ...Set[T]) Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, len(s.members)),
	}

	copy(t.members, s.members)

	for _, set := range sets {
		t.Remove(set.Members()...)
	}

	return t
}
