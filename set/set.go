// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
)

// Set represents a set abstract data type.
type Set[T any] interface {
	Add(...T)
	Remove(...T)
	Cardinality() int
	IsEmpty() bool
	Contains(T) bool
	Equals(Set[T]) bool
	Members() []T
	Clone() Set[T]
	CloneEmpty() Set[T]
	Union(...Set[T]) Set[T]
	Intersection(...Set[T]) Set[T]
	Difference(...Set[T]) Set[T]
	String() string
}

type set[T any] struct {
	equal   generic.EqualFunc[T]
	members []T
}

// New creates a new empty set.
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

func (s *set[T]) Cardinality() int {
	return len(s.members)
}

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

func (s *set[T]) Equals(t Set[T]) bool {
	for _, m := range s.members {
		if !t.Contains(m) {
			return false
		}
	}

	for _, m := range t.Members() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) Members() []T {
	members := make([]T, len(s.members))
	copy(members, s.members)

	return members
}

func (s *set[T]) Clone() Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, len(s.members)),
	}

	copy(t.members, s.members)

	return t
}

func (s *set[T]) CloneEmpty() Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, 0),
	}

	return t
}

func (s *set[T]) Union(sets ...Set[T]) Set[T] {
	t := s.Clone()

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
	t := s.Clone()

	for _, set := range sets {
		t.Remove(set.Members()...)
	}

	return t
}

func (s *set[T]) String() string {
	strs := make([]string, len(s.members))
	for i, m := range s.members {
		strs[i] = fmt.Sprintf("%v", m)
	}

	return fmt.Sprintf("{%s}", strings.Join(strs, ", "))
}

// Powerset creates and returns the power set of a given set.
// The power set of a set is the set of all subsets, including the empty set and the set itself.
//
//	Set[T]      A set
//	Set[Set[T]  The power set (the set of all subsets)
func Powerset[T any](s Set[T]) Set[Set[T]] {
	setEqFunc := func(a, b Set[T]) bool { return a.Equals(b) }

	// The power set
	PS := New[Set[T]](setEqFunc)

	// Recurssion end condition
	if s.Cardinality() == 0 {
		// Single empty set
		es := s.CloneEmpty()
		PS.Add(es)
		return PS
	}

	mems := s.Members()
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(mems[0])
	tail.Add(mems[1:]...)

	// For every subset of s[1:]
	for _, ss := range Powerset[T](tail).Members() {
		PS.Add(ss)             // Add every subset to the power set
		PS.Add(head.Union(ss)) // Prepend s[0] to every subset and then add it to the power set
	}

	return PS
}

// Partitions creates and returns the set of all partitions for a given set.
// A partition of a set is a grouping of its elements into non-empty subsets,
// in such a way that every element is included in exactly one subset.
//
//	Set[T]            A set
//	Set[Set[T]        A partition (a set of non-empty disjoint subsets with every element included)
//	Set[Set[Set[T]]]  The set of all partitions
func Partitions[T any](s Set[T]) Set[Set[Set[T]]] {
	setEqFunc := func(a, b Set[T]) bool { return a.Equals(b) }
	partEqFunc := func(a, b Set[Set[T]]) bool { return a.Equals(b) }

	// The set of all partitions
	Ps := New[Set[Set[T]]](partEqFunc)

	// Recurssion end condition
	if s.Cardinality() == 0 {
		// Single empty partition
		P := New[Set[T]](setEqFunc)
		Ps.Add(P)
		return Ps
	}

	mems := s.Members()
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(mems[0])
	tail.Add(mems[1:]...)

	// For every partition of s[1:]
	for _, P := range Partitions[T](tail).Members() {
		Pmems := P.Members()

		// Prepend s[0] to the curret partition
		Q := New[Set[T]](setEqFunc)
		Q.Add(head.Clone())
		Q.Add(Pmems...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the curret partition
		for i := range Pmems {
			Q := New[Set[T]](setEqFunc)
			Q.Add(Pmems[0:i]...)
			Q.Add(head.Union(Pmems[i]))
			Q.Add(Pmems[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}
