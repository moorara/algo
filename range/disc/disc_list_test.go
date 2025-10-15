package disc

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNewRangeList(t *testing.T) {
	tests := []struct {
		name           string
		rs             []Range[int]
		expectedRanges []Range[int]
		expectedString string
	}{
		{
			name: "CurrentHiOnLastHi",
			rs: []Range[int]{
				{20, 40},
				{100, 200},
				{200, 200},
				{200, 400},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 400},
			},
			expectedString: "[20, 40] [100, 400]",
		},
		{
			name: "CurrentHiBeforeLastHi",
			rs: []Range[int]{
				{20, 40},
				{100, 600},
				{200, 300},
				{400, 600},
				{500, 700},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 700},
			},
			expectedString: "[20, 40] [100, 700]",
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			rs: []Range[int]{
				{20, 40},
				{100, 199},
				{200, 200},
				{201, 300},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 300},
			},
			expectedString: "[20, 40] [100, 300]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewRangeList(tc.rs...)

			assert.Equal(t, tc.expectedRanges, l.ranges)
			assert.Equal(t, tc.expectedString, l.String())
		})
	}
}

func TestNewRangeListWithFormat(t *testing.T) {
	format := func(ranges iter.Seq[Range[int]]) string {
		strs := make([]string, 0)
		for r := range ranges {
			strs = append(strs, fmt.Sprintf("[%d..%d]", r.Lo, r.Hi))
		}
		return strings.Join(strs, "\n")
	}

	tests := []struct {
		name           string
		format         FormatList[int]
		rs             []Range[int]
		expectedRanges []Range[int]
		expectedString string
	}{
		{
			name:   "CurrentHiOnLastHi",
			format: format,
			rs: []Range[int]{
				{20, 40},
				{100, 200},
				{200, 200},
				{200, 400},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 400},
			},
			expectedString: "[20..40]\n[100..400]",
		},
		{
			name:   "CurrentHiBeforeLastHi",
			format: format,
			rs: []Range[int]{
				{20, 40},
				{100, 600},
				{200, 300},
				{400, 600},
				{500, 700},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 700},
			},
			expectedString: "[20..40]\n[100..700]",
		},
		{
			name:   "CurrentHiAdjacentToLastHi",
			format: format,
			rs: []Range[int]{
				{20, 40},
				{100, 199},
				{200, 200},
				{201, 300},
			},
			expectedRanges: []Range[int]{
				{20, 40},
				{100, 300},
			},
			expectedString: "[20..40]\n[100..300]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewRangeListWithFormat(tc.format, tc.rs...)

			assert.Equal(t, tc.expectedRanges, l.ranges)
			assert.Equal(t, tc.expectedString, l.String())
		})
	}
}

func TestRangeList_String(t *testing.T) {
	tests := []struct {
		name           string
		l              *RangeList[int]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
			expectedString: "[20, 40] [100, 400]",
		},
		{
			name: "WithCustomFormat",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: func(ranges iter.Seq[Range[int]]) string {
					strs := make([]string, 0)
					for r := range ranges {
						strs = append(strs, fmt.Sprintf("[%d..%d]", r.Lo, r.Hi))
					}
					return strings.Join(strs, "\n")
				},
			},
			expectedString: "[20..40]\n[100..400]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.l.String())
		})
	}
}

func TestRangeList_Clone(t *testing.T) {
	tests := []struct {
		name string
		l    *RangeList[int]
	}{
		{
			name: "OK",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.l.Clone()

			assert.True(t, clone.Equal(tc.l))
		})
	}
}

func TestRangeList_Equal(t *testing.T) {
	l := &RangeList[int]{
		ranges: []Range[int]{
			{20, 40},
			{100, 400},
		},
		format: defaultFormatList[int],
	}

	tests := []struct {
		name          string
		l             *RangeList[int]
		rhs           *RangeList[int]
		expectedEqual bool
	}{
		{
			name: "NotEqual_DiffLens",
			l:    l,
			rhs: &RangeList[int]{
				ranges: []Range[int]{},
				format: defaultFormatList[int],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			l:    l,
			rhs: &RangeList[int]{
				ranges: []Range[int]{
					{10, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			l:    l,
			rhs: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.l.Equal(tc.rhs))
		})
	}
}

func TestRangeList_Size(t *testing.T) {
	tests := []struct {
		name         string
		l            *RangeList[int]
		expectedSize int
	}{
		{
			name: "OK",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
			expectedSize: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSize, tc.l.Size())
		})
	}
}

func TestRangeList_Get(t *testing.T) {
	l := &RangeList[int]{
		ranges: []Range[int]{
			{0, 9},
			{10, 20},
			{200, 400},
		},
		format: defaultFormatList[int],
	}

	tests := []struct {
		l             *RangeList[int]
		val           int
		expectedOK    bool
		expectedRange Range[int]
	}{
		{l: l, val: -1, expectedOK: false, expectedRange: Range[int]{}},
		{l: l, val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}},
		{l: l, val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}},
		{l: l, val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}},
		{l: l, val: 10, expectedOK: true, expectedRange: Range[int]{10, 20}},
		{l: l, val: 15, expectedOK: true, expectedRange: Range[int]{10, 20}},
		{l: l, val: 20, expectedOK: true, expectedRange: Range[int]{10, 20}},
		{l: l, val: 100, expectedOK: false, expectedRange: Range[int]{}},
		{l: l, val: 200, expectedOK: true, expectedRange: Range[int]{200, 400}},
		{l: l, val: 300, expectedOK: true, expectedRange: Range[int]{200, 400}},
		{l: l, val: 400, expectedOK: true, expectedRange: Range[int]{200, 400}},
		{l: l, val: 500, expectedOK: false, expectedRange: Range[int]{}},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r, ok := tc.l.Get(tc.val)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRange, r)
		})
	}
}

func TestRangeList_Add(t *testing.T) {
	tests := []struct {
		name           string
		l              *RangeList[int]
		rs             []Range[int]
		expectedRanges []Range[int]
	}{
		{
			name: "CurrentHiOnLastHi",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 400},
				},
				format: defaultFormatList[int],
			},
			rs: []Range[int]{
				{0, 9},
				{50, 60},
				{60, 60},
				{60, 80},
				{1000, 2000},
			},
			expectedRanges: []Range[int]{
				{0, 9},
				{20, 40},
				{50, 80},
				{100, 400},
				{1000, 2000},
			},
		},
		{
			name: "CurrentHiBeforeLastHi",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 700},
				},
				format: defaultFormatList[int],
			},
			rs: []Range[int]{
				{0, 9},
				{50, 80},
				{55, 65},
				{70, 80},
				{75, 90},
				{1000, 2000},
			},
			expectedRanges: []Range[int]{
				{0, 9},
				{20, 40},
				{50, 90},
				{100, 700},
				{1000, 2000},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{20, 40},
					{100, 300},
				},
				format: defaultFormatList[int],
			},
			rs: []Range[int]{
				{0, 9},
				{60, 69},
				{70, 70},
				{71, 80},
				{1000, 2000},
			},
			expectedRanges: []Range[int]{
				{0, 9},
				{20, 40},
				{60, 80},
				{100, 300},
				{1000, 2000},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.l.Add(tc.rs...)

			assert.Equal(t, tc.expectedRanges, tc.l.ranges)
		})
	}
}

func TestRangeList_Remove(t *testing.T) {
	ranges := []Range[int]{
		{100, 200},
		{300, 400},
		{500, 600},
		{700, 800},
		{900, 1000},
	}

	tests := []struct {
		name           string
		l              *RangeList[int]
		rs             []Range[int]
		expectedRanges []Range[int]
	}{
		{
			name: "None",
			l:    NewRangeList(ranges...),
			rs:   nil,
			expectedRanges: []Range[int]{
				{100, 200},
				{300, 400},
				{500, 600},
				{700, 800},
				{900, 1000},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			l:    NewRangeList(ranges...),
			rs: []Range[int]{
				{40, 60},
				{240, 260},
				{440, 460},
				{640, 660},
				{840, 860},
				{1040, 1060},
			},
			expectedRanges: []Range[int]{
				{100, 200},
				{300, 400},
				{500, 600},
				{700, 800},
				{900, 1000},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			l:    NewRangeList(ranges...),
			rs: []Range[int]{
				{80, 100},
				{200, 300},
				{400, 500},
				{600, 700},
				{800, 900},
				{1000, 1020},
			},
			expectedRanges: []Range[int]{
				{101, 199},
				{301, 399},
				{501, 599},
				{701, 799},
				{901, 999},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			l:    NewRangeList(ranges...),
			rs: []Range[int]{
				{80, 120},
				{180, 320},
				{380, 520},
				{580, 720},
				{780, 920},
				{980, 1020},
			},
			expectedRanges: []Range[int]{
				{121, 179},
				{321, 379},
				{521, 579},
				{721, 779},
				{921, 979},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			l:    NewRangeList(ranges...),
			rs: []Range[int]{
				{140, 160},
				{340, 360},
				{540, 560},
				{740, 760},
				{940, 960},
			},
			expectedRanges: []Range[int]{
				{100, 139},
				{161, 200},
				{300, 339},
				{361, 400},
				{500, 539},
				{561, 600},
				{700, 739},
				{761, 800},
				{900, 939},
				{961, 1000},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			l:    NewRangeList(ranges...),
			rs: []Range[int]{
				{250, 450},
				{450, 850},
			},
			expectedRanges: []Range[int]{
				{100, 200},
				{900, 1000},
			},
		},
		{
			name:           "All",
			l:              NewRangeList(ranges...),
			rs:             ranges,
			expectedRanges: []Range[int]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.l.Remove(tc.rs...)

			assert.Equal(t, tc.expectedRanges, tc.l.ranges)
		})
	}
}

func TestRangeList_All(t *testing.T) {
	tests := []struct {
		name        string
		l           *RangeList[int]
		expectedAll []Range[int]
	}{
		{
			name: "OK",
			l: &RangeList[int]{
				ranges: []Range[int]{
					{0, 9},
					{10, 20},
					{200, 400},
				},
				format: defaultFormatList[int],
			},
			expectedAll: []Range[int]{
				{0, 9},
				{10, 20},
				{200, 400},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			all := generic.Collect1(tc.l.All())

			assert.Equal(t, tc.expectedAll, all)
		})
	}
}
