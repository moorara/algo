package predictive

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

func TestParsingTableError(t *testing.T) {
	tests := []struct {
		name          string
		e             *ParsingTableError
		expectedError string
	}{
		{
			name: "OK",
			e: &ParsingTableError{
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
