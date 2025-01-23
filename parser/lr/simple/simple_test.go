package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
)

var starts = []grammar.NonTerminal{
	"E′",
	"grammar′",
}

var prods = [][]*grammar.Production{
	{
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
	{
		{Head: "grammar′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("grammar")}},                           // grammar′ → grammar
		{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
		{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT")}},       // name → "grammar" IDENT
		{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}},   // decls → decls decl
		{Head: "decls", Body: grammar.E}, // decls → ε
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                  // decl → token
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                   // decl → rule
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING")}}, // token → TOKEN "=" STRING
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX")}},  // token → TOKEN "=" REGEX
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs")}}, // rule → lhs "=" rhs
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("=")}},                             // rule → lhs "="
		{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // lhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                         // rhs → rhs rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},  // rhs → rhs "|" rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},       // rhs → "(" rhs ")"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},       // rhs → "[" rhs "]"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},       // rhs → "{" rhs "}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},     // rhs → "{{" rhs "}}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // rhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                    // rhs → term
		{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                  // nonterm → IDENT
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                     // term → TOKEN
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                    // term → STRING
	},
}

var grammars = []*grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[0][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[1][1:],
		"grammar",
	),
}

func TestNew(t *testing.T) {
	tests := []struct {
		name                 string
		L                    lexer.Lexer
		G                    *grammar.CFG
		expectedErrorStrings []string
	}{
		{
			name:                 "Success",
			L:                    nil,
			G:                    grammars[0],
			expectedErrorStrings: nil,
		},
		{
			name: "None_SLR(1)_Grammar",
			L:    nil,
			G:    grammars[1],
			expectedErrorStrings: []string{
				`failed to construct the SLR parsing table: 20 errors occurred:`,
				`shift/reduce conflict at ACTION[2, "("]`,
				`shift/reduce conflict at ACTION[2, "IDENT"]`,
				`shift/reduce conflict at ACTION[2, "STRING"]`,
				`shift/reduce conflict at ACTION[2, "TOKEN"]`,
				`shift/reduce conflict at ACTION[2, "["]`,
				`shift/reduce conflict at ACTION[2, "{"]`,
				`shift/reduce conflict at ACTION[2, "{{"]`,
				`shift/reduce conflict at ACTION[2, "|"]`,
				`shift/reduce conflict at ACTION[7, "IDENT"]`,
				`shift/reduce conflict at ACTION[7, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "("]`,
				`shift/reduce conflict at ACTION[14, "IDENT"]`,
				`shift/reduce conflict at ACTION[14, "STRING"]`,
				`shift/reduce conflict at ACTION[14, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "["]`,
				`shift/reduce conflict at ACTION[14, "{"]`,
				`shift/reduce conflict at ACTION[14, "{{"]`,
				`shift/reduce conflict at ACTION[14, "|"]`,
				`shift/reduce conflict at ACTION[19, "IDENT"]`,
				`shift/reduce conflict at ACTION[19, "TOKEN"]`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			p, err := New(tc.L, tc.G)

			if len(tc.expectedErrorStrings) == 0 {
				assert.NotNil(t, p)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, p)
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}
