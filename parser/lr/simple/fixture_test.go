package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

var starts = []grammar.NonTerminal{
	"E′",
	"grammar′",
}

var prods = [][]*grammar.Production{
	{
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
	{
		{Head: "grammar′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("grammar")}},                           // grammar′ → grammar
		{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
		{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT")}},       // name → "grammar" IDENT
		{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}},   // decls → decls decl
		{Head: "decls", Body: grammar.E}, // decls → ε
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                  // decl → token
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                   // decl → rule
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING")}}, // token → TOKEN "=" STRING
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX")}},  // token → TOKEN "=" REGEX
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs")}}, // rule → lhs "=" rhs
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("=")}},                             // rule → lhs "="
		{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // lhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                         // rhs → rhs rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},  // rhs → rhs "|" rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},       // rhs → "(" rhs ")"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},       // rhs → "[" rhs "]"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},       // rhs → "{" rhs "}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},     // rhs → "{{" rhs "}}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // rhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                    // rhs → term
		{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                  // nonterm → IDENT
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                     // term → TOKEN
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                    // term → STRING
	},
}

var grammars = []*grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[0][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[1][1:],
		"grammar",
	),
}

var statemaps = []lr.StateMap{
	{
		{ // I0
			&lr.Item0{Production: prods[0][0], Start: starts[0], Dot: 0}, // E′ → •E
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 0}, // E → •E + T
			&lr.Item0{Production: prods[0][2], Start: starts[0], Dot: 0}, // E → •T
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 0}, // T → •T * F
			&lr.Item0{Production: prods[0][4], Start: starts[0], Dot: 0}, // T → •F
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 0}, // F → •( E )
			&lr.Item0{Production: prods[0][6], Start: starts[0], Dot: 0}, // F → •id
		},
		{ // I1
			&lr.Item0{Production: prods[0][0], Start: starts[0], Dot: 1}, // E′ → E•
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 1}, // E → E•+ T
		},
		{ // I2
			&lr.Item0{Production: prods[0][2], Start: starts[0], Dot: 1}, // E → T•
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 1}, // T → T•* F
		},
		{ // I3
			&lr.Item0{Production: prods[0][4], Start: starts[0], Dot: 1}, // T → F•
		},
		{ // I4
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 1}, // F → (•E )
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 0}, // E → •E + T
			&lr.Item0{Production: prods[0][2], Start: starts[0], Dot: 0}, // E → •T
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 0}, // T → •T * F
			&lr.Item0{Production: prods[0][4], Start: starts[0], Dot: 0}, // T → •F
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 0}, // F → •( E )
			&lr.Item0{Production: prods[0][6], Start: starts[0], Dot: 0}, // F → •id
		},
		{ // I5
			&lr.Item0{Production: prods[0][6], Start: starts[0], Dot: 1}, // F → id•
		},
		{ // I6
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 2}, // E → E +•T
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 0}, // T → •T * F
			&lr.Item0{Production: prods[0][4], Start: starts[0], Dot: 0}, // T → •F
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 0}, // F → •( E )
			&lr.Item0{Production: prods[0][6], Start: starts[0], Dot: 0}, // F → •id
		},
		{ // I7
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 2}, // T → T *•F
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 0}, // F → •( E )
			&lr.Item0{Production: prods[0][6], Start: starts[0], Dot: 0}, // F → •id
		},
		{ // I8
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 1}, // E → E•+ T
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 2}, // F → ( E•)
		},
		{ // I9
			&lr.Item0{Production: prods[0][1], Start: starts[0], Dot: 3}, // E → E + T•
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 1}, // T → T•* F
		},
		{ // I10
			&lr.Item0{Production: prods[0][3], Start: starts[0], Dot: 3}, // T → T * F•
		},
		{ // I11
			&lr.Item0{Production: prods[0][5], Start: starts[0], Dot: 3}, // F → ( E )•
		},
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		statemaps[0],
		[]grammar.Terminal{"(", ")", "*", "+", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt0.AddACTION(0, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(1, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(2, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(2, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(3, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(3, "*", &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(3, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(4, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, "*", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(5, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(6, ")", &lr.Action{Type: lr.SHIFT, State: 3})
	pt0.AddACTION(6, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(7, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(7, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(8, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})
	pt0.AddACTION(8, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(8, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})
	pt0.AddACTION(9, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(9, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(10, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][6]})
	pt0.AddACTION(10, "*", &lr.Action{Type: lr.REDUCE, Production: prods[0][6]})
	pt0.AddACTION(10, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][6]})
	pt0.AddACTION(10, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][6]})
	pt0.AddACTION(11, ")", &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(11, "*", &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(11, "+", &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(11, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})

	pt0.SetGOTO(0, "E", 1)
	pt0.SetGOTO(0, "T", 8)
	pt0.SetGOTO(0, "F", 11)
	pt0.SetGOTO(5, "T", 2)
	pt0.SetGOTO(5, "F", 11)
	pt0.SetGOTO(7, "F", 4)
	pt0.SetGOTO(9, "E", 6)
	pt0.SetGOTO(9, "T", 8)
	pt0.SetGOTO(9, "F", 11)

	return []*lr.ParsingTable{pt0}
}
