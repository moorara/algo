package parser

import (
	"errors"
	"testing"

	"github.com/moorara/algo/lexer"
	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	tests := []struct {
		name          string
		e             *ParseError
		expectedError string
	}{
		{
			name: "WithDescription",
			e: &ParseError{
				Description: "error on parsing the input string",
			},
			expectedError: "error on parsing the input string",
		},
		{
			name: "WithDescriptionAndPos",
			e: &ParseError{
				Description: "unacceptable input symbol",
				Pos: lexer.Position{
					Filename: "test",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test:8:27: unacceptable input symbol",
		},
		{
			name: "WithCause",
			e: &ParseError{
				Cause: errors.New("grammar is ambiguous"),
				Pos: lexer.Position{
					Filename: "test",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test:8:27: grammar is ambiguous",
		},
		{
			name: "WithDescriptionAndCauseAndPos",
			e: &ParseError{
				Description: "error on parsing the input string",
				Cause:       errors.New("grammar is ambiguous"),
				Pos: lexer.Position{
					Filename: "test",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test:8:27: error on parsing the input string: grammar is ambiguous",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
		assert.Equal(t, tc.e.Cause, tc.e.Unwrap())
	}
}
