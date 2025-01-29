package canonical

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

var starts = []grammar.NonTerminal{
	"S′",
	"E′",
	"grammar′",
}

var prods = [][]*grammar.Production{
	{
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                          // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")}}, // S → CC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")}},    // C → cC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("d")}},                              // C → d
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
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
		prods[0][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		prods[1][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[2][1:],
		"grammar",
	),
}

var statemaps = []lr.StateMap{
	{
		{ //I0
			&lr.Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},     // S′ → •S, $
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
		},
		{ //I1
			&lr.Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{ //I2
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
		},
		{ //I3
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
		},
		{ //I4
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
		},
		{ //I5
			&lr.Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
		},
		{ //I6
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
		},
		{ //I7
			&lr.Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
		},
		{ //I8
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
		},
		{ //I9
			&lr.Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
		},
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		statemaps[0],
		[]grammar.Terminal{"c", "d", grammar.Endmarker},
		[]grammar.NonTerminal{"S", "C"},
	)

	pt0.AddACTION(0, "c", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(0, "d", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(9, "c", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(9, "d", &lr.Action{Type: lr.SHIFT, State: 8})
	pt0.AddACTION(5, "c", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(5, "d", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(7, "c", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(7, "d", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(6, "c", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "d", &lr.Action{Type: lr.SHIFT, State: 8})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(2, "c", &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})
	pt0.AddACTION(2, "d", &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "C", 9)
	pt0.SetGOTO(9, "C", 4)
	pt0.SetGOTO(5, "C", 2)
	pt0.SetGOTO(6, "C", 3)

	return []*lr.ParsingTable{pt0}
}
