package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestDFAs() []*DFA {
	d0 := NewDFA(0, States{3})
	d0.Add(0, 'a', 1)
	d0.Add(0, 'b', 0)
	d0.Add(1, 'a', 1)
	d0.Add(1, 'b', 2)
	d0.Add(2, 'a', 1)
	d0.Add(2, 'b', 3)
	d0.Add(3, 'a', 1)
	d0.Add(3, 'b', 0)

	d1 := NewDFA(0, States{4})
	d1.Add(0, 'a', 1)
	d1.Add(0, 'b', 2)
	d1.Add(1, 'a', 1)
	d1.Add(1, 'b', 3)
	d1.Add(2, 'a', 1)
	d1.Add(2, 'b', 2)
	d1.Add(3, 'a', 1)
	d1.Add(3, 'b', 4)
	d1.Add(4, 'a', 1)
	d1.Add(4, 'b', 2)

	d2 := NewDFA(0, States{8})
	d2.Add(0, 'a', 1)
	d2.Add(0, 'b', 0)
	d2.Add(1, 'a', 1)
	d2.Add(1, 'b', 2)
	d2.Add(2, 'a', 1)
	d2.Add(2, 'b', 3)
	d2.Add(3, 'a', 1)
	d2.Add(3, 'b', 0)
	d2.Add(3, 'a', 4)
	d2.Add(4, 'a', 5)
	d2.Add(4, 'b', 6)
	d2.Add(5, 'a', 5)
	d2.Add(5, 'b', 7)
	d2.Add(6, 'a', 5)
	d2.Add(6, 'b', 6)
	d2.Add(7, 'a', 5)
	d2.Add(7, 'b', 8)
	d2.Add(8, 'a', 5)
	d2.Add(8, 'b', 6)

	d3 := NewDFA(0, States{3})
	d3.Add(0, 'a', 1)
	d3.Add(0, 'b', 0)
	d3.Add(1, 'a', 1)
	d3.Add(1, 'b', 2)
	d3.Add(2, 'a', 1)
	d3.Add(2, 'b', 3)
	d3.Add(3, 'a', 1)
	d3.Add(3, 'b', 0)

	return []*DFA{d0, d1, d2, d3}
}

func TestNewDFA(t *testing.T) {
	d := getTestDFAs()[0]

	dfa := NewDFA(d.Start, d.Final)
	assert.NotNil(t, dfa)
}

func TestDFA_Add(t *testing.T) {
	dfa := NewDFA(0, States{1, 2})

	tests := []struct {
		name string
		d    *DFA
		s    State
		a    Symbol
		next State
	}{
		{
			name: "NewState",
			d:    dfa,
			s:    State(0),
			a:    'a',
			next: State(1),
		},
		{
			name: "ExistingState",
			d:    dfa,
			s:    State(0),
			a:    'b',
			next: State(2),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.d.Add(tc.s, tc.a, tc.next)
		})
	}
}

func TestDFA_Next(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name          string
		d             *DFA
		s             State
		a             Symbol
		expectedState State
	}{
		{
			name:          "First",
			d:             dfas[0],
			s:             State(0),
			a:             'a',
			expectedState: State(1),
		},
		{
			name:          "Second",
			d:             dfas[1],
			s:             State(0),
			a:             'b',
			expectedState: State(2),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.d.Next(tc.s, tc.a)
			assert.Equal(t, tc.expectedState, state)
		})
	}
}

func TestDFA_States(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		d              *DFA
		expectedStates States
	}{
		{
			name:           "Empty",
			d:              NewDFA(0, States{1}),
			expectedStates: States{0, 1},
		},
		{
			name:           "First",
			d:              dfas[0],
			expectedStates: States{0, 1, 2, 3},
		},
		{
			name:           "Second",
			d:              dfas[1],
			expectedStates: States{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.d.States().Equals(tc.expectedStates))
		})
	}
}

func TestDFA_LastState(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name              string
		d                 *DFA
		expectedLastState State
	}{
		{
			name:              "First",
			d:                 dfas[0],
			expectedLastState: State(3),
		},
		{
			name:              "Second",
			d:                 dfas[1],
			expectedLastState: State(4),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedLastState, tc.d.LastState())
		})
	}
}

func TestDFA_Symbols(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name            string
		d               *DFA
		expectedSymbols Symbols
	}{
		{
			name:            "First",
			d:               dfas[0],
			expectedSymbols: Symbols{'a', 'b'},
		},
		{
			name:            "Second",
			d:               dfas[1],
			expectedSymbols: Symbols{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.d.Symbols())
		})
	}
}

func TestDFA_Join(t *testing.T) {
	dfas := getTestDFAs()

	type edge struct {
		s    State
		a    Symbol
		next State
	}

	tests := []struct {
		name           string
		d              *DFA
		dfa            *DFA
		newFinal       States
		extraTrans     []edge
		expectedStates States
		expectedStart  State
		expectedFinal  States
		expectedDFA    *DFA
	}{
		{
			name:     "OK",
			d:        dfas[0],
			dfa:      dfas[1],
			newFinal: States{8},
			extraTrans: []edge{
				{3, 'a', 4},
			},
			expectedStates: States{4, 5, 6, 7, 8},
			expectedStart:  State(4),
			expectedFinal:  States{8},
			expectedDFA:    dfas[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states, start, final := tc.d.Join(tc.dfa)

			tc.d.Final = tc.newFinal
			for _, e := range tc.extraTrans {
				tc.d.Add(e.s, e.a, e.next)
			}

			assert.True(t, states.Equals(tc.expectedStates))
			assert.Equal(t, tc.expectedStart, start)
			assert.True(t, final.Equals(tc.expectedFinal))
			assert.True(t, tc.d.Equals(tc.expectedDFA))
		})
	}
}

func TestDFA_Accept(t *testing.T) {
	dfa := getTestDFAs()[0]

	tests := []struct {
		name           string
		d              *DFA
		s              String
		expectedResult bool
	}{
		{
			name:           "Accepted",
			d:              dfa,
			s:              ToString("aabbababb"),
			expectedResult: true,
		},
		{
			name:           "NotAccepted",
			d:              dfa,
			s:              ToString("aabab"),
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.d.Accept(tc.s)
			assert.Equal(t, tc.expectedResult, b)
		})
	}
}

func TestDFA_ToNFA(t *testing.T) {
	dfas := getTestDFAs()
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		d           *DFA
		expectedNFA *NFA
	}{
		{
			name:        "OK",
			d:           dfas[3],
			expectedNFA: nfas[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.d.ToNFA()
			assert.True(t, nfa.Equals(tc.expectedNFA))
		})
	}
}

func TestDFA_Minimize(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name:        "OK",
			d:           dfas[1],
			expectedDFA: dfas[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.Minimize()
			assert.True(t, dfa.Equals(tc.expectedDFA))
		})
	}
}

func TestDFA_Equals(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		d              *DFA
		dfa            *DFA
		expectedEquals bool
	}{
		{
			name:           "Equal",
			d:              dfas[0],
			dfa:            dfas[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			d:              dfas[1],
			dfa:            dfas[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.d.Equals(tc.dfa))
		})
	}
}

func TestDFA_Graphviz(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name             string
		d                *DFA
		expectedGraphviz string
	}{
		{
			name: "Empty",
			d:    NewDFA(0, States{1}),
			expectedGraphviz: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=doublecircle];

  start -> 0 [];
}`,
		},
		{
			name: "First",
			d:    dfas[0],
			expectedGraphviz: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=circle];
  3 [label="3", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="a"];
  0 -> 0 [label="b"];
  1 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 1 [label="a"];
  2 -> 3 [label="b"];
  3 -> 1 [label="a"];
  3 -> 0 [label="b"];
}`,
		},
		{
			name: "Second",
			d:    dfas[1],
			expectedGraphviz: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=circle];
  3 [label="3", shape=circle];
  4 [label="4", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="a"];
  0 -> 2 [label="b"];
  1 -> 1 [label="a"];
  1 -> 3 [label="b"];
  2 -> 1 [label="a"];
  2 -> 2 [label="b"];
  3 -> 1 [label="a"];
  3 -> 4 [label="b"];
  4 -> 1 [label="a"];
  4 -> 2 [label="b"];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedGraphviz, tc.d.Graphviz())
		})
	}
}
