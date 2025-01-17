package predictive

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

func getTestParsingTables() []*parsingTable {
	pt0 := newParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "E′", "T", "T′", "F"},
	)

	pt0.AddProduction("E", "id", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.AddProduction("E", "(", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.AddProduction("E′", "+", grammar.Production{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")}})
	pt0.AddProduction("E′", ")", grammar.Production{Head: "E′", Body: grammar.E})
	pt0.AddProduction("E′", grammar.Endmarker, grammar.Production{Head: "E′", Body: grammar.E})
	pt0.AddProduction("T", "id", grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.AddProduction("T", "(", grammar.Production{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.AddProduction("T′", "+", grammar.Production{Head: "T′", Body: grammar.E})
	pt0.AddProduction("T′", "*", grammar.Production{Head: "T′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("F"), grammar.NonTerminal("T′")}})
	pt0.AddProduction("T′", ")", grammar.Production{Head: "T′", Body: grammar.E})
	pt0.AddProduction("T′", grammar.Endmarker, grammar.Production{Head: "T′", Body: grammar.E})
	pt0.AddProduction("F", "id", grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}})
	pt0.AddProduction("F", "(", grammar.Production{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}})

	pt0.SetSync("E", ")", true)
	pt0.SetSync("E", grammar.Endmarker, true)
	pt0.SetSync("T", "+", true)
	pt0.SetSync("T", ")", true)
	pt0.SetSync("T", grammar.Endmarker, true)
	pt0.SetSync("F", "+", true)
	pt0.SetSync("F", "*", true)
	pt0.SetSync("F", ")", true)
	pt0.SetSync("F", grammar.Endmarker, true)

	pt1 := newParsingTable(
		[]grammar.Terminal{"a", "b", "e", "i", "t"},
		[]grammar.NonTerminal{"S", "S′", "E"},
	)

	pt1.AddProduction("S", "a", grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("a")}})
	pt1.AddProduction("S", "i", grammar.Production{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.Terminal("i"), grammar.NonTerminal("E"), grammar.Terminal("t"), grammar.NonTerminal("S"), grammar.NonTerminal("S′")}})
	pt1.AddProduction("S′", "e", grammar.Production{Head: "S′", Body: grammar.E})
	pt1.AddProduction("S′", "e", grammar.Production{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.Terminal("e"), grammar.NonTerminal("S")}})
	pt1.AddProduction("S′", grammar.Endmarker, grammar.Production{Head: "S′", Body: grammar.E})
	pt1.AddProduction("E", "b", grammar.Production{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("b")}})

	pt2 := newParsingTable(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	return []*parsingTable{pt0, pt1, pt2}
}

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    grammar.CFG
		expectedTable        ParsingTable
		expectedErrorStrings []string
	}{
		{
			name: "1st",
			G:    grammars[0],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[E, "("]`,
				`multiple productions in parsing table at M[E, "-"]`,
				`multiple productions in parsing table at M[E, "id"]`,
			},
		},
		{
			name: "2nd",
			G:    grammars[1],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[E, "("]`,
				`multiple productions in parsing table at M[E, "id"]`,
				`multiple productions in parsing table at M[T, "("]`,
				`multiple productions in parsing table at M[T, "id"]`,
			},
		},
		{
			name:                 "3rd",
			expectedTable:        pt[0],
			G:                    grammars[2],
			expectedErrorStrings: nil,
		},
		{
			name: "4th",
			G:    grammars[3],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[decls, "IDENT"]`,
				`multiple productions in parsing table at M[decls, "TOKEN"]`,
				`multiple productions in parsing table at M[rule, "IDENT"]`,
				`multiple productions in parsing table at M[token, "TOKEN"]`,
				`multiple productions in parsing table at M[rhs, "("]`,
				`multiple productions in parsing table at M[rhs, "IDENT"]`,
				`multiple productions in parsing table at M[rhs, "STRING"]`,
				`multiple productions in parsing table at M[rhs, "TOKEN"]`,
				`multiple productions in parsing table at M[rhs, "["]`,
				`multiple productions in parsing table at M[rhs, "{"]`,
				`multiple productions in parsing table at M[rhs, "{{"]`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table := BuildParsingTable(tc.G)
			err := table.Error()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equals(tc.expectedTable))
			} else {
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}

func TestParsingTable_String(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name               string
		pt                 *parsingTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			pt:   pt[0],
			expectedSubstrings: []string{
				`┌──────────────┬────────────────────────────────────────────────────────────────────────────┐`,
				`│              │                                  Terminal                                  │`,
				`├──────────────┼───────────────┬───────────────┬───────────────┬────────┬──────────┬────────┤`,
				`│ Non-Terminal │      "+"      │      "*"      │      "("      │  ")"   │   "id"   │   $    │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      E       │               │               │   E → T E′    │  sync  │ E → T E′ │  sync  │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      E′      │ E′ → "+" T E′ │               │               │ E′ → ε │          │ E′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T       │     sync      │               │   T → F T′    │  sync  │ T → F T′ │  sync  │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T′      │    T′ → ε     │ T′ → "*" F T′ │               │ T′ → ε │          │ T′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      F       │     sync      │     sync      │ F → "(" E ")" │  sync  │ F → "id" │  sync  │`,
				`└──────────────┴───────────────┴───────────────┴───────────────┴────────┴──────────┴────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.pt.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}

func TestParsingTable_Equals(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name           string
		pt             *parsingTable
		rhs            ParsingTable
		expectedEquals bool
	}{
		{
			name:           "Equal",
			pt:             pt[0],
			rhs:            pt[0],
			expectedEquals: true,
		},
		{
			name:           "NotEqual",
			pt:             pt[1],
			rhs:            pt[2],
			expectedEquals: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEquals, tc.pt.Equals(tc.rhs))
		})
	}
}

func TestParsingTable_Error(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		pt                   *parsingTable
		expectedErrorStrings []string
	}{
		{
			name:                 "NoError",
			pt:                   pt[0],
			expectedErrorStrings: nil,
		},
		{
			name: "Error",
			pt:   pt[1],
			expectedErrorStrings: []string{
				`multiple productions in parsing table at M[S′, "e"]`,
				`S′ → "e" S`,
				`S′ → ε`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.pt.Error()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}

func TestParsingTable_AddProduction(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *parsingTable
		A          grammar.NonTerminal
		a          grammar.Terminal
		prod       grammar.Production
		expectedOK bool
	}{
		{
			name: "OK",
			pt:   pt[2],
			A:    grammar.NonTerminal("F"),
			a:    grammar.Terminal("("),
			prod: grammar.Production{
				Head: "F",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")},
			},
			expectedOK: true,
		},
		{
			name: "IsSync",
			pt:   pt[0],
			A:    grammar.NonTerminal("F"),
			a:    grammar.Terminal(")"),
			prod: grammar.Production{
				Head: "F",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")},
			},
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.pt.AddProduction(tc.A, tc.a, tc.prod)
			assert.Equal(t, tc.expectedOK, ok)

			if tc.expectedOK {
				e, ok := tc.pt.getEntry(tc.A, tc.a)
				assert.True(t, ok)
				assert.True(t, e.Productions.Contains(tc.prod))
			}
		})
	}
}

func TestParsingTable_SetSync(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *parsingTable
		A          grammar.NonTerminal
		a          grammar.Terminal
		sync       bool
		expectedOK bool
	}{
		{
			name:       "OK",
			pt:         pt[0],
			A:          grammar.NonTerminal("F"),
			a:          grammar.Terminal(")"),
			sync:       true,
			expectedOK: true,
		},
		{
			name:       "HasProduction",
			pt:         pt[0],
			A:          grammar.NonTerminal("F"),
			a:          grammar.Terminal("("),
			sync:       false,
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.pt.SetSync(tc.A, tc.a, tc.sync)
			assert.Equal(t, tc.expectedOK, ok)

			if tc.expectedOK {
				e, ok := tc.pt.getEntry(tc.A, tc.a)
				assert.True(t, ok)
				assert.Equal(t, tc.sync, e.Sync)
			}
		})
	}
}

func TestParsingTable_IsEmpty(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name            string
		pt              *parsingTable
		A               grammar.NonTerminal
		a               grammar.Terminal
		expectedIsEmpty bool
	}{
		{
			name:            "Empty",
			pt:              pt[0],
			A:               grammar.NonTerminal("E"),
			a:               grammar.Terminal("+"),
			expectedIsEmpty: true,
		},
		{
			name:            "NotEmpty",
			pt:              pt[0],
			A:               grammar.NonTerminal("E"),
			a:               grammar.Terminal("id"),
			expectedIsEmpty: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsEmpty, tc.pt.IsEmpty(tc.A, tc.a))
		})
	}
}

func TestParsingTable_IsSync(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name           string
		pt             *parsingTable
		A              grammar.NonTerminal
		a              grammar.Terminal
		expectedIsSync bool
	}{
		{
			name:           "Sync",
			pt:             pt[0],
			A:              grammar.NonTerminal("E"),
			a:              grammar.Terminal(")"),
			expectedIsSync: true,
		},
		{
			name:           "NotSync",
			pt:             pt[0],
			A:              grammar.NonTerminal("E"),
			a:              grammar.Terminal("*"),
			expectedIsSync: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsSync, tc.pt.IsSync(tc.A, tc.a))
		})
	}
}

func TestParsingTable_GetProduction(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name               string
		pt                 *parsingTable
		A                  grammar.NonTerminal
		a                  grammar.Terminal
		expectedOK         bool
		expectedProduction grammar.Production
	}{
		{
			name:               "Empty",
			pt:                 pt[0],
			A:                  grammar.NonTerminal("E"),
			a:                  grammar.Terminal("+"),
			expectedOK:         false,
			expectedProduction: grammar.Production{},
		},
		{
			name:       "OK",
			pt:         pt[0],
			A:          grammar.NonTerminal("E′"),
			a:          grammar.Terminal("+"),
			expectedOK: true,
			expectedProduction: grammar.Production{
				Head: "E′",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prod, ok := tc.pt.GetProduction(tc.A, tc.a)

			assert.Equal(t, tc.expectedOK, ok)
			assert.True(t, prod.Equals(tc.expectedProduction))
		})
	}
}

func TestParsingTable_GetProductions(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                string
		pt                  *parsingTable
		A                   grammar.NonTerminal
		a                   grammar.Terminal
		expectedOK          bool
		expectedProductions set.Set[grammar.Production]
	}{
		{
			name:                "Empty",
			pt:                  pt[0],
			A:                   grammar.NonTerminal("E"),
			a:                   grammar.Terminal("+"),
			expectedOK:          false,
			expectedProductions: nil,
		},
		{
			name:       "OK",
			pt:         pt[0],
			A:          grammar.NonTerminal("E′"),
			a:          grammar.Terminal("+"),
			expectedOK: true,
			expectedProductions: set.New(grammar.EqProduction,
				grammar.Production{
					Head: "E′",
					Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")},
				},
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prods, ok := tc.pt.GetProductions(tc.A, tc.a)

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, prods.Equals(tc.expectedProductions))
			} else {
				assert.False(t, ok)
				assert.Nil(t, prods)
			}
		})
	}
}
