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
// The second return value indicates if the intersection is non-empty.
func (r Range[T]) Intersect(rr Range[T]) (Range[T], bool) {
	if !rr.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rr))
	}

	res := Range[T]{
		Lo: max(r.Lo, rr.Lo),
		Hi: min(r.Hi, rr.Hi),
	}

	if res.Valid() {
		return res, true
	}

	return Range[T]{}, false
}

// Subtract returns the subtraction of two discrete ranges.
// The second return value indicates if the subtraction is non-empty.
func (r Range[T]) Subtract(rr Range[T]) []Range[T] {
	if !rr.Valid() {
		panic(fmt.Sprintf("invalid range: %s", rr))
	}

	var res []Range[T]

	left := Range[T]{
		Lo: r.Lo,
		Hi: min(r.Hi, rr.Lo-1),
	}

	if left.Valid() {
		res = append(res, left)
	}

	right := Range[T]{
		Lo: max(r.Lo, rr.Hi+1),
		Hi: r.Hi,
	}

	if right.Valid() {
		res = append(res, right)
	}

	return res
}
