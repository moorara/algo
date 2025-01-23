package canonical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

func getTestItemSets() []lr.ItemSet {
	I0 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		// Non-Kernels
		&LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	)

	I1 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
	)

	I2 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
		// Non-Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	)

	I3 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
		// Non-Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	)

	I4 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
	)

	I5 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
	)

	I6 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
		// Non-Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	)

	I7 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
	)

	I8 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
	)

	I9 := lr.NewItemSet(
		// Kernels
		&LR1Item{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
	)

	return []lr.ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9}
}

func TestLR1Item_String(t *testing.T) {
	tests := []struct {
		name           string
		i              *LR1Item
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: &LR1Item{
				Production: &grammar.Production{Head: "S", Body: grammar.E},
				Start:      starts[0],
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
			expectedString: `S → •, $`,
		},
		{
			name:           "Initial",
			i:              &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S′ → •S, $`,
		},
		{
			name:           "DotAtLeft",
			i:              &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S → •C C, $`,
		},
		{
			name:           "DotInMiddle",
			i:              &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker},
			expectedString: `S → C•C, $`,
		},
		{
			name:           "DotAtRight",
			i:              &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker},
			expectedString: `S → C C•, $`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestLR1Item_Equals(t *testing.T) {
	tests := []struct {
		name           string
		i              *LR1Item
		rhs            *LR1Item
		expectedEquals bool
	}{
		{
			name:           "Equal",
			i:              &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			i:              &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.i.Equals(tc.rhs))
		})
	}
}

func TestLR1Item_Compare(t *testing.T) {
	tests := []struct {
		name            string
		i               *LR1Item
		rhs             *LR1Item
		expectedCompare int
	}{
		{
			name:            "FirstInitial",
			i:               &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			rhs:             &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondInitial",
			i:               &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedCompare: 1,
		},
		{
			name:            "FirstKernel",
			i:               &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedCompare: -1,
		},
		{
			name:            "SecondKernel",
			i:               &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			rhs:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: 1,
		},
		{
			name:            "FirstHead",
			i:               &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondHead",
			i:               &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstDot",
			i:               &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: -1,
		},
		{
			name:            "SecondDot",
			i:               &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstProduction",
			i:               &LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			rhs:             &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			expectedCompare: -1,
		},
		{
			name:            "SecondProduction",
			i:               &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &LR1Item{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			expectedCompare: 1,
		},
		{
			name:            "FirstLookahead",
			i:               &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			expectedCompare: -1,
		},
		{
			name:            "SecondLookahead",
			i:               &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			rhs:             &LR1Item{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			expectedCompare: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmp := tc.i.Compare(tc.rhs)

			assert.Equal(t, tc.expectedCompare, cmp)
		})
	}
}

func TestLR1Item_IsInitial(t *testing.T) {
	tests := []struct {
		name             string
		i                *LR1Item
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsInitial())
		})
	}
}

func TestLR1Item_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                *LR1Item
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // E → •CC, $
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel())
		})
	}
}

func TestLR1Item_IsComplete(t *testing.T) {
	tests := []struct {
		name               string
		i                  *LR1Item
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestLR1Item_IsFinal(t *testing.T) {
	tests := []struct {
		name            string
		i               *LR1Item
		expectedIsFinal bool
	}{
		{
			name:            "Final",
			i:               &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsFinal: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsFinal, tc.i.IsFinal())
		})
	}
}

func TestLR1Item_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 *LR1Item
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedDotSymbol: grammar.NonTerminal("S"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedDotSymbol: nil,
			expectedOK:        false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			X, ok := tc.i.DotSymbol()

			if tc.expectedOK {
				assert.True(t, X.Equals(tc.expectedDotSymbol))
				assert.True(t, ok)
			} else {
				assert.Nil(t, X)
				assert.False(t, ok)
			}
		})
	}
}

func TestLR1Item_Next(t *testing.T) {
	tests := []struct {
		name         string
		i            *LR1Item
		expectedNext lr.Item
		expectedOK   bool
	}{
		{
			name:         "Initial",
			i:            &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedOK:   true,
			expectedNext: &LR1Item{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{
			name:         "Complete",
			i:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•
			expectedOK:   false,
			expectedNext: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item, ok := tc.i.Next()

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, item.Equals(tc.expectedNext))
			} else {
				assert.False(t, ok)
				assert.Nil(t, item)
			}
		})
	}
}

func TestLR1Item_GetαPrefix(t *testing.T) {
	tests := []struct {
		name          string
		i             *LR1Item
		expectedAlpha grammar.String[grammar.Symbol]
	}{
		{
			name:          "DotAtLeft",
			i:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedAlpha: grammar.E,
		},
		{
			name:          "DotInMiddle",
			i:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedAlpha: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:          "DotAtRight",
			i:             &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedAlpha: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			alpha := tc.i.GetαPrefix()

			assert.True(t, alpha.Equals(tc.expectedAlpha))
		})
	}
}

func TestLR1Item_GetβSuffix(t *testing.T) {
	tests := []struct {
		name         string
		i            *LR1Item
		expectedBeta grammar.String[grammar.Symbol]
	}{
		{
			name:         "DotAtLeft",
			i:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedBeta: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:         "DotInMiddle",
			i:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedBeta: grammar.E,
		},
		{
			name:         "DotAtRight",
			i:            &LR1Item{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedBeta: grammar.E,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			beta := tc.i.GetβSuffix()

			assert.True(t, beta.Equals(tc.expectedBeta))
		})
	}
}
