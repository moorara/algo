package predictive_test

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser/predictive"
)

func Example_parse() {
	src := strings.NewReader(`
		(price + tax * quantity) * 
			(discount + shipping) * 
		(weight + volume) + total
	`)

	l, err := parsertest.NewExprLexer(src)
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

	parser := predictive.New(G, l)

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

	l, err := parsertest.NewExprLexer(src)
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

	parser := predictive.New(G, l)

	ast, err := parser.ParseAndBuildAST()
	if err != nil {
		panic(err)
	}

	fmt.Println(ast.DOT())
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

	table := predictive.BuildParsingTable(G)
	if err := table.Error(); err != nil {
		panic(err)
	}

	fmt.Println(table)
}
