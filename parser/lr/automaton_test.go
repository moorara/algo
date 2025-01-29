package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

func TestAugment(t *testing.T) {
	tests := []struct {
		name        string
		G           *grammar.CFG
		expectedCFG *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				prods[2],
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

func TestCalculator0_G(t *testing.T) {
	tests := []struct {
		name string
		c    *calculator0
	}{
		{
			name: "OK",
			c: &calculator0{
				augG: augment(grammars[2]),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			G := tc.c.G()

			assert.True(t, G.Equal(tc.c.augG))
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
				augG: augment(grammars[2]),
			},
			expectedInitial: &Item0{
				Production: prods[2][0],
				Start:      starts[2],
				Dot:        0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equal(tc.expectedInitial))
		})
	}
}

func TestCalculator0_CLOSURE(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name            string
		c               *calculator0
		I               ItemSet
		expectedCLOSURE ItemSet
	}{
		{
			name: "OK",
			c: &calculator0{
				augG: augment(grammars[2]),
			},
			I: NewItemSet(
				&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equal(tc.expectedCLOSURE))
		})
	}
}

func TestCalculator1_G(t *testing.T) {
	tests := []struct {
		name string
		c    *calculator1
	}{
		{
			name: "OK",
			c: &calculator1{
				augG: augment(grammars[0]),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			G := tc.c.G()

			assert.True(t, G.Equal(tc.c.augG))
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
				augG: augment(grammars[0]),
			},
			expectedInitial: &Item1{
				Production: prods[0][0],
				Start:      starts[0],
				Dot:        0,
				Lookahead:  grammar.Endmarker,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equal(tc.expectedInitial))
		})
	}
}

func TestCalculator1_CLOSURE(t *testing.T) {
	s := getTestLR1ItemSets()
	g := augment(grammars[0])

	tests := []struct {
		name            string
		c               *calculator1
		I               ItemSet
		expectedCLOSURE ItemSet
	}{
		{
			name: "OK",
			c: &calculator1{
				augG:  g,
				FIRST: g.ComputeFIRST(),
			},
			I: NewItemSet(
				&Item1{Production: prods[0][0], Start: starts[0], Dot: 0, Lookahead: grammar.Endmarker}, // S′ → •S, $
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equal(tc.expectedCLOSURE))
		})
	}
}

func TestAutomaton_GOTO(t *testing.T) {
	s := getTestLR0ItemSets()

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
				calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			I:            s[0],
			X:            grammar.Terminal("("),
			expectedGOTO: s[4],
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
	s := getTestLR0ItemSets()

	tests := []struct {
		name              string
		a                 *automaton
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &automaton{
				calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			expectedCanonical: NewItemSetCollection(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.a.Canonical()
			assert.True(t, C.Equal(tc.expectedCanonical))
		})
	}
}

func TestNewLR0Automaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR0Automaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*automaton).calculator)
			assert.NotEmpty(t, calc.(*automaton).calculator.(*calculator0).augG)
		})
	}
}

func TestNewLR1Automaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR1Automaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*automaton).calculator)
			assert.NotEmpty(t, calc.(*automaton).calculator.(*calculator1).augG)
			assert.NotNil(t, calc.(*automaton).calculator.(*calculator1).FIRST)
		})
	}
}

func TestKernelAutomaton_GOTO(t *testing.T) {
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
				calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			I: NewItemSet(
				&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E
			),
			X: grammar.Terminal("("),
			expectedGOTO: NewItemSet(
				&Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E ),
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
	tests := []struct {
		name              string
		a                 *kernelAutomaton
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &kernelAutomaton{
				calculator: &calculator0{
					augG: augment(grammars[2]),
				},
			},
			expectedCanonical: NewItemSetCollection(
				NewItemSet(
					&Item0{Production: prods[2][0], Start: starts[2], Dot: 0}, // E′ → •E,
				),
				NewItemSet(
					&Item0{Production: prods[2][0], Start: starts[2], Dot: 1}, // E′ → E•
					&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E•+ T
				),
				NewItemSet(
					&Item0{Production: prods[2][2], Start: starts[2], Dot: 1}, // E → T•
					&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
				),
				NewItemSet(
					&Item0{Production: prods[2][4], Start: starts[2], Dot: 1}, // T → F•
				),
				NewItemSet(
					&Item0{Production: prods[2][5], Start: starts[2], Dot: 1}, // F → (•E )
				),
				NewItemSet(
					&Item0{Production: prods[2][6], Start: starts[2], Dot: 1}, // F → id•
				),
				NewItemSet(
					&Item0{Production: prods[2][1], Start: starts[2], Dot: 2}, // E → E +•T
				),
				NewItemSet(
					&Item0{Production: prods[2][3], Start: starts[2], Dot: 2}, // T → T *•F
				),
				NewItemSet(
					&Item0{Production: prods[2][1], Start: starts[2], Dot: 1}, // E → E• + T
					&Item0{Production: prods[2][5], Start: starts[2], Dot: 2}, // F → ( E•)
				),
				NewItemSet(
					&Item0{Production: prods[2][1], Start: starts[2], Dot: 3}, // E → E + T•
					&Item0{Production: prods[2][3], Start: starts[2], Dot: 1}, // T → T•* F
				),
				NewItemSet(
					&Item0{Production: prods[2][3], Start: starts[2], Dot: 3}, // T → T * F•
				),
				NewItemSet(
					&Item0{Production: prods[2][5], Start: starts[2], Dot: 3}, // F → ( E )•
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

func TestNewLR0KernelAutomaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR0KernelAutomaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*kernelAutomaton).calculator)
			assert.NotEmpty(t, calc.(*kernelAutomaton).calculator.(*calculator0).augG)
		})
	}
}

func TestNewLR1KernelAutomaton(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[2],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR1KernelAutomaton(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.(*kernelAutomaton).calculator)
			assert.NotEmpty(t, calc.(*kernelAutomaton).calculator.(*calculator1).augG)
			assert.NotNil(t, calc.(*kernelAutomaton).calculator.(*calculator1).FIRST)
		})
	}
}
