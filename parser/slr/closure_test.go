package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

func TestInitial(t *testing.T) {
	g0 := lr.Augment(grammars[0])

	tests := []struct {
		name            string
		augG            grammar.CFG
		expectedInitial lr.Item
	}{
		{
			name: "OK",
			augG: g0,
			expectedInitial: LR0Item{
				Production: &prods[0][0],
				Start:      &starts[0],
				Dot:        0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.augG.Verify())
			initial := Initial(tc.augG)
			assert.True(t, initial.Equals(tc.expectedInitial))
		})
	}
}

func TestLR0Closure(t *testing.T) {
	s := getTestItemSets()
	g := lr.Augment(grammars[0])

	tests := []struct {
		name            string
		augG            grammar.CFG
		I               lr.ItemSet
		expectedCLOSURE lr.ItemSet
	}{
		{
			name: "OK",
			augG: g,
			I: lr.NewItemSet(
				LR0Item{Production: &prods[0][0], Start: &g.Start, Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.augG.Verify())
			CLOSURE := LR0Closure(tc.augG)
			J := CLOSURE(tc.I)
			assert.True(t, J.Equals(tc.expectedCLOSURE))
		})
	}
}
