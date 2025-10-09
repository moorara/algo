package set

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNewSorted(t *testing.T) {
	tests := []struct {
		name            string
		compare         generic.CompareFunc[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name:            "OK",
			compare:         generic.NewCompareFunc[string](),
			vals:            []string{"d", "c", "b", "a"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSorted(tc.compare, tc.vals...)

			assert.NotNil(t, s)
			assert.Equal(t, tc.expectedMembers, s.(*sorted[string]).members)
			assert.NotNil(t, s.(*sorted[string]).compare)
			assert.NotNil(t, s.(*sorted[string]).format)
		})
	}
}

func TestNewSortedWithFormat(t *testing.T) {
	tests := []struct {
		name            string
		compare         generic.CompareFunc[string]
		format          Format[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name:            "OK",
			compare:         generic.NewCompareFunc[string](),
			format:          defaultFormat[string],
			vals:            []string{"d", "c", "b", "a"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSortedWithFormat(tc.compare, tc.format, tc.vals...)

			assert.NotNil(t, s)
			assert.Equal(t, tc.expectedMembers, s.(*sorted[string]).members)
			assert.NotNil(t, s.(*sorted[string]).compare)
			assert.NotNil(t, s.(*sorted[string]).format)
		})
	}
}

func TestSorted_String(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name           string
		s              *sorted[string]
		expectedString string
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
				format:  defaultFormat[string],
			},
			expectedString: "{}",
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
				format:  defaultFormat[string],
			},
			expectedString: "{a, b, c, d}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.s.String())
		})
	}
}

func TestSorted_Equal(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		t        Set[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			t: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			expected: true,
		},
		{
			name: "Equal",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			t: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			expected: true,
		},
		{
			name: "NotEqual",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			t: &sorted[string]{
				members: []string{"c", "d", "e", "f"},
				compare: cmpFunc,
			},
			expected: false,
		},
		{
			name: "NotEqual_Subset",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			t: &sorted[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				compare: cmpFunc,
			},
			expected: false,
		},
		{
			name: "NotEqual_Superset",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				compare: cmpFunc,
			},
			t: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.Equal(tc.t)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_Clone(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name string
		s    *sorted[string]
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Clone()
			assert.True(t, set.Equal(tc.s))
		})
	}
}

func TestSorted_CloneEmpty(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		expected Set[string]
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			expected: NewSorted[string](cmpFunc),
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			expected: NewSorted[string](cmpFunc),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.CloneEmpty()
			assert.True(t, set.Equal(tc.expected))
		})
	}
}

func TestSorted_Size(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		expected int
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			expected: 0,
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.s.Size()
			assert.Equal(t, tc.expected, size)
		})
	}
}

func TestSorted_IsEmpty(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			expected: true,
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.s.IsEmpty())
		})
	}
}

func TestSorted_Add(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name            string
		s               *sorted[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			vals:            []string{"d", "c", "b", "a"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			vals:            []string{"d", "c", "a"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Add(tc.vals...)
			assert.Equal(t, tc.expectedMembers, tc.s.members)
		})
	}
}

func TestSorted_Remove(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name            string
		s               *sorted[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			vals:            []string{"b", "a"},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			vals:            []string{"c", "a"},
			expectedMembers: []string{"b", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Remove(tc.vals...)
			assert.Equal(t, tc.expectedMembers, tc.s.members)
		})
	}
}

func TestSorted_RemoveAll(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name string
		s    *sorted[string]
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.RemoveAll()
			assert.Zero(t, tc.s.Size())
			assert.True(t, tc.s.IsEmpty())
		})
	}
}

func TestSorted_Contains(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		vals     []string
		expected bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			vals:     []string{"b", "c"},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.Contains(tc.vals...)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_All(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name            string
		s               *sorted[string]
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			members := generic.Collect1(tc.s.All())
			assert.Equal(t, tc.expectedMembers, members)
		})
	}
}

func TestSorted_AnyMatch(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *sorted[string]
		p        generic.Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &sorted[string]{
				members: []string{"a", "B", "c", "d"},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.AnyMatch(tc.p)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_AllMatch(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *sorted[string]
		p        generic.Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: true,
		},
		{
			name: "NonEmpty_No",
			s: &sorted[string]{
				members: []string{"A", "B", "c", "D"},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &sorted[string]{
				members: []string{"A", "B", "C", "D"},
				compare: cmpFunc,
			},
			p:        predicate,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.AllMatch(tc.p)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_FirstMatch(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name          string
		s             *sorted[string]
		p             generic.Predicate1[string]
		expectedValue string
		expectedOK    bool
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			p:             predicate,
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name: "NoMatch",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			p:             predicate,
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name: "OK",
			s: &sorted[string]{
				members: []string{"a", "b", "C", "d"},
				compare: cmpFunc,
			},
			p:             predicate,
			expectedValue: "C",
			expectedOK:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := tc.s.FirstMatch(tc.p)
			assert.Equal(t, tc.expectedValue, val)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestSorted_SelectMatch(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name             string
		s                *sorted[string]
		p                generic.Predicate1[string]
		expectedSelected Set[string]
	}{
		{
			name: "Empty",
			s: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
			p: predicate,
			expectedSelected: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
		},
		{
			name: "SelectNone",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			p: predicate,
			expectedSelected: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
		},
		{
			name: "SelectSome",
			s: &sorted[string]{
				members: []string{"A", "b", "C", "d"},
				compare: cmpFunc,
			},
			p: predicate,
			expectedSelected: &sorted[string]{
				members: []string{"A", "C"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			selected := tc.s.SelectMatch(tc.p)

			assert.True(t, selected.(*sorted[string]).Equal(tc.expectedSelected))
		})
	}
}

func TestSorted_PartitionMatch(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name              string
		s                 *sorted[string]
		p                 generic.Predicate1[string]
		expectedMatched   Set[string]
		expectedUnmatched Set[string]
	}{
		{
			name: "OK",
			s: &sorted[string]{
				members: []string{"A", "b", "C", "d"},
				compare: cmpFunc,
			},
			p: predicate,
			expectedMatched: &sorted[string]{
				members: []string{"A", "C"},
				compare: cmpFunc,
			},
			expectedUnmatched: &sorted[string]{
				members: []string{"b", "d"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matched, unmatched := tc.s.PartitionMatch(tc.p)

			assert.True(t, matched.(*sorted[string]).Equal(tc.expectedMatched))
			assert.True(t, unmatched.(*sorted[string]).Equal(tc.expectedUnmatched))
		})
	}
}

func TestSorted_IsSubset(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		superset Set[string]
		expected bool
	}{
		{
			name: "Subset",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			superset: &sorted[string]{
				members: []string{"a", "b", "x", "y"},
				compare: cmpFunc,
			},
			expected: true,
		},
		{
			name: "NotSubset",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			superset: &sorted[string]{
				members: []string{"x", "y"},
				compare: cmpFunc,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.IsSubset(tc.superset)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_IsSuperset(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		superset Set[string]
		expected bool
	}{
		{
			name: "Superset",
			s: &sorted[string]{
				members: []string{"a", "b", "x", "y"},
				compare: cmpFunc,
			},
			superset: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			expected: true,
		},
		{
			name: "NotSuperset",
			s: &sorted[string]{
				members: []string{"a", "b", "x", "y"},
				compare: cmpFunc,
			},
			superset: &sorted[string]{
				members: []string{"x", "y", "z"},
				compare: cmpFunc,
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.IsSuperset(tc.superset)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSorted_Union(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"c", "d"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"e", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				compare: cmpFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"c", "e"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"d", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Union(tc.sets...)
			assert.True(t, set.Equal(tc.expected))
		})
	}
}

func TestSorted_Intersection(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"c", "d"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"e", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{},
				compare: cmpFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"b", "e"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"b", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{"b"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Intersection(tc.sets...)
			assert.True(t, set.Equal(tc.expected))
		})
	}
}

func TestSorted_Difference(t *testing.T) {
	cmpFunc := generic.NewCompareFunc[string]()

	tests := []struct {
		name     string
		s        *sorted[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"c", "d"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"e", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &sorted[string]{
				members: []string{"a", "b", "c", "d"},
				compare: cmpFunc,
			},
			sets: []Set[string]{
				&sorted[string]{
					members: []string{"c", "e"},
					compare: cmpFunc,
				},
				&sorted[string]{
					members: []string{"d", "f"},
					compare: cmpFunc,
				},
			},
			expected: &sorted[string]{
				members: []string{"a", "b"},
				compare: cmpFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Difference(tc.sets...)
			assert.True(t, set.Equal(tc.expected))
		})
	}
}
