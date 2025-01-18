package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/sort"
)

func getTestItemSets() []ItemSet {
	I0 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
		// Non-Kernels
		Item{Production: &prods[0][1], Dot: 0}, // E → •E + T
		Item{Production: &prods[0][2], Dot: 0}, // E → •T
		Item{Production: &prods[0][3], Dot: 0}, // T → •T * F
		Item{Production: &prods[0][4], Dot: 0}, // T → •F
		Item{Production: &prods[0][5], Dot: 0}, // F → •( E )
		Item{Production: &prods[0][6], Dot: 0}, // F → •id
	)

	I1 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][0], Initial: true, Dot: 1}, // E′ → E•
		Item{Production: &prods[0][1], Dot: 1},                // E → E•+ T
	)

	I2 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][2], Dot: 1}, // E → T•
		Item{Production: &prods[0][3], Dot: 1}, // T → T•* F
	)

	I3 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][4], Dot: 1}, // T → F•
	)

	I4 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][5], Dot: 1}, // F → (•E )
		// Non-Kernels
		Item{Production: &prods[0][1], Dot: 0}, // E → •E + T
		Item{Production: &prods[0][2], Dot: 0}, // E → •T
		Item{Production: &prods[0][3], Dot: 0}, // T → •T * F
		Item{Production: &prods[0][4], Dot: 0}, // T → •F
		Item{Production: &prods[0][5], Dot: 0}, // F → •( E )
		Item{Production: &prods[0][6], Dot: 0}, // F → •id
	)

	I5 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][6], Dot: 1}, // F → id•
	)

	I6 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][1], Dot: 2}, // E → E +•T
		// Non-Kernels
		Item{Production: &prods[0][3], Dot: 0}, // T → •T * F
		Item{Production: &prods[0][4], Dot: 0}, // T → •F
		Item{Production: &prods[0][5], Dot: 0}, // F → •( E )
		Item{Production: &prods[0][6], Dot: 0}, // F → •id
	)

	I7 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][3], Dot: 2}, // T → T *•F
		// Non-Kernels
		Item{Production: &prods[0][5], Dot: 0}, // F → •( E )
		Item{Production: &prods[0][6], Dot: 0}, // F → •id
	)

	I8 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][1], Dot: 1}, // E → E• + T
		Item{Production: &prods[0][5], Dot: 2}, // F → ( E•)
	)

	I9 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][1], Dot: 3}, // E → E + T•
		Item{Production: &prods[0][3], Dot: 1}, // T → T•* F
	)

	I10 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][3], Dot: 3}, // T → T * F•
	)

	I11 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][5], Dot: 3}, // F → ( E )•
	)

	return []ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func TestNewItemSetCollection(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name                      string
		sets                      []ItemSet
		expectedItemSetCollection ItemSetCollection
	}{
		{
			name:                      "OK",
			sets:                      []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedItemSetCollection: NewItemSetCollection(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSetCollection(tc.sets...)
			assert.NotNil(t, s)
			assert.True(t, s.Equals(tc.expectedItemSetCollection))
		})
	}
}

func TestNewItemSet(t *testing.T) {
	tests := []struct {
		name            string
		items           []Item
		expectedItemSet ItemSet
	}{
		{
			name: "OK",
			items: []Item{
				{Production: &prods[0][0], Initial: true, Dot: 1}, // E′ → E•
				{Production: &prods[0][1], Dot: 1},                // E → E•+ T
			},
			expectedItemSet: NewItemSet(
				Item{Production: &prods[0][0], Initial: true, Dot: 1}, // E′ → E•
				Item{Production: &prods[0][1], Dot: 1},                // E → E•+ T
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSet(tc.items...)
			assert.NotNil(t, s)
			assert.True(t, s.Equals(tc.expectedItemSet))
		})
	}
}

func TestItem_String(t *testing.T) {
	tests := []struct {
		name           string
		i              Item
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: Item{
				Production: &grammar.Production{
					Head: "A",
					Body: grammar.E,
				},
				Dot: 0,
			},
			expectedString: `A → •`,
		},
		{
			name:           "Initial",
			i:              Item{Production: &prods[0][0], Initial: true, Dot: 0},
			expectedString: `E′ → •E`,
		},
		{
			name:           "DotAtLeft",
			i:              Item{Production: &prods[0][1], Dot: 0},
			expectedString: `E → •E "+" T`,
		},
		{
			name:           "DotInMiddle",
			i:              Item{Production: &prods[0][1], Dot: 2},
			expectedString: `E → E "+"•T`,
		},
		{
			name:           "DotAtRight",
			i:              Item{Production: &prods[0][1], Dot: 3},
			expectedString: `E → E "+" T•`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem_Equals(t *testing.T) {
	tests := []struct {
		name           string
		i              Item
		rhs            Item
		expectedEquals bool
	}{
		{
			name:           "Equal",
			i:              Item{Production: &prods[0][1], Dot: 1}, // E → E•+ T
			rhs:            Item{Production: &prods[0][1], Dot: 1}, // E → E•+ T
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			i:              Item{Production: &prods[0][1], Dot: 1}, // E → E•+ T
			rhs:            Item{Production: &prods[0][1], Dot: 2}, // E → E +•T
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.i.Equals(tc.rhs))
		})
	}
}

func TestItem_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                Item
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                Item{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                Item{Production: &prods[0][1], Dot: 2}, // E → E +•T
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                Item{Production: &prods[0][1], Dot: 0}, // E → •E + T
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel())
		})
	}
}

func TestItem_IsComplete(t *testing.T) {
	tests := []struct {
		name               string
		i                  Item
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  Item{Production: &prods[0][1], Dot: 3}, // E → E + T•
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  Item{Production: &prods[0][1], Dot: 2}, // E → E +•T
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestItem_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 Item
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 Item{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
			expectedDotSymbol: grammar.NonTerminal("E"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 Item{Production: &prods[0][1], Dot: 3}, // E → E + T•
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

func TestItem_NextItem(t *testing.T) {
	tests := []struct {
		name             string
		i                Item
		expectedNextItem Item
		expectedOK       bool
	}{
		{
			name:             "Initial",
			i:                Item{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
			expectedNextItem: Item{Production: &prods[0][0], Initial: true, Dot: 1}, // E′ → E•
			expectedOK:       true,
		},
		{
			name:             "Complete",
			i:                Item{Production: &prods[0][1], Dot: 3}, // E → E + T•
			expectedNextItem: Item{},
			expectedOK:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			item, ok := tc.i.NextItem()

			assert.Equal(t, tc.expectedNextItem, item)
			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestCmpItem(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name          string
		items         []Item
		expectedItems []Item
	}{
		{
			name:  "I₀",
			items: generic.Collect1(s[0].All()),
			expectedItems: []Item{
				// Kernels
				{Production: &prods[0][0], Initial: true, Dot: 0}, // E′ → •E
				// Non-Kernels
				{Production: &prods[0][1], Dot: 0}, // E → •E + T
				{Production: &prods[0][2], Dot: 0}, // E → •T
				{Production: &prods[0][5], Dot: 0}, // F → •( E )
				{Production: &prods[0][6], Dot: 0}, // F → •id
				{Production: &prods[0][3], Dot: 0}, // T → •T * F
				{Production: &prods[0][4], Dot: 0}, // T → •F
			},
		},
		{
			name:  "I₈",
			items: generic.Collect1(s[8].All()),
			expectedItems: []Item{
				{Production: &prods[0][1], Dot: 1}, // E → E• + T
				{Production: &prods[0][5], Dot: 2}, // F → ( E•)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.items, cmpItem)
			assert.Equal(t, tc.expectedItems, tc.items)
		})
	}
}

func TestCmpItemSet(t *testing.T) {
	s := getTestItemSets()

	tests := []struct {
		name         string
		sets         []ItemSet
		expectedSets []ItemSet
	}{
		{
			name:         "OK",
			sets:         []ItemSet{s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]},
			expectedSets: []ItemSet{s[0], s[1], s[8], s[2], s[4], s[5], s[9], s[3], s[6], s[7], s[11], s[10]},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.sets, cmpItemSet)
			assert.Equal(t, tc.expectedSets, tc.sets)
		})
	}
}
