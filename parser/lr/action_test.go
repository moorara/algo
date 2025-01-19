package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

var actions = []Action{
	{
		Type:  SHIFT,
		State: StatePtr(4),
	},
	{
		Type: REDUCE,
		Production: &grammar.Production{ // E → T
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")},
		},
	},
	{
		Type: REDUCE,
		Production: &grammar.Production{ // F → id
			Head: "F",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
	},
}

func StatePtr(s State) *State {
	return &s
}

func TestEqAction(t *testing.T) {
	tests := []struct {
		name           string
		lhs            Action
		rhs            Action
		expectedEquals bool
	}{
		{
			name:           "Equal",
			lhs:            actions[0],
			rhs:            actions[0],
			expectedEquals: true,
		},
		{
			name:           "DifferentTypes",
			lhs:            actions[0],
			rhs:            actions[1],
			expectedEquals: false,
		},
		{
			name: "DifferentStates",
			lhs:  actions[0],
			rhs: Action{
				Type:  SHIFT,
				State: StatePtr(6),
			},
			expectedEquals: false,
		},
		{
			name:           "DifferentProductions",
			lhs:            actions[1],
			rhs:            actions[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEquals, EqAction(tc.lhs, tc.rhs))
	}
}

func TestAction_String(t *testing.T) {
	tests := []struct {
		name           string
		a              Action
		expectedString string
	}{
		{
			name:           "SHIFT",
			a:              actions[0],
			expectedString: "SHIFT 4",
		},
		{
			name:           "REDUCE",
			a:              actions[1],
			expectedString: "REDUCE E → T",
		},
		{
			name: "ACCEPT",
			a: Action{
				Type: ACCEPT,
			},
			expectedString: "ACCEPT",
		},
		{
			name: "ERROR",
			a: Action{
				Type: ERROR,
			},
			expectedString: "ERROR",
		},
		{
			name:           "INVALID",
			a:              Action{},
			expectedString: "INVALID ACTION(0)",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedString, tc.a.String())
	}
}

func TestAction_Equals(t *testing.T) {
	tests := []struct {
		name           string
		a              Action
		rhs            Action
		expectedEquals bool
	}{
		{
			name:           "Equal",
			a:              actions[0],
			rhs:            actions[0],
			expectedEquals: true,
		},
		{
			name:           "DifferentTypes",
			a:              actions[0],
			rhs:            actions[1],
			expectedEquals: false,
		},
		{
			name: "DifferentStates",
			a:    actions[0],
			rhs: Action{
				Type:  SHIFT,
				State: StatePtr(6),
			},
			expectedEquals: false,
		},
		{
			name:           "DifferentProductions",
			a:              actions[1],
			rhs:            actions[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEquals, tc.a.Equals(tc.rhs))
	}
}
