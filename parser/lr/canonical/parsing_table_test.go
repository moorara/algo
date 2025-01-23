package canonical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]grammar.Terminal{"c", "d"},
		[]grammar.NonTerminal{"S", "C"},
	)

	pt0.AddACTION(0, "c", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(0, "d", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(1, grammar.Endmarker, lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(9, "c", lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(9, "d", lr.Action{Type: lr.SHIFT, State: 8})
	pt0.AddACTION(5, "c", lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(5, "d", lr.Action{Type: lr.SHIFT, State: 7})
	pt0.AddACTION(7, "c", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(7, "d", lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(4, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][1]})
	pt0.AddACTION(6, "c", lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "d", lr.Action{Type: lr.SHIFT, State: 8})
	pt0.AddACTION(8, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][3]})
	pt0.AddACTION(2, "c", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(2, "d", lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})
	pt0.AddACTION(3, grammar.Endmarker, lr.Action{Type: lr.REDUCE, Production: &prods[0][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "C", 9)
	pt0.SetGOTO(9, "C", 4)
	pt0.SetGOTO(5, "C", 2)
	pt0.SetGOTO(6, "C", 3)

	pt1 := lr.NewParsingTable(
		[]lr.State{},
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
			assert.NotNil(t, calc.FIRST)
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
			expectedInitial: LR1Item{
				Production: &prods[0][0],
				Start:      &starts[0],
				Dot:        0,
				Lookahead:  grammar.Endmarker,
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
	g := lr.Augment(grammars[0])

	tests := []struct {
		name            string
		c               *calculator
		I               lr.ItemSet
		expectedCLOSURE lr.ItemSet
	}{
		{
			name: "OK",
			c: &calculator{
				augG:  g,
				FIRST: g.ComputeFIRST(),
			},
			I: lr.NewItemSet(
				LR1Item{Production: &prods[0][0], Start: &starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
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
				`8 errors occurred:`,
				`shift/reduce conflict at ACTION[2, "*"]`,
				`shift/reduce conflict at ACTION[2, "+"]`,
				`shift/reduce conflict at ACTION[3, "*"]`,
				`shift/reduce conflict at ACTION[3, "+"]`,
				`shift/reduce conflict at ACTION[4, "*"]`,
				`shift/reduce conflict at ACTION[4, "+"]`,
				`shift/reduce conflict at ACTION[5, "*"]`,
				`shift/reduce conflict at ACTION[5, "+"]`,
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
