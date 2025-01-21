package slr

import "github.com/moorara/algo/lexer"

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
