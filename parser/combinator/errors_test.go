package combinator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxError(t *testing.T) {
	tests := []struct {
		name          string
		e             *syntaxError
		expectedError string
	}{
		{
			name: "OK",
			e: &syntaxError{
				Pos:  0,
				Rune: 'a',
			},
			expectedError: "0: unexpected rune 'a'",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, tc.e.Error())
		})
	}
}

func TestSemanticError(t *testing.T) {
	tests := []struct {
		name           string
		e              *semanticError
		expectedError  string
		expectedUnwrap error
	}{
		{
			name: "OK",
			e: &semanticError{
				Pos: 0,
				Err: errors.New("invalid range"),
			},
			expectedError:  "0: invalid range",
			expectedUnwrap: errors.New("invalid range"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, tc.e.Error())
			assert.Equal(t, tc.expectedUnwrap, tc.e.Unwrap())
		})
	}
}
