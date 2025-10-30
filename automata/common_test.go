package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
)

func TestFormatRange(t *testing.T) {
	tests := []struct {
		name           string
		r              disc.Range[Symbol]
		expectedString string
	}{
		{
			name:           "Empty",
			r:              disc.Range[Symbol]{Lo: E, Hi: E},
			expectedString: "[ε..ε]",
		},
		{
			name:           "Zero",
			r:              disc.Range[Symbol]{Lo: 0, Hi: 0},
			expectedString: "[NUL..NUL]",
		},
		{
			name:           "HorizontalTab",
			r:              disc.Range[Symbol]{Lo: '\t', Hi: '\t'},
			expectedString: "[\\t..\\t]",
		},
		{
			name:           "Newline",
			r:              disc.Range[Symbol]{Lo: '\n', Hi: '\n'},
			expectedString: "[\\n..\\n]",
		},
		{
			name:           "VerticalTab",
			r:              disc.Range[Symbol]{Lo: '\v', Hi: '\v'},
			expectedString: "[\\v..\\v]",
		},
		{
			name:           "FormFeed",
			r:              disc.Range[Symbol]{Lo: '\f', Hi: '\f'},
			expectedString: "[\\f..\\f]",
		},
		{
			name:           "CarriageReturn",
			r:              disc.Range[Symbol]{Lo: '\r', Hi: '\r'},
			expectedString: "[\\r..\\r]",
		},
		{
			name:           "Space",
			r:              disc.Range[Symbol]{Lo: ' ', Hi: ' '},
			expectedString: "[SP..SP]",
		},
		{
			name:           "Digit",
			r:              disc.Range[Symbol]{Lo: '0', Hi: '9'},
			expectedString: "[0..9]",
		},
		{
			name:           "Letter",
			r:              disc.Range[Symbol]{Lo: 'A', Hi: 'Z'},
			expectedString: "[A..Z]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, formatRange(tc.r))
		})
	}
}

func TestNewRangeSet(t *testing.T) {
	tests := []struct {
		name           string
		rs             []disc.Range[Symbol]
		expectedString string
	}{
		{
			name: "OK",
			rs: []disc.Range[Symbol]{
				{Lo: 'A', Hi: 'F'},
				{Lo: 'a', Hi: 'f'},
			},
			expectedString: "[A..F], [a..f]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rs := newRangeSet(tc.rs...)

			assert.NotNil(t, rs)
			assert.Equal(t, tc.expectedString, rs.String())
		})
	}
}

func TestNewRangeList(t *testing.T) {
	tests := []struct {
		name           string
		rs             []disc.Range[Symbol]
		expectedString string
	}{
		{
			name: "OK",
			rs: []disc.Range[Symbol]{
				{Lo: 'A', Hi: 'M'},
				{Lo: 'N', Hi: 'Z'},
				{Lo: 'a', Hi: 'm'},
				{Lo: 'n', Hi: 'z'},
			},
			expectedString: "[A..Z], [a..z]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rs := newRangeList(tc.rs...)

			assert.NotNil(t, rs)
			assert.Equal(t, tc.expectedString, rs.String())
		})
	}
}

func TestNewRangeMapping(t *testing.T) {
	tests := []struct {
		name           string
		pairs          []disc.RangeValue[Symbol, classID]
		expectedString string
	}{
		{
			name: "OK",
			pairs: []disc.RangeValue[Symbol, classID]{
				{Range: disc.Range[Symbol]{Lo: 'A', Hi: 'F'}, Value: 0},
				{Range: disc.Range[Symbol]{Lo: 'a', Hi: 'f'}, Value: 0},
			},
			expectedString: `Ranges:
  [A..F]: 0
  [a..f]: 0
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rm := newRangeMapping(tc.pairs)

			assert.NotNil(t, rm)
			assert.Equal(t, tc.expectedString, rm.String())
		})
	}
}

func TestNewClassMapping(t *testing.T) {
	tests := []struct {
		name           string
		pairs          []generic.KeyValue[classID, rangeSet]
		expectedString string
	}{
		{
			name: "OK",
			pairs: []generic.KeyValue[classID, rangeSet]{
				{Key: 0, Val: newRangeSet(disc.Range[Symbol]{Lo: 'A', Hi: 'F'}, disc.Range[Symbol]{Lo: 'a', Hi: 'f'})},
			},
			expectedString: "{<0:[A..F], [a..f]>}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cm := newClassMapping(tc.pairs)

			assert.NotNil(t, cm)
			assert.Equal(t, tc.expectedString, cm.String())
		})
	}
}

func TestStateManager(t *testing.T) {
	tests := []struct {
		name          string
		last          State
		id            int
		s             State
		expectedState State
	}{
		{
			name:          "OK",
			last:          10,
			id:            0,
			s:             1,
			expectedState: 11,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sm := newStateManager(tc.last)

			// Test creating a new state
			state := sm.GetOrCreateState(tc.id, tc.s)
			assert.Equal(t, tc.expectedState, state)

			// Test retrieving the same state again
			state = sm.GetOrCreateState(tc.id, tc.s)
			assert.Equal(t, tc.expectedState, state)
		})
	}
}

func TestGenerateStatePermutations(t *testing.T) {
	tests := []struct {
		name                 string
		states               []State
		start                int
		end                  int
		result               bool
		expectedResult       bool
		expectedPermutations [][]State
	}{
		{
			name:           "OK",
			states:         []State{0, 1, 2},
			start:          0,
			end:            2,
			result:         true,
			expectedResult: true,
			expectedPermutations: [][]State{
				{0, 1, 2},
				{0, 2, 1},
				{1, 0, 2},
				{1, 2, 0},
				{2, 1, 0},
				{2, 0, 1},
			},
		},
		{
			name:           "ReturnEarly",
			states:         []State{0, 1, 2},
			start:          0,
			end:            2,
			result:         false,
			expectedResult: false,
			expectedPermutations: [][]State{
				{0, 1, 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			perms := [][]State{}
			result := generateStatePermutations(tc.states, tc.start, tc.end, func(perm []State) bool {
				clone := make([]State, len(perm))
				copy(clone, perm)

				perms = append(perms, clone)
				return tc.result
			})

			assert.Equal(t, tc.expectedResult, result)
			assert.Equal(t, tc.expectedPermutations, perms)
		})
	}
}
