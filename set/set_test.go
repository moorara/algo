package set

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name                string
		addValues           []string
		removeValues        []string
		expectedMembers     []string
		expectedIsEmpty     bool
		expectedCardinality int
	}{
		{
			name:                "NonEmpty",
			addValues:           []string{"a", "b", "c", "d", "e", "f"},
			removeValues:        []string{"a", "c", "e"},
			expectedMembers:     []string{"b", "d", "f"},
			expectedIsEmpty:     false,
			expectedCardinality: 3,
		},
		{
			name:                "Empty",
			addValues:           []string{"a", "b", "c", "d", "e", "f"},
			removeValues:        []string{"a", "b", "c", "d", "e", "f"},
			expectedMembers:     []string{},
			expectedIsEmpty:     true,
			expectedCardinality: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			equal := generic.NewEqualFunc[string]()
			set := New[string](equal)

			// Th set is initially empty
			assert.True(t, set.IsEmpty())
			assert.Zero(t, set.Cardinality())

			set.Add(tc.addValues...)
			set.Remove(tc.removeValues...)

			members := set.Members()
			assert.Equal(t, tc.expectedMembers, members)

			for _, member := range tc.expectedMembers {
				assert.True(t, set.Contains(member))
			}

			assert.Equal(t, tc.expectedIsEmpty, set.IsEmpty())
			assert.Equal(t, tc.expectedCardinality, set.Cardinality())
		})
	}
}
