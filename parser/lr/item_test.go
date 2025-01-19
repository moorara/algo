package lr

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/sort"
	"github.com/stretchr/testify/assert"
)

var start = grammar.NonTerminal("E′")

var prods = []grammar.Production{
	{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
	{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
	{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
	{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
	{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
	{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
	{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
}

func getTestItemSets() []ItemSet {
	I0 := NewItemSet(
		// Kernels
		Item0{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
		// Non-Kernels
		Item0{Production: &prods[1], Start: &start, Dot: 0}, // E → •E + T
		Item0{Production: &prods[2], Start: &start, Dot: 0}, // E → •T
		Item0{Production: &prods[3], Start: &start, Dot: 0}, // T → •T * F
		Item0{Production: &prods[4], Start: &start, Dot: 0}, // T → •F
		Item0{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
		Item0{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
	)

	I1 := NewItemSet(
		// Kernels
		Item0{Production: &prods[0], Start: &start, Dot: 1}, // E′ → E•
		Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E•+ T
	)

	I2 := NewItemSet(
		// Kernels
		Item0{Production: &prods[2], Start: &start, Dot: 1}, // E → T•
		Item0{Production: &prods[3], Start: &start, Dot: 1}, // T → T•* F
	)

	I3 := NewItemSet(
		// Kernels
		Item0{Production: &prods[4], Start: &start, Dot: 1}, // T → F•
	)

	I4 := NewItemSet(
		// Kernels
		Item0{Production: &prods[5], Start: &start, Dot: 1}, // F → (•E )
		// Non-Kernels
		Item0{Production: &prods[1], Start: &start, Dot: 0}, // E → •E + T
		Item0{Production: &prods[2], Start: &start, Dot: 0}, // E → •T
		Item0{Production: &prods[3], Start: &start, Dot: 0}, // T → •T * F
		Item0{Production: &prods[4], Start: &start, Dot: 0}, // T → •F
		Item0{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
		Item0{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
	)

	I5 := NewItemSet(
		// Kernels
		Item0{Production: &prods[6], Start: &start, Dot: 1}, // F → id•
	)

	I6 := NewItemSet(
		// Kernels
		Item0{Production: &prods[1], Start: &start, Dot: 2}, // E → E +•T
		// Non-Kernels
		Item0{Production: &prods[3], Start: &start, Dot: 0}, // T → •T * F
		Item0{Production: &prods[4], Start: &start, Dot: 0}, // T → •F
		Item0{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
		Item0{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
	)

	I7 := NewItemSet(
		// Kernels
		Item0{Production: &prods[3], Start: &start, Dot: 2}, // T → T *•F
		// Non-Kernels
		Item0{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
		Item0{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
	)

	I8 := NewItemSet(
		// Kernels
		Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E• + T
		Item0{Production: &prods[5], Start: &start, Dot: 2}, // F → ( E•)
	)

	I9 := NewItemSet(
		// Kernels
		Item0{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
		Item0{Production: &prods[3], Start: &start, Dot: 1}, // T → T•* F
	)

	I10 := NewItemSet(
		// Kernels
		Item0{Production: &prods[3], Start: &start, Dot: 3}, // T → T * F•
	)

	I11 := NewItemSet(
		// Kernels
		Item0{Production: &prods[5], Start: &start, Dot: 3}, // F → ( E )•
	)

	return []ItemSet{I0, I1, I2, I3, I4, I5, I6, I7, I8, I9, I10, I11}
}

func TestNewItemSetCollection(t *testing.T) {
	s := getTestItemSets()

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
				Item0{Production: &prods[0], Start: &start, Dot: 1}, // E′ → E•
				Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E•+ T
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
	s := getTestItemSets()

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
		i              Item0
		expectedString string
	}{
		{
			name: "EmptyProduction",
			i: Item0{
				Production: &grammar.Production{Head: "E", Body: grammar.E},
				Start:      &start,
				Dot:        0,
			},
			expectedString: `E → •`,
		},
		{
			name:           "Initial",
			i:              Item0{Production: &prods[0], Start: &start, Dot: 0},
			expectedString: `E′ → •E`,
		},
		{
			name:           "DotAtLeft",
			i:              Item0{Production: &prods[1], Start: &start, Dot: 0},
			expectedString: `E → •E "+" T`,
		},
		{
			name:           "DotInMiddle",
			i:              Item0{Production: &prods[1], Start: &start, Dot: 2},
			expectedString: `E → E "+"•T`,
		},
		{
			name:           "DotAtRight",
			i:              Item0{Production: &prods[1], Start: &start, Dot: 3},
			expectedString: `E → E "+" T•`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem0_Equals(t *testing.T) {
	tests := []struct {
		name           string
		i              Item0
		rhs            Item0
		expectedEquals bool
	}{
		{
			name:           "Equal",
			i:              Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E•+ T
			rhs:            Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E•+ T
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			i:              Item0{Production: &prods[1], Start: &start, Dot: 1}, // E → E•+ T
			rhs:            Item0{Production: &prods[1], Start: &start, Dot: 2}, // E → E +•T
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.i.Equals(tc.rhs))
		})
	}
}

func TestItem0_Compare(t *testing.T) {
	tests := []struct {
		name          string
		items         []Item0
		expectedItems []Item0
	}{
		{
			name: "I₀",
			items: []Item0{
				{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
				{Production: &prods[1], Start: &start, Dot: 0}, // E → •E + T
				{Production: &prods[2], Start: &start, Dot: 0}, // E → •T
				{Production: &prods[3], Start: &start, Dot: 0}, // T → •T * F
				{Production: &prods[4], Start: &start, Dot: 0}, // T → •F
				{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
				{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
			},
			expectedItems: []Item0{
				{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
				{Production: &prods[1], Start: &start, Dot: 0}, // E → •E + T
				{Production: &prods[2], Start: &start, Dot: 0}, // E → •T
				{Production: &prods[5], Start: &start, Dot: 0}, // F → •( E )
				{Production: &prods[6], Start: &start, Dot: 0}, // F → •id
				{Production: &prods[3], Start: &start, Dot: 0}, // T → •T * F
				{Production: &prods[4], Start: &start, Dot: 0}, // T → •F
			},
		},
		{
			name: "I₈",
			items: []Item0{
				{Production: &prods[1], Start: &start, Dot: 1}, // E → E• + T
				{Production: &prods[5], Start: &start, Dot: 2}, // F → ( E•)
			},
			expectedItems: []Item0{
				{Production: &prods[5], Start: &start, Dot: 2}, // F → ( E•)
				{Production: &prods[1], Start: &start, Dot: 1}, // E → E• + T
			},
		},
		{
			name: "I₉",
			items: []Item0{
				{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
				{Production: &prods[3], Start: &start, Dot: 1}, // T → T•* F
			},
			expectedItems: []Item0{
				{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
				{Production: &prods[3], Start: &start, Dot: 1}, // T → T•* F
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.items, func(lhs, rhs Item0) int {
				return lhs.Compare(rhs)
			})

			assert.Equal(t, tc.expectedItems, tc.items)
		})
	}
}

func TestItem0_IsInitial(t *testing.T) {
	tests := []struct {
		name             string
		i                Item0
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                Item0{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                Item0{Production: &prods[0], Start: &start, Dot: 1}, // E′ → E•
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
		i                Item0
		expectedIsKernel bool
	}{
		{
			name:             "Initial",
			i:                Item0{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                Item0{Production: &prods[1], Start: &start, Dot: 2}, // E → E +•T
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                Item0{Production: &prods[1], Start: &start, Dot: 0}, // E → •E + T
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
		i                  Item0
		expectedIsComplete bool
	}{
		{
			name:               "Complete",
			i:                  Item0{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  Item0{Production: &prods[1], Start: &start, Dot: 2}, // E → E +•T
			expectedIsComplete: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsComplete, tc.i.IsComplete())
		})
	}
}

func TestItem0_DotSymbol(t *testing.T) {
	tests := []struct {
		name              string
		i                 Item0
		expectedDotSymbol grammar.Symbol
		expectedOK        bool
	}{
		{
			name:              "Initial",
			i:                 Item0{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
			expectedDotSymbol: grammar.NonTerminal("E"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 Item0{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
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

func TestItem0_Next(t *testing.T) {
	tests := []struct {
		name         string
		i            Item0
		expectedNext Item
		expectedOK   bool
	}{
		{
			name:         "Initial",
			i:            Item0{Production: &prods[0], Start: &start, Dot: 0}, // E′ → •E
			expectedNext: Item0{Production: &prods[0], Start: &start, Dot: 1}, // E′ → E•
			expectedOK:   true,
		},
		{
			name:         "Complete",
			i:            Item0{Production: &prods[1], Start: &start, Dot: 3}, // E → E + T•
			expectedNext: Item0{},
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
