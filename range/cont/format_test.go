package cont

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rangesToSeq[T Continuous](ranges []Range[T]) iter.Seq[Range[T]] {
	return func(yield func(Range[T]) bool) {
		for _, v := range ranges {
			if !yield(v) {
				return
			}
		}
	}
}

func rangeValsToSeq2[T Continuous, V any](pairs []rangeValue[T, V]) iter.Seq2[Range[T], V] {
	return func(yield func(Range[T], V) bool) {
		for _, p := range pairs {
			if !yield(p.Range, p.Value) {
				return
			}
		}
	}
}

func TestDefaultFormatList(t *testing.T) {
	tests := []struct {
		name           string
		all            iter.Seq[Range[float64]]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangesToSeq[float64](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangesToSeq([]Range[float64]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangesToSeq([]Range[float64]{
				{Bound[float64]{2.2, false}, Bound[float64]{4.4, false}},
			}),
			expectedString: "[2.2, 4.4]",
		},
		{
			name: "Many",
			all: rangesToSeq([]Range[float64]{
				{Bound[float64]{2.2, false}, Bound[float64]{4.4, false}},
				{Bound[float64]{6.9, false}, Bound[float64]{7.0, true}},
				{Bound[float64]{7.0, true}, Bound[float64]{9.9, false}},
				{Bound[float64]{9.99, true}, Bound[float64]{9.999, true}},
			}),
			expectedString: "[2.2, 4.4] [6.9, 7) (7, 9.9] (9.99, 9.999)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, defaultFormatList(tc.all))
		})
	}
}

func TestDefaultFormatMap(t *testing.T) {
	tests := []struct {
		name           string
		all            iter.Seq2[Range[float64], rune]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangeValsToSeq2[float64, rune](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangeValsToSeq2([]rangeValue[float64, rune]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangeValsToSeq2([]rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.2, false}, Bound[float64]{4.4, false}}, 'a'},
			}),
			expectedString: "[2.2, 4.4]:97",
		},
		{
			name: "Many",
			all: rangeValsToSeq2([]rangeValue[float64, rune]{
				{Range[float64]{Bound[float64]{2.2, false}, Bound[float64]{4.4, false}}, 'a'},
				{Range[float64]{Bound[float64]{6.9, false}, Bound[float64]{7.0, true}}, 'b'},
				{Range[float64]{Bound[float64]{7.0, true}, Bound[float64]{9.9, false}}, 'c'},
				{Range[float64]{Bound[float64]{9.99, true}, Bound[float64]{9.999, true}}, 'd'},
			}),
			expectedString: "[2.2, 4.4]:97 [6.9, 7):98 (7, 9.9]:99 (9.99, 9.999):100",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, defaultFormatMap(tc.all))
		})
	}
}
