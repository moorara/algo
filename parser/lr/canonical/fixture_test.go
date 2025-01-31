package canonical

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/parser/lr"
)

var statemaps = []lr.StateMap{
	{
		{ //I0
			&lr.Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},     // S′ → •S, $
			&lr.Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
		},
		{ //I1
			&lr.Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{ //I2
			&lr.Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
		},
		{ //I3
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
		},
		{ //I4
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
		},
		{ //I5
			&lr.Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
		},
		{ //I6
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
		},
		{ //I7
			&lr.Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
		},
		{ //I8
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
		},
		{ //I9
			&lr.Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
		},
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		statemaps[0].States(),
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
	pt0.AddACTION(7, "c", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][3]})
	pt0.AddACTION(7, "d", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][3]})
	pt0.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][1]})
	pt0.AddACTION(6, "c", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "d", &lr.Action{Type: lr.SHIFT, State: 8})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][3]})
	pt0.AddACTION(2, "c", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][2]})
	pt0.AddACTION(2, "d", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][2]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[1][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "C", 9)
	pt0.SetGOTO(9, "C", 4)
	pt0.SetGOTO(5, "C", 2)
	pt0.SetGOTO(6, "C", 3)

	return []*lr.ParsingTable{pt0}
}
