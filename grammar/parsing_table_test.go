package grammar

import (
	"testing"

	"github.com/moorara/algo/set"
	"github.com/stretchr/testify/assert"
)

func getTestParsingTables() []*parsingTable {
	pt0 := NewParsingTable(
		[]Terminal{"+", "*", "(", ")", "id"},
		[]NonTerminal{"E", "E′", "T", "T′", "F"},
	).(*parsingTable)

	pt0.Add("E", "id", Production{"E", String[Symbol]{NonTerminal("T"), NonTerminal("E′")}})
	pt0.Add("E", "(", Production{"E", String[Symbol]{NonTerminal("T"), NonTerminal("E′")}})
	pt0.Add("E′", "+", Production{"E′", String[Symbol]{Terminal("+"), NonTerminal("T"), NonTerminal("E′")}})
	pt0.Add("E′", ")", Production{"E′", ε})
	pt0.Add("E′", endmarker, Production{"E′", ε})
	pt0.Add("T", "id", Production{"T", String[Symbol]{NonTerminal("F"), NonTerminal("T′")}})
	pt0.Add("T", "(", Production{"T", String[Symbol]{NonTerminal("F"), NonTerminal("T′")}})
	pt0.Add("T′", "+", Production{"T′", ε})
	pt0.Add("T′", "*", Production{"T′", String[Symbol]{Terminal("*"), NonTerminal("F"), NonTerminal("T′")}})
	pt0.Add("T′", ")", Production{"T′", ε})
	pt0.Add("T′", endmarker, Production{"T′", ε})
	pt0.Add("F", "id", Production{"F", String[Symbol]{Terminal("id")}})
	pt0.Add("F", "(", Production{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}})

	pt1 := NewParsingTable(
		[]Terminal{"a", "b", "e", "i", "t"},
		[]NonTerminal{"S", "S′", "E"},
	).(*parsingTable)

	pt1.Add("S", "a", Production{"S", String[Symbol]{Terminal("a")}})
	pt1.Add("S", "i", Production{"S", String[Symbol]{Terminal("i"), NonTerminal("E"), Terminal("t"), NonTerminal("S"), NonTerminal("S′")}})
	pt1.Add("S′", "e", Production{"S′", ε})
	pt1.Add("S′", "e", Production{"S′", String[Symbol]{Terminal("e"), NonTerminal("S")}})
	pt1.Add("S′", endmarker, Production{"S′", ε})
	pt1.Add("E", "b", Production{"E", String[Symbol]{Terminal("b")}})

	pt2 := NewParsingTable(
		[]Terminal{"+", "*", "(", ")", "id"},
		[]NonTerminal{"E", "T", "F"},
	).(*parsingTable)

	return []*parsingTable{pt0, pt1, pt2}
}

func TestNewParsingTable(t *testing.T) {
	tests := []struct {
		name                 string
		terminals            []Terminal
		nonTerminals         []NonTerminal
		expectedTerminals    []Terminal
		expectedNonTerminals []NonTerminal
	}{
		{
			name:                 "OK",
			terminals:            []Terminal{"+", "*", "(", ")", "id"},
			nonTerminals:         []NonTerminal{"E", "E′", "T", "T′", "F"},
			expectedTerminals:    []Terminal{"+", "*", "(", ")", "id", endmarker},
			expectedNonTerminals: []NonTerminal{"E", "E′", "T", "T′", "F"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pt := NewParsingTable(tc.terminals, tc.nonTerminals).(*parsingTable)

			assert.NotNil(t, pt)
			assert.Equal(t, tc.expectedTerminals, pt.terminals)
			assert.Equal(t, tc.expectedNonTerminals, pt.nonTerminals)
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
				`│      E       │               │               │   E → T E′    │        │ E → T E′ │        │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      E′      │ E′ → "+" T E′ │               │               │ E′ → ε │          │ E′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T       │               │               │   T → F T′    │        │ T → F T′ │        │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      T′      │    T′ → ε     │ T′ → "*" F T′ │               │ T′ → ε │          │ T′ → ε │`,
				`├──────────────┼───────────────┼───────────────┼───────────────┼────────┼──────────┼────────┤`,
				`│      F       │               │               │ F → "(" E ")" │        │ F → "id" │        │`,
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

func TestParsingTable_Add(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *parsingTable
		A          NonTerminal
		a          Terminal
		prod       Production
		expectedOK bool
	}{
		{
			name: "OK",
			pt:   pt[2],
			A:    NonTerminal("F"),
			a:    Terminal("("),
			prod: Production{"F", String[Symbol]{Terminal("("), NonTerminal("E"), Terminal(")")}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.pt.Add(tc.A, tc.a, tc.prod)
		})
	}
}

func TestParsingTable_Get(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                string
		pt                  *parsingTable
		A                   NonTerminal
		a                   Terminal
		expectedProductions set.Set[Production]
	}{
		{
			name: "OK",
			pt:   pt[0],
			A:    NonTerminal("E′"),
			a:    Terminal("+"),
			expectedProductions: set.New(eqProduction,
				Production{"E′", String[Symbol]{Terminal("+"), NonTerminal("T"), NonTerminal("E′")}},
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prods := tc.pt.Get(tc.A, tc.a)
			assert.True(t, prods.Equals(tc.expectedProductions))
		})
	}
}

func TestParsingTable_CheckErrors(t *testing.T) {
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
			err := tc.pt.CheckErrors()

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
