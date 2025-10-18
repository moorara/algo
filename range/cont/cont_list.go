package cont

import (
	"bytes"
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/generic"
)

// RangeList represents a list of continuous ranges.
// The ranges are always non-overlapping and sorted.
type RangeList[T Continuous] interface {
	fmt.Stringer
	generic.Equaler[RangeList[T]]
	generic.Cloner[RangeList[T]]

	Size() int
	Find(T) (Range[T], bool)
	Add(...Range[T])
	Remove(...Range[T])
	All() iter.Seq[Range[T]]
}

// RangeListOpts holds optional settings for creating a RangeList.
type RangeListOpts[T Continuous] struct {
	Format FormatListFunc[T]
}

// FormatListFunc is a function type for formatting a range list into a single string representation.
type FormatListFunc[T Continuous] func(iter.Seq[Range[T]]) string

func defaultFormatList[T Continuous](all iter.Seq[Range[T]]) string {
	var b bytes.Buffer

	for r := range all {
		fmt.Fprintf(&b, "%s ", r)
	}

	// Remove the last space
	if b.Len() > 0 {
		b.Truncate(b.Len() - 1)
	}

	return b.String()
}

// rangeList is a concrete implementation of RangeList interface.
type rangeList[T Continuous] struct {
	ranges []Range[T]
	format FormatListFunc[T]
}

// NewRangeList creates a new range list from the given ranges.
// It panics if any of the provided ranges are invalid.
//
// Ranges stored in the list are always non-overlapping and sorted.
//
// When a new range overlaps existing ranges, the ranges are merged.
func NewRangeList[T Continuous](opts RangeListOpts[T], rs ...Range[T]) RangeList[T] {
	for _, r := range rs {
		if !r.Valid() {
			panic(fmt.Sprintf("invalid range: %s", r))
		}
	}

	if opts.Format == nil {
		opts.Format = defaultFormatList[T]
	}

	l := &rangeList[T]{
		ranges: rs,
		format: opts.Format,
	}

	// Sort ranges by their low bound ascending
	slices.SortFunc(l.ranges, func(lhs, rhs Range[T]) int {
		return compareLoLo(lhs.Lo, rhs.Lo)
	})

	// Merge overlapping and adjacent ranges
	l.mergeRanges()

	return l
}

// searchRanges performs a binary search to find the index of the range that contains the given value.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func (l *rangeList[T]) searchRanges(v Bound[T]) (int, bool) {
	lo, hi := 0, len(l.ranges)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if compareLoLo(v, l.ranges[mid].Lo) < 0 {
			hi = mid - 1
		} else if compareHiLo(l.ranges[mid].Hi, v) < 0 {
			lo = mid + 1
		} else {
			return mid, true
		}
	}

	return lo, false
}

// mergeRanges merges overlapping or adjacent ranges in the sorted list of ranges.
func (l *rangeList[T]) mergeRanges() {
	merged := make([]Range[T], 0, len(l.ranges))

	for _, curr := range l.ranges {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if compareLoHi(curr.Lo, last.Hi) <= 0 {
			if compareHiHi(curr.Hi, last.Hi) > 0 {
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
func (l *rangeList[T]) String() string {
	return l.format(l.All())
}

// Clone implements the generic.Cloner interface.
func (l *rangeList[T]) Clone() RangeList[T] {
	ll := &rangeList[T]{
		ranges: make([]Range[T], len(l.ranges)),
		format: l.format,
	}

	copy(ll.ranges, l.ranges)

	return ll
}

// Equal implements the generic.Equaler interface.
func (l *rangeList[T]) Equal(rhs RangeList[T]) bool {
	ll, ok := rhs.(*rangeList[T])
	if !ok {
		return false
	}

	if len(l.ranges) != len(ll.ranges) {
		return false
	}

	for i, r := range l.ranges {
		if !r.Equal(ll.ranges[i]) {
			return false
		}
	}

	return true
}

// Size returns the number of ranges in the range list.
func (l *rangeList[T]) Size() int {
	return len(l.ranges)
}

// Find returns the range that includes the given value.
// The second return value indicates if such a range exists.
func (l *rangeList[T]) Find(v T) (Range[T], bool) {
	if i, ok := l.searchRanges(Bound[T]{v, false}); ok {
		return l.ranges[i], true
	}

	return Range[T]{}, false
}

// Add inserts the given ranges to the range list.
// It panics if any of the given ranges are invalid.
func (l *rangeList[T]) Add(rs ...Range[T]) {
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
func (l *rangeList[T]) Remove(rs ...Range[T]) {
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
func (l *rangeList[T]) All() iter.Seq[Range[T]] {
	return func(yield func(Range[T]) bool) {
		for _, r := range l.ranges {
			if !yield(r) {
				return
			}
		}
	}
}
