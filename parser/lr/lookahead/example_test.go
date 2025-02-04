package lookahead_test

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/lexer/input"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/parser/lr/lookahead"
	"github.com/moorara/algo/parser/lr/simple"
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
		[]grammar.NonTerminal{"E", "T", "F"},
		[]*grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
		},
		"E",
	)

	parser, err := lookahead.New(l, G, lr.PrecedenceLevels{})
	if err != nil {
		panic(err)
	}

	err = parser.Parse(
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
		[]grammar.NonTerminal{"E", "T", "F"},
		[]*grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
		},
		"E",
	)

	parser, err := lookahead.New(l, G, lr.PrecedenceLevels{})
	if err != nil {
		panic(err)
	}

	ast, err := parser.ParseAndBuildAST()
	if err != nil {
		panic(err)
	}

	fmt.Println(ast.DOT())
}

func Example_parseAndEvaluate() {
	src := strings.NewReader(`69 + 9  * 3`)

	l, err := NewExprLexer(src)
	if err != nil {
		panic(err)
	}

	prods := []*grammar.Production{
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("num")}},                                                   // F → num
	}

	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "num"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods,
		"E",
	)

	parser, err := simple.New(l, G, lr.PrecedenceLevels{})
	if err != nil {
		panic(err)
	}

	val, err := parser.ParseAndEvaluate(func(p *grammar.Production, rhs []*lr.Value) (any, error) {
		switch {
		case p.Equal(prods[0]):
			E := rhs[0].Val.(int)
			T := rhs[2].Val.(int)
			return E + T, nil

		case p.Equal(prods[1]):
			return rhs[0].Val, nil

		case p.Equal(prods[2]):
			T := rhs[0].Val.(int)
			F := rhs[2].Val.(int)
			return T * F, nil

		case p.Equal(prods[3]):
			return rhs[0].Val, nil

		case p.Equal(prods[4]):
			return rhs[1].Val, nil

		case p.Equal(prods[5]):
			return strconv.Atoi(rhs[0].Val.(string))

		default:
			return fmt.Errorf("unexpected production: %s", p), nil
		}
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}

func Example_buildParsingTable() {
	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		[]*grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
			{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
			{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
		},
		"E",
	)

	table, err := lookahead.BuildParsingTable(G, lr.PrecedenceLevels{})
	if err != nil {
		panic(err)
	}

	fmt.Println(table)
}

// You can copy-paste the output of this example into https://edotor.net to view the result.
func Example_ambiguousGrammar() {
	src := strings.NewReader(`foo + bar   * baz`)

	l, err := NewExprLexer(src)
	if err != nil {
		panic(err)
	}

	G := grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		[]*grammar.Production{
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
			{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
		},
		"E",
	)

	precedences := lr.PrecedenceLevels{
		{
			Associativity: lr.LEFT,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("*"),
				lr.PrecedenceHandleForTerminal("/"),
			),
		},
		{
			Associativity: lr.LEFT,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("+"),
				lr.PrecedenceHandleForTerminal("-"),
			),
		},
	}

	parser, err := simple.New(l, G, precedences)
	if err != nil {
		panic(err)
	}

	ast, err := parser.ParseAndBuildAST()
	if err != nil {
		panic(err)
	}

	fmt.Println(ast.DOT())
}

func Example_ebnf() {
	G := grammar.NewCFG(
		[]grammar.Terminal{
			"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}",
			"grammar", "@left", "@right", "@none",
			"IDENT", "TOKEN", "STRING", "REGEX", "PREDEF",
		},
		[]grammar.NonTerminal{
			"grammar", "name", "decls", "decl", "token", "directive",
			"handles", "rule", "lhs", "rhs", "nonterm", "term",
		},
		[]*grammar.Production{ // G5
			{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
			{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT")}},       // name → "grammar" IDENT
			{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}},   // decls → decls decl
			{Head: "decls", Body: grammar.E}, // decls → ε
			{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                  // decl → token
			{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("directive")}},                                              // decl → directive
			{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                   // decl → rule
			{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING")}}, // token → TOKEN "=" STRING
			{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX")}},  // token → TOKEN "=" REGEX
			{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("PREDEF")}}, // token → TOKEN "=" PREDEF
			{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@left"), grammar.NonTerminal("handles")}},                // directive → "@left" handles
			{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@right"), grammar.NonTerminal("handles")}},               // directive → "@right" handles
			{Head: "directive", Body: grammar.String[grammar.Symbol]{grammar.Terminal("@none"), grammar.NonTerminal("handles")}},                // directive → "@none" handles
			{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("handles"), grammar.NonTerminal("term")}},                // handles → handles term
			{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("handles"), grammar.NonTerminal("rule")}},                // handles → handles rule
			{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                // handles → term
			{Head: "handles", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                // handles → rule
			{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs")}}, // rule → lhs "=" rhs
			{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("=")}},                             // rule → lhs "="
			{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // lhs → nonterm
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // rhs → nonterm
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                    // rhs → term
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                         // rhs → rhs rhs
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},       // rhs → "(" rhs ")"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},       // rhs → "[" rhs "]"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},       // rhs → "{" rhs "}"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},     // rhs → "{{" rhs "}}"
			{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},  // rhs → rhs "|" rhs
			{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                  // nonterm → IDENT
			{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                     // term → TOKEN
			{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                    // term → STRING
		},
		"grammar",
	)

	precedences := lr.PrecedenceLevels{
		{
			Associativity: lr.LEFT,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForProduction(&grammar.Production{
					Head: "rhs",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")},
				}),
			),
		},
		{
			Associativity: lr.LEFT,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("("),
				lr.PrecedenceHandleForTerminal("["),
				lr.PrecedenceHandleForTerminal("{"),
				lr.PrecedenceHandleForTerminal("{{"),
				lr.PrecedenceHandleForTerminal("IDENT"),
				lr.PrecedenceHandleForTerminal("TOKEN"),
				lr.PrecedenceHandleForTerminal("STRING"),
			),
		},
		{
			Associativity: lr.RIGHT,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("|"),
			),
		},
		{
			Associativity: lr.NONE,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("="),
			),
		},
		{
			Associativity: lr.NONE,
			Handles: lr.NewPrecedenceHandles(
				lr.PrecedenceHandleForTerminal("@left"),
				lr.PrecedenceHandleForTerminal("@right"),
				lr.PrecedenceHandleForTerminal("@none"),
			),
		},
	}

	table, err := lookahead.BuildParsingTable(G, precedences)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)
}
