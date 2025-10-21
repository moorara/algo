package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/range/disc"
)

var testNFA = []*NFA{
	// (a+|b+)
	{
		start: 0,
		final: NewStates(2, 4),
		classes: disc.NewRangeMap(eqClassID, classesOpts, []disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(1, 3)).
			Add(1, 1, NewStates(2)).
			Add(2, 1, NewStates(2)).
			Add(3, 2, NewStates(4)).
			Add(4, 2, NewStates(4)),
	},
	// ([A-Za-z_][0-9A-Za-z_]*)|[0-9]+|(0x[0-9A-Fa-f]+)|[ \t\n]+|[+\-*/=]
	{
		start: 0,
		final: NewStates(1, 2, 5),
		classes: disc.NewRangeMap(eqClassID, classesOpts, []disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'F'}, Value: 2},
			{Range: disc.Range[Symbol]{Lo: 'G', Hi: 'W'}, Value: 3},
			{Range: disc.Range[Symbol]{Lo: 'X', Hi: 'X'}, Value: 4},
			{Range: disc.Range[Symbol]{Lo: 'Y', Hi: 'Z'}, Value: 3},
			{Range: disc.Range[Symbol]{Lo: '_', Hi: '_'}, Value: 3},
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'f'}, Value: 2},
			{Range: disc.Range[Symbol]{Lo: 'g', Hi: 'w'}, Value: 3},
			{Range: disc.Range[Symbol]{Lo: 'x', Hi: 'x'}, Value: 4},
			{Range: disc.Range[Symbol]{Lo: 'y', Hi: 'z'}, Value: 3},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(2, 3)).
			Add(0, 1, NewStates(2)).
			Add(0, 2, NewStates(1)).
			Add(0, 3, NewStates(1)).
			Add(0, 4, NewStates(1)).
			Add(1, 0, NewStates(1)).
			Add(1, 1, NewStates(1)).
			Add(1, 2, NewStates(1)).
			Add(1, 3, NewStates(1)).
			Add(1, 4, NewStates(1)).
			Add(2, 0, NewStates(2)).
			Add(2, 1, NewStates(2)).
			Add(3, 0, NewStates(2)).
			Add(3, 1, NewStates(2)).
			Add(3, 4, NewStates(4)).
			Add(4, 0, NewStates(5)).
			Add(4, 1, NewStates(5)).
			Add(4, 2, NewStates(5)).
			Add(5, 0, NewStates(5)).
			Add(5, 1, NewStates(5)).
			Add(5, 2, NewStates(5)),
	},
}

func TestNFABuilder(t *testing.T) {
	type transition struct {
		s          State
		start, end Symbol
		next       []State
	}

	tests := []struct {
		name        string
		start       State
		final       []State
		trans       []transition
		expectedNFA *NFA
	}{
		{
			name:  "Simple",
			start: 0,
			final: []State{2, 4},
			trans: []transition{
				{s: 0, start: E, end: E, next: []State{1, 3}},
				{s: 1, start: 'a', end: 'a', next: []State{2}},
				{s: 2, start: 'a', end: 'a', next: []State{2}},
				{s: 3, start: 'b', end: 'b', next: []State{4}},
				{s: 4, start: 'b', end: 'b', next: []State{4}},
			},
			expectedNFA: testNFA[0],
		},
		{
			name:  "ID_NUM_WS_OP",
			start: 0,
			final: []State{1, 2, 5},
			trans: []transition{
				{s: 0, start: '0', end: '0', next: []State{3}},
				{s: 0, start: '0', end: '9', next: []State{2}},
				{s: 0, start: 'A', end: 'Z', next: []State{1}},
				{s: 0, start: 'a', end: 'z', next: []State{1}},
				{s: 0, start: '_', end: '_', next: []State{1}},
				{s: 1, start: '0', end: '9', next: []State{1}},
				{s: 1, start: 'A', end: 'Z', next: []State{1}},
				{s: 1, start: 'a', end: 'z', next: []State{1}},
				{s: 1, start: '_', end: '_', next: []State{1}},
				{s: 2, start: '0', end: '9', next: []State{2}},
				{s: 3, start: '0', end: '9', next: []State{2}},
				{s: 3, start: 'X', end: 'X', next: []State{4}},
				{s: 3, start: 'x', end: 'x', next: []State{4}},
				{s: 4, start: '0', end: '9', next: []State{5}},
				{s: 4, start: 'A', end: 'F', next: []State{5}},
				{s: 4, start: 'a', end: 'f', next: []State{5}},
				{s: 5, start: '0', end: '9', next: []State{5}},
				{s: 5, start: 'A', end: 'F', next: []State{5}},
				{s: 5, start: 'a', end: 'f', next: []State{5}},
			},
			expectedNFA: testNFA[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(NFABuilder).SetStart(tc.start).SetFinal(tc.final...)

			for _, tr := range tc.trans {
				b.AddTransition(tr.s, tr.start, tr.end, tr.next)
			}

			t.Run("Build", func(t *testing.T) {
				nfa := b.Build()
				assert.True(t, nfa.Equal(tc.expectedNFA), "Expected:\n%s\nGot:\n%s", tc.expectedNFA.trans, nfa.trans)
			})
		})
	}
}

func TestNFA_String(t *testing.T) {
	tests := []struct {
		name           string
		n              *NFA
		expectedString string
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedString: `Start state: 0
Final states: 2, 4
Equivalence Classes:
  [ε..ε]: 0
  [a..a]: 1
  [b..b]: 2
Transitions:
  0 --0--> {1, 3}
  1 --1--> {2}
  2 --1--> {2}
  3 --2--> {4}
  4 --2--> {4}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.n.String())
		})
	}
}

func TestNFA_Clone(t *testing.T) {
	n := testNFA[0].Clone()
	n.states = []State{0, 1, 2, 3, 4}

	tests := []struct {
		name string
		n    *NFA
	}{
		{
			name: "OK",
			n:    n,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.n.Clone()

			assert.NotSame(t, clone, tc.n)
			assert.True(t, clone.Equal(tc.n))
		})
	}
}

func TestNFA_Equal(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		rhs           *NFA
		expectedEqual bool
	}{
		{
			name:          "NotEqual_Nil",
			n:             testNFA[0],
			rhs:           nil,
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffStart",
			n:    testNFA[0],
			rhs: &NFA{
				start: 1,
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffFinal",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(1, 3),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffClasses",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				classes: disc.NewRangeMap(eqClassID, classesOpts, []disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 2},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffTransitions",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				classes: disc.NewRangeMap(eqClassID, classesOpts, []disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)).
					Add(4, 2, NewStates(4)),
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			n:    testNFA[0],
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				classes: disc.NewRangeMap(eqClassID, classesOpts, []disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 3)).
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)).
					Add(4, 2, NewStates(4)),
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.n.Equal(tc.rhs))
		})
	}
}

func TestNFA_Start(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		expectedStart State
	}{
		{
			name:          "OK",
			n:             testNFA[0],
			expectedStart: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStart, tc.n.Start())
		})
	}
}

func TestNFA_Final(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		expectedFinal []State
	}{
		{
			name:          "OK",
			n:             testNFA[0],
			expectedFinal: []State{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedFinal, tc.n.Final())
		})
	}
}

func TestNFA_States(t *testing.T) {
	tests := []struct {
		name           string
		n              *NFA
		expectedStates []State
	}{
		{
			name:           "OK",
			n:              testNFA[0],
			expectedStates: []State{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.n.States())
		})
	}
}
