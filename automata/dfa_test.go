package automata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
)

var testDFA = []*DFA{
	// 1(0|1)*
	{
		start: 0,
		final: NewStates(1),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: '1', Hi: '1'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 1, 1).
			Add(1, 0, 1).
			Add(1, 1, 1),
	},
	// ([A-Za-z_][0-9A-Za-z_]*)|[0-9]+|(0x[0-9A-Fa-f]+)|[ \t\n]+|[+\-*/=]
	{
		start: 0,
		final: NewStates(1, 2, 3, 5),
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
		trans: newDFATransitionTable().
			Add(0, 0, 3).
			Add(0, 1, 2).
			Add(0, 2, 1).
			Add(0, 3, 1).
			Add(0, 4, 1).
			Add(1, 0, 1).
			Add(1, 1, 1).
			Add(1, 2, 1).
			Add(1, 3, 1).
			Add(1, 4, 1).
			Add(2, 0, 2).
			Add(2, 1, 2).
			Add(3, 0, 2).
			Add(3, 1, 2).
			Add(3, 4, 4).
			Add(4, 0, 5).
			Add(4, 1, 5).
			Add(4, 2, 5).
			Add(5, 0, 5).
			Add(5, 1, 5).
			Add(5, 2, 5),
	},
}

func TestDFABuilder(t *testing.T) {
	type transition struct {
		s          State
		start, end Symbol
		next       State
	}

	tests := []struct {
		name        string
		start       State
		final       []State
		trans       []transition
		expectedDFA *DFA
	}{
		{
			name:  "Simple",
			start: 0,
			final: []State{1},
			trans: []transition{
				{s: 0, start: '1', end: '1', next: 1},
				{s: 1, start: '0', end: '1', next: 1},
			},
			expectedDFA: testDFA[0],
		},
		{
			name:  "ID_NUM_WS_OP",
			start: 0,
			final: []State{1, 2, 3, 5},
			trans: []transition{
				{s: 0, start: '0', end: '0', next: 3},
				{s: 0, start: '1', end: '9', next: 2},
				{s: 0, start: 'A', end: 'Z', next: 1},
				{s: 0, start: 'a', end: 'z', next: 1},
				{s: 0, start: '_', end: '_', next: 1},
				{s: 1, start: '0', end: '9', next: 1},
				{s: 1, start: 'A', end: 'Z', next: 1},
				{s: 1, start: 'a', end: 'z', next: 1},
				{s: 1, start: '_', end: '_', next: 1},
				{s: 2, start: '0', end: '9', next: 2},
				{s: 3, start: '0', end: '9', next: 2},
				{s: 3, start: 'X', end: 'X', next: 4},
				{s: 3, start: 'x', end: 'x', next: 4},
				{s: 4, start: '0', end: '9', next: 5},
				{s: 4, start: 'A', end: 'F', next: 5},
				{s: 4, start: 'a', end: 'f', next: 5},
				{s: 5, start: '0', end: '9', next: 5},
				{s: 5, start: 'A', end: 'F', next: 5},
				{s: 5, start: 'a', end: 'f', next: 5},
			},
			expectedDFA: testDFA[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(DFABuilder).SetStart(tc.start).SetFinal(tc.final...)

			for _, tr := range tc.trans {
				b.AddTransition(tr.s, tr.start, tr.end, tr.next)
			}

			t.Run("Build", func(t *testing.T) {
				dfa := b.Build()
				assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA.trans, dfa.trans)
			})
		})
	}
}

func TestDFA_String(t *testing.T) {
	tests := []struct {
		name           string
		d              *DFA
		expectedString string
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedString: `Start state: 0
Final states: 1
Transitions:
  0 -- [1..1] --> 1
  1 -- [0..0] --> 1
  1 -- [1..1] --> 1
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.d.String())
		})
	}
}

func TestDFA_Clone(t *testing.T) {
	tests := []struct {
		name string
		d    *DFA
	}{
		{
			name: "OK",
			d:    testDFA[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.d.Clone()

			assert.NotSame(t, clone, tc.d)
			assert.True(t, clone.Equal(tc.d))
		})
	}
}

func TestDFA_Equal(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		rhs           *DFA
		expectedEqual bool
	}{
		{
			name:          "NotEqual_Nil",
			d:             testDFA[0],
			rhs:           nil,
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffStart",
			d:    testDFA[0],
			rhs: &DFA{
				start: 1,
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffFinal",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(2),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffClasses",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '1'}, Value: 0},
				}),
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffTransitions",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '1'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(0, 1, 1).
					Add(1, 0, 1).
					Add(1, 1, 1),
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			d:    testDFA[0],
			rhs: &DFA{
				start: 0,
				final: NewStates(1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '0'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '1', Hi: '1'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 1, 1).
					Add(1, 0, 1).
					Add(1, 1, 1),
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.d.Equal(tc.rhs))
		})
	}
}

func TestDFA_Start(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		expectedStart State
	}{
		{
			name:          "OK",
			d:             testDFA[0],
			expectedStart: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStart, tc.d.Start())
		})
	}
}

func TestDFA_Final(t *testing.T) {
	tests := []struct {
		name          string
		d             *DFA
		expectedFinal []State
	}{
		{
			name:          "OK",
			d:             testDFA[0],
			expectedFinal: []State{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedFinal, tc.d.Final())
		})
	}
}

func TestDFA_States(t *testing.T) {
	tests := []struct {
		name           string
		d              *DFA
		expectedStates []State
	}{
		{
			name:           "OK",
			d:              testDFA[0],
			expectedStates: []State{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedStates, tc.d.States())
		})
	}
}

func TestDFA_Symbols(t *testing.T) {
	tests := []struct {
		name            string
		d               *DFA
		expectedSymbols []disc.Range[Symbol]
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedSymbols: []disc.Range[Symbol]{
				{Lo: '0', Hi: '0'},
				{Lo: '1', Hi: '1'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSymbols, tc.d.Symbols())
		})
	}
}

func TestDFA_classes(t *testing.T) {
	tests := []struct {
		name            string
		d               *DFA
		expectedClasses classMapping
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedClasses: newClassMapping([]generic.KeyValue[classID, rangeSet]{
				{Key: 0, Val: newRangeSet(disc.Range[Symbol]{Lo: '0', Hi: '0'})},
				{Key: 1, Val: newRangeSet(disc.Range[Symbol]{Lo: '1', Hi: '1'})},
			}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.d.classes().Equal(tc.expectedClasses))
		})
	}
}

func TestDFA_Transitions(t *testing.T) {
	type transition struct {
		s      State
		ranges []disc.Range[Symbol]
		next   State
	}

	eqTransition := func(a, b transition) bool {
		return a.s == b.s && reflect.DeepEqual(a.ranges, b.ranges) && a.next == b.next
	}

	tests := []struct {
		name          string
		d             *DFA
		expectedTrans set.Set[transition]
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedTrans: set.New(eqTransition,
				transition{0, []disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
				transition{1, []disc.Range[Symbol]{{Lo: '0', Hi: '0'}}, 1},
				transition{1, []disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := set.New(eqTransition)
			for s, seq := range tc.d.Transitions() {
				for ranges, next := range seq {
					trans.Add(transition{s, ranges, next})
				}
			}

			assert.True(t, trans.Equal(tc.expectedTrans))
		})
	}
}

func TestDFA_TransitionsFrom(t *testing.T) {
	type transition struct {
		ranges []disc.Range[Symbol]
		next   State
	}

	eqTransition := func(a, b transition) bool {
		return reflect.DeepEqual(a.ranges, b.ranges) && a.next == b.next
	}

	tests := []struct {
		name          string
		d             *DFA
		s             State
		expectedTrans set.Set[transition]
	}{
		{
			name: "OK",
			d:    testDFA[0],
			s:    1,
			expectedTrans: set.New(eqTransition,
				transition{[]disc.Range[Symbol]{{Lo: '0', Hi: '0'}}, 1},
				transition{[]disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := set.New(eqTransition)
			for ranges, next := range tc.d.TransitionsFrom(tc.s) {
				trans.Add(transition{ranges, next})
			}

			assert.True(t, trans.Equal(tc.expectedTrans))
		})
	}
}
