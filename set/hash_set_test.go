package set

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

func TestHashSet(t *testing.T) {
	tests := getSetTests()

	for _, tc := range tests {
		hashInt := hash.HashFuncForInt[int](nil)
		eqInt := generic.NewEqualFunc[int]()
		set := NewHashSet(hashInt, eqInt, HashSetOpts{})

		runSetTest(t, set, tc)
	}
}

func TestNewHashSet_Panic(t *testing.T) {
	hashInt := hash.HashFuncForInt[int](nil)
	eqInt := generic.NewEqualFunc[int]()

	assert.PanicsWithValue(t, "The hash set capacity must be a prime number", func() {
		NewHashSet(hashInt, eqInt, HashSetOpts{
			InitialCap: 69,
		})
	})
}
