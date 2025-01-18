package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stateRep string

func (s stateRep) Equals(rhs stateRep) bool {
	return s == rhs
}

func TestAction_For(t *testing.T) {
	tests := []struct {
		name          string
		m             StateMap[stateRep]
		v             stateRep
		expectedState State
	}{
		{
			name:          "OK",
			m:             []stateRep{"S₀", "S₁", "S₂", "S₃", "S₄", "S₅", "S₆", "S₇", "S₈", "S₉"},
			v:             "S₇",
			expectedState: State(7),
		},
		{
			name:          "Error",
			m:             []stateRep{"S₀", "S₁", "S₂", "S₃", "S₄", "S₅", "S₆", "S₇", "S₈", "S₉"},
			v:             "S₁₀",
			expectedState: ErrState,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedState, tc.m.For(tc.v))
	}
}

func TestAction_All(t *testing.T) {
	tests := []struct {
		name           string
		m              StateMap[stateRep]
		expectedStates []State
	}{
		{
			name:           "OK",
			m:              []stateRep{"S₀", "S₁", "S₂", "S₃", "S₄", "S₅", "S₆", "S₇", "S₈", "S₉"},
			expectedStates: []State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedStates, tc.m.All())
	}
}
