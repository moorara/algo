package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/range/disc"
)

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
		{ // (a+|b+)
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
			expectedNFA: &NFA{
				start: 0,
				final: NewStates(2, 4),
				classes: disc.NewRangeMap(eqClassID, nil, []disc.RangeValue[Symbol, classID]{
					{Range: disc.Range[Symbol]{Lo: E, Hi: E}, Value: 0},
					{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'a'}, Value: 1},
					{Range: disc.Range[Symbol]{Lo: 'b', Hi: 'b'}, Value: 2},
				}),
			},
		},
		{ // ([A-Za-z_][0-9A-Za-z_]*)|[0-9]+|(0x[0-9A-Fa-f]+)|[ \t\n]+|[+\-*/=]
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
			expectedNFA: &NFA{
				start: 0,
				final: NewStates(1, 2, 5),
				classes: disc.NewRangeMap(eqClassID, nil, []disc.RangeValue[Symbol, classID]{
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
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := new(NFABuilder).SetStart(tc.start).SetFinal(tc.final...)

			for _, tr := range tc.trans {
				b.AddTransition(tr.s, tr.start, tr.end, tr.next)
			}

			t.Run("Build", func(t *testing.T) {
				assert.True(t, b.Build().Equal(tc.expectedNFA))
			})
		})
	}
}
