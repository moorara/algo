package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

func getTestItemSets() []set.Set[Item] {
	I0 := NewItemSet(
		// Kernels
		Item{Production: &prods[0][0], Dot: 0}, // E′ → •E
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
		Item{Production: &prods[0][0], Dot: 1}, // E′ → E•
		Item{Production: &prods[0][1], Dot: 1}, // E → E•+ T
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

	return []set.Set[Item]{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func TestNewItemSet(t *testing.T) {
	tests := []struct {
		name            string
		items           []Item
		expectedItemSet set.Set[Item]
	}{
		{
			name: "OK",
			items: []Item{
				{Production: &prods[0][0], Dot: 1}, // E′ → E•
				{Production: &prods[0][1], Dot: 1}, // E → E•+ T
			},
			expectedItemSet: NewItemSet(
				Item{Production: &prods[0][0], Dot: 1}, // E′ → E•
				Item{Production: &prods[0][1], Dot: 1}, // E → E•+ T
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
			name: "DotAtLeft",
			i: Item{
				Production: &prods[0][1],
				Dot:        0,
			},
			expectedString: `E → •E "+" T`,
		},
		{
			name: "DotInMiddle",
			i: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
			expectedString: `E → E "+"•T`,
		},
		{
			name: "DotAtRight",
			i: Item{
				Production: &prods[0][1],
				Dot:        3,
			},
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
		j              Item
		expectedEquals bool
	}{
		{
			name: "Equal",
			i: Item{
				Production: &prods[0][1],
				Dot:        1,
			},
			j: Item{
				Production: &prods[0][1],
				Dot:        1,
			},
			expectedEquals: true,
		},
		{
			name: "NotEqual",
			i: Item{
				Production: &prods[0][1],
				Dot:        1,
			},
			j: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.i.Equals(tc.j))
		})
	}
}

func TestItem_IsKernel(t *testing.T) {
	tests := []struct {
		name             string
		i                Item
		init             Item
		expectedIsKernel bool
	}{
		{
			name: "InitialItem",
			i: Item{
				Production: &prods[0][0],
				Dot:        0,
			},
			init: Item{
				Production: &prods[0][0],
				Dot:        0,
			},
			expectedIsKernel: true,
		},
		{
			name: "Kernel",
			i: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
			init: Item{
				Production: &prods[0][0],
				Dot:        0,
			},
			expectedIsKernel: true,
		},
		{
			name: "NonKernel",
			i: Item{
				Production: &prods[0][1],
				Dot:        0,
			},
			init: Item{
				Production: &prods[0][0],
				Dot:        0,
			},
			expectedIsKernel: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsKernel, tc.i.IsKernel(tc.init))
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
			name: "Complete",
			i: Item{
				Production: &prods[0][1],
				Dot:        3,
			},
			expectedIsComplete: true,
		},
		{
			name: "NotComplete",
			i: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
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
			name: "OK",
			i: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
			expectedDotSymbol: grammar.NonTerminal("T"),
			expectedOK:        true,
		},
		{
			name: "Complete",
			i: Item{
				Production: &prods[0][1],
				Dot:        3,
			},
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
			name: "OK",
			i: Item{
				Production: &prods[0][1],
				Dot:        2,
			},
			expectedNextItem: Item{
				Production: &prods[0][1],
				Dot:        3,
			},
			expectedOK: true,
		},
		{
			name: "Complete",
			i: Item{
				Production: &prods[0][1],
				Dot:        3,
			},
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
