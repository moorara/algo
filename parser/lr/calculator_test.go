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
			G:    grammars[1],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "*", "(", ")", "id"},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				prods[1],
				"E′",
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			augG := Augment(tc.G)
			assert.True(t, augG.Equals(tc.expectedCFG))
		})
	}
}

func TestAutomatonCalculator_GOTO(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name         string
		a            *AutomatonCalculator
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name: `GOTO(I₀,E)`,
			a: &AutomatonCalculator{
				Calculator: &calculator0{
					augG: Augment(grammars[1]),
				},
			},
			I:            s[0],
			X:            grammar.NonTerminal("E"),
			expectedGOTO: s[1],
		},
		{
			name: `GOTO(I₀,T)`,
			a: &AutomatonCalculator{
				Calculator: &calculator0{
					augG: Augment(grammars[1]),
				},
			},
			I:            s[0],
			X:            grammar.NonTerminal("T"),
			expectedGOTO: s[2],
		},
		{
			name: `GOTO(I₀,F)`,
			a: &AutomatonCalculator{
				Calculator: &calculator0{
					augG: Augment(grammars[1]),
				},
			},
			I:            s[0],
			X:            grammar.NonTerminal("F"),
			expectedGOTO: s[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := tc.a.GOTO(tc.I, tc.X)
			assert.True(t, J.Equals(tc.expectedGOTO))
		})
	}
}

func TestAutomatonCalculator_Canonical(t *testing.T) {
	s := getTestLR0ItemSets()

	tests := []struct {
		name              string
		a                 *AutomatonCalculator
		expectedCanonical ItemSetCollection
	}{
		{
			name: "OK",
			a: &AutomatonCalculator{
				Calculator: &calculator0{
					augG: Augment(grammars[1]),
				},
			},
			expectedCanonical: NewItemSetCollection(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7], s[8], s[9], s[10], s[11]),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := tc.a.Canonical()
			assert.True(t, C.Equals(tc.expectedCanonical))
		})
	}
}

func TestNewLR0AutomatonCalculator(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR0AutomatonCalculator(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.Calculator)
			assert.NotEmpty(t, calc.Calculator.(*calculator0).augG)
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
				augG: Augment(grammars[1]),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			G := tc.c.G()

			assert.True(t, G.Equals(tc.c.augG))
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
				augG: Augment(grammars[1]),
			},
			expectedInitial: &Item0{
				Production: prods[1][0],
				Start:      starts[1],
				Dot:        0,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			initial := tc.c.Initial()

			assert.True(t, initial.Equals(tc.expectedInitial))
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
				augG: Augment(grammars[1]),
			},
			I: NewItemSet(
				&Item0{Production: prods[1][0], Start: starts[1], Dot: 0}, // E′ → •E
			),
			expectedCLOSURE: s[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			J := tc.c.CLOSURE(tc.I)

			assert.True(t, J.Equals(tc.expectedCLOSURE))
		})
	}
}

func TestNewLR1AutomatonCalculator(t *testing.T) {
	tests := []struct {
		name string
		G    *grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			calc := NewLR1AutomatonCalculator(tc.G)

			assert.NotNil(t, calc)
			assert.NotNil(t, calc.Calculator)
			assert.NotEmpty(t, calc.Calculator.(*calculator1).augG)
			assert.NotNil(t, calc.Calculator.(*calculator1).FIRST)
		})
	}
}

func TestCalculator_G(t *testing.T) {
	tests := []struct {
		name string
		c    *calculator1
	}{
		{
			name: "OK",
			c: &calculator1{
				augG: Augment(grammars[0]),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.c.augG.Verify())
			G := tc.c.G()

			assert.True(t, G.Equals(tc.c.augG))
		})
	}
}

func TestCalculator_Initial(t *testing.T) {
	tests := []struct {
		name            string
		c               *calculator1
		expectedInitial Item
	}{
		{
			name: "OK",
			c: &calculator1{
				augG: Augment(grammars[0]),
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

			assert.True(t, initial.Equals(tc.expectedInitial))
		})
	}
}

func TestCalculator_CLOSURE(t *testing.T) {
	s := getTestLR1ItemSets()
	g := Augment(grammars[0])

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

			assert.True(t, J.Equals(tc.expectedCLOSURE))
		})
	}
}
