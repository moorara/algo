package set

import (
	"testing"

	"github.com/moorara/algo/generic"
)

func TestSet(t *testing.T) {
	tests := getSetTests()

	for _, tc := range tests {
		eqInt := generic.NewEqualFunc[int]()
		set := New(eqInt)

		runSetTest(t, set, tc)
	}
}
