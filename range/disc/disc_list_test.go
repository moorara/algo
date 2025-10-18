package disc

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func rangesToSeq[T Discrete](ranges []Range[T]) iter.Seq[Range[T]] {
	return func(yield func(Range[T]) bool) {
		for _, v := range ranges {
			if !yield(v) {
				return
			}
		}
	}
}

func TestDefaultFormatList(t *testing.T) {
	tests := []struct {
		name           string
		all            iter.Seq[Range[int]]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangesToSeq[int](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangesToSeq([]Range[int]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangesToSeq([]Range[int]{
				{2, 4},
			}),
			expectedString: "[2, 4]",
		},
		{
			name: "Many",
			all: rangesToSeq([]Range[int]{
				{2, 4},
				{6, 8},
				{10, 10},
				{16, 20},
			}),
			expectedString: "[2, 4] [6, 8] [10, 10] [16, 20]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, defaultFormatList(tc.all))
		})
	}
}

func TestNewRangeList(t *testing.T) {
	format := func(ranges iter.Seq[Range[int]]) string {
		ss := make([]string, 0)
		for r := range ranges {
			ss = append(ss, fmt.Sprintf("[%d..%d]", r.Lo, r.Hi))
		}
		return strings.Join(ss, "\n")
	}

	tests := []struct {
		name           string
		opts           RangeListOpts[int]
		rs             []Range[int]
		expectedRanges []Range[int]
		expectedString string
	}{
		{
			name: "CurrentHiOnLastHi",
			opts: RangeListOpts[int]{},
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
			name: "CurrentHiOnLastHi_CustomFormat",
			opts: RangeListOpts[int]{
				Format: format,
			},
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
			name: "CurrentHiBeforeLastHi",
			opts: RangeListOpts[int]{},
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
			name: "CurrentHiBeforeLastHi_CustomFormat",
			opts: RangeListOpts[int]{
				Format: format,
			},
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
			name: "CurrentHiAdjacentToLastHi",
			opts: RangeListOpts[int]{},
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
		{
			name: "CurrentHiAdjacentToLastHi_CustomFormat",
			opts: RangeListOpts[int]{
				Format: format,
			},
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
			l := NewRangeList(tc.opts, tc.rs...).(*rangeList[int])

			assert.Equal(t, tc.expectedRanges, l.ranges)
			assert.Equal(t, tc.expectedString, l.String())
		})
	}
}

func TestRangeList_String(t *testing.T) {
	tests := []struct {
		name           string
		l              *rangeList[int]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			l: &rangeList[int]{
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
			l: &rangeList[int]{
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
		l    *rangeList[int]
	}{
		{
			name: "OK",
			l: &rangeList[int]{
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
	l := &rangeList[int]{
		ranges: []Range[int]{
			{20, 40},
			{100, 400},
		},
		format: defaultFormatList[int],
	}

	tests := []struct {
		name          string
		l             *rangeList[int]
		rhs           RangeList[int]
		expectedEqual bool
	}{
		{
			name:          "NotEqual_DiffTypes",
			l:             l,
			rhs:           nil,
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffLens",
			l:    l,
			rhs: &rangeList[int]{
				ranges: []Range[int]{},
				format: defaultFormatList[int],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			l:    l,
			rhs: &rangeList[int]{
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
			rhs: &rangeList[int]{
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
		l            *rangeList[int]
		expectedSize int
	}{
		{
			name: "OK",
			l: &rangeList[int]{
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

func TestRangeList_Find(t *testing.T) {
	l := &rangeList[int]{
		ranges: []Range[int]{
			{0, 9},
			{10, 20},
			{200, 400},
		},
		format: defaultFormatList[int],
	}

	tests := []struct {
		l             *rangeList[int]
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
			r, ok := tc.l.Find(tc.val)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRange, r)
		})
	}
}

func TestRangeList_Add(t *testing.T) {
	tests := []struct {
		name           string
		l              *rangeList[int]
		rs             []Range[int]
		expectedRanges []Range[int]
	}{
		{
			name: "CurrentHiOnLastHi",
			l: &rangeList[int]{
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
			l: &rangeList[int]{
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
			l: &rangeList[int]{
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
		l              *rangeList[int]
		rs             []Range[int]
		expectedRanges []Range[int]
	}{
		{
			name: "None",
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
			rs: nil,
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
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
			name: "All",
			l: &rangeList[int]{
				ranges: append([]Range[int]{}, ranges...),
				format: defaultFormatList[int],
			},
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
		l           *rangeList[int]
		expectedAll []Range[int]
	}{
		{
			name: "OK",
			l: &rangeList[int]{
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
