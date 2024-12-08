package generic

import "golang.org/x/exp/constraints"

// Min returns the smaller of two values.
// It works with any type T that stisfies the constraints.Ordered constraint,
// which includes types like int, uint, float, and string.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two values.
// It works with any type T that stisfies the constraints.Ordered constraint,
// which includes types like int, uint, float, and string.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
