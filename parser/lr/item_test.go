package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/sort"
)

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
				Start:      "E′",
				Dot:        0,
			},
			expectedString: `E → •`,
		},
		{
			name:           "Initial",
			i:              &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0},
			expectedString: `E′ → •E`,
		},
		{
			name:           "DotAtLeft",
			i:              &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0},
			expectedString: `E → •E "+" T`,
		},
		{
			name:           "DotInMiddle",
			i:              &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2},
			expectedString: `E → E "+"•T`,
		},
		{
			name:           "DotAtRight",
			i:              &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3},
			expectedString: `E → E "+" T•`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem0_Hash(t *testing.T) {
	tests := []struct {
		name         string
		i            *Item0
		expectedHash uint64
	}{
		{
			name: "EmptyProduction",
			i: &Item0{
				Production: &grammar.Production{Head: "E", Body: grammar.E},
				Start:      "E′",
				Dot:        0,
			},
			expectedHash: 0x53601cabdae2cd11,
		},
		{
			name:         "Initial",
			i:            &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0},
			expectedHash: 0x1220d07375b21726,
		},
		{
			name:         "DotAtLeft",
			i:            &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0},
			expectedHash: 0xdf028a641306cfaf,
		},
		{
			name:         "DotInMiddle",
			i:            &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2},
			expectedHash: 0x545d934f6f5f29f9,
		},
		{
			name:         "DotAtRight",
			i:            &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3},
			expectedHash: 0x8f0b17c51d8b571e,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedHash, tc.i.Hash())
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
			i:             &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
			rhs:           &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			i:             &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
			rhs:           &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
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
				{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
				{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
				{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
				{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
				{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
				{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
			},
			expectedItems: []*Item0{
				{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
				{Production: parsertest.Prods[3][2], Start: "E′", Dot: 0}, // E → •T
				{Production: parsertest.Prods[3][5], Start: "E′", Dot: 0}, // F → •( E )
				{Production: parsertest.Prods[3][6], Start: "E′", Dot: 0}, // F → •id
				{Production: parsertest.Prods[3][3], Start: "E′", Dot: 0}, // T → •T * F
				{Production: parsertest.Prods[3][4], Start: "E′", Dot: 0}, // T → •F
			},
		},
		{
			name: "I₈",
			items: []*Item0{
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E• + T
				{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
			},
			expectedItems: []*Item0{
				{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E• + T
			},
		},
		{
			name: "I₉",
			items: []*Item0{
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
				{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
			},
			expectedItems: []*Item0{
				{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
				{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
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
			i:                &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
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
			i:                &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 0}, // E → •E + T
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
			i:                  &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
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
			i:               &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
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
			i:                 &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			expectedDotSymbol: grammar.NonTerminal("E"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
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
			i:            &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			expectedOK:   true,
			expectedNext: &Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
		},
		{
			name:         "Complete",
			i:            &Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
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
			i:             &Item0{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1}, // S → C•C
			lookahead:     grammar.Endmarker,
			expectedItem1: &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
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
				Start:      "S′",
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
			expectedString: `S → •, $`,
		},
		{
			name:           "Initial",
			i:              &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S′ → •S, $`,
		},
		{
			name:           "DotAtLeft",
			i:              &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},
			expectedString: `S → •C C, $`,
		},
		{
			name:           "DotInMiddle",
			i:              &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker},
			expectedString: `S → C•C, $`,
		},
		{
			name:           "DotAtRight",
			i:              &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker},
			expectedString: `S → C C•, $`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.i.String())
		})
	}
}

func TestItem1_Hash(t *testing.T) {
	tests := []struct {
		name         string
		i            *Item1
		expectedHash uint64
	}{
		{
			name: "EmptyProduction",
			i: &Item1{
				Production: &grammar.Production{Head: "S", Body: grammar.E},
				Start:      "S′",
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
			expectedHash: 0xe8dbda3d0a9c1cc1,
		},
		{
			name:         "Initial",
			i:            &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},
			expectedHash: 0xa3e7c4c7838e3ef6,
		},
		{
			name:         "DotAtLeft",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker},
			expectedHash: 0x35c3ab3505bcedb7,
		},
		{
			name:         "DotInMiddle",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker},
			expectedHash: 0xa421ea49649eee3a,
		},
		{
			name:         "DotAtRight",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker},
			expectedHash: 0x1280295dc380eebd,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedHash, tc.i.Hash())
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
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:           &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:           &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
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
			i:               &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			rhs:             &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondInitial",
			i:               &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedCompare: 1,
		},
		{
			name:            "FirstKernel",
			i:               &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedCompare: -1,
		},
		{
			name:            "SecondKernel",
			i:               &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			rhs:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: 1,
		},
		{
			name:            "FirstHead",
			i:               &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			rhs:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: -1,
		},
		{
			name:            "SecondHead",
			i:               &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstDot",
			i:               &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			rhs:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedCompare: -1,
		},
		{
			name:            "SecondDot",
			i:               &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			rhs:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedCompare: 1,
		},
		{
			name:            "FirstProduction",
			i:               &Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			rhs:             &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			expectedCompare: -1,
		},
		{
			name:            "SecondProduction",
			i:               &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &Item1{Production: parsertest.Prods[1][2], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •cC, c
			expectedCompare: 1,
		},
		{
			name:            "FirstLookahead",
			i:               &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
			rhs:             &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			expectedCompare: -1,
		},
		{
			name:            "SecondLookahead",
			i:               &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("d")}, // C → •d, d
			rhs:             &Item1{Production: parsertest.Prods[1][3], Start: "S′", Dot: 0, Lookahead: grammar.Terminal("c")}, // C → •d, c
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
			i:                &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "NotInitial",
			i:                &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
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
			i:                &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedIsKernel: true,
		},
		{
			name:             "Kernel",
			i:                &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedIsKernel: true,
		},
		{
			name:             "NonKernel",
			i:                &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // E → •CC, $
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
			i:                  &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
			expectedIsComplete: true,
		},
		{
			name:               "NotComplete",
			i:                  &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
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
			i:               &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
			expectedIsFinal: true,
		},
		{
			name:            "NotFinal",
			i:               &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
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
			i:                 &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedDotSymbol: grammar.NonTerminal("S"),
			expectedOK:        true,
		},
		{
			name:              "Complete",
			i:                 &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
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
			i:            &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			expectedOK:   true,
			expectedNext: &Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S′ → S•, $
		},
		{
			name:         "Complete",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•
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
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedAlpha: grammar.E,
		},
		{
			name:          "DotInMiddle",
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedAlpha: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:          "DotAtRight",
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
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
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S → •CC, $
			expectedBeta: grammar.String[grammar.Symbol]{grammar.NonTerminal("C"), grammar.NonTerminal("C")},
		},
		{
			name:         "DotInMiddle",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedBeta: grammar.String[grammar.Symbol]{grammar.NonTerminal("C")},
		},
		{
			name:         "DotAtRight",
			i:            &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 2, Lookahead: grammar.Endmarker}, // S → CC•, $
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
			i:             &Item1{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1, Lookahead: grammar.Endmarker}, // S → C•C, $
			expectedItem0: &Item0{Production: parsertest.Prods[1][1], Start: "S′", Dot: 1},                               // S → C•C
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

func TestNewItemSet(t *testing.T) {
	tests := []struct {
		name               string
		items              []Item
		expectedSubstrings []string
	}{
		{
			name: "OK",
			items: []Item{
				&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
				&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
			},
			expectedSubstrings: []string{
				`┌─────────────┐`,
				`│ E′ → E•     │`,
				`│ E → E•"+" T │`,
				`└─────────────┘`,
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

			str := s.String()
			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}

func TestCmpItemSet(t *testing.T) {
	tests := []struct {
		name         string
		sets         []ItemSet
		expectedSets []ItemSet
	}{
		{
			name: "OK",
			sets: []ItemSet{
				LR0ItemSets[0],
				LR0ItemSets[1],
				LR0ItemSets[2],
				LR0ItemSets[3],
				LR0ItemSets[4],
				LR0ItemSets[5],
				LR0ItemSets[6],
				LR0ItemSets[7],
				LR0ItemSets[8],
				LR0ItemSets[9],
				LR0ItemSets[10],
				LR0ItemSets[11],
			},
			expectedSets: []ItemSet{
				LR0ItemSets[0],
				LR0ItemSets[1],
				LR0ItemSets[9],
				LR0ItemSets[11],
				LR0ItemSets[10],
				LR0ItemSets[6],
				LR0ItemSets[8],
				LR0ItemSets[7],
				LR0ItemSets[2],
				LR0ItemSets[4],
				LR0ItemSets[5],
				LR0ItemSets[3],
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sort.Quick(tc.sets, cmpItemSet)
			assert.Equal(t, tc.expectedSets, tc.sets)
		})
	}
}

func TestItemSetStringer(t *testing.T) {
	state := State(0)

	tests := []struct {
		name               string
		ss                 *itemSetStringer
		expectedSubstrings []string
	}{
		{
			name: "WithoutState",
			ss: &itemSetStringer{
				items: generic.Collect1(LR0ItemSets[0].All()),
			},
			expectedSubstrings: []string{
				`┌────────────────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
			},
		},
		{
			name: "WithState",
			ss: &itemSetStringer{
				state: &state,
				items: generic.Collect1(LR0ItemSets[0].All()),
			},
			expectedSubstrings: []string{
				`┌──────[0]───────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.ss.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}

func TestNewItemSetCollection(t *testing.T) {
	tests := []struct {
		name               string
		sets               []ItemSet
		expectedSubstrings []string
	}{
		{
			name: "OK",
			sets: []ItemSet{
				LR0ItemSets[0],
				LR0ItemSets[1],
				LR0ItemSets[2],
				LR0ItemSets[3],
				LR0ItemSets[4],
				LR0ItemSets[5],
				LR0ItemSets[6],
				LR0ItemSets[7],
				LR0ItemSets[8],
				LR0ItemSets[9],
				LR0ItemSets[10],
				LR0ItemSets[11],
			},
			expectedSubstrings: []string{
				`┌──────[0]───────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[1]───────┐`,
				`│ E′ → E•        │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[2]───────┐`,
				`│ E → T•         │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[3]───────┐`,
				`│ T → F•         │`,
				`└────────────────┘`,
				`┌──────[4]───────┐`,
				`│ F → "("•E ")"  │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[5]───────┐`,
				`│ F → "id"•      │`,
				`└────────────────┘`,
				`┌──────[6]───────┐`,
				`│ E → E "+"•T    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[7]───────┐`,
				`│ T → T "*"•F    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`└────────────────┘`,
				`┌──────[8]───────┐`,
				`│ F → "(" E•")"  │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[9]───────┐`,
				`│ E → E "+" T•   │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[10]──────┐`,
				`│ T → T "*" F•   │`,
				`└────────────────┘`,
				`┌──────[11]──────┐`,
				`│ F → "(" E ")"• │`,
				`└────────────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewItemSetCollection(tc.sets...)

			assert.NotNil(t, s)

			for _, expectedItemSet := range tc.sets {
				assert.True(t, s.Contains(expectedItemSet))
			}

			str := s.String()
			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}

func TestItemSetCollectionStringer(t *testing.T) {
	tests := []struct {
		name               string
		cs                 *itemSetCollectionStringer
		expectedSubstrings []string
	}{
		{
			name: "OK",
			cs: &itemSetCollectionStringer{
				sets: []ItemSet{
					LR0ItemSets[0],
					LR0ItemSets[1],
					LR0ItemSets[2],
					LR0ItemSets[3],
					LR0ItemSets[4],
					LR0ItemSets[5],
					LR0ItemSets[6],
					LR0ItemSets[7],
					LR0ItemSets[8],
					LR0ItemSets[9],
					LR0ItemSets[10],
					LR0ItemSets[11],
				},
			},
			expectedSubstrings: []string{
				`┌──────[0]───────┐`,
				`│ E′ → •E        │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[1]───────┐`,
				`│ E′ → E•        │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[2]───────┐`,
				`│ E → T•         │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[3]───────┐`,
				`│ T → F•         │`,
				`└────────────────┘`,
				`┌──────[4]───────┐`,
				`│ F → "("•E ")"  │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ E → •E "+" T   │`,
				`│ E → •T         │`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[5]───────┐`,
				`│ F → "id"•      │`,
				`└────────────────┘`,
				`┌──────[6]───────┐`,
				`│ E → E "+"•T    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`│ T → •T "*" F   │`,
				`│ T → •F         │`,
				`└────────────────┘`,
				`┌──────[7]───────┐`,
				`│ T → T "*"•F    │`,
				`├╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌╌┤`,
				`│ F → •"(" E ")" │`,
				`│ F → •"id"      │`,
				`└────────────────┘`,
				`┌──────[8]───────┐`,
				`│ F → "(" E•")"  │`,
				`│ E → E•"+" T    │`,
				`└────────────────┘`,
				`┌──────[9]───────┐`,
				`│ E → E "+" T•   │`,
				`│ T → T•"*" F    │`,
				`└────────────────┘`,
				`┌──────[10]──────┐`,
				`│ T → T "*" F•   │`,
				`└────────────────┘`,
				`┌──────[11]──────┐`,
				`│ F → "(" E ")"• │`,
				`└────────────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			str := tc.cs.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, str, expectedSubstring)
			}
		})
	}
}
