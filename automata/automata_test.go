package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStates_Contains(t *testing.T) {
	tests := []struct {
		name           string
		s              States
		t              State
		expectedResult bool
	}{
		{
			name:           "No",
			s:              States{2, 4},
			t:              3,
			expectedResult: false,
		},
		{
			name:           "Yes",
			s:              States{2, 4},
			t:              4,
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.s.Contains(tc.t))
		})
	}
}

func TestStates_Equals(t *testing.T) {
	tests := []struct {
		name           string
		s              States
		t              States
		expectedResult bool
	}{
		{
			name:           "No",
			s:              States{2},
			t:              States{2, 4},
			expectedResult: false,
		},
		{
			name:           "Yes",
			s:              States{2, 4},
			t:              States{4, 2},
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.s.Equals(tc.t))
		})
	}
}

func TestSymbols_Contains(t *testing.T) {
	tests := []struct {
		name           string
		s              Symbols
		t              Symbol
		expectedResult bool
	}{
		{
			name:           "No",
			s:              Symbols{'b', 'd'},
			t:              'c',
			expectedResult: false,
		},
		{
			name:           "Yes",
			s:              Symbols{'b', 'd'},
			t:              'd',
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.s.Contains(tc.t))
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
