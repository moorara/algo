package simple_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/parser/lr/simple"
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

	parser, err := simple.New(l, G)
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

	l, err := parsertest.NewExprLexer(src)
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

	parser, err := simple.New(l, G)
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

	l, err := parsertest.NewExprLexer(src)
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

	parser, err := simple.New(l, G)
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

	table, err := simple.BuildParsingTable(G)
	if err != nil {
		panic(err)
	}

	fmt.Println(table)
}
