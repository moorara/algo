// Package generic provides types, interfaces, and functions to support generic programming use cases.
package generic

import "golang.org/x/exp/constraints"

// Cloner is an interface that defines a method for cloning an object (the prototype pattern).
type Cloner[T any] interface {
	Clone() T
}

// Equaler is a generic interface for determining equality two objects of the same type.
type Equaler[T any] interface {
	Equal(T) bool
}

// EqualFunc defines a generic function type for checking equality between two values of the same type.
// The function takes two arguments of type T and returns true if they are considered equal, or false otherwise.
type EqualFunc[T any] func(T, T) bool

// NewEqualFunc returns a generic equality function for any type that satisfies the comparable constraint.
func NewEqualFunc[T comparable]() EqualFunc[T] {
	return func(lhs, rhs T) bool {
		return lhs == rhs
	}
}

// Comparer is a generic interface for comparing two objects of the same type
// and establishing an order between them.
// The Compare method returns a negative value if the current object is less than the given object,
// zero if they are equal, and a positive value if the current object is greater.
type Comparer[T any] interface {
	Compare(T) int
}

// CompareFunc defines a generic function type for comparing two values of the same type.
// The function takes two arguments of type T and returns:
//   - A negative integer if the first value is less than the second,
//   - Zero if the two values are equal,
//   - A positive integer if the first value is greater than the second.
type CompareFunc[T any] func(T, T) int

// NewCompareFunc returns a generic comparison function
// for any type that satisfies the constraints.Ordered interface.
// The returned function compares two values of type T and returns:
//   - -1 if lhs is less than rhs,
//   - 1 if lhs is greater than rhs,
//   - 0 if lhs is equal to rhs.
//
// This is useful for implementing custom sorting or comparison logic.
func NewCompareFunc[T constraints.Ordered]() CompareFunc[T] {
	return func(lhs, rhs T) int {
		switch {
		case lhs < rhs:
			return -1
		case lhs > rhs:
			return 1
		default:
			return 0
		}
	}
}

// NewReverseCompareFunc returns a generic reverse comparison function
// for any type that satisfies the constraints.Ordered interface.
// The returned function compares two values of type T and returns:
//   - 1 if lhs is less than rhs,
//   - -1 if lhs is greater than rhs,
//   - 0 if lhs is equal to rhs.
//
// This is useful for implementing reverse sorting or inverted comparison logic.
func NewReverseCompareFunc[T constraints.Ordered]() CompareFunc[T] {
	return func(lhs, rhs T) int {
		switch {
		case lhs < rhs:
			return 1
		case lhs > rhs:
			return -1
		default:
			return 0
		}
	}
}
