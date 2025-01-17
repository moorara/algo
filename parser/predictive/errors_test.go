package predictive

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/set"
)

func TestParseError(t *testing.T) {
	tests := []struct {
		name          string
		e             *ParseError
		expectedError string
	}{
		{
			name: "WithoutPos",
			e: &ParseError{
				description: "cannot construct the parsing table",
			},
			expectedError: "cannot construct the parsing table",
		},
		{
			name: "WithDescription",
			e: &ParseError{
				description: "unacceptable input symbol",
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
				cause: &parsingTableError{
					NonTerminal: grammar.NonTerminal("decls"),
					Terminal:    grammar.Terminal("IDENT"),
					Productions: set.New(grammar.EqProduction,
						grammar.Production{
							Head: "decls",
							Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")},
						},
						grammar.Production{
							Head: "decls",
							Body: grammar.E,
						},
					),
				},
				Pos: lexer.Position{
					Filename: "test",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test:8:27: multiple productions in parsing table at M[decls, \"IDENT\"]:\n  decls → decls decl\n  decls → ε\n",
		},
		{
			name: "WithDescriptionAndCause",
			e: &ParseError{
				description: "parsing issue",
				cause: &parsingTableError{
					NonTerminal: grammar.NonTerminal("decls"),
					Terminal:    grammar.Terminal("IDENT"),
					Productions: set.New(grammar.EqProduction,
						grammar.Production{
							Head: "decls",
							Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")},
						},
						grammar.Production{
							Head: "decls",
							Body: grammar.E,
						},
					),
				},
				Pos: lexer.Position{
					Filename: "test",
					Offset:   69,
					Line:     8,
					Column:   27,
				},
			},
			expectedError: "test:8:27: parsing issue: multiple productions in parsing table at M[decls, \"IDENT\"]:\n  decls → decls decl\n  decls → ε\n",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
		assert.Equal(t, tc.e.cause, tc.e.Unwrap())
	}
}

func TestParsingTableError(t *testing.T) {
	tests := []struct {
		name          string
		e             *parsingTableError
		expectedError string
	}{
		{
			name: "OK",
			e: &parsingTableError{
				NonTerminal: grammar.NonTerminal("decls"),
				Terminal:    grammar.Terminal("IDENT"),
				Productions: set.New(grammar.EqProduction,
					grammar.Production{
						Head: "decls",
						Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")},
					},
					grammar.Production{
						Head: "decls",
						Body: grammar.E,
					},
				),
			},
			expectedError: "multiple productions in parsing table at M[decls, \"IDENT\"]:\n  decls → decls decl\n  decls → ε\n",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}
