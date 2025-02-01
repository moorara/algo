package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
)

func TestAugment(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		expectedCFG *grammar.CFG
	}{
		{
			name: "OK",
			G:    parsertest.Grammars[3],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				parsertest.Prods[3],
				"E′",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			augG := augment(tc.G)
			assert.True(t, augG.Equal(tc.expectedCFG))
		})
	}
}

func TestNewGrammarWithLR0(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR0(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &automaton{})
			assert.IsType(t, G.Automaton.(*automaton).calculator, &calculator0{})
		})
	}
}

func TestNewGrammarWithLR1(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR1(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &automaton{})
			assert.IsType(t, G.Automaton.(*automaton).calculator, &calculator1{})
		})
	}
}

func TestNewGrammarWithLR0Kernel(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR0Kernel(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &kernelAutomaton{})
			assert.IsType(t, G.Automaton.(*kernelAutomaton).calculator, &calculator0{})
		})
	}
}

func TestNewGrammarWithLR1Kernel(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		precedences PrecedenceLevels
	}{
		{
			name:        "OK",
			G:           parsertest.Grammars[3],
			precedences: PrecedenceLevels{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			G := NewGrammarWithLR1Kernel(tc.G, tc.precedences)

			assert.NotNil(t, G)
			assert.NotNil(t, G.CFG)
			assert.NotNil(t, G.PrecedenceLevels)
			assert.NotNil(t, G.Automaton)
			assert.IsType(t, G.Automaton, &kernelAutomaton{})
			assert.IsType(t, G.Automaton.(*kernelAutomaton).calculator, &calculator1{})
		})
	}
}

func TestPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name           string
		handles        []*PrecedenceHandle
		expectedString string
	}{
		{
			name:           "Zero",
			handles:        []*PrecedenceHandle{},
			expectedString: ``,
		},
		{
			name: "One",
			handles: []*PrecedenceHandle{
				handles[1],
			},
			expectedString: `"|"`,
		},
		{
			name: "Two",
			handles: []*PrecedenceHandle{
				handles[1],
				handles[2],
			},
			expectedString: `"(", "|"`,
		},
		{
			name: "Three",
			handles: []*PrecedenceHandle{
				handles[1],
				handles[2],
				handles[3],
			},
			expectedString: `"(", "[", "|"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewPrecedenceHandles(tc.handles...)

			assert.Equal(t, tc.expectedString, h.String())
		})
	}
}

func TestCmpPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name            string
		lhs             PrecedenceHandles
		rhs             PrecedenceHandles
		expectedCompare int
	}{
		{
			name: "FirstShorter",
			lhs: NewPrecedenceHandles(
				handles[1],
			),
			rhs: NewPrecedenceHandles(
				handles[1],
				handles[2],
			),
			expectedCompare: -1,
		},
		{
			name: "FirstLonger",
			lhs: NewPrecedenceHandles(
				handles[1],
				handles[2],
			),
			rhs: NewPrecedenceHandles(
				handles[1],
			),
			expectedCompare: 1,
		},
		{
			name: "EqualLength",
			lhs: NewPrecedenceHandles(
				handles[1],
				handles[2],
			),
			rhs: NewPrecedenceHandles(
				handles[1],
				handles[3],
			),
			expectedCompare: -1,
		},
		{
			name: "Equal",
			lhs: NewPrecedenceHandles(
				handles[1],
				handles[3],
			),
			rhs: NewPrecedenceHandles(
				handles[1],
				handles[3],
			),
			expectedCompare: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCompare, cmpPrecedenceHandles(tc.lhs, tc.rhs))
		})
	}
}

func TestPrecedenceHandle(t *testing.T) {
	type EqualTest struct {
		rhs           *PrecedenceHandle
		expectedEqual bool
	}

	tests := []struct {
		name                 string
		h                    *PrecedenceHandle
		expectedIsTerminal   bool
		expectedIsProduction bool
		expectedString       string
		equalTests           []EqualTest
	}{
		{
			name:                 "Terminal",
			h:                    handles[6],
			expectedIsTerminal:   true,
			expectedIsProduction: false,
			expectedString:       `"IDENT"`,
			equalTests: []EqualTest{
				{
					rhs:           handles[6],
					expectedEqual: true,
				},
				{
					rhs:           handles[7],
					expectedEqual: false,
				},
				{
					rhs:           handles[9],
					expectedEqual: false,
				},
			},
		},
		{
			name:                 "Production",
			h:                    handles[9],
			expectedIsTerminal:   false,
			expectedIsProduction: true,
			expectedString:       `rhs = rhs rhs`,
			equalTests: []EqualTest{
				{
					rhs:           handles[9],
					expectedEqual: true,
				},
				{
					rhs:           handles[10],
					expectedEqual: false,
				},
				{
					rhs:           handles[6],
					expectedEqual: false,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsTerminal, tc.h.IsTerminal())
			assert.Equal(t, tc.expectedIsProduction, tc.h.IsProduction())
			assert.Equal(t, tc.expectedString, tc.h.String())

			t.Run("Equal", func(t *testing.T) {
				for _, test := range tc.equalTests {
					assert.Equal(t, test.expectedEqual, tc.h.Equal(test.rhs))
				}
			})
		})
	}
}

func TestCmpPrecedenceHandle(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *PrecedenceHandle
		rhs             *PrecedenceHandle
		expectedCompare int
	}{
		{
			name:            "BothTerminal",
			lhs:             handles[6],
			rhs:             handles[7],
			expectedCompare: -1,
		},
		{
			name:            "BothProduction",
			lhs:             handles[9],
			rhs:             handles[10],
			expectedCompare: 1,
		},
		{
			name:            "Mixed",
			lhs:             handles[6],
			rhs:             handles[9],
			expectedCompare: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCompare, cmpPrecedenceHandle(tc.lhs, tc.rhs))

		})
	}
}

func TestAutomaton_GOTO(t *testing.T) {
	G := augment(parsertest.Grammars[3])

	tests := []struct {
		name         string
		a            *automaton
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name: `GOTO(I₀,"(")`,
			a: &automaton{
				G: G,
				calculator: &calculator0{
					G: G,
				},
			},
			I:            LR0ItemSets[0],
			X:            grammar.Terminal("("),
			expectedGOTO: LR0ItemSets[4],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.a.GOTO(tc.I, tc.X)
			assert.True(t, J.Equal(tc.expectedGOTO))
		})
	}
}

func TestAutomaton_Canonical(t *testing.T) {
	G := augment(parsertest.Grammars[3])

	tests := []struct {
		name              string
		a                 *automaton
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &automaton{
				G: G,
				calculator: &calculator0{
					G: G,
				},
			},
			expectedCanonical: NewItemSetCollection(
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
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.a.Canonical()
			assert.True(t, C.Equal(tc.expectedCanonical))
		})
	}
}

func TestKernelAutomaton_GOTO(t *testing.T) {
	G := augment(parsertest.Grammars[3])

	tests := []struct {
		name         string
		a            *kernelAutomaton
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name: `GOTO(I₀,"(")`,
			a: &kernelAutomaton{
				G: G,
				calculator: &calculator0{
					G: G,
				},
			},
			I: NewItemSet(
				&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			),
			X: grammar.Terminal("("),
			expectedGOTO: NewItemSet(
				&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1}, // F → (•E ),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.a.GOTO(tc.I, tc.X)
			assert.True(t, J.Equal(tc.expectedGOTO))
		})
	}
}

func TestKernelAutomaton_Canonical(t *testing.T) {
	G := augment(parsertest.Grammars[3])

	tests := []struct {
		name              string
		a                 *kernelAutomaton
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &kernelAutomaton{
				G: G,
				calculator: &calculator0{
					G: G,
				},
			},
			expectedCanonical: NewItemSetCollection(
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E,
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 1}, // E′ → E•
					&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E•+ T
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][2], Start: "E′", Dot: 1}, // E → T•
					&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][4], Start: "E′", Dot: 1}, // T → F•
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 1}, // F → (•E )
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][6], Start: "E′", Dot: 1}, // F → id•
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 2}, // E → E +•T
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 2}, // T → T *•F
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 1}, // E → E• + T
					&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 2}, // F → ( E•)
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][1], Start: "E′", Dot: 3}, // E → E + T•
					&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 1}, // T → T•* F
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][3], Start: "E′", Dot: 3}, // T → T * F•
				),
				NewItemSet(
					&Item0{Production: parsertest.Prods[3][5], Start: "E′", Dot: 3}, // F → ( E )•
				),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.a.Canonical()
			assert.True(t, C.Equal(tc.expectedCanonical))
		})
	}
}

func TestCalculator0_Initial(t *testing.T) {
	tests := []struct {
		name            string
		c               *calculator0
		expectedInitial Item
	}{
		{
			name: "OK",
			c: &calculator0{
				G: augment(parsertest.Grammars[3]),
			},
			expectedInitial: &Item0{
				Production: parsertest.Prods[3][0],
				Start:      "E′",
				Dot:        0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.G.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equal(tc.expectedInitial))
		})
	}
}

func TestCalculator0_CLOSURE(t *testing.T) {
	tests := []struct {
		name            string
		c               *calculator0
		I               ItemSet
		expectedCLOSURE ItemSet
	}{
		{
			name: "OK",
			c: &calculator0{
				G: augment(parsertest.Grammars[3]),
			},
			I: NewItemSet(
				&Item0{Production: parsertest.Prods[3][0], Start: "E′", Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: LR0ItemSets[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.G.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equal(tc.expectedCLOSURE))
		})
	}
}

func TestCalculator1_Initial(t *testing.T) {
	tests := []struct {
		name            string
		c               *calculator1
		expectedInitial Item
	}{
		{
			name: "OK",
			c: &calculator1{
				G: augment(parsertest.Grammars[1]),
			},
			expectedInitial: &Item1{
				Production: parsertest.Prods[1][0],
				Start:      "S′",
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.G.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equal(tc.expectedInitial))
		})
	}
}

func TestCalculator1_CLOSURE(t *testing.T) {
	g := augment(parsertest.Grammars[1])

	tests := []struct {
		name            string
		c               *calculator1
		I               ItemSet
		expectedCLOSURE ItemSet
	}{
		{
			name: "OK",
			c: &calculator1{
				G:     g,
				FIRST: g.ComputeFIRST(),
			},
			I: NewItemSet(
				&Item1{Production: parsertest.Prods[1][0], Start: "S′", Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			),
			expectedCLOSURE: LR1ItemSets[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.G.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equal(tc.expectedCLOSURE))
		})
	}
}
