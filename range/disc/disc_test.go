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

	type adjacentTest[T Discrete] struct {
		rr             Range[T]
		expectedBefore bool
		expectedAfter  bool
	}

	type intersectTest[T Discrete] struct {
		rr             Range[T]
		expectedOK     bool
		expectedResult Range[T]
	}

	tests := []struct {
		name           string
		r              Range[int]
		expectedValid  bool
		expectedString string
		equalTests     []equalTest[int]
		adjacentTests  []adjacentTest[int]
		intersectTests []intersectTest[int]
	}{
		{
			name:           "Invalid",
			r:              Range[int]{4, 2},
			expectedValid:  false,
			expectedString: "[4, 2]",
			equalTests:     nil,
			adjacentTests:  nil,
			intersectTests: nil,
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
			adjacentTests: []adjacentTest[int]{
				{rr: Range[int]{0, 0}, expectedBefore: false, expectedAfter: false},
				{rr: Range[int]{1, 1}, expectedBefore: false, expectedAfter: true},
				{rr: Range[int]{3, 3}, expectedBefore: true, expectedAfter: false},
				{rr: Range[int]{4, 5}, expectedBefore: false, expectedAfter: false},
			},
			intersectTests: []intersectTest[int]{
				{rr: Range[int]{0, 1}, expectedOK: false, expectedResult: Range[int]{}},
				{rr: Range[int]{1, 3}, expectedOK: true, expectedResult: Range[int]{2, 2}},
				{rr: Range[int]{2, 2}, expectedOK: true, expectedResult: Range[int]{2, 2}},
				{rr: Range[int]{3, 4}, expectedOK: false, expectedResult: Range[int]{}},
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
				{rr: Range[int]{0, 1}, expectedOK: false, expectedResult: Range[int]{}},
				{rr: Range[int]{0, 2}, expectedOK: true, expectedResult: Range[int]{2, 2}},
				{rr: Range[int]{1, 5}, expectedOK: true, expectedResult: Range[int]{2, 4}},
				{rr: Range[int]{2, 4}, expectedOK: true, expectedResult: Range[int]{2, 4}},
				{rr: Range[int]{3, 5}, expectedOK: true, expectedResult: Range[int]{3, 4}},
				{rr: Range[int]{5, 6}, expectedOK: false, expectedResult: Range[int]{}},
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

			for i, tc := range tc.adjacentTests {
				t.Run(fmt.Sprintf("Adjacent/%d", i), func(t *testing.T) {
					before, after := r.Adjacent(tc.rr)

					assert.Equal(t, tc.expectedBefore, before)
					assert.Equal(t, tc.expectedAfter, after)
				})
			}

			for i, tc := range tc.equalTests {
				t.Run(fmt.Sprintf("Equal/%d", i), func(t *testing.T) {
					assert.Equal(t, tc.expectedEqual, r.Equal(tc.rhs))
				})
			}

			for i, tc := range tc.intersectTests {
				t.Run(fmt.Sprintf("Intersect/%d", i), func(t *testing.T) {
					result, ok := r.Intersect(tc.rr)

					assert.Equal(t, tc.expectedOK, ok)
					assert.Equal(t, tc.expectedResult, result)
				})
			}
		})
	}
}
