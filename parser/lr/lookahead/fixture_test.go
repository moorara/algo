package lookahead

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

var starts = []grammar.NonTerminal{
	"S′",
	"E′",
	"E′",
	"grammar′",
}

var prods = [][]*grammar.Production{
	{
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                                                 // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L"), grammar.Terminal("="), grammar.NonTerminal("R")}}, // S → L = R
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("R")}},                                                  // S → R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("R")}},                           // L → *R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // L → id
		{Head: "R", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L")}},                                                  // R → L
	},
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
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
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
		[]grammar.Terminal{"=", "*", "id"},
		[]grammar.NonTerminal{"S", "L", "R"},
		prods[0][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[1][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		prods[2][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[3][1:],
		"grammar",
	),
}

var statemaps = []lr.StateMap{
	{
		{ // I0
			&lr.Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		},
		{ // I1
			&lr.Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{ // I2
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 3, Lookahead: grammar.Endmarker}, // S → L "=" R•, $
		},
		{ // I3
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 2, Lookahead: "="},               // L → "*" R•, =
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // L → "*" R•, $
		},
		{ // I4
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → L "="•R, $
		},
		{ // I5
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: "="},               // L → "*"•R, =
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // L → "*"•R, $
		},
		{ // I6
			&lr.Item1{Production: prods[0][4], Start: starts[0], Dot: 1, Lookahead: "="},               // L → "id"•, =
			&lr.Item1{Production: prods[0][4], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // L → "id"•, $
		},
		{ // I7
			&lr.Item1{Production: prods[0][5], Start: starts[0], Dot: 1, Lookahead: "="},               // R → L•, =
			&lr.Item1{Production: prods[0][5], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // R → L•, $
		},
		{ // I8
			&lr.Item1{Production: prods[0][5], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // R → L•, $
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → L•"=" R, $
		},
		{ // I9
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → R•, $
		},
	},
	{
		{ // I0
			&lr.Item1{Production: prods[1][0], Start: starts[1], Dot: 0, Lookahead: grammar.Endmarker}, // E′ → •E, $
		},
		{ // I1
			&lr.Item1{Production: prods[1][0], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // E′ → E•, $
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 1, Lookahead: "+"},               // E → E•+ T, +
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // E → E•+ T, $
		},
		{ // I2
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 3, Lookahead: ")"},               // E → E + T•, )
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 3, Lookahead: "+"},               // E → E + T•, +
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 3, Lookahead: grammar.Endmarker}, // E → E + T•, $
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: ")"},               // T → T•* F, )
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: "*"},               // T → T•* F, *
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: "+"},               // T → T•* F, +
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // T → T•* F, $
		},
		{ // I3
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 3, Lookahead: ")"},               // F → ( E )•, )
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 3, Lookahead: "*"},               // F → ( E )•, *
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 3, Lookahead: "+"},               // F → ( E )•, +
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 3, Lookahead: grammar.Endmarker}, // F → ( E )•, $
		},
		{ // I4
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 3, Lookahead: ")"},               // T → T * F•, )
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 3, Lookahead: "*"},               // T → T * F•, *
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 3, Lookahead: "+"},               // T → T * F•, +
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 3, Lookahead: grammar.Endmarker}, // T → T * F•, $
		},
		{ // I5
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 2, Lookahead: ")"},               // E → E +•T, )
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 2, Lookahead: "+"},               // E → E +•T, +
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 2, Lookahead: grammar.Endmarker}, // E → E +•T, $
		},
		{ // I6
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 2, Lookahead: ")"},               // F → ( E•), )
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 2, Lookahead: "*"},               // F → ( E•), *
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 2, Lookahead: "+"},               // F → ( E•), +
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 2, Lookahead: grammar.Endmarker}, // F → ( E•), $
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 1, Lookahead: ")"},               // E → E•+ T, )
			&lr.Item1{Production: prods[1][1], Start: starts[1], Dot: 1, Lookahead: "+"},               // E → E•+ T, +
		},
		{ // I7
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 2, Lookahead: ")"},               // T → T *•F, )
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 2, Lookahead: "*"},               // T → T *•F, *
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 2, Lookahead: "+"},               // T → T *•F, +
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 2, Lookahead: grammar.Endmarker}, // T → T *•F, $
		},
		{ // I8
			&lr.Item1{Production: prods[1][2], Start: starts[1], Dot: 1, Lookahead: ")"},               // E → T•, )
			&lr.Item1{Production: prods[1][2], Start: starts[1], Dot: 1, Lookahead: "+"},               // E → T•, +
			&lr.Item1{Production: prods[1][2], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // E → T•, $
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: ")"},               // T → T•* F, )
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: "*"},               // T → T•* F, *
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: "+"},               // T → T•* F, +
			&lr.Item1{Production: prods[1][3], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // T → T•* F, $
		},
		{ // I9
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 1, Lookahead: ")"},               // F → (•E ), )
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 1, Lookahead: "*"},               // F → (•E ), *
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 1, Lookahead: "+"},               // F → (•E ), +
			&lr.Item1{Production: prods[1][5], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // F → (•E ), $
		},
		{ // I10
			&lr.Item1{Production: prods[1][6], Start: starts[1], Dot: 1, Lookahead: ")"},               // F → id•, )
			&lr.Item1{Production: prods[1][6], Start: starts[1], Dot: 1, Lookahead: "*"},               // F → id•, *
			&lr.Item1{Production: prods[1][6], Start: starts[1], Dot: 1, Lookahead: "+"},               // F → id•, +
			&lr.Item1{Production: prods[1][6], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // F → id•, $
		},
		{ // I11
			&lr.Item1{Production: prods[1][4], Start: starts[1], Dot: 1, Lookahead: ")"},               // T → F•, )
			&lr.Item1{Production: prods[1][4], Start: starts[1], Dot: 1, Lookahead: "*"},               // T → F•, *
			&lr.Item1{Production: prods[1][4], Start: starts[1], Dot: 1, Lookahead: "+"},               // T → F•, +
			&lr.Item1{Production: prods[1][4], Start: starts[1], Dot: 1, Lookahead: grammar.Endmarker}, // T → F•, $
		},
	},
}

var kernelmaps = []lr.StateMap{
	{
		{
			&lr.Item0{Production: prods[0][0], Start: `S′`, Dot: 0}, // S′ → •S
		},
		{
			&lr.Item0{Production: prods[0][0], Start: `S′`, Dot: 1}, // S′ → S•
		},
		{
			&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 3}, // S → L "=" R•
		},
		{
			&lr.Item0{Production: prods[0][3], Start: `S′`, Dot: 2}, // L → "*" R•
		},
		{
			&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 2}, // S → L "="•R
		},
		{
			&lr.Item0{Production: prods[0][3], Start: `S′`, Dot: 1}, // L → "*"•R
		},
		{
			&lr.Item0{Production: prods[0][4], Start: `S′`, Dot: 1}, // L → "id"•
		},
		{
			&lr.Item0{Production: prods[0][5], Start: `S′`, Dot: 1}, // R → L•
			&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 1}, // S → L•"=" R
		},
		{
			&lr.Item0{Production: prods[0][5], Start: `S′`, Dot: 1}, // R → L•
		},
		{
			&lr.Item0{Production: prods[0][2], Start: `S′`, Dot: 1}, // S → R•
		},
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		statemaps[0],
		[]grammar.Terminal{"=", "*", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"S", "L", "R"},
	)

	pt0.AddACTION(0, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(3, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(4, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(5, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(6, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(7, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(7, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(8, "=", &lr.Action{Type: lr.SHIFT, State: 4})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(9, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "L", 8)
	pt0.SetGOTO(0, "R", 9)
	pt0.SetGOTO(4, "L", 7)
	pt0.SetGOTO(4, "R", 2)
	pt0.SetGOTO(5, "L", 7)
	pt0.SetGOTO(5, "R", 3)

	pt1 := lr.NewParsingTable(
		statemaps[1],
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt1.AddACTION(0, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(1, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt1.AddACTION(2, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(2, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(2, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(3, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(4, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(5, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(6, ")", &lr.Action{Type: lr.SHIFT, State: 3})
	pt1.AddACTION(6, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(7, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(7, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(8, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(8, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(8, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(9, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(9, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(10, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(11, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})

	pt1.SetGOTO(0, "E", 1)
	pt1.SetGOTO(0, "T", 8)
	pt1.SetGOTO(0, "F", 11)
	pt1.SetGOTO(5, "T", 2)
	pt1.SetGOTO(5, "F", 11)
	pt1.SetGOTO(7, "F", 4)
	pt1.SetGOTO(9, "E", 6)
	pt1.SetGOTO(9, "T", 8)
	pt1.SetGOTO(9, "F", 11)

	return []*lr.ParsingTable{pt0, pt1}
}
