// Package set implements a set data structure.
// A set is a collection of objects without any particular order.
package set

import (
	"fmt"
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

// Powerset creates and returns the power set of a given set.
// The power set of a set is the set of all subsets, including the empty set and the set itself.
//
//   - Set[T]:      A set
//   - Set[Set[T]]: The power set (the set of all subsets)
//
// The power set implementation is chosen based on the concrete type of the input set.
// For custom Set implementations, the basic set is used to construct the power set.
func Powerset[T any](s Set[T]) Set[Set[T]] {
	switch s.(type) {
	case *set[T]:
		return powerset(s)
	case *stableSet[T]:
		return stablePowerset(s)
	case *sortedSet[T]:
		return sortedPowerset(s)
	case *hashSet[T]:
		return hashPowerset(s)
	default:
		return powerset(s)
	}
}

// Partitions creates and returns the set of all partitions for a given set.
// A partition of a set is a grouping of its elements into non-empty subsets,
// in such a way that every element is included in exactly one subset.
//
//   - Set[T]:           A set
//   - Set[Set[T]]:      A partition (a set of non-empty disjoint subsets with every element included)
//   - Set[Set[Set[T]]]: The set of all partitions
//
// The partition implementation is chosen based on the concrete type of the input set.
// For custom Set implementations, the basic set is used to construct partitions and the set of all partitions.
func Partitions[T any](s Set[T]) Set[Set[Set[T]]] {
	switch s.(type) {
	case *set[T]:
		return partitions(s)
	case *stableSet[T]:
		return stablePartitions(s)
	case *sortedSet[T]:
		return sortedPartitions(s)
	case *hashSet[T]:
		return hashPartitions(s)
	default:
		return partitions(s)
	}
}
