// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"fmt"
	"iter"
	"math/rand"
	"time"

	"github.com/moorara/algo/generic"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Set represents a set abstract data type.
type Set[T any] interface {
	fmt.Stringer
	generic.Equaler[Set[T]]
	generic.Cloner[Set[T]]
	generic.Collection1[T]

	CloneEmpty() Set[T]
	IsSubset(Set[T]) bool
	IsSuperset(Set[T]) bool
	Union(...Set[T]) Set[T]
	Intersection(...Set[T]) Set[T]
	Difference(...Set[T]) Set[T]
}

type set[T any] struct {
	members []T
	equal   generic.EqualFunc[T]
	format  StringFormat[T]
}

// New creates a new set.
func New[T any](equal generic.EqualFunc[T], vals ...T) Set[T] {
	s := &set[T]{
		members: make([]T, 0),
		equal:   equal,
		format:  defaultStringFormat[T],
	}

	s.Add(vals...)

	return s
}

// NewWithFormat creates a new set with a custom format for String method.
func NewWithFormat[T any](equal generic.EqualFunc[T], format StringFormat[T], vals ...T) Set[T] {
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
	for _, m := range s.members {
		if !rhs.Contains(m) {
			return false
		}
	}

	for m := range rhs.All() {
		if !s.Contains(m) {
			return false
		}
	}

	return true
}

func (s *set[T]) Clone() Set[T] {
	t := &set[T]{
		members: make([]T, len(s.members)),
		equal:   s.equal,
		format:  s.format,
	}

	copy(t.members, s.members)

	return t
}

func (s *set[T]) CloneEmpty() Set[T] {
	t := &set[T]{
		members: make([]T, 0),
		equal:   s.equal,
		format:  s.format,
	}

	return t
}

func (s *set[T]) Size() int {
	return len(s.members)
}

func (s *set[T]) IsEmpty() bool {
	return len(s.members) == 0
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

	var zeroT T
	return zeroT, false
}

func (s *set[T]) SelectMatch(p generic.Predicate1[T]) generic.Collection1[T] {
	newS := New[T](s.equal)

	for _, m := range s.members {
		if p(m) {
			newS.Add(m)
		}
	}

	return newS
}

func (s *set[T]) PartitionMatch(p generic.Predicate1[T]) (generic.Collection1[T], generic.Collection1[T]) {
	matched := New[T](s.equal)
	unmatched := New[T](s.equal)

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
		members: make([]T, 0),
		equal:   s.equal,
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

// Powerset creates and returns the power set of a given set.
// The power set of a set is the set of all subsets, including the empty set and the set itself.
//
//	Set[T]      A set
//	Set[Set[T]  The power set (the set of all subsets)
func Powerset[T any](s Set[T]) Set[Set[T]] {
	setEqFunc := func(a, b Set[T]) bool { return a.Equal(b) }

	// The power set
	PS := New[Set[T]](setEqFunc)

	// Recurssion end condition
	if s.Size() == 0 {
		// Single empty set
		es := s.CloneEmpty()
		PS.Add(es)
		return PS
	}

	members := generic.Collect1(s.All())
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
	setEqFunc := func(a, b Set[T]) bool { return a.Equal(b) }
	partEqFunc := func(a, b Set[Set[T]]) bool { return a.Equal(b) }

	// The set of all partitions
	Ps := New[Set[Set[T]]](partEqFunc)

	// Recurssion end condition
	if s.Size() == 0 {
		// Single empty partition
		P := New[Set[T]](setEqFunc)
		Ps.Add(P)
		return Ps
	}

	members := generic.Collect1(s.All())
	head, tail := s.CloneEmpty(), s.CloneEmpty()
	head.Add(members[0])
	tail.Add(members[1:]...)

	// For every partition of s[1:]
	for P := range Partitions[T](tail).All() {
		Pmembers := generic.Collect1(P.All())

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
