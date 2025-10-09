package disc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeMap(t *testing.T) {
	type addTest[K Discrete, V any] struct {
		key Range[K]
		val V
	}

	type equalTest[K Discrete, V any] struct {
		rhs           *RangeMap[K, V]
		expectedEqual bool
	}

	type getTest[K Discrete, V any] struct {
		key           K
		expectedOK    bool
		expectedRange Range[K]
		expectedValue V
	}

	tests := []struct {
		name           string
		pairs          map[Range[int]]rune
		addTests       []addTest[int, rune]
		equalTests     []equalTest[int, rune]
		getTests       []getTest[int, rune]
		expectedAll    []generic.KeyValue[Range[int], rune]
		expectedString string
	}{
		{
			name: "CurrentHiOnLastHi_Merging",
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'a',
				{200, 400}: 'a',
			},
			addTests: []addTest[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 60}, 'A'},
				{Range[int]{60, 60}, 'A'},
				{Range[int]{60, 80}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			equalTests: []equalTest[int, rune]{
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{50, 80}, 'A'},
							{Range[int]{100, 400}, 'a'},
							{Range[int]{1000, 2000}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{50, 80}, 'A'},
							{Range[int]{100, 400}, 'a'},
							{Range[int]{1000, 2000}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int, rune]{
				{key: -10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 20, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 30, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 40, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 44, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 50, expectedOK: true, expectedRange: Range[int]{50, 80}, expectedValue: 'A'},
				{key: 60, expectedOK: true, expectedRange: Range[int]{50, 80}, expectedValue: 'A'},
				{key: 80, expectedOK: true, expectedRange: Range[int]{50, 80}, expectedValue: 'A'},
				{key: 99, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 100, expectedOK: true, expectedRange: Range[int]{100, 400}, expectedValue: 'a'},
				{key: 200, expectedOK: true, expectedRange: Range[int]{100, 400}, expectedValue: 'a'},
				{key: 400, expectedOK: true, expectedRange: Range[int]{100, 400}, expectedValue: 'a'},
				{key: 500, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 4000, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{20, 40}, Val: '@'},
				{Key: Range[int]{50, 80}, Val: 'A'},
				{Key: Range[int]{100, 400}, Val: 'a'},
				{Key: Range[int]{1000, 2000}, Val: '$'},
			},
			expectedString: "[0, 9]:35 [20, 40]:64 [50, 80]:65 [100, 400]:97 [1000, 2000]:36",
		},
		{
			name: "CurrentHiOnLastHi_Splitting",
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 200}: 'a',
				{200, 200}: 'b',
				{200, 300}: 'b',
				{300, 400}: 'c',
			},
			addTests: []addTest[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 60}, 'A'},
				{Range[int]{60, 60}, 'B'},
				{Range[int]{60, 70}, 'B'},
				{Range[int]{70, 80}, 'C'},
				{Range[int]{1000, 2000}, '$'},
			},
			equalTests: []equalTest[int, rune]{
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{50, 59}, 'A'},
							{Range[int]{60, 69}, 'B'},
							{Range[int]{70, 80}, 'C'},
							{Range[int]{100, 199}, 'a'},
							{Range[int]{200, 299}, 'b'},
							{Range[int]{300, 400}, 'c'},
							{Range[int]{1000, 2000}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
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
					expectedEqual: true,
				},
			},
			getTests: []getTest[int, rune]{
				{key: -10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 20, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 30, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 40, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 44, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 50, expectedOK: true, expectedRange: Range[int]{50, 59}, expectedValue: 'A'},
				{key: 55, expectedOK: true, expectedRange: Range[int]{50, 59}, expectedValue: 'A'},
				{key: 59, expectedOK: true, expectedRange: Range[int]{50, 59}, expectedValue: 'A'},
				{key: 60, expectedOK: true, expectedRange: Range[int]{60, 69}, expectedValue: 'B'},
				{key: 66, expectedOK: true, expectedRange: Range[int]{60, 69}, expectedValue: 'B'},
				{key: 69, expectedOK: true, expectedRange: Range[int]{60, 69}, expectedValue: 'B'},
				{key: 70, expectedOK: true, expectedRange: Range[int]{70, 80}, expectedValue: 'C'},
				{key: 77, expectedOK: true, expectedRange: Range[int]{70, 80}, expectedValue: 'C'},
				{key: 80, expectedOK: true, expectedRange: Range[int]{70, 80}, expectedValue: 'C'},
				{key: 99, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 100, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 150, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 199, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 200, expectedOK: true, expectedRange: Range[int]{200, 299}, expectedValue: 'b'},
				{key: 250, expectedOK: true, expectedRange: Range[int]{200, 299}, expectedValue: 'b'},
				{key: 299, expectedOK: true, expectedRange: Range[int]{200, 299}, expectedValue: 'b'},
				{key: 300, expectedOK: true, expectedRange: Range[int]{300, 400}, expectedValue: 'c'},
				{key: 350, expectedOK: true, expectedRange: Range[int]{300, 400}, expectedValue: 'c'},
				{key: 400, expectedOK: true, expectedRange: Range[int]{300, 400}, expectedValue: 'c'},
				{key: 500, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 4000, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{20, 40}, Val: '@'},
				{Key: Range[int]{50, 59}, Val: 'A'},
				{Key: Range[int]{60, 69}, Val: 'B'},
				{Key: Range[int]{70, 80}, Val: 'C'},
				{Key: Range[int]{100, 199}, Val: 'a'},
				{Key: Range[int]{200, 299}, Val: 'b'},
				{Key: Range[int]{300, 400}, Val: 'c'},
				{Key: Range[int]{1000, 2000}, Val: '$'},
			},
			expectedString: "[0, 9]:35 [20, 40]:64 [50, 59]:65 [60, 69]:66 [70, 80]:67 [100, 199]:97 [200, 299]:98 [300, 400]:99 [1000, 2000]:36",
		},
		{
			name: "CurrentHiBeforeLastHi_Merging",
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'a',
				{400, 600}: 'a',
				{500, 700}: 'a',
			},
			addTests: []addTest[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 80}, 'A'},
				{Range[int]{55, 65}, 'A'},
				{Range[int]{70, 80}, 'A'},
				{Range[int]{75, 90}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			equalTests: []equalTest[int, rune]{
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{50, 90}, 'A'},
							{Range[int]{100, 700}, 'a'},
							{Range[int]{1000, 2000}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{50, 90}, 'A'},
							{Range[int]{100, 700}, 'a'},
							{Range[int]{1000, 2000}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int, rune]{
				{key: -10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 20, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 30, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 40, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 44, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 50, expectedOK: true, expectedRange: Range[int]{50, 90}, expectedValue: 'A'},
				{key: 70, expectedOK: true, expectedRange: Range[int]{50, 90}, expectedValue: 'A'},
				{key: 90, expectedOK: true, expectedRange: Range[int]{50, 90}, expectedValue: 'A'},
				{key: 99, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 100, expectedOK: true, expectedRange: Range[int]{100, 700}, expectedValue: 'a'},
				{key: 500, expectedOK: true, expectedRange: Range[int]{100, 700}, expectedValue: 'a'},
				{key: 700, expectedOK: true, expectedRange: Range[int]{100, 700}, expectedValue: 'a'},
				{key: 800, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 4000, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{20, 40}, Val: '@'},
				{Key: Range[int]{50, 90}, Val: 'A'},
				{Key: Range[int]{100, 700}, Val: 'a'},
				{Key: Range[int]{1000, 2000}, Val: '$'},
			},
			expectedString: "[0, 9]:35 [20, 40]:64 [50, 90]:65 [100, 700]:97 [1000, 2000]:36",
		},
		{
			name: "CurrentHiBeforeLastHi_Splitting",
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 600}: 'a',
				{200, 300}: 'b',
				{400, 600}: 'b',
				{500, 700}: 'c',
			},
			addTests: []addTest[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{50, 80}, 'A'},
				{Range[int]{55, 65}, 'B'},
				{Range[int]{70, 80}, 'B'},
				{Range[int]{75, 90}, 'C'},
				{Range[int]{1000, 2000}, '$'},
			},
			equalTests: []equalTest[int, rune]{
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
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
							{Range[int]{1000, 2000}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
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
					expectedEqual: true,
				},
			},
			getTests: []getTest[int, rune]{
				{key: -10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 20, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 30, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 40, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 44, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 50, expectedOK: true, expectedRange: Range[int]{50, 54}, expectedValue: 'A'},
				{key: 52, expectedOK: true, expectedRange: Range[int]{50, 54}, expectedValue: 'A'},
				{key: 54, expectedOK: true, expectedRange: Range[int]{50, 54}, expectedValue: 'A'},
				{key: 55, expectedOK: true, expectedRange: Range[int]{55, 65}, expectedValue: 'B'},
				{key: 60, expectedOK: true, expectedRange: Range[int]{55, 65}, expectedValue: 'B'},
				{key: 65, expectedOK: true, expectedRange: Range[int]{55, 65}, expectedValue: 'B'},
				{key: 66, expectedOK: true, expectedRange: Range[int]{66, 69}, expectedValue: 'A'},
				{key: 68, expectedOK: true, expectedRange: Range[int]{66, 69}, expectedValue: 'A'},
				{key: 69, expectedOK: true, expectedRange: Range[int]{66, 69}, expectedValue: 'A'},
				{key: 70, expectedOK: true, expectedRange: Range[int]{70, 74}, expectedValue: 'B'},
				{key: 72, expectedOK: true, expectedRange: Range[int]{70, 74}, expectedValue: 'B'},
				{key: 74, expectedOK: true, expectedRange: Range[int]{70, 74}, expectedValue: 'B'},
				{key: 75, expectedOK: true, expectedRange: Range[int]{75, 90}, expectedValue: 'C'},
				{key: 80, expectedOK: true, expectedRange: Range[int]{75, 90}, expectedValue: 'C'},
				{key: 90, expectedOK: true, expectedRange: Range[int]{75, 90}, expectedValue: 'C'},
				{key: 99, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 100, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 150, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 199, expectedOK: true, expectedRange: Range[int]{100, 199}, expectedValue: 'a'},
				{key: 200, expectedOK: true, expectedRange: Range[int]{200, 300}, expectedValue: 'b'},
				{key: 250, expectedOK: true, expectedRange: Range[int]{200, 300}, expectedValue: 'b'},
				{key: 300, expectedOK: true, expectedRange: Range[int]{200, 300}, expectedValue: 'b'},
				{key: 301, expectedOK: true, expectedRange: Range[int]{301, 399}, expectedValue: 'a'},
				{key: 360, expectedOK: true, expectedRange: Range[int]{301, 399}, expectedValue: 'a'},
				{key: 399, expectedOK: true, expectedRange: Range[int]{301, 399}, expectedValue: 'a'},
				{key: 400, expectedOK: true, expectedRange: Range[int]{400, 499}, expectedValue: 'b'},
				{key: 444, expectedOK: true, expectedRange: Range[int]{400, 499}, expectedValue: 'b'},
				{key: 499, expectedOK: true, expectedRange: Range[int]{400, 499}, expectedValue: 'b'},
				{key: 500, expectedOK: true, expectedRange: Range[int]{500, 700}, expectedValue: 'c'},
				{key: 600, expectedOK: true, expectedRange: Range[int]{500, 700}, expectedValue: 'c'},
				{key: 700, expectedOK: true, expectedRange: Range[int]{500, 700}, expectedValue: 'c'},
				{key: 800, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 4000, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{20, 40}, Val: '@'},
				{Key: Range[int]{50, 54}, Val: 'A'},
				{Key: Range[int]{55, 65}, Val: 'B'},
				{Key: Range[int]{66, 69}, Val: 'A'},
				{Key: Range[int]{70, 74}, Val: 'B'},
				{Key: Range[int]{75, 90}, Val: 'C'},
				{Key: Range[int]{100, 199}, Val: 'a'},
				{Key: Range[int]{200, 300}, Val: 'b'},
				{Key: Range[int]{301, 399}, Val: 'a'},
				{Key: Range[int]{400, 499}, Val: 'b'},
				{Key: Range[int]{500, 700}, Val: 'c'},
				{Key: Range[int]{1000, 2000}, Val: '$'},
			},
			expectedString: "[0, 9]:35 [20, 40]:64 [50, 54]:65 [55, 65]:66 [66, 69]:65 [70, 74]:66 [75, 90]:67 [100, 199]:97 [200, 300]:98 [301, 399]:97 [400, 499]:98 [500, 700]:99 [1000, 2000]:36",
		},
		{
			name: "CurrentHiAdjacentToLastHi_Merging",
			pairs: map[Range[int]]rune{
				{20, 40}:   '@',
				{100, 199}: 'a',
				{200, 200}: 'a',
				{201, 300}: 'a',
			},
			addTests: []addTest[int, rune]{
				{Range[int]{0, 9}, '#'},
				{Range[int]{60, 69}, 'A'},
				{Range[int]{70, 70}, 'A'},
				{Range[int]{71, 80}, 'A'},
				{Range[int]{1000, 2000}, '$'},
			},
			equalTests: []equalTest[int, rune]{
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{60, 80}, 'A'},
							{Range[int]{100, 300}, 'a'},
							{Range[int]{1000, 2000}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[int, rune]{
						pairs: []rangeValue[int, rune]{
							{Range[int]{0, 9}, '#'},
							{Range[int]{20, 40}, '@'},
							{Range[int]{60, 80}, 'A'},
							{Range[int]{100, 300}, 'a'},
							{Range[int]{1000, 2000}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[int, rune]{
				{key: -10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
				{key: 10, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 20, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 30, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 40, expectedOK: true, expectedRange: Range[int]{20, 40}, expectedValue: '@'},
				{key: 50, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 60, expectedOK: true, expectedRange: Range[int]{60, 80}, expectedValue: 'A'},
				{key: 70, expectedOK: true, expectedRange: Range[int]{60, 80}, expectedValue: 'A'},
				{key: 80, expectedOK: true, expectedRange: Range[int]{60, 80}, expectedValue: 'A'},
				{key: 99, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 100, expectedOK: true, expectedRange: Range[int]{100, 300}, expectedValue: 'a'},
				{key: 200, expectedOK: true, expectedRange: Range[int]{100, 300}, expectedValue: 'a'},
				{key: 300, expectedOK: true, expectedRange: Range[int]{100, 300}, expectedValue: 'a'},
				{key: 500, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
				{key: 1000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 1500, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 2000, expectedOK: true, expectedRange: Range[int]{1000, 2000}, expectedValue: '$'},
				{key: 4000, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{20, 40}, Val: '@'},
				{Key: Range[int]{60, 80}, Val: 'A'},
				{Key: Range[int]{100, 300}, Val: 'a'},
				{Key: Range[int]{1000, 2000}, Val: '$'},
			},
			expectedString: "[0, 9]:35 [20, 40]:64 [60, 80]:65 [100, 300]:97 [1000, 2000]:36",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var m *RangeMap[int, rune]

			t.Run("NewRangeMap", func(t *testing.T) {
				eq := generic.NewEqualFunc[rune]()
				m = NewRangeMap(eq, tc.pairs)
			})

			t.Run("Clone", func(t *testing.T) {
				clone := m.Clone()
				assert.True(t, clone.Equal(m))
			})

			for i, tc := range tc.addTests {
				t.Run(fmt.Sprintf("Add/%d", i), func(t *testing.T) {
					m.Add(tc.key, tc.val)
				})
			}

			for i, tc := range tc.equalTests {
				t.Run(fmt.Sprintf("Equal/%d", i), func(t *testing.T) {
					assert.Equal(t, tc.expectedEqual, m.Equal(tc.rhs))
				})
			}

			for i, tc := range tc.getTests {
				t.Run(fmt.Sprintf("Get/%d", i), func(t *testing.T) {
					r, v, ok := m.Get(tc.key)

					assert.Equal(t, tc.expectedOK, ok)
					assert.Equal(t, tc.expectedRange, r)
					assert.Equal(t, tc.expectedValue, v)
				})
			}

			t.Run("From", func(t *testing.T) {
				all := generic.Collect2(m.All())
				assert.Equal(t, tc.expectedAll, all)
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, m.String())
			})
		})
	}
}
