package lr

import "github.com/moorara/algo/grammar"

var starts = []grammar.NonTerminal{
	"S′",
	"S′",
	"E′",
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
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
		prods[0][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "*", "id"},
		[]grammar.NonTerminal{"S", "L", "R"},
		prods[1][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[2][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		prods[3][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[4][1:],
		"grammar",
	),
}

var LR0ItemSets = []ItemSet{
	NewItemSet( // I0
		// Kernels
		&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
		// Non-Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	),
	NewItemSet( // I1
		// Kernels
		&Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
	),
	NewItemSet( // I2
		// Kernels
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 1}, // E → T•
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
	),
	NewItemSet( // I3
		// Kernels
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 1}, // T → F•
	),
	NewItemSet( // I4
		// Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E )
		// Non-Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	),
	NewItemSet( // I5
		// Kernels
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 1}, // F → id•
	),
	NewItemSet( // I6
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
		// Non-Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	),
	NewItemSet( // I7
		// Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 2}, // T → T *•F
		// Non-Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	),
	NewItemSet( // I8
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E• + T
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
	),
	NewItemSet( // I9
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
	),
	NewItemSet( // I10
		// Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 3}, // T → T * F•
	),
	NewItemSet( // I11
		// Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 3}, // F → ( E )•
	),
}

var LR1ItemSets = []ItemSet{
	NewItemSet( //I0
		// Kernels
		&Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		// Non-Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	),
	NewItemSet( //I1
		// Kernels
		&Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
	),
	NewItemSet( //I2
		// Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	),
	NewItemSet( //I3
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	),
	NewItemSet( //I4
		// Kernels
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
	),
	NewItemSet( //I5
		// Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
	),
	NewItemSet( //I6
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	),
	NewItemSet( //I7
		// Kernels
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
	),
	NewItemSet( //I8
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
	),
	NewItemSet( //I9
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
	),
}

var statemaps = []StateMap{
	{
		{
			&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
			&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
			&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
			&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		},
		{
			&Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
		},
		{
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
		},
		{
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 3}, // F → ( E )•
		},
		{
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 3}, // T → T * F•
		},
		{
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
			&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
			&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		},
		{
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
		},
		{
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 2}, // T → T *•F
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
			&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
		},
		{
			&Item0{Production: prods[2][2], Start: starts[2], Dot: 1}, // E → T•
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
		},
		{
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E )
			&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
			&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
			&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
			&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
			&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
			&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		},
		{
			&Item0{Production: prods[2][6], Start: starts[2], Dot: 1}, // F → id•
		},
		{
			&Item0{Production: prods[2][4], Start: starts[2], Dot: 1}, // T → F•
		},
	},
}

var actions = []*Action{
	{
		Type:  SHIFT,
		State: 5,
	},
	{
		Type:  SHIFT,
		State: 7,
	},
	{
		Type: REDUCE,
		Production: &grammar.Production{ // E → T
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")},
		},
	},
	{
		Type: REDUCE,
		Production: &grammar.Production{ // F → id
			Head: "F",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")},
		},
	},
	{
		Type: ACCEPT,
	},
	{
		Type: ERROR,
	},
}

func getTestParsingTables() []*ParsingTable {
	pt0 := NewParsingTable(
		statemaps[0],
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt0.AddACTION(0, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(0, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(1, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &Action{Type: ACCEPT})
	pt0.AddACTION(2, "+", &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(2, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(2, ")", &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(2, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(3, "+", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, "*", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, ")", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(4, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(4, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(5, "+", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, "*", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, ")", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(6, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(6, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(7, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(7, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(8, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(8, ")", &Action{Type: SHIFT, State: 11})
	pt0.AddACTION(9, "+", &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(9, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(9, ")", &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(9, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(10, "+", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, "*", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, ")", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(11, "+", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, "*", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, ")", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][5]})

	pt0.SetGOTO(0, "E", 1)
	pt0.SetGOTO(0, "T", 2)
	pt0.SetGOTO(0, "F", 3)
	pt0.SetGOTO(4, "E", 8)
	pt0.SetGOTO(4, "T", 2)
	pt0.SetGOTO(4, "F", 3)
	pt0.SetGOTO(6, "T", 9)
	pt0.SetGOTO(6, "F", 3)
	pt0.SetGOTO(7, "F", 10)

	pt1 := NewParsingTable(
		[][]Item{{}, {}, {}, {}, {}, {}, {}},
		[]grammar.Terminal{"a", "b", "c", "d", grammar.Endmarker},
		[]grammar.NonTerminal{"A", "B", "C", "D"},
	)

	pt1.AddACTION(0, "a", &Action{
		Type:  SHIFT,
		State: 5,
	})

	pt1.AddACTION(0, "a", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "A",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("a"), grammar.NonTerminal("A")},
		},
	})

	pt1.AddACTION(1, "b", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "B",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("b"), grammar.NonTerminal("B")},
		},
	})

	pt1.AddACTION(1, "b", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "C",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")},
		},
	})

	return []*ParsingTable{pt0, pt1}
}
