package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestSymbolRange(t *testing.T) {
	tests := []struct {
		name           string
		r              SymbolRange
		expectedString string
		equal          SymbolRange
		expectedEqual  bool
	}{
		{
			name:           "Equal_Empty",
			r:              SymbolRange{Start: E, End: E},
			expectedString: "[ε]",
			equal:          SymbolRange{Start: E, End: E},
			expectedEqual:  true,
		},
		{
			name:           "Equal_Range",
			r:              SymbolRange{Start: '0', End: '9'},
			expectedString: "[0..9]",
			equal:          SymbolRange{Start: '0', End: '9'},
			expectedEqual:  true,
		},
		{
			name:           "NotEqual_Range",
			r:              SymbolRange{Start: 'a', End: 'z'},
			expectedString: "[a..z]",
			equal:          SymbolRange{Start: 'α', End: 'ω'},
			expectedEqual:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("Validate", func(t *testing.T) {
				tc.r.Validate()
				assert.Equal(t, tc.expectedString, tc.r.String())
				assert.Equal(t, tc.expectedEqual, tc.r.Equal(tc.equal))
			})
		})
	}
}
