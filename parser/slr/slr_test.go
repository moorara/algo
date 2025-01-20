package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

var starts = []grammar.NonTerminal{
	"E′",
}

var prods = [][]grammar.Production{
	{
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
}

var grammars = []grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[0][1:],
		"E",
	),
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]grammar.Terminal{"(", ")", "*", "+", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt0.AddACTION(0, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(0, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(1, "+", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(1, grammar.Endmarker, lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(2, "*", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(2, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(2, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(3, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(3, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][5]})
	pt0.AddACTION(4, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(5, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(5, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(6, ")", lr.Action{Type: lr.SHIFT, State: 3})
	pt0.AddACTION(6, "+", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(7, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(7, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(8, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(8, "*", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(8, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(8, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(9, "(", lr.Action{Type: lr.SHIFT, State: 9})
	pt0.AddACTION(9, "id", lr.Action{Type: lr.SHIFT, State: 10})
	pt0.AddACTION(10, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(10, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][6]})
	pt0.AddACTION(11, ")", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, "*", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, "+", lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})
	pt0.AddACTION(11, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][4]})

	pt0.SetGOTO(0, "E", 1)
	pt0.SetGOTO(0, "T", 8)
	pt0.SetGOTO(0, "F", 11)
	pt0.SetGOTO(5, "T", 2)
	pt0.SetGOTO(5, "F", 11)
	pt0.SetGOTO(7, "F", 4)
	pt0.SetGOTO(9, "E", 6)
	pt0.SetGOTO(9, "T", 8)
	pt0.SetGOTO(9, "F", 11)

	pt1 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36},
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "GRAMMAR", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
	)

	return []*lr.ParsingTable{pt0, pt1}
}

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name          string
		G             grammar.CFG
		expectedTable *lr.ParsingTable
	}{
		{
			name:          "OK",
			G:             grammars[0],
			expectedTable: pt[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table := BuildParsingTable(tc.G)

			assert.True(t, table.Equals(tc.expectedTable))
		})
	}
}
