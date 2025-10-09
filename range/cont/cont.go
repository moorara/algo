// Package cont provides algorithms and data structures for continuous ranges.
package cont

import (
	"fmt"
)

// Continuous represents continuous numerical types.
type Continuous interface {
	~float32 | ~float64
}

// Bound represents a bound in a continuous range.
type Bound[T Continuous] struct {
	Val  T
	Open bool
}

// Range represents a range of continuous values.
type Range[T Continuous] struct {
	Lo Bound[T]
	Hi Bound[T]
}

// Valid determines if the range is valid.
func (r Range[T]) Valid() bool {
	return r.Lo.Val < r.Hi.Val ||
		(r.Lo.Val == r.Hi.Val && !r.Lo.Open && !r.Hi.Open)
}

// String implements the fmt.Stringer interface.
func (r Range[T]) String() string {
	var lo, hi string

	if r.Lo.Open {
		lo = fmt.Sprintf("(%v", r.Lo.Val)
	} else {
		lo = fmt.Sprintf("[%v", r.Lo.Val)
	}

	if r.Hi.Open {
		hi = fmt.Sprintf("%v)", r.Hi.Val)
	} else {
		hi = fmt.Sprintf("%v]", r.Hi.Val)
	}

	return fmt.Sprintf("%s, %s", lo, hi)
}

// Equal implements the generic.Equaler interface.
func (r Range[T]) Equal(rhs Range[T]) bool {
	return r.Lo.Val == rhs.Lo.Val && r.Lo.Open == rhs.Lo.Open &&
		r.Hi.Val == rhs.Hi.Val && r.Hi.Open == rhs.Hi.Open
}

// Adjacent checks if two continuous ranges are adjacent.
// The first return value indicates if r is immediately before rr.
// The second return value indicates if r is immediately after rr.
func (r Range[T]) Adjacent(rr Range[T]) (bool, bool) {
	return r.Hi.Val == rr.Lo.Val && r.Hi.Open != rr.Lo.Open,
		rr.Hi.Val == r.Lo.Val && rr.Hi.Open != r.Lo.Open
}

// Intersect returns the intersection of two continuous ranges.
// The second return value indicates if the intersection is non-empty.
func (r Range[T]) Intersect(rr Range[T]) (Range[T], bool) {
	res := Range[T]{}

	// Determine the low bound (max of low bounds)
	if r.Lo.Val > rr.Lo.Val {
		res.Lo = r.Lo
	} else if r.Lo.Val < rr.Lo.Val {
		res.Lo = rr.Lo
	} else { // r.Lo.Val == rr.Lo.Val
		res.Lo = Bound[T]{
			Val:  r.Lo.Val,
			Open: r.Lo.Open || rr.Lo.Open,
		}
	}

	// Determine the high bound (min of high bounds)
	if r.Hi.Val < rr.Hi.Val {
		res.Hi = r.Hi
	} else if r.Hi.Val > rr.Hi.Val {
		res.Hi = rr.Hi
	} else { // r.Hi.Val == rr.Hi.Val
		res.Hi = Bound[T]{
			Val:  r.Hi.Val,
			Open: r.Hi.Open || rr.Hi.Open,
		}
	}

	if res.Valid() {
		return res, true
	}

	return Range[T]{}, false
}
