package predictive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/set"
)

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    *grammar.CFG
		expectedTable        *ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:          "E→TE′",
			G:             parsertest.Grammars[0],
			expectedTable: pt[0],
		},
		{
			name: "E→E+T",
			G:    parsertest.Grammars[3],
			expectedErrorStrings: []string{
				`4 errors occurred:`,
				`multiple productions at M[E, "("]:`,
				`E → E "+" T`,
				`E → T`,
				`multiple productions at M[E, "id"]:`,
				`E → E "+" T`,
				`E → T`,
				`multiple productions at M[T, "("]:`,
				`T → T "*" F`,
				`T → F`,
				`multiple productions at M[T, "id"]:`,
				`T → T "*" F`,
				`T → F`,
			},
		},
		{
			name: "E→E+E",
			G:    parsertest.Grammars[4],
			expectedErrorStrings: []string{
				`2 errors occurred:`,
				`multiple productions at M[E, "("]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "(" E ")"`,
				`multiple productions at M[E, "id"]:`,
				`E → E "*" E`,
				`E → E "+" E`,
				`E → "id"`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table, err := BuildParsingTable(tc.G)

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equal(tc.expectedTable))
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

func TestParsingTable_addProduction(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *ParsingTable
		A          grammar.NonTerminal
		a          grammar.Terminal
		prod       *grammar.Production
		expectedOK bool
	}{
		{
			name: "OK",
			pt:   pt[2],
			A:    grammar.NonTerminal("F"),
			a:    grammar.Terminal("("),
			prod: &grammar.Production{
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
			prod: &grammar.Production{
				Head: "F",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")},
			},
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.pt.addProduction(tc.A, tc.a, tc.prod)
			assert.Equal(t, tc.expectedOK, ok)

			if tc.expectedOK {
				e, ok := tc.pt.getEntry(tc.A, tc.a)
				assert.True(t, ok)
				assert.True(t, e.Productions.Contains(tc.prod))
			}
		})
	}
}

func TestParsingTable_setSync(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *ParsingTable
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
			ok := tc.pt.setSync(tc.A, tc.a, tc.sync)
			assert.Equal(t, tc.expectedOK, ok)

			if tc.expectedOK {
				e, ok := tc.pt.getEntry(tc.A, tc.a)
				assert.True(t, ok)
				assert.Equal(t, tc.sync, e.Sync)
			}
		})
	}
}

func TestParsingTable_String(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name               string
		pt                 *ParsingTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			pt:   pt[0],
			expectedSubstrings: []string{
				`┌──────────────┬────────────────────────────────────────────────────────────────────────────┐`,
				`│              │                                  Terminal                                  │`,
				`│ Non-Terminal ├───────────────┬───────────────┬───────────────┬────────┬──────────┬────────┤`,
				`│              │      "+"      │      "*"      │      "("      │  ")"   │   "id"   │   $    │`,
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

func TestParsingTable_Equal(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name          string
		pt            *ParsingTable
		rhs           *ParsingTable
		expectedEqual bool
	}{
		{
			name:          "Equal",
			pt:            pt[0],
			rhs:           pt[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			pt:            pt[1],
			rhs:           pt[2],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.pt.Equal(tc.rhs))
		})
	}
}

func TestParsingTable_Conflicts(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		pt                   *ParsingTable
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
				`multiple productions at M[S′, "e"]`,
				`S′ → "e" S`,
				`S′ → ε`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.pt.Conflicts()

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

func TestParsingTable_IsEmpty(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name            string
		pt              *ParsingTable
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
		pt             *ParsingTable
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
		pt                 *ParsingTable
		A                  grammar.NonTerminal
		a                  grammar.Terminal
		expectedOK         bool
		expectedProduction *grammar.Production
	}{
		{
			name:               "Empty",
			pt:                 pt[0],
			A:                  grammar.NonTerminal("E"),
			a:                  grammar.Terminal("+"),
			expectedOK:         false,
			expectedProduction: nil,
		},
		{
			name:       "OK",
			pt:         pt[0],
			A:          grammar.NonTerminal("E′"),
			a:          grammar.Terminal("+"),
			expectedOK: true,
			expectedProduction: &grammar.Production{
				Head: "E′",
				Body: grammar.String[grammar.Symbol]{grammar.Terminal("+"), grammar.NonTerminal("T"), grammar.NonTerminal("E′")},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prod, ok := tc.pt.GetProduction(tc.A, tc.a)

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, prod.Equal(tc.expectedProduction))
			} else {
				assert.False(t, ok)
				assert.Nil(t, prod)
			}
		})
	}
}

func TestParsingTableError(t *testing.T) {
	tests := []struct {
		name          string
		e             *parsingTableError
		expectedError string
	}{
		{
			name: "OK",
			e: &parsingTableError{
				NonTerminal: grammar.NonTerminal("decls"),
				Terminal:    grammar.Terminal("IDENT"),
				Productions: set.New(grammar.EqProduction,
					&grammar.Production{
						Head: "decls",
						Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")},
					},
					&grammar.Production{
						Head: "decls",
						Body: grammar.E,
					},
				),
			},
			expectedError: "multiple productions at M[decls, \"IDENT\"]:\n  decls → decls decl\n  decls → ε\n",
		},
	}

	for _, tc := range tests {
		assert.EqualError(t, tc.e, tc.expectedError)
	}
}

func TestTableStringer(t *testing.T) {
	tests := []struct {
		name               string
		ts                 *tableStringer[string, string]
		expectedSubstrings []string
	}{
		{
			name: "OK",
			ts: &tableStringer[string, string]{
				K1Title:  "None-Terminal",
				K1Values: []string{"A", "B", "C", "D"},
				K2Title:  "Input",
				K2Values: []string{"a", "b", "c", "d"},
				GetK1K2: func(k1 string, k2 string) string {
					return fmt.Sprintf("next(%s,%s)", k1, k2)
				},
			},
			expectedSubstrings: []string{
				`┌───────────────┬───────────────────────────────────────────────┐`,
				`│               │                     Input                     │`,
				`│ None-Terminal ├───────────┬───────────┬───────────┬───────────┤`,
				`│               │     a     │     b     │     c     │     d     │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       A       │ next(A,a) │ next(A,b) │ next(A,c) │ next(A,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       B       │ next(B,a) │ next(B,b) │ next(B,c) │ next(B,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       C       │ next(C,a) │ next(C,b) │ next(C,c) │ next(C,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       D       │ next(D,a) │ next(D,b) │ next(D,c) │ next(D,d) │`,
				`└───────────────┴───────────┴───────────┴───────────┴───────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.ts.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}
