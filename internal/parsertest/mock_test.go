package parsertest

import (
	"errors"
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
