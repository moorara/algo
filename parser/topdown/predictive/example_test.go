package predictive_test

import (
	"fmt"
	"io"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/lexer/input"
	"github.com/moorara/algo/parser/topdown/predictive"
)

type exprLexer struct {
	in *input.Input
}

func NewExprLexer(src io.Reader) (lexer.Lexer, error) {
	in, err := input.New("expression", src, 4096)
	if err != nil {
		return nil, err
	}

	return &exprLexer{
		in: in,
	}, nil
}

func (l *exprLexer) NextToken() (lexer.Token, error) {
	var state int
	var r rune
	var err error

	// Reads runes from the input and feeds them into the DFA.
	for r, err = l.in.Next(); err == nil; r, err = l.in.Next() {
		state = l.advanceDFA(state, r)
		if token, ok := l.evalDFA(state, r); ok {
			return token, nil
		}
	}

	// Process last lexeme.
	if err == io.EOF && state == 5 {
		lexeme, pos := l.in.Lexeme()
		return lexer.Token{Terminal: grammar.Terminal("id"), Lexeme: lexeme, Pos: pos}, nil
	}

	return lexer.Token{}, err
}

// advanceDFA simulates a determinist finite automata.
func (l *exprLexer) advanceDFA(state int, r rune) int {
	switch state {
	case 0:
		switch r {
		case ' ', '\t', '\n':
			return 1
		case '+', '-', '*', '/', '(', ')':
			return 3
		// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 4
		}

	case 1:
		switch r {
		case ' ', '\t', '\n':
			return 1
		case '+', '-', '*', '/', '(', ')':
			return 2
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 2
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 2
		}

	case 2:
		switch r {
		case ' ', '\t', '\n':
			return 1
		case '+', '-', '*', '/', '(', ')':
			return 3
		// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 4
		}

	case 3:
		switch r {
		case ' ', '\t', '\n':
			return 1
		case '+', '-', '*', '/', '(', ')':
			return 3
		// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 4
		}

	case 4:
		switch r {
		case ' ', '\t', '\n':
			return 6
		case '+', '-', '*', '/', '(', ')':
			return 6
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 5
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 5
		}

	case 5:
		switch r {
		case ' ', '\t', '\n':
			return 6
		case '+', '-', '*', '/', '(', ')':
			return 6
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 5
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 5
		}

	case 6:
		switch r {
		case ' ', '\t', '\n':
			return 1
		case '+', '-', '*', '/', '(', ')':
			return 3
		// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			return 4
		}
	}

	return -1
}

// evalDFA evaluates the state of the DFA after processing the rune r.
func (l *exprLexer) evalDFA(state int, r rune) (lexer.Token, bool) {
	switch state {
	// Ignore whitespaces
	case 2:
		l.in.Retract()
		l.in.Skip()

	// Operations
	case 3:
		lexeme, pos := l.in.Lexeme()

		switch r {
		case '+':
			return lexer.Token{Terminal: grammar.Terminal("+"), Lexeme: lexeme, Pos: pos}, true
		case '-':
			return lexer.Token{Terminal: grammar.Terminal("-"), Lexeme: lexeme, Pos: pos}, true
		case '*':
			return lexer.Token{Terminal: grammar.Terminal("*"), Lexeme: lexeme, Pos: pos}, true
		case '/':
			return lexer.Token{Terminal: grammar.Terminal("/"), Lexeme: lexeme, Pos: pos}, true
		case '(':
			return lexer.Token{Terminal: grammar.Terminal("("), Lexeme: lexeme, Pos: pos}, true
		case ')':
			return lexer.Token{Terminal: grammar.Terminal(")"), Lexeme: lexeme, Pos: pos}, true
		}

	// Identifier
	case 6:
		l.in.Retract()
		lexeme, pos := l.in.Lexeme()
		return lexer.Token{Terminal: grammar.Terminal("id"), Lexeme: lexeme, Pos: pos}, true
	}

	return lexer.Token{}, false
}

func Example() {
	src := strings.NewReader(`
		(price + tax * quantity) * 
			discount + shipping * 
		(weight + volume) + total
	`)

	l, err := NewExprLexer(src)
	if err != nil {
		panic(err)
	}

	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},                         // E → T E′
			{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}}, // E′ → + T E′
			{Head: "E′", Body: grammar.E}, // E′ → ε
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}},                         // T → F T′
			{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}}, // T′ → * F T′
			{Head: "T′", Body: grammar.E}, // T′ → ε
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}, // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                 // F → id
		},
		"E",
	)

	parser := predictive.New(G, l)
	err = parser.Parse(func(P grammar.Production, token lexer.Token) {
		fmt.Printf("%s\n%s\n\n", P, token)
	})

	if err != nil {
		panic(err)
	}
}
