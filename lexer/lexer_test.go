package lexer

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	tests := []struct {
		name           string
		t              Token
		expectedString string
	}{
		{
			name: "OK",
			t: Token{
				Terminal: grammar.Terminal("ID"),
				Lexeme:   "name",
				Pos: Position{
					Filename: "test_file",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedString: `"ID" <name, test_file:8:27>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.t.String())
		})
	}
}

func TestPosition(t *testing.T) {
	tests := []struct {
		name           string
		p              Position
		expectedString string
		expectedIsZero bool
	}{
		{
			name:           "Zero",
			p:              Position{},
			expectedString: `0`,
			expectedIsZero: true,
		},
		{
			name: "WithoutLineAndColumn",
			p: Position{
				Filename: "test_file",
				Offset:   69,
			},
			expectedString: `test_file:69`,
			expectedIsZero: false,
		},
		{
			name: "WithLineAndColumn",
			p: Position{
				Filename: "test_file",
				Offset:   69,
				Line:     8,
				Column:   27,
			},
			expectedString: `test_file:8:27`,
			expectedIsZero: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
			assert.Equal(t, tc.expectedIsZero, tc.p.IsZero())
		})
	}
}
