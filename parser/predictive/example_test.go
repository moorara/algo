package predictive_test

import (
	"fmt"
	"io"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/lexer/input"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/predictive"
)

// ExprLexer is an implementation of lexer.Lexer for basic math expressions.
type ExprLexer struct {
	in    *input.Input
	state int
}

func NewExprLexer(src io.Reader) (lexer.Lexer, error) {
	in, err := input.New("expression", src, 4096)
	if err != nil {
		return nil, err
	}

	return &ExprLexer{
		in: in,
	}, nil
}

func (l *ExprLexer) NextToken() (lexer.Token, error) {
	var r rune
	var err error

	// Reads runes from the input and feeds them into the DFA.
	for r, err = l.in.Next(); err == nil; r, err = l.in.Next() {
		if token, ok := l.advanceDFA(r); ok {
			return token, nil
		}
	}

	// Process last lexeme.
	if err == io.EOF {
		return l.finalizeDFA()
	}

	return lexer.Token{}, err
}

// advanceDFA simulates a deterministic finite automata to identify tokens.
func (l *ExprLexer) advanceDFA(r rune) (lexer.Token, bool) {
	// Determine the next state based on the current state and input.
	switch l.state {
	case 0:
		switch r {
		case ' ', '\t', '\n':
			l.state = 0
		case '+', '-', '*', '/', '(', ')':
			l.state = 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.state = 2
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			l.state = 4
		}

	case 2:
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.state = 2
		case ' ', '\t', '\n',
			'+', '-', '*', '/', '(', ')',
			'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			l.state = 3
		}

	case 4:
		switch r {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			l.state = 4
		case ' ', '\t', '\n',
			'+', '-', '*', '/', '(', ')',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.state = 5
		}

	default:
		panic("WTF?")
	}

	// Create and return a token based on the current state.
	switch l.state {
	case 0:
		l.in.Skip()

	// +  -  *  /  (  )
	case 1:
		l.state = 0
		lexeme, pos := l.in.Lexeme()
		return lexer.Token{Terminal: grammar.Terminal(r), Lexeme: lexeme, Pos: pos}, true

		// Number
	case 3:
		l.state = 0
		l.in.Retract()
		lexeme, pos := l.in.Lexeme()
		return lexer.Token{Terminal: grammar.Terminal("num"), Lexeme: lexeme, Pos: pos}, true

		// Identifier
	case 5:
		l.state = 0
		l.in.Retract()
		lexeme, pos := l.in.Lexeme()
		return lexer.Token{Terminal: grammar.Terminal("id"), Lexeme: lexeme, Pos: pos}, true
	}

	return lexer.Token{}, false
}

// finalizeDFA is called after all inputs have been processed by the DFA.
// It generates the final token based on the current state of the lexer.
func (l *ExprLexer) finalizeDFA() (lexer.Token, error) {
	lexeme, pos := l.in.Lexeme()

	switch l.state {
	case 2:
		l.state = 0
		return lexer.Token{Terminal: grammar.Terminal("num"), Lexeme: lexeme, Pos: pos}, nil
	case 4:
		l.state = 0
		return lexer.Token{Terminal: grammar.Terminal("id"), Lexeme: lexeme, Pos: pos}, nil
	default:
		return lexer.Token{}, io.EOF
	}
}

func Example_parse() {
	src := strings.NewReader(`
		(price + tax * quantity) * 
			(discount + shipping) * 
		(weight + volume) + total
	`)

	l, err := NewExprLexer(src)
	if err != nil {
		panic(err)
	}

	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]*grammar.Production{
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

	p := predictive.New(G, l)

	err = p.Parse(
		func(token *lexer.Token) error {
			fmt.Printf("Token: %s\n", token)
			return nil
		},
		func(prod *grammar.Production) error {
			fmt.Printf("Production: %s\n", prod)
			return nil
		},
	)

	if err != nil {
		panic(err)
	}
}

// You can copy-paste the output of this example into https://edotor.net to view the result.
func Example_parseAndBuildAST() {
	src := strings.NewReader(`
		(price + tax * quantity) * 
			(discount + shipping) * 
		(weight + volume) + total
	`)

	l, err := NewExprLexer(src)
	if err != nil {
		panic(err)
	}

	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]*grammar.Production{
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

	p := predictive.New(G, l)

	root, err := p.ParseAndBuildAST()
	if err != nil {
		panic(err)
	}

	n := root.(*parser.InternalNode)
	fmt.Println(n.DOT())
}

func Example_buildParsingTable() {
	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		[]*grammar.Production{
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

	table, err := predictive.BuildParsingTable(G)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)
}
