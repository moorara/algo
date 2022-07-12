package common

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// EqualFunc is a function for checking equality of two values of the same type.
type EqualFunc[T any] func(T, T) bool

// NewEqualFunc creates a new comparator function for standard Go types.
func NewEqualFunc[T comparable]() EqualFunc[T] {
	return func(lhs, rhs T) bool {
		return lhs == rhs
	}
}

// CompareFunc is a function for comparing two values of the same type.
type CompareFunc[T any] func(T, T) int

// NewCompareFunc creates a new comparator function for standard Go types.
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
