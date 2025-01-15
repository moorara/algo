package input

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/lexer"
)

func TestInputError(t *testing.T) {
	tests := []struct {
		name          string
		e             *InputError
		expectedError string
	}{
		{
			name: "WithoutLineAndColumn",
			e: &InputError{
				Description: "invalid utf-8 rune",
				Pos: lexer.Position{
					Filename: "test_file",
					Offset:   69,
				},
			},
			expectedError: "test_file:69: invalid utf-8 rune",
		},
		{
			name: "WithLineAndColumn",
			e: &InputError{
				Description: "invalid utf-8 rune",
				Pos: lexer.Position{
					Filename: "test_file",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test_file:8:27: invalid utf-8 rune",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}
