package disc

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rangesToSeq[T Discrete](ranges []Range[T]) iter.Seq[Range[T]] {
	return func(yield func(Range[T]) bool) {
		for _, v := range ranges {
			if !yield(v) {
				return
			}
		}
	}
}

func rangeValsToSeq2[T Discrete, V any](pairs []rangeValue[T, V]) iter.Seq2[Range[T], V] {
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
		all            iter.Seq[Range[int]]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangesToSeq[int](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangesToSeq([]Range[int]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangesToSeq([]Range[int]{
				{2, 4},
			}),
			expectedString: "[2, 4]",
		},
		{
			name: "Many",
			all: rangesToSeq([]Range[int]{
				{2, 4},
				{6, 8},
				{10, 10},
				{16, 20},
			}),
			expectedString: "[2, 4] [6, 8] [10, 10] [16, 20]",
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
		all            iter.Seq2[Range[int], rune]
		expectedString string
	}{
		{
			name:           "Nil",
			all:            rangeValsToSeq2[int, rune](nil),
			expectedString: "",
		},
		{
			name:           "Zero",
			all:            rangeValsToSeq2([]rangeValue[int, rune]{}),
			expectedString: "",
		},
		{
			name: "One",
			all: rangeValsToSeq2([]rangeValue[int, rune]{
				{Range[int]{2, 4}, 'a'},
			}),
			expectedString: "[2, 4]:97",
		},
		{
			name: "Many",
			all: rangeValsToSeq2([]rangeValue[int, rune]{
				{Range[int]{2, 4}, 'a'},
				{Range[int]{6, 8}, 'b'},
				{Range[int]{10, 10}, 'c'},
				{Range[int]{16, 20}, 'd'},
			}),
			expectedString: "[2, 4]:97 [6, 8]:98 [10, 10]:99 [16, 20]:100",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, defaultFormatMap(tc.all))
		})
	}
}
