package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestNFAs() []*NFA {
	n0 := NewNFA(0, States{2, 4})
	n0.Add(0, E, States{1, 3})
	n0.Add(1, 'a', States{2})
	n0.Add(2, 'a', States{2})
	n0.Add(3, 'b', States{4})
	n0.Add(4, 'b', States{4})

	n1 := NewNFA(0, States{10})
	n1.Add(0, E, States{1, 7})
	n1.Add(1, E, States{2, 4})
	n1.Add(2, 'a', States{3})
	n1.Add(3, E, States{6})
	n1.Add(4, 'b', States{5})
	n1.Add(5, E, States{6})
	n1.Add(6, E, States{1, 7})
	n1.Add(7, 'a', States{8})
	n1.Add(8, 'b', States{9})
	n1.Add(9, 'b', States{10})

	n2 := NewNFA(0, States{15})
	n2.Add(0, E, States{1, 3})
	n2.Add(1, 'a', States{2})
	n2.Add(2, 'a', States{2})
	n2.Add(3, 'b', States{4})
	n2.Add(4, 'b', States{4})
	n2.Add(2, E, States{5})
	n2.Add(4, E, States{5})
	n2.Add(5, E, States{6, 12})
	n2.Add(6, E, States{7, 9})
	n2.Add(7, 'a', States{8})
	n2.Add(8, E, States{11})
	n2.Add(9, 'b', States{10})
	n2.Add(10, E, States{11})
	n2.Add(11, E, States{6, 12})
	n2.Add(12, 'a', States{13})
	n2.Add(13, 'b', States{14})
	n2.Add(14, 'b', States{15})

	return []*NFA{n0, n1, n2}
}

func TestNewNFA(t *testing.T) {
	n := getTestNFAs()[0]

	nfa := NewNFA(n.Start, n.Final)
	assert.NotNil(t, nfa)
}

func TestNFA_Add(t *testing.T) {
	nfa := NewNFA(0, States{1, 2, 3, 4})

	tests := []struct {
		name string
		n    *NFA
		s    State
		a    Symbol
		next States
	}{
		{
			name: "NewState",
			n:    nfa,
			s:    State(0),
			a:    'a',
			next: States{1, 2},
		},
		{
			name: "ExistingState",
			n:    nfa,
			s:    State(0),
			a:    'b',
			next: States{3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.n.Add(tc.s, tc.a, tc.next)
		})
	}
}

func TestNFA_Next(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		n              *NFA
		s              State
		a              Symbol
		expectedStates States
	}{
		{
			name:           "First",
			n:              nfas[0],
			s:              State(0),
			a:              E,
			expectedStates: States{1, 3},
		},
		{
			name:           "Second",
			n:              nfas[1],
			s:              State(1),
			a:              E,
			expectedStates: States{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states := tc.n.Next(tc.s, tc.a)
			assert.Equal(t, tc.expectedStates, states)
		})
	}
}

func TestNFA_States(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		n              *NFA
		expectedStates States
	}{
		{
			name:           "First",
			n:              nfas[0],
			expectedStates: States{0, 1, 2, 3, 4},
		},
		{
			name:           "Second",
			n:              nfas[1],
			expectedStates: States{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.n.States())
		})
	}
}

func TestNFA_LastState(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name              string
		n                 *NFA
		expectedLastState State
	}{
		{
			name:              "First",
			n:                 nfas[0],
			expectedLastState: State(4),
		},
		{
			name:              "Second",
			n:                 nfas[1],
			expectedLastState: State(10),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedLastState, tc.n.LastState())
		})
	}
}

func TestNFA_Symbols(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name            string
		n               *NFA
		expectedSymbols Symbols
	}{
		{
			name:            "First",
			n:               nfas[0],
			expectedSymbols: Symbols{'a', 'b'},
		},
		{
			name:            "Second",
			n:               nfas[1],
			expectedSymbols: Symbols{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.n.Symbols())
		})
	}
}

func TestNFA_Join(t *testing.T) {
	nfas := getTestNFAs()

	type edge struct {
		s    State
		a    Symbol
		next States
	}

	tests := []struct {
		name           string
		n              *NFA
		nfa            *NFA
		newFinal       States
		extraTrans     []edge
		expectedStates States
		expectedStart  State
		expectedFinal  States
		expectedNFA    *NFA
	}{
		{
			name:     "OK",
			n:        nfas[0],
			nfa:      nfas[1],
			newFinal: States{15},
			extraTrans: []edge{
				{2, E, States{5}},
				{4, E, States{5}},
			},
			expectedStates: States{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expectedStart:  State(5),
			expectedFinal:  States{15},
			expectedNFA:    nfas[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states, start, final := tc.n.Join(tc.nfa)

			tc.n.Final = tc.newFinal
			for _, e := range tc.extraTrans {
				tc.n.Add(e.s, e.a, e.next)
			}

			assert.Equal(t, tc.expectedStates, states)
			assert.Equal(t, tc.expectedStart, start)
			assert.Equal(t, tc.expectedFinal, final)
			assert.True(t, tc.n.Equals(tc.expectedNFA))
		})
	}
}

func TestNFA_Accept(t *testing.T) {
	nfa := getTestNFAs()[0]

	tests := []struct {
		name           string
		n              *NFA
		s              String
		expectedResult bool
	}{
		{
			name:           "Accepted",
			n:              nfa,
			s:              ToString("aaaa"),
			expectedResult: true,
		},
		{
			name:           "NotAccepted",
			n:              nfa,
			s:              ToString("abbb"),
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.n.Accept(tc.s)
			assert.Equal(t, tc.expectedResult, b)
		})
	}
}

func TestNFA_ToDFA(t *testing.T) {
	dfas := getTestDFAs()
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		n           *NFA
		expectedDFA *DFA
	}{
		{
			name:        "OK",
			n:           nfas[1],
			expectedDFA: dfas[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.n.ToDFA()
			assert.True(t, dfa.Equals(tc.expectedDFA))
		})
	}
}

func TestNFA_Equals(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		n              *NFA
		nfa            *NFA
		expectedEquals bool
	}{
		{
			name:           "Equal",
			n:              nfas[0],
			nfa:            nfas[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			n:              nfas[1],
			nfa:            nfas[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.n.Equals(tc.nfa))
		})
	}
}

func TestNFA_Graphviz(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name             string
		n                *NFA
		expectedGraphviz string
	}{
		{
			name: "First",
			n:    nfas[0],
			expectedGraphviz: `strict digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=doublecircle];
  3 [label="3", shape=circle];
  4 [label="4", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="ε"];
  0 -> 3 [label="ε"];
  1 -> 2 [label="a"];
  2 -> 2 [label="a"];
  3 -> 4 [label="b"];
  4 -> 4 [label="b"];
}`,
		},
		{
			name: "Second",
			n:    nfas[1],
			expectedGraphviz: `strict digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=circle];
  3 [label="3", shape=circle];
  4 [label="4", shape=circle];
  5 [label="5", shape=circle];
  6 [label="6", shape=circle];
  7 [label="7", shape=circle];
  8 [label="8", shape=circle];
  9 [label="9", shape=circle];
  10 [label="10", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="ε"];
  0 -> 7 [label="ε"];
  1 -> 2 [label="ε"];
  1 -> 4 [label="ε"];
  2 -> 3 [label="a"];
  3 -> 6 [label="ε"];
  4 -> 5 [label="b"];
  5 -> 6 [label="ε"];
  6 -> 1 [label="ε"];
  6 -> 7 [label="ε"];
  7 -> 8 [label="a"];
  8 -> 9 [label="b"];
  9 -> 10 [label="b"];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedGraphviz, tc.n.Graphviz())
		})
	}
}
