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

type (
	// MockLexer is an implementation of lexer.Lexer for testing purposes.
	MockLexer struct {
		NextTokenIndex int
		NextTokenMocks []NextTokenMock
	}

	NextTokenMock struct {
		OutToken lexer.Token
		OutError error
	}
)

func (m *MockLexer) NextToken() (lexer.Token, error) {
	i := m.NextTokenIndex
	m.NextTokenIndex++
	return m.NextTokenMocks[i].OutToken, m.NextTokenMocks[i].OutError
}

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		G     *grammar.CFG
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
		tokenF               parser.TokenFunc
		prodF                parser.ProductionFunc
		expectedErrorStrings []string
	}{
		{
			name: "None_LL(1)_Grammar",
			p: &predictiveParser{
				G:     grammars[0],
				lexer: new(MockLexer),
			},
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`multiple productions at M[E, "-"]`,
				`multiple productions at M[E, "("]`,
				`multiple productions at M[E, "id"]`,
			},
		},
		{
			name: "EmptyString",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
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
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`unacceptable input <"+", +> for non-terminal E`,
			},
		},
		{
			name: "TokenFuncError",
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
			tokenF: func(*lexer.Token) error { return errors.New("invalid semantic") },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`invalid semantic`,
			},
		},
		{
			name: "ProductionFuncError",
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return errors.New("invalid semantic") },
			expectedErrorStrings: []string{
				`invalid semantic`,
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
				G:     grammars[0],
				lexer: new(MockLexer),
			},
			expectedAST: nil,
			expectedErrorStrings: []string{
				`multiple productions at M[E, "-"]`,
				`multiple productions at M[E, "("]`,
				`multiple productions at M[E, "id"]`,
			},
		},
		{
			name: "EmptyString",
			p: &predictiveParser{
				G: grammars[2],
				lexer: &MockLexer{
					NextTokenMocks: []NextTokenMock{
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
				G: grammars[2],
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
			expectedAST: nil,
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
