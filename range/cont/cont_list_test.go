package cont

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
		rs             []Range[float64]
		expectedRanges []Range[float64]
		expectedString string
	}{
		{
			name: "CurrentHiOnLastHi",
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
			},
			expectedString: "[2, 4] [10, 40]",
		},
		{
			name: "CurrentHiBeforeLastHi",
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
			},
			expectedString: "[2, 4] [10, 70]",
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}},
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, true}, Bound[float64]{30.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
			},
			expectedString: "[2, 4] [10, 30]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewRangeList(tc.rs...).(*rangeList[float64])

			assert.Equal(t, tc.expectedRanges, l.ranges)
			assert.Equal(t, tc.expectedString, l.String())
		})
	}
}

func TestNewRangeListWithFormat(t *testing.T) {
	format := func(ranges iter.Seq[Range[float64]]) string {
		strs := make([]string, 0)
		for r := range ranges {
			strs = append(strs, r.String())
		}
		return strings.Join(strs, "\n")
	}

	tests := []struct {
		name           string
		format         FormatList[float64]
		rs             []Range[float64]
		expectedRanges []Range[float64]
		expectedString string
	}{
		{
			name:   "CurrentHiOnLastHi",
			format: format,
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
			},
			expectedString: "[2, 4]\n[10, 40]",
		},
		{
			name:   "CurrentHiBeforeLastHi",
			format: format,
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
			},
			expectedString: "[2, 4]\n[10, 70]",
		},
		{
			name:   "CurrentHiAdjacentToLastHi",
			format: format,
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}},
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, true}, Bound[float64]{30.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
			},
			expectedString: "[2, 4]\n[10, 30]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := NewRangeListWithFormat(tc.format, tc.rs...).(*rangeList[float64])

			assert.Equal(t, tc.expectedRanges, l.ranges)
			assert.Equal(t, tc.expectedString, l.String())
		})
	}
}

func TestRangeList_String(t *testing.T) {
	tests := []struct {
		name           string
		l              *rangeList[float64]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
			},
			expectedString: "[2, 4] [10, 40]",
		},
		{
			name: "WithCustomFormat",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: func(ranges iter.Seq[Range[float64]]) string {
					strs := make([]string, 0)
					for r := range ranges {
						strs = append(strs, r.String())
					}
					return strings.Join(strs, "\n")
				},
			},
			expectedString: "[2, 4]\n[10, 40]",
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
		l    *rangeList[float64]
	}{
		{
			name: "OK",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
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
	l := &rangeList[float64]{
		ranges: []Range[float64]{
			{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
			{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
		},
		format: defaultFormatList[float64],
	}

	tests := []struct {
		name          string
		l             *rangeList[float64]
		rhs           RangeList[float64]
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
			rhs: &rangeList[float64]{
				ranges: []Range[float64]{},
				format: defaultFormatList[float64],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			l:    l,
			rhs: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			l:    l,
			rhs: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
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
		l            *rangeList[float64]
		expectedSize int
	}{
		{
			name: "OK",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
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
	l := &rangeList[float64]{
		ranges: []Range[float64]{
			{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
			{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}},
			{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}},
			{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}},
		},
		format: defaultFormatList[float64],
	}

	tests := []struct {
		l             *rangeList[float64]
		val           float64
		expectedOK    bool
		expectedRange Range[float64]
	}{
		{l: l, val: -1.0, expectedOK: false, expectedRange: Range[float64]{}},
		{l: l, val: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
		{l: l, val: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
		{l: l, val: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
		{l: l, val: 1.0, expectedOK: false, expectedRange: Range[float64]{}},
		{l: l, val: 1.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}},
		{l: l, val: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}},
		{l: l, val: 10.0, expectedOK: false, expectedRange: Range[float64]{}},
		{l: l, val: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}},
		{l: l, val: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}},
		{l: l, val: 40.0, expectedOK: false, expectedRange: Range[float64]{}},
		{l: l, val: 60.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}},
		{l: l, val: 80.0, expectedOK: false, expectedRange: Range[float64]{}},
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
		l              *rangeList[float64]
		rs             []Range[float64]
		expectedRanges []Range[float64]
	}{
		{
			name: "CurrentHiOnLastHi",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				},
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
		},
		{
			name: "CurrentHiBeforeLastHi",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
				},
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}},
				{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
					{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
				},
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}},
				{Bound[float64]{7.0, false}, Bound[float64]{7.0, false}},
				{Bound[float64]{7.0, true}, Bound[float64]{8.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
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
	ranges := []Range[float64]{
		{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
		{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}},
		{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}},
		{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
		{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
	}

	tests := []struct {
		name           string
		l              *rangeList[float64]
		rs             []Range[float64]
		expectedRanges []Range[float64]
	}{
		{
			name: "None",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: nil,
			expectedRanges: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{4.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{24.0, false}, Bound[float64]{26.0, false}},
				{Bound[float64]{44.0, false}, Bound[float64]{46.0, false}},
				{Bound[float64]{64.0, false}, Bound[float64]{66.0, false}},
				{Bound[float64]{84.0, false}, Bound[float64]{86.0, false}},
				{Bound[float64]{104.0, false}, Bound[float64]{106.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{10.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{40.0, false}, Bound[float64]{50.0, false}},
				{Bound[float64]{60.0, false}, Bound[float64]{70.0, false}},
				{Bound[float64]{80.0, false}, Bound[float64]{90.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{102.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{10.0, true}, Bound[float64]{20.0, true}},
				{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}},
				{Bound[float64]{50.0, true}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, true}, Bound[float64]{100.0, true}},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{12.0, false}},
				{Bound[float64]{18.0, false}, Bound[float64]{32.0, false}},
				{Bound[float64]{38.0, false}, Bound[float64]{52.0, false}},
				{Bound[float64]{58.0, false}, Bound[float64]{72.0, false}},
				{Bound[float64]{78.0, false}, Bound[float64]{92.0, false}},
				{Bound[float64]{98.0, false}, Bound[float64]{102.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{12.0, true}, Bound[float64]{18.0, true}},
				{Bound[float64]{32.0, true}, Bound[float64]{38.0, true}},
				{Bound[float64]{52.0, true}, Bound[float64]{58.0, true}},
				{Bound[float64]{72.0, true}, Bound[float64]{78.0, true}},
				{Bound[float64]{92.0, true}, Bound[float64]{98.0, true}},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{14.0, true}, Bound[float64]{16.0, true}},
				{Bound[float64]{34.0, true}, Bound[float64]{36.0, true}},
				{Bound[float64]{54.0, true}, Bound[float64]{56.0, true}},
				{Bound[float64]{74.0, true}, Bound[float64]{76.0, true}},
				{Bound[float64]{94.0, true}, Bound[float64]{96.0, true}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{14.0, false}},
				{Bound[float64]{16.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{30.0, true}, Bound[float64]{34, false}},
				{Bound[float64]{36.0, false}, Bound[float64]{40.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{54.0, false}},
				{Bound[float64]{56.0, false}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{74.0, false}},
				{Bound[float64]{76.0, false}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, false}, Bound[float64]{94.0, false}},
				{Bound[float64]{96.0, false}, Bound[float64]{100.0, false}},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs: []Range[float64]{
				{Bound[float64]{25.0, false}, Bound[float64]{45.0, false}},
				{Bound[float64]{45.0, true}, Bound[float64]{85.0, false}},
			},
			expectedRanges: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
			},
		},
		{
			name: "All",
			l: &rangeList[float64]{
				ranges: append([]Range[float64]{}, ranges...),
				format: defaultFormatList[float64],
			},
			rs:             ranges,
			expectedRanges: []Range[float64]{},
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
		l           *rangeList[float64]
		expectedAll []Range[float64]
	}{
		{
			name: "OK",
			l: &rangeList[float64]{
				ranges: []Range[float64]{
					{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
					{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}},
					{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}},
					{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}},
				},
				format: defaultFormatList[float64],
			},
			expectedAll: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}},
				{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}},
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
