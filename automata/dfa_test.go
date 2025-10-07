package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

var testDFA = []*DFA{
	// 1(0|1)*
	{
		start: 0,
		final: NewStates(1),
		trans: newDFATransitionTable(
			map[State][]rangeState{
				0: {
					{SymbolRange{Start: '1', End: '1'}, 1},
				},
				1: {
					{SymbolRange{Start: '0', End: '1'}, 1},
				},
			},
		),
	},
	// ab+|ba+
	{
		start: 0,
		final: NewStates(2, 4),
		trans: newDFATransitionTable(
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
		),
	},
	// (a|b)*abb
	{
		start: 0,
		final: NewStates(4),
		trans: newDFATransitionTable(
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
		),
	},
}

func TestDFABuilder(t *testing.T) {
	tests := []struct {
		name        string
		start       State
		final       []State
		trans       map[State]map[SymbolRange]State
		expectedDFA *DFA
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{1},
			trans: map[State]map[SymbolRange]State{
				0: {SymbolRange{Start: '1', End: '1'}: 1},
				1: {SymbolRange{Start: '0', End: '1'}: 1},
			},
			expectedDFA: testDFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(DFABuilder).SetStart(tc.start).SetFinal(tc.final...)
			for s, sub := range tc.trans {
				for r, next := range sub {
					b.AddTransition(s, r.Start, r.End, next)
				}
			}

			t.Run("Build", func(t *testing.T) {
				assert.True(t, b.Build().Equal(tc.expectedDFA))
			})
		})
	}
}

func TestDFA_String(t *testing.T) {
	tests := []struct {
		name           string
		d              *DFA
		expectedString string
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedString: `Start state: 0
Final states: 1
Transitions:
  0 --[1]--> 1
  1 --[0..1]--> 1
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.d.String())
		})
	}
}

func TestDFA_Clone(t *testing.T) {
	dfa := testDFA[0].Clone()
	dfa.states = []State{0, 1}
	dfa.symbols = []SymbolRange{{Start: '0', End: '1'}}

	tests := []struct {
		name string
		d    *DFA
	}{
		{
			name: "OK",
			d:    dfa,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.d.Clone()

			assert.NotSame(t, clone, tc.d)
			assert.True(t, clone.Equal(tc.d))
		})
	}
}

func TestDFA_Equal(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		rhs           *DFA
		expectedEqual bool
	}{
		{
			name:          "Equal",
			d:             testDFA[0],
			rhs:           testDFA[0].Clone(),
			expectedEqual: true,
		},
		{
			name: "NotEqual_DiffStart",
			d:    testDFA[0],
			rhs: &DFA{
				start: 1,
				final: NewStates(),
				trans: newDFATransitionTable(nil),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffFinal",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(0),
				trans: newDFATransitionTable(nil),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffTrans",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(1),
				trans: newDFATransitionTable(nil),
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.d.Equal(tc.rhs))
		})
	}
}

func TestDFA_Start(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		expectedStart State
	}{
		{
			name:          "OK",
			d:             testDFA[0],
			expectedStart: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStart, tc.d.Start())
		})
	}
}

func TestDFA_Final(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		expectedFinal []State
	}{
		{
			name:          "OK",
			d:             testDFA[0],
			expectedFinal: []State{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedFinal, tc.d.Final())
		})
	}
}

func TestDFA_States(t *testing.T) {
	tests := []struct {
		name           string
		d              *DFA
		expectedStates []State
	}{
		{
			name:           "OK",
			d:              testDFA[0],
			expectedStates: []State{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.d.States())
		})
	}
}

func TestDFA_Symbols(t *testing.T) {
	tests := []struct {
		name            string
		d               *DFA
		expectedSymbols []SymbolRange
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedSymbols: []SymbolRange{
				{'0', '1'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.d.Symbols())
		})
	}
}

func TestDFA_Transitions(t *testing.T) {
	type transition struct {
		s    State
		r    SymbolRange
		next State
	}

	tests := []struct {
		name                string
		d                   *DFA
		expectedTransitions []transition
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedTransitions: []transition{
				{0, SymbolRange{Start: '1', End: '1'}, 1},
				{1, SymbolRange{Start: '0', End: '1'}, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			all := []transition{}
			for s, pairs := range tc.d.Transitions() {
				for r, next := range pairs {
					all = append(all, transition{s, r, next})
				}
			}

			assert.Equal(t, tc.expectedTransitions, all)
		})
	}
}

func TestDFA_TransitionsFrom(t *testing.T) {
	tests := []struct {
		name                    string
		d                       *DFA
		s                       State
		expectedTransitionsFrom []generic.KeyValue[SymbolRange, State]
	}{
		{
			name: "OK",
			d:    testDFA[0],
			s:    0,
			expectedTransitionsFrom: []generic.KeyValue[SymbolRange, State]{
				{Key: SymbolRange{Start: '1', End: '1'}, Val: 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			from := generic.Collect2(tc.d.TransitionsFrom(tc.s))
			assert.Equal(t, tc.expectedTransitionsFrom, from)
		})
	}
}

func TestDFA_DOT(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedDOT string
	}{
		{
			name: "OK",
			d:    testDFA[1],
			expectedDOT: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1"];
  2 [label="2", shape=doublecircle];
  3 [label="3"];
  4 [label="4", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="[a]"];
  0 -> 3 [label="[b]"];
  1 -> 2 [label="[b]"];
  2 -> 2 [label="[b]"];
  3 -> 4 [label="[a]"];
  4 -> 4 [label="[a]"];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.d.DOT())
		})
	}
}
