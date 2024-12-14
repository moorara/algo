package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestDFAs() []*DFA {
	// (a|b)*abb
	d0 := NewDFA(0, States{3})
	d0.Add(0, 'a', 1)
	d0.Add(0, 'b', 0)
	d0.Add(1, 'a', 1)
	d0.Add(1, 'b', 2)
	d0.Add(2, 'a', 1)
	d0.Add(2, 'b', 3)
	d0.Add(3, 'a', 1)
	d0.Add(3, 'b', 0)

	// (a|b)*abb
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

	// ab+|ba+
	d2 := NewDFA(0, States{2, 4})
	d2.Add(0, 'a', 1)
	d2.Add(1, 'b', 2)
	d2.Add(2, 'b', 2)
	d2.Add(0, 'b', 3)
	d2.Add(3, 'a', 4)
	d2.Add(4, 'a', 4)

	// (ab)+
	d3 := NewDFA(0, States{2})
	d3.Add(0, 'a', 1)
	d3.Add(1, 'b', 2)
	d3.Add(2, 'a', 1)

	// ab(a|b)*
	d4 := NewDFA(0, States{2})
	d4.Add(0, 'a', 1)
	d4.Add(0, 'b', 3)
	d4.Add(1, 'a', 4)
	d4.Add(1, 'b', 2)
	d4.Add(2, 'a', 2)
	d4.Add(2, 'b', 2)
	d4.Add(3, 'a', 3)
	d4.Add(3, 'b', 3)
	d4.Add(4, 'a', 4)
	d4.Add(4, 'b', 4)

	// ab(a|b)*
	d5 := NewDFA(0, States{2})
	d5.Add(0, 'a', 1)
	d5.Add(1, 'b', 2)
	d5.Add(2, 'a', 2)
	d5.Add(2, 'b', 2)

	return []*DFA{d0, d1, d2, d3, d4, d5}
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
		{
			name:          "Invalid",
			d:             dfas[0],
			s:             State(0),
			a:             'c',
			expectedState: State(-1),
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
			d:           dfas[2],
			expectedNFA: nfas[2],
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
			expectedDFA: dfas[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.Minimize()
			assert.True(t, dfa.Isomorphic(tc.expectedDFA))
		})
	}
}

func TestDFA_WithoutDeadStates(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name:        "OK",
			d:           dfas[4],
			expectedDFA: dfas[5],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.WithoutDeadStates()
			assert.True(t, dfa.Equals(tc.expectedDFA))
		})
	}
}

func TestDFA_Equals(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		d              *DFA
		rhs            *DFA
		expectedEquals bool
	}{
		{
			name:           "Equal",
			d:              dfas[0],
			rhs:            dfas[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			d:              dfas[0],
			rhs:            dfas[1],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.d.Equals(tc.rhs))
		})
	}
}

func TestDFA_Isomorphic(t *testing.T) {
	// aa*|bb*
	d0 := NewDFA(0, States{1, 2})
	d0.Add(0, 'a', 1)
	d0.Add(1, 'a', 1)
	d0.Add(0, 'b', 2)
	d0.Add(2, 'b', 2)

	d1 := NewDFA(0, States{1})

	d2 := NewDFA(0, States{2, 4})
	d2.Add(0, 'a', 1)
	d2.Add(1, 'a', 2)
	d2.Add(2, 'a', 2)
	d2.Add(0, 'b', 3)
	d2.Add(3, 'b', 4)
	d2.Add(4, 'b', 4)

	d3 := NewDFA(0, States{1, 2})
	d3.Add(0, 'a', 1)
	d3.Add(1, 'c', 1)
	d3.Add(0, 'b', 2)
	d3.Add(2, 'c', 2)

	d4 := NewDFA(0, States{1, 2})
	d4.Add(0, 'a', 1)
	d4.Add(1, 'a', 1)
	d4.Add(0, 'b', 2)

	d5 := NewDFA(2, States{0, 1})
	d5.Add(2, 'a', 0)
	d5.Add(0, 'a', 0)
	d5.Add(2, 'b', 1)
	d5.Add(1, 'b', 1)

	tests := []struct {
		name               string
		d                  *DFA
		rhs                *DFA
		expectedIsomorphic bool
	}{
		{
			name:               "FinalStatesNotEqualSize",
			d:                  d0,
			rhs:                d1,
			expectedIsomorphic: false,
		},
		{
			name:               "StatesNotEqualSize",
			d:                  d0,
			rhs:                d2,
			expectedIsomorphic: false,
		},
		{
			name:               "SymbolsNotEqual",
			d:                  d0,
			rhs:                d3,
			expectedIsomorphic: false,
		},
		{
			name:               "DegreesNotEqual",
			d:                  d0,
			rhs:                d4,
			expectedIsomorphic: false,
		},
		{
			name:               "Equal",
			d:                  d0,
			rhs:                d0,
			expectedIsomorphic: true,
		},
		{
			name:               "Isomorphic",
			d:                  d0,
			rhs:                d5,
			expectedIsomorphic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsomorphic, tc.d.Isomorphic(tc.rhs))
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
			name:             "Empty",
			d:                NewDFA(0, States{1}),
			expectedGraphviz: dfaEmpty,
		},
		{
			name:             "First",
			d:                dfas[0],
			expectedGraphviz: dfa01,
		},
		{
			name:             "Second",
			d:                dfas[1],
			expectedGraphviz: dfa02,
		},
		{
			name:             "Third",
			d:                dfas[3],
			expectedGraphviz: dfa03,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedGraphviz, tc.d.Graphviz())
		})
	}
}

var dfaEmpty = `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=doublecircle];

  start -> 0 [];
}`

var dfa01 = `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=circle];
  3 [label="3", shape=doublecircle];

  start -> 0 [];
  0 -> 0 [label="b"];
  0 -> 1 [label="a"];
  1 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 1 [label="a"];
  2 -> 3 [label="b"];
  3 -> 0 [label="b"];
  3 -> 1 [label="a"];
}`

var dfa02 = `digraph "DFA" {
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
}`

var dfa03 = `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [];

  start [style=invis];
  0 [label="0", shape=circle];
  1 [label="1", shape=circle];
  2 [label="2", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="a"];
  1 -> 2 [label="b"];
  2 -> 1 [label="a"];
}`
