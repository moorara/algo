package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncs(t *testing.T) {
	t.Run("EqStates", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     States
			expected bool
		}{
			{
				name:     "NotEqual_BothNil",
				a:        nil,
				b:        nil,
				expected: true,
			},
			{
				name:     "NotEqual_FirstNotNilSecondNil",
				a:        NewStates(1, 2),
				b:        nil,
				expected: false,
			},
			{
				name:     "NotEqual_FirstNilSecondNotNil",
				a:        nil,
				b:        NewStates(3, 4),
				expected: false,
			},
			{
				name:     "NotEqual_BothNotNil",
				a:        NewStates(1, 2),
				b:        NewStates(3, 4),
				expected: false,
			},
			{
				name:     "Equal",
				a:        NewStates(1, 2),
				b:        NewStates(1, 2),
				expected: true,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, EqStates(tc.a, tc.b))
			})
		}
	})

	t.Run("EqSymbols", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     Symbols
			expected bool
		}{
			{
				name:     "NotEqual_BothNil",
				a:        nil,
				b:        nil,
				expected: true,
			},
			{
				name:     "NotEqual_FirstNotNilSecondNil",
				a:        NewSymbols('a', 'b'),
				b:        nil,
				expected: false,
			},
			{
				name:     "NotEqual_FirstNilSecondNotNil",
				a:        nil,
				b:        NewSymbols('c', 'd'),
				expected: false,
			},
			{
				name:     "NotEqual_BothNotNil",
				a:        NewSymbols('a', 'b'),
				b:        NewSymbols('c', 'd'),
				expected: false,
			},
			{
				name:     "Equal",
				a:        NewSymbols('a', 'b'),
				b:        NewSymbols('a', 'b'),
				expected: true,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, EqSymbols(tc.a, tc.b))
			})
		}
	})

	t.Run("unionStates", func(t *testing.T) {
		tests := []struct {
			name          string
			a, b          States
			expectedUnion States
		}{
			{
				name:          "BothNil",
				a:             nil,
				b:             nil,
				expectedUnion: nil,
			},
			{
				name:          "FirstNotNilSecondNil",
				a:             NewStates(1, 2),
				b:             nil,
				expectedUnion: NewStates(1, 2),
			},
			{
				name:          "FirstNilSecondNotNil",
				a:             nil,
				b:             NewStates(2, 3),
				expectedUnion: NewStates(2, 3),
			},
			{
				name:          "BothNotNil",
				a:             NewStates(1, 2),
				b:             NewStates(2, 3),
				expectedUnion: NewStates(1, 2, 3),
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				union := unionStates(tc.a, tc.b)

				if tc.expectedUnion == nil {
					assert.Nil(t, union)
				} else {
					assert.True(t, union.Equal(tc.expectedUnion), "expected: %s\ngot: %s\n", tc.expectedUnion, union)
				}
			})
		}
	})
}

func TestNewStates(t *testing.T) {
	tests := []struct {
		name string
		s    []State
	}{
		{
			name: "OK",
			s:    []State{0, 1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewStates(tc.s...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.s...))
		})
	}
}

func TestNewSymbols(t *testing.T) {
	tests := []struct {
		name string
		a    []Symbol
	}{
		{
			name: "OK",
			a:    []Symbol{'a', 'b', 'c', 'd'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewSymbols(tc.a...)

			assert.NotNil(t, set)
			assert.True(t, set.Contains(tc.a...))
		})
	}
}
