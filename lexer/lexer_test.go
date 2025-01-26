package lexer

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/stretchr/testify/assert"
)

func TestToken_String(t *testing.T) {
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

func TestToken_Equal(t *testing.T) {
	tests := []struct {
		name          string
		t             Token
		rhs           Token
		expectedEqual bool
	}{
		{
			name: "Equal",
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
			rhs: Token{
				Terminal: grammar.Terminal("ID"),
				Lexeme:   "name",
				Pos: Position{
					Filename: "test_file",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			t: Token{
				Terminal: grammar.Terminal("ID"),
				Lexeme:   "name",
				Pos: Position{
					Filename: "foo",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			rhs: Token{
				Terminal: grammar.Terminal("ID"),
				Lexeme:   "name",
				Pos: Position{
					Filename: "bar",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.t.Equal(tc.rhs))
		})
	}
}

func TestPosition_Strong(t *testing.T) {
	tests := []struct {
		name           string
		p              Position
		expectedString string
	}{
		{
			name:           "Zero",
			p:              Position{},
			expectedString: `0`,
		},
		{
			name: "WithoutLineAndColumn",
			p: Position{
				Filename: "test_file",
				Offset:   69,
			},
			expectedString: `test_file:69`,
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
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
		})
	}
}

func TestPosition_Equal(t *testing.T) {
	tests := []struct {
		name          string
		p             Position
		rhs           Position
		expectedEqual bool
	}{
		{
			name: "Equal",
			p: Position{
				Filename: "test_file",
				Offset:   69,
				Line:     8,
				Column:   27,
			},
			rhs: Position{
				Filename: "test_file",
				Offset:   69,
				Line:     8,
				Column:   27,
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			p: Position{
				Filename: "foo",
				Offset:   69,
				Line:     8,
				Column:   27,
			},
			rhs: Position{
				Filename: "bar",
				Offset:   69,
				Line:     8,
				Column:   27,
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestPosition_IsZero(t *testing.T) {
	tests := []struct {
		name           string
		p              Position
		expectedIsZero bool
	}{
		{
			name:           "Zero",
			p:              Position{},
			expectedIsZero: true,
		},
		{
			name: "NotZero",
			p: Position{
				Filename: "test_file",
				Offset:   69,
			},
			expectedIsZero: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsZero, tc.p.IsZero())
		})
	}
}
