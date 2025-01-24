package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/sort"
)

func getTestLR0ItemSets() []ItemSet {
	I0 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
		// Non-Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	)

	I1 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
	)

	I2 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 1}, // E → T•
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
	)

	I3 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 1}, // T → F•
	)

	I4 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E )
		// Non-Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
		&Item0{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	)

	I5 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 1}, // F → id•
	)

	I6 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
		// Non-Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
		&Item0{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	)

	I7 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 2}, // T → T *•F
		// Non-Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
		&Item0{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
	)

	I8 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E• + T
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
	)

	I9 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
	)

	I10 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][3], Start: starts[2], Dot: 3}, // T → T * F•
	)

	I11 := NewItemSet(
		// Kernels
		&Item0{Production: prods[2][5], Start: starts[2], Dot: 3}, // F → ( E )•
	)

	return []ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func getTestLR1ItemSets() []ItemSet {
	I0 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
		// Non-Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},     // S → •CC, $
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	)

	I1 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
	)

	I2 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	)

	I3 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → c•C, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → c•C, d
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •cC, d
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
	)

	I4 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("c")}, // C → d•, c
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Terminal("d")}, // C → d•, d
	)

	I5 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
	)

	I6 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → c•C, $
		// Non-Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •cC, $
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // C → •d, $
	)

	I7 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][3], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // C → d•, $
	)

	I8 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("c")}, // C → cC•, c
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Terminal("d")}, // C → cC•, d
	)

	I9 := NewItemSet(
		// Kernels
		&Item1{Production: prods[0][2], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // C → cC•, $
	)

	return []ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9}
}

func TestNewItemSetCollection(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name string
		sets []ItemSet
	}{
		{
			name: "OK",
			sets: []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSetCollection(tc.sets...)

			assert.NotNil(t, s)
			for _, expectedItemSet := range tc.sets {
				assert.True(t, s.Contains(expectedItemSet))
			}
		})
	}
}

func TestNewItemSet(t *testing.T) {
	tests := []struct {
		name  string
		items []Item
	}{
		{
			name: "OK",
			items: []Item{
				&Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
				&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSet(tc.items...)

			assert.NotNil(t, s)
			for _, expectedItem := range tc.items {
				assert.True(t, s.Contains(expectedItem))
			}
		})
	}
}

func TestCmpItemSet(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name         string
		sets         []ItemSet
		expectedSets []ItemSet
	}{
		{
			name:         "OK",
			sets:         []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedSets: []ItemSet{s[0], s[1], s[9], s[11], s[10], s[6], s[8], s[7], s[2], s[4], s[5], s[3]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.sets, cmpItemSet)
			assert.Equal(t, tc.expectedSets, tc.sets)
		})
	}
}

func TestItem0_String(t *testing.T) {
	tests := []struct {
		name           string
		i              *Item0
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: &Item0{
				Production: &grammar.Production{Head: "E", Body: grammar.E},
				Start:      starts[2],
				Dot:        0,
			},
			expectedString: `E → •`,
		},
		{
			name:           "Initial",
			i:              &Item0{Production: prods[2][0], Start: starts[2], Dot: 0},
			expectedString: `E′ → •E`,
		},
		{
			name:           "DotAtLeft",
			i:              &Item0{Production: prods[2][1], Start: starts[2], Dot: 0},
			expectedString: `E → •E "+" T`,
		},
		{
			name:           "DotInMiddle",
			i:              &Item0{Production: prods[2][1], Start: starts[2], Dot: 2},
			expectedString: `E → E "+"•T`,
		},
		{
			name:           "DotAtRight",
			i:              &Item0{Production: prods[2][1], Start: starts[2], Dot: 3},
			expectedString: `E → E "+" T•`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem0_Equal(t *testing.T) {
	tests := []struct {
		name          string
		i             *Item0
		rhs           *Item0
		expectedEqual bool
	}{
		{
			name:          "Equal",
			i:             &Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
			rhs:           &Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			i:             &Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
			rhs:           &Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.i.Equal(tc.rhs))
		})
	}
}

func TestItem0_Compare(t *testing.T) {
	tests := []struct {
		name          string
		items         []*Item0
		expectedItems []*Item0
	}{
		{
			name: "I₀",
			items: []*Item0{
				{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
				{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
				{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
				{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
				{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
				{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
				{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
			},
			expectedItems: []*Item0{
				{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
				{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
				{Production: prods[2][2], Start: starts[2], Dot: 0}, // E → •T
				{Production: prods[2][5], Start: starts[2], Dot: 0}, // F → •( E )
				{Production: prods[2][6], Start: starts[2], Dot: 0}, // F → •id
				{Production: prods[2][3], Start: starts[2], Dot: 0}, // T → •T * F
				{Production: prods[2][4], Start: starts[2], Dot: 0}, // T → •F
			},
		},
		{
			name: "I₈",
			items: []*Item0{
				{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E• + T
				{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
			},
			expectedItems: []*Item0{
				{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
				{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E• + T
			},
		},
		{
			name: "I₉",
			items: []*Item0{
				{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
				{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
			},
			expectedItems: []*Item0{
				{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
				{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.items, func(lhs, rhs *Item0) int {
				return lhs.Compare(rhs)
			})

			assert.Equal(t, tc.expectedItems, tc.items)
		})
	}
}

func TestItem0_IsInitial(t *testing.T) {
	tests := []struct {
		name             string
		i                *Item0
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                &Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsInitial())
		})
	}
}

func TestItem0_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                *Item0
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                &Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                &Item0{Production: prods[2][1], Start: starts[2], Dot: 0}, // E → •E + T
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel())
		})
	}
}

func TestItem0_IsComplete(t *testing.T) {
	tests := []struct {
		name               string
		i                  *Item0
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  &Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  &Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestItem0_IsFinal(t *testing.T) {
	tests := []struct {
		name            string
		i               *Item0
		expectedIsFinal bool
	}{
		{
			name:            "Final",
			i:               &Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               &Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			expectedIsFinal: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsFinal, tc.i.IsFinal())
		})
	}
}

func TestItem0_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 *Item0
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 &Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			expectedDotSymbol: grammar.NonTerminal("E"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 &Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
			expectedDotSymbol: nil,
			expectedOK:        false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			X, ok := tc.i.DotSymbol()

			if tc.expectedOK {
				assert.True(t, X.Equal(tc.expectedDotSymbol))
				assert.True(t, ok)
			} else {
				assert.Nil(t, X)
				assert.False(t, ok)
			}
		})
	}
}

func TestItem0_Next(t *testing.T) {
	tests := []struct {
		name         string
		i            *Item0
		expectedOK   bool
		expectedNext Item
	}{
		{
			name:         "Initial",
			i:            &Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			expectedOK:   true,
			expectedNext: &Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
		},
		{
			name:         "Complete",
			i:            &Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
			expectedOK:   false,
			expectedNext: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item, ok := tc.i.Next()

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, item.Equal(tc.expectedNext))
			} else {
				assert.False(t, ok)
				assert.Nil(t, item)
			}
		})
	}
}

func TestItem0_Item1(t *testing.T) {
	tests := []struct {
		name          string
		i             *Item0
		lookahead     grammar.Terminal
		expectedItem1 *Item1
	}{
		{
			name:          "OK",
			i:             &Item0{Production: prods[0][1], Start: starts[0], Dot: 1}, // S → C•C
			lookahead:     grammar.Endmarker,
			expectedItem1: &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item0 := tc.i.Item1(tc.lookahead)

			if tc.expectedItem1 == nil {
				assert.Nil(t, item0)
			} else {
				assert.True(t, item0.Equal(tc.expectedItem1))
			}
		})
	}
}

func TestItem1_String(t *testing.T) {
	tests := []struct {
		name           string
		i              *Item1
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: &Item1{
				Production: &grammar.Production{Head: "S", Body: grammar.E},
				Start:      starts[0],
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
			expectedString: `S → •, $`,
		},
		{
			name:           "Initial",
			i:              &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S′ → •S, $`,
		},
		{
			name:           "DotAtLeft",
			i:              &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S → •C C, $`,
		},
		{
			name:           "DotInMiddle",
			i:              &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker},
			expectedString: `S → C•C, $`,
		},
		{
			name:           "DotAtRight",
			i:              &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker},
			expectedString: `S → C C•, $`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem1_Equal(t *testing.T) {
	tests := []struct {
		name          string
		i             *Item1
		rhs           *Item1
		expectedEqual bool
	}{
		{
			name:          "Equal",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:           &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:           &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.i.Equal(tc.rhs))
		})
	}
}

func TestItem1_Compare(t *testing.T) {
	tests := []struct {
		name            string
		i               *Item1
		rhs             *Item1
		expectedCompare int
	}{
		{
			name:            "FirstInitial",
			i:               &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			rhs:             &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondInitial",
			i:               &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedCompare: 1,
		},
		{
			name:            "FirstKernel",
			i:               &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedCompare: -1,
		},
		{
			name:            "SecondKernel",
			i:               &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			rhs:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: 1,
		},
		{
			name:            "FirstHead",
			i:               &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondHead",
			i:               &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstDot",
			i:               &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: -1,
		},
		{
			name:            "SecondDot",
			i:               &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstProduction",
			i:               &Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			rhs:             &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			expectedCompare: -1,
		},
		{
			name:            "SecondProduction",
			i:               &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &Item1{Production: prods[0][2], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			expectedCompare: 1,
		},
		{
			name:            "FirstLookahead",
			i:               &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			expectedCompare: -1,
		},
		{
			name:            "SecondLookahead",
			i:               &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			rhs:             &Item1{Production: prods[0][3], Start: starts[0], Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
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

func TestItem1_IsInitial(t *testing.T) {
	tests := []struct {
		name             string
		i                *Item1
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsInitial())
		})
	}
}

func TestItem1_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                *Item1
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // E → •CC, $
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel())
		})
	}
}

func TestItem1_IsComplete(t *testing.T) {
	tests := []struct {
		name               string
		i                  *Item1
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestItem1_IsFinal(t *testing.T) {
	tests := []struct {
		name            string
		i               *Item1
		expectedIsFinal bool
	}{
		{
			name:            "Final",
			i:               &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsFinal: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsFinal, tc.i.IsFinal())
		})
	}
}

func TestItem1_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 *Item1
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedDotSymbol: grammar.NonTerminal("S"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedDotSymbol: nil,
			expectedOK:        false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			X, ok := tc.i.DotSymbol()

			if tc.expectedOK {
				assert.True(t, X.Equal(tc.expectedDotSymbol))
				assert.True(t, ok)
			} else {
				assert.Nil(t, X)
				assert.False(t, ok)
			}
		})
	}
}

func TestItem1_Next(t *testing.T) {
	tests := []struct {
		name         string
		i            *Item1
		expectedNext Item
		expectedOK   bool
	}{
		{
			name:         "Initial",
			i:            &Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedOK:   true,
			expectedNext: &Item1{Production: prods[0][0], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{
			name:         "Complete",
			i:            &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•
			expectedOK:   false,
			expectedNext: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item, ok := tc.i.Next()

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, item.Equal(tc.expectedNext))
			} else {
				assert.False(t, ok)
				assert.Nil(t, item)
			}
		})
	}
}

func TestItem1_GetPrefix(t *testing.T) {
	tests := []struct {
		name          string
		i             *Item1
		expectedAlpha grammar.String[grammar.Symbol]
	}{
		{
			name:          "DotAtLeft",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedAlpha: grammar.E,
		},
		{
			name:          "DotInMiddle",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedAlpha: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:          "DotAtRight",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedAlpha: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			alpha := tc.i.GetPrefix()

			assert.True(t, alpha.Equal(tc.expectedAlpha))
		})
	}
}

func TestItem1_GetSuffix(t *testing.T) {
	tests := []struct {
		name         string
		i            *Item1
		expectedBeta grammar.String[grammar.Symbol]
	}{
		{
			name:         "DotAtLeft",
			i:            &Item1{Production: prods[0][1], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedBeta: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")},
		},
		{
			name:         "DotInMiddle",
			i:            &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedBeta: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:         "DotAtRight",
			i:            &Item1{Production: prods[0][1], Start: starts[0], Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedBeta: grammar.E,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			beta := tc.i.GetSuffix()

			assert.True(t, beta.Equal(tc.expectedBeta))
		})
	}
}

func TestItem1_Item0(t *testing.T) {
	tests := []struct {
		name          string
		i             *Item1
		expectedItem0 *Item0
	}{
		{
			name:          "OK",
			i:             &Item1{Production: prods[0][1], Start: starts[0], Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedItem0: &Item0{Production: prods[0][1], Start: starts[0], Dot: 1},                               // S → C•C
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item0 := tc.i.Item0()

			if tc.expectedItem0 == nil {
				assert.Nil(t, item0)
			} else {
				assert.True(t, item0.Equal(tc.expectedItem0))
			}
		})
	}
}
