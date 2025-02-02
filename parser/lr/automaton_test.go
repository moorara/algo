package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
)

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
