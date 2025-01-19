package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

var starts = []grammar.NonTerminal{
	"E′",
}

var prods = [][]grammar.Production{
	{
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
}

var grammars = []grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[0],
		"E",
	),
}

func TestAugment(t *testing.T) {
	tests := []struct {
		name        string
		G           grammar.CFG
		expectedCFG grammar.CFG
	}{
		{
			name: "OK",
			G:    grammars[0],
			expectedCFG: grammar.NewCFG(
				[]grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
				[]grammar.NonTerminal{"E′", "E", "T", "F"},
				[]grammar.Production{
					{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
					{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
					{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
					{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
					{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
					{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
					{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
				},
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

func TestGOTO(t *testing.T) {
	s := getTestItemSets()

	mockClosure := func(I ItemSet) ItemSet {
		return I
	}

	tests := []struct {
		name         string
		CLOSURE      ClosureFunc
		I            ItemSet
		X            grammar.Symbol
		expectedGOTO ItemSet
	}{
		{
			name:         `GOTO(I₀,E)`,
			CLOSURE:      mockClosure,
			I:            s[0],
			X:            grammar.NonTerminal("E"),
			expectedGOTO: s[1],
		},
		{
			name:         `GOTO(I₀,T)`,
			CLOSURE:      mockClosure,
			I:            s[0],
			X:            grammar.NonTerminal("T"),
			expectedGOTO: s[2],
		},
		{
			name:         `GOTO(I₀,F)`,
			CLOSURE:      mockClosure,
			I:            s[0],
			X:            grammar.NonTerminal("F"),
			expectedGOTO: s[3],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			J := GOTO(tc.CLOSURE, tc.I, tc.X)
			assert.True(t, J.Equals(tc.expectedGOTO))
		})
	}
}

func TestCanonical(t *testing.T) {
	g := Augment(grammars[0])

	mockClosure := func(I ItemSet) ItemSet {
		return I
	}

	tests := []struct {
		name              string
		augG              grammar.CFG
		initial           Item
		CLOSURE           ClosureFunc
		expectedCanonical ItemSetCollection
	}{
		{
			name:    "OK",
			augG:    g,
			initial: mockItem("E′→•E"),
			CLOSURE: mockClosure,
			expectedCanonical: NewItemSetCollection(
				NewItemSet(
					mockItem("E′→•E"),
				),
				NewItemSet(
					mockItem("E′→E•"),
				),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			C := Canonical(tc.augG, tc.initial, tc.CLOSURE)
			assert.True(t, C.Equals(tc.expectedCanonical))
		})
	}
}
