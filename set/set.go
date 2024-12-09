// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	. "github.com/moorara/algo/generic"
)

// Set represents a set abstract data type.
type Set[T any] interface {
	fmt.Stringer
	Equaler[Set[T]]
	Collection1[T]

	Add(...T)
	Remove(...T)
	Cardinality() int
	IsEmpty() bool
	Contains(...T) bool
	Clone() Set[T]
	CloneEmpty() Set[T]
	Union(...Set[T]) Set[T]
	Intersection(...Set[T]) Set[T]
	Difference(...Set[T]) Set[T]

	Filter(Predicate1[T]) Set[T]
}

type set[T any] struct {
	equal   EqualFunc[T]
	members []T
}

// New creates a new empty set.
func New[T any](equal EqualFunc[T]) Set[T] {
	return &set[T]{
		equal:   equal,
		members: make([]T, 0),
	}
}

func (s *set[T]) find(v T) int {
	for i, m := range s.members {
		if s.equal(m, v) {
			return i
		}
	}

	return -1
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
		if i := s.find(v); i != -1 {
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

func (s *set[T]) Contains(vals ...T) bool {
	for _, v := range vals {
		if s.find(v) == -1 {
			return false
		}
	}
	return true
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
		for m := range set.All() {
			t.Add(m)
		}
	}

	return t
}

func (s *set[T]) Intersection(sets ...Set[T]) Set[T] {
	t := &set[T]{
		equal:   s.equal,
		members: make([]T, 0),
	}

	for _, m := range s.members {
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
		for m := range set.All() {
			t.Remove(m)
		}
	}

	return t
}

func (s *set[T]) String() string {
	elems := make([]string, len(s.members))
	for i, m := range s.members {
		elems[i] = fmt.Sprintf("%v", m)
	}

	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}

func (s *set[T]) Equals(t Set[T]) bool {
	for _, m := range s.members {
		if !t.Contains(m) {
			return false
		}
	}

	for m := range t.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, m := range s.members {
			if !yield(m) {
				break
			}
		}
	}
}

func (s *set[T]) AnyMatch(p Predicate1[T]) bool {
	for _, m := range s.members {
		if p(m) {
			return true
		}
	}

	return false
}

func (s *set[T]) AllMatch(p Predicate1[T]) bool {
	for _, m := range s.members {
		if !p(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) Filter(p Predicate1[T]) Set[T] {
	members := make([]T, 0)
	for _, m := range s.members {
		if p(m) {
			members = append(members, m)
		}
	}

	return &set[T]{
		equal:   s.equal,
		members: members,
	}
}

// Transformer is a function that converts a value of type T to a value of type U.
type Transformer[T, U any] func(T) U

// Transform converts the elements of a set from one type (T) to another type (U).
// It applies the Transformer function (f) to each element of the input set (s) to produce a new set of the transformed type (U).
//
// The caller must provide an equality function (equal) to compare elements of the new set (U).
func (f Transformer[T, U]) Transform(s Set[T], equal EqualFunc[U]) Set[U] {
	members := make([]U, 0)
	for m := range s.All() {
		members = append(members, f(m))
	}

	return &set[U]{
		equal:   equal,
		members: members,
	}
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

	members := slices.Collect(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every subset of s[1:]
	for subset := range Powerset[T](tail).All() {
		PS.Add(subset)             // Add every subset to the power set
		PS.Add(head.Union(subset)) // Prepend s[0] to every subset and then add it to the power set
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

	members := slices.Collect(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range Partitions[T](tail).All() {
		Pmembers := slices.Collect(P.All())

		// Prepend s[0] to the curret partition
		Q := New[Set[T]](setEqFunc)
		Q.Add(head.Clone())
		Q.Add(Pmembers...)
		Ps.Add(Q)

		// Prepend s[0] to every subset in the curret partition
		for i := range Pmembers {
			Q := New[Set[T]](setEqFunc)
			Q.Add(Pmembers[0:i]...)
			Q.Add(head.Union(Pmembers[i]))
			Q.Add(Pmembers[i+1:]...)
			Ps.Add(Q)
		}
	}

	return Ps
}
