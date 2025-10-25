package automata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
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
	// (a|b)*abb
	{
		start: 0,
		final: NewStates(4),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 0, 1).
			Add(0, 1, 2).
			Add(1, 0, 1).
			Add(1, 1, 3).
			Add(2, 0, 1).
			Add(2, 1, 2).
			Add(3, 0, 1).
			Add(3, 1, 4).
			Add(4, 0, 1).
			Add(4, 1, 2),
	},
	// (ab)+
	{
		start: 0,
		final: NewStates(2),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 0, 1).
			Add(1, 1, 2).
			Add(2, 0, 1),
	},
	// ab+|ba+
	{
		start: 0,
		final: NewStates(2, 4),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 0, 1).
			Add(1, 1, 2).
			Add(2, 1, 2).
			Add(0, 1, 3).
			Add(3, 0, 4).
			Add(4, 0, 4),
	},
	// ab(a|b)*
	{
		start: 0,
		final: NewStates(2),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 0, 1).
			Add(1, 1, 2).
			Add(2, 0, 2).
			Add(2, 1, 2),
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
			expectedDFA: testDFA[5],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := NewDFABuilder().SetStart(tc.start).SetFinal(tc.final)

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

func TestDFA_Isomorphic(t *testing.T) {
	dfa := &DFA{
		start: 0,
		final: NewStates(1, 2),
		ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
			{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
			{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
		}),
		trans: newDFATransitionTable().
			Add(0, 0, 1).
			Add(1, 0, 1).
			Add(0, 1, 2).
			Add(2, 1, 2),
	}

	tests := []struct {
		name               string
		d                  *DFA
		rhs                *DFA
		expectedIsomorphic bool
	}{
		{
			name: "DiffFinalLens",
			d:    dfa,
			rhs: &DFA{
				start: 0,
				final: NewStates(2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable(),
			},
			expectedIsomorphic: false,
		},
		{
			name: "DiffStateLens",
			d:    dfa,
			rhs: &DFA{
				start: 0,
				final: NewStates(1, 2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(1, 0, 1).
					Add(0, 1, 2).
					Add(2, 1, 2).
					Add(2, 0, 3).
					Add(2, 1, 3),
			},
			expectedIsomorphic: false,
		},
		{
			name: "DiffAlphabetLens",
			d:    dfa,
			rhs: &DFA{
				start: 0,
				final: NewStates(1, 2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'c', Hi: 'c'}, Value: 2},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(1, 0, 1).
					Add(0, 1, 2).
					Add(2, 1, 2),
			},
			expectedIsomorphic: false,
		},
		{
			name: "AlphabetNotEqual",
			d:    dfa,
			rhs: &DFA{
				start: 0,
				final: NewStates(1, 2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'z', Hi: 'z'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(1, 0, 1).
					Add(0, 1, 2).
					Add(2, 1, 2),
			},
			expectedIsomorphic: false,
		},
		{
			name: "SortedDegreesNotEqual",
			d:    dfa,
			rhs: &DFA{
				start: 0,
				final: NewStates(1, 2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(1, 0, 1).
					Add(0, 1, 2),
			},
			expectedIsomorphic: false,
		},
		{
			name:               "Equal",
			d:                  dfa,
			rhs:                dfa,
			expectedIsomorphic: true,
		},
		{
			name: "Isomorphic",
			d:    dfa,
			rhs: &DFA{
				start: 2,
				final: NewStates(0, 1),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(2, 0, 0).
					Add(0, 0, 0).
					Add(2, 1, 1).
					Add(1, 1, 1),
			},
			expectedIsomorphic: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsomorphic, tc.d.Isomorphic(tc.rhs))
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

	tests := []struct {
		name          string
		d             *DFA
		expectedTrans []transition
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedTrans: []transition{
				{0, []disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
				{1, []disc.Range[Symbol]{{Lo: '0', Hi: '0'}}, 1},
				{1, []disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := []transition{}
			for s, seq := range tc.d.Transitions() {
				for ranges, next := range seq {
					trans = append(trans, transition{s, ranges, next})
				}
			}

			assert.True(t, reflect.DeepEqual(trans, tc.expectedTrans))
		})
	}
}

func TestDFA_TransitionsFrom(t *testing.T) {
	type transition struct {
		ranges []disc.Range[Symbol]
		next   State
	}

	tests := []struct {
		name          string
		d             *DFA
		s             State
		expectedTrans []transition
	}{
		{
			name: "OK",
			d:    testDFA[0],
			s:    1,
			expectedTrans: []transition{
				{[]disc.Range[Symbol]{{Lo: '0', Hi: '0'}}, 1},
				{[]disc.Range[Symbol]{{Lo: '1', Hi: '1'}}, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			trans := []transition{}
			for ranges, next := range tc.d.TransitionsFrom(tc.s) {
				trans = append(trans, transition{ranges, next})
			}

			assert.True(t, reflect.DeepEqual(trans, tc.expectedTrans))
		})
	}
}

func TestDFA_Minimize(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name: "OK",
			d:    testDFA[1],
			expectedDFA: &DFA{
				start: 0,
				final: NewStates(3),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(0, 1, 0).
					Add(1, 0, 1).
					Add(1, 1, 2).
					Add(2, 0, 1).
					Add(2, 1, 3).
					Add(3, 0, 1).
					Add(3, 1, 0),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.Minimize()

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
		})
	}
}

func TestBuildGroupTransitions(t *testing.T) {
	tests := []struct {
		name               string
		d                  *DFA
		P                  *partition
		G                  group
		expectedGroupTrans *dfaTransitionTable
	}{
		{
			name: "OK",
			d:    testDFA[3],
			P: &partition{
				groups: newGroups(
					group{States: NewStates(0, 1, 3), Rep: 0},
					group{States: NewStates(2, 4), Rep: 1},
				),
				nextRep: 2,
			},
			G: group{States: NewStates(2, 4), Rep: 1},
			expectedGroupTrans: newDFATransitionTable().
				Add(2, 1, 1).
				Add(4, 0, 1),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Gtrans := buildGroupTransitions(tc.d, tc.P, tc.G)

			assert.True(t, Gtrans.Equal(tc.expectedGroupTrans))
		})
	}
}

func TestPartitionGroups(t *testing.T) {
	tests := []struct {
		name              string
		P                 *partition
		Gtrans            *dfaTransitionTable
		expectedPartition *partition
	}{
		{
			name: "OK",
			P: &partition{
				groups:  newGroups(),
				nextRep: 0,
			},
			Gtrans: newDFATransitionTable().
				Add(2, 1, 1).
				Add(4, 0, 1),
			expectedPartition: &partition{
				groups: newGroups(
					group{States: NewStates(2), Rep: 0},
					group{States: NewStates(4), Rep: 1},
				),
				nextRep: 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			partitionGroup(tc.P, tc.Gtrans)

			assert.True(t, tc.P.Equal(tc.expectedPartition))
		})
	}
}

func TestDFA_EliminateDeadStates(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name: "OK",
			d: &DFA{ // ab(a|b)*
				start: 0,
				final: NewStates(2),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(0, 1, 3).
					Add(1, 0, 4).
					Add(1, 1, 2).
					Add(2, 0, 2).
					Add(2, 1, 2).
					Add(3, 0, 3).
					Add(3, 1, 3).
					Add(4, 0, 4).
					Add(4, 1, 4),
			},
			expectedDFA: testDFA[4],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.EliminateDeadStates()

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
		})
	}
}

func TestDFA_ReindexStates(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedDFA *DFA
	}{
		{
			name: "OK",
			d: &DFA{ // (0|1)+\.(0|1)+
				start: 0,
				final: NewStates(3, 4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '.', Hi: '.'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '1'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 1, 3).
					Add(3, 1, 3).
					Add(3, 0, 1).
					Add(1, 1, 4).
					Add(4, 1, 4),
			},
			expectedDFA: &DFA{ // (0|1)+\.(0|1)+
				start: 0,
				final: NewStates(1, 3),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: '.', Hi: '.'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: '0', Hi: '1'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 1, 1).
					Add(1, 1, 1).
					Add(1, 0, 2).
					Add(2, 1, 3).
					Add(3, 1, 3),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa := tc.d.ReindexStates()

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
		})
	}
}

func TestDFA_Union(t *testing.T) {
	tests := []struct {
		name             string
		d                *DFA
		ds               []*DFA
		expectedDFA      *DFA
		expectedFinalMap [][]State
	}{
		{
			name: "OK",
			d:    testDFA[2],
			ds: []*DFA{
				testDFA[3],
				testDFA[4],
			},
			expectedDFA: &DFA{
				start: 0,
				final: NewStates(3, 4, 5, 6, 7, 8),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(0, 1, 2).
					Add(1, 1, 3).
					Add(2, 0, 4).
					Add(3, 0, 5).
					Add(3, 1, 6).
					Add(4, 0, 4).
					Add(5, 0, 7).
					Add(5, 1, 8).
					Add(6, 0, 7).
					Add(6, 1, 6).
					Add(7, 0, 7).
					Add(7, 1, 7).
					Add(8, 0, 5).
					Add(8, 1, 7),
			},
			expectedFinalMap: [][]State{
				{3, 8},
				{3, 4, 6},
				{3, 5, 6, 7, 8},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa, finalMap := tc.d.Union(tc.ds...)

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
			assert.Equal(t, tc.expectedFinalMap, finalMap, "Expected:\n%v\nGot:\n%v", tc.expectedFinalMap, finalMap)
		})
	}
}

func TestUnionDFA(t *testing.T) {
	tests := []struct {
		name             string
		ds               []*DFA
		expectedDFA      *DFA
		expectedFinalMap [][]State
	}{
		{
			name: "OK",
			ds: []*DFA{
				testDFA[2],
				testDFA[3],
				testDFA[4],
			},
			expectedDFA: &DFA{
				start: 0,
				final: NewStates(3, 4, 5, 6, 7, 8),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newDFATransitionTable().
					Add(0, 0, 1).
					Add(0, 1, 2).
					Add(1, 1, 3).
					Add(2, 0, 4).
					Add(3, 0, 5).
					Add(3, 1, 6).
					Add(4, 0, 4).
					Add(5, 0, 7).
					Add(5, 1, 8).
					Add(6, 0, 7).
					Add(6, 1, 6).
					Add(7, 0, 7).
					Add(7, 1, 7).
					Add(8, 0, 5).
					Add(8, 1, 7),
			},
			expectedFinalMap: [][]State{
				{3, 8},
				{3, 4, 6},
				{3, 5, 6, 7, 8},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dfa, finalMap := UnionDFA(tc.ds...)

			assert.True(t, dfa.Equal(tc.expectedDFA), "Expected:\n%s\nGot:\n%s", tc.expectedDFA, dfa)
			assert.Equal(t, tc.expectedFinalMap, finalMap, "Expected:\n%v\nGot:\n%v", tc.expectedFinalMap, finalMap)
		})
	}
}

func TestDFA_ToNFA(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedNFA *NFA
	}{
		{
			name: "OK",
			d:    testDFA[1],
			expectedNFA: &NFA{
				start: 0,
				final: NewStates(4),
				ranges: newRangeMapping([]disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 1},
				}),
				trans: newNFATransitionTable().
					Add(0, 0, NewStates(1)).
					Add(0, 1, NewStates(2)).
					Add(1, 0, NewStates(1)).
					Add(1, 1, NewStates(3)).
					Add(2, 0, NewStates(1)).
					Add(2, 1, NewStates(2)).
					Add(3, 0, NewStates(1)).
					Add(3, 1, NewStates(4)).
					Add(4, 0, NewStates(1)).
					Add(4, 1, NewStates(2)),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nfa := tc.d.ToNFA()

			assert.True(t, nfa.Equal(tc.expectedNFA), "Expected:\n%s\nGot:\n%s", tc.expectedNFA, nfa)
		})
	}
}

func TestDFA_DOT(t *testing.T) {
	tests := []struct {
		name        string
		d           *DFA
		expectedDOT string
	}{
		{
			name: "OK",
			d:    testDFA[0],
			expectedDOT: `digraph "DFA" {
  rankdir=LR;
  concentrate=false;
  node [shape=circle];

  start [style=invis];
  0 [label="0"];
  1 [label="1", shape=doublecircle];

  start -> 0 [];
  0 -> 1 [label="[1..1]"];
  1 -> 1 [label="[0..0]"];
  1 -> 1 [label="[1..1]"];
}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedDOT, tc.d.DOT())
		})
	}
}

func TestDFA_Runner(t *testing.T) {
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
			r := tc.d.Runner()

			assert.NotNil(t, r)
			assert.NotNil(t, r.trans)
			assert.Equal(t, tc.d.start, r.start)
			assert.True(t, r.final.Equal(tc.d.final))
			assert.True(t, r.ranges.Equal(tc.d.ranges))
			assert.NotNil(t, r.trans)
		})
	}
}

func TestDFARunner_Next(t *testing.T) {
	runner := testDFA[0].Runner()

	tests := []struct {
		name         string
		r            *DFARunner
		s            State
		a            Symbol
		expectedNext State
	}{
		{
			name:         "NotOK",
			r:            runner,
			s:            0,
			a:            '0',
			expectedNext: -1,
		},
		{
			name:         "OK",
			r:            runner,
			s:            0,
			a:            '1',
			expectedNext: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedNext, tc.r.Next(tc.s, tc.a))
		})
	}
}

func TestDFARunner_Accept(t *testing.T) {
	runner := testDFA[5].Runner()

	tests := []struct {
		name           string
		r              *DFARunner
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
