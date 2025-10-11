package disc

import (
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/generic"
)

// rangeValue associates a discrete range with a value.
type rangeValue[K Discrete, V any] struct {
	Range[K]
	Value V
}

// RangeMap represents a map from discrete ranges to values.
// The ranges are always non-overlapping and sorted.
type RangeMap[K Discrete, V any] struct {
	pairs  []rangeValue[K, V]
	eqVal  generic.EqualFunc[V]
	format FormatMap[K, V]
}

// NewRangeMap creates a new range map from the given ranges.
// It panics if any of the given ranges are invalid.
func NewRangeMap[K Discrete, V any](eqVal generic.EqualFunc[V], pairs map[Range[K]]V) *RangeMap[K, V] {
	m := &RangeMap[K, V]{
		pairs:  make([]rangeValue[K, V], 0, len(pairs)),
		eqVal:  eqVal,
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
		return int(lhs.Lo) - int(rhs.Lo)
	})

	// Merge and/or split overlapping and adjacent ranges
	m.mergeAndSplitRanges()

	return m
}

// searchRanges performs a binary search to find the index of the range that contains the given key.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func (m *RangeMap[K, V]) searchRanges(k K) (int, bool) {
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
func (m *RangeMap[K, V]) mergeAndSplitRanges() {
	merged := make([]rangeValue[K, V], 0, len(m.pairs))

	for _, curr := range m.pairs {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if curr.Lo <= last.Hi {
			if curr.Hi < last.Hi {
				if m.eqVal(last.Value, curr.Value) {
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
				if m.eqVal(last.Value, curr.Value) {
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
				if m.eqVal(last.Value, curr.Value) {
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
		} else if before, _ := last.Range.Adjacent(curr.Range); before && m.eqVal(last.Value, curr.Value) {
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
func (m *RangeMap[K, V]) String() string {
	return m.format(m.All())
}

// Clone implements the generic.Cloner interface.
func (m *RangeMap[K, V]) Clone() *RangeMap[K, V] {
	mm := &RangeMap[K, V]{
		pairs: make([]rangeValue[K, V], len(m.pairs)),
		eqVal: m.eqVal,
	}

	copy(mm.pairs, m.pairs)

	return mm
}

// Equal implements the generic.Equaler interface.
func (m *RangeMap[K, V]) Equal(rhs *RangeMap[K, V]) bool {
	if len(m.pairs) != len(rhs.pairs) {
		return false
	}

	for i, p := range m.pairs {
		if !p.Range.Equal(rhs.pairs[i].Range) || !m.eqVal(p.Value, rhs.pairs[i].Value) {
			return false
		}
	}

	return true
}

// Get returns the range and its associated value that includes the given key.
// The third return value indicates if such a range exists.
func (m *RangeMap[K, V]) Get(k K) (Range[K], V, bool) {
	if i, ok := m.searchRanges(k); ok {
		return m.pairs[i].Range, m.pairs[i].Value, true
	}

	var zero V
	return Range[K]{}, zero, false
}

// Add inserts the given range to the range map.
// It panics if any of the given range are invalid.
func (m *RangeMap[K, V]) Add(k Range[K], v V) {
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
func (m *RangeMap[K, V]) Remove(k Range[K]) {
	if !k.Valid() {
		panic(fmt.Sprintf("invalid range: %s", k))
	}

	for i, _ := m.searchRanges(k.Lo); i < len(m.pairs); i++ {

	}
}

// All returns an iterator over all range-value pairs in the range map.
func (m *RangeMap[K, V]) All() iter.Seq2[Range[K], V] {
	return func(yield func(Range[K], V) bool) {
		for _, p := range m.pairs {
			if !yield(p.Range, p.Value) {
				return
			}
		}
	}
}
