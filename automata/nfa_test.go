package automata

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/stretchr/testify/assert"
)

func getTestNFAs() []*NFA {
	// aa*|bb*
	n0 := NewNFA(0, []State{2, 4})
	n0.Add(0, E, []State{1, 3})
	n0.Add(1, 'a', []State{2})
	n0.Add(2, 'a', []State{2})
	n0.Add(3, 'b', []State{4})
	n0.Add(4, 'b', []State{4})

	// (a|b)*abb
	n1 := NewNFA(0, []State{10})
	n1.Add(0, E, []State{1, 7})
	n1.Add(1, E, []State{2, 4})
	n1.Add(2, 'a', []State{3})
	n1.Add(3, E, []State{6})
	n1.Add(4, 'b', []State{5})
	n1.Add(5, E, []State{6})
	n1.Add(6, E, []State{1, 7})
	n1.Add(7, 'a', []State{8})
	n1.Add(8, 'b', []State{9})
	n1.Add(9, 'b', []State{10})

	// ab+|ba+
	n2 := NewNFA(0, []State{2, 4})
	n2.Add(0, 'a', []State{1})
	n2.Add(1, 'b', []State{2})
	n2.Add(2, 'b', []State{2})
	n2.Add(0, 'b', []State{3})
	n2.Add(3, 'a', []State{4})
	n2.Add(4, 'a', []State{4})

	// (ab)+
	n3 := NewNFA(0, []State{2})
	n3.Add(0, 'a', []State{1})
	n3.Add(1, 'b', []State{2})
	n3.Add(2, 'a', []State{1})

	// (ab+|ba+)*
	n4 := NewNFA(0, []State{1})
	n4.Add(0, E, []State{1})
	n4.Add(0, E, []State{2})
	n4.Add(2, 'a', []State{3})
	n4.Add(2, 'b', []State{4})
	n4.Add(3, 'b', []State{5})
	n4.Add(4, 'a', []State{6})
	n4.Add(5, 'b', []State{5})
	n4.Add(6, 'a', []State{6})
	n4.Add(5, E, []State{1})
	n4.Add(5, E, []State{2})
	n4.Add(6, E, []State{1})
	n4.Add(6, E, []State{2})

	// ab+|ba+|(ab)+
	n5 := NewNFA(0, []State{1})
	n5.Add(0, E, []State{2})
	n5.Add(0, E, []State{7})
	n5.Add(2, 'a', []State{3})
	n5.Add(2, 'b', []State{4})
	n5.Add(3, 'b', []State{5})
	n5.Add(4, 'a', []State{6})
	n5.Add(5, 'b', []State{5})
	n5.Add(5, E, []State{1})
	n5.Add(6, 'a', []State{6})
	n5.Add(6, E, []State{1})
	n5.Add(7, 'a', []State{8})
	n5.Add(8, 'b', []State{9})
	n5.Add(9, 'a', []State{8})
	n5.Add(9, E, []State{1})

	// (ab+|ba+)(ab)+
	n6 := NewNFA(0, []State{6})
	n6.Add(0, 'a', []State{1})
	n6.Add(0, 'b', []State{2})
	n6.Add(1, 'b', []State{3})
	n6.Add(2, 'a', []State{4})
	n6.Add(3, 'a', []State{5})
	n6.Add(3, 'b', []State{3})
	n6.Add(4, 'a', []State{4})
	n6.Add(4, 'a', []State{5})
	n6.Add(5, 'b', []State{6})
	n6.Add(6, 'a', []State{5})

	return []*NFA{n0, n1, n2, n3, n4, n5, n6}
}

func TestNewNFA(t *testing.T) {
	tests := []struct {
		name  string
		start State
		final []State
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := NewNFA(tc.start, tc.final)

			assert.NotNil(t, nfa)
			assert.Equal(t, tc.start, nfa.Start)
			assert.True(t, nfa.Final.Contains(tc.final...))
		})
	}
}

func Test_newNFA(t *testing.T) {
	tests := []struct {
		name  string
		start State
		final States
	}{
		{
			name:  "OK",
			start: 0,
			final: NewStates(2, 4),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := newNFA(tc.start, tc.final)

			assert.NotNil(t, nfa)
			assert.Equal(t, tc.start, nfa.Start)
			assert.True(t, nfa.Final.Equal(tc.final))
		})
	}
}

func TestNFA_String(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		n              *NFA
		expectedString string
	}{
		{
			name: "OK",
			n:    nfas[1],
			expectedString: `Start state: 0
Final states: 10
Transitions:
  (0, ε) --> {1, 7}
  (1, ε) --> {2, 4}
  (2, a) --> {3}
  (3, ε) --> {6}
  (4, b) --> {5}
  (5, ε) --> {6}
  (6, ε) --> {1, 7}
  (7, a) --> {8}
  (8, b) --> {9}
  (9, b) --> {10}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.n.String())
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

func TestNFA_Clone(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name string
		n    *NFA
	}{
		{
			name: "OK",
			n:    nfas[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.n.Clone()
			assert.True(t, nfa.Equal(tc.n))
		})
	}
}

func TestNFA_Add(t *testing.T) {
	nfa := NewNFA(0, []State{1, 2, 3, 4})

	tests := []struct {
		name string
		n    *NFA
		s    State
		a    Symbol
		next []State
	}{
		{
			name: "NewState",
			n:    nfa,
			s:    State(0),
			a:    'a',
			next: []State{1, 2},
		},
		{
			name: "ExistingState",
			n:    nfa,
			s:    State(0),
			a:    'b',
			next: []State{3, 4},
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
		expectedStates []State
	}{
		{
			name:           "First",
			n:              nfas[0],
			s:              State(0),
			a:              E,
			expectedStates: []State{1, 3},
		},
		{
			name:           "Second",
			n:              nfas[1],
			s:              State(1),
			a:              E,
			expectedStates: []State{2, 4},
		},
		{
			name:           "Invalid",
			n:              nfas[0],
			s:              State(0),
			a:              'c',
			expectedStates: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states := tc.n.Next(tc.s, tc.a)
			assert.Equal(t, tc.expectedStates, states)
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
			s:              toString("aaaa"),
			expectedResult: true,
		},
		{
			name:           "NotAccepted",
			n:              nfa,
			s:              toString("abbb"),
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

func TestNFA_States(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		n              *NFA
		expectedStates []State
	}{
		{
			name:           "Empty",
			n:              NewNFA(0, []State{1, 2}),
			expectedStates: []State{0, 1, 2},
		},
		{
			name:           "First",
			n:              nfas[0],
			expectedStates: []State{0, 1, 2, 3, 4},
		},
		{
			name:           "Second",
			n:              nfas[1],
			expectedStates: []State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.n.States())
		})
	}
}

func TestNFA_Symbols(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name            string
		n               *NFA
		expectedSymbols []Symbol
	}{
		{
			name:            "First",
			n:               nfas[0],
			expectedSymbols: []Symbol{'a', 'b'},
		},
		{
			name:            "Second",
			n:               nfas[1],
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
	nfas := getTestNFAs()

	tests := []struct {
		name                string
		n                   *NFA
		expectedTransitions []*Transition[[]State]
	}{
		{
			name: "First",
			n:    nfas[0],
			expectedTransitions: []*Transition[[]State]{
				{0, E, []State{1, 3}},
				{1, 'a', []State{2}},
				{2, 'a', []State{2}},
				{3, 'b', []State{4}},
				{4, 'b', []State{4}},
			},
		},
		{
			name: "Second",
			n:    nfas[1],
			expectedTransitions: []*Transition[[]State]{
				{0, E, []State{1, 7}},
				{1, E, []State{2, 4}},
				{2, 'a', []State{3}},
				{3, E, []State{6}},
				{4, 'b', []State{5}},
				{5, E, []State{6}},
				{6, E, []State{1, 7}},
				{7, 'a', []State{8}},
				{8, 'b', []State{9}},
				{9, 'b', []State{10}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			transitions := generic.Collect1(tc.n.Transitions())
			assert.Equal(t, tc.expectedTransitions, transitions)
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

func TestNFA_Isomorphic(t *testing.T) {
	// aa*|bb*
	n0 := NewNFA(0, []State{2, 4})
	n0.Add(0, E, []State{1, 3})
	n0.Add(1, 'a', []State{2})
	n0.Add(2, 'a', []State{2})
	n0.Add(3, 'b', []State{4})
	n0.Add(4, 'b', []State{4})

	n1 := NewNFA(0, []State{2})

	n2 := NewNFA(0, []State{3, 6})
	n2.Add(0, E, []State{1, 4})
	n2.Add(1, 'a', []State{2})
	n2.Add(2, 'a', []State{3})
	n2.Add(3, 'a', []State{3})
	n2.Add(4, 'b', []State{5})
	n2.Add(5, 'b', []State{6})
	n2.Add(6, 'b', []State{6})

	n3 := NewNFA(0, []State{2, 4})
	n3.Add(0, E, []State{1, 3})
	n3.Add(1, 'a', []State{2})
	n3.Add(2, 'c', []State{2})
	n3.Add(3, 'b', []State{4})
	n3.Add(4, 'c', []State{4})

	n4 := NewNFA(0, []State{2, 4})
	n4.Add(0, E, []State{1, 3})
	n4.Add(1, 'a', []State{2})
	n4.Add(2, 'a', []State{2})
	n4.Add(3, 'b', []State{4})

	n5 := NewNFA(4, []State{0, 1})
	n5.Add(4, E, []State{2, 3})
	n5.Add(2, 'a', []State{0})
	n5.Add(0, 'a', []State{0})
	n5.Add(3, 'b', []State{1})
	n5.Add(1, 'b', []State{1})

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
			name: "Empty",
			n:    NewNFA(0, []State{1}),
			expectedDOT: `digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1", shape=doublecircle];

  start -> 0 [];
}`,
		},
		{
			name: "First",
			n:    nfas[0],
			expectedDOT: `digraph "NFA" {
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
			expectedDOT: `digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1"];
  2 [label="2"];
  3 [label="3"];
  4 [label="4"];
  5 [label="5"];
  6 [label="6"];
  7 [label="7"];
  8 [label="8"];
  9 [label="9"];
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
		{
			name: "Third",
			n:    nfas[3],
			expectedDOT: `digraph "NFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1"];
  2 [label="2", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 1 [label="a"];
}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.n.DOT())
		})
	}
}
