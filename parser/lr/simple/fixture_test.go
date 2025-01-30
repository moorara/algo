package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/parser/lr"
)

var statemaps = []lr.StateMap{
	{
		{ // I0
			&lr.Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
			&lr.Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&lr.Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&lr.Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
		},
		{ // I1
			&lr.Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
		},
		{ // I2
			&lr.Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1}, // E → T•
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
		},
		{ // I3
			&lr.Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1}, // T → F•
		},
		{ // I4
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1}, // F → (•E )
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
			&lr.Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&lr.Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&lr.Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
		},
		{ // I5
			&lr.Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1}, // F → id•
		},
		{ // I6
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
			&lr.Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&lr.Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
		},
		{ // I7
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2}, // T → T *•F
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
			&lr.Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
		},
		{ // I8
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
		},
		{ // I9
			&lr.Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
		},
		{ // I10
			&lr.Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3}, // T → T * F•
		},
		{ // I11
			&lr.Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3}, // F → ( E )•
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
	pt0.AddACTION(2, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(2, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(2, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt0.AddACTION(3, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(3, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(3, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt0.AddACTION(4, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(4, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(4, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt0.AddACTION(5, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(6, ")", &lr.Action{Type: lr.SHIFT, State: 3})
	pt0.AddACTION(6, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(7, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(7, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(8, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(8, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(8, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt0.AddACTION(9, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(9, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(10, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(10, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(10, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(10, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt0.AddACTION(11, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(11, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(11, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt0.AddACTION(11, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})

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
