package automata

import (
	"testing"

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

	return []*DFA{d0, d1, d2, d3, d4, d5}
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
