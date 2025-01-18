package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

func TestCFG_AugmentCFG(t *testing.T) {
	tests := []struct {
		name              string
		G                 grammar.CFG
		expectedAugmented AugmentedCFG
	}{
		{
			name: "OK",
			G:    grammars[0],
			expectedAugmented: AugmentedCFG{
				CFG: grammar.NewCFG(
					[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
					[]grammar.NonTerminal{"E′", "E", "T", "F"},
					prods[0],
					"E′",
				),
				Initial: Item{
					Production: &prods[0][0],
					Initial:    true,
					Dot:        0,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			augG := AugmentCFG(tc.G)
			assert.True(t, augG.Equals(tc.expectedAugmented))
		})
	}
}

func TestAugmentedCFG_Equals(t *testing.T) {
	g0 := AugmentCFG(grammars[0])
	g1 := AugmentCFG(grammars[1])

	tests := []struct {
		name           string
		g              AugmentedCFG
		rhs            AugmentedCFG
		expectedEquals bool
	}{
		{
			name:           "Equal",
			g:              g0,
			rhs:            g0,
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			g:              g0,
			rhs:            g1,
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.g.Verify())
			assert.NoError(t, tc.rhs.Verify())
			assert.Equal(t, tc.expectedEquals, tc.g.Equals(tc.rhs))
		})
	}
}

func TestAugmentedCFG_CLOSURE(t *testing.T) {
	s := getTestItemSets()
	g := AugmentCFG(grammars[0])

	tests := []struct {
		name            string
		g               AugmentedCFG
		I               ItemSet
		expectedCLOSURE ItemSet
	}{
		{
			name: "OK",
			g:    g,
			I: NewItemSet(
				Item{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.g.CLOSURE(tc.I)
			assert.True(t, J.Equals(tc.expectedCLOSURE))
		})
	}
}

func TestAugmentedCFG_GOTO(t *testing.T) {
	s := getTestItemSets()
	g := AugmentCFG(grammars[0])

	tests := []struct {
		name         string
		g            AugmentedCFG
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name:         `GOTO(I₀,E)`,
			g:            g,
			I:            s[0],
			X:            grammar.NonTerminal("E"),
			expectedGOTO: s[1],
		},
		{
			name:         `GOTO(I₀,T)`,
			g:            g,
			I:            s[0],
			X:            grammar.NonTerminal("T"),
			expectedGOTO: s[2],
		},
		{
			name:         `GOTO(I₀,F)`,
			g:            g,
			I:            s[0],
			X:            grammar.NonTerminal("F"),
			expectedGOTO: s[3],
		},
		{
			name:         `GOTO(I₀,"(")`,
			g:            g,
			I:            s[0],
			X:            grammar.Terminal("("),
			expectedGOTO: s[4],
		},
		{
			name:         `GOTO(I₀,"id")`,
			g:            g,
			I:            s[0],
			X:            grammar.Terminal("id"),
			expectedGOTO: s[5],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.g.GOTO(tc.I, tc.X)
			assert.True(t, J.Equals(tc.expectedGOTO))
		})
	}
}

func TestAugmentedCFG_CanonicalLR0Collection(t *testing.T) {
	s := getTestItemSets()
	g := AugmentCFG(grammars[0])

	tests := []struct {
		name               string
		g                  AugmentedCFG
		expectedCollection ItemSetCollection
	}{
		{
			name:               "OK",
			g:                  g,
			expectedCollection: NewItemSetCollection(s...),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.g.CanonicalLR0Collection()
			assert.True(t, C.Equals(tc.expectedCollection))
		})
	}
}
