package predictive

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		G     *grammar.CFG
		lexer lexer.Lexer
	}{
		{
			name:  "OK",
			G:     parsertest.Grammars[0],
			lexer: new(parsertest.MockLexer),
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
		tokenF               parser.TokenFunc
		prodF                parser.ProductionFunc
		expectedErrorStrings []string
	}{
		{
			name: "None_LL(1)_Grammar",
			p: &predictiveParser{
				G:     parsertest.Grammars[4],
				lexer: new(parsertest.MockLexer),
			},
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`failed to construct the predictive parsing table: 2 errors occurred:`,
				`multiple productions at M[E, "("]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "(" E ")"`,
				`multiple productions at M[E, "id"]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "id"`,
			},
		},
		{
			name: "EmptyString",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
						{OutError: io.EOF},
					},
				},
			},
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`unacceptable input <$, > for non-terminal E`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
						{OutError: errors.New("cannot read rune")},
					},
				},
			},
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`cannot read rune`,
			},
		},
		{
			name: "Second_NextToken_Fails",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`input failed`,
			},
		},
		{
			name: "Invalid_Input",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`unacceptable input <"+", +> for non-terminal E`,
			},
		},
		{
			name: "TokenFuncError",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			tokenF: func(*lexer.Token) error { return errors.New("invalid semantic") },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`invalid semantic`,
			},
		},
		{
			name: "ProductionFuncError",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return errors.New("invalid semantic") },
			expectedErrorStrings: []string{
				`invalid semantic`,
			},
		},
		{
			name: "Success",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			tokenF:               func(*lexer.Token) error { return nil },
			prodF:                func(*grammar.Production) error { return nil },
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.p.G.Verify())
			err := tc.p.Parse(tc.tokenF, tc.prodF)

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

func TestPredictiveParser_ParseAndBuildAST(t *testing.T) {
	tests := []struct {
		name                 string
		p                    *predictiveParser
		expectedAST          parser.Node
		expectedErrorStrings []string
	}{
		{
			name: "None_LL(1)_Grammar",
			p: &predictiveParser{
				G:     parsertest.Grammars[4],
				lexer: new(parsertest.MockLexer),
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`failed to construct the predictive parsing table: 2 errors occurred:`,
				`multiple productions at M[E, "("]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "(" E ")"`,
				`multiple productions at M[E, "id"]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "id"`,
			},
		},
		{
			name: "EmptyString",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
						{OutError: io.EOF},
					},
				},
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`unacceptable input <$, > for non-terminal E`,
			},
		},
		{
			name: "First_NextToken_Fails",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
				`unacceptable input <"+", +> for non-terminal E`,
			},
		},
		{
			name: "Success",
			p: &predictiveParser{
				G: parsertest.Grammars[0],
				lexer: &parsertest.MockLexer{
					NextTokenMocks: []parsertest.NextTokenMock{
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
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")},
				},
				Children: []parser.Node{
					&parser.InternalNode{
						NonTerminal: "T",
						Production: &grammar.Production{
							Head: "T",
							Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")},
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
							&parser.InternalNode{
								NonTerminal: "T′",
								Production: &grammar.Production{
									Head: "T′",
									Body: grammar.E,
								},
							},
						},
					},
					&parser.InternalNode{
						NonTerminal: "E′",
						Production: &grammar.Production{
							Head: "E′",
							Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")},
						},
						Children: []parser.Node{
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
									Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")},
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
									&parser.InternalNode{
										NonTerminal: "T′",
										Production: &grammar.Production{
											Head: "T′",
											Body: grammar.E,
										},
									},
								},
							},
							&parser.InternalNode{
								NonTerminal: "E′",
								Production: &grammar.Production{
									Head: "E′",
									Body: grammar.E,
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
			ast, err := tc.p.ParseAndBuildAST()

			if len(tc.expectedErrorStrings) == 0 {
				assert.True(t, ast.Equal(tc.expectedAST))
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
