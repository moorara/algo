package parsertest

import "github.com/moorara/algo/grammar"

var Prods = [][]*grammar.Production{
	{ // G0
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}},                         // E → T E′
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}}, // E′ → + T E′
		{Head: "E′", Body: grammar.E}, // E′ → ε
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}},                         // T → F T′
		{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}}, // T′ → * F T′
		{Head: "T′", Body: grammar.E}, // T′ → ε
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}}, // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                 // F → id
	},
	{ // G1
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                          // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")}}, // S → CC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")}},    // C → cC
		{Head: "C", Body: grammar.String[grammar.Symbol]{grammar.Terminal("d")}},                              // C → d
	},
	{ // G2
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                                                 // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L"), grammar.Terminal("="), grammar.NonTerminal("R")}}, // S → L = R
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("R")}},                                                  // S → R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("R")}},                           // L → *R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // L → id
		{Head: "R", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L")}},                                                  // R → L
	},
	{ // G3
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
	{ // G4
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
	},
}

var Grammars = []*grammar.CFG{
	// G0
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
		Prods[0],
		"E",
	),
	// G1
	grammar.NewCFG(
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
		Prods[1][1:],
		"S",
	),
	// G2
	grammar.NewCFG(
		[]grammar.Terminal{"=", "*", "id"},
		[]grammar.NonTerminal{"S", "L", "R"},
		Prods[2][1:],
		"S",
	),
	// G3
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		Prods[3][1:],
		"E",
	),
	// G4
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		Prods[4][1:],
		"E",
	),
}
