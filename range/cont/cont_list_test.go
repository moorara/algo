package cont

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeList(t *testing.T) {
	type equalTest[T Continuous] struct {
		rhs           *RangeList[T]
		expectedEqual bool
	}

	type getTest[T Continuous] struct {
		val           T
		expectedOK    bool
		expectedRange Range[T]
	}

	tests := []struct {
		name           string
		rs             []Range[float64]
		addRanges      []Range[float64]
		equalTests     []equalTest[float64]
		getTests       []getTest[float64]
		expectedAll    []Range[float64]
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
			addRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			equalTests: []equalTest[float64]{
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{400.0, false}},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64]{
				{val: -1, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 1.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 4.4, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}},
				{val: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}},
				{val: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}},
				{val: 9.9, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}},
				{val: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}},
				{val: 40.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}},
				{val: 50.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 400.0, expectedOK: false, expectedRange: Range[float64]{}},
			},
			expectedAll: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedString: "[0, 0.9] [2, 4] [5, 8] [10, 40] [100, 200]",
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
			addRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}},
				{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			equalTests: []equalTest[float64]{
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{400.0, false}},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64]{
				{val: -1, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 1.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 4.4, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}},
				{val: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}},
				{val: 9.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}},
				{val: 9.9, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}},
				{val: 50.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}},
				{val: 70.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}},
				{val: 80.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 400.0, expectedOK: false, expectedRange: Range[float64]{}},
			},
			expectedAll: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedString: "[0, 0.9] [2, 4] [5, 9] [10, 70] [100, 200]",
		},
		{
			name: "CurrentHiAdjacentToLastHi",
			rs: []Range[float64]{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}},
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{20.0, true}, Bound[float64]{30.0, false}},
			},
			addRanges: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}},
				{Bound[float64]{7.0, false}, Bound[float64]{7.0, false}},
				{Bound[float64]{7.0, true}, Bound[float64]{8.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			equalTests: []equalTest[float64]{
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{400.0, false}},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeList[float64]{
						ranges: []Range[float64]{
							{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
							{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
							{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
							{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
							{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64]{
				{val: -1, expectedOK: false, expectedRange: Range[float64]{}},
				{val: .0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: .5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: .9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}},
				{val: 1.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}},
				{val: 5.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}},
				{val: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}},
				{val: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}},
				{val: 9.9, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}},
				{val: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}},
				{val: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}},
				{val: 50.0, expectedOK: false, expectedRange: Range[float64]{}},
				{val: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}},
				{val: 400.0, expectedOK: false, expectedRange: Range[float64]{}},
			},
			expectedAll: []Range[float64]{
				{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}},
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}},
				{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}},
				{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}},
			},
			expectedString: "[0, 0.9] [2, 4] [6, 8] [10, 30] [100, 200]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var l *RangeList[float64]

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
