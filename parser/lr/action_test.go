package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqAction(t *testing.T) {
	tests := []struct {
		name          string
		lhs           *Action
		rhs           *Action
		expectedEqual bool
	}{
		{
			name:          "Equal",
			lhs:           actions[0],
			rhs:           actions[0],
			expectedEqual: true,
		},
		{
			name:          "DifferentTypes",
			lhs:           actions[0],
			rhs:           actions[2],
			expectedEqual: false,
		},
		{
			name:          "DifferentStates",
			lhs:           actions[0],
			rhs:           actions[1],
			expectedEqual: false,
		},
		{
			name:          "DifferentProductions",
			lhs:           actions[2],
			rhs:           actions[3],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEqual, eqAction(tc.lhs, tc.rhs))
	}
}

func TestAction_String(t *testing.T) {
	tests := []struct {
		name           string
		a              *Action
		expectedString string
	}{
		{
			name:           "SHIFT",
			a:              actions[0],
			expectedString: "SHIFT 5",
		},
		{
			name:           "REDUCE",
			a:              actions[2],
			expectedString: "REDUCE E → T",
		},
		{
			name:           "ACCEPT",
			a:              actions[4],
			expectedString: "ACCEPT",
		},
		{
			name:           "ERROR",
			a:              actions[5],
			expectedString: "ERROR",
		},
		{
			name:           "INVALID",
			a:              &Action{},
			expectedString: "INVALID ACTION(0)",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedString, tc.a.String())
	}
}

func TestAction_Equal(t *testing.T) {
	tests := []struct {
		name          string
		a             *Action
		rhs           *Action
		expectedEqual bool
	}{
		{
			name:          "Equal",
			a:             actions[0],
			rhs:           actions[0],
			expectedEqual: true,
		},
		{
			name:          "DifferentTypes",
			a:             actions[0],
			rhs:           actions[2],
			expectedEqual: false,
		},
		{
			name:          "DifferentStates",
			a:             actions[0],
			rhs:           actions[1],
			expectedEqual: false,
		},
		{
			name:          "DifferentProductions",
			a:             actions[2],
			rhs:           actions[3],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEqual, tc.a.Equal(tc.rhs))
	}
}

func TestCmpAction(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *Action
		rhs             *Action
		expectedCompare int
	}{
		{
			name:            "ByStates",
			lhs:             actions[0],
			rhs:             actions[1],
			expectedCompare: -2,
		},
		{
			name:            "ByProductions",
			lhs:             actions[2],
			rhs:             actions[3],
			expectedCompare: -1,
		},
		{
			name:            "ByTypes",
			lhs:             actions[4],
			rhs:             actions[5],
			expectedCompare: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmp := cmpAction(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expectedCompare, cmp)
		})
	}
}
