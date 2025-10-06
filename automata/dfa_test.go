package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

var testDFA = []*DFA{
	{
		start: 0,
		final: NewStates(1),
		trans: newDFATransitionTable(
			map[State][]rangeState{
				0: {
					{SymbolRange{Start: '1', End: '1'}, 0},
				},
				1: {
					{SymbolRange{Start: '0', End: '1'}, 1},
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
				0: {SymbolRange{Start: '1', End: '1'}: 0},
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
  0 --[1]--> 0
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
	dfa.states = NewStates(0, 1)
	dfa.symbols = NewSymbols('0', '1')

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
		expectedSymbols []Symbol
	}{
		{
			name:            "OK",
			d:               testDFA[0],
			expectedSymbols: []Symbol{'0', '1'},
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
				{0, SymbolRange{Start: '1', End: '1'}, 0},
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
				{Key: SymbolRange{Start: '1', End: '1'}, Val: 0},
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
