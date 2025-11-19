package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAction_String(t *testing.T) {
	tests := []struct {
		name           string
		a              *Action
		expectedString string
	}{
		{
			name:           "ACCEPT",
			a:              actions[0][0],
			expectedString: `ACCEPT`,
		},
		{
			name:           "ERROR",
			a:              actions[0][1],
			expectedString: `ERROR`,
		},
		{
			name:           "SHIFT",
			a:              actions[0][2],
			expectedString: `SHIFT 5`,
		},
		{
			name:           "REDUCE",
			a:              actions[0][4],
			expectedString: `REDUCE E → E "+" E`,
		},
		{
			name:           "INVALID",
			a:              &Action{},
			expectedString: `INVALID ACTION(0)`,
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
			a:             actions[0][2],
			rhs:           actions[0][2],
			expectedEqual: true,
		},
		{
			name:          "DifferentTypes",
			a:             actions[0][2],
			rhs:           actions[0][5],
			expectedEqual: false,
		},
		{
			name:          "DifferentStates",
			a:             actions[0][2],
			rhs:           actions[0][3],
			expectedEqual: false,
		},
		{
			name:          "DifferentProductions",
			a:             actions[0][4],
			rhs:           actions[0][5],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEqual, tc.a.Equal(tc.rhs))
	}
}

func TestEqAction(t *testing.T) {
	tests := []struct {
		name          string
		lhs           *Action
		rhs           *Action
		expectedEqual bool
	}{
		{
			name:          "Equal",
			lhs:           actions[0][2],
			rhs:           actions[0][2],
			expectedEqual: true,
		},
		{
			name:          "DifferentTypes",
			lhs:           actions[0][2],
			rhs:           actions[0][5],
			expectedEqual: false,
		},
		{
			name:          "DifferentStates",
			lhs:           actions[0][2],
			rhs:           actions[0][3],
			expectedEqual: false,
		},
		{
			name:          "DifferentProductions",
			lhs:           actions[0][4],
			rhs:           actions[0][5],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		eq := eqAction(tc.lhs, tc.rhs)
		assert.Equal(t, tc.expectedEqual, eq)
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
			lhs:             actions[0][2],
			rhs:             actions[0][3],
			expectedCompare: -1,
		},
		{
			name:            "ByProductions",
			lhs:             actions[0][4],
			rhs:             actions[0][5],
			expectedCompare: 1,
		},
		{
			name:            "ByTypes",
			lhs:             actions[0][0],
			rhs:             actions[0][1],
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

func TestHashAction(t *testing.T) {
	tests := []struct {
		name         string
		a            *Action
		expectedHash uint64
	}{
		{
			name:         "Accept",
			a:            actions[0][0],
			expectedHash: 0xc2d1a9d27d012431,
		},
		{
			name:         "Error",
			a:            actions[0][1],
			expectedHash: 0x312fe8e6dbe324b4,
		},
		{
			name:         "Shift",
			a:            actions[0][2],
			expectedHash: 0xb78c1f6261a04e4,
		},
		{
			name:         "Reduce",
			a:            actions[0][4],
			expectedHash: 0xa88475f72fafcd4b,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hash := hashAction(tc.a)
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}
