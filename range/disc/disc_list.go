package disc

import (
	"fmt"
	"iter"
	"slices"
)

// RangeList represents a list of discrete ranges.
// The ranges are always non-overlapping and sorted.
type RangeList[T Discrete] struct {
	ranges []Range[T]
	format FormatList[T]
}

// NewRangeList creates a new range list from the given ranges.
// It panics if any of the given ranges are invalid.
func NewRangeList[T Discrete](rs ...Range[T]) *RangeList[T] {
	for _, r := range rs {
		if !r.Valid() {
			panic(fmt.Sprintf("invalid range: %s", r))
		}
	}

	l := &RangeList[T]{
		ranges: rs,
		format: defaultFormatList[T],
	}

	// Sort ranges by their low bound ascending
	slices.SortFunc(l.ranges, func(lhs, rhs Range[T]) int {
		return int(lhs.Lo) - int(rhs.Lo)
	})

	// Merge overlapping and adjacent ranges
	l.mergeRanges()

	return l
}

// NewRangeListWithFormat creates a new range list with a custom format function from the given ranges.
// It panics if any of the given ranges are invalid.
func NewRangeListWithFormat[T Discrete](format FormatList[T], rs ...Range[T]) *RangeList[T] {
	l := NewRangeList(rs...)
	l.format = format

	return l
}

// searchRanges performs a binary search to find the index of the range that contains the given value.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func (l *RangeList[T]) searchRanges(v T) (int, bool) {
	lo, hi := 0, len(l.ranges)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if v < l.ranges[mid].Lo {
			hi = mid - 1
		} else if l.ranges[mid].Hi < v {
			lo = mid + 1
		} else {
			return mid, true
		}
	}

	return lo, false
}

// mergeRanges merges overlapping or adjacent ranges in the sorted list of ranges.
func (l *RangeList[T]) mergeRanges() {
	merged := make([]Range[T], 0, len(l.ranges))

	for _, curr := range l.ranges {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if curr.Lo <= last.Hi {
			if curr.Hi > last.Hi {
				// Case curr.Lo < last.Hi && curr.Hi > last.Hi
				//
				//   last:  |_____|_____|     |    ---->    |_________________|
				//   curr:        |___________|    ---->
				//
				// Case curr.Lo == last.Hi && curr.Hi > last.Hi
				//
				//   last:  |___________|     |    ---->    |_________________|
				//   curr:              |_____|    ---->
				//

				last.Hi = curr.Hi
			}
		} else if before, _ := last.Adjacent(curr); before {
			// Case last.Hi is immediately before curr.Lo
			//
			//   last:  |__________||     |    ---->    |_________________|
			//   curr:              |_____|    ---->
			//

			last.Hi = curr.Hi
		} else {
			merged = append(merged, curr)
		}
	}

	l.ranges = merged
}

// String implements the fmt.Stringer interface.
func (l *RangeList[T]) String() string {
	return l.format(l.All())
}

// Clone implements the generic.Cloner interface.
func (l *RangeList[T]) Clone() *RangeList[T] {
	ll := &RangeList[T]{
		ranges: make([]Range[T], len(l.ranges)),
	}

	copy(ll.ranges, l.ranges)

	return ll
}

// Equal implements the generic.Equaler interface.
func (l *RangeList[T]) Equal(rhs *RangeList[T]) bool {
	if len(l.ranges) != len(rhs.ranges) {
		return false
	}

	for i, r := range l.ranges {
		if !r.Equal(rhs.ranges[i]) {
			return false
		}
	}

	return true
}

// Size returns the number of ranges in the range list.
func (l *RangeList[T]) Size() int {
	return len(l.ranges)
}

// Get returns the range that includes the given value.
// The second return value indicates if such a range exists.
func (l *RangeList[T]) Get(v T) (Range[T], bool) {
	if i, ok := l.searchRanges(v); ok {
		return l.ranges[i], true
	}

	return Range[T]{}, false
}

// Add inserts the given ranges to the range list.
// It panics if any of the given ranges are invalid.
func (l *RangeList[T]) Add(rs ...Range[T]) {
	for _, r := range rs {
		if !r.Valid() {
			panic(fmt.Sprintf("invalid range: %s", r))
		}

		// Find the insertion point
		i, ok := l.searchRanges(r.Lo)
		if ok {
			i++
		}

		// Insert the new entry at position i
		l.ranges = append(l.ranges, Range[T]{})
		copy(l.ranges[i+1:], l.ranges[i:])
		l.ranges[i] = r

		// Merge overlapping and adjacent ranges
		l.mergeRanges()
	}
}

// Remove deletes the given ranges from the range list.
// It panics if any of the given ranges are invalid.
func (l *RangeList[T]) Remove(rs ...Range[T]) {
	for _, r := range rs {
		if !r.Valid() {
			panic(fmt.Sprintf("invalid range: %s", r))
		}

		i, _ := l.searchRanges(r.Lo)

		for i < len(l.ranges) {
			left, right := l.ranges[i].Subtract(r)

			if !left.Empty && !right.Empty {
				// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
				//
				//   |______________|
				//        |____|
				//   |___|      |___|
				//     l          r
				//

				l.ranges[i] = left.Range
				l.ranges = append(l.ranges, Range[T]{})
				copy(l.ranges[i+2:], l.ranges[i+1:])
				l.ranges[i+1] = right.Range
				break
			} else if !left.Empty {
				// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
				//
				//   |____________|        |____________|              |____________|
				//          |_____|               |___________|                     |_____|
				//   |_____|               |_____|                     |___________|
				//      l                     l                              l
				//

				l.ranges[i] = left.Range
				i++
			} else if !right.Empty {
				// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
				//
				//   |____________|
				//   |_____|
				//          |_____|
				//             r
				//
				// Case r.Lo < ranges[i].Lo
				//
				//         |_________|             |_________|             |_________|
				//   |___|                    |____|                  |_________|
				//         |_________|              |________|                   |___|
				//              r                        r                         r
				//

				l.ranges[i] = right.Range
				break
			} else {
				// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
				//
				//   |_________|        |_________|
				//   |_________|        |______________|
				//        ∅                  ∅
				//
				// Case r.Lo < ranges[i].Lo
				//
				//        |_________|             |_________|
				//   |______________|        |___________________|
				//           ∅                         ∅
				//

				l.ranges = append(l.ranges[:i], l.ranges[i+1:]...)
			}
		}
	}
}

// All returns an iterator over all ranges in the range list.
func (l *RangeList[T]) All() iter.Seq[Range[T]] {
	return func(yield func(Range[T]) bool) {
		for _, r := range l.ranges {
			if !yield(r) {
				return
			}
		}
	}
}
