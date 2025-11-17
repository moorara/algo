package disc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	type equalTest[T Discrete] struct {
		rhs           Range[T]
		expectedEqual bool
	}

	type includesTest[T Discrete] struct {
		val              T
		expectedIncludes bool
	}

	type adjacentTest[T Discrete] struct {
		rr             Range[T]
		expectedBefore bool
		expectedAfter  bool
	}

	type intersectTest[T Discrete] struct {
		rr             Range[T]
		expectedResult RangeOrEmpty[T]
	}

	type subtractTest[T Discrete] struct {
		rr            Range[T]
		expectedLeft  RangeOrEmpty[T]
		expectedRight RangeOrEmpty[T]
	}

	tests := []struct {
		name           string
		r              Range[int]
		expectedValid  bool
		expectedString string
		equalTests     []equalTest[int]
		includesTests  []includesTest[int]
		adjacentTests  []adjacentTest[int]
		intersectTests []intersectTest[int]
		subtractTests  []subtractTest[int]
	}{
		{
			name:           "Invalid",
			r:              Range[int]{4, 2},
			expectedValid:  false,
			expectedString: "[4, 2]",
			equalTests:     nil,
			includesTests:  nil,
			adjacentTests:  nil,
			intersectTests: nil,
			subtractTests:  nil,
		},
		{
			name:           "EqualBounds",
			r:              Range[int]{2, 2},
			expectedValid:  true,
			expectedString: "[2, 2]",
			equalTests: []equalTest[int]{
				{rhs: Range[int]{1, 2}, expectedEqual: false},
				{rhs: Range[int]{2, 2}, expectedEqual: true},
				{rhs: Range[int]{2, 3}, expectedEqual: false},
			},
			includesTests: []includesTest[int]{
				{val: 1, expectedIncludes: false},
				{val: 2, expectedIncludes: true},
				{val: 3, expectedIncludes: false},
			},
			adjacentTests: []adjacentTest[int]{
				{rr: Range[int]{0, 0}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{1, 1}, expectedBefore: false, expectedAfter: true},
				{rr: Range[int]{3, 3}, expectedBefore: true, expectedAfter: false},
				{rr: Range[int]{4, 5}, expectedBefore: false, expectedAfter: false},
			},
			intersectTests: []intersectTest[int]{
				{rr: Range[int]{0, 1}, expectedResult: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{1, 3}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{2, 2}}},
				{rr: Range[int]{2, 2}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{2, 2}}},
				{rr: Range[int]{3, 4}, expectedResult: RangeOrEmpty[int]{Empty: true}},
			},
			subtractTests: []subtractTest[int]{
				{rr: Range[int]{0, 1}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{2, 2}}},
				{rr: Range[int]{1, 2}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{2, 2}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{2, 3}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{3, 4}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 2}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
			},
		},
		{
			name:           "DiffBounds",
			r:              Range[int]{2, 4},
			expectedValid:  true,
			expectedString: "[2, 4]",
			equalTests: []equalTest[int]{
				{rhs: Range[int]{1, 4}, expectedEqual: false},
				{rhs: Range[int]{2, 4}, expectedEqual: true},
				{rhs: Range[int]{2, 5}, expectedEqual: false},
			},
			includesTests: []includesTest[int]{
				{val: 1, expectedIncludes: false},
				{val: 2, expectedIncludes: true},
				{val: 3, expectedIncludes: true},
				{val: 4, expectedIncludes: true},
				{val: 5, expectedIncludes: false},
			},
			adjacentTests: []adjacentTest[int]{
				{rr: Range[int]{0, 0}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{0, 1}, expectedBefore: false, expectedAfter: true},
				{rr: Range[int]{0, 2}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{3, 3}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{4, 5}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{5, 6}, expectedBefore: true, expectedAfter: false},
				{rr: Range[int]{6, 6}, expectedBefore: false, expectedAfter: false},
			},
			intersectTests: []intersectTest[int]{
				{rr: Range[int]{0, 1}, expectedResult: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{0, 2}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{2, 2}}},
				{rr: Range[int]{1, 5}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{2, 4}}},
				{rr: Range[int]{2, 4}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{2, 4}}},
				{rr: Range[int]{3, 5}, expectedResult: RangeOrEmpty[int]{Range: Range[int]{3, 4}}},
				{rr: Range[int]{5, 6}, expectedResult: RangeOrEmpty[int]{Empty: true}},
			},
			subtractTests: []subtractTest[int]{
				{rr: Range[int]{0, 1}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{2, 4}}},
				{rr: Range[int]{1, 1}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{2, 4}}},
				{rr: Range[int]{1, 2}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{3, 4}}},
				{rr: Range[int]{1, 3}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{4, 4}}},
				{rr: Range[int]{1, 4}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{1, 5}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{2, 2}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{3, 4}}},
				{rr: Range[int]{2, 3}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{4, 4}}},
				{rr: Range[int]{2, 4}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{2, 5}, expectedLeft: RangeOrEmpty[int]{Empty: true}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{3, 3}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 2}}, expectedRight: RangeOrEmpty[int]{Range: Range[int]{4, 4}}},
				{rr: Range[int]{3, 4}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 2}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{3, 5}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 2}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{4, 4}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 3}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{4, 5}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 3}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{5, 5}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 4}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
				{rr: Range[int]{5, 6}, expectedLeft: RangeOrEmpty[int]{Range: Range[int]{2, 4}}, expectedRight: RangeOrEmpty[int]{Empty: true}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.r

			t.Run("Valid", func(t *testing.T) {
				assert.Equal(t, tc.expectedValid, r.Valid())
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, r.String())
			})

			for i, tc := range tc.equalTests {
				t.Run(fmt.Sprintf("Equal/%d", i), func(t *testing.T) {
					assert.Equal(t, tc.expectedEqual, r.Equal(tc.rhs))
				})
			}

			for i, tc := range tc.includesTests {
				t.Run(fmt.Sprintf("Includes/%d", i), func(t *testing.T) {
					assert.Equal(t, tc.expectedIncludes, r.Includes(tc.val))
				})
			}

			for i, tc := range tc.adjacentTests {
				t.Run(fmt.Sprintf("Adjacent/%d", i), func(t *testing.T) {
					before, after := r.Adjacent(tc.rr)

					assert.Equal(t, tc.expectedBefore, before)
					assert.Equal(t, tc.expectedAfter, after)
				})
			}

			for i, tc := range tc.intersectTests {
				t.Run(fmt.Sprintf("Intersect/%d", i), func(t *testing.T) {
					res := r.Intersect(tc.rr)

					assert.Equal(t, tc.expectedResult, res)
				})
			}

			for i, tc := range tc.subtractTests {
				t.Run(fmt.Sprintf("Subtract/%d", i), func(t *testing.T) {
					left, right := r.Subtract(tc.rr)

					assert.Equal(t, tc.expectedLeft, left)
					assert.Equal(t, tc.expectedRight, right)
				})
			}
		})
	}
}

func TestEqRange(t *testing.T) {
	tests := []struct {
		name          string
		lhs, rhs      Range[int]
		expectedEqual bool
	}{
		{lhs: Range[int]{2, 4}, rhs: Range[int]{3, 4}, expectedEqual: false},
		{lhs: Range[int]{2, 4}, rhs: Range[int]{2, 4}, expectedEqual: true},
		{lhs: Range[int]{2, 4}, rhs: Range[int]{1, 4}, expectedEqual: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, EqRange(tc.lhs, tc.rhs))
		})
	}
}

func TestCmpRange(t *testing.T) {
	tests := []struct {
		name            string
		lhs, rhs        Range[int]
		expectedCompare int
	}{
		{lhs: Range[int]{2, 4}, rhs: Range[int]{3, 4}, expectedCompare: -1},
		{lhs: Range[int]{2, 4}, rhs: Range[int]{2, 4}, expectedCompare: 0},
		{lhs: Range[int]{2, 4}, rhs: Range[int]{1, 4}, expectedCompare: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCompare, CmpRange(tc.lhs, tc.rhs))
		})
	}
}
