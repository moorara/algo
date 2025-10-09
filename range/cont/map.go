package cont

import "github.com/moorara/algo/generic"

// rangeValue associates a continuous range with a value.
type rangeValue[K Continuous, V generic.Equaler[V]] struct {
	Range[K]
	Value V
}

type RangeMap[K Continuous, V generic.Equaler[V]] struct {
	pairs []rangeValue[K, V]
}
