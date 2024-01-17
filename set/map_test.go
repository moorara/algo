package set

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestMapper_Map(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	mapper := func(s string) string {
		return strings.ToUpper(s)
	}

	tests := []struct {
		name     string
		f        Mapper[string, string]
		s        *set[string]
		expected Set[string]
	}{
		{
			name: "Empty",
			f:    mapper,
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
		},
		{
			name: "NonEmpty",
			f:    mapper,
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"A", "B", "C", "D"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.f.Map(tc.s, eqFunc)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}
