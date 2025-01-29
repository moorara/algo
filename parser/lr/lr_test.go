package lr

import (
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/stretchr/testify/assert"
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

func TestParser_Parse(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		p                    *Parser
		tokenF               parser.TokenFunc
		prodF                parser.ProductionFunc
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`no action exists in the parsing table for ACTION[0, $]`,
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`no action exists in the parsing table for ACTION[0, "+"]`,
			},
		},
		{
			name: "TokenFuncError",
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
			tokenF: func(*lexer.Token) error { return errors.New("invalid semantic") },
			prodF:  func(*grammar.Production) error { return nil },
			expectedErrorStrings: []string{
				`invalid semantic`,
			},
		},
		{
			name: "ProductionFuncError",
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
			tokenF: func(*lexer.Token) error { return nil },
			prodF:  func(*grammar.Production) error { return errors.New("invalid semantic") },
			expectedErrorStrings: []string{
				`invalid semantic`,
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
			tokenF:               func(*lexer.Token) error { return nil },
			prodF:                func(*grammar.Production) error { return nil },
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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

func TestParser_ParseAndBuildAST(t *testing.T) {
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
				`no action exists in the parsing table for ACTION[0, $]`,
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
				`no action exists in the parsing table for ACTION[0, "+"]`,
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

func TestParser_ParseAndEvaluate(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		p                    *Parser
		eval                 EvaluateFunc
		expectedValue        *Value
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
			eval:          func(*grammar.Production, []*Value) (any, error) { return nil, nil },
			expectedValue: nil,
			expectedErrorStrings: []string{
				`no action exists in the parsing table for ACTION[0, $]`,
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
			eval:          func(*grammar.Production, []*Value) (any, error) { return nil, nil },
			expectedValue: nil,
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
			eval:          func(*grammar.Production, []*Value) (any, error) { return nil, nil },
			expectedValue: nil,
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
			eval:          func(*grammar.Production, []*Value) (any, error) { return nil, nil },
			expectedValue: nil,
			expectedErrorStrings: []string{
				`no action exists in the parsing table for ACTION[0, "+"]`,
			},
		},
		{
			name: "EvaluateFuncError",
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
			eval:          func(*grammar.Production, []*Value) (any, error) { return nil, errors.New("invalid semantic") },
			expectedValue: nil,
			expectedErrorStrings: []string{
				`invalid semantic`,
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
			eval: func(*grammar.Production, []*Value) (any, error) { return 69, nil },
			expectedValue: &Value{
				Val: 69,
				Pos: &lexer.Position{
					Filename: "test",
					Offset:   0,
					Line:     1,
					Column:   1,
				},
			},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.p.ParseAndEvaluate(tc.eval)

			if len(tc.expectedErrorStrings) == 0 {
				assert.True(t, reflect.DeepEqual(val, tc.expectedValue))
				assert.NoError(t, err)
			} else {
				assert.Nil(t, val)
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	tests := []struct {
		name           string
		v              *Value
		expectedString string
	}{
		{
			name:           "Zero",
			v:              &Value{},
			expectedString: `<nil>`,
		},
		{
			name: "OK",
			v: &Value{
				Val: 69,
				Pos: &lexer.Position{
					Filename: "test",
					Offset:   0,
					Line:     1,
					Column:   1,
				},
			},
			expectedString: `69 <test:1:1>`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.v.String())
		})
	}
}
