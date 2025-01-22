package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

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
				augG: lr.Augment(grammars[0]),
			},
			I: lr.NewItemSet(
				LR0Item{Production: &prods[0][0], Start: &g.Start, Dot: 0}, // E′ → •E
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
			table := BuildParsingTable(tc.G)
			err := table.Error()

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
