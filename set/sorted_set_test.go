package set

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func TestSortedSet(t *testing.T) {
	tests := getSetTests()

	for _, tc := range tests {
		cmpInt := generic.NewCompareFunc[int]()
		set := NewSortedSet(cmpInt)

		runSetTest(t, set, tc)
	}
}
