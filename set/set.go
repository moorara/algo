// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import "github.com/moorara/algo/common"

// Set represents a set abstract data type.
type Set[T any] interface {
	Add(...T)
	Remove(...T)
	Contains(T) bool
	Members() []T
	IsEmpty() bool
	Cardinality() int
}

type set[T any] struct {
	equal common.EqualFunc[T]

	members []T
}

// New creates a new empty set
func New[T any](equal common.EqualFunc[T]) Set[T] {
	return &set[T]{
		equal: equal,

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
	return s.members
}

// IsEmpty determines whether or not the set is an empty set.
func (s *set[T]) IsEmpty() bool {
	return len(s.members) == 0
}

// Cardinality returns the number of members of the set.
func (s *set[T]) Cardinality() int {
	return len(s.members)
}
