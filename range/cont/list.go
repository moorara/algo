package cont

// RangeList represents a list of continuous ranges.
// The ranges are always non-overlapping and sorted.
type RangeList[T Continuous] struct {
	ranges []Range[T]
}
