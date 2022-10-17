package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestNFAs() []*NFA {
	t1 := NewNTrans()
	t1.Add(0, E, States{1, 3})
	t1.Add(1, 'a', States{2})
	t1.Add(2, 'a', States{2})
	t1.Add(3, 'b', States{4})
	t1.Add(4, 'b', States{4})

	t2 := NewNTrans()
	t2.Add(0, E, States{1, 7})
	t2.Add(1, E, States{2, 4})
	t2.Add(2, 'a', States{3})
	t2.Add(3, E, States{6})
	t2.Add(4, 'b', States{5})
	t2.Add(5, E, States{6})
	t2.Add(6, E, States{1, 7})
	t2.Add(7, 'a', States{8})
	t2.Add(8, 'b', States{9})
	t2.Add(9, 'b', States{10})

	t3 := NewNTrans()
	t3.Add(0, E, States{1, 3})
	t3.Add(1, 'a', States{2})
	t3.Add(2, 'a', States{2})
	t3.Add(3, 'b', States{4})
	t3.Add(4, 'b', States{4})
	t3.Add(2, E, States{5})
	t3.Add(4, E, States{5})
	t3.Add(5, E, States{6, 12})
	t3.Add(6, E, States{7, 9})
	t3.Add(7, 'a', States{8})
	t3.Add(8, E, States{11})
	t3.Add(9, 'b', States{10})
	t3.Add(10, E, States{11})
	t3.Add(11, E, States{6, 12})
	t3.Add(12, 'a', States{13})
	t3.Add(13, 'b', States{14})
	t3.Add(14, 'b', States{15})

	return []*NFA{
		{
			trans: t1,
			start: State(0),
			final: States{2, 4},
		},
		{
			trans: t2,
			start: State(0),
			final: States{10},
		},
		{
			trans: t3,
			start: State(0),
			final: States{15},
		},
	}
}

func TestNewNTrans(t *testing.T) {
	ntrans := NewNTrans()
	assert.NotNil(t, ntrans)
}

func TestNTrans_Add(t *testing.T) {
	ntrans := NewNTrans()

	tests := []struct {
		name   string
		ntrans *NTrans
		s      State
		a      Symbol
		next   States
	}{
		{
			name:   "NewState",
			ntrans: ntrans,
			s:      State(0),
			a:      'a',
			next:   States{1, 2},
		},
		{
			name:   "ExistingState",
			ntrans: ntrans,
			s:      State(0),
			a:      'b',
			next:   States{3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.ntrans.Add(tc.s, tc.a, tc.next)
		})
	}
}

func TestNTrans_Next(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		ntrans         *NTrans
		s              State
		a              Symbol
		expectedStates States
	}{
		{
			name:           "First",
			ntrans:         nfas[0].trans,
			s:              State(0),
			a:              E,
			expectedStates: States{1, 3},
		},
		{
			name:           "Second",
			ntrans:         nfas[1].trans,
			s:              State(1),
			a:              E,
			expectedStates: States{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states := tc.ntrans.Next(tc.s, tc.a)
			assert.Equal(t, tc.expectedStates, states)
		})
	}
}

func TestNTrans_States(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name           string
		ntrans         *NTrans
		expectedStates States
	}{
		{
			name:           "First",
			ntrans:         nfas[0].trans,
			expectedStates: States{0, 1, 2, 3, 4},
		},
		{
			name:           "Second",
			ntrans:         nfas[1].trans,
			expectedStates: States{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.ntrans.States())
		})
	}
}

func TestNTrans_Symbols(t *testing.T) {
	nfas := getTestNFAs()

	tests := []struct {
		name            string
		ntrans          *NTrans
		expectedSymbols Symbols
	}{
		{
			name:            "First",
			ntrans:          nfas[0].trans,
			expectedSymbols: Symbols{'a', 'b'},
		},
		{
			name:            "Second",
			ntrans:          nfas[1].trans,
			expectedSymbols: Symbols{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.ntrans.Symbols())
		})
	}
}

func TestNewNFA(t *testing.T) {
	n := getTestNFAs()[0]

	nfa := NewNFA(n.trans, n.start, n.final)
	assert.NotNil(t, nfa)
}

func TestNFA_UpdateFinal(t *testing.T) {
	nfa := getTestNFAs()[0]

	tests := []struct {
		name  string
		n     *NFA
		final States
	}{
		{
			name:  "OK",
			n:     nfa,
			final: States{4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.n.UpdateFinal(tc.final)
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
		extraTrans     []edge
		newFinal       States
		expectedStates States
		expectedNFA    *NFA
	}{
		{
			name: "OK",
			n:    nfas[0],
			nfa:  nfas[1],
			extraTrans: []edge{
				{2, E, States{5}},
				{4, E, States{5}},
			},
			newFinal:       States{15},
			expectedStates: States{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expectedNFA:    nfas[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			states := tc.n.Join(tc.nfa)
			for _, e := range tc.extraTrans {
				tc.n.trans.Add(e.s, e.a, e.next)
			}
			tc.n.UpdateFinal(tc.newFinal)

			assert.Equal(t, tc.expectedStates, states)
			// This is a trick to avoid comparing the symbol tables with their internal structures.
			assert.Equal(t, tc.expectedNFA.Graphviz(), tc.n.Graphviz())
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

			// This is a trick to avoid comparing the symbol tables with their internal structures.
			assert.Equal(t, tc.expectedDFA.Graphviz(), dfa.Graphviz())
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