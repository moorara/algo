package disc

import (
	"fmt"
	"iter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
)

func TestNewRangeMap(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()
	keep := func(a, b rune) rune {
		return a
	}
	combine := func(a, b rune) rune {
		return a + b
	}

	tests := []struct {
		name          string
		equal         generic.EqualFunc[rune]
		opts          *RangeMapOpts[int, rune]
		pairs         []RangeValue[int, rune]
		expectedPairs []RangeValue[int, rune]
	}{
		{
			name:  "CurrentHiBeforeLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 400}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{200, 400}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 100}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 300}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 400}, 'b'},
				{Range[int]{401, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{200, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 400}, 'b'},
				{Range[int]{401, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 100}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 100}, 'b'},
				{Range[int]{101, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 300}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 300}, 'b'},
				{Range[int]{301, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{200, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 100}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 300}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 400}, 'a' + 'b'},
				{Range[int]{401, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{200, 400}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 199}, 'a'},
				{Range[int]{200, 400}, 'a' + 'b'},
				{Range[int]{401, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 100}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 100}, 'a' + 'b'},
				{Range[int]{101, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 300}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 300}, 'a' + 'b'},
				{Range[int]{301, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 500}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 500}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 500}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 500}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 500}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 499}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 500}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 499}, 'a'},
				{Range[int]{500, 500}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 500}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 700}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 700}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 700}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 700}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 700}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 700}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 700}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 700}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 499}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: keep},
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{100, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a' + 'b'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{300, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 500}, 'a' + 'b'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 499}, 'a'},
				{Range[int]{500, 500}, 'a' + 'b'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[int, rune]{Resolve: combine},
			pairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a'},
				{Range[int]{500, 700}, 'b'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{500, 500}, 'a' + 'b'},
				{Range[int]{501, 700}, 'b'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 100}, 'a'},
				{Range[int]{101, 300}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 300}, 'a'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{201, 300}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 300}, 'a'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 299}, 'a'},
				{Range[int]{300, 300}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 300}, 'a'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{299, 299}, 'a'},
				{Range[int]{300, 300}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{299, 300}, 'a'},
			},
		},
		{
			name:  "DisjointRanges",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'a'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'a'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewRangeMap(tc.equal, tc.opts, tc.pairs).(*rangeMap[int, rune])

			assert.Equal(t, tc.expectedPairs, m.pairs)
		})
	}
}

func TestRangeMap_String(t *testing.T) {
	tests := []struct {
		name           string
		m              *rangeMap[int, rune]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			m: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedString: "[20, 40]:64 [100, 199]:97 [200, 299]:98 [300, 400]:99",
		},
		{
			name: "WithCustomFormat",
			m: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal: generic.NewEqualFunc[rune](),
				format: func(ranges iter.Seq2[Range[int], rune]) string {
					strs := make([]string, 0)
					for r, v := range ranges {
						strs = append(strs, fmt.Sprintf("[%d..%d] --> %c", r.Lo, r.Hi, v))
					}
					return strings.Join(strs, "\n")
				},
				resolve: defaultResolve[rune],
			},
			expectedString: "[20..40] --> @\n[100..199] --> a\n[200..299] --> b\n[300..400] --> c",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.m.String())
		})
	}
}

func TestRangeMap_Clone(t *testing.T) {
	tests := []struct {
		name string
		m    *rangeMap[int, rune]
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clone := tc.m.Clone()

			assert.True(t, clone.Equal(tc.m))
		})
	}
}

func TestRangeMap_Equal(t *testing.T) {
	m := &rangeMap[int, rune]{
		pairs: []RangeValue[int, rune]{
			{Range[int]{20, 40}, '@'},
			{Range[int]{100, 199}, 'a'},
			{Range[int]{200, 299}, 'b'},
			{Range[int]{300, 400}, 'c'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[int, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		rhs           RangeMap[int, rune]
		expectedEqual bool
	}{
		{
			name:          "NotEqual_DiffTypes",
			m:             m,
			rhs:           nil,
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffLens",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{10, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffValues",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '*'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			m:    m,
			rhs: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.m.Equal(tc.rhs))
		})
	}
}

func TestRangeMap_Size(t *testing.T) {
	tests := []struct {
		name         string
		m            *rangeMap[int, rune]
		expectedSize int
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{20, 40}, '@'},
					{Range[int]{100, 199}, 'a'},
					{Range[int]{200, 299}, 'b'},
					{Range[int]{300, 400}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedSize: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedSize, tc.m.Size())
		})
	}
}

func TestRangeMap_Find(t *testing.T) {
	m := &rangeMap[int, rune]{
		pairs: []RangeValue[int, rune]{
			{Range[int]{0, 9}, '#'},
			{Range[int]{10, 20}, '@'},
			{Range[int]{200, 400}, 'a'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[int, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		m             *rangeMap[int, rune]
		val           int
		expectedOK    bool
		expectedRange Range[int]
		expectedValue rune
	}{
		{m: m, val: -1, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
		{m: m, val: 0, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 5, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 9, expectedOK: true, expectedRange: Range[int]{0, 9}, expectedValue: '#'},
		{m: m, val: 10, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 15, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 20, expectedOK: true, expectedRange: Range[int]{10, 20}, expectedValue: '@'},
		{m: m, val: 100, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
		{m: m, val: 200, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 300, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 400, expectedOK: true, expectedRange: Range[int]{200, 400}, expectedValue: 'a'},
		{m: m, val: 500, expectedOK: false, expectedRange: Range[int]{}, expectedValue: 0},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			r, v, ok := tc.m.Find(tc.val)

			assert.Equal(t, tc.expectedOK, ok)
			assert.Equal(t, tc.expectedRange, r)
			assert.Equal(t, tc.expectedValue, v)
		})
	}
}

func TestRangeMap_Add(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()
	keep := func(a, b rune) rune {
		return a
	}
	combine := func(a, b rune) rune {
		return a + b
	}

	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		pairs         []RangeValue[int, rune]
		expectedPairs []RangeValue[int, rune]
	}{
		{
			name: "CurrentHiBeforeLastHi/Case01/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 40}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{20, 40}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 10}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 30}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 40}, 'B'},
				{Range[int]{41, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{20, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 19}, 'A'},
				{Range[int]{20, 40}, 'B'},
				{Range[int]{41, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 10}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 10}, 'B'},
				{Range[int]{11, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 30}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 30}, 'B'},
				{Range[int]{31, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{20, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 10}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 30}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 40}, 'A' + 'B'},
				{Range[int]{41, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{20, 40}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 19}, 'A'},
				{Range[int]{20, 40}, 'A' + 'B'},
				{Range[int]{41, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 10}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 10}, 'A' + 'B'},
				{Range[int]{11, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 30}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 30}, 'A' + 'B'},
				{Range[int]{31, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 50}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 50}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 50}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 50}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 50}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 49}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 50}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 49}, 'A'},
				{Range[int]{50, 50}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 50}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 70}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 70}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 70}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 70}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 70}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 70}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 70}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 70}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 49}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: keep,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{10, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A' + 'B'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{30, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 50}, 'A' + 'B'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 49}, 'A'},
				{Range[int]{50, 50}, 'A' + 'B'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: combine,
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A'},
				{Range[int]{50, 70}, 'B'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{50, 50}, 'A' + 'B'},
				{Range[int]{51, 70}, 'B'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case01/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 10}, 'A'},
				{Range[int]{11, 30}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 30}, 'A'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case02/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 20}, 'A'},
				{Range[int]{21, 30}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 30}, 'A'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case03/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 29}, 'A'},
				{Range[int]{30, 30}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 30}, 'A'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case04/EqualValues",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{29, 29}, 'A'},
				{Range[int]{30, 30}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{29, 30}, 'A'},
			},
		},
		{
			name: "DisjointRanges",
			m: &rangeMap[int, rune]{
				pairs:   []RangeValue[int, rune]{},
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[int, rune]{
				{Range[int]{10, 20}, 'A'},
				{Range[int]{30, 40}, 'A'},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{10, 20}, 'A'},
				{Range[int]{30, 40}, 'A'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, p := range tc.pairs {
				tc.m.Add(p.Range, p.Value)
			}

			assert.Equal(t, tc.expectedPairs, tc.m.pairs)
		})
	}
}

func TestRangeMap_Remove(t *testing.T) {
	equal := generic.NewEqualFunc[rune]()
	pairs := []RangeValue[int, rune]{
		{Range[int]{100, 200}, 'a'},
		{Range[int]{300, 400}, 'b'},
		{Range[int]{500, 600}, 'c'},
		{Range[int]{700, 800}, 'd'},
		{Range[int]{900, 1000}, 'e'},
	}

	tests := []struct {
		name          string
		m             *rangeMap[int, rune]
		keys          []Range[int]
		expectedPairs []RangeValue[int, rune]
	}{
		{
			name: "None",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: nil,
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'b'},
				{Range[int]{500, 600}, 'c'},
				{Range[int]{700, 800}, 'd'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{40, 60},
				{240, 260},
				{440, 460},
				{640, 660},
				{840, 860},
				{1040, 1060},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{300, 400}, 'b'},
				{Range[int]{500, 600}, 'c'},
				{Range[int]{700, 800}, 'd'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{80, 100},
				{200, 300},
				{400, 500},
				{600, 700},
				{800, 900},
				{1000, 1020},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{101, 199}, 'a'},
				{Range[int]{301, 399}, 'b'},
				{Range[int]{501, 599}, 'c'},
				{Range[int]{701, 799}, 'd'},
				{Range[int]{901, 999}, 'e'},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{80, 120},
				{180, 320},
				{380, 520},
				{580, 720},
				{780, 920},
				{980, 1020},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{121, 179}, 'a'},
				{Range[int]{321, 379}, 'b'},
				{Range[int]{521, 579}, 'c'},
				{Range[int]{721, 779}, 'd'},
				{Range[int]{921, 979}, 'e'},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{140, 160},
				{340, 360},
				{540, 560},
				{740, 760},
				{940, 960},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 139}, 'a'},
				{Range[int]{161, 200}, 'a'},
				{Range[int]{300, 339}, 'b'},
				{Range[int]{361, 400}, 'b'},
				{Range[int]{500, 539}, 'c'},
				{Range[int]{561, 600}, 'c'},
				{Range[int]{700, 739}, 'd'},
				{Range[int]{761, 800}, 'd'},
				{Range[int]{900, 939}, 'e'},
				{Range[int]{961, 1000}, 'e'},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{250, 450},
				{450, 850},
			},
			expectedPairs: []RangeValue[int, rune]{
				{Range[int]{100, 200}, 'a'},
				{Range[int]{900, 1000}, 'e'},
			},
		},
		{
			name: "All",
			m: &rangeMap[int, rune]{
				pairs:   append([]RangeValue[int, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[int]{
				{100, 200},
				{300, 400},
				{500, 600},
				{700, 800},
				{900, 1000},
			},
			expectedPairs: []RangeValue[int, rune]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for _, k := range tc.keys {
				tc.m.Remove(k)
			}

			assert.Equal(t, tc.expectedPairs, tc.m.pairs)
		})
	}
}

func TestRangeMap_All(t *testing.T) {
	tests := []struct {
		name        string
		m           *rangeMap[int, rune]
		expectedAll []generic.KeyValue[Range[int], rune]
	}{
		{
			name: "OK",
			m: &rangeMap[int, rune]{
				pairs: []RangeValue[int, rune]{
					{Range[int]{0, 9}, '#'},
					{Range[int]{10, 20}, '@'},
					{Range[int]{200, 400}, 'a'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[int, rune],
				resolve: defaultResolve[rune],
			},
			expectedAll: []generic.KeyValue[Range[int], rune]{
				{Key: Range[int]{0, 9}, Val: '#'},
				{Key: Range[int]{10, 20}, Val: '@'},
				{Key: Range[int]{200, 400}, Val: 'a'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			all := generic.Collect2(tc.m.All())

			assert.Equal(t, tc.expectedAll, all)
		})
	}
}
