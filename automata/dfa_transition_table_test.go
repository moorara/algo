package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeState(t *testing.T) {
	tests := []struct {
		name           string
		rs             rangeState
		expectedString string
		equal          rangeState
		expectedEqual  bool
	}{
		{
			name: "Equal",
			rs: rangeState{
				SymbolRange{Start: '0', End: '9'},
				1,
			},
			expectedString: "[0..9] → 1",
			equal: rangeState{
				SymbolRange{Start: '0', End: '9'},
				1,
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			rs: rangeState{
				SymbolRange{Start: '0', End: '9'},
				1,
			},
			expectedString: "[0..9] → 1",
			equal: rangeState{
				SymbolRange{Start: '0', End: '9'},
				3,
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.rs.String())
			assert.Equal(t, tc.expectedEqual, tc.rs.Equal(tc.equal))
		})
	}
}

func TestDFATransitionTable(t *testing.T) {
	type equalTest struct {
		rhs           *dfaTransitionTable
		expectedEqual bool
	}

	type addTest struct {
		s     State
		start Symbol
		end   Symbol
		next  State
	}

	type nextTest struct {
		s            State
		a            Symbol
		expectedNext State
		expectedOK   bool
	}

	type fromTest struct {
		s            State
		expectedFrom []generic.KeyValue[SymbolRange, State]
	}

	type transition struct {
		s    State
		r    SymbolRange
		next State
	}

	tests := []struct {
		name           string
		trans          map[State][]rangeState
		equalTests     []equalTest
		addTests       []addTest
		nextTests      []nextTest
		fromTests      []fromTest
		expectedAll    []transition
		expectedString string
	}{
		{
			name: "CurrentEndOnLastEnd_SameStates",
			trans: map[State][]rangeState{
				0: {
					{SymbolRange{Start: '0', End: '9'}, 0},
					{SymbolRange{Start: 'a', End: 'n'}, 10},
					{SymbolRange{Start: 'n', End: 'n'}, 10},
					{SymbolRange{Start: 'n', End: 'z'}, 10},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newDFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'z'}, 100},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'z'}, 10},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: 20},
				{s: 0, start: 'A', end: 'N', next: 30},
				{s: 0, start: 'N', end: 'N', next: 30},
				{s: 0, start: 'N', end: 'Z', next: 30},
				{s: 0, start: 'α', end: 'δ', next: 40},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: 0, expectedOK: true},
				{s: 0, a: '5', expectedNext: 0, expectedOK: true},
				{s: 0, a: '9', expectedNext: 0, expectedOK: true},
				{s: 0, a: 'a', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'm', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'x', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'z', expectedNext: 10, expectedOK: true},
				{s: 0, a: '+', expectedNext: 20, expectedOK: true},
				{s: 0, a: ',', expectedNext: 20, expectedOK: true},
				{s: 0, a: '-', expectedNext: 20, expectedOK: true},
				{s: 0, a: 'A', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'M', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'X', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'α', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'β', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: 40, expectedOK: true},
				{s: 0, a: '#', expectedNext: -1, expectedOK: false},
				{s: 0, a: '@', expectedNext: -1, expectedOK: false},
				{s: 0, a: '_', expectedNext: -1, expectedOK: false},
				{s: 0, a: '|', expectedNext: -1, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: -1, expectedOK: false},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, State]{
						{Key: SymbolRange{'+', '-'}, Val: 20},
						{Key: SymbolRange{'0', '9'}, Val: 0},
						{Key: SymbolRange{'A', 'Z'}, Val: 30},
						{Key: SymbolRange{'a', 'z'}, Val: 10},
						{Key: SymbolRange{'α', 'δ'}, Val: 40},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, 20},
				{0, SymbolRange{'0', '9'}, 0},
				{0, SymbolRange{'A', 'Z'}, 30},
				{0, SymbolRange{'a', 'z'}, 10},
				{0, SymbolRange{'α', 'δ'}, 40},
			},
			expectedString: `Transitions:
  0 --[+..-]--> 20
  0 --[0..9]--> 0
  0 --[A..Z]--> 30
  0 --[a..z]--> 10
  0 --[α..δ]--> 40
`,
		},
		{
			name: "CurrentEndOnLastEnd_DiffStates",
			trans: map[State][]rangeState{
				0: {
					{SymbolRange{Start: '0', End: '9'}, 0},
					{SymbolRange{Start: 'a', End: 'n'}, 10},
					{SymbolRange{Start: 'n', End: 'n'}, 11},
					{SymbolRange{Start: 'n', End: 'p'}, 11},
					{SymbolRange{Start: 'p', End: 'z'}, 12},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newDFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 11},
								{SymbolRange{Start: 'n', End: 'p'}, 11},
								{SymbolRange{Start: 'p', End: 'z'}, 120},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'n'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 11},
								{SymbolRange{Start: 'n', End: 'p'}, 11},
								{SymbolRange{Start: 'p', End: 'z'}, 12},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: 20},
				{s: 0, start: 'A', end: 'N', next: 30},
				{s: 0, start: 'N', end: 'N', next: 31},
				{s: 0, start: 'N', end: 'P', next: 31},
				{s: 0, start: 'P', end: 'Z', next: 32},
				{s: 0, start: 'α', end: 'δ', next: 40},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: 0, expectedOK: true},
				{s: 0, a: '5', expectedNext: 0, expectedOK: true},
				{s: 0, a: '9', expectedNext: 0, expectedOK: true},
				{s: 0, a: 'a', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'i', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'j', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'm', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'n', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'o', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'p', expectedNext: 12, expectedOK: true},
				{s: 0, a: 'q', expectedNext: 12, expectedOK: true},
				{s: 0, a: 'x', expectedNext: 12, expectedOK: true},
				{s: 0, a: 'z', expectedNext: 12, expectedOK: true},
				{s: 0, a: '+', expectedNext: 20, expectedOK: true},
				{s: 0, a: ',', expectedNext: 20, expectedOK: true},
				{s: 0, a: '-', expectedNext: 20, expectedOK: true},
				{s: 0, a: 'A', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'I', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'J', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'M', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'N', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'O', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'P', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'Q', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'X', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'α', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'β', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: 40, expectedOK: true},
				{s: 0, a: '#', expectedNext: -1, expectedOK: false},
				{s: 0, a: '@', expectedNext: -1, expectedOK: false},
				{s: 0, a: '_', expectedNext: -1, expectedOK: false},
				{s: 0, a: '|', expectedNext: -1, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: -1, expectedOK: false},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, State]{
						{Key: SymbolRange{'+', '-'}, Val: 20},
						{Key: SymbolRange{'0', '9'}, Val: 0},
						{Key: SymbolRange{'A', 'M'}, Val: 30},
						{Key: SymbolRange{'N', 'O'}, Val: 31},
						{Key: SymbolRange{'P', 'Z'}, Val: 32},
						{Key: SymbolRange{'a', 'm'}, Val: 10},
						{Key: SymbolRange{'n', 'o'}, Val: 11},
						{Key: SymbolRange{'p', 'z'}, Val: 12},
						{Key: SymbolRange{'α', 'δ'}, Val: 40},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, 20},
				{0, SymbolRange{'0', '9'}, 0},
				{0, SymbolRange{'A', 'M'}, 30},
				{0, SymbolRange{'N', 'O'}, 31},
				{0, SymbolRange{'P', 'Z'}, 32},
				{0, SymbolRange{'a', 'm'}, 10},
				{0, SymbolRange{'n', 'o'}, 11},
				{0, SymbolRange{'p', 'z'}, 12},
				{0, SymbolRange{'α', 'δ'}, 40},
			},
			expectedString: `Transitions:
  0 --[+..-]--> 20
  0 --[0..9]--> 0
  0 --[A..M]--> 30
  0 --[N..O]--> 31
  0 --[P..Z]--> 32
  0 --[a..m]--> 10
  0 --[n..o]--> 11
  0 --[p..z]--> 12
  0 --[α..δ]--> 40
`,
		},
		{
			name: "CurrentEndBeforeLastEnd_SameStates",
			trans: map[State][]rangeState{
				0: {
					{SymbolRange{Start: '0', End: '9'}, 0},
					{SymbolRange{Start: 'a', End: 'w'}, 10},
					{SymbolRange{Start: 'i', End: 'm'}, 10},
					{SymbolRange{Start: 'j', End: 'w'}, 10},
					{SymbolRange{Start: 'k', End: 'z'}, 10},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newDFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'w'}, 10},
								{SymbolRange{Start: 'i', End: 'm'}, 10},
								{SymbolRange{Start: 'j', End: 'w'}, 10},
								{SymbolRange{Start: 'k', End: 'z'}, 100},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'w'}, 10},
								{SymbolRange{Start: 'i', End: 'm'}, 10},
								{SymbolRange{Start: 'j', End: 'w'}, 10},
								{SymbolRange{Start: 'k', End: 'z'}, 10},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: 20},
				{s: 0, start: 'A', end: 'W', next: 30},
				{s: 0, start: 'I', end: 'M', next: 30},
				{s: 0, start: 'J', end: 'W', next: 30},
				{s: 0, start: 'K', end: 'Z', next: 30},
				{s: 0, start: 'α', end: 'δ', next: 40},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: 0, expectedOK: true},
				{s: 0, a: '5', expectedNext: 0, expectedOK: true},
				{s: 0, a: '9', expectedNext: 0, expectedOK: true},
				{s: 0, a: 'a', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'm', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'x', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'z', expectedNext: 10, expectedOK: true},
				{s: 0, a: '+', expectedNext: 20, expectedOK: true},
				{s: 0, a: ',', expectedNext: 20, expectedOK: true},
				{s: 0, a: '-', expectedNext: 20, expectedOK: true},
				{s: 0, a: 'A', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'M', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'X', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'α', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'β', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: 40, expectedOK: true},
				{s: 0, a: '#', expectedNext: -1, expectedOK: false},
				{s: 0, a: '@', expectedNext: -1, expectedOK: false},
				{s: 0, a: '_', expectedNext: -1, expectedOK: false},
				{s: 0, a: '|', expectedNext: -1, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: -1, expectedOK: false},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, State]{
						{Key: SymbolRange{'+', '-'}, Val: 20},
						{Key: SymbolRange{'0', '9'}, Val: 0},
						{Key: SymbolRange{'A', 'Z'}, Val: 30},
						{Key: SymbolRange{'a', 'z'}, Val: 10},
						{Key: SymbolRange{'α', 'δ'}, Val: 40},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, 20},
				{0, SymbolRange{'0', '9'}, 0},
				{0, SymbolRange{'A', 'Z'}, 30},
				{0, SymbolRange{'a', 'z'}, 10},
				{0, SymbolRange{'α', 'δ'}, 40},
			},
			expectedString: `Transitions:
  0 --[+..-]--> 20
  0 --[0..9]--> 0
  0 --[A..Z]--> 30
  0 --[a..z]--> 10
  0 --[α..δ]--> 40
`,
		},
		{
			name: "CurrentEndBeforeLastEnd_DiffStates",
			trans: map[State][]rangeState{
				0: {
					{SymbolRange{Start: '0', End: '9'}, 0},
					{SymbolRange{Start: 'a', End: 'w'}, 10},
					{SymbolRange{Start: 'i', End: 'm'}, 11},
					{SymbolRange{Start: 'r', End: 'w'}, 11},
					{SymbolRange{Start: 'v', End: 'z'}, 12},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newDFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'w'}, 10},
								{SymbolRange{Start: 'i', End: 'm'}, 11},
								{SymbolRange{Start: 'r', End: 'w'}, 11},
								{SymbolRange{Start: 'v', End: 'z'}, 120},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'w'}, 10},
								{SymbolRange{Start: 'i', End: 'm'}, 11},
								{SymbolRange{Start: 'r', End: 'w'}, 11},
								{SymbolRange{Start: 'v', End: 'z'}, 12},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: 20},
				{s: 0, start: 'A', end: 'W', next: 30},
				{s: 0, start: 'I', end: 'M', next: 31},
				{s: 0, start: 'R', end: 'W', next: 31},
				{s: 0, start: 'V', end: 'Z', next: 32},
				{s: 0, start: 'α', end: 'δ', next: 40},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: 0, expectedOK: true},
				{s: 0, a: '5', expectedNext: 0, expectedOK: true},
				{s: 0, a: '9', expectedNext: 0, expectedOK: true},
				{s: 0, a: 'a', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'e', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'h', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'i', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'j', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'm', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'n', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'p', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'q', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'r', expectedNext: 11, expectedOK: true},
				{s: 0, a: 's', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'u', expectedNext: 11, expectedOK: true},
				{s: 0, a: 'v', expectedNext: 12, expectedOK: true},
				{s: 0, a: 'x', expectedNext: 12, expectedOK: true},
				{s: 0, a: 'z', expectedNext: 12, expectedOK: true},
				{s: 0, a: '+', expectedNext: 20, expectedOK: true},
				{s: 0, a: ',', expectedNext: 20, expectedOK: true},
				{s: 0, a: '-', expectedNext: 20, expectedOK: true},
				{s: 0, a: 'A', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'E', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'H', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'I', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'J', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'M', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'N', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'P', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'Q', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'R', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'S', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'U', expectedNext: 31, expectedOK: true},
				{s: 0, a: 'V', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'X', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: 32, expectedOK: true},
				{s: 0, a: 'α', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'β', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: 40, expectedOK: true},
				{s: 0, a: '#', expectedNext: -1, expectedOK: false},
				{s: 0, a: '@', expectedNext: -1, expectedOK: false},
				{s: 0, a: '_', expectedNext: -1, expectedOK: false},
				{s: 0, a: '|', expectedNext: -1, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: -1, expectedOK: false},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, State]{
						{Key: SymbolRange{'+', '-'}, Val: 20},
						{Key: SymbolRange{'0', '9'}, Val: 0},
						{Key: SymbolRange{'A', 'H'}, Val: 30},
						{Key: SymbolRange{'I', 'M'}, Val: 31},
						{Key: SymbolRange{'N', 'Q'}, Val: 30},
						{Key: SymbolRange{'R', 'U'}, Val: 31},
						{Key: SymbolRange{'V', 'Z'}, Val: 32},
						{Key: SymbolRange{'a', 'h'}, Val: 10},
						{Key: SymbolRange{'i', 'm'}, Val: 11},
						{Key: SymbolRange{'n', 'q'}, Val: 10},
						{Key: SymbolRange{'r', 'u'}, Val: 11},
						{Key: SymbolRange{'v', 'z'}, Val: 12},
						{Key: SymbolRange{'α', 'δ'}, Val: 40},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, 20},
				{0, SymbolRange{'0', '9'}, 0},
				{0, SymbolRange{'A', 'H'}, 30},
				{0, SymbolRange{'I', 'M'}, 31},
				{0, SymbolRange{'N', 'Q'}, 30},
				{0, SymbolRange{'R', 'U'}, 31},
				{0, SymbolRange{'V', 'Z'}, 32},
				{0, SymbolRange{'a', 'h'}, 10},
				{0, SymbolRange{'i', 'm'}, 11},
				{0, SymbolRange{'n', 'q'}, 10},
				{0, SymbolRange{'r', 'u'}, 11},
				{0, SymbolRange{'v', 'z'}, 12},
				{0, SymbolRange{'α', 'δ'}, 40},
			},
			expectedString: `Transitions:
  0 --[+..-]--> 20
  0 --[0..9]--> 0
  0 --[A..H]--> 30
  0 --[I..M]--> 31
  0 --[N..Q]--> 30
  0 --[R..U]--> 31
  0 --[V..Z]--> 32
  0 --[a..h]--> 10
  0 --[i..m]--> 11
  0 --[n..q]--> 10
  0 --[r..u]--> 11
  0 --[v..z]--> 12
  0 --[α..δ]--> 40
`,
		},
		{
			name: "CurrentEndAdjacentToLastEnd_SameStates",
			trans: map[State][]rangeState{
				0: {
					{SymbolRange{Start: '0', End: '9'}, 0},
					{SymbolRange{Start: 'a', End: 'm'}, 10},
					{SymbolRange{Start: 'n', End: 'n'}, 10},
					{SymbolRange{Start: 'o', End: 'z'}, 10},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newDFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'm'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 10},
								{SymbolRange{Start: 'o', End: 'z'}, 100},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newDFATransitionTable(
						map[State][]rangeState{
							0: {
								{SymbolRange{Start: '0', End: '9'}, 0},
								{SymbolRange{Start: 'a', End: 'm'}, 10},
								{SymbolRange{Start: 'n', End: 'n'}, 10},
								{SymbolRange{Start: 'o', End: 'z'}, 10},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: 20},
				{s: 0, start: 'A', end: 'M', next: 30},
				{s: 0, start: 'N', end: 'N', next: 30},
				{s: 0, start: 'O', end: 'Z', next: 30},
				{s: 0, start: 'α', end: 'δ', next: 40},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: 0, expectedOK: true},
				{s: 0, a: '5', expectedNext: 0, expectedOK: true},
				{s: 0, a: '9', expectedNext: 0, expectedOK: true},
				{s: 0, a: 'a', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'i', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'm', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'n', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'o', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'x', expectedNext: 10, expectedOK: true},
				{s: 0, a: 'z', expectedNext: 10, expectedOK: true},
				{s: 0, a: '+', expectedNext: 20, expectedOK: true},
				{s: 0, a: ',', expectedNext: 20, expectedOK: true},
				{s: 0, a: '-', expectedNext: 20, expectedOK: true},
				{s: 0, a: 'A', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'I', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'M', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'N', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'O', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'X', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: 30, expectedOK: true},
				{s: 0, a: 'α', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'β', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: 40, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: 40, expectedOK: true},
				{s: 0, a: '#', expectedNext: -1, expectedOK: false},
				{s: 0, a: '@', expectedNext: -1, expectedOK: false},
				{s: 0, a: '_', expectedNext: -1, expectedOK: false},
				{s: 0, a: '|', expectedNext: -1, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: -1, expectedOK: false},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, State]{
						{Key: SymbolRange{'+', '-'}, Val: 20},
						{Key: SymbolRange{'0', '9'}, Val: 0},
						{Key: SymbolRange{'A', 'Z'}, Val: 30},
						{Key: SymbolRange{'a', 'z'}, Val: 10},
						{Key: SymbolRange{'α', 'δ'}, Val: 40},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, 20},
				{0, SymbolRange{'0', '9'}, 0},
				{0, SymbolRange{'A', 'Z'}, 30},
				{0, SymbolRange{'a', 'z'}, 10},
				{0, SymbolRange{'α', 'δ'}, 40},
			},
			expectedString: `Transitions:
  0 --[+..-]--> 20
  0 --[0..9]--> 0
  0 --[A..Z]--> 30
  0 --[a..z]--> 10
  0 --[α..δ]--> 40
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newDFATransitionTable(tc.trans)

			t.Run("Clone", func(t *testing.T) {
				clone := m.Clone()
				assert.True(t, clone.Equal(m))
			})

			t.Run("Equal", func(t *testing.T) {
				for _, tc := range tc.equalTests {
					assert.Equal(t, tc.expectedEqual, m.Equal(tc.rhs))
				}
			})

			t.Run("Add", func(t *testing.T) {
				for _, tc := range tc.addTests {
					m.Add(tc.s, tc.start, tc.end, tc.next)
				}
			})

			t.Run("Next", func(t *testing.T) {
				for _, tc := range tc.nextTests {
					state, ok := m.Next(tc.s, tc.a)
					assert.Equal(t, tc.expectedOK, ok)
					assert.Equal(t, tc.expectedNext, state, "From state %s on symbol %q expected %d, but got %d", tc.s, tc.a, tc.expectedNext, state)
				}
			})

			t.Run("From", func(t *testing.T) {
				for _, tc := range tc.fromTests {
					from := generic.Collect2(m.From(tc.s))
					assert.Equal(t, tc.expectedFrom, from)
				}
			})

			t.Run("All", func(t *testing.T) {
				all := []transition{}
				for s, pairs := range m.All() {
					for r, next := range pairs {
						all = append(all, transition{s, r, next})
					}
				}

				assert.Equal(t, tc.expectedAll, all)
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, m.String())
			})
		})
	}
}
