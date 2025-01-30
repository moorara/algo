package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/parser/lr"
)

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    *grammar.CFG
		expectedTable        *lr.ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:          "E→E+T",
			G:             parsertest.Grammars[3],
			expectedTable: pt[0],
		},
		{
			name: "EBNF",
			G:    parsertest.Grammars[5],
			expectedErrorStrings: []string{
				`20 errors occurred:`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "("]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "IDENT"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "STRING"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "TOKEN"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "["]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "{"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "{{"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "|"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[7, "IDENT"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[7, "TOKEN"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "("]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "IDENT"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "STRING"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "TOKEN"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "["]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "{"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "{{"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[14, "|"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[19, "IDENT"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[19, "TOKEN"]`,
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
