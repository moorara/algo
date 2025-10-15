package cont

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	type equalTest[T Continuous] struct {
		rhs           Range[T]
		expectedEqual bool
	}

	type adjacentTest[T Continuous] struct {
		rr             Range[T]
		expectedBefore bool
		expectedAfter  bool
	}

	type intersectTest[T Continuous] struct {
		rr             Range[T]
		expectedResult RangeOrEmpty[T]
	}

	type subtractTest[T Continuous] struct {
		rr            Range[T]
		expectedLeft  RangeOrEmpty[T]
		expectedRight RangeOrEmpty[T]
	}

	tests := []struct {
		name           string
		r              Range[float64]
		expectedValid  bool
		expectedString string
		equalTests     []equalTest[float64]
		adjacentTests  []adjacentTest[float64]
		intersectTests []intersectTest[float64]
		subtractTests  []subtractTest[float64]
	}{
		{
			name: "Invalid_HiLessThanLo",
			r: Range[float64]{
				Lo: Bound[float64]{4, false},
				Hi: Bound[float64]{2, false},
			},
			expectedValid:  false,
			expectedString: "[4, 2]",
			equalTests:     nil,
			adjacentTests:  nil,
			intersectTests: nil,
			subtractTests:  nil,
		},
		{
			name: "Invalid_EqualBounds_HiBoundOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, false},
				Hi: Bound[float64]{2, true},
			},
			expectedValid:  false,
			expectedString: "[2, 2)",
			equalTests:     nil,
			adjacentTests:  nil,
			intersectTests: nil,
			subtractTests:  nil,
		},
		{
			name: "Invalid_EqualBounds_LoBoundOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, true},
				Hi: Bound[float64]{2, false},
			},
			expectedValid:  false,
			expectedString: "(2, 2]",
			equalTests:     nil,
			adjacentTests:  nil,
			intersectTests: nil,
			subtractTests:  nil,
		},
		{
			name: "Invalid_EqualBounds_BothBoundsOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, true},
				Hi: Bound[float64]{2, true},
			},
			expectedValid:  false,
			expectedString: "(2, 2)",
			equalTests:     nil,
			adjacentTests:  nil,
			intersectTests: nil,
			subtractTests:  nil,
		},
		{
			name: "EqualBounds",
			r: Range[float64]{
				Lo: Bound[float64]{2, false},
				Hi: Bound[float64]{2, false},
			},
			expectedValid:  true,
			expectedString: "[2, 2]",
			equalTests: []equalTest[float64]{
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: true,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedEqual: false,
				},
			},
			adjacentTests: []adjacentTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedBefore: false,
					expectedAfter:  true,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: true,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
			},
			intersectTests: []intersectTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
			},
			subtractTests: []subtractTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
			},
		},
		{
			name: "DiffBounds_BothBoundsClosed",
			r: Range[float64]{
				Lo: Bound[float64]{2, false},
				Hi: Bound[float64]{4, false},
			},
			expectedValid:  true,
			expectedString: "[2, 4]",
			equalTests: []equalTest[float64]{
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: true,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedEqual: false,
				},
			},
			adjacentTests: []adjacentTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedBefore: false,
					expectedAfter:  true,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: true,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
			},
			intersectTests: []intersectTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
			},
			subtractTests: []subtractTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, true},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
			},
		},
		{
			name: "DiffBounds_HiBoundOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, false},
				Hi: Bound[float64]{4, true},
			},
			expectedValid:  true,
			expectedString: "[2, 4)",
			equalTests: []equalTest[float64]{
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: true,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedEqual: false,
				},
			},
			adjacentTests: []adjacentTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedBefore: false,
					expectedAfter:  true,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: true,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
			},
			intersectTests: []intersectTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
			},
			subtractTests: []subtractTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{2, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, true},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, false},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
			},
		},
		{
			name: "DiffBounds_LoBoundOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, true},
				Hi: Bound[float64]{4, false},
			},
			expectedValid:  true,
			expectedString: "(2, 4]",
			equalTests: []equalTest[float64]{
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: true,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedEqual: false,
				},
			},
			adjacentTests: []adjacentTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedBefore: false,
					expectedAfter:  true,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: true,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
			},
			intersectTests: []intersectTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
			},
			subtractTests: []subtractTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{4, false},
							Hi: Bound[float64]{4, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, true},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
			},
		},
		{
			name: "DiffBounds_BothBoundsOpen",
			r: Range[float64]{
				Lo: Bound[float64]{2, true},
				Hi: Bound[float64]{4, true},
			},
			expectedValid:  true,
			expectedString: "(2, 4)",
			equalTests: []equalTest[float64]{
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedEqual: false,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedEqual: true,
				},
				{
					rhs: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedEqual: false,
				},
			},
			adjacentTests: []adjacentTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedBefore: false,
					expectedAfter:  true,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: true,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedBefore: false,
					expectedAfter:  false,
				},
			},
			intersectTests: []intersectTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedResult: RangeOrEmpty[float64]{Empty: true},
				},
			},
			subtractTests: []subtractTest[float64]{
				{
					rr: Range[float64]{
						Lo: Bound[float64]{0, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{1, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{1, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{2, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, true},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, false},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{2, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft:  RangeOrEmpty[float64]{Empty: true},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{3, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{3, true},
							Hi: Bound[float64]{4, true},
						},
					},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, true},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{3, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{3, false},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{4, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{4, true},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{5, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, false},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
				{
					rr: Range[float64]{
						Lo: Bound[float64]{5, true},
						Hi: Bound[float64]{6, false},
					},
					expectedLeft: RangeOrEmpty[float64]{
						Range: Range[float64]{
							Lo: Bound[float64]{2, true},
							Hi: Bound[float64]{4, true},
						},
					},
					expectedRight: RangeOrEmpty[float64]{Empty: true},
				},
			},
		},
	}

	for _, tc := range tests {
		r := tc.r

		t.Run(tc.name, func(t *testing.T) {
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

func TestCompareFuncs(t *testing.T) {
	type compareTest[T Continuous] struct {
		lhs, rhs        Bound[T]
		expectedCompare int
	}

	t.Run("compareLoLo", func(t *testing.T) {
		tests := []compareTest[float64]{
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 0,
			},
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, true},
				expectedCompare: -1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, true},
				expectedCompare: 0,
			},
		}

		for i, tc := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				assert.Equal(t, tc.expectedCompare, compareLoLo(tc.lhs, tc.rhs))
			})
		}
	})

	t.Run("compareLoHi", func(t *testing.T) {
		tests := []compareTest[float64]{
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 0,
			},
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, true},
				expectedCompare: 1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, true},
				expectedCompare: 1,
			},
		}

		for i, tc := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				assert.Equal(t, tc.expectedCompare, compareLoHi(tc.lhs, tc.rhs))
			})
		}
	})

	t.Run("compareHiLo", func(t *testing.T) {
		tests := []compareTest[float64]{
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 0,
			},
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, true},
				expectedCompare: -1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, false},
				expectedCompare: -1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, true},
				expectedCompare: -1,
			},
		}

		for i, tc := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				assert.Equal(t, tc.expectedCompare, compareHiLo(tc.lhs, tc.rhs))
			})
		}
	})

	t.Run("compareHiHi", func(t *testing.T) {
		tests := []compareTest[float64]{
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, false},
				expectedCompare: 0,
			},
			{
				lhs:             Bound[float64]{2, false},
				rhs:             Bound[float64]{2, true},
				expectedCompare: 1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, false},
				expectedCompare: -1,
			},
			{
				lhs:             Bound[float64]{2, true},
				rhs:             Bound[float64]{2, true},
				expectedCompare: 0,
			},
		}

		for i, tc := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				assert.Equal(t, tc.expectedCompare, compareHiHi(tc.lhs, tc.rhs))
			})
		}
	})
}
