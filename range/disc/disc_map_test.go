package disc

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func rangeValsToSeq2[T Discrete, V any](pairs []rangeValue[T, V]) iter.Seq2[Range[T], V] {
	return func(yield func(Range[T], V) bool) {
		for _, p := range pairs {
			if !yield(p.Range, p.Value) {
				return
			}
		}
	}
}

func TestDefaultFormatMap(t *testing.T) {
	tests := []struct {
		name           string
		all            iter.Seq2[Range[int], rune]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangeValsToSeq2[int, rune](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangeValsToSeq2([]rangeValue[int, rune]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangeValsToSeq2([]rangeValue[int, rune]{
				{Range[int]{2, 4}, 'a'},
			}),
			expectedString: "[2, 4]:97",
		},
		{
			name: "Many",
			all: rangeValsToSeq2([]rangeValue[int, rune]{
				{Range[int]{2, 4}, 'a'},
				{Range[int]{6, 8}, 'b'},
				{Range[int]{10, 10}, 'c'},
				{Range[int]{16, 20}, 'd'},
			}),
			expectedString: "[2, 4]:97 [6, 8]:98 [10, 10]:99 [16, 20]:100",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, defaultFormatMap(tc.all))
		})
	}
}

func TestNewRangeMap(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()
	format := func(ranges iter.Seq2[Range[int], rune]) string {
		ss := make([]string, 0)
		for r, v := range ranges {
			ss = append(ss, fmt.Sprintf("[%d..%d] --> %c", r.Lo, r.Hi, v))
		}
		return strings.Join(ss, "\n")
	}

	tests := []struct {
		name           string
		equal          generic.EqualFunc[rune]
		opts           RangeMapOpts[int, rune]
		pairs          map[Range[int]]rune
		expectedPairs  []rangeValue[int, rune]
		expectedString string
	}{
		{
			name:  "CurrentHiOnLastHi_Merging",
			equal: equal,
			opts:  RangeMapOpts[int, rune]{},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'a',
				{200, 400}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 400}, 'a'},
			},
			expectedString: "[20, 40]:64 [100, 400]:97",
		},
		{
			name:  "CurrentHiOnLastHi_Merging_CustomFormat",
			equal: equal,
			opts: RangeMapOpts[int, rune]{
				Format: format,
			},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'a',
				{200, 400}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 400}, 'a'},
			},
			expectedString: "[20..40] --> @\n[100..400] --> a",
		},
		{
			name:  "CurrentHiOnLastHi_Splitting",
			equal: equal,
			opts:  RangeMapOpts[int, rune]{},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'b',
				{200, 300}: 'b',
				{300, 400}: 'c',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 299}, 'b'},
				{Range[int]{300, 400}, 'c'},
			},
			expectedString: "[20, 40]:64 [100, 199]:97 [200, 299]:98 [300, 400]:99",
		},
		{
			name:  "CurrentHiOnLastHi_Splitting_CustomFormat",
			equal: equal,
			opts: RangeMapOpts[int, rune]{
				Format: format,
			},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'b',
				{200, 300}: 'b',
				{300, 400}: 'c',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 299}, 'b'},
				{Range[int]{300, 400}, 'c'},
			},
			expectedString: "[20..40] --> @\n[100..199] --> a\n[200..299] --> b\n[300..400] --> c",
		},
		{
			name:  "CurrentHiBeforeLastHi_Merging",
			equal: equal,
			opts:  RangeMapOpts[int, rune]{},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'a',
				{400, 600}: 'a',
				{500, 700}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 700}, 'a'},
			},
			expectedString: "[20, 40]:64 [100, 700]:97",
		},
		{
			name:  "CurrentHiBeforeLastHi_Merging_CustomFormat",
			equal: equal,
			opts: RangeMapOpts[int, rune]{
				Format: format,
			},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'a',
				{400, 600}: 'a',
				{500, 700}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 700}, 'a'},
			},
			expectedString: "[20..40] --> @\n[100..700] --> a",
		},
		{
			name:  "CurrentHiBeforeLastHi_Splitting",
			equal: equal,
			opts:  RangeMapOpts[int, rune]{},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'b',
				{400, 600}: 'b',
				{500, 700}: 'c',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 300}, 'b'},
				{Range[int]{301, 399}, 'a'},
				{Range[int]{400, 499}, 'b'},
				{Range[int]{500, 700}, 'c'},
			},
			expectedString: "[20, 40]:64 [100, 199]:97 [200, 300]:98 [301, 399]:97 [400, 499]:98 [500, 700]:99",
		},
		{
			name:  "CurrentHiBeforeLastHi_Splitting_CustomFormat",
			equal: equal,
			opts: RangeMapOpts[int, rune]{
				Format: format,
			},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'b',
				{400, 600}: 'b',
				{500, 700}: 'c',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 300}, 'b'},
				{Range[int]{301, 399}, 'a'},
				{Range[int]{400, 499}, 'b'},
				{Range[int]{500, 700}, 'c'},
			},
			expectedString: "[20..40] --> @\n[100..199] --> a\n[200..300] --> b\n[301..399] --> a\n[400..499] --> b\n[500..700] --> c",
		},
		{
			name:  "CurrentHiAdjacentToLastHi_Merging",
			equal: equal,
			opts:  RangeMapOpts[int, rune]{},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 199}: 'a',
				{200, 200}: 'a',
				{201, 300}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 300}, 'a'},
			},
			expectedString: "[20, 40]:64 [100, 300]:97",
		},
		{
			name:  "CurrentHiAdjacentToLastHi_Merging_CustomFormat",
			equal: equal,
			opts: RangeMapOpts[int, rune]{
				Format: format,
			},
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 199}: 'a',
				{200, 200}: 'a',
				{201, 300}: 'a',
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{20, 40}, '@'},
				{Range[int]{100, 300}, 'a'},
			},
			expectedString: "[20..40] --> @\n[100..300] --> a",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewRangeMap(tc.equal, tc.opts, tc.pairs).(*rangeMap[int, rune])

			assert.Equal(t, tc.expectedPairs, m.pairs)
			assert.Equal(t, tc.expectedString, m.String())
		})
	}
}

func TestRangeMap_String(t *testing.T) {
	tests := []struct {
		name           string
		m              *rangeMap[int, rune]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedString: "[20, 40]:64 [100, 199]:97 [200, 299]:98 [300, 400]:99",
		},
		{
			name: "WithCustomFormat",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal: generic.NewEqualFunc[rune](),
				format: func(ranges iter.Seq2[Range[int], rune]) string {
					strs := make([]string, 0)
					for r, v := range ranges {
						strs = append(strs, fmt.Sprintf("[%d..%d] --> %c", r.Lo, r.Hi, v))
					}
					return strings.Join(strs, "\n")
				},
				resolve: defaultResolve[rune],
			},
			expectedString: "[20..40] --> @\n[100..199] --> a\n[200..299] --> b\n[300..400] --> c",
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
		m    *rangeMap[int, rune]
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
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
	m := &rangeMap[int, rune]{
		pairs: []rangeValue[int, rune]{
			{Range[int]{20, 40}, '@'},
			{Range[int]{100, 199}, 'a'},
			{Range[int]{200, 299}, 'b'},
			{Range[int]{300, 400}, 'c'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[int, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		rhs           RangeMap[int, rune]
		expectedEqual bool
	}{
		{
			name:          "NotEqual_DiffTypes",
			m:             m,
			rhs:           nil,
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffLens",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs:   []rangeValue[int, rune]{},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{10, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffValues",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '*'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
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
		m            *rangeMap[int, rune]
		expectedSize int
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
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

func TestRangeMap_Find(t *testing.T) {
	m := &rangeMap[int, rune]{
		pairs: []rangeValue[int, rune]{
			{Range[int]{0, 9}, '#'},
			{Range[int]{10, 20}, '@'},
			{Range[int]{200, 400}, 'a'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[int, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		m             *rangeMap[int, rune]
		val           int
		expectedOK    bool
		expectedRange Range[int]
		expectedValue rune
	}{
		{m: m, val: -1, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
		{m: m, val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 10, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 15, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 20, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 100, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
		{m: m, val: 200, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 300, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 400, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 500, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r, v, ok := tc.m.Find(tc.val)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRange, r)
			assert.Equal(t, tc.expectedValue, v)
		})
	}
}

func TestRangeMap_Add(t *testing.T) {
	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		pairs         []rangeValue[int, rune]
		expectedPairs []rangeValue[int, rune]
	}{
		{
			name: "CurrentHiOnLastHi_Merging",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 400}, 'a'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 60}, 'A'},
				{Range[int]{60, 60}, 'A'},
				{Range[int]{60, 80}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{20, 40}, '@'},
				{Range[int]{50, 80}, 'A'},
				{Range[int]{100, 400}, 'a'},
				{Range[int]{1000, 2000}, '$'},
			},
		},
		{
			name: "CurrentHiOnLastHi_Splitting",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 60}, 'A'},
				{Range[int]{60, 60}, 'B'},
				{Range[int]{60, 70}, 'B'},
				{Range[int]{70, 80}, 'C'},
				{Range[int]{1000, 2000}, '$'},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{20, 40}, '@'},
				{Range[int]{50, 59}, 'A'},
				{Range[int]{60, 69}, 'B'},
				{Range[int]{70, 80}, 'C'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 299}, 'b'},
				{Range[int]{300, 400}, 'c'},
				{Range[int]{1000, 2000}, '$'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi_Merging",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 700}, 'a'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 80}, 'A'},
				{Range[int]{55, 65}, 'A'},
				{Range[int]{70, 80}, 'A'},
				{Range[int]{75, 90}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{20, 40}, '@'},
				{Range[int]{50, 90}, 'A'},
				{Range[int]{100, 700}, 'a'},
				{Range[int]{1000, 2000}, '$'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi_Splitting",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 300}, 'b'},
					{Range[int]{301, 399}, 'a'},
					{Range[int]{400, 499}, 'b'},
					{Range[int]{500, 700}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 80}, 'A'},
				{Range[int]{55, 65}, 'B'},
				{Range[int]{70, 80}, 'B'},
				{Range[int]{75, 90}, 'C'},
				{Range[int]{1000, 2000}, '$'},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{20, 40}, '@'},
				{Range[int]{50, 54}, 'A'},
				{Range[int]{55, 65}, 'B'},
				{Range[int]{66, 69}, 'A'},
				{Range[int]{70, 74}, 'B'},
				{Range[int]{75, 90}, 'C'},
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 300}, 'b'},
				{Range[int]{301, 399}, 'a'},
				{Range[int]{400, 499}, 'b'},
				{Range[int]{500, 700}, 'c'},
				{Range[int]{1000, 2000}, '$'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi_Merging",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 300}, 'a'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{60, 69}, 'A'},
				{Range[int]{70, 70}, 'A'},
				{Range[int]{71, 80}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{20, 40}, '@'},
				{Range[int]{60, 80}, 'A'},
				{Range[int]{100, 300}, 'a'},
				{Range[int]{1000, 2000}, '$'},
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
	pairs := []rangeValue[int, rune]{
		{Range[int]{100, 200}, 'a'},
		{Range[int]{300, 400}, 'b'},
		{Range[int]{500, 600}, 'c'},
		{Range[int]{700, 800}, 'd'},
		{Range[int]{900, 1000}, 'e'},
	}

	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		keys          []Range[int]
		expectedPairs []rangeValue[int, rune]
	}{
		{
			name: "None",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: nil,
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'b'},
				{Range[int]{500, 600}, 'c'},
				{Range[int]{700, 800}, 'd'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{40, 60},
				{240, 260},
				{440, 460},
				{640, 660},
				{840, 860},
				{1040, 1060},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'b'},
				{Range[int]{500, 600}, 'c'},
				{Range[int]{700, 800}, 'd'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{80, 100},
				{200, 300},
				{400, 500},
				{600, 700},
				{800, 900},
				{1000, 1020},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{101, 199}, 'a'},
				{Range[int]{301, 399}, 'b'},
				{Range[int]{501, 599}, 'c'},
				{Range[int]{701, 799}, 'd'},
				{Range[int]{901, 999}, 'e'},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{80, 120},
				{180, 320},
				{380, 520},
				{580, 720},
				{780, 920},
				{980, 1020},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{121, 179}, 'a'},
				{Range[int]{321, 379}, 'b'},
				{Range[int]{521, 579}, 'c'},
				{Range[int]{721, 779}, 'd'},
				{Range[int]{921, 979}, 'e'},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{140, 160},
				{340, 360},
				{540, 560},
				{740, 760},
				{940, 960},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{100, 139}, 'a'},
				{Range[int]{161, 200}, 'a'},
				{Range[int]{300, 339}, 'b'},
				{Range[int]{361, 400}, 'b'},
				{Range[int]{500, 539}, 'c'},
				{Range[int]{561, 600}, 'c'},
				{Range[int]{700, 739}, 'd'},
				{Range[int]{761, 800}, 'd'},
				{Range[int]{900, 939}, 'e'},
				{Range[int]{961, 1000}, 'e'},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{250, 450},
				{450, 850},
			},
			expectedPairs: []rangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			name: "All",
			m: &rangeMap[int, rune]{
				pairs:   append([]rangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{100, 200},
				{300, 400},
				{500, 600},
				{700, 800},
				{900, 1000},
			},
			expectedPairs: []rangeValue[int, rune]{},
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
		m           *rangeMap[int, rune]
		expectedAll []generic.KeyValue[Range[int], rune]
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []rangeValue[int, rune]{
					{Range[int]{0, 9}, '#'},
					{Range[int]{10, 20}, '@'},
					{Range[int]{200, 400}, 'a'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{10, 20}, Val: '@'},
				{Key: Range[int]{200, 400}, Val: 'a'},
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
