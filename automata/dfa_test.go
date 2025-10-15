package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDFA = []*DFA{
	// 1(0|1)*
	{
		start: 0,
		final: NewStates(1),
		/* trans: newDFATransitionTable(
			map[State][]rangeState{
				0: {
					{SymbolRange{Start: '1', End: '1'}, 1},
				},
				1: {
					{SymbolRange{Start: '0', End: '1'}, 1},
				},
			},
		), */
	},
	// ab+|ba+
	{
		start: 0,
		final: NewStates(2, 4),
		/* trans: newDFATransitionTable(
			map[State][]rangeState{
				0: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 3},
				},
				1: {
					{SymbolRange{Start: 'b', End: 'b'}, 2},
				},
				2: {
					{SymbolRange{Start: 'b', End: 'b'}, 2},
				},
				3: {
					{SymbolRange{Start: 'a', End: 'a'}, 4},
				},
				4: {
					{SymbolRange{Start: 'a', End: 'a'}, 4},
				},
			},
		), */
	},
	// (a|b)*abb
	{
		start: 0,
		final: NewStates(4),
		/* trans: newDFATransitionTable(
			map[State][]rangeState{
				0: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 2},
				},
				1: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 3},
				},
				2: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 2},
				},
				3: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 4},
				},
				4: {
					{SymbolRange{Start: 'a', End: 'a'}, 1},
					{SymbolRange{Start: 'b', End: 'b'}, 2},
				},
			},
		), */
	},
}

func TestDFABuilder(t *testing.T) {
	type transition struct {
		s          State
		start, end Symbol
		next       State
	}

	tests := []struct {
		name        string
		start       State
		final       []State
		trans       []transition
		expectedDFA *DFA
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{1},
			trans: []transition{
				{s: 0, start: '1', end: '1', next: 1},
				{s: 1, start: '0', end: '1', next: 1},
			},
			expectedDFA: testDFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(DFABuilder).SetStart(tc.start).SetFinal(tc.final...)

			for _, tr := range tc.trans {
				b.AddTransition(tr.s, tr.start, tr.end, tr.next)
			}

			t.Run("Build", func(t *testing.T) {
				assert.True(t, b.Build().Equal(tc.expectedDFA))
			})
		})
	}
}
