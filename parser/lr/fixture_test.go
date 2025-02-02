package lr

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
)

var LR0ItemSets = []ItemSet{
	NewItemSet( // I0
		// Kernels
		&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
		// Non-Kernels
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
		&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
		&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
		&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
	),
	NewItemSet( // I1
		// Kernels
		&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
	),
	NewItemSet( // I2
		// Kernels
		&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1}, // E → T•
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
	),
	NewItemSet( // I3
		// Kernels
		&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1}, // T → F•
	),
	NewItemSet( // I4
		// Kernels
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1}, // F → (•E )
		// Non-Kernels
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
		&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
		&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
		&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
	),
	NewItemSet( // I5
		// Kernels
		&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1}, // F → id•
	),
	NewItemSet( // I6
		// Kernels
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
		// Non-Kernels
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
		&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
		&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
	),
	NewItemSet( // I7
		// Kernels
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2}, // T → T *•F
		// Non-Kernels
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
		&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
	),
	NewItemSet( // I8
		// Kernels
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E• + T
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
	),
	NewItemSet( // I9
		// Kernels
		&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
	),
	NewItemSet( // I10
		// Kernels
		&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3}, // T → T * F•
	),
	NewItemSet( // I11
		// Kernels
		&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3}, // F → ( E )•
	),
}

var LR1ItemSets = []ItemSet{
	NewItemSet( //I0
		// Kernels
		&Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		// Non-Kernels
		&Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	),
	NewItemSet( //I1
		// Kernels
		&Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
	),
	NewItemSet( //I2
		// Kernels
		&Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
		// Non-Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	),
	NewItemSet( //I3
		// Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
		// Non-Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	),
	NewItemSet( //I4
		// Kernels
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
	),
	NewItemSet( //I5
		// Kernels
		&Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
	),
	NewItemSet( //I6
		// Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
		// Non-Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	),
	NewItemSet( //I7
		// Kernels
		&Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
	),
	NewItemSet( //I8
		// Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
	),
	NewItemSet( //I9
		// Kernels
		&Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
	),
}

var statemaps = []StateMap{
	{
		{
			&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
			&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		},
		{
			&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
		},
		{
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
		},
		{
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3}, // F → ( E )•
		},
		{
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3}, // T → T * F•
		},
		{
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		},
		{
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
		},
		{
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2}, // T → T *•F
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
		},
		{
			&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1}, // E → T•
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
		},
		{
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1}, // F → (•E )
			&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
			&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
			&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
			&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
		},
		{
			&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1}, // F → id•
		},
		{
			&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1}, // T → F•
		},
	},
}

var actions = [][]*Action{
	{
		{ // 0
			Type: ACCEPT,
		},
		{ // 1
			Type: ERROR,
		},
		{ // 2
			Type:  SHIFT,
			State: 5,
		},
		{ // 3
			Type:  SHIFT,
			State: 6,
		},
		{ // 4
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "E",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
			},
		},
		{ // 5
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "E",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
			},
		},
		{ // 6
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "E",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("<"), grammar.NonTerminal("E")},
			},
		},
	},
	{
		{ // 0
			Type: ACCEPT,
		},
		{ // 1
			Type: ERROR,
		},
		{ // 2
			Type:  SHIFT,
			State: 13,
		},
		{ // 3
			Type:  SHIFT,
			State: 27,
		},
		{ // 4
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "rhs",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")},
			},
		},
		{ // 5
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "rhs",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")},
			},
		},
		{ // 6
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "rhs",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")},
			},
		},
		{ // 7
			Type: REDUCE,
			Production: &grammar.Production{
				Head: "nonterm",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")},
			},
		},
	},
}

var handles = [][]*PrecedenceHandle{
	{
		{Terminal: TerminalPtr("+")}, // 0
		{Terminal: TerminalPtr("-")}, // 1
		{Terminal: TerminalPtr("*")}, // 2
		{Terminal: TerminalPtr("/")}, // 3
		{Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")},
		}},
		{Production: &grammar.Production{
			Head: "E",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")},
		}},
	},
	{
		{Terminal: TerminalPtr("=")},      // 0
		{Terminal: TerminalPtr("|")},      // 1
		{Terminal: TerminalPtr("(")},      // 2
		{Terminal: TerminalPtr("[")},      // 3
		{Terminal: TerminalPtr("{")},      // 4
		{Terminal: TerminalPtr("{{")},     // 5
		{Terminal: TerminalPtr("IDENT")},  // 6
		{Terminal: TerminalPtr("TOKEN")},  // 7
		{Terminal: TerminalPtr("STRING")}, // 8
		{Production: &grammar.Production{
			Head: "rhs",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")},
		}},
		{Production: &grammar.Production{
			Head: "rhs",
			Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")},
		}},
	},
}

var precedences = []PrecedenceLevels{
	{ // 0
		{
			Associativity: LEFT,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("*"),
				PrecedenceHandleForTerminal("/"),
			),
		},
		{
			Associativity: LEFT,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("+"),
				PrecedenceHandleForTerminal("-"),
			),
		},
		{
			Associativity: NONE,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("<"),
				PrecedenceHandleForTerminal(">"),
			),
		},
	},
	{ // 1
		{
			Associativity: LEFT,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForProduction(&grammar.Production{
					Head: "rhs",
					Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")},
				}),
			),
		},
		{
			Associativity: LEFT,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("("),
				PrecedenceHandleForTerminal("["),
				PrecedenceHandleForTerminal("{"),
				PrecedenceHandleForTerminal("{{"),
				PrecedenceHandleForTerminal("IDENT"),
				PrecedenceHandleForTerminal("TOKEN"),
				PrecedenceHandleForTerminal("STRING"),
			),
		},
		{
			Associativity: RIGHT,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("|"),
			),
		},
		{
			Associativity: NONE,
			Handles: NewPrecedenceHandles(
				PrecedenceHandleForTerminal("="),
			),
		},
	},
}

func getTestParsingTables() []*ParsingTable {
	pt0 := NewParsingTable(
		statemaps[0].States(),
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
		PrecedenceLevels{},
	)

	pt0.AddACTION(0, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(0, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(1, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &Action{Type: ACCEPT})
	pt0.AddACTION(2, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(2, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(2, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(2, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(3, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(3, "*", &Action{Type: REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(3, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(3, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(4, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(4, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(5, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(5, "*", &Action{Type: REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(5, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(5, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(6, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(6, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(7, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(7, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(8, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(8, ")", &Action{Type: SHIFT, State: 11})
	pt0.AddACTION(9, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(9, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(9, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(9, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(10, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(10, "*", &Action{Type: REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(10, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(10, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(11, "+", &Action{Type: REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(11, "*", &Action{Type: REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(11, ")", &Action{Type: REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(11, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[3][5]})

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
		[]State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E"},
		precedences[0],
	)

	pt1.AddACTION(0, "(", &Action{Type: SHIFT, State: 8})
	pt1.AddACTION(0, "id", &Action{Type: SHIFT, State: 9})
	pt1.AddACTION(1, "*", &Action{Type: SHIFT, State: 5})
	pt1.AddACTION(1, "+", &Action{Type: SHIFT, State: 6})
	pt1.AddACTION(1, grammar.Endmarker, &Action{Type: ACCEPT})
	pt1.AddACTION(2, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt1.AddACTION(2, "*", &Action{Type: SHIFT, State: 5})
	pt1.AddACTION(2, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt1.AddACTION(2, "+", &Action{Type: SHIFT, State: 6})
	pt1.AddACTION(2, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt1.AddACTION(2, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt1.AddACTION(3, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt1.AddACTION(3, "*", &Action{Type: SHIFT, State: 5})
	pt1.AddACTION(3, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt1.AddACTION(3, "+", &Action{Type: SHIFT, State: 6})
	pt1.AddACTION(3, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt1.AddACTION(3, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt1.AddACTION(4, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt1.AddACTION(4, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt1.AddACTION(4, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt1.AddACTION(4, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt1.AddACTION(5, "(", &Action{Type: SHIFT, State: 8})
	pt1.AddACTION(5, "id", &Action{Type: SHIFT, State: 9})
	pt1.AddACTION(6, "(", &Action{Type: SHIFT, State: 8})
	pt1.AddACTION(6, "id", &Action{Type: SHIFT, State: 9})
	pt1.AddACTION(7, ")", &Action{Type: SHIFT, State: 4})
	pt1.AddACTION(7, "*", &Action{Type: SHIFT, State: 5})
	pt1.AddACTION(7, "+", &Action{Type: SHIFT, State: 6})
	pt1.AddACTION(8, "(", &Action{Type: SHIFT, State: 8})
	pt1.AddACTION(8, "id", &Action{Type: SHIFT, State: 9})
	pt1.AddACTION(9, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt1.AddACTION(9, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt1.AddACTION(9, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt1.AddACTION(9, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})

	pt1.SetGOTO(0, "E", 1)
	pt1.SetGOTO(5, "E", 2)
	pt1.SetGOTO(6, "E", 3)
	pt1.SetGOTO(8, "E", 7)

	pt2 := NewParsingTable(
		[]State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E"},
		PrecedenceLevels{},
	)

	pt2.AddACTION(0, "(", &Action{Type: SHIFT, State: 8})
	pt2.AddACTION(0, "id", &Action{Type: SHIFT, State: 9})
	pt2.AddACTION(1, "*", &Action{Type: SHIFT, State: 5})
	pt2.AddACTION(1, "+", &Action{Type: SHIFT, State: 6})
	pt2.AddACTION(1, grammar.Endmarker, &Action{Type: ACCEPT})
	pt2.AddACTION(2, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt2.AddACTION(2, "*", &Action{Type: SHIFT, State: 5})
	pt2.AddACTION(2, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt2.AddACTION(2, "+", &Action{Type: SHIFT, State: 6})
	pt2.AddACTION(2, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt2.AddACTION(2, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][2]})
	pt2.AddACTION(3, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt2.AddACTION(3, "*", &Action{Type: SHIFT, State: 5})
	pt2.AddACTION(3, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt2.AddACTION(3, "+", &Action{Type: SHIFT, State: 6})
	pt2.AddACTION(3, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt2.AddACTION(3, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][1]})
	pt2.AddACTION(4, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt2.AddACTION(4, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt2.AddACTION(4, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt2.AddACTION(4, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][3]})
	pt2.AddACTION(5, "(", &Action{Type: SHIFT, State: 8})
	pt2.AddACTION(5, "id", &Action{Type: SHIFT, State: 9})
	pt2.AddACTION(6, "(", &Action{Type: SHIFT, State: 8})
	pt2.AddACTION(6, "id", &Action{Type: SHIFT, State: 9})
	pt2.AddACTION(7, ")", &Action{Type: SHIFT, State: 4})
	pt2.AddACTION(7, "*", &Action{Type: SHIFT, State: 5})
	pt2.AddACTION(7, "+", &Action{Type: SHIFT, State: 6})
	pt2.AddACTION(8, "(", &Action{Type: SHIFT, State: 8})
	pt2.AddACTION(8, "id", &Action{Type: SHIFT, State: 9})
	pt2.AddACTION(9, ")", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt2.AddACTION(9, "*", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt2.AddACTION(9, "+", &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})
	pt2.AddACTION(9, grammar.Endmarker, &Action{Type: REDUCE, Production: parsertest.Prods[4][4]})

	pt2.SetGOTO(0, "E", 1)
	pt2.SetGOTO(5, "E", 2)
	pt2.SetGOTO(6, "E", 3)
	pt2.SetGOTO(8, "E", 7)

	return []*ParsingTable{pt0, pt1, pt2}
}

func TerminalPtr(t grammar.Terminal) *grammar.Terminal {
	return &t
}
