package slr

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/lr"
)

var starts = []grammar.NonTerminal{
	"E′",
	"grammar′",
}

var prods = [][]grammar.Production{
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
}

var grammars = []grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[0][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[1][1:],
		"grammar",
	),
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]grammar.Terminal{"(", ")", "*", "+", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt0.AddACTION(0, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(0, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(1, "+", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(1, grammar.Endmarker, lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(2, "*", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(2, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(2, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(3, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(4, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(5, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(5, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(6, ")", lr.Action{Type: lr.SHIFT, State: 3})
	pt0.AddACTION(6, "+", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(7, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(7, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(8, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(8, "*", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(8, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(8, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(9, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(9, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(10, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(11, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})

	pt0.SetGOTO(0, "E", 1)
	pt0.SetGOTO(0, "T", 8)
	pt0.SetGOTO(0, "F", 11)
	pt0.SetGOTO(5, "T", 2)
	pt0.SetGOTO(5, "F", 11)
	pt0.SetGOTO(7, "F", 4)
	pt0.SetGOTO(9, "E", 6)
	pt0.SetGOTO(9, "T", 8)
	pt0.SetGOTO(9, "F", 11)

	pt1 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36},
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
	)

	return []*lr.ParsingTable{pt0, pt1}
}

func TestSLRParser_Parse(t *testing.T) {
	tests := []struct {
		name                 string
		p                    *slrParser
		prodF                parser.ProductionFunc
		tokenF               parser.TokenFunc
		expectedErrorStrings []string
	}{
		{
			name: "None_SLR(1)_Grammar",
			p: &slrParser{
				G:     grammars[1],
				lexer: new(MockLexer),
			},
			prodF:  func(grammar.Production) {},
			tokenF: func(lexer.Token) {},
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
		{
			name: "EmptyString",
			p: &slrParser{
				G: grammars[0],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: io.EOF},
					},
				},
			},
			prodF:  func(grammar.Production) {},
			tokenF: func(lexer.Token) {},
			expectedErrorStrings: []string{
				`no action for ACTION[0, $]`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &slrParser{
				G: grammars[0],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
			},
			prodF:  func(grammar.Production) {},
			tokenF: func(lexer.Token) {},
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &slrParser{
				G: grammars[0],
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
			prodF:  func(grammar.Production) {},
			tokenF: func(lexer.Token) {},
			expectedErrorStrings: []string{
				`input stream failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &slrParser{
				G: grammars[0],
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
			prodF:  func(grammar.Production) {},
			tokenF: func(lexer.Token) {},
			expectedErrorStrings: []string{
				`no action for ACTION[0, "+"]`,
			},
		},
		{
			name: "Success",
			p: &slrParser{
				G: grammars[0],
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
			prodF:                func(grammar.Production) {},
			tokenF:               func(lexer.Token) {},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.p.G.Verify())
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

func TestSLRParser_ParseAST(t *testing.T) {
	tests := []struct {
		name                 string
		p                    *slrParser
		expectedAST          parser.Node
		expectedErrorStrings []string
	}{
		{
			name: "None_SLR(1)_Grammar",
			p: &slrParser{
				G:     grammars[1],
				lexer: new(MockLexer),
			},
			expectedAST: nil,
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
		{
			name: "EmptyString",
			p: &slrParser{
				G: grammars[0],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: io.EOF},
					},
				},
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`no action for ACTION[0, $]`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &slrParser{
				G: grammars[0],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &slrParser{
				G: grammars[0],
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
						{OutError: errors.New("input failed")},
					},
				},
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`input failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &slrParser{
				G: grammars[0],
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
			expectedAST: nil,
			expectedErrorStrings: []string{
				`no action for ACTION[0, "+"]`,
			},
		},
		{
			name: "Success",
			p: &slrParser{
				G: grammars[0],
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
			assert.NoError(t, tc.p.G.Verify())
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
