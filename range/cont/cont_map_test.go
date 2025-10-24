package cont

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
		opts          *RangeMapOpts[float64, rune]
		pairs         []RangeValue[float64, rune]
		expectedPairs []RangeValue[float64, rune]
	}{
		{
			name:  "CurrentHiBeforeLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{10.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{40.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{10.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/DefaultResolver",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: keep},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{50.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			equal: equal,
			opts:  &RangeMapOpts[float64, rune]{Resolve: combine},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{70.0, false}}, 'b'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{50.0, false}}, 'a' + 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{70.0, false}}, 'b'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case01/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{10.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{10.0, true}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case02/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{20.0, true}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
		},
		{
			name:  "CurrentHiAdjacentToLastHi/Case03/EqualValues",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{30.0, false}}, 'a'},
			},
		},
		{
			name:  "DisjointRanges",
			equal: equal,
			opts:  nil,
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'a'},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := NewRangeMap(tc.equal, tc.opts, tc.pairs...).(*rangeMap[float64, rune])

			assert.Equal(t, tc.expectedPairs, m.pairs)
		})
	}
}

func TestRangeMap_String(t *testing.T) {
	tests := []struct {
		name           string
		m              *rangeMap[float64, rune]
		expectedString string
	}{
		{
			name: "WithDefaultFormat",
			m: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			expectedString: "[2, 4]:64 [10, 20):97 [20, 30):98 [30, 40]:99",
		},
		{
			name: "WithCustomFormat",
			m: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal: generic.NewEqualFunc[rune](),
				format: func(ranges iter.Seq2[Range[float64], rune]) string {
					strs := make([]string, 0)
					for r, v := range ranges {
						strs = append(strs, fmt.Sprintf("%s --> %c", r.String(), v))
					}
					return strings.Join(strs, "\n")
				},
				resolve: defaultResolve[rune],
			},
			expectedString: "[2, 4] --> @\n[10, 20) --> a\n[20, 30) --> b\n[30, 40] --> c",
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
		m    *rangeMap[float64, rune]
	}{
		{
			name: "OK",
			m: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
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
	m := &rangeMap[float64, rune]{
		pairs: []RangeValue[float64, rune]{
			{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
			{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
			{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
			{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[float64, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		name          string
		m             *rangeMap[float64, rune]
		rhs           RangeMap[float64, rune]
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
			rhs: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffRanges",
			m:    m,
			rhs: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "NotEqual_DiffValues",
			m:    m,
			rhs: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '*'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			expectedEqual: false,
		},
		{
			name: "Equal",
			m:    m,
			rhs: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
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
		m            *rangeMap[float64, rune]
		expectedSize int
	}{
		{
			name: "OK",
			m: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, '@'},
					{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{30.0, true}}, 'b'},
					{Range[float64]{Bound[float64]{30.0, false}, Bound[float64]{40.0, false}}, 'c'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
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
	m := &rangeMap[float64, rune]{
		pairs: []RangeValue[float64, rune]{
			{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
			{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, '@'},
			{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, 'a'},
			{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, 'b'},
		},
		equal:   generic.NewEqualFunc[rune](),
		format:  defaultFormatMap[float64, rune],
		resolve: defaultResolve[rune],
	}

	tests := []struct {
		m             *rangeMap[float64, rune]
		val           float64
		expectedOK    bool
		expectedRange Range[float64]
		expectedValue rune
	}{
		{m: m, val: -1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, val: 0.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, val: 0.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, val: 0.9, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, expectedValue: '#'},
		{m: m, val: 1.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, val: 1.5, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, expectedValue: '@'},
		{m: m, val: 2.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, expectedValue: '@'},
		{m: m, val: 10.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, val: 20.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
		{m: m, val: 30.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, expectedValue: 'a'},
		{m: m, val: 40.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
		{m: m, val: 60.0, expectedOK: true, expectedRange: Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, expectedValue: 'b'},
		{m: m, val: 80.0, expectedOK: false, expectedRange: Range[float64]{}, expectedValue: 0},
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
		m             *rangeMap[float64, rune]
		pairs         []RangeValue[float64, rune]
		expectedPairs []RangeValue[float64, rune]
	}{
		{
			name: "CurrentHiBeforeLastHi/Case01/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{4.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{2.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{4.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'B'},
				{Range[float64]{Bound[float64]{3.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{4.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{4.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{2.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, false}, Bound[float64]{4.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{4.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiBeforeLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{3.0, true}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiOnLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/DefaultResolver",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveExistingValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: keep,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case01/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case02/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case03/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{5.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAfterLastHi/Case04/DiffValues/CustomResolver/ResolveCombinedValue",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: combine,
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{7.0, false}}, 'B'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{5.0, false}, Bound[float64]{5.0, false}}, 'A' + 'B'},
				{Range[float64]{Bound[float64]{5.0, true}, Bound[float64]{7.0, false}}, 'B'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case01/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{1.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{3.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case02/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{2.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{2.0, true}, Bound[float64]{3.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, false}}, 'A'},
			},
		},
		{
			name: "CurrentHiAdjacentToLastHi/Case03/EqualValues",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, true}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{3.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{3.0, false}}, 'A'},
			},
		},
		{
			name: "DisjointRanges",
			m: &rangeMap[float64, rune]{
				pairs:   []RangeValue[float64, rune]{},
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			pairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{2.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{4.0, false}}, 'A'},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{1.0, false}, Bound[float64]{2.0, false}}, 'A'},
				{Range[float64]{Bound[float64]{3.0, false}, Bound[float64]{4.0, false}}, 'A'},
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
	pairs := []RangeValue[float64, rune]{
		{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
		{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}, 'b'},
		{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}, 'c'},
		{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
		{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
	}

	tests := []struct {
		name          string
		m             *rangeMap[float64, rune]
		keys          []Range[float64]
		expectedPairs []RangeValue[float64, rune]
	}{
		{
			name: "None",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: nil,
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: No Overlapping
			//
			//        |________|        |________|        |________|        |________|        |________|
			//  |__|              |__|              |__|              |__|              |__|              |__|
			//
			name: "NoOverlapping",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{4.0, false}, Bound[float64]{6.0, false}},
				{Bound[float64]{24.0, false}, Bound[float64]{26.0, false}},
				{Bound[float64]{44.0, false}, Bound[float64]{46.0, false}},
				{Bound[float64]{64.0, false}, Bound[float64]{66.0, false}},
				{Bound[float64]{84.0, false}, Bound[float64]{86.0, false}},
				{Bound[float64]{104.0, false}, Bound[float64]{106.0, false}},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: Overlapping Bounds
			//
			//        |________|        |________|        |________|        |________|        |________|
			//     |__|        |________|        |________|        |________|        |________|        |__|
			//
			name: "OverlappingBounds",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{10.0, false}},
				{Bound[float64]{20.0, false}, Bound[float64]{30.0, false}},
				{Bound[float64]{40.0, false}, Bound[float64]{50.0, false}},
				{Bound[float64]{60.0, false}, Bound[float64]{70.0, false}},
				{Bound[float64]{80.0, false}, Bound[float64]{90.0, false}},
				{Bound[float64]{100.0, false}, Bound[float64]{102.0, false}},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, true}, Bound[float64]{20.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{40.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, true}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, true}, Bound[float64]{100.0, true}}, 'e'},
			},
		},
		{
			// Case: Overlapping Ranges
			//
			//        |________|        |________|        |________|        |________|        |________|
			//      |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|    |___|
			//
			name: "OverlappingRanges",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{8.0, false}, Bound[float64]{12.0, false}},
				{Bound[float64]{18.0, false}, Bound[float64]{32.0, false}},
				{Bound[float64]{38.0, false}, Bound[float64]{52.0, false}},
				{Bound[float64]{58.0, false}, Bound[float64]{72.0, false}},
				{Bound[float64]{78.0, false}, Bound[float64]{92.0, false}},
				{Bound[float64]{98.0, false}, Bound[float64]{102.0, false}},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{12.0, true}, Bound[float64]{18.0, true}}, 'a'},
				{Range[float64]{Bound[float64]{32.0, true}, Bound[float64]{38.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{52.0, true}, Bound[float64]{58.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{72.0, true}, Bound[float64]{78.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{92.0, true}, Bound[float64]{98.0, true}}, 'e'},
			},
		},
		{
			// Case: Subsets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//           |__|              |__|              |__|              |__|              |__|
			//
			name: "Subsets",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{14.0, true}, Bound[float64]{16.0, true}},
				{Bound[float64]{34.0, true}, Bound[float64]{36.0, true}},
				{Bound[float64]{54.0, true}, Bound[float64]{56.0, true}},
				{Bound[float64]{74.0, true}, Bound[float64]{76.0, true}},
				{Bound[float64]{94.0, true}, Bound[float64]{96.0, true}},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{14.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{16.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{30.0, true}, Bound[float64]{34, false}}, 'b'},
				{Range[float64]{Bound[float64]{36.0, false}, Bound[float64]{40.0, false}}, 'b'},
				{Range[float64]{Bound[float64]{50.0, false}, Bound[float64]{54.0, false}}, 'c'},
				{Range[float64]{Bound[float64]{56.0, false}, Bound[float64]{60.0, true}}, 'c'},
				{Range[float64]{Bound[float64]{70.0, true}, Bound[float64]{74.0, false}}, 'd'},
				{Range[float64]{Bound[float64]{76.0, false}, Bound[float64]{80.0, true}}, 'd'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{94.0, false}}, 'e'},
				{Range[float64]{Bound[float64]{96.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			// Case: Supersets
			//
			//        |________|        |________|        |________|        |________|        |________|
			//                      |________________||__________________________________|
			//
			name: "Supersets",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{25.0, false}, Bound[float64]{45.0, false}},
				{Bound[float64]{45.0, true}, Bound[float64]{85.0, false}},
			},
			expectedPairs: []RangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}}, 'a'},
				{Range[float64]{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}}, 'e'},
			},
		},
		{
			name: "All",
			m: &rangeMap[float64, rune]{
				pairs:   append([]RangeValue[float64, rune]{}, pairs...),
				equal:   equal,
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			keys: []Range[float64]{
				{Bound[float64]{10.0, false}, Bound[float64]{20.0, false}},
				{Bound[float64]{30.0, true}, Bound[float64]{40.0, false}},
				{Bound[float64]{50.0, false}, Bound[float64]{60.0, true}},
				{Bound[float64]{70.0, true}, Bound[float64]{80.0, true}},
				{Bound[float64]{90.0, false}, Bound[float64]{100.0, false}},
			},
			expectedPairs: []RangeValue[float64, rune]{},
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
		m           *rangeMap[float64, rune]
		expectedAll []generic.KeyValue[Range[float64], rune]
	}{
		{
			name: "OK",
			m: &rangeMap[float64, rune]{
				pairs: []RangeValue[float64, rune]{
					{Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, '#'},
					{Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, '@'},
					{Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, 'a'},
					{Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, 'b'},
				},
				equal:   generic.NewEqualFunc[rune](),
				format:  defaultFormatMap[float64, rune],
				resolve: defaultResolve[rune],
			},
			expectedAll: []generic.KeyValue[Range[float64], rune]{
				{Key: Range[float64]{Bound[float64]{0.0, false}, Bound[float64]{0.9, false}}, Val: '#'},
				{Key: Range[float64]{Bound[float64]{1.0, true}, Bound[float64]{2.0, false}}, Val: '@'},
				{Key: Range[float64]{Bound[float64]{20.0, false}, Bound[float64]{40.0, true}}, Val: 'a'},
				{Key: Range[float64]{Bound[float64]{40.0, true}, Bound[float64]{80.0, true}}, Val: 'b'},
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
