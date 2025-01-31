package lookahead

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/parser/lr"
)

var statemaps = []lr.StateMap{
	{
		{ // I0
			&lr.Item1{Production: parsertest.Prods[2][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		},
		{ // I1
			&lr.Item1{Production: parsertest.Prods[2][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{ // I2
			&lr.Item1{Production: parsertest.Prods[2][1], Start: "S′", Dot: 3, Lookahead: grammar.Endmarker}, // S → L "=" R•, $
		},
		{ // I3
			&lr.Item1{Production: parsertest.Prods[2][3], Start: "S′", Dot: 2, Lookahead: "="},               // L → "*" R•, =
			&lr.Item1{Production: parsertest.Prods[2][3], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // L → "*" R•, $
		},
		{ // I4
			&lr.Item1{Production: parsertest.Prods[2][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → L "="•R, $
		},
		{ // I5
			&lr.Item1{Production: parsertest.Prods[2][3], Start: "S′", Dot: 1, Lookahead: "="},               // L → "*"•R, =
			&lr.Item1{Production: parsertest.Prods[2][3], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // L → "*"•R, $
		},
		{ // I6
			&lr.Item1{Production: parsertest.Prods[2][4], Start: "S′", Dot: 1, Lookahead: "="},               // L → "id"•, =
			&lr.Item1{Production: parsertest.Prods[2][4], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // L → "id"•, $
		},
		{ // I7
			&lr.Item1{Production: parsertest.Prods[2][5], Start: "S′", Dot: 1, Lookahead: "="},               // R → L•, =
			&lr.Item1{Production: parsertest.Prods[2][5], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // R → L•, $
		},
		{ // I8
			&lr.Item1{Production: parsertest.Prods[2][5], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // R → L•, $
			&lr.Item1{Production: parsertest.Prods[2][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → L•"=" R, $
		},
		{ // I9
			&lr.Item1{Production: parsertest.Prods[2][2], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → R•, $
		},
	},
	{
		{ // I0
			&lr.Item1{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0, Lookahead: grammar.Endmarker}, // E′ → •E, $
		},
		{ // I1
			&lr.Item1{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // E′ → E•, $
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1, Lookahead: "+"},               // E → E•+ T, +
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // E → E•+ T, $
		},
		{ // I2
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3, Lookahead: ")"},               // E → E + T•, )
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3, Lookahead: "+"},               // E → E + T•, +
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3, Lookahead: grammar.Endmarker}, // E → E + T•, $
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: ")"},               // T → T•* F, )
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: "*"},               // T → T•* F, *
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: "+"},               // T → T•* F, +
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // T → T•* F, $
		},
		{ // I3
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3, Lookahead: ")"},               // F → ( E )•, )
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3, Lookahead: "*"},               // F → ( E )•, *
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3, Lookahead: "+"},               // F → ( E )•, +
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3, Lookahead: grammar.Endmarker}, // F → ( E )•, $
		},
		{ // I4
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3, Lookahead: ")"},               // T → T * F•, )
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3, Lookahead: "*"},               // T → T * F•, *
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3, Lookahead: "+"},               // T → T * F•, +
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3, Lookahead: grammar.Endmarker}, // T → T * F•, $
		},
		{ // I5
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2, Lookahead: ")"},               // E → E +•T, )
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2, Lookahead: "+"},               // E → E +•T, +
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2, Lookahead: grammar.Endmarker}, // E → E +•T, $
		},
		{ // I6
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2, Lookahead: ")"},               // F → ( E•), )
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2, Lookahead: "*"},               // F → ( E•), *
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2, Lookahead: "+"},               // F → ( E•), +
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2, Lookahead: grammar.Endmarker}, // F → ( E•), $
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1, Lookahead: ")"},               // E → E•+ T, )
			&lr.Item1{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1, Lookahead: "+"},               // E → E•+ T, +
		},
		{ // I7
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2, Lookahead: ")"},               // T → T *•F, )
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2, Lookahead: "*"},               // T → T *•F, *
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2, Lookahead: "+"},               // T → T *•F, +
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2, Lookahead: grammar.Endmarker}, // T → T *•F, $
		},
		{ // I8
			&lr.Item1{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1, Lookahead: ")"},               // E → T•, )
			&lr.Item1{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1, Lookahead: "+"},               // E → T•, +
			&lr.Item1{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // E → T•, $
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: ")"},               // T → T•* F, )
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: "*"},               // T → T•* F, *
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: "+"},               // T → T•* F, +
			&lr.Item1{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // T → T•* F, $
		},
		{ // I9
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1, Lookahead: ")"},               // F → (•E ), )
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1, Lookahead: "*"},               // F → (•E ), *
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1, Lookahead: "+"},               // F → (•E ), +
			&lr.Item1{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // F → (•E ), $
		},
		{ // I10
			&lr.Item1{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1, Lookahead: ")"},               // F → id•, )
			&lr.Item1{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1, Lookahead: "*"},               // F → id•, *
			&lr.Item1{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1, Lookahead: "+"},               // F → id•, +
			&lr.Item1{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // F → id•, $
		},
		{ // I11
			&lr.Item1{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1, Lookahead: ")"},               // T → F•, )
			&lr.Item1{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1, Lookahead: "*"},               // T → F•, *
			&lr.Item1{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1, Lookahead: "+"},               // T → F•, +
			&lr.Item1{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1, Lookahead: grammar.Endmarker}, // T → F•, $
		},
	},
}

var kernelmaps = []lr.StateMap{
	{
		{
			&lr.Item0{Production: parsertest.Prods[2][0], Start: `S′`, Dot: 0}, // S′ → •S
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][0], Start: `S′`, Dot: 1}, // S′ → S•
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][1], Start: `S′`, Dot: 3}, // S → L "=" R•
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][3], Start: `S′`, Dot: 2}, // L → "*" R•
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][1], Start: `S′`, Dot: 2}, // S → L "="•R
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][3], Start: `S′`, Dot: 1}, // L → "*"•R
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][4], Start: `S′`, Dot: 1}, // L → "id"•
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][5], Start: `S′`, Dot: 1}, // R → L•
			&lr.Item0{Production: parsertest.Prods[2][1], Start: `S′`, Dot: 1}, // S → L•"=" R
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][5], Start: `S′`, Dot: 1}, // R → L•
		},
		{
			&lr.Item0{Production: parsertest.Prods[2][2], Start: `S′`, Dot: 1}, // S → R•
		},
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		statemaps[0].States(),
		[]grammar.Terminal{"=", "*", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"S", "L", "R"},
	)

	pt0.AddACTION(0, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][1]})
	pt0.AddACTION(3, "=", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][3]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][3]})
	pt0.AddACTION(4, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(4, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(5, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "=", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][4]})
	pt0.AddACTION(6, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][4]})
	pt0.AddACTION(7, "=", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][5]})
	pt0.AddACTION(7, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][5]})
	pt0.AddACTION(8, "=", &lr.Action{Type: lr.SHIFT, State: 4})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][5]})
	pt0.AddACTION(9, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[2][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "L", 8)
	pt0.SetGOTO(0, "R", 9)
	pt0.SetGOTO(4, "L", 7)
	pt0.SetGOTO(4, "R", 2)
	pt0.SetGOTO(5, "L", 7)
	pt0.SetGOTO(5, "R", 3)

	pt1 := lr.NewParsingTable(
		statemaps[1].States(),
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt1.AddACTION(0, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(1, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt1.AddACTION(2, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt1.AddACTION(2, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(2, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt1.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][1]})
	pt1.AddACTION(3, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt1.AddACTION(3, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt1.AddACTION(3, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt1.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][5]})
	pt1.AddACTION(4, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt1.AddACTION(4, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt1.AddACTION(4, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt1.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][3]})
	pt1.AddACTION(5, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(6, ")", &lr.Action{Type: lr.SHIFT, State: 3})
	pt1.AddACTION(6, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(7, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(7, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(8, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt1.AddACTION(8, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(8, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt1.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][2]})
	pt1.AddACTION(9, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(9, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(10, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt1.AddACTION(10, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt1.AddACTION(10, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt1.AddACTION(10, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][6]})
	pt1.AddACTION(11, ")", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt1.AddACTION(11, "*", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt1.AddACTION(11, "+", &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})
	pt1.AddACTION(11, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: parsertest.Prods[3][4]})

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
