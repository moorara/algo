package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

var actions = []Action{
	{
		Type:  SHIFT,
		State: 5,
	},
	{
		Type:  SHIFT,
		State: 7,
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
	{
		Type: ACCEPT,
	},
	{
		Type: ERROR,
	},
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
			rhs:            actions[2],
			expectedEquals: false,
		},
		{
			name:           "DifferentStates",
			lhs:            actions[0],
			rhs:            actions[1],
			expectedEquals: false,
		},
		{
			name:           "DifferentProductions",
			lhs:            actions[2],
			rhs:            actions[3],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEquals, eqAction(tc.lhs, tc.rhs))
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
			rhs:            actions[2],
			expectedEquals: false,
		},
		{
			name:           "DifferentStates",
			a:              actions[0],
			rhs:            actions[1],
			expectedEquals: false,
		},
		{
			name:           "DifferentProductions",
			a:              actions[2],
			rhs:            actions[3],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expectedEquals, tc.a.Equals(tc.rhs))
	}
}

func TestCmpAction(t *testing.T) {
	tests := []struct {
		name            string
		lhs             Action
		rhs             Action
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
