package disc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeList(t *testing.T) {
	type equalTest[T Discrete] struct {
		rhs           *RangeList[T]
		expectedEqual bool
	}

	type getTest[T Discrete] struct {
		val           T
		expectedOK    bool
		expectedRange Range[T]
	}

	tests := []struct {
		name           string
		rs             []Range[int]
		addRanges      []Range[int]
		equalTests     []equalTest[int]
		getTests       []getTest[int]
		expectedAll    []Range[int]
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
			addRanges: []Range[int]{
				{0, 9},
				{50, 60},
				{60, 60},
				{60, 80},
				{1000, 2000},
			},
			equalTests: []equalTest[int]{
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{50, 80},
							{100, 400},
							{1000, 4000},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{50, 80},
							{100, 400},
							{1000, 2000},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int]{
				{val: -10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 20, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 30, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 40, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 44, expectedOK: false, expectedRange: Range[int]{}},
				{val: 50, expectedOK: true, expectedRange: Range[int]{50, 80}},
				{val: 60, expectedOK: true, expectedRange: Range[int]{50, 80}},
				{val: 80, expectedOK: true, expectedRange: Range[int]{50, 80}},
				{val: 99, expectedOK: false, expectedRange: Range[int]{}},
				{val: 100, expectedOK: true, expectedRange: Range[int]{100, 400}},
				{val: 200, expectedOK: true, expectedRange: Range[int]{100, 400}},
				{val: 400, expectedOK: true, expectedRange: Range[int]{100, 400}},
				{val: 500, expectedOK: false, expectedRange: Range[int]{}},
				{val: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 4000, expectedOK: false, expectedRange: Range[int]{}},
			},
			expectedAll: []Range[int]{
				{0, 9},
				{20, 40},
				{50, 80},
				{100, 400},
				{1000, 2000},
			},
			expectedString: "[0, 9] [20, 40] [50, 80] [100, 400] [1000, 2000]",
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
			addRanges: []Range[int]{
				{0, 9},
				{50, 80},
				{55, 65},
				{70, 80},
				{75, 90},
				{1000, 2000},
			},
			equalTests: []equalTest[int]{
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{50, 90},
							{100, 700},
							{1000, 4000},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{50, 90},
							{100, 700},
							{1000, 2000},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int]{
				{val: -10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 20, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 30, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 40, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 44, expectedOK: false, expectedRange: Range[int]{}},
				{val: 50, expectedOK: true, expectedRange: Range[int]{50, 90}},
				{val: 70, expectedOK: true, expectedRange: Range[int]{50, 90}},
				{val: 90, expectedOK: true, expectedRange: Range[int]{50, 90}},
				{val: 99, expectedOK: false, expectedRange: Range[int]{}},
				{val: 100, expectedOK: true, expectedRange: Range[int]{100, 700}},
				{val: 500, expectedOK: true, expectedRange: Range[int]{100, 700}},
				{val: 700, expectedOK: true, expectedRange: Range[int]{100, 700}},
				{val: 800, expectedOK: false, expectedRange: Range[int]{}},
				{val: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 4000, expectedOK: false, expectedRange: Range[int]{}},
			},
			expectedAll: []Range[int]{
				{0, 9},
				{20, 40},
				{50, 90},
				{100, 700},
				{1000, 2000},
			},
			expectedString: "[0, 9] [20, 40] [50, 90] [100, 700] [1000, 2000]",
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			rs: []Range[int]{
				{20, 40},
				{100, 199},
				{200, 200},
				{201, 300},
			},
			addRanges: []Range[int]{
				{0, 9},
				{60, 69},
				{70, 70},
				{71, 80},
				{1000, 2000},
			},
			equalTests: []equalTest[int]{
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{60, 80},
							{100, 300},
							{1000, 4000},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[int]{
						ranges: []Range[int]{
							{0, 9},
							{20, 40},
							{60, 80},
							{100, 300},
							{1000, 2000},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int]{
				{val: -10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}},
				{val: 10, expectedOK: false, expectedRange: Range[int]{}},
				{val: 20, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 30, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 40, expectedOK: true, expectedRange: Range[int]{20, 40}},
				{val: 50, expectedOK: false, expectedRange: Range[int]{}},
				{val: 60, expectedOK: true, expectedRange: Range[int]{60, 80}},
				{val: 70, expectedOK: true, expectedRange: Range[int]{60, 80}},
				{val: 80, expectedOK: true, expectedRange: Range[int]{60, 80}},
				{val: 99, expectedOK: false, expectedRange: Range[int]{}},
				{val: 100, expectedOK: true, expectedRange: Range[int]{100, 300}},
				{val: 200, expectedOK: true, expectedRange: Range[int]{100, 300}},
				{val: 300, expectedOK: true, expectedRange: Range[int]{100, 300}},
				{val: 500, expectedOK: false, expectedRange: Range[int]{}},
				{val: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}},
				{val: 4000, expectedOK: false, expectedRange: Range[int]{}},
			},
			expectedAll: []Range[int]{
				{0, 9},
				{20, 40},
				{60, 80},
				{100, 300},
				{1000, 2000},
			},
			expectedString: "[0, 9] [20, 40] [60, 80] [100, 300] [1000, 2000]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var l *RangeList[int]

			t.Run("NewRangeList", func(t *testing.T) {
				l = NewRangeList(tc.rs...)
			})

			t.Run("Clone", func(t *testing.T) {
				clone := l.Clone()
				assert.True(t, clone.Equal(l))
			})

			t.Run("Add", func(t *testing.T) {
				l.Add(tc.addRanges...)
			})

			for i, tc := range tc.equalTests {
				t.Run(fmt.Sprintf("Equal/%d", i), func(t *testing.T) {
					assert.Equal(t, tc.expectedEqual, l.Equal(tc.rhs))
				})
			}

			for i, tc := range tc.getTests {
				t.Run(fmt.Sprintf("Get/%d", i), func(t *testing.T) {
					r, ok := l.Get(tc.val)

					assert.Equal(t, tc.expectedOK, ok)
					assert.Equal(t, tc.expectedRange, r)
				})
			}

			t.Run("All", func(t *testing.T) {
				all := generic.Collect1(l.All())
				assert.Equal(t, tc.expectedAll, all)
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, l.String())
			})
		})
	}
}
