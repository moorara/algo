package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
)

func TestAugment(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		expectedCFG *grammar.CFG
	}{
		{
			name: "OK",
			G:    parsertest.Grammars[3],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				parsertest.Prods[3],
				"E′",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			augG := augment(tc.G)
			assert.True(t, augG.Equal(tc.expectedCFG))
		})
	}
}

func TestNewGrammarWithLR0(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR0(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &automaton{})
			assert.IsType(t, G.Automaton.(*automaton).calculator, &calculator0{})
		})
	}
}

func TestNewGrammarWithLR1(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR1(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &automaton{})
			assert.IsType(t, G.Automaton.(*automaton).calculator, &calculator1{})
		})
	}
}

func TestNewGrammarWithLR0Kernel(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR0Kernel(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &kernelAutomaton{})
			assert.IsType(t, G.Automaton.(*kernelAutomaton).calculator, &calculator0{})
		})
	}
}

func TestNewGrammarWithLR1Kernel(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR1Kernel(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &kernelAutomaton{})
			assert.IsType(t, G.Automaton.(*kernelAutomaton).calculator, &calculator1{})
		})
	}
}
