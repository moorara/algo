package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbols_Contains(t *testing.T) {
	tests := []struct {
		name             string
		s                Symbols
		t                Symbol
		expectedContains bool
	}{
		{
			name:             "Yes",
			s:                Symbols{'a', 'b'},
			t:                'b',
			expectedContains: true,
		},
		{
			name:             "No",
			s:                Symbols{'a', 'b'},
			t:                'c',
			expectedContains: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedContains, tc.s.Contains(tc.t))
		})
	}
}

func TestSymbols_Equal(t *testing.T) {
	tests := []struct {
		name          string
		s             Symbols
		rhs           Symbols
		expectedEqual bool
	}{
		{
			name:          "Equal",
			s:             Symbols{'a', 'b'},
			rhs:           Symbols{'b', 'a'},
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			s:             Symbols{'a', 'b'},
			rhs:           Symbols{'a'},
			expectedEqual: false,
		},
		{
			name:          "NotEqual",
			s:             Symbols{'a'},
			rhs:           Symbols{'a', 'b'},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.s.Equal(tc.rhs))
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name           string
		s              string
		expectedString String
	}{
		{
			name:           "OK",
			s:              "ababb",
			expectedString: String{'a', 'b', 'a', 'b', 'b'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, ToString(tc.s))
		})
	}
}
