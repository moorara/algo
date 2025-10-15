package cont

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNewRangeMap(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()

	tests := []struct {
		name           string
		equal          generic.EqualFunc[rune]
		pairs          map[Range[float64]]rune
		expectedPairs  []rangeValue[float64, rune]
		expectedString string
	}{
		{
			name:  "CurrentHiOnLastHi_Merging",
			equal: equal,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}: 'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
			expectedString: "[2, 4]:64 [10, 40]:97",
		},
		{
			name:  "CurrentHiOnLastHi_Splitting",
			equal: equal,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'b',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}: 'c',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
			},
			expectedString: "[2, 4]:64 [10, 20):97 [20, 30):98 [30, 40]:99",
		},
		{
			name:  "CurrentHiBeforeLastHi_Merging",
			equal: equal,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'a',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedString: "[2, 4]:64 [10, 70]:97",
		},
		{
			name:  "CurrentHiBeforeLastHi_Splitting",
			equal: equal,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'b',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'c',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
			},
			expectedString: "[2, 4]:64 [10, 20):97 [20, 30]:98 (30, 40):97 [40, 50):98 [50, 70]:99",
		},
		{
			name:  "CurrentHiAdjacentToLastHi_Merging",
			equal: equal,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}:  'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20, true}, Bound[float64]{30.0, false}}:    'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedString: "[2, 4]:64 [10, 30]:97",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewRangeMap(tc.equal, tc.pairs)

			assert.Equal(t, tc.expectedPairs, m.pairs)
			assert.Equal(t, tc.expectedString, m.String())
		})
	}
}

func TestNewRangeMapWithFormat(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()

	format := func(ranges iter.Seq2[Range[float64], rune]) string {
		strs := make([]string, 0)
		for r, v := range ranges {
			strs = append(strs, fmt.Sprintf("%s --> %c", r.String(), v))
		}
		return strings.Join(strs, "\n")
	}

	tests := []struct {
		name           string
		equal          generic.EqualFunc[rune]
		format         FormatMap[float64, rune]
		pairs          map[Range[float64]]rune
		expectedPairs  []rangeValue[float64, rune]
		expectedString string
	}{
		{
			name:   "CurrentHiOnLastHi_Merging",
			equal:  equal,
			format: format,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}: 'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
			expectedString: "[2, 4] --> @\n[10, 40] --> a",
		},
		{
			name:   "CurrentHiOnLastHi_Splitting",
			equal:  equal,
			format: format,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'b',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}: 'c',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
			},
			expectedString: "[2, 4] --> @\n[10, 20) --> a\n[20, 30) --> b\n[30, 40] --> c",
		},
		{
			name:   "CurrentHiBeforeLastHi_Merging",
			equal:  equal,
			format: format,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'a',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedString: "[2, 4] --> @\n[10, 70] --> a",
		},
		{
			name:   "CurrentHiBeforeLastHi_Splitting",
			equal:  equal,
			format: format,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'b',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'c',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
			},
			expectedString: "[2, 4] --> @\n[10, 20) --> a\n[20, 30] --> b\n(30, 40) --> a\n[40, 50) --> b\n[50, 70] --> c",
		},
		{
			name:   "CurrentHiAdjacentToLastHi_Merging",
			equal:  equal,
			format: format,
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}:  'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20, true}, Bound[float64]{30.0, false}}:    'a',
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedString: "[2, 4] --> @\n[10, 30] --> a",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewRangeMapWithFormat(tc.equal, tc.format, tc.pairs)

			assert.Equal(t, tc.expectedPairs, m.pairs)
			assert.Equal(t, tc.expectedString, m.String())
		})
	}
}

func TestRangeMap_String(t *testing.T) {
	tests := []struct {
		name           string
		m              *RangeMap[float64, rune]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedString: "[2, 4]:64 [10, 20):97 [20, 30):98 [30, 40]:99",
		},
		{
			name: "WithCustomFormat",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal: generic.NewEqualFunc[rune](),
				format: func(ranges iter.Seq2[Range[float64], rune]) string {
					strs := make([]string, 0)
					for r, v := range ranges {
						strs = append(strs, fmt.Sprintf("%s --> %c", r.String(), v))
					}
					return strings.Join(strs, "\n")
				},
			},
			expectedString: "[2, 4] --> @\n[10, 20) --> a\n[20, 30) --> b\n[30, 40] --> c",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.m.String())
		})
	}
}

func TestRangeMap_Clone(t *testing.T) {
	tests := []struct {
		name string
		m    *RangeMap[float64, rune]
	}{
		{
			name: "OK",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.m.Clone()

			assert.True(t, clone.Equal(tc.m))
		})
	}
}

func TestRangeMap_Equal(t *testing.T) {
	m := &RangeMap[float64, rune]{
		pairs: []rangeValue[float64, rune]{
			{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
			{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
			{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
			{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
		},
		equal:  generic.NewEqualFunc[rune](),
		format: defaultFormatMap[float64, rune],
	}

	tests := []struct {
		name          string
		m             *RangeMap[float64, rune]
		rhs           *RangeMap[float64, rune]
		expectedEqual bool
	}{
		{
			name: "NotEqual_DiffLens",
			m:    m,
			rhs: &RangeMap[float64, rune]{
				pairs:  []rangeValue[float64, rune]{},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			m:    m,
			rhs: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffValues",
			m:    m,
			rhs: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '*'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			m:    m,
			rhs: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.m.Equal(tc.rhs))
		})
	}
}

func TestRangeMap_Size(t *testing.T) {
	tests := []struct {
		name         string
		m            *RangeMap[float64, rune]
		expectedSize int
	}{
		{
			name: "OK",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedSize: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSize, tc.m.Size())
		})
	}
}

func TestRangeMap_Get(t *testing.T) {
	m := &RangeMap[float64, rune]{
		pairs: []rangeValue[float64, rune]{
			{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
			{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, '@'},
			{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, 'a'},
			{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, 'b'},
		},
		equal:  generic.NewEqualFunc[rune](),
		format: defaultFormatMap[float64, rune],
	}

	tests := []struct {
		m             *RangeMap[float64, rune]
		key           float64
		expectedOK    bool
		expectedRange Range[float64]
		expectedValue rune
	}{
		{m: m, key: -1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, key: 1.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, expectedValue: '@'},
		{m: m, key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, expectedValue: '@'},
		{m: m, key: 10.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, key: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
		{m: m, key: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
		{m: m, key: 40.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, key: 60.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, expectedValue: 'b'},
		{m: m, key: 80.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r, v, ok := tc.m.Get(tc.key)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRange, r)
			assert.Equal(t, tc.expectedValue, v)
		})
	}
}

func TestRangeMap_Add(t *testing.T) {
	tests := []struct {
		name          string
		m             *RangeMap[float64, rune]
		pairs         []rangeValue[float64, rune]
		expectedPairs []rangeValue[float64, rune]
	}{
		{
			name: "CurrentHiOnLastHi_Merging",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			pairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
		},
		{
			name: "CurrentHiOnLastHi_Splitting",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			pairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, 'B'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi_Merging",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			pairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi_Splitting",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			pairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, 'A'},
				{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'B'},
				{Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, 'B'},
				{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi_Merging",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			pairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, true}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, p := range tc.pairs {
				tc.m.Add(p.Range, p.Value)
			}

			assert.Equal(t, tc.expectedPairs, tc.m.pairs)
		})
	}
}

func TestRangeMap_Remove(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()
	pairs := map[Range[float64]]rune{
		{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}:  'a',
		{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}:   'b',
		{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}:   'c',
		{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}:    'd',
		{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}: 'e',
	}

	tests := []struct {
		name          string
		m             *RangeMap[float64, rune]
		keys          []Range[float64]
		expectedPairs []rangeValue[float64, rune]
	}{
		{
			name: "None",
			m:    NewRangeMap(equal, pairs),
			keys: nil,
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{4.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{24.0, false}, Bound[float64]{26.0, false}},
				{Bound[float64]{44.0, false}, Bound[float64]{46.0, false}},
				{Bound[float64]{64.0, false}, Bound[float64]{66.0, false}},
				{Bound[float64]{84.0, false}, Bound[float64]{86.0, false}},
				{Bound[float64]{104.0, false}, Bound[float64]{106.0, false}},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{10.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{40.0, false}, Bound[float64]{50.0, false}},
				{Bound[float64]{60.0, false}, Bound[float64]{70.0, false}},
				{Bound[float64]{80.0, false}, Bound[float64]{90.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{102.0, false}},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, true}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, true}, Bound[float64]{100.0, true}}, 'e'},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{12.0, false}},
				{Bound[float64]{18.0, false}, Bound[float64]{32.0, false}},
				{Bound[float64]{38.0, false}, Bound[float64]{52.0, false}},
				{Bound[float64]{58.0, false}, Bound[float64]{72.0, false}},
				{Bound[float64]{78.0, false}, Bound[float64]{92.0, false}},
				{Bound[float64]{98.0, false}, Bound[float64]{102.0, false}},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{12.0, true}, Bound[float64]{18.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{32.0, true}, Bound[float64]{38.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{52.0, true}, Bound[float64]{58.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{72.0, true}, Bound[float64]{78.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{92.0, true}, Bound[float64]{98.0, true}}, 'e'},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{14.0, true}, Bound[float64]{16.0, true}},
				{Bound[float64]{34.0, true}, Bound[float64]{36.0, true}},
				{Bound[float64]{54.0, true}, Bound[float64]{56.0, true}},
				{Bound[float64]{74.0, true}, Bound[float64]{76.0, true}},
				{Bound[float64]{94.0, true}, Bound[float64]{96.0, true}},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{14.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{16.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{34, false}}, 'b'},
				{Range[float64]{Bound[float64]{36.0, false}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{54.0, false}}, 'c'},
				{Range[float64]{Bound[float64]{56.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{74.0, false}}, 'd'},
				{Range[float64]{Bound[float64]{76.0, false}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{94.0, false}}, 'e'},
				{Range[float64]{Bound[float64]{96.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{25.0, false}, Bound[float64]{45.0, false}},
				{Bound[float64]{45.0, true}, Bound[float64]{85.0, false}},
			},
			expectedPairs: []rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			name: "All",
			m:    NewRangeMap(equal, pairs),
			keys: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
			},
			expectedPairs: []rangeValue[float64, rune]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, k := range tc.keys {
				tc.m.Remove(k)
			}

			assert.Equal(t, tc.expectedPairs, tc.m.pairs)
		})
	}
}

func TestRangeMap_All(t *testing.T) {
	tests := []struct {
		name        string
		m           *RangeMap[float64, rune]
		expectedAll []generic.KeyValue[Range[float64], rune]
	}{
		{
			name: "OK",
			m: &RangeMap[float64, rune]{
				pairs: []rangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
					{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, '@'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, 'b'},
				},
				equal:  generic.NewEqualFunc[rune](),
				format: defaultFormatMap[float64, rune],
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, Val: 'b'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			all := generic.Collect2(tc.m.All())

			assert.Equal(t, tc.expectedAll, all)
		})
	}
}
