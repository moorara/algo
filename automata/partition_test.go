package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqGroup(t *testing.T) {
	tests := []struct {
		name          string
		a, b          group
		expectedEqual bool
	}{
		{
			name:          "NotEqual",
			a:             group{States: NewStates(2, 4, 8), Rep: 2},
			b:             group{States: NewStates(3, 5, 7), Rep: 3},
			expectedEqual: false,
		},
		{
			name:          "Equal_SameRepresentative",
			a:             group{States: NewStates(2, 4, 8), Rep: 2},
			b:             group{States: NewStates(2, 4, 8), Rep: 2},
			expectedEqual: true,
		},
		{
			name:          "Equal_DiffRepresentative",
			a:             group{States: NewStates(2, 4, 8), Rep: 2},
			b:             group{States: NewStates(2, 4, 8), Rep: 4},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, eqGroup(tc.a, tc.b))
		})
	}
}

func TestCmpGroup(t *testing.T) {
	tests := []struct {
		name            string
		a, b            group
		expectedCompare int
	}{
		{
			name:            "LessThan",
			a:               group{States: NewStates(2, 4, 8), Rep: 2},
			b:               group{States: NewStates(3, 5, 7), Rep: 3},
			expectedCompare: -1,
		},
		{
			name:            "GreaterThan",
			a:               group{States: NewStates(3, 5, 7), Rep: 3},
			b:               group{States: NewStates(2, 4, 8), Rep: 2},
			expectedCompare: 1,
		},
		{
			name:            "Equal_SameRepresentative",
			a:               group{States: NewStates(2, 4, 8), Rep: 2},
			b:               group{States: NewStates(2, 4, 8), Rep: 2},
			expectedCompare: 0,
		},
		{
			name:            "Equal_DiffRepresentative",
			a:               group{States: NewStates(2, 4, 8), Rep: 2},
			b:               group{States: NewStates(2, 4, 8), Rep: 4},
			expectedCompare: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCompare, cmpGroup(tc.a, tc.b))
		})
	}
}

func TestHashGroup(t *testing.T) {
	tests := []struct {
		name         string
		g            group
		expectedHash uint64
	}{
		{
			name:         "OK",
			g:            group{States: NewStates(2, 4, 8), Rep: 2},
			expectedHash: 0xf36910cafdc3d8bb,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedHash, hashGroup(tc.g))
		})
	}
}

func TestNewGroups(t *testing.T) {
	tests := []struct {
		name           string
		gs             []group
		expectedString string
	}{
		{
			name: "OK",
			gs: []group{
				{States: NewStates(2, 4, 8), Rep: 2},
				{States: NewStates(3, 5, 7), Rep: 3},
			},
			expectedString: "{{2, 4, 8}, {3, 5, 7}}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gs := newGroups(tc.gs...)

			assert.NotNil(t, gs)
			assert.Equal(t, tc.expectedString, gs.String())
		})
	}
}

func TestNewPartition(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		P := newPartition()

		assert.NotNil(t, P)
		assert.NotNil(t, P.groups)
	})
}

func TestPartition_Equal(t *testing.T) {
	tests := []struct {
		name          string
		p             *partition
		rhs           *partition
		expectedEqual bool
	}{
		{
			name: "Equal",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
				),
				nextRep: 2,
			},
			rhs: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
				),
				nextRep: 2,
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
				),
				nextRep: 2,
			},
			rhs: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1), Rep: 1},
					group{States: NewStates(2), Rep: 2},
				),
				nextRep: 3,
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestPartition_Add(t *testing.T) {
	tests := []struct {
		name              string
		p                 *partition
		groups            []States
		expectedPartition *partition
	}{
		{
			name: "OK",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
				),
				nextRep: 1,
			},
			groups: []States{
				NewStates(1, 2),
				NewStates(3, 4),
			},
			expectedPartition: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
					group{States: NewStates(3, 4), Rep: 2},
				),
				nextRep: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Add(tc.groups...)

			assert.True(t, tc.p.Equal(tc.expectedPartition))
		})
	}
}

func TestPartition_FindRep(t *testing.T) {
	tests := []struct {
		name        string
		p           *partition
		s           State
		expectedRep State
	}{
		{
			name: "Found",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
					group{States: NewStates(3, 4), Rep: 2},
				),
				nextRep: 4,
			},
			s:           State(4),
			expectedRep: State(2),
		},
		{
			name: "NotFound",
			p: &partition{
				groups: newGroups(
					group{States: NewStates(0), Rep: 0},
					group{States: NewStates(1, 2), Rep: 1},
					group{States: NewStates(3, 4), Rep: 2},
				),
				nextRep: 4,
			},
			s:           State(8),
			expectedRep: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rep := tc.p.FindRep(tc.s)

			assert.Equal(t, tc.expectedRep, rep)
		})
	}
}
