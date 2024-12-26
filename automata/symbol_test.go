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

func TestSymbols_Equals(t *testing.T) {
	tests := []struct {
		name           string
		s              Symbols
		rhs            Symbols
		expectedEquals bool
	}{
		{
			name:           "Equal",
			s:              Symbols{'a', 'b'},
			rhs:            Symbols{'b', 'a'},
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			s:              Symbols{'a', 'b'},
			rhs:            Symbols{'a'},
			expectedEquals: false,
		},
		{
			name:           "NotEqual",
			s:              Symbols{'a'},
			rhs:            Symbols{'a', 'b'},
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.s.Equals(tc.rhs))
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
