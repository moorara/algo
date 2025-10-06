package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

var testNFA = []*NFA{
	{
		start: 0,
		final: NewStates(2, 4),
		trans: newNFATransitionTable(
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
		),
	},
}

func TestNFABuilder(t *testing.T) {
	tests := []struct {
		name        string
		start       State
		final       []State
		trans       map[State]map[SymbolRange][]State
		expectedNFA *NFA
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{2, 4},
			trans: map[State]map[SymbolRange][]State{
				0: {SymbolRange{Start: E, End: E}: []State{1, 3}},
				1: {SymbolRange{Start: 'a', End: 'a'}: []State{2}},
				2: {SymbolRange{Start: 'a', End: 'a'}: []State{2}},
				3: {SymbolRange{Start: 'b', End: 'b'}: []State{4}},
				4: {SymbolRange{Start: 'b', End: 'b'}: []State{4}},
			},
			expectedNFA: testNFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(NFABuilder).SetStart(tc.start).SetFinal(tc.final...)
			for s, sub := range tc.trans {
				for r, next := range sub {
					b.AddTransition(s, r.Start, r.End, next)
				}
			}

			t.Run("Build", func(t *testing.T) {
				assert.True(t, b.Build().Equal(tc.expectedNFA))
			})
		})
	}
}

func TestNFA_String(t *testing.T) {
	tests := []struct {
		name           string
		n              *NFA
		expectedString string
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedString: `Start state: 0
Final states: 2, 4
Transitions:
  0 --[ε]--> {1, 3}
  1 --[a]--> {2}
  2 --[a]--> {2}
  3 --[b]--> {4}
  4 --[b]--> {4}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.n.String())
		})
	}
}

func TestNFA_Clone(t *testing.T) {
	nfa := testNFA[0].Clone()
	nfa.states = NewStates(0, 1, 2, 3, 4)
	nfa.symbols = NewSymbols('a', 'b')

	tests := []struct {
		name string
		n    *NFA
	}{
		{
			name: "OK",
			n:    nfa,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.n.Clone()

			assert.NotSame(t, clone, tc.n)
			assert.True(t, clone.Equal(tc.n))
		})
	}
}

func TestNFA_Equal(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		rhs           *NFA
		expectedEqual bool
	}{
		{
			name:          "Equal",
			n:             testNFA[0],
			rhs:           testNFA[0].Clone(),
			expectedEqual: true,
		},
		{
			name: "NotEqual_DiffStart",
			n:    testNFA[0],
			rhs: &NFA{
				start: 1,
				final: NewStates(),
				trans: newNFATransitionTable(nil),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffFinal",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(0),
				trans: newNFATransitionTable(nil),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffTran",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				trans: newNFATransitionTable(nil),
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.n.Equal(tc.rhs))
		})
	}
}

func TestNFA_Start(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		expectedStart State
	}{
		{
			name:          "OK",
			n:             testNFA[0],
			expectedStart: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStart, tc.n.Start())
		})
	}
}

func TestNFA_Final(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		expectedFinal []State
	}{
		{
			name:          "OK",
			n:             testNFA[0],
			expectedFinal: []State{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedFinal, tc.n.Final())
		})
	}
}

func TestNFA_States(t *testing.T) {
	tests := []struct {
		name           string
		n              *NFA
		expectedStates []State
	}{

		{
			name:           "OK",
			n:              testNFA[0],
			expectedStates: []State{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.n.States())
		})
	}
}

func TestNFA_Symbols(t *testing.T) {
	tests := []struct {
		name            string
		n               *NFA
		expectedSymbols []Symbol
	}{
		{
			name:            "OK",
			n:               testNFA[0],
			expectedSymbols: []Symbol{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.n.Symbols())
		})
	}
}

func TestNFA_Transitions(t *testing.T) {
	type transition struct {
		s    State
		r    SymbolRange
		next []State
	}

	tests := []struct {
		name                string
		n                   *NFA
		expectedTransitions []transition
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedTransitions: []transition{
				{0, SymbolRange{Start: E, End: E}, []State{1, 3}},
				{1, SymbolRange{Start: 'a', End: 'a'}, []State{2}},
				{2, SymbolRange{Start: 'a', End: 'a'}, []State{2}},
				{3, SymbolRange{Start: 'b', End: 'b'}, []State{4}},
				{4, SymbolRange{Start: 'b', End: 'b'}, []State{4}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			all := []transition{}
			for s, pairs := range tc.n.Transitions() {
				for r, next := range pairs {
					all = append(all, transition{s, r, next})
				}
			}

			assert.Equal(t, tc.expectedTransitions, all)
		})
	}
}

func TestNFA_TransitionsFrom(t *testing.T) {
	tests := []struct {
		name                    string
		n                       *NFA
		s                       State
		expectedTransitionsFrom []generic.KeyValue[SymbolRange, []State]
	}{
		{
			name: "OK",
			n:    testNFA[0],
			s:    0,
			expectedTransitionsFrom: []generic.KeyValue[SymbolRange, []State]{
				{Key: SymbolRange{Start: E, End: E}, Val: []State{1, 3}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			from := generic.Collect2(tc.n.TransitionsFrom(tc.s))
			assert.Equal(t, tc.expectedTransitionsFrom, from)
		})
	}
}
