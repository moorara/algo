package parsertest

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/stretchr/testify/assert"
)

func TestMockLexer(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		l := &MockLexer{
			NextTokenMocks: []NextTokenMock{
				// First token
				{OutError: errors.New("cannot read rune")},
			},
		}

		token, err := l.NextToken()

		assert.Zero(t, token)
		assert.EqualError(t, err, "cannot read rune")
	})

	t.Run("OK", func(t *testing.T) {
		l := &MockLexer{
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
			},
		}

		token, err := l.NextToken()

		assert.NoError(t, err)
		assert.True(t, token.Equal(lexer.Token{
			Terminal: grammar.Terminal("id"),
			Lexeme:   "a",
			Pos: lexer.Position{
				Filename: "test",
				Offset:   0,
				Line:     1,
				Column:   1,
			},
		}))
	})
}

func TestExprLexer(t *testing.T) {
	tests := []struct {
		name           string
		src            io.Reader
		expectedTokens []lexer.Token
	}{
		{
			name: "ID",
			src:  strings.NewReader(`a + b   *c`),
			expectedTokens: []lexer.Token{
				{
					Terminal: "id",
					Lexeme:   "a",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   0,
						Line:     1,
						Column:   1,
					},
				},
				{
					Terminal: "+",
					Lexeme:   "+",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   2,
						Line:     1,
						Column:   3,
					},
				},
				{
					Terminal: "id",
					Lexeme:   "b",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   4,
						Line:     1,
						Column:   5,
					},
				},
				{
					Terminal: "*",
					Lexeme:   "*",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   8,
						Line:     1,
						Column:   9,
					},
				},
				{
					Terminal: "id",
					Lexeme:   "c",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   9,
						Line:     1,
						Column:   10,
					},
				},
			},
		},
		{
			name: "Number",
			src:  strings.NewReader(`69 + 9   *3`),
			expectedTokens: []lexer.Token{
				{
					Terminal: "num",
					Lexeme:   "69",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   0,
						Line:     1,
						Column:   1,
					},
				},
				{
					Terminal: "+",
					Lexeme:   "+",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   3,
						Line:     1,
						Column:   4,
					},
				},
				{
					Terminal: "num",
					Lexeme:   "9",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   5,
						Line:     1,
						Column:   6,
					},
				},
				{
					Terminal: "*",
					Lexeme:   "*",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   9,
						Line:     1,
						Column:   10,
					},
				},
				{
					Terminal: "num",
					Lexeme:   "3",
					Pos: lexer.Position{
						Filename: "expression",
						Offset:   10,
						Line:     1,
						Column:   11,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l, err := NewExprLexer(tc.src)
			assert.NoError(t, err)

			for _, expectedToken := range tc.expectedTokens {
				token, err := l.NextToken()
				assert.NoError(t, err)
				assert.True(t, token.Equal(expectedToken))
			}
		})
	}
}
