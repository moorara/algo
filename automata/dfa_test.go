package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestDFAs() []*DFA {
	t1 := NewDTrans()
	t1.Add(0, 'a', 1)
	t1.Add(0, 'b', 0)
	t1.Add(1, 'a', 1)
	t1.Add(1, 'b', 2)
	t1.Add(2, 'a', 1)
	t1.Add(2, 'b', 3)
	t1.Add(3, 'a', 1)
	t1.Add(3, 'b', 0)

	t2 := NewDTrans()
	t2.Add(0, 'a', 1)
	t2.Add(0, 'b', 2)
	t2.Add(1, 'a', 1)
	t2.Add(1, 'b', 3)
	t2.Add(2, 'a', 1)
	t2.Add(2, 'b', 2)
	t2.Add(3, 'a', 1)
	t2.Add(3, 'b', 4)
	t2.Add(4, 'a', 1)
	t2.Add(4, 'b', 2)

	return []*DFA{
		{
			trans: t1,
			start: State(0),
			final: States{3},
		},
		{
			trans: t2,
			start: State(0),
			final: States{4},
		},
	}
}

func TestNewDTrans(t *testing.T) {
	dtrans := NewDTrans()
	assert.NotNil(t, dtrans)
}

func TestDTrans_Add(t *testing.T) {
	dtrans := NewDTrans()

	tests := []struct {
		name   string
		dtrans *DTrans
		s      State
		a      Symbol
		next   State
	}{
		{
			name:   "NewState",
			dtrans: dtrans,
			s:      State(0),
			a:      'a',
			next:   State(1),
		},
		{
			name:   "ExistingState",
			dtrans: dtrans,
			s:      State(0),
			a:      'b',
			next:   State(2),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.dtrans.Add(tc.s, tc.a, tc.next)
		})
	}
}

func TestDTrans_Next(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name          string
		dtrans        *DTrans
		s             State
		a             Symbol
		expectedState State
	}{
		{
			name:          "First",
			dtrans:        dfas[0].trans,
			s:             State(0),
			a:             'a',
			expectedState: State(1),
		},
		{
			name:          "Second",
			dtrans:        dfas[1].trans,
			s:             State(0),
			a:             'b',
			expectedState: State(2),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.dtrans.Next(tc.s, tc.a)
			assert.Equal(t, tc.expectedState, state)
		})
	}
}

func TestDTrans_States(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name           string
		dtrans         *DTrans
		expectedStates States
	}{
		{
			name:           "First",
			dtrans:         dfas[0].trans,
			expectedStates: States{0, 1, 2, 3},
		},
		{
			name:           "Second",
			dtrans:         dfas[1].trans,
			expectedStates: States{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.dtrans.States())
		})
	}
}

func TestDTrans_Symbols(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name            string
		dtrans          *DTrans
		expectedSymbols Symbols
	}{
		{
			name:            "First",
			dtrans:          dfas[0].trans,
			expectedSymbols: Symbols{'a', 'b'},
		},
		{
			name:            "Second",
			dtrans:          dfas[1].trans,
			expectedSymbols: Symbols{'a', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.dtrans.Symbols())
		})
	}
}

func TestNewDFA(t *testing.T) {
	d := getTestDFAs()[0]

	dfa := NewDFA(d.trans, d.start, d.final)
	assert.NotNil(t, dfa)
}

func TestDFA_Accept(t *testing.T) {
	dfa := getTestDFAs()[0]

	tests := []struct {
		name           string
		dfa            *DFA
		s              String
		expectedResult bool
	}{
		{
			name:           "Accepted",
			dfa:            dfa,
			s:              ToString("aabbababb"),
			expectedResult: true,
		},
		{
			name:           "NotAccepted",
			dfa:            dfa,
			s:              ToString("aabab"),
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.dfa.Accept(tc.s)
			assert.Equal(t, tc.expectedResult, b)
		})
	}
}

func TestDFA_Graphviz(t *testing.T) {
	dfas := getTestDFAs()

	tests := []struct {
		name             string
		dfa              *DFA
		expectedGraphviz string
	}{
		{
			name: "First",
			dfa:  dfas[0],
			expectedGraphviz: `strict digraph "DFA" {
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
			dfa:  dfas[1],
			expectedGraphviz: `strict digraph "DFA" {
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
			assert.Equal(t, tc.expectedGraphviz, tc.dfa.Graphviz())
		})
	}
}
