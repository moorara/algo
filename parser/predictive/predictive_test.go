package predictive

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
)

var grammars = []grammar.CFG{
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

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		G     grammar.CFG
		lexer lexer.Lexer
	}{
		{
			name:  "OK",
			G:     grammars[2],
			lexer: new(MockLexer),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			p := New(tc.G, tc.lexer)
			assert.NotNil(t, p)
		})
	}
}

func TestPredictiveParser_Parse(t *testing.T) {
	tests := []struct {
		name                 string
		p                    *predictiveParser
		yield                parser.Action
		expectedErrorStrings []string
	}{
		{
			name: "None_LL(1)_Grammar",
			p: &predictiveParser{
				G:     grammars[0],
				lexer: new(MockLexer),
			},
			yield: func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[E, "-"]`,
				`multiple productions in parsing table at M[E, "("]`,
				`multiple productions in parsing table at M[E, "id"]`,
			},
		},
		{
			name: "Input_Empty",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: io.EOF},
					},
				},
			},
			yield:                func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: nil,
		},
		{
			name: "First_NextToken_Fails",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
			},
			yield: func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						// First token
						{
							OutToken: lexer.Token{
								Terminal: grammar.Terminal("id"),
								Lexeme:   "a",
								Pos: lexer.Position{
									Filename: "test",
									Offset:   0,
									Line:     1,
									Column:   1,
								},
							},
						},
						// EOF
						{OutError: errors.New("input stream failed")},
					},
				},
			},
			yield: func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: []string{
				`input stream failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						// First token
						{
							OutToken: lexer.Token{
								Terminal: grammar.Terminal("+"),
								Lexeme:   "+",
								Pos: lexer.Position{
									Filename: "test",
									Offset:   0,
									Line:     1,
									Column:   1,
								},
							},
						},
					},
				},
			},
			yield: func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: []string{
				`unacceptable input <"+", +> for non-terminal E`,
			},
		},
		{
			name: "Success",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						// First token
						{
							OutToken: lexer.Token{
								Terminal: grammar.Terminal("id"),
								Lexeme:   "a",
								Pos: lexer.Position{
									Filename: "test",
									Offset:   0,
									Line:     1,
									Column:   1,
								},
							},
						},
						// Second token
						{
							OutToken: lexer.Token{
								Terminal: grammar.Terminal("+"),
								Lexeme:   "+",
								Pos: lexer.Position{
									Filename: "test",
									Offset:   2,
									Line:     1,
									Column:   3,
								},
							},
						},
						// Third token
						{
							OutToken: lexer.Token{
								Terminal: grammar.Terminal("id"),
								Lexeme:   "b",
								Pos: lexer.Position{
									Filename: "test",
									Offset:   4,
									Line:     1,
									Column:   5,
								},
							},
						},
						// EOF
						{OutError: io.EOF},
					},
				},
			},
			yield:                func(grammar.Production, lexer.Token) {},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.p.G.Verify())
			err := tc.p.Parse(tc.yield)

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
