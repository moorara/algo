package lr

import (
	"errors"
	"io"
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/stretchr/testify/assert"
)

var starts = []grammar.NonTerminal{
	"S′",
	"E′",
	"E′",
	"grammar′",
}

var prods = [][]*grammar.Production{
	{
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                          // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")}}, // S → CC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")}},    // C → cC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("d")}},                              // C → d
	},
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
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
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
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
		prods[0][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[1][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		prods[2][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[3][1:],
		"grammar",
	),
}

func TestParser_Parse(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		p                    *Parser
		prodF                parser.ProductionFunc
		tokenF               parser.TokenFunc
		expectedErrorStrings []string
	}{
		{
			name: "EmptyString",
			p: &Parser{
				L: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: io.EOF},
					},
				},
				T: pt[0],
			},
			prodF:  func(*grammar.Production) {},
			tokenF: func(*lexer.Token) {},
			expectedErrorStrings: []string{
				`no action for ACTION[0, $]`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &Parser{
				L: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
				T: pt[0],
			},
			prodF:  func(*grammar.Production) {},
			tokenF: func(*lexer.Token) {},
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &Parser{
				L: &MockLexer{
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
				T: pt[0],
			},
			prodF:  func(*grammar.Production) {},
			tokenF: func(*lexer.Token) {},
			expectedErrorStrings: []string{
				`input stream failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &Parser{
				L: &MockLexer{
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
				T: pt[0],
			},
			prodF:  func(*grammar.Production) {},
			tokenF: func(*lexer.Token) {},
			expectedErrorStrings: []string{
				`no action for ACTION[0, "+"]`,
			},
		},
		{
			name: "Success",
			p: &Parser{
				L: &MockLexer{
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
				T: pt[0],
			},
			prodF:                func(*grammar.Production) {},
			tokenF:               func(*lexer.Token) {},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.p.Parse(tc.prodF, tc.tokenF)

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

func TestParser_ParseAST(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		p                    *Parser
		expectedAST          parser.Node
		expectedErrorStrings []string
	}{
		{
			name: "EmptyString",
			p: &Parser{
				L: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: io.EOF},
					},
				},
				T: pt[0],
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`no action for ACTION[0, $]`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &Parser{
				L: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
				T: pt[0],
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &Parser{
				L: &MockLexer{
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
						{OutError: errors.New("input failed")},
					},
				},
				T: pt[0],
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`input failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &Parser{
				L: &MockLexer{
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
				T: pt[0],
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`no action for ACTION[0, "+"]`,
			},
		},
		{
			name: "Success",
			p: &Parser{
				L: &MockLexer{
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
				T: pt[0],
			},
			expectedAST: &parser.InternalNode{
				NonTerminal: "E",
				Production: &grammar.Production{
					Head: "E",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")},
				},
				Children: []parser.Node{
					&parser.InternalNode{
						NonTerminal: "E",
						Production: &grammar.Production{
							Head: "E",
							Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")},
						},
						Children: []parser.Node{
							&parser.InternalNode{
								NonTerminal: "T",
								Production: &grammar.Production{
									Head: "T",
									Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")},
								},
								Children: []parser.Node{
									&parser.InternalNode{
										NonTerminal: "F",
										Production: &grammar.Production{
											Head: "F",
											Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
										},
										Children: []parser.Node{
											&parser.LeafNode{
												Terminal: "id",
												Lexeme:   "a",
												Position: lexer.Position{
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
						},
					},
					&parser.LeafNode{
						Terminal: "+",
						Lexeme:   "+",
						Position: lexer.Position{
							Filename: "test",
							Offset:   2,
							Line:     1,
							Column:   3,
						},
					},
					&parser.InternalNode{
						NonTerminal: "T",
						Production: &grammar.Production{
							Head: "T",
							Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")},
						},
						Children: []parser.Node{
							&parser.InternalNode{
								NonTerminal: "F",
								Production: &grammar.Production{
									Head: "F",
									Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
								},
								Children: []parser.Node{
									&parser.LeafNode{
										Terminal: "id",
										Lexeme:   "b",
										Position: lexer.Position{
											Filename: "test",
											Offset:   4,
											Line:     1,
											Column:   5,
										},
									},
								},
							},
						},
					},
				},
			},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ast, err := tc.p.ParseAST()

			if len(tc.expectedErrorStrings) == 0 {
				assert.True(t, ast.Equals(tc.expectedAST))
				assert.NoError(t, err)
			} else {
				assert.Nil(t, ast)
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}
