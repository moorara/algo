package parsertest

import (
	"github.com/moorara/algo/lexer"
)

// MockLexer is an implementation of lexer.Lexer for testing purposes.
type MockLexer struct {
	NextTokenIndex int
	NextTokenMocks []NextTokenMock
}

type NextTokenMock struct {
	OutToken lexer.Token
	OutError error
}

func (m *MockLexer) NextToken() (lexer.Token, error) {
	i := m.NextTokenIndex
	m.NextTokenIndex++
	return m.NextTokenMocks[i].OutToken, m.NextTokenMocks[i].OutError
}
