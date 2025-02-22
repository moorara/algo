package automata

import (
	"testing"

	"github.com/moorara/algo/generic"
	"github.com/stretchr/testify/assert"
)

func getTestDFAs() []*DFA {
	// (a|b)*abb
	d0 := NewDFA(0, []State{3})
	d0.Add(0, 'a', 1)
	d0.Add(0, 'b', 0)
	d0.Add(1, 'a', 1)
	d0.Add(1, 'b', 2)
	d0.Add(2, 'a', 1)
	d0.Add(2, 'b', 3)
	d0.Add(3, 'a', 1)
	d0.Add(3, 'b', 0)

	// (a|b)*abb
	d1 := NewDFA(0, []State{4})
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
	d2 := NewDFA(0, []State{2, 4})
	d2.Add(0, 'a', 1)
	d2.Add(1, 'b', 2)
	d2.Add(2, 'b', 2)
	d2.Add(0, 'b', 3)
	d2.Add(3, 'a', 4)
	d2.Add(4, 'a', 4)

	// (ab)+
	d3 := NewDFA(0, []State{2})
	d3.Add(0, 'a', 1)
	d3.Add(1, 'b', 2)
	d3.Add(2, 'a', 1)

	// ab(a|b)*
	d4 := NewDFA(0, []State{2})
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
	d5 := NewDFA(0, []State{2})
	d5.Add(0, 'a', 1)
	d5.Add(1, 'b', 2)
	d5.Add(2, 'a', 2)
	d5.Add(2, 'b', 2)

	// (0|1)+.(0|1)+
	d6 := NewDFA(0, []State{3, 4})
	d6.Add(0, '0', 3)
	d6.Add(0, '1', 3)
	d6.Add(3, '0', 3)
	d6.Add(3, '1', 3)
	d6.Add(3, '.', 1)
	d6.Add(1, '0', 4)
	d6.Add(1, '1', 4)
	d6.Add(4, '0', 4)
	d6.Add(4, '1', 4)

	// (0|1)+.(0|1)+
	d7 := NewDFA(0, []State{1, 3})
	d7.Add(0, '0', 1)
	d7.Add(0, '1', 1)
	d7.Add(1, '0', 1)
	d7.Add(1, '1', 1)
	d7.Add(1, '.', 2)
	d7.Add(2, '0', 3)
	d7.Add(2, '1', 3)
	d7.Add(3, '0', 3)
	d7.Add(3, '1', 3)

	return []*DFA{d0, d1, d2, d3, d4, d5, d6, d7}
}

func TestNewDFA(t *testing.T) {
	tests := []struct {
		name  string
		start State
		final []State
	}{
		{
			name:  "OK",
			start: 0,
			final: []State{3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := NewDFA(tc.start, tc.final)

			assert.NotNil(t, dfa)
			assert.Equal(t, tc.start, dfa.Start)
			assert.True(t, dfa.Final.Contains(tc.final...))
		})
	}
}

func Test_newDFA(t *testing.T) {
	tests := []struct {
		name  string
		start State
		final States
	}{
		{
			name:  "OK",
			start: 0,
			final: NewStates(3),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := newDFA(tc.start, tc.final)

			assert.NotNil(t, dfa)
			assert.Equal(t, tc.start, dfa.Start)
			assert.True(t, dfa.Final.Equal(tc.final))
		})
	}
}

func TestDFA_String(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		d              *DFA
		expectedString string
	}{
		{
			name: "OK",
			d:    dfas[1],
			expectedString: `Start state: 0
Final states: 4
Transitions:
  (0, a) --> 1
  (0, b) --> 2
  (1, a) --> 1
  (1, b) --> 3
  (2, a) --> 1
  (2, b) --> 2
  (3, a) --> 1
  (3, b) --> 4
  (4, a) --> 1
  (4, b) --> 2
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.d.String())
		})
	}
}

func TestDFA_Equal(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name          string
		d             *DFA
		rhs           *DFA
		expectedEqual bool
	}{
		{
			name:          "Equal",
			d:             dfas[0],
			rhs:           dfas[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			d:             dfas[0],
			rhs:           dfas[1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.d.Equal(tc.rhs))
		})
	}
}

func TestDFA_Clone(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name string
		d    *DFA
	}{
		{
			name: "OK",
			d:    dfas[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.Clone()
			assert.True(t, dfa.Equal(tc.d))
		})
	}
}

func TestDFA_Add(t *testing.T) {
	dfa := NewDFA(0, []State{1, 2})

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
			s:              toString("aabbababb"),
			expectedResult: true,
		},
		{
			name:           "NotAccepted",
			d:              dfa,
			s:              toString("aabab"),
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

func TestDFA_States(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		d              *DFA
		expectedStates []State
	}{
		{
			name:           "Empty",
			d:              NewDFA(0, []State{1}),
			expectedStates: []State{0, 1},
		},
		{
			name:           "First",
			d:              dfas[0],
			expectedStates: []State{0, 1, 2, 3},
		},
		{
			name:           "Second",
			d:              dfas[1],
			expectedStates: []State{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.d.States())
		})
	}
}

func TestDFA_Symbols(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name            string
		d               *DFA
		expectedSymbols []Symbol
	}{
		{
			name:            "First",
			d:               dfas[0],
			expectedSymbols: []Symbol{'a', 'b'},
		},
		{
			name:            "Second",
			d:               dfas[1],
			expectedSymbols: []Symbol{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.d.Symbols())
		})
	}
}

func TestDFA_Transitions(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name                string
		d                   *DFA
		expectedTransitions []*Transition[State]
	}{
		{
			name: "First",
			d:    dfas[0],
			expectedTransitions: []*Transition[State]{
				{0, 'a', 1},
				{0, 'b', 0},
				{1, 'a', 1},
				{1, 'b', 2},
				{2, 'a', 1},
				{2, 'b', 3},
				{3, 'a', 1},
				{3, 'b', 0},
			},
		},
		{
			name: "Second",
			d:    dfas[1],
			expectedTransitions: []*Transition[State]{
				{0, 'a', 1},
				{0, 'b', 2},
				{1, 'a', 1},
				{1, 'b', 3},
				{2, 'a', 1},
				{2, 'b', 2},
				{3, 'a', 1},
				{3, 'b', 4},
				{4, 'a', 1},
				{4, 'b', 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			transitions := generic.Collect1(tc.d.Transitions())
			assert.Equal(t, tc.expectedTransitions, transitions)
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
			assert.True(t, nfa.Equal(tc.expectedNFA))
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

func TestDFA_EliminateDeadStates(t *testing.T) {
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
			dfa := tc.d.EliminateDeadStates()
			assert.True(t, dfa.Equal(tc.expectedDFA))
		})
	}
}

func TestDFA_ReindexStates(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name:        "OK",
			d:           dfas[6],
			expectedDFA: dfas[7],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.ReindexStates()
			assert.True(t, dfa.Equal(tc.expectedDFA))
		})
	}
}

func TestDFA_Isomorphic(t *testing.T) {
	// aa*|bb*
	d0 := NewDFA(0, []State{1, 2})
	d0.Add(0, 'a', 1)
	d0.Add(1, 'a', 1)
	d0.Add(0, 'b', 2)
	d0.Add(2, 'b', 2)

	d1 := NewDFA(0, []State{1})

	d2 := NewDFA(0, []State{2, 4})
	d2.Add(0, 'a', 1)
	d2.Add(1, 'a', 2)
	d2.Add(2, 'a', 2)
	d2.Add(0, 'b', 3)
	d2.Add(3, 'b', 4)
	d2.Add(4, 'b', 4)

	d3 := NewDFA(0, []State{1, 2})
	d3.Add(0, 'a', 1)
	d3.Add(1, 'c', 1)
	d3.Add(0, 'b', 2)
	d3.Add(2, 'c', 2)

	d4 := NewDFA(0, []State{1, 2})
	d4.Add(0, 'a', 1)
	d4.Add(1, 'a', 1)
	d4.Add(0, 'b', 2)

	d5 := NewDFA(2, []State{0, 1})
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

func TestDFA_DOT(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name        string
		d           *DFA
		expectedDOT string
	}{
		{
			name: "Empty",
			d:    NewDFA(0, []State{1}),
			expectedDOT: `digraph "DFA" {
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
			d:    dfas[0],
			expectedDOT: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1"];
  2 [label="2"];
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
}`,
		},
		{
			name: "Second",
			d:    dfas[1],
			expectedDOT: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1"];
  2 [label="2"];
  3 [label="3"];
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
		{
			name: "Third",
			d:    dfas[3],
			expectedDOT: `digraph "DFA" {
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
			assert.Equal(t, tc.expectedDOT, tc.d.DOT())
		})
	}
}

func TestCombineDFA(t *testing.T) {
	dfas := getTestDFAs()

	dfa := NewDFA(0, []State{3, 4, 5, 6, 7, 8})
	dfa.Add(0, 'a', 1)
	dfa.Add(0, 'b', 2)
	dfa.Add(1, 'b', 3)
	dfa.Add(2, 'a', 4)
	dfa.Add(3, 'a', 5)
	dfa.Add(3, 'b', 6)
	dfa.Add(4, 'a', 4)
	dfa.Add(5, 'a', 7)
	dfa.Add(5, 'b', 8)
	dfa.Add(6, 'a', 7)
	dfa.Add(6, 'b', 6)
	dfa.Add(7, 'a', 7)
	dfa.Add(7, 'b', 7)
	dfa.Add(8, 'a', 5)
	dfa.Add(8, 'b', 7)

	tests := []struct {
		name             string
		ds               []*DFA
		expectedDFA      *DFA
		expectedFinalMap [][]State
	}{
		{
			name: "OK",
			ds: []*DFA{
				dfas[2],
				dfas[3],
				dfas[5],
			},
			expectedDFA: dfa,
			expectedFinalMap: [][]State{
				{3, 4, 6},
				{3, 8},
				{3, 5, 6, 7, 8},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa, finalMap := CombineDFA(tc.ds...)

			assert.True(t, dfa.Equal(tc.expectedDFA))
			assert.Equal(t, tc.expectedFinalMap, finalMap)
		})
	}
}
