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
