package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testNFA = []*NFA{
	// (a+|b+)
	{
		start: 0,
		final: NewStates(2, 4),
		/* trans: newNFATransitionTable(
			map[State][]rangeStates{
				0: {
					{SymbolRange{Start: E, End: E}, NewStates(1, 3)},
				},
				1: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(2)},
				},
				2: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(2)},
				},
				3: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(4)},
				},
				4: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(4)},
				},
			},
		), */
	},
	// ab+|ba+
	{
		start: 0,
		final: NewStates(2, 4),
		/* trans: newNFATransitionTable(
			map[State][]rangeStates{
				0: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(1)},
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(3)},
				},
				1: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(2)},
				},
				2: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(2)},
				},
				3: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(4)},
				},
				4: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(4)},
				},
			},
		), */
	},
	// (a|b)*abb
	{
		start: 0,
		final: NewStates(10),
		/* trans: newNFATransitionTable(
			map[State][]rangeStates{
				0: {
					{SymbolRange{Start: E, End: E}, NewStates(1, 7)},
				},
				1: {
					{SymbolRange{Start: E, End: E}, NewStates(2, 4)},
				},
				2: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(3)},
				},
				3: {
					{SymbolRange{Start: E, End: E}, NewStates(6)},
				},
				4: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(5)},
				},
				5: {
					{SymbolRange{Start: E, End: E}, NewStates(6)},
				},
				6: {
					{SymbolRange{Start: E, End: E}, NewStates(1, 7)},
				},
				7: {
					{SymbolRange{Start: 'a', End: 'a'}, NewStates(8)},
				},
				8: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(9)},
				},
				9: {
					{SymbolRange{Start: 'b', End: 'b'}, NewStates(10)},
				},
			},
		), */
	},
}

func TestNFABuilder(t *testing.T) {
	type transition struct {
		s          State
		start, end Symbol
		next       []State
	}

	tests := []struct {
		name        string
		start       State
		final       []State
		trans       []transition
		expectedNFA *NFA
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{2, 4},
			trans: []transition{
				{s: 0, start: E, end: E, next: []State{1, 3}},
				{s: 1, start: 'a', end: 'a', next: []State{2}},
				{s: 2, start: 'a', end: 'a', next: []State{2}},
				{s: 3, start: 'b', end: 'b', next: []State{4}},
				{s: 4, start: 'b', end: 'b', next: []State{4}},
			},
			expectedNFA: testNFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(NFABuilder).SetStart(tc.start).SetFinal(tc.final...)

			for _, tr := range tc.trans {
				b.AddTransition(tr.s, tr.start, tr.end, tr.next)
			}

			t.Run("Build", func(t *testing.T) {
				assert.True(t, b.Build().Equal(tc.expectedNFA))
			})
		})
	}
}
