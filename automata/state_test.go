package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStates_Contains(t *testing.T) {
	tests := []struct {
		name           string
		s              States
		t              State
		expectedResult bool
	}{
		{
			name:           "No",
			s:              States{2, 4},
			t:              3,
			expectedResult: false,
		},
		{
			name:           "Yes",
			s:              States{2, 4},
			t:              4,
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.s.Contains(tc.t))
		})
	}
}

func TestStates_Equals(t *testing.T) {
	tests := []struct {
		name           string
		s              States
		rhs            States
		expectedEquals bool
	}{
		{
			name:           "Equal",
			s:              States{2, 4},
			rhs:            States{4, 2},
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			s:              States{2, 4},
			rhs:            States{2},
			expectedEquals: false,
		},
		{
			name:           "NotEqual",
			s:              States{2},
			rhs:            States{2, 4},
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.s.Equals(tc.rhs))
		})
	}
}

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		name                 string
		states               States
		start                int
		end                  int
		expectedReturn       bool
		expectedPermutations []States
	}{
		{
			name:   "OK",
			states: States{0, 1, 2},
			start:  0,
			end:    2,
			expectedPermutations: []States{
				{0, 1, 2},
				{0, 2, 1},
				{1, 0, 2},
				{1, 2, 0},
				{2, 1, 0},
				{2, 0, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, generatePermutations(tc.states, tc.start, tc.end, func(perm States) bool {
				assert.Contains(t, tc.expectedPermutations, perm)
				return true
			}))
		})
	}
}

func TestStateFactory(t *testing.T) {
	tests := []struct {
		name          string
		last          State
		id            int
		s             State
		expectedState State
	}{
		{
			name:          "OK",
			last:          10,
			id:            0,
			s:             1,
			expectedState: 11,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			factory := newStateFactory(tc.last)
			state := factory.StateFor(tc.id, tc.s)
			assert.Equal(t, tc.expectedState, state)
		})
	}
}
