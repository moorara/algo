package set

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func TestStableSet(t *testing.T) {
	tests := getSetTests()

	for _, tc := range tests {
		eqInt := generic.NewEqualFunc[int]()
		set := NewStableSet(eqInt)

		runSetTest(t, set, tc)
	}
}
