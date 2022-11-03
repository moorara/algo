package set

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNew(t *testing.T) {
	set := New(generic.NewEqualFunc[string]())
	assert.NotNil(t, set)
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name            string
		s               *set[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			vals:            []string{"a", "b", "c", "d"},
			expectedMembers: []string{"a", "b", "c", "d"},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
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
	tests := []struct {
		name            string
		s               *set[string]
		vals            []string
		expectedMembers []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			vals:            []string{"a", "b"},
			expectedMembers: []string{},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
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

func TestSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		s        *set[string]
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			expected: true,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
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
	tests := []struct {
		name     string
		s        *set[string]
		val      string
		expected bool
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			val:      "c",
			expected: false,
		},
		{
			name: "NonEmpty_No",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b"},
			},
			val:      "c",
			expected: false,
		},
		{
			name: "NonEmpty_Yes",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b", "c", "d"},
			},
			val:      "c",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.s.Contains(tc.val)
			assert.Equal(t, tc.expected, b)
		})
	}
}

func TestSet_Members(t *testing.T) {
	tests := []struct {
		name     string
		s        *set[string]
		expected []string
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			expected: []string{},
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b", "c", "d"},
			},
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mems := tc.s.Members()
			assert.Equal(t, tc.expected, mems)
		})
	}
}

func TestSet_Cardinality(t *testing.T) {
	tests := []struct {
		name     string
		s        *set[string]
		expected int
	}{
		{
			name: "Empty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{},
			},
			expected: 0,
		},
		{
			name: "NonEmpty",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
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

func TestSet_Union(t *testing.T) {
	tests := []struct {
		name            string
		s               *set[string]
		sets            []Set[string]
		expectedMembers []string
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"e", "f"},
				},
			},
			expectedMembers: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"c", "e"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"d", "f"},
				},
			},
			expectedMembers: []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Union(tc.sets...)
			assert.Equal(t, tc.expectedMembers, set.Members())
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name            string
		s               *set[string]
		sets            []Set[string]
		expectedMembers []string
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"e", "f"},
				},
			},
			expectedMembers: []string{},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"b", "e"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"b", "f"},
				},
			},
			expectedMembers: []string{"b"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Intersection(tc.sets...)
			assert.Equal(t, tc.expectedMembers, set.Members())
		})
	}
}

func TestSet_Difference(t *testing.T) {
	tests := []struct {
		name            string
		s               *set[string]
		sets            []Set[string]
		expectedMembers []string
	}{
		{
			name: "Disjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"c", "d"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"e", "f"},
				},
			},
			expectedMembers: []string{"a", "b"},
		},
		{
			name: "NotDisjoint",
			s: &set[string]{
				equal:   generic.NewEqualFunc[string](),
				members: []string{"a", "b", "c", "d"},
			},
			sets: []Set[string]{
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"c", "e"},
				},
				&set[string]{
					equal:   generic.NewEqualFunc[string](),
					members: []string{"d", "f"},
				},
			},
			expectedMembers: []string{"a", "b"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.s.Difference(tc.sets...)
			assert.Equal(t, tc.expectedMembers, set.Members())
		})
	}
}
