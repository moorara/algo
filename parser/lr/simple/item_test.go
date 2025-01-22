package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/sort"
)

func getTestItemSets() []lr.ItemSet {
	I0 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
		// Non-Kernels
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 0}, // E → •E + T
		LR0Item{Production: &prods[0][2], Start: &starts[0], Dot: 0}, // E → •T
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 0}, // T → •T * F
		LR0Item{Production: &prods[0][4], Start: &starts[0], Dot: 0}, // T → •F
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
		LR0Item{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
	)

	I1 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 1}, // E′ → E•
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E•+ T
	)

	I2 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][2], Start: &starts[0], Dot: 1}, // E → T•
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 1}, // T → T•* F
	)

	I3 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][4], Start: &starts[0], Dot: 1}, // T → F•
	)

	I4 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 1}, // F → (•E )
		// Non-Kernels
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 0}, // E → •E + T
		LR0Item{Production: &prods[0][2], Start: &starts[0], Dot: 0}, // E → •T
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 0}, // T → •T * F
		LR0Item{Production: &prods[0][4], Start: &starts[0], Dot: 0}, // T → •F
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
		LR0Item{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
	)

	I5 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][6], Start: &starts[0], Dot: 1}, // F → id•
	)

	I6 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 2}, // E → E +•T
		// Non-Kernels
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 0}, // T → •T * F
		LR0Item{Production: &prods[0][4], Start: &starts[0], Dot: 0}, // T → •F
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
		LR0Item{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
	)

	I7 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 2}, // T → T *•F
		// Non-Kernels
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
		LR0Item{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
	)

	I8 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E• + T
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 2}, // F → ( E•)
	)

	I9 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 1}, // T → T•* F
	)

	I10 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][3], Start: &starts[0], Dot: 3}, // T → T * F•
	)

	I11 := lr.NewItemSet(
		// Kernels
		LR0Item{Production: &prods[0][5], Start: &starts[0], Dot: 3}, // F → ( E )•
	)

	return []lr.ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func TestLR0Item_String(t *testing.T) {
	tests := []struct {
		name           string
		i              LR0Item
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: LR0Item{
				Production: &grammar.Production{Head: "E", Body: grammar.E},
				Start:      &starts[0],
				Dot:        0,
			},
			expectedString: `E → •`,
		},
		{
			name:           "Initial",
			i:              LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0},
			expectedString: `E′ → •E`,
		},
		{
			name:           "DotAtLeft",
			i:              LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 0},
			expectedString: `E → •E "+" T`,
		},
		{
			name:           "DotInMiddle",
			i:              LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 2},
			expectedString: `E → E "+"•T`,
		},
		{
			name:           "DotAtRight",
			i:              LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 3},
			expectedString: `E → E "+" T•`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestLR0Item_Equals(t *testing.T) {
	tests := []struct {
		name           string
		i              LR0Item
		rhs            LR0Item
		expectedEquals bool
	}{
		{
			name:           "Equal",
			i:              LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E•+ T
			rhs:            LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E•+ T
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			i:              LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E•+ T
			rhs:            LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 2}, // E → E +•T
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.i.Equals(tc.rhs))
		})
	}
}

func TestLR0Item_Compare(t *testing.T) {
	tests := []struct {
		name          string
		items         []LR0Item
		expectedItems []LR0Item
	}{
		{
			name: "I₀",
			items: []LR0Item{
				{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
				{Production: &prods[0][1], Start: &starts[0], Dot: 0}, // E → •E + T
				{Production: &prods[0][2], Start: &starts[0], Dot: 0}, // E → •T
				{Production: &prods[0][3], Start: &starts[0], Dot: 0}, // T → •T * F
				{Production: &prods[0][4], Start: &starts[0], Dot: 0}, // T → •F
				{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
				{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
			},
			expectedItems: []LR0Item{
				{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
				{Production: &prods[0][1], Start: &starts[0], Dot: 0}, // E → •E + T
				{Production: &prods[0][2], Start: &starts[0], Dot: 0}, // E → •T
				{Production: &prods[0][5], Start: &starts[0], Dot: 0}, // F → •( E )
				{Production: &prods[0][6], Start: &starts[0], Dot: 0}, // F → •id
				{Production: &prods[0][3], Start: &starts[0], Dot: 0}, // T → •T * F
				{Production: &prods[0][4], Start: &starts[0], Dot: 0}, // T → •F
			},
		},
		{
			name: "I₈",
			items: []LR0Item{
				{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E• + T
				{Production: &prods[0][5], Start: &starts[0], Dot: 2}, // F → ( E•)
			},
			expectedItems: []LR0Item{
				{Production: &prods[0][5], Start: &starts[0], Dot: 2}, // F → ( E•)
				{Production: &prods[0][1], Start: &starts[0], Dot: 1}, // E → E• + T
			},
		},
		{
			name: "I₉",
			items: []LR0Item{
				{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
				{Production: &prods[0][3], Start: &starts[0], Dot: 1}, // T → T•* F
			},
			expectedItems: []LR0Item{
				{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
				{Production: &prods[0][3], Start: &starts[0], Dot: 1}, // T → T•* F
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.items, func(lhs, rhs LR0Item) int {
				return lhs.Compare(rhs)
			})

			assert.Equal(t, tc.expectedItems, tc.items)
		})
	}
}

func TestLR0Item_IsInitial(t *testing.T) {
	tests := []struct {
		name             string
		i                LR0Item
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 1}, // E′ → E•
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsInitial())
		})
	}
}

func TestLR0Item_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                LR0Item
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 2}, // E → E +•T
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 0}, // E → •E + T
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel())
		})
	}
}

func TestLR0Item_IsComplete(t *testing.T) {
	tests := []struct {
		name               string
		i                  LR0Item
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 2}, // E → E +•T
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestLR0Item_IsFinal(t *testing.T) {
	tests := []struct {
		name            string
		i               LR0Item
		expectedIsFinal bool
	}{
		{
			name:            "Final",
			i:               LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 1}, // E′ → E•
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			expectedIsFinal: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsFinal, tc.i.IsFinal())
		})
	}
}

func TestLR0Item_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 LR0Item
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			expectedDotSymbol: grammar.NonTerminal("E"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
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

func TestLR0Item_Next(t *testing.T) {
	tests := []struct {
		name         string
		i            LR0Item
		expectedNext lr.Item
		expectedOK   bool
	}{
		{
			name:         "Initial",
			i:            LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 0}, // E′ → •E
			expectedNext: LR0Item{Production: &prods[0][0], Start: &starts[0], Dot: 1}, // E′ → E•
			expectedOK:   true,
		},
		{
			name:         "Complete",
			i:            LR0Item{Production: &prods[0][1], Start: &starts[0], Dot: 3}, // E → E + T•
			expectedNext: LR0Item{},
			expectedOK:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item, ok := tc.i.Next()

			assert.Equal(t, tc.expectedNext, item)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}
