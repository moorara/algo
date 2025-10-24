// Package disc provides algorithms and data structures for discrete ranges.
package disc

import "fmt"

// Discrete represents discrete numerical types.
type Discrete interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Range represents a range of discrete values.
// Discrete range bounds are always inclusive.
type Range[T Discrete] struct {
	Lo T
	Hi T
}

// Valid determines if the range is valid.
func (r Range[T]) Valid() bool {
	return r.Lo <= r.Hi
}

// String implements the fmt.Stringer interface.
func (r Range[T]) String() string {
	return fmt.Sprintf("[%v, %v]", r.Lo, r.Hi)
}

// Equal implements the generic.Equaler interface.
func (r Range[T]) Equal(rhs Range[T]) bool {
	if !rhs.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rhs))
	}

	return r.Lo == rhs.Lo && r.Hi == rhs.Hi
}

// Includes checks if the discrete range includes the given value.
func (r Range[T]) Includes(v T) bool {
	return r.Lo <= v && v <= r.Hi
}

// Adjacent checks if two discrete ranges are adjacent.
// The first return value indicates if r is immediately before rr.
// The second return value indicates if r is immediately after rr.
func (r Range[T]) Adjacent(rr Range[T]) (bool, bool) {
	if !rr.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rr))
	}

	return r.Hi+1 == rr.Lo, rr.Hi+1 == r.Lo
}

// Intersect returns the intersection of two discrete ranges.
func (r Range[T]) Intersect(rr Range[T]) RangeOrEmpty[T] {
	if !rr.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rr))
	}

	res := Range[T]{
		Lo: max(r.Lo, rr.Lo),
		Hi: min(r.Hi, rr.Hi),
	}

	if res.Valid() {
		return RangeOrEmpty[T]{Range: res}
	}

	return RangeOrEmpty[T]{Empty: true}
}

// Subtract returns the subtraction of two discrete ranges.
// It returns two ranges representing the left and right parts of the subtraction.
func (r Range[T]) Subtract(rr Range[T]) (RangeOrEmpty[T], RangeOrEmpty[T]) {
	if !rr.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rr))
	}

	left := Range[T]{
		Lo: r.Lo,
		Hi: min(r.Hi, rr.Lo-1),
	}

	right := Range[T]{
		Lo: max(r.Lo, rr.Hi+1),
		Hi: r.Hi,
	}

	if left.Valid() && right.Valid() {
		return RangeOrEmpty[T]{Range: left}, RangeOrEmpty[T]{Range: right}
	} else if left.Valid() {
		return RangeOrEmpty[T]{Range: left}, RangeOrEmpty[T]{Empty: true}
	} else if right.Valid() {
		return RangeOrEmpty[T]{Empty: true}, RangeOrEmpty[T]{Range: right}
	}

	return RangeOrEmpty[T]{Empty: true}, RangeOrEmpty[T]{Empty: true}
}

// RangeOrEmpty represents a discrete range that can be empty.
type RangeOrEmpty[T Discrete] struct {
	Range[T]
	Empty bool
}

// EqRange compares two discrete ranges for equality.
func EqRange[T Discrete](lhs, rhs Range[T]) bool {
	return lhs.Equal(rhs)
}
