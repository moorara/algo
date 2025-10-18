package disc

import (
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/generic"
)

// RangeMap represents a map from discrete ranges to values.
// The ranges are always non-overlapping and sorted.
type RangeMap[K Discrete, V any] interface {
	fmt.Stringer
	generic.Equaler[RangeMap[K, V]]
	generic.Cloner[RangeMap[K, V]]

	Size() int
	Find(K) (Range[K], V, bool)
	Add(Range[K], V)
	Remove(Range[K])
	All() iter.Seq2[Range[K], V]
}

// rangeValue associates a discrete range with a value.
type rangeValue[K Discrete, V any] struct {
	Range[K]
	Value V
}

// rangeMap is a concrete implementation of RangeMap interface.
type rangeMap[K Discrete, V any] struct {
	pairs  []rangeValue[K, V]
	equal  generic.EqualFunc[V]
	format FormatMap[K, V]
}

// NewRangeMap creates a new range map from the given ranges.
// It panics if any of the provided ranges are invalid.
//
// Ranges stored in the map are always non-overlapping and sorted.
//
// When a new range overlaps existing ranges, overlapping portions are resolved as follows:
//
//   - If the existing range's value equals the new range's value, the ranges are merged.
//   - If the values differ, the new range's value takes precedence and the existing range is split.
func NewRangeMap[K Discrete, V any](equal generic.EqualFunc[V], pairs map[Range[K]]V) RangeMap[K, V] {
	m := &rangeMap[K, V]{
		pairs:  make([]rangeValue[K, V], 0, len(pairs)),
		equal:  equal,
		format: defaultFormatMap[K, V],
	}

	for r, v := range pairs {
		if !r.Valid() {
			panic(fmt.Sprintf("invalid range: %s", r))
		}

		m.pairs = append(m.pairs, rangeValue[K, V]{
			Range: r,
			Value: v,
		})
	}

	// Sort ranges by their low bound ascending
	slices.SortFunc(m.pairs, func(lhs, rhs rangeValue[K, V]) int {
		return int(lhs.Lo - rhs.Lo)
	})

	// Merge and/or split overlapping and adjacent ranges
	m.mergeAndSplitRanges()

	return m
}

// NewRangeMap creates a new range map with a custom format function from the given ranges.
// It panics if any of the provided ranges are invalid.
//
// Ranges stored in the map are always non-overlapping and sorted.
//
// When a new range overlaps existing ranges, overlapping portions are resolved as follows:
//
//   - If the existing range's value equals the new range's value, the ranges are merged.
//   - If the values differ, the new range's value takes precedence and the existing range is split.
func NewRangeMapWithFormat[K Discrete, V any](equal generic.EqualFunc[V], format FormatMap[K, V], pairs map[Range[K]]V) RangeMap[K, V] {
	m := NewRangeMap(equal, pairs).(*rangeMap[K, V])
	m.format = format

	return m
}

// searchRanges performs a binary search to find the index of the range that contains the given key.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func (m *rangeMap[K, V]) searchRanges(k K) (int, bool) {
	lo, hi := 0, len(m.pairs)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if k < m.pairs[mid].Lo {
			hi = mid - 1
		} else if m.pairs[mid].Hi < k {
			lo = mid + 1
		} else {
			return mid, true
		}
	}

	return lo, false
}

// mergeAndSplitRanges merges overlapping or adjacent ranges in the sorted list of ranges.
func (m *rangeMap[K, V]) mergeAndSplitRanges() {
	merged := make([]rangeValue[K, V], 0, len(m.pairs))

	for _, curr := range m.pairs {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if curr.Lo <= last.Hi {
			if curr.Hi < last.Hi {
				if m.equal(last.Value, curr.Value) {
					// Case curr.Lo < last.Hi && curr.Hi < last.Hi && last.Value == curr.Value:
					//
					//   last:  |_____|_____|_____|  Value: A    ---->    |_________________|  Value: A
					//   curr:        |_____|        Value: A    ---->
					//
					// Impossible case of curr.Lo == last.Hi && curr.Hi < last.Hi
					//
				} else {
					// Case curr.Lo < last.Hi && curr.Hi < last.Hi && last.Value != curr.Value:
					//
					//   last:  |_____|_____|_____|  Value: A    ---->    |____||     ||    |  Value: A
					//   curr:        |_____|        Value: B    ---->          |_____||    |  Value: B
					//                                           ---->                 |____|  Value: A
					//
					// Impossible case of curr.Lo == last.Hi && curr.Hi < last.Hi
					//

					lastEnd := last.Hi
					last.Hi = curr.Lo - 1
					merged = append(merged, curr)
					merged = append(merged, rangeValue[K, V]{
						Range: Range[K]{Lo: curr.Hi + 1, Hi: lastEnd},
						Value: last.Value,
					})
				}
			} else if curr.Hi == last.Hi {
				if m.equal(last.Value, curr.Value) {
					// Case curr.Lo < last.Hi && curr.Hi == last.Hi && last.Value == curr.Value:
					//
					//   last:  |_____|___________|  Value: A    ---->    |_________________|  Value: A
					//   curr:        |___________|  Value: A    ---->
					//
					// Case curr.Lo == last.Hi && curr.Hi == last.Hi && last.Value == curr.Value:
					//
					//   last:  |_________________|  Value: A    ---->    |_________________|  Value: A
					//   curr:                    |  Value: A    ---->
					//
				} else {
					// Case curr.Lo < last.Hi && curr.Hi == last.Hi && last.Value != curr.Value:
					//
					//   last:  |_____|___________|  Value: A    ---->    |____||           |  Value: A
					//   curr:        |___________|  Value: B    ---->          |___________|  Value: B
					//
					// Case curr.Lo == last.Hi && curr.Hi == last.Hi && last.Value != curr.Value:
					//
					//   last:  |_________________|  Value: A    ---->    |________________||  Value: A
					//   curr:                    |  Value: B    ---->                      |  Value: B
					//

					last.Hi = curr.Lo - 1
					merged = append(merged, curr)
				}
			} else /* if curr.Hi > last.Hi */ {
				if m.equal(last.Value, curr.Value) {
					// Case curr.Lo < last.Hi && curr.Hi > last.Hi && last.Value == curr.Value:
					//
					//   last:  |_____|_____|     |  Value: A    ---->    |_________________|  Value: A
					//   curr:        |___________|  Value: A    ---->
					//
					// Case curr.Lo == last.Hi && curr.Hi > last.Hi && last.Value == curr.Value:
					//
					//   last:  |___________|     |  Value: A    ---->    |_________________|  Value: A
					//   curr:              |_____|  Value: A    ---->
					//

					last.Hi = curr.Hi
				} else {
					// Case curr.Lo < last.Hi && curr.Hi > last.Hi && last.Value != curr.Value:
					//
					//   last:  |_____|_____|     |  Value: A    ---->    |____||           |  Value: A
					//   curr:        |___________|  Value: B    ---->          |___________|  Value: B
					//
					// Case curr.Lo == last.Hi && curr.Hi > last.Hi && last.Value != curr.Value:
					//
					//   last:  |___________|     |  Value: A    ---->    |__________||     |  Value: A
					//   curr:              |_____|  Value: B    ---->                |_____|  Value: B
					//

					last.Hi = curr.Lo - 1
					merged = append(merged, curr)
				}
			}
		} else if before, _ := last.Range.Adjacent(curr.Range); before && m.equal(last.Value, curr.Value) {
			// Case last.Hi is immediately before curr.Lo && last.Value == curr.Value:
			//
			//   last:  |__________||     |  Value: A    ---->    |_________________|  Value: A
			//   curr:              |_____|  Value: A    ---->
			//

			last.Hi = curr.Hi
		} else {
			merged = append(merged, curr)
		}
	}

	m.pairs = merged
}

// String implements the fmt.Stringer interface.
func (m *rangeMap[K, V]) String() string {
	return m.format(m.All())
}

// Clone implements the generic.Cloner interface.
func (m *rangeMap[K, V]) Clone() RangeMap[K, V] {
	mm := &rangeMap[K, V]{
		pairs:  make([]rangeValue[K, V], len(m.pairs)),
		equal:  m.equal,
		format: m.format,
	}

	copy(mm.pairs, m.pairs)

	return mm
}

// Equal implements the generic.Equaler interface.
func (m *rangeMap[K, V]) Equal(rhs RangeMap[K, V]) bool {
	mm, ok := rhs.(*rangeMap[K, V])
	if !ok {
		return false
	}

	if len(m.pairs) != len(mm.pairs) {
		return false
	}

	for i, p := range m.pairs {
		if !p.Range.Equal(mm.pairs[i].Range) || !m.equal(p.Value, mm.pairs[i].Value) {
			return false
		}
	}

	return true
}

// Size returns the number of ranges in the range map.
func (m *rangeMap[K, V]) Size() int {
	return len(m.pairs)
}

// Find returns the range and its associated value that includes the given key.
// The third return value indicates if such a range exists.
func (m *rangeMap[K, V]) Find(v K) (Range[K], V, bool) {
	if i, ok := m.searchRanges(v); ok {
		return m.pairs[i].Range, m.pairs[i].Value, true
	}

	var zero V
	return Range[K]{}, zero, false
}

// Add inserts the given range to the range map.
// It panics if any of the given range are invalid.
func (m *rangeMap[K, V]) Add(k Range[K], v V) {
	p := rangeValue[K, V]{
		Range: k,
		Value: v,
	}

	if !p.Valid() {
		panic(fmt.Sprintf("invalid range: %s", p.Range))
	}

	// Find the insertion point
	i, ok := m.searchRanges(p.Lo)
	if ok {
		i++
	}

	// Insert the new entry at position i
	m.pairs = append(m.pairs, rangeValue[K, V]{})
	copy(m.pairs[i+1:], m.pairs[i:])
	m.pairs[i] = p

	// Merge and/or split overlapping and adjacent ranges
	m.mergeAndSplitRanges()
}

// Remove deletes the given range from the range map.
// It panics if any of the given range are invalid.
func (m *rangeMap[K, V]) Remove(k Range[K]) {
	if !k.Valid() {
		panic(fmt.Sprintf("invalid range: %s", k))
	}

	i, _ := m.searchRanges(k.Lo)

	for i < len(m.pairs) {
		left, right := m.pairs[i].Range.Subtract(k)

		if !left.Empty && !right.Empty {
			// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
			//
			//   |______________|
			//        |____|
			//   |___|      |___|
			//     l          r
			//

			m.pairs[i].Range = left.Range
			m.pairs = append(m.pairs, rangeValue[K, V]{})
			copy(m.pairs[i+2:], m.pairs[i+1:])
			m.pairs[i+1].Range = right.Range
			m.pairs[i+1].Value = m.pairs[i].Value
			break
		} else if !left.Empty {
			// Case ranges[i].Lo <= r.Lo <= ranges[i].Hi
			//
			//   |____________|        |____________|              |____________|
			//          |_____|               |___________|                     |_____|
			//   |_____|               |_____|                     |___________|
			//      l                     l                              l
			//

			m.pairs[i].Range = left.Range
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

			m.pairs[i].Range = right.Range
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

			m.pairs = append(m.pairs[:i], m.pairs[i+1:]...)
		}
	}
}

// All returns an iterator over all range-value pairs in the range map.
func (m *rangeMap[K, V]) All() iter.Seq2[Range[K], V] {
	return func(yield func(Range[K], V) bool) {
		for _, p := range m.pairs {
			if !yield(p.Range, p.Value) {
				return
			}
		}
	}
}
