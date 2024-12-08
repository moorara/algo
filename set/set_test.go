package set

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/moorara/algo/generic"
)

func TestNew(t *testing.T) {
	set := New(NewEqualFunc[string]())
	assert.NotNil(t, set)
}

func TestSet_Add(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			vals:            []string{"a", "b", "c", "d"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			vals:            []string{"a", "c", "d"},
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

func TestSet_Remove(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			vals:            []string{"a", "b"},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
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

func TestSet_Cardinality(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		expected int
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: 0,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			card := tc.s.Cardinality()
			assert.Equal(t, tc.expected, card)
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: true,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.IsEmpty()
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSet_Contains(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		vals     []string
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
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

func TestSet_Clone(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name string
		s    *set[string]
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Clone()
			assert.True(t, set.Equals(tc.s))
		})
	}
}

func TestSet_CloneEmpty(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		expected Set[string]
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: New[string](eqFunc),
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expected: New[string](eqFunc),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.CloneEmpty()
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestSet_Union(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"e", "f"},
				},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d", "e", "f"},
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"c", "e"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"d", "f"},
				},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d", "e", "f"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Union(tc.sets...)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"e", "f"},
				},
			},
			expected: &set[string]{
				members: []string{},
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"b", "e"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"b", "f"},
				},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"b"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Intersection(tc.sets...)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestSet_Difference(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		sets     []Set[string]
		expected Set[string]
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"e", "f"},
				},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   eqFunc,
					members: []string{"c", "e"},
				},
				&set[string]{
					equal:   eqFunc,
					members: []string{"d", "f"},
				},
			},
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Difference(tc.sets...)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestSet_String(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		expected string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: "{}",
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expected: "{a, b, c, d}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.s.String()
			assert.Equal(t, tc.expected, str)
		})
	}
}

func TestSet_Equals(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name     string
		s        *set[string]
		t        Set[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			t: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: true,
		},
		{
			name: "Equal",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			t: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expected: true,
		},
		{
			name: "NotEqual",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			t: &set[string]{
				equal:   eqFunc,
				members: []string{"c", "d", "e", "f"},
			},
			expected: false,
		},
		{
			name: "NotEqual_Subset",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			t: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d", "e", "f"},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.Equals(tc.t)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSet_All(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expectedMembers: nil,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			members := slices.Collect(tc.s.All())
			assert.Equal(t, tc.expectedMembers, members)
		})
	}
}

func TestSet_AnyMatch(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *set[string]
		p        Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "B", "c", "d"},
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

func TestSet_AllMatch(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *set[string]
		p        Predicate1[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			p:        predicate,
			expected: true,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"A", "B", "c", "D"},
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"A", "B", "C", "D"},
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

func TestSet_Filter(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name     string
		s        *set[string]
		p        Predicate1[string]
		expected Set[string]
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			p: predicate,
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
		},
		{
			name: "SelectNone",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c", "d"},
			},
			p: predicate,
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
		},
		{
			name: "SelectSome",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"A", "c", "C", "d"},
			},
			p: predicate,
			expected: &set[string]{
				equal:   eqFunc,
				members: []string{"A", "C"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Filter(tc.p)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestMapper_Transform(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	mapper := func(s string) string {
		return strings.ToUpper(s)
	}

	tests := []struct {
		name     string
		f        Transformer[string, string]
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
			set := tc.f.Transform(tc.s, eqFunc)
			assert.True(t, set.Equals(tc.expected))
		})
	}
}

func TestSet_Powerset(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	setEqFunc := func(a, b Set[string]) bool { return a.Equals(b) }

	tests := []struct {
		name     string
		s        Set[string]
		expected Set[Set[string]]
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: &set[Set[string]]{
				equal: setEqFunc,
				members: []Set[string]{
					&set[string]{
						equal:   eqFunc,
						members: []string{},
					},
				},
			},
		},
		{
			name: "OneElement",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a"},
			},
			expected: &set[Set[string]]{
				equal: setEqFunc,
				members: []Set[string]{
					&set[string]{
						equal:   eqFunc,
						members: []string{},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a"},
					},
				},
			},
		},
		{
			name: "TwoElements",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			expected: &set[Set[string]]{
				equal: setEqFunc,
				members: []Set[string]{
					&set[string]{
						equal:   eqFunc,
						members: []string{},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"b"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a", "b"},
					},
				},
			},
		},
		{
			name: "ThreeElements",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c"},
			},
			expected: &set[Set[string]]{
				equal: setEqFunc,
				members: []Set[string]{
					&set[string]{
						equal:   eqFunc,
						members: []string{},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"b"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"c"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a", "b"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a", "c"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"b", "c"},
					},
					&set[string]{
						equal:   eqFunc,
						members: []string{"a", "b", "c"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ps := Powerset[string](tc.s)
			assert.True(t, ps.Equals(tc.expected))
		})
	}
}

func TestSet_Partitions(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	setEqFunc := func(a, b Set[string]) bool { return a.Equals(b) }
	partEqFunc := func(a, b Set[Set[string]]) bool { return a.Equals(b) }

	tests := []struct {
		name     string
		s        Set[string]
		expected Set[Set[Set[string]]]
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{},
			},
			expected: &set[Set[Set[string]]]{
				equal: partEqFunc,
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						equal:   setEqFunc,
						members: []Set[string]{},
					},
				},
			},
		},
		{
			name: "OneElement",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a"},
			},
			expected: &set[Set[Set[string]]]{
				equal: partEqFunc,
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a"},
							},
						},
					},
				},
			},
		},
		{
			name: "TwoElements",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b"},
			},
			expected: &set[Set[Set[string]]]{
				equal: partEqFunc,
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"b"},
							},
						},
					},
					&set[Set[string]]{ // 2nd partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a", "b"},
							},
						},
					},
				},
			},
		},
		{
			name: "ThreeElements",
			s: &set[string]{
				equal:   eqFunc,
				members: []string{"a", "b", "c"},
			},
			expected: &set[Set[Set[string]]]{
				equal: partEqFunc,
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"b"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"c"},
							},
						},
					},
					&set[Set[string]]{ // 2nd partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a", "b"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"c"},
							},
						},
					},
					&set[Set[string]]{ // 3rd partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"b"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"a", "c"},
							},
						},
					},
					&set[Set[string]]{ // 4th partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a"},
							},
							&set[string]{
								equal:   eqFunc,
								members: []string{"b", "c"},
							},
						},
					},
					&set[Set[string]]{ // 5th partition
						equal: setEqFunc,
						members: []Set[string]{
							&set[string]{
								equal:   eqFunc,
								members: []string{"a", "b", "c"},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := Partitions[string](tc.s)
			assert.True(t, parts.Equals(tc.expected))
		})
	}
}
