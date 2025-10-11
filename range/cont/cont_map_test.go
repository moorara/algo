package cont

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestRangeMap(t *testing.T) {
	type addTest[K Continuous, V any] struct {
		key Range[K]
		val V
	}

	type equalTest[K Continuous, V any] struct {
		rhs           *RangeMap[K, V]
		expectedEqual bool
	}

	type getTest[K Continuous, V any] struct {
		key           K
		expectedOK    bool
		expectedRange Range[K]
		expectedValue V
	}

	tests := []struct {
		name           string
		pairs          map[Range[float64]]rune
		addTests       []addTest[float64, rune]
		equalTests     []equalTest[float64, rune]
		getTests       []getTest[float64, rune]
		expectedAll    []generic.KeyValue[Range[float64], rune]
		expectedString string
	}{
		{
			name: "CurrentHiOnLastHi_Merging",
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}: 'a',
			},
			addTests: []addTest[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			equalTests: []equalTest[float64, rune]{
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64, rune]{
				{key: -1, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.4, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 9.9, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'a'},
				{key: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'a'},
				{key: 40.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'a'},
				{key: 50.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 400.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, Val: '$'},
			},
			expectedString: "[0, 0.9]:35 [2, 4]:64 [5, 8]:65 [10, 40]:97 [100, 200]:36",
		},
		{
			name: "CurrentHiOnLastHi_Splitting",
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'b',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}: 'c',
			},
			addTests: []addTest[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{6.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			equalTests: []equalTest[float64, rune]{
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, 'A'},
							{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, 'B'},
							{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'C'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
							{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
							{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, 'A'},
							{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, 'B'},
							{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'C'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
							{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
							{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64, rune]{
				{key: -1, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.4, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, expectedValue: 'A'},
				{key: 5.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, expectedValue: 'A'},
				{key: 5.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, expectedValue: 'A'},
				{key: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, expectedValue: 'B'},
				{key: 6.6, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, expectedValue: 'B'},
				{key: 6.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, expectedValue: 'B'},
				{key: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'C'},
				{key: 7.7, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'C'},
				{key: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'C'},
				{key: 9.9, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 15.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 19.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, expectedValue: 'b'},
				{key: 25.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, expectedValue: 'b'},
				{key: 29.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, expectedValue: 'b'},
				{key: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'c'},
				{key: 35.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'c'},
				{key: 40.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, expectedValue: 'c'},
				{key: 50.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 400.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{6.0, true}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, Val: 'B'},
				{Key: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, Val: 'C'},
				{Key: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, Val: 'b'},
				{Key: Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, Val: 'c'},
				{Key: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, Val: '$'},
			},
			expectedString: "[0, 0.9]:35 [2, 4]:64 [5, 6):65 [6, 7):66 [7, 8]:67 [10, 20):97 [20, 30):98 [30, 40]:99 [100, 200]:36",
		},
		{
			name: "CurrentHiBeforeLastHi_Merging",
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'a',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'a',
			},
			addTests: []addTest[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			equalTests: []equalTest[float64, rune]{
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64, rune]{
				{key: -1, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.4, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, expectedValue: 'A'},
				{key: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, expectedValue: 'A'},
				{key: 9.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, expectedValue: 'A'},
				{key: 9.9, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'a'},
				{key: 50.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'a'},
				{key: 70.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'a'},
				{key: 80.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 400.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{9.0, false}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, Val: '$'},
			},
			expectedString: "[0, 0.9]:35 [2, 4]:64 [5, 9]:65 [10, 70]:97 [100, 200]:36",
		},
		{
			name: "CurrentHiBeforeLastHi_Splitting",
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{60.0, false}}: 'a',
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}: 'b',
				{Bound[float64]{40.0, false}, Bound[float64]{60.0, false}}: 'b',
				{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}: 'c',
			},
			addTests: []addTest[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{8.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'C'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			equalTests: []equalTest[float64, rune]{
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, 'A'},
							{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'B'},
							{Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, 'A'},
							{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, 'B'},
							{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'C'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
							{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
							{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
							{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
							{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, 'A'},
							{Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, 'B'},
							{Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, 'A'},
							{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, 'B'},
							{Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, 'C'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, 'a'},
							{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, 'b'},
							{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'a'},
							{Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, 'b'},
							{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'c'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64, rune]{
				{key: -1, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.4, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 5.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, expectedValue: 'A'},
				{key: 5.2, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, expectedValue: 'A'},
				{key: 5.4, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, expectedValue: 'A'},
				{key: 5.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, expectedValue: 'B'},
				{key: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, expectedValue: 'B'},
				{key: 6.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, expectedValue: 'B'},
				{key: 6.6, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, expectedValue: 'A'},
				{key: 6.8, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, expectedValue: 'A'},
				{key: 6.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, expectedValue: 'A'},
				{key: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, expectedValue: 'B'},
				{key: 7.2, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, expectedValue: 'B'},
				{key: 7.4, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, expectedValue: 'B'},
				{key: 7.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, expectedValue: 'C'},
				{key: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, expectedValue: 'C'},
				{key: 9.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, expectedValue: 'C'},
				{key: 9.9, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 15.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 19.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, expectedValue: 'a'},
				{key: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'b'},
				{key: 25.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'b'},
				{key: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'b'},
				{key: 30.1, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
				{key: 36.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
				{key: 39.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
				{key: 40.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, expectedValue: 'b'},
				{key: 44.4, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, expectedValue: 'b'},
				{key: 49.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, expectedValue: 'b'},
				{key: 50.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'c'},
				{key: 60.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'c'},
				{key: 70.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, expectedValue: 'c'},
				{key: 80.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 400.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.5, true}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{5.5, false}, Bound[float64]{6.5, false}}, Val: 'B'},
				{Key: Range[float64]{Bound[float64]{6.5, true}, Bound[float64]{7.0, true}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.5, true}}, Val: 'B'},
				{Key: Range[float64]{Bound[float64]{7.5, false}, Bound[float64]{9.0, false}}, Val: 'C'},
				{Key: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20, true}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}}, Val: 'b'},
				{Key: Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{40.0, false}, Bound[float64]{50.0, true}}, Val: 'b'},
				{Key: Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, Val: 'c'},
				{Key: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, Val: '$'},
			},
			expectedString: "[0, 0.9]:35 [2, 4]:64 [5, 5.5):65 [5.5, 6.5]:66 (6.5, 7):65 [7, 7.5):66 [7.5, 9]:67 [10, 20):97 [20, 30]:98 (30, 40):97 [40, 50):98 [50, 70]:99 [100, 200]:36",
		},
		{
			name: "CurrentHiAdjacentToLastHi_Merging",
			pairs: map[Range[float64]]rune{
				{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}:   '@',
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}:  'a',
				{Bound[float64]{20.0, false}, Bound[float64]{20.0, false}}: 'a',
				{Bound[float64]{20, true}, Bound[float64]{30.0, false}}:    'a',
			},
			addTests: []addTest[float64, rune]{
				{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
				{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{7.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, false}, Bound[float64]{7.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{7.0, true}, Bound[float64]{8.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
			},
			equalTests: []equalTest[float64, rune]{
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '%'},
						},
					},
					expectedEqual: false,
				},
				{
					rhs: &RangeMap[float64, rune]{
						pairs: []rangeValue[float64, rune]{
							{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
							{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
							{Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, 'A'},
							{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
							{Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, '$'},
						},
					},
					expectedEqual: true,
				},
			},
			getTests: []getTest[float64, rune]{
				{key: -1, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
				{key: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 3.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 4.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, expectedValue: '@'},
				{key: 5.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 6.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 7.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 8.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, expectedValue: 'A'},
				{key: 9.9, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 10.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'a'},
				{key: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'a'},
				{key: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, expectedValue: 'a'},
				{key: 50.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
				{key: 100.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 150.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 200.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, expectedValue: '$'},
				{key: 400.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{6.0, false}, Bound[float64]{8.0, false}}, Val: 'A'},
				{Key: Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{100.0, false}, Bound[float64]{200.0, false}}, Val: '$'},
			},
			expectedString: "[0, 0.9]:35 [2, 4]:64 [6, 8]:65 [10, 30]:97 [100, 200]:36",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var m *RangeMap[float64, rune]

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

			t.Run("All", func(t *testing.T) {
				all := generic.Collect2(m.All())
				assert.Equal(t, tc.expectedAll, all)
			})

			t.Run("String", func(t *testing.T) {
				assert.Equal(t, tc.expectedString, m.String())
			})
		})
	}
}
