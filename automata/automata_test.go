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

	t.Run("CmpStates", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     States
			expected int
		}{
			{
				name:     "BothNil",
				a:        nil,
				b:        nil,
				expected: 0,
			},
			{
				name:     "FirstNotNilSecondNil",
				a:        NewStates(1, 2),
				b:        nil,
				expected: 1,
			},
			{
				name:     "FirstNilSecondNotNil",
				a:        nil,
				b:        NewStates(3, 4),
				expected: -1,
			},
			{
				name:     "BothNotNil_Equal",
				a:        NewStates(2, 4),
				b:        NewStates(2, 4),
				expected: 0,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_ByFirstElement",
				a:        NewStates(1, 4),
				b:        NewStates(2, 4),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_BySecondElement",
				a:        NewStates(2, 3),
				b:        NewStates(2, 4),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_ByLength",
				a:        NewStates(2, 4),
				b:        NewStates(2, 4, 8),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_ByFirstElement",
				a:        NewStates(2, 4),
				b:        NewStates(1, 4),
				expected: 1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_BySecondElement",
				a:        NewStates(2, 4),
				b:        NewStates(2, 3),
				expected: 1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_ByLength",
				a:        NewStates(2, 4, 8),
				b:        NewStates(2, 4),
				expected: 1,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, CmpStates(tc.a, tc.b))
			})
		}
	})

	t.Run("HashStates", func(t *testing.T) {
		tests := []struct {
			name     string
			ss       States
			expected uint64
		}{
			{
				name:     "Ascending",
				ss:       NewStates(1, 2, 3, 4),
				expected: 0x36f8f2218060dd88,
			},
			{
				name:     "Descending",
				ss:       NewStates(4, 3, 2, 1),
				expected: 0x36f8f2218060dd88,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, HashStates(tc.ss))
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

	t.Run("CmpSymbols", func(t *testing.T) {
		tests := []struct {
			name     string
			a, b     Symbols
			expected int
		}{
			{
				name:     "BothNil",
				a:        nil,
				b:        nil,
				expected: 0,
			},
			{
				name:     "FirstNotNilSecondNil",
				a:        NewSymbols('a', 'b'),
				b:        nil,
				expected: 1,
			},
			{
				name:     "FirstNilSecondNotNil",
				a:        nil,
				b:        NewSymbols('c', 'd'),
				expected: -1,
			},
			{
				name:     "BothNotNil_Equal",
				a:        NewSymbols('b', 'd'),
				b:        NewSymbols('b', 'd'),
				expected: 0,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_ByFirstElement",
				a:        NewSymbols('a', 'd'),
				b:        NewSymbols('b', 'd'),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_BySecondElement",
				a:        NewSymbols('b', 'c'),
				b:        NewSymbols('b', 'd'),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstLessThanSecond_ByLength",
				a:        NewSymbols('b', 'd'),
				b:        NewSymbols('b', 'd', 'h'),
				expected: -1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_ByFirstElement",
				a:        NewSymbols('b', 'd'),
				b:        NewSymbols('a', 'd'),
				expected: 1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_BySecondElement",
				a:        NewSymbols('b', 'd'),
				b:        NewSymbols('b', 'c'),
				expected: 1,
			},
			{
				name:     "BothNotNil_FirstGreaterThanSecond_ByLength",
				a:        NewSymbols('b', 'd', 'h'),
				b:        NewSymbols('b', 'd'),
				expected: 1,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, CmpSymbols(tc.a, tc.b))
			})
		}
	})

	t.Run("HashSymbols", func(t *testing.T) {
		tests := []struct {
			name     string
			ss       Symbols
			expected uint64
		}{
			{
				name:     "Ascending",
				ss:       NewSymbols('a', 'b', 'c', 'd'),
				expected: 0x146340012207fa8,
			},
			{
				name:     "Descending",
				ss:       NewSymbols('d', 'c', 'b', 'a'),
				expected: 0x146340012207fa8,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.expected, HashSymbols(tc.ss))
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
					assert.True(t, union.Equal(tc.expectedUnion), "Expected:\n%s\nGot:\n%s", tc.expectedUnion, union)
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
