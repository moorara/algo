package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

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

func TestNewCalculator(t *testing.T) {
	tests := []struct {
		name string
		G    grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewCalculator(tc.G).(*calculator)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.augG)
		})
	}
}

func TestCalculator_G(t *testing.T) {
	tests := []struct {
		name string
		c    *calculator
	}{
		{
			name: "OK",
			c: &calculator{
				augG: lr.Augment(grammars[0]),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			G := tc.c.G()

			assert.True(t, G.Equals(tc.c.augG))
		})
	}
}

func TestCalculator_Initial(t *testing.T) {
	tests := []struct {
		name            string
		c               *calculator
		expectedInitial lr.Item
	}{
		{
			name: "OK",
			c: &calculator{
				augG: lr.Augment(grammars[0]),
			},
			expectedInitial: LR0Item{
				Production: &prods[0][0],
				Start:      &starts[0],
				Dot:        0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equals(tc.expectedInitial))
		})
	}
}

func TestCalculator_CLOSURE(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name            string
		c               *calculator
		I               lr.ItemSet
		expectedCLOSURE lr.ItemSet
	}{
		{
			name: "OK",
			c: &calculator{
				augG: lr.Augment(grammars[0]),
			},
			I: lr.NewItemSet(
				LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equals(tc.expectedCLOSURE))
		})
	}
}

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    grammar.CFG
		expectedTable        *lr.ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:          "1st",
			G:             grammars[0],
			expectedTable: pt[0],
		},
		{
			name: "2nd",
			G:    grammars[1],
			expectedErrorStrings: []string{
				`20 errors occurred:`,
				`shift/reduce conflict at ACTION[2, "("]`,
				`shift/reduce conflict at ACTION[2, "IDENT"]`,
				`shift/reduce conflict at ACTION[2, "STRING"]`,
				`shift/reduce conflict at ACTION[2, "TOKEN"]`,
				`shift/reduce conflict at ACTION[2, "["]`,
				`shift/reduce conflict at ACTION[2, "{"]`,
				`shift/reduce conflict at ACTION[2, "{{"]`,
				`shift/reduce conflict at ACTION[2, "|"]`,
				`shift/reduce conflict at ACTION[7, "IDENT"]`,
				`shift/reduce conflict at ACTION[7, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "("]`,
				`shift/reduce conflict at ACTION[14, "IDENT"]`,
				`shift/reduce conflict at ACTION[14, "STRING"]`,
				`shift/reduce conflict at ACTION[14, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "["]`,
				`shift/reduce conflict at ACTION[14, "{"]`,
				`shift/reduce conflict at ACTION[14, "{{"]`,
				`shift/reduce conflict at ACTION[14, "|"]`,
				`shift/reduce conflict at ACTION[19, "IDENT"]`,
				`shift/reduce conflict at ACTION[19, "TOKEN"]`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table, err := BuildParsingTable(tc.G)

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equals(tc.expectedTable))
			} else {
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}
