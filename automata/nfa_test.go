package automata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
)

var testNFA = []*NFA{
	// (a+|b+)
	{
		start: 0,
		final: NewStates(2, 4),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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
	// (a|b)*abb
	{
		start: 0,
		final: NewStates(10),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(1, 7)).
			Add(1, 0, NewStates(2, 4)).
			Add(2, 1, NewStates(3)).
			Add(3, 0, NewStates(6)).
			Add(4, 2, NewStates(5)).
			Add(5, 0, NewStates(6)).
			Add(6, 0, NewStates(1, 7)).
			Add(7, 1, NewStates(8)).
			Add(8, 2, NewStates(9)).
			Add(9, 2, NewStates(10)),
	},
	// [A-Z][A-Za-z]*
	{
		start: 0,
		final: NewStates(1),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'Z'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'z'}, Value: 1},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(1)).
			Add(1, 0, NewStates(1)).
			Add(1, 1, NewStates(1)),
	},
	// 0|[1-9][0-9]*
	{
		start: 0,
		final: NewStates(1, 2),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 1},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(1)).
			Add(0, 1, NewStates(2)).
			Add(2, 0, NewStates(2)).
			Add(2, 1, NewStates(2)),
	},
	// 0|0x[0-9A-Fa-f]+
	{
		start: 0,
		final: NewStates(1, 3),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'F'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'X', Hi: 'X'}, Value: 2},
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'f'}, Value: 1},
			{Range: disc.Range[Symbol]{Lo: 'x', Hi: 'x'}, Value: 2},
		}),
		trans: newNFATransitionTable().
			Add(0, 0, NewStates(1)).
			Add(1, 2, NewStates(2)).
			Add(2, 0, NewStates(3)).
			Add(2, 1, NewStates(3)).
			Add(3, 0, NewStates(3)).
			Add(3, 1, NewStates(3)),
	},
	// ([A-Za-z_][0-9A-Za-z_]*)|[0-9]+|(0x[0-9A-Fa-f]+)|[ \t\n]+|[+\-*/=]
	{
		start: 0,
		final: NewStates(1, 2, 5),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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
			expectedNFA: testNFA[5],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := NewNFABuilder().SetStart(tc.start).SetFinal(tc.final)

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
Transitions:
  0 -- [ε..ε] --> {1, 3}
  1 -- [a..a] --> {2}
  2 -- [a..a] --> {2}
  3 -- [b..b] --> {4}
  4 -- [b..b] --> {4}
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
	tests := []struct {
		name string
		n    *NFA
	}{
		{
			name: "OK",
			n:    testNFA[0],
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
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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

func TestNFA_Isomorphic(t *testing.T) {
	nfa := &NFA{
		start: 0,
		final: NewStates(2, 4),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
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
	}

	tests := []struct {
		name               string
		n                  *NFA
		rhs                *NFA
		expectedIsomorphic bool
	}{
		{
			name: "DiffFinalLens",
			n:    nfa,
			rhs: &NFA{
				start: 0,
				final: NewStates(2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable(),
			},
			expectedIsomorphic: false,
		},
		{
			name: "DiffStateLens",
			n:    nfa,
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 3)).
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)).
					Add(4, 2, NewStates(4)).
					Add(4, 1, NewStates(5)).
					Add(4, 2, NewStates(5)),
			},
			expectedIsomorphic: false,
		},
		{
			name: "DiffAlphabetLens",
			n:    nfa,
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
					{Range: disc.Range[Symbol]{Lo: 'c', Hi: 'c'}, Value: 3},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 3)).
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)).
					Add(4, 2, NewStates(4)),
			},
			expectedIsomorphic: false,
		},
		{
			name: "AlphabetNotEqual",
			n:    nfa,
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'z', Hi: 'z'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 3)).
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)).
					Add(4, 2, NewStates(4)),
			},
			expectedIsomorphic: false,
		},
		{
			name: "SortedDegreesNotEqual",
			n:    nfa,
			rhs: &NFA{
				start: 0,
				final: NewStates(2, 4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 3)).
					Add(1, 1, NewStates(2)).
					Add(2, 1, NewStates(2)).
					Add(3, 2, NewStates(4)),
			},
			expectedIsomorphic: false,
		},
		{
			name:               "Equal",
			n:                  nfa,
			rhs:                nfa,
			expectedIsomorphic: true,
		},
		{
			name: "Isomorphic",
			n:    nfa,
			rhs: &NFA{
				start: 4,
				final: NewStates(0, 1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(4, 0, NewStates(2, 3)).
					Add(2, 1, NewStates(0)).
					Add(0, 1, NewStates(0)).
					Add(3, 2, NewStates(1)).
					Add(1, 2, NewStates(1)),
			},
			expectedIsomorphic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsomorphic, tc.n.Isomorphic(tc.rhs))
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

func TestNFA_Symbols(t *testing.T) {
	tests := []struct {
		name            string
		n               *NFA
		expectedSymbols []disc.Range[Symbol]
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedSymbols: []disc.Range[Symbol]{
				{Lo: 'a', Hi: 'a'},
				{Lo: 'b', Hi: 'b'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.n.Symbols())
		})
	}
}

func TestNFA_classes(t *testing.T) {
	tests := []struct {
		name            string
		n               *NFA
		expectedClasses classMapping
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedClasses: newClassMapping([]generic.KeyValue[classID, rangeSet]{
				{Key: 0, Val: newRangeSet(disc.Range[Symbol]{Lo: E, Hi: E})},
				{Key: 1, Val: newRangeSet(disc.Range[Symbol]{Lo: 'a', Hi: 'a'})},
				{Key: 2, Val: newRangeSet(disc.Range[Symbol]{Lo: 'b', Hi: 'b'})},
			}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.n.classes().Equal(tc.expectedClasses))
		})
	}
}

func TestNFA_Transitions(t *testing.T) {
	type transition struct {
		s      State
		ranges []disc.Range[Symbol]
		next   []State
	}

	tests := []struct {
		name          string
		n             *NFA
		expectedTrans []transition
	}{
		{
			name: "OK",
			n:    testNFA[0],
			expectedTrans: []transition{
				{0, []disc.Range[Symbol]{{Lo: E, Hi: E}}, []State{1, 3}},
				{1, []disc.Range[Symbol]{{Lo: 'a', Hi: 'a'}}, []State{2}},
				{2, []disc.Range[Symbol]{{Lo: 'a', Hi: 'a'}}, []State{2}},
				{3, []disc.Range[Symbol]{{Lo: 'b', Hi: 'b'}}, []State{4}},
				{4, []disc.Range[Symbol]{{Lo: 'b', Hi: 'b'}}, []State{4}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := []transition{}
			for s, seq := range tc.n.Transitions() {
				for ranges, next := range seq {
					trans = append(trans, transition{s, ranges, next})
				}
			}

			assert.True(t, reflect.DeepEqual(trans, tc.expectedTrans))
		})
	}
}

func TestNFA_TransitionsFrom(t *testing.T) {
	type transition struct {
		ranges []disc.Range[Symbol]
		next   []State
	}

	tests := []struct {
		name          string
		n             *NFA
		s             State
		expectedTrans []transition
	}{
		{
			name: "OK",
			n:    testNFA[0],
			s:    0,
			expectedTrans: []transition{
				{[]disc.Range[Symbol]{{Lo: E, Hi: E}}, []State{1, 3}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := []transition{}
			for ranges, next := range tc.n.TransitionsFrom(tc.s) {
				trans = append(trans, transition{ranges, next})
			}

			assert.True(t, reflect.DeepEqual(trans, tc.expectedTrans))
		})
	}
}

func TestNFA_Star(t *testing.T) {
	tests := []struct {
		name         string
		n            *NFA
		expectedStar *NFA
	}{
		{
			name: "OK",
			n:    testNFA[3],
			expectedStar: &NFA{
				start: 0,
				final: NewStates(1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 2},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1, 2)).
					Add(2, 1, NewStates(3)).
					Add(2, 2, NewStates(4)).
					Add(3, 0, NewStates(1, 2)).
					Add(4, 0, NewStates(1, 2)).
					Add(4, 1, NewStates(4)).
					Add(4, 2, NewStates(4)),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			star := tc.n.Star()

			assert.True(t, star.Equal(tc.expectedStar), "Expected:\n%s\nGot:\n%s", tc.expectedStar, star)
		})
	}
}

func TestNFA_Union(t *testing.T) {
	tests := []struct {
		name          string
		n             *NFA
		ns            []*NFA
		expectedUnion *NFA
	}{
		{
			name: "OK",
			n:    testNFA[3],
			ns:   []*NFA{testNFA[4]},
			expectedUnion: &NFA{
				start: 0,
				final: NewStates(1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 2},
					{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'F'}, Value: 3},
					{Range: disc.Range[Symbol]{Lo: 'X', Hi: 'X'}, Value: 4},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'f'}, Value: 3},
					{Range: disc.Range[Symbol]{Lo: 'x', Hi: 'x'}, Value: 4},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(2, 5)).
					Add(2, 1, NewStates(3)).
					Add(2, 2, NewStates(4)).
					Add(3, 0, NewStates(1)).
					Add(4, 0, NewStates(1)).
					Add(4, 1, NewStates(4)).
					Add(4, 2, NewStates(4)).
					Add(5, 1, NewStates(6)).
					Add(6, 0, NewStates(1)).
					Add(6, 4, NewStates(7)).
					Add(7, 1, NewStates(8)).
					Add(7, 2, NewStates(8)).
					Add(7, 3, NewStates(8)).
					Add(8, 0, NewStates(1)).
					Add(8, 1, NewStates(8)).
					Add(8, 2, NewStates(8)).
					Add(8, 3, NewStates(8)),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			union := tc.n.Union(tc.ns...)

			assert.True(t, union.Equal(tc.expectedUnion), "Expected:\n%s\nGot:\n%s", tc.expectedUnion, union)
		})
	}
}

func TestNFA_Concat(t *testing.T) {
	tests := []struct {
		name           string
		n              *NFA
		ns             []*NFA
		expectedConcat *NFA
	}{
		{
			name: "OK",
			n:    testNFA[2],
			ns:   []*NFA{testNFA[3]},
			expectedConcat: &NFA{
				start: 0,
				final: NewStates(2, 3),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '9'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'Z'}, Value: 2},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'z'}, Value: 3},
				}),
				trans: newNFATransitionTable().
					Add(0, 2, NewStates(1)).
					Add(1, 0, NewStates(2)).
					Add(1, 1, NewStates(3)).
					Add(1, 2, NewStates(1)).
					Add(1, 3, NewStates(1)).
					Add(3, 0, NewStates(3)).
					Add(3, 1, NewStates(3)),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			concat := tc.n.Concat(tc.ns...)

			assert.True(t, concat.Equal(tc.expectedConcat), "Expected:\n%s\nGot:\n%s", tc.expectedConcat, concat)
		})
	}
}

func TestNFA_ToDFA(t *testing.T) {
	tests := []struct {
		name        string
		n           *NFA
		expectedDFA *DFA
	}{
		{
			name:        "OK",
			n:           testNFA[1],
			expectedDFA: testDFA[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.n.ToDFA()

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
		})
	}
}

func TestNFA_DOT(t *testing.T) {
	tests := []struct {
		name        string
		n           *NFA
		expectedDOT string
	}{
		{
			name: "OK",
			n:    testNFA[0],
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
  0 -> 1 [label="[ε..ε]"];
  0 -> 3 [label="[ε..ε]"];
  1 -> 2 [label="[a..a]"];
  2 -> 2 [label="[a..a]"];
  3 -> 4 [label="[b..b]"];
  4 -> 4 [label="[b..b]"];
}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.n.DOT())
		})
	}
}

func TestNFA_Runner(t *testing.T) {
	tests := []struct {
		name string
		n    *NFA
	}{
		{
			name: "OK",
			n:    testNFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.n.Runner()

			assert.NotNil(t, r)
			assert.NotNil(t, r.trans)
			assert.Equal(t, tc.n.start, r.start)
			assert.True(t, r.final.Equal(tc.n.final))
			assert.True(t, r.ranges.Equal(tc.n.ranges))
		})
	}
}

func TestNFARunner_Next(t *testing.T) {
	runner := testNFA[0].Runner()

	tests := []struct {
		name         string
		r            *NFARunner
		s            State
		a            Symbol
		expectedNext []State
	}{
		{
			name:         "NotOK",
			r:            runner,
			s:            0,
			a:            'a',
			expectedNext: nil,
		},
		{
			name:         "OK",
			r:            runner,
			s:            0,
			a:            E,
			expectedNext: []State{1, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedNext, tc.r.Next(tc.s, tc.a))
		})
	}
}

func TestNFARunner_Accept(t *testing.T) {
	runner := testNFA[5].Runner()

	tests := []struct {
		name           string
		r              *NFARunner
		s              String
		expectedAccept bool
	}{
		{
			name:           "EmptyString",
			r:              runner,
			s:              String{},
			expectedAccept: false,
		},
		{
			name:           "NotAccepted",
			r:              runner,
			s:              String{'0', '1', '_', 'I', 'd'},
			expectedAccept: false,
		},
		{
			name:           "Accepted",
			r:              runner,
			s:              String{'I', 'd', '_', '0', '1'},
			expectedAccept: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAccept, tc.r.Accept(tc.s))
		})
	}
}
