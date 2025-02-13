package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStates(t *testing.T) {
	tests := []struct {
		name string
		s    []State
	}{
		{
			name: "OK",
			s:    []State{0, 1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewStates(tc.s...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.s...))
		})
	}
}

func TestNewSymbols(t *testing.T) {
	tests := []struct {
		name string
		a    []Symbol
	}{
		{
			name: "OK",
			a:    []Symbol{'a', 'b', 'c', 'd'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewSymbols(tc.a...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.a...))
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name           string
		s              string
		expectedString String
	}{
		{
			name:           "OK",
			s:              "ababb",
			expectedString: String{'a', 'b', 'a', 'b', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, toString(tc.s))
		})
	}
}

func TestStateManager(t *testing.T) {
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
			sm := newStateManager(tc.last)
			state := sm.GetOrCreateState(tc.id, tc.s)
			assert.Equal(t, tc.expectedState, state)
		})
	}
}

func TestGeneratePermutations(t *testing.T) {
	tests := []struct {
		name                 string
		states               []State
		start                int
		end                  int
		expectedReturn       bool
		expectedPermutations [][]State
	}{
		{
			name:   "OK",
			states: []State{0, 1, 2},
			start:  0,
			end:    2,
			expectedPermutations: [][]State{
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
			assert.True(t, generatePermutations(tc.states, tc.start, tc.end, func(perm []State) bool {
				assert.Contains(t, tc.expectedPermutations, perm)
				return true
			}))
		})
	}
}
