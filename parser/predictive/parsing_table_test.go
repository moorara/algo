package predictive

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/stretchr/testify/assert"
)

var CFGrammars = []grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]grammar.NonTerminal{"S", "E"},
		[]grammar.Production{
			{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                  // S → E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("-"), grammar.NonTerminal("E")}}, // E → E - E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("/"), grammar.NonTerminal("E")}}, // E → E / E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("-"), grammar.NonTerminal("E")}},                           // E → - E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
		},
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]grammar.NonTerminal{"S", "E", "T", "F"},
		[]grammar.Production{
			{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                  // S → E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("-"), grammar.NonTerminal("T")}}, // E → E - T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("/"), grammar.NonTerminal("F")}}, // T → T / F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
		},
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},                         // E → T E′
			{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}}, // E′ → + T E′
			{Head: "E′", Body: grammar.E}, // E′ → ε
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}},                         // T → F T′
			{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}}, // T′ → * F T′
			{Head: "T′", Body: grammar.E}, // T′ → ε
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}, // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                 // F → id
		},
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		[]grammar.Production{
			{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
			{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("GRAMMAR"), grammar.Terminal("IDENT")}},       // name → GRAMMAR IDENT
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
		"grammar",
	),
}

func getTestParsingTables() []*parsingTable {
	pt0 := newParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
	)

	pt0.Add("E", "id", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.Add("E", "(", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.Add("E′", "+", grammar.Production{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.Add("E′", ")", grammar.Production{Head: "E′", Body: grammar.E})
	pt0.Add("E′", grammar.Endmarker, grammar.Production{Head: "E′", Body: grammar.E})
	pt0.Add("T", "id", grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.Add("T", "(", grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.Add("T′", "+", grammar.Production{Head: "T′", Body: grammar.E})
	pt0.Add("T′", "*", grammar.Production{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.Add("T′", ")", grammar.Production{Head: "T′", Body: grammar.E})
	pt0.Add("T′", grammar.Endmarker, grammar.Production{Head: "T′", Body: grammar.E})
	pt0.Add("F", "id", grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}})
	pt0.Add("F", "(", grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}})

	pt1 := newParsingTable(
		[]grammar.Terminal{"a", "b", "e", "i", "t"},
		[]grammar.NonTerminal{"S", "S′", "E"},
	)

	pt1.Add("S", "a", grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("a")}})
	pt1.Add("S", "i", grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("i"), grammar.NonTerminal("E"), grammar.Terminal("t"), grammar.NonTerminal("S"), grammar.NonTerminal("S′")}})
	pt1.Add("S′", "e", grammar.Production{Head: "S′", Body: grammar.E})
	pt1.Add("S′", "e", grammar.Production{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("e"), grammar.NonTerminal("S")}})
	pt1.Add("S′", grammar.Endmarker, grammar.Production{Head: "S′", Body: grammar.E})
	pt1.Add("E", "b", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("b")}})

	pt2 := newParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	return []*parsingTable{pt0, pt1, pt2}
}

func TestNewParsingTable(t *testing.T) {
	tests := []struct {
		name                 string
		g                    grammar.CFG
		expectedErrorStrings []string
	}{
		{
			name: "1st",
			g:    CFGrammars[0],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[E, "("]`,
				`multiple productions in parsing table at M[E, "-"]`,
				`multiple productions in parsing table at M[E, "id"]`,
			},
		},
		{
			name: "2nd",
			g:    CFGrammars[1],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[E, "("]`,
				`multiple productions in parsing table at M[E, "id"]`,
				`multiple productions in parsing table at M[T, "("]`,
				`multiple productions in parsing table at M[T, "id"]`,
			},
		},
		{
			name:                 "3rd",
			g:                    CFGrammars[2],
			expectedErrorStrings: nil,
		},
		{
			name: "4th",
			g:    CFGrammars[3],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[decls, "IDENT"]`,
				`multiple productions in parsing table at M[decls, "TOKEN"]`,
				`multiple productions in parsing table at M[rule, "IDENT"]`,
				`multiple productions in parsing table at M[token, "TOKEN"]`,
				`multiple productions in parsing table at M[rhs, "("]`,
				`multiple productions in parsing table at M[rhs, "IDENT"]`,
				`multiple productions in parsing table at M[rhs, "STRING"]`,
				`multiple productions in parsing table at M[rhs, "TOKEN"]`,
				`multiple productions in parsing table at M[rhs, "["]`,
				`multiple productions in parsing table at M[rhs, "{"]`,
				`multiple productions in parsing table at M[rhs, "{{"]`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			table := NewParsingTable(tc.g)
			err := table.CheckErrors()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}

func TestParsingTable_String(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name               string
		pt                 *parsingTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			pt:   pt[0],
			expectedSubstrings: []string{
				`┌──────────────┬────────────────────────────────────────────────────────────────────────────┐`,
				`│              │                                  Terminal                                  │`,
				`├──────────────┼───────────────┬───────────────┬───────────────┬────────┬──────────┬────────┤`,
				`│ Non-Terminal │      "+"      │      "*"      │      "("      │  ")"   │   "id"   │   $    │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      E       │               │               │   E → T E′    │        │ E → T E′ │        │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      E′      │ E′ → "+" T E′ │               │               │ E′ → ε │          │ E′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T       │               │               │   T → F T′    │        │ T → F T′ │        │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T′      │    T′ → ε     │ T′ → "*" F T′ │               │ T′ → ε │          │ T′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      F       │               │               │ F → "(" E ")" │        │ F → "id" │        │`,
				`└──────────────┴───────────────┴───────────────┴───────────────┴────────┴──────────┴────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.pt.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}

func TestParsingTable_Equals(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name           string
		pt             *parsingTable
		rhs            ParsingTable
		expectedEquals bool
	}{
		{
			name:           "Equal",
			pt:             pt[0],
			rhs:            pt[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			pt:             pt[1],
			rhs:            pt[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.pt.Equals(tc.rhs))
		})
	}
}

func TestParsingTable_Add(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *parsingTable
		A          grammar.NonTerminal
		a          grammar.Terminal
		prod       grammar.Production
		expectedOK bool
	}{
		{
			name: "OK",
			pt:   pt[2],
			A:    grammar.NonTerminal("F"),
			a:    grammar.Terminal("("),
			prod: grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.pt.Add(tc.A, tc.a, tc.prod)
		})
	}
}

func TestParsingTable_Get(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                string
		pt                  *parsingTable
		A                   grammar.NonTerminal
		a                   grammar.Terminal
		expectedProductions set.Set[grammar.Production]
	}{
		{
			name: "OK",
			pt:   pt[0],
			A:    grammar.NonTerminal("E′"),
			a:    grammar.Terminal("+"),
			expectedProductions: set.New(grammar.EqProduction,
				grammar.Production{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prods := tc.pt.Get(tc.A, tc.a)
			assert.True(t, prods.Equals(tc.expectedProductions))
		})
	}
}

func TestParsingTable_CheckErrors(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		pt                   *parsingTable
		expectedErrorStrings []string
	}{
		{
			name:                 "NoError",
			pt:                   pt[0],
			expectedErrorStrings: nil,
		},
		{
			name: "Error",
			pt:   pt[1],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[S′, "e"]`,
				`S′ → "e" S`,
				`S′ → ε`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.pt.CheckErrors()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}
