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

func TestSet_String(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		expectedStrings []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expectedStrings: []string{},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expectedStrings: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.s.String()
			for _, expectedString := range tc.expectedStrings {
				assert.Contains(t, str, expectedString)
			}
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
				members: []string{},
				equal:   eqFunc,
			},
			t: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "Equal",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expected: true,
		},
		{
			name: "NotEqual",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &set[string]{
				members: []string{"c", "d", "e", "f"},
				equal:   eqFunc,
			},
			expected: false,
		},
		{
			name: "NotEqual_Subset",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			t: &set[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
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

func TestSet_Size(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name         string
		s            *set[string]
		expectedSize int
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expectedSize: 0,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expectedSize: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSize, tc.s.Size())
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		expectedIsEmpty bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			expectedIsEmpty: true,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			expectedIsEmpty: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsEmpty, tc.s.IsEmpty())
		})
	}
}

func TestSet_Add(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		ss              []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			ss:              []string{"a", "b", "c", "d"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			ss:              []string{"a", "c", "d"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Add(tc.ss...)
			assert.Equal(t, tc.expectedMembers, tc.s.members)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	eqFunc := NewEqualFunc[string]()

	tests := []struct {
		name            string
		s               *set[string]
		ss              []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			ss:              []string{"a", "b"},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			ss:              []string{"a", "c"},
			expectedMembers: []string{"b", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Remove(tc.ss...)
			assert.Equal(t, tc.expectedMembers, tc.s.members)
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
				members: []string{},
				equal:   eqFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			vals:     []string{"c"},
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
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
				members: []string{},
				equal:   eqFunc,
			},
			expectedMembers: nil,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			members := slices.Collect(tc.s.All())

			for _, expectedMember := range tc.expectedMembers {
				assert.Contains(t, members, expectedMember)
			}

			for _, member := range members {
				assert.Contains(t, tc.expectedMembers, member)
			}
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
				members: []string{},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
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
				members: []string{},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: true,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				members: []string{"A", "B", "c", "D"},
				equal:   eqFunc,
			},
			p:        predicate,
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
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

func TestSet_SelectMatch(t *testing.T) {
	eqFunc := NewEqualFunc[string]()
	predicate := func(s string) bool {
		return strings.ToUpper(s) == s
	}

	tests := []struct {
		name        string
		s           *set[string]
		p           Predicate1[string]
		expectedSet Set[string]
	}{
		{
			name: "Empty",
			s: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSet: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "SelectNone",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSet: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "SelectSome",
			s: &set[string]{
				members: []string{"A", "c", "C", "d"},
				equal:   eqFunc,
			},
			p: predicate,
			expectedSet: &set[string]{
				members: []string{"A", "C"},
				equal:   eqFunc,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.SelectMatch(tc.p).(*set[string])
			assert.True(t, set.Equals(tc.expectedSet))
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
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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
				members: []string{},
				equal:   eqFunc,
			},
			expected: New[string](eqFunc),
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
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
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"c", "e"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"d", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{"a", "b", "c", "d", "e", "f"},
				equal:   eqFunc,
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
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"b", "e"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"b", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{"b"},
				equal:   eqFunc,
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
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"c", "d"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"e", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				members: []string{"a", "b", "c", "d"},
				equal:   eqFunc,
			},
			sets: []Set[string]{
				&set[string]{
					members: []string{"c", "e"},
					equal:   eqFunc,
				},
				&set[string]{
					members: []string{"d", "f"},
					equal:   eqFunc,
				},
			},
			expected: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
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
				members: []string{},
				equal:   eqFunc,
			},
			expected: &set[Set[string]]{
				members: []Set[string]{
					&set[string]{
						members: []string{},
						equal:   eqFunc,
					},
				},
				equal: setEqFunc,
			},
		},
		{
			name: "OneElement",
			s: &set[string]{
				members: []string{"a"},
				equal:   eqFunc,
			},
			expected: &set[Set[string]]{
				members: []Set[string]{
					&set[string]{
						members: []string{},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a"},
						equal:   eqFunc,
					},
				},
				equal: setEqFunc,
			},
		},
		{
			name: "TwoElements",
			s: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			expected: &set[Set[string]]{
				members: []Set[string]{
					&set[string]{
						members: []string{},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"b"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a", "b"},
						equal:   eqFunc,
					},
				},
				equal: setEqFunc,
			},
		},
		{
			name: "ThreeElements",
			s: &set[string]{
				members: []string{"a", "b", "c"},
				equal:   eqFunc,
			},
			expected: &set[Set[string]]{
				members: []Set[string]{
					&set[string]{
						members: []string{},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"b"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"c"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a", "b"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a", "c"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"b", "c"},
						equal:   eqFunc,
					},
					&set[string]{
						members: []string{"a", "b", "c"},
						equal:   eqFunc,
					},
				},
				equal: setEqFunc,
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
				members: []string{},
				equal:   eqFunc,
			},
			expected: &set[Set[Set[string]]]{
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						members: []Set[string]{},
						equal:   setEqFunc,
					},
				},
				equal: partEqFunc,
			},
		},
		{
			name: "OneElement",
			s: &set[string]{
				members: []string{"a"},
				equal:   eqFunc,
			},
			expected: &set[Set[Set[string]]]{
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
				},
				equal: partEqFunc,
			},
		},
		{
			name: "TwoElements",
			s: &set[string]{
				members: []string{"a", "b"},
				equal:   eqFunc,
			},
			expected: &set[Set[Set[string]]]{
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"b"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
					&set[Set[string]]{ // 2nd partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a", "b"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
				},
				equal: partEqFunc,
			},
		},
		{
			name: "ThreeElements",
			s: &set[string]{
				members: []string{"a", "b", "c"},
				equal:   eqFunc,
			},
			expected: &set[Set[Set[string]]]{
				members: []Set[Set[string]]{
					&set[Set[string]]{ // 1st partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"b"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"c"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
					&set[Set[string]]{ // 2nd partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a", "b"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"c"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
					&set[Set[string]]{ // 3rd partition
						members: []Set[string]{
							&set[string]{
								members: []string{"b"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"a", "c"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
					&set[Set[string]]{ // 4th partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a"},
								equal:   eqFunc,
							},
							&set[string]{
								members: []string{"b", "c"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
					&set[Set[string]]{ // 5th partition
						members: []Set[string]{
							&set[string]{
								members: []string{"a", "b", "c"},
								equal:   eqFunc,
							},
						},
						equal: setEqFunc,
					},
				},
				equal: partEqFunc,
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
