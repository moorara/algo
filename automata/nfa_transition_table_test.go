package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeStates(t *testing.T) {
	tests := []struct {
		name           string
		rs             rangeStates
		expectedString string
		equal          rangeStates
		expectedEqual  bool
	}{
		{
			name: "Equal",
			rs: rangeStates{
				SymbolRange{Start: '0', End: '9'},
				NewStates(1, 2),
			},
			expectedString: "[0..9] → {1, 2}",
			equal: rangeStates{
				SymbolRange{Start: '0', End: '9'},
				NewStates(1, 2),
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			rs: rangeStates{
				SymbolRange{Start: '0', End: '9'},
				NewStates(1, 2),
			},
			expectedString: "[0..9] → {1, 2}",
			equal: rangeStates{
				SymbolRange{Start: '0', End: '9'},
				NewStates(3, 4),
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

func TestNFATransitionTable(t *testing.T) {
	type equalTest struct {
		rhs           *nfaTransitionTable
		expectedEqual bool
	}

	type addTest struct {
		s     State
		start Symbol
		end   Symbol
		next  []State
	}

	type nextTest struct {
		s            State
		a            Symbol
		expectedNext []State
		expectedOK   bool
	}

	type nextOnRangeTest struct {
		s             State
		r             SymbolRange
		expectedPairs []rangeStates
		expectedOK    bool
	}

	type fromTest struct {
		s            State
		expectedFrom []generic.KeyValue[SymbolRange, []State]
	}

	type transition struct {
		s    State
		r    SymbolRange
		next []State
	}

	tests := []struct {
		name                 string
		trans                map[State][]rangeStates
		equalTests           []equalTest
		addTests             []addTest
		nextTests            []nextTest
		nextOnRangeTests     []nextOnRangeTest
		fromTests            []fromTest
		expectedAll          []transition
		expectedSymbolRanges []SymbolRange
		expectedString       string
	}{
		{
			name: "CurrentEndOnLastEnd_SameStates",
			trans: map[State][]rangeStates{
				0: {
					{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
					{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
					{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
					{SymbolRange{Start: 'n', End: 'z'}, NewStates(10, 11)},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newNFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'z'}, NewStates(10, 110)},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'z'}, NewStates(10, 11)},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: []State{20, 21}},
				{s: 0, start: 'A', end: 'N', next: []State{30, 31}},
				{s: 0, start: 'N', end: 'N', next: []State{30, 31}},
				{s: 0, start: 'N', end: 'Z', next: []State{30, 31}},
				{s: 0, start: 'α', end: 'δ', next: []State{40, 41}},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '5', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '9', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: 'a', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'm', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'x', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'z', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: '+', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: ',', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: '-', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: 'A', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'M', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'X', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'α', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'β', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: '#', expectedNext: nil, expectedOK: false},
				{s: 0, a: '@', expectedNext: nil, expectedOK: false},
				{s: 0, a: '_', expectedNext: nil, expectedOK: false},
				{s: 0, a: '|', expectedNext: nil, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: nil, expectedOK: false},
			},
			nextOnRangeTests: []nextOnRangeTest{
				{
					s:             0,
					r:             SymbolRange{'ε', 'ω'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s:             1,
					r:             SymbolRange{'A', 'Z'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s: 0,
					r: SymbolRange{'A', 'z'},
					expectedPairs: []rangeStates{
						{SymbolRange{'A', 'Z'}, NewStates(30, 31)},
						{SymbolRange{'a', 'z'}, NewStates(10, 11)},
					},
					expectedOK: true,
				},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, []State]{
						{Key: SymbolRange{'+', '-'}, Val: []State{20, 21}},
						{Key: SymbolRange{'0', '9'}, Val: []State{0, 1}},
						{Key: SymbolRange{'A', 'Z'}, Val: []State{30, 31}},
						{Key: SymbolRange{'a', 'z'}, Val: []State{10, 11}},
						{Key: SymbolRange{'α', 'δ'}, Val: []State{40, 41}},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, []State{20, 21}},
				{0, SymbolRange{'0', '9'}, []State{0, 1}},
				{0, SymbolRange{'A', 'Z'}, []State{30, 31}},
				{0, SymbolRange{'a', 'z'}, []State{10, 11}},
				{0, SymbolRange{'α', 'δ'}, []State{40, 41}},
			},
			expectedSymbolRanges: []SymbolRange{
				{'+', '-'},
				{'0', '9'},
				{'A', 'Z'},
				{'a', 'z'},
				{'α', 'δ'},
			},
			expectedString: `Transitions:
  0 --[+..-]--> {20, 21}
  0 --[0..9]--> {0, 1}
  0 --[A..Z]--> {30, 31}
  0 --[a..z]--> {10, 11}
  0 --[α..δ]--> {40, 41}
`,
		},
		{
			name: "CurrentEndOnLastEnd_DiffStates",
			trans: map[State][]rangeStates{
				0: {
					{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
					{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
					{SymbolRange{Start: 'n', End: 'n'}, NewStates(12, 13)},
					{SymbolRange{Start: 'n', End: 'p'}, NewStates(12, 13)},
					{SymbolRange{Start: 'p', End: 'z'}, NewStates(14, 15)},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newNFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'n'}, NewStates(12, 13)},
								{SymbolRange{Start: 'n', End: 'p'}, NewStates(12, 13)},
								{SymbolRange{Start: 'p', End: 'z'}, NewStates(14, 150)},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'n'}, NewStates(10, 11)},
								{SymbolRange{Start: 'n', End: 'n'}, NewStates(12, 13)},
								{SymbolRange{Start: 'n', End: 'p'}, NewStates(12, 13)},
								{SymbolRange{Start: 'p', End: 'z'}, NewStates(14, 15)},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: []State{20, 21}},
				{s: 0, start: 'A', end: 'N', next: []State{30, 31}},
				{s: 0, start: 'N', end: 'N', next: []State{32, 33}},
				{s: 0, start: 'N', end: 'P', next: []State{32, 33}},
				{s: 0, start: 'P', end: 'Z', next: []State{34, 35}},
				{s: 0, start: 'α', end: 'δ', next: []State{40, 41}},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '5', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '9', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: 'a', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'i', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'j', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'm', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'n', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'o', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'p', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: 'q', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: 'x', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: 'z', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: '+', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: ',', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: '-', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: 'A', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'I', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'J', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'M', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'N', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'O', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'P', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'Q', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'X', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'α', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'β', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: '#', expectedNext: nil, expectedOK: false},
				{s: 0, a: '@', expectedNext: nil, expectedOK: false},
				{s: 0, a: '_', expectedNext: nil, expectedOK: false},
				{s: 0, a: '|', expectedNext: nil, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: nil, expectedOK: false},
			},
			nextOnRangeTests: []nextOnRangeTest{
				{
					s:             0,
					r:             SymbolRange{'ε', 'ω'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s:             1,
					r:             SymbolRange{'A', 'Z'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s: 0,
					r: SymbolRange{'A', 'z'},
					expectedPairs: []rangeStates{
						{SymbolRange{'A', 'M'}, NewStates(30, 31)},
						{SymbolRange{'N', 'O'}, NewStates(32, 33)},
						{SymbolRange{'P', 'Z'}, NewStates(34, 35)},
						{SymbolRange{'a', 'm'}, NewStates(10, 11)},
						{SymbolRange{'n', 'o'}, NewStates(12, 13)},
						{SymbolRange{'p', 'z'}, NewStates(14, 15)},
					},
					expectedOK: true,
				},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, []State]{
						{Key: SymbolRange{'+', '-'}, Val: []State{20, 21}},
						{Key: SymbolRange{'0', '9'}, Val: []State{0, 1}},
						{Key: SymbolRange{'A', 'M'}, Val: []State{30, 31}},
						{Key: SymbolRange{'N', 'O'}, Val: []State{32, 33}},
						{Key: SymbolRange{'P', 'Z'}, Val: []State{34, 35}},
						{Key: SymbolRange{'a', 'm'}, Val: []State{10, 11}},
						{Key: SymbolRange{'n', 'o'}, Val: []State{12, 13}},
						{Key: SymbolRange{'p', 'z'}, Val: []State{14, 15}},
						{Key: SymbolRange{'α', 'δ'}, Val: []State{40, 41}},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, []State{20, 21}},
				{0, SymbolRange{'0', '9'}, []State{0, 1}},
				{0, SymbolRange{'A', 'M'}, []State{30, 31}},
				{0, SymbolRange{'N', 'O'}, []State{32, 33}},
				{0, SymbolRange{'P', 'Z'}, []State{34, 35}},
				{0, SymbolRange{'a', 'm'}, []State{10, 11}},
				{0, SymbolRange{'n', 'o'}, []State{12, 13}},
				{0, SymbolRange{'p', 'z'}, []State{14, 15}},
				{0, SymbolRange{'α', 'δ'}, []State{40, 41}},
			},
			expectedSymbolRanges: []SymbolRange{
				{'+', '-'},
				{'0', '9'},
				{'A', 'Z'},
				{'a', 'z'},
				{'α', 'δ'},
			},
			expectedString: `Transitions:
  0 --[+..-]--> {20, 21}
  0 --[0..9]--> {0, 1}
  0 --[A..M]--> {30, 31}
  0 --[N..O]--> {32, 33}
  0 --[P..Z]--> {34, 35}
  0 --[a..m]--> {10, 11}
  0 --[n..o]--> {12, 13}
  0 --[p..z]--> {14, 15}
  0 --[α..δ]--> {40, 41}
`,
		},
		{
			name: "CurrentEndBeforeLastEnd_SameStates",
			trans: map[State][]rangeStates{
				0: {
					{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
					{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
					{SymbolRange{Start: 'i', End: 'm'}, NewStates(10, 11)},
					{SymbolRange{Start: 'j', End: 'w'}, NewStates(10, 11)},
					{SymbolRange{Start: 'k', End: 'z'}, NewStates(10, 11)},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newNFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'i', End: 'm'}, NewStates(10, 11)},
								{SymbolRange{Start: 'j', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'k', End: 'z'}, NewStates(10, 110)},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'i', End: 'm'}, NewStates(10, 11)},
								{SymbolRange{Start: 'j', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'k', End: 'z'}, NewStates(10, 11)},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: []State{20, 21}},
				{s: 0, start: 'A', end: 'W', next: []State{30, 31}},
				{s: 0, start: 'I', end: 'M', next: []State{30, 31}},
				{s: 0, start: 'J', end: 'W', next: []State{30, 31}},
				{s: 0, start: 'K', end: 'Z', next: []State{30, 31}},
				{s: 0, start: 'α', end: 'δ', next: []State{40, 41}},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '5', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '9', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: 'a', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'm', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'x', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'z', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: '+', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: ',', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: '-', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: 'A', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'M', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'X', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'α', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'β', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: '#', expectedNext: nil, expectedOK: false},
				{s: 0, a: '@', expectedNext: nil, expectedOK: false},
				{s: 0, a: '_', expectedNext: nil, expectedOK: false},
				{s: 0, a: '|', expectedNext: nil, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: nil, expectedOK: false},
			},
			nextOnRangeTests: []nextOnRangeTest{
				{
					s:             0,
					r:             SymbolRange{'ε', 'ω'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s:             1,
					r:             SymbolRange{'A', 'Z'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s: 0,
					r: SymbolRange{'A', 'z'},
					expectedPairs: []rangeStates{
						{SymbolRange{'A', 'Z'}, NewStates(30, 31)},
						{SymbolRange{'a', 'z'}, NewStates(10, 11)},
					},
					expectedOK: true,
				},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, []State]{
						{Key: SymbolRange{'+', '-'}, Val: []State{20, 21}},
						{Key: SymbolRange{'0', '9'}, Val: []State{0, 1}},
						{Key: SymbolRange{'A', 'Z'}, Val: []State{30, 31}},
						{Key: SymbolRange{'a', 'z'}, Val: []State{10, 11}},
						{Key: SymbolRange{'α', 'δ'}, Val: []State{40, 41}},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, []State{20, 21}},
				{0, SymbolRange{'0', '9'}, []State{0, 1}},
				{0, SymbolRange{'A', 'Z'}, []State{30, 31}},
				{0, SymbolRange{'a', 'z'}, []State{10, 11}},
				{0, SymbolRange{'α', 'δ'}, []State{40, 41}},
			},
			expectedSymbolRanges: []SymbolRange{
				{'+', '-'},
				{'0', '9'},
				{'A', 'Z'},
				{'a', 'z'},
				{'α', 'δ'},
			},
			expectedString: `Transitions:
  0 --[+..-]--> {20, 21}
  0 --[0..9]--> {0, 1}
  0 --[A..Z]--> {30, 31}
  0 --[a..z]--> {10, 11}
  0 --[α..δ]--> {40, 41}
`,
		},
		{
			name: "CurrentEndBeforeLastEnd_DiffStates",
			trans: map[State][]rangeStates{
				0: {
					{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
					{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
					{SymbolRange{Start: 'i', End: 'm'}, NewStates(12, 13)},
					{SymbolRange{Start: 'r', End: 'w'}, NewStates(12, 13)},
					{SymbolRange{Start: 'v', End: 'z'}, NewStates(14, 15)},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newNFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'i', End: 'm'}, NewStates(12, 13)},
								{SymbolRange{Start: 'r', End: 'w'}, NewStates(12, 13)},
								{SymbolRange{Start: 'v', End: 'z'}, NewStates(14, 150)},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								{SymbolRange{Start: 'a', End: 'w'}, NewStates(10, 11)},
								{SymbolRange{Start: 'i', End: 'm'}, NewStates(12, 13)},
								{SymbolRange{Start: 'r', End: 'w'}, NewStates(12, 13)},
								{SymbolRange{Start: 'v', End: 'z'}, NewStates(14, 15)},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: []State{20, 21}},
				{s: 0, start: 'A', end: 'W', next: []State{30, 31}},
				{s: 0, start: 'I', end: 'M', next: []State{32, 33}},
				{s: 0, start: 'R', end: 'W', next: []State{32, 33}},
				{s: 0, start: 'V', end: 'Z', next: []State{34, 35}},
				{s: 0, start: 'α', end: 'δ', next: []State{40, 41}},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '5', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '9', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: 'a', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'e', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'h', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'i', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'j', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'm', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'n', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'p', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'q', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'r', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 's', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'u', expectedNext: []State{12, 13}, expectedOK: true},
				{s: 0, a: 'v', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: 'x', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: 'z', expectedNext: []State{14, 15}, expectedOK: true},
				{s: 0, a: '+', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: ',', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: '-', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: 'A', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'E', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'H', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'I', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'J', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'M', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'N', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'P', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'Q', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'R', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'S', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'U', expectedNext: []State{32, 33}, expectedOK: true},
				{s: 0, a: 'V', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'X', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: []State{34, 35}, expectedOK: true},
				{s: 0, a: 'α', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'β', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: '#', expectedNext: nil, expectedOK: false},
				{s: 0, a: '@', expectedNext: nil, expectedOK: false},
				{s: 0, a: '_', expectedNext: nil, expectedOK: false},
				{s: 0, a: '|', expectedNext: nil, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: nil, expectedOK: false},
			},
			nextOnRangeTests: []nextOnRangeTest{
				{
					s:             0,
					r:             SymbolRange{'ε', 'ω'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s:             1,
					r:             SymbolRange{'A', 'Z'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s: 0,
					r: SymbolRange{'A', 'z'},
					expectedPairs: []rangeStates{
						{SymbolRange{'A', 'H'}, NewStates(30, 31)},
						{SymbolRange{'I', 'M'}, NewStates(32, 33)},
						{SymbolRange{'N', 'Q'}, NewStates(30, 31)},
						{SymbolRange{'R', 'U'}, NewStates(32, 33)},
						{SymbolRange{'V', 'Z'}, NewStates(34, 35)},
						{SymbolRange{'a', 'h'}, NewStates(10, 11)},
						{SymbolRange{'i', 'm'}, NewStates(12, 13)},
						{SymbolRange{'n', 'q'}, NewStates(10, 11)},
						{SymbolRange{'r', 'u'}, NewStates(12, 13)},
						{SymbolRange{'v', 'z'}, NewStates(14, 15)},
					},
					expectedOK: true,
				},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, []State]{
						{Key: SymbolRange{'+', '-'}, Val: []State{20, 21}},
						{Key: SymbolRange{'0', '9'}, Val: []State{0, 1}},
						{Key: SymbolRange{'A', 'H'}, Val: []State{30, 31}},
						{Key: SymbolRange{'I', 'M'}, Val: []State{32, 33}},
						{Key: SymbolRange{'N', 'Q'}, Val: []State{30, 31}},
						{Key: SymbolRange{'R', 'U'}, Val: []State{32, 33}},
						{Key: SymbolRange{'V', 'Z'}, Val: []State{34, 35}},
						{Key: SymbolRange{'a', 'h'}, Val: []State{10, 11}},
						{Key: SymbolRange{'i', 'm'}, Val: []State{12, 13}},
						{Key: SymbolRange{'n', 'q'}, Val: []State{10, 11}},
						{Key: SymbolRange{'r', 'u'}, Val: []State{12, 13}},
						{Key: SymbolRange{'v', 'z'}, Val: []State{14, 15}},
						{Key: SymbolRange{'α', 'δ'}, Val: []State{40, 41}},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, []State{20, 21}},
				{0, SymbolRange{'0', '9'}, []State{0, 1}},
				{0, SymbolRange{'A', 'H'}, []State{30, 31}},
				{0, SymbolRange{'I', 'M'}, []State{32, 33}},
				{0, SymbolRange{'N', 'Q'}, []State{30, 31}},
				{0, SymbolRange{'R', 'U'}, []State{32, 33}},
				{0, SymbolRange{'V', 'Z'}, []State{34, 35}},
				{0, SymbolRange{'a', 'h'}, []State{10, 11}},
				{0, SymbolRange{'i', 'm'}, []State{12, 13}},
				{0, SymbolRange{'n', 'q'}, []State{10, 11}},
				{0, SymbolRange{'r', 'u'}, []State{12, 13}},
				{0, SymbolRange{'v', 'z'}, []State{14, 15}},
				{0, SymbolRange{'α', 'δ'}, []State{40, 41}},
			},
			expectedSymbolRanges: []SymbolRange{
				{'+', '-'},
				{'0', '9'},
				{'A', 'Z'},
				{'a', 'z'},
				{'α', 'δ'},
			},
			expectedString: `Transitions:
  0 --[+..-]--> {20, 21}
  0 --[0..9]--> {0, 1}
  0 --[A..H]--> {30, 31}
  0 --[I..M]--> {32, 33}
  0 --[N..Q]--> {30, 31}
  0 --[R..U]--> {32, 33}
  0 --[V..Z]--> {34, 35}
  0 --[a..h]--> {10, 11}
  0 --[i..m]--> {12, 13}
  0 --[n..q]--> {10, 11}
  0 --[r..u]--> {12, 13}
  0 --[v..z]--> {14, 15}
  0 --[α..δ]--> {40, 41}
`,
		},
		{
			name: "CurrentEndAdjacentToLastEnd_SameStates",
			trans: map[State][]rangeStates{
				0: {
					{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
					{SymbolRange{Start: 'a', End: 'm'}, NewStates(10, 11)},
					{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
					{SymbolRange{Start: 'o', End: 'z'}, NewStates(10, 11)},
				},
			},
			equalTests: []equalTest{
				{
					rhs:           newNFATransitionTable(nil),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								rangeStates{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								rangeStates{SymbolRange{Start: 'a', End: 'm'}, NewStates(10, 11)},
								rangeStates{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
								rangeStates{SymbolRange{Start: 'o', End: 'z'}, NewStates(10, 110)},
							},
						},
					),
					expectedEqual: false,
				},
				{
					rhs: newNFATransitionTable(
						map[State][]rangeStates{
							0: {
								rangeStates{SymbolRange{Start: '0', End: '9'}, NewStates(0, 1)},
								rangeStates{SymbolRange{Start: 'a', End: 'm'}, NewStates(10, 11)},
								rangeStates{SymbolRange{Start: 'n', End: 'n'}, NewStates(10, 11)},
								rangeStates{SymbolRange{Start: 'o', End: 'z'}, NewStates(10, 11)},
							},
						},
					),
					expectedEqual: true,
				},
			},
			addTests: []addTest{
				{s: 0, start: '+', end: '-', next: []State{20, 21}},
				{s: 0, start: 'A', end: 'M', next: []State{30, 31}},
				{s: 0, start: 'N', end: 'N', next: []State{30, 31}},
				{s: 0, start: 'O', end: 'Z', next: []State{30, 31}},
				{s: 0, start: 'α', end: 'δ', next: []State{40, 41}},
			},
			nextTests: []nextTest{
				{s: 0, a: '0', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '5', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: '9', expectedNext: []State{0, 1}, expectedOK: true},
				{s: 0, a: 'a', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'i', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'm', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'n', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'o', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'x', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: 'z', expectedNext: []State{10, 11}, expectedOK: true},
				{s: 0, a: '+', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: ',', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: '-', expectedNext: []State{20, 21}, expectedOK: true},
				{s: 0, a: 'A', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'I', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'M', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'N', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'O', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'X', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'Z', expectedNext: []State{30, 31}, expectedOK: true},
				{s: 0, a: 'α', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'β', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'γ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: 'δ', expectedNext: []State{40, 41}, expectedOK: true},
				{s: 0, a: '#', expectedNext: nil, expectedOK: false},
				{s: 0, a: '@', expectedNext: nil, expectedOK: false},
				{s: 0, a: '_', expectedNext: nil, expectedOK: false},
				{s: 0, a: '|', expectedNext: nil, expectedOK: false},
				{s: 0, a: 'ω', expectedNext: nil, expectedOK: false},
			},
			nextOnRangeTests: []nextOnRangeTest{
				{
					s:             0,
					r:             SymbolRange{'ε', 'ω'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s:             1,
					r:             SymbolRange{'A', 'Z'},
					expectedPairs: nil,
					expectedOK:    false,
				},
				{
					s: 0,
					r: SymbolRange{'A', 'z'},
					expectedPairs: []rangeStates{
						{SymbolRange{'A', 'Z'}, NewStates(30, 31)},
						{SymbolRange{'a', 'z'}, NewStates(10, 11)},
					},
					expectedOK: true,
				},
			},
			fromTests: []fromTest{
				{
					s: 0,
					expectedFrom: []generic.KeyValue[SymbolRange, []State]{
						{Key: SymbolRange{'+', '-'}, Val: []State{20, 21}},
						{Key: SymbolRange{'0', '9'}, Val: []State{0, 1}},
						{Key: SymbolRange{'A', 'Z'}, Val: []State{30, 31}},
						{Key: SymbolRange{'a', 'z'}, Val: []State{10, 11}},
						{Key: SymbolRange{'α', 'δ'}, Val: []State{40, 41}},
					},
				},
			},
			expectedAll: []transition{
				{0, SymbolRange{'+', '-'}, []State{20, 21}},
				{0, SymbolRange{'0', '9'}, []State{0, 1}},
				{0, SymbolRange{'A', 'Z'}, []State{30, 31}},
				{0, SymbolRange{'a', 'z'}, []State{10, 11}},
				{0, SymbolRange{'α', 'δ'}, []State{40, 41}},
			},
			expectedSymbolRanges: []SymbolRange{
				{'+', '-'},
				{'0', '9'},
				{'A', 'Z'},
				{'a', 'z'},
				{'α', 'δ'},
			},
			expectedString: `Transitions:
  0 --[+..-]--> {20, 21}
  0 --[0..9]--> {0, 1}
  0 --[A..Z]--> {30, 31}
  0 --[a..z]--> {10, 11}
  0 --[α..δ]--> {40, 41}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newNFATransitionTable(tc.trans)

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
					states, ok := m.Next(tc.s, tc.a)

					assert.Equal(t, tc.expectedOK, ok)
					assert.Equal(t, tc.expectedNext, states, "From state %s on symbol %q expected %v, but got %v", tc.s, tc.a, tc.expectedNext, states)
				}
			})

			t.Run("NextOnRange", func(t *testing.T) {
				for _, tc := range tc.nextOnRangeTests {
					pairs, ok := m.NextOnRange(tc.s, tc.r)

					assert.Equal(t, tc.expectedOK, ok)
					assert.Len(t, pairs, len(tc.expectedPairs))
					for i, pair := range pairs {
						assert.True(t, pair.Equal(tc.expectedPairs[i]))
					}
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

			t.Run("SymbolRanges", func(t *testing.T) {
				symbols := m.SymbolRanges()
				assert.Equal(t, tc.expectedSymbolRanges, symbols)
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, m.String())
			})
		})
	}
}
