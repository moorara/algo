package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestNFAs() []*NFA {
	// aa*|bb*
	n0 := NewNFA(0, States{2, 4})
	n0.Add(0, E, States{1, 3})
	n0.Add(1, 'a', States{2})
	n0.Add(2, 'a', States{2})
	n0.Add(3, 'b', States{4})
	n0.Add(4, 'b', States{4})

	// (a|b)*abb
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

	// ab+|ba+
	n2 := NewNFA(0, States{2, 4})
	n2.Add(0, 'a', States{1})
	n2.Add(1, 'b', States{2})
	n2.Add(2, 'b', States{2})
	n2.Add(0, 'b', States{3})
	n2.Add(3, 'a', States{4})
	n2.Add(4, 'a', States{4})

	// (ab)+
	n3 := NewNFA(0, States{2})
	n3.Add(0, 'a', States{1})
	n3.Add(1, 'b', States{2})
	n3.Add(2, 'a', States{1})

	// (ab+|ba+)*
	n4 := NewNFA(0, States{1})
	n4.Add(0, E, States{1})
	n4.Add(0, E, States{2})
	n4.Add(2, 'a', States{3})
	n4.Add(2, 'b', States{4})
	n4.Add(3, 'b', States{5})
	n4.Add(4, 'a', States{6})
	n4.Add(5, 'b', States{5})
	n4.Add(6, 'a', States{6})
	n4.Add(5, E, States{1})
	n4.Add(5, E, States{2})
	n4.Add(6, E, States{1})
	n4.Add(6, E, States{2})

	// ab+|ba+|(ab)+
	n5 := NewNFA(0, States{1})
	n5.Add(0, E, States{2})
	n5.Add(0, E, States{7})
	n5.Add(2, 'a', States{3})
	n5.Add(2, 'b', States{4})
	n5.Add(3, 'b', States{5})
	n5.Add(4, 'a', States{6})
	n5.Add(5, 'b', States{5})
	n5.Add(5, E, States{1})
	n5.Add(6, 'a', States{6})
	n5.Add(6, E, States{1})
	n5.Add(7, 'a', States{8})
	n5.Add(8, 'b', States{9})
	n5.Add(9, 'a', States{8})
	n5.Add(9, E, States{1})

	// (ab+|ba+)(ab)+
	n6 := NewNFA(0, States{6})
	n6.Add(0, 'a', States{1})
	n6.Add(0, 'b', States{2})
	n6.Add(1, 'b', States{3})
	n6.Add(2, 'a', States{4})
	n6.Add(3, 'a', States{5})
	n6.Add(3, 'b', States{3})
	n6.Add(4, 'a', States{4})
	n6.Add(4, 'a', States{5})
	n6.Add(5, 'b', States{6})
	n6.Add(6, 'a', States{5})

	return []*NFA{n0, n1, n2, n3, n4, n5, n6}
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
			name:           "Empty",
			n:              NewNFA(0, States{1, 2}),
			expectedStates: States{0, 1, 2},
		},
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
			assert.True(t, tc.n.States().Equal(tc.expectedStates))
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

func TestNFA_Star(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		n           *NFA
		expectedNFA *NFA
	}{
		{
			name:        "OK",
			n:           nfas[2],
			expectedNFA: nfas[4],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.n.Star()
			assert.True(t, nfa.Equal(tc.expectedNFA))
		})
	}
}

func TestNFA_Union(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		n           *NFA
		ns          []*NFA
		expectedNFA *NFA
	}{
		{
			name:        "OK",
			n:           nfas[2],
			ns:          nfas[3:4],
			expectedNFA: nfas[5],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.n.Union(tc.ns...)
			assert.True(t, nfa.Equal(tc.expectedNFA))
		})
	}
}

func TestNFA_Concat(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		n           *NFA
		ns          []*NFA
		expectedNFA *NFA
	}{
		{
			name:        "OK",
			n:           nfas[2],
			ns:          nfas[3:4],
			expectedNFA: nfas[6],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.n.Concat(tc.ns...)
			assert.True(t, nfa.Equal(tc.expectedNFA))
		})
	}
}

func TestNFA_ToDFA(t *testing.T) {
	nfas := getTestNFAs()
	dfas := getTestDFAs()

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
			assert.True(t, dfa.Equal(tc.expectedDFA))
		})
	}
}

func TestNFA_Equal(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name          string
		n             *NFA
		rhs           *NFA
		expectedEqual bool
	}{
		{
			name:          "Equal",
			n:             nfas[0],
			rhs:           nfas[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			n:             nfas[0],
			rhs:           nfas[1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.n.Equal(tc.rhs))
		})
	}
}

func TestNFA_Isomorphic(t *testing.T) {
	// aa*|bb*
	n0 := NewNFA(0, States{2, 4})
	n0.Add(0, E, States{1, 3})
	n0.Add(1, 'a', States{2})
	n0.Add(2, 'a', States{2})
	n0.Add(3, 'b', States{4})
	n0.Add(4, 'b', States{4})

	n1 := NewNFA(0, States{2})

	n2 := NewNFA(0, States{3, 6})
	n2.Add(0, E, States{1, 4})
	n2.Add(1, 'a', States{2})
	n2.Add(2, 'a', States{3})
	n2.Add(3, 'a', States{3})
	n2.Add(4, 'b', States{5})
	n2.Add(5, 'b', States{6})
	n2.Add(6, 'b', States{6})

	n3 := NewNFA(0, States{2, 4})
	n3.Add(0, E, States{1, 3})
	n3.Add(1, 'a', States{2})
	n3.Add(2, 'c', States{2})
	n3.Add(3, 'b', States{4})
	n3.Add(4, 'c', States{4})

	n4 := NewNFA(0, States{2, 4})
	n4.Add(0, E, States{1, 3})
	n4.Add(1, 'a', States{2})
	n4.Add(2, 'a', States{2})
	n4.Add(3, 'b', States{4})

	n5 := NewNFA(4, States{0, 1})
	n5.Add(4, E, States{2, 3})
	n5.Add(2, 'a', States{0})
	n5.Add(0, 'a', States{0})
	n5.Add(3, 'b', States{1})
	n5.Add(1, 'b', States{1})

	tests := []struct {
		name               string
		n                  *NFA
		rhs                *NFA
		expectedIsomorphic bool
	}{
		{
			name:               "FinalStatesNotEqualSize",
			n:                  n0,
			rhs:                n1,
			expectedIsomorphic: false,
		},
		{
			name:               "StatesNotEqualSize",
			n:                  n0,
			rhs:                n2,
			expectedIsomorphic: false,
		},
		{
			name:               "SymbolsNotEqual",
			n:                  n0,
			rhs:                n3,
			expectedIsomorphic: false,
		},
		{
			name:               "DegreesNotEqual",
			n:                  n0,
			rhs:                n4,
			expectedIsomorphic: false,
		},
		{
			name:               "Equal",
			n:                  n0,
			rhs:                n0,
			expectedIsomorphic: true,
		},
		{
			name:               "Isomorphic",
			n:                  n0,
			rhs:                n5,
			expectedIsomorphic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsomorphic, tc.n.Isomorphic(tc.rhs))
		})
	}
}

func TestNFA_DOT(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name        string
		n           *NFA
		expectedDOT string
	}{
		{
			name:        "Empty",
			n:           NewNFA(0, States{1, 2}),
			expectedDOT: nfaEmpty,
		},
		{
			name:        "First",
			n:           nfas[0],
			expectedDOT: nfa01,
		},
		{
			name:        "Second",
			n:           nfas[1],
			expectedDOT: nfa02,
		},
		{
			name:        "Third",
			n:           nfas[3],
			expectedDOT: nfa03,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.n.DOT())
		})
	}
}

var nfaEmpty = `digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=doublecircle];
  2 [label="2", shape=doublecircle];

  start -> 0 [];
}`

var nfa01 = `digraph "NFA" {
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
}`

var nfa02 = `digraph "NFA" {
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
}`

var nfa03 = `digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 1 [label="a"];
}`
