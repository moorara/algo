package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/set"
)

func TestCNFError(t *testing.T) {
	tests := []struct {
		name          string
		e             *CNFError
		expectedError string
	}{
		{
			name: "OK",
			e: &CNFError{
				P: Production{"rule", String[Symbol]{NonTerminal("lhs"), Terminal("="), NonTerminal("rhs")}},
			},
			expectedError: `production rule → lhs "=" rhs is neither a binary rule, a terminal rule, nor S → ε`,
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}

func TestLL1Error(t *testing.T) {
	tests := []struct {
		name          string
		e             *LL1Error
		expectedError string
	}{
		{
			name: "OK",
			e: &LL1Error{
				Description: "ε is in FIRST(β), but FOLLOW(A) and FIRST(α) are not disjoint sets",
				A:           NonTerminal("decls"),
				Alpha:       String[Symbol]{NonTerminal("decls"), NonTerminal("decl")},
				Beta:        ε,
				FOLLOWA: &TerminalsAndEndmarker{
					Terminals:         set.New(eqTerminal, "IDENT", "TOKEN"),
					IncludesEndmarker: true,
				},
				FIRSTα: &TerminalsAndEmpty{
					Terminals: set.New(eqTerminal, "IDENT", "TOKEN"),
				},
				FIRSTβ: &TerminalsAndEmpty{
					Terminals:     set.New(eqTerminal),
					IncludesEmpty: true,
				},
			},
			expectedError: "ε is in FIRST(β), but FOLLOW(A) and FIRST(α) are not disjoint sets:\n  decls → decls decl | ε\n    FOLLOW(decls): {\"IDENT\", \"TOKEN\", $}\n    FIRST(decls decl): {\"IDENT\", \"TOKEN\"}\n    FIRST(ε): {ε}\n",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}

func TestParsingTableError(t *testing.T) {
	tests := []struct {
		name          string
		e             *ParsingTableError
		expectedError string
	}{
		{
			name: "OK",
			e: &ParsingTableError{
				NonTerminal: NonTerminal("decls"),
				Terminal:    Terminal("IDENT"),
				Productions: set.New(eqProduction,
					Production{"decls", String[Symbol]{NonTerminal("decls"), NonTerminal("decl")}},
					Production{"decls", ε},
				),
			},
			expectedError: "multiple productions in parsing table at M[decls, \"IDENT\"]:\n  decls → decls decl\n  decls → ε\n",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}
