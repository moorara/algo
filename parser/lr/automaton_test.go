package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

func TestAugment(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		expectedCFG *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "*", "(", ")", "id"},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				prods[2],
				"E′",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			augG := augment(tc.G)
			assert.True(t, augG.Equals(tc.expectedCFG))
		})
	}
}

func TestAutomaton_GOTO(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name         string
		a            *automaton
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name: `GOTO(I₀,E)`,
			a: &automaton{
				Calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			I:            s[0],
			X:            grammar.NonTerminal("E"),
			expectedGOTO: s[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.a.GOTO(tc.I, tc.X)
			assert.True(t, J.Equals(tc.expectedGOTO))
		})
	}
}

func TestAutomaton_Canonical(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name              string
		a                 *automaton
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &automaton{
				Calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			expectedCanonical: NewItemSetCollection(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.a.Canonical()
			assert.True(t, C.Equals(tc.expectedCanonical))
		})
	}
}

func TestNewLR0Automaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR0Automaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*automaton).Calculator)
			assert.NotEmpty(t, calc.(*automaton).Calculator.(*calculator0).augG)
		})
	}
}

func TestNewLR1Automaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR1Automaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*automaton).Calculator)
			assert.NotEmpty(t, calc.(*automaton).Calculator.(*calculator1).augG)
			assert.NotNil(t, calc.(*automaton).Calculator.(*calculator1).FIRST)
		})
	}
}
