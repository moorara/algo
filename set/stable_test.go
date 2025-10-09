package set

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNewStable(t *testing.T) {
	tests := []struct {
		name            string
		equal           generic.EqualFunc[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name:            "OK",
			equal:           generic.NewEqualFunc[string](),
			vals:            []string{"a", "c", "b", "d"},
			expectedMembers: []string{"a", "c", "b", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStable(tc.equal, tc.vals...)

			assert.NotNil(t, s)
			assert.Equal(t, tc.expectedMembers, s.(*stable[string]).members)
			assert.NotNil(t, s.(*stable[string]).equal)
			assert.NotNil(t, s.(*stable[string]).format)
		})
	}
}

func TestNewStableWithFormat(t *testing.T) {
	tests := []struct {
		name            string
		equal           generic.EqualFunc[string]
		format          Format[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name:            "OK",
			equal:           generic.NewEqualFunc[string](),
			format:          defaultFormat[string],
			vals:            []string{"b", "c", "a", "d"},
			expectedMembers: []string{"b", "c", "a", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewStableWithFormat(tc.equal, tc.format, tc.vals...)

			assert.NotNil(t, s)
			assert.Equal(t, tc.expectedMembers, s.(*stable[string]).members)
			assert.NotNil(t, s.(*stable[string]).equal)
			assert.NotNil(t, s.(*stable[string]).format)
		})
	}
}

func TestStable_String(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name           string
		s              *stable[string]
		expectedString string
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
				format:  defaultFormat[string],
			},
			expectedString: "{}",
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"b", "d", "a", "c"},
				equal:   eqFunc,
				format:  defaultFormat[string],
			},
			expectedString: "{b, d, a, c}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.s.String())
		})
	}
}

func TestStable_Equal(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		t        Set[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			t: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "Equal",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "NotEqual",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &stable[string]{
				members: []string{"c", "d", "e", "f"},
				equal:   eqFunc,
			},
			expected: false,
		},
		{
			name: "NotEqual_Subset",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &stable[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
			},
			expected: false,
		},
		{
			name: "NotEqual_Superset",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
			},
			t: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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

func TestStable_Clone(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name string
		s    *stable[string]
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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

func TestStable_CloneEmpty(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		expected Set[string]
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expected: New[string](eqFunc),
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expected: New[string](eqFunc),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.CloneEmpty()
			assert.True(t, set.Equal(tc.expected))
		})
	}
}

func TestStable_Size(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		expected int
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expected: 0,
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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

func TestStable_IsEmpty(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
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

func TestStable_Add(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *stable[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			vals:            []string{"b", "c", "a", "d"},
			expectedMembers: []string{"b", "c", "a", "d"},
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			vals:            []string{"c", "a", "d"},
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

func TestStable_Remove(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *stable[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			vals:            []string{"a", "b"},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			vals:            []string{"a", "c"},
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

func TestStable_RemoveAll(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name string
		s    *stable[string]
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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

func TestStable_Contains(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		vals     []string
		expected bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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

func TestStable_All(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *stable[string]
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expectedMembers: nil,
		},
		{
			name: "NonEmpty",
			s: &stable[string]{
				members: []string{"b", "c", "d", "a"},
				equal:   eqFunc,
			},
			expectedMembers: []string{"b", "c", "d", "a"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			members := generic.Collect1(tc.s.All())

			for _, expectedMember := range tc.expectedMembers {
				assert.Contains(t, members, expectedMember)
			}

			for _, member := range members {
				assert.Contains(t, tc.expectedMembers, member)
			}
		})
	}
}

func TestStable_AnyMatch(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *stable[string]
		p        generic.Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &stable[string]{
				members: []string{"a", "B", "c", "d"},
				equal:   eqFunc,
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

func TestStable_AllMatch(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *stable[string]
		p        generic.Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: true,
		},
		{
			name: "NonEmpty_No",
			s: &stable[string]{
				members: []string{"A", "B", "c", "D"},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &stable[string]{
				members: []string{"A", "B", "C", "D"},
				equal:   eqFunc,
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

func TestStable_FirstMatch(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name          string
		s             *stable[string]
		p             generic.Predicate1[string]
		expectedValue string
		expectedOK    bool
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			p:             predicate,
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name: "NoMatch",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			p:             predicate,
			expectedValue: "",
			expectedOK:    false,
		},
		{
			name: "OK",
			s: &stable[string]{
				members: []string{"a", "b", "C", "d"},
				equal:   eqFunc,
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

func TestStable_SelectMatch(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name             string
		s                *stable[string]
		p                generic.Predicate1[string]
		expectedSelected Set[string]
	}{
		{
			name: "Empty",
			s: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSelected: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "SelectNone",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSelected: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "SelectSome",
			s: &stable[string]{
				members: []string{"A", "b", "C", "d"},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSelected: &stable[string]{
				members: []string{"A", "C"},
				equal:   eqFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			selected := tc.s.SelectMatch(tc.p)

			assert.True(t, selected.(*stable[string]).Equal(tc.expectedSelected))
		})
	}
}

func TestStable_PartitionMatch(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name              string
		s                 *stable[string]
		p                 generic.Predicate1[string]
		expectedMatched   Set[string]
		expectedUnmatched Set[string]
	}{
		{
			name: "OK",
			s: &stable[string]{
				members: []string{"A", "b", "C", "d"},
				equal:   eqFunc,
			},
			p: predicate,
			expectedMatched: &stable[string]{
				members: []string{"A", "C"},
				equal:   eqFunc,
			},
			expectedUnmatched: &stable[string]{
				members: []string{"b", "d"},
				equal:   eqFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			matched, unmatched := tc.s.PartitionMatch(tc.p)

			assert.True(t, matched.(*stable[string]).Equal(tc.expectedMatched))
			assert.True(t, unmatched.(*stable[string]).Equal(tc.expectedUnmatched))
		})
	}
}

func TestStable_IsSubset(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		superset Set[string]
		expected bool
	}{
		{
			name: "Subset",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			superset: &stable[string]{
				members: []string{"a", "b", "x", "y"},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "NotSubset",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			superset: &stable[string]{
				members: []string{"x", "y"},
				equal:   eqFunc,
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

func TestStable_IsSuperset(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		superset Set[string]
		expected bool
	}{
		{
			name: "Superset",
			s: &stable[string]{
				members: []string{"a", "b", "x", "y"},
				equal:   eqFunc,
			},
			superset: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "NotSuperset",
			s: &stable[string]{
				members: []string{"a", "b", "x", "y"},
				equal:   eqFunc,
			},
			superset: &stable[string]{
				members: []string{"x", "y", "z"},
				equal:   eqFunc,
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

func TestStable_Union(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"c", "e"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"d", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
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

func TestStable_Intersection(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"b", "e"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"b", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{"b"},
				equal:   eqFunc,
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

func TestStable_Difference(t *testing.T) {
	eqFunc := generic.NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *stable[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &stable[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&stable[string]{
					members: []string{"c", "e"},
					equal:   eqFunc,
				},
				&stable[string]{
					members: []string{"d", "f"},
					equal:   eqFunc,
				},
			},
			expected: &stable[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
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
