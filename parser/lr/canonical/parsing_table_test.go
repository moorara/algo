package canonical

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
			name:          "S→CC",
			G:             parsertest.Grammars[1],
			expectedTable: pt[0],
		},
		{
			name: "E→E+E",
			G:    parsertest.Grammars[4],
			expectedErrorStrings: []string{
				`8 errors occurred:`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "*"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[2, "+"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[3, "*"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[3, "+"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[4, "*"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[4, "+"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[5, "*"]`,
				`AMBIGUOUS Grammar: shift/reduce conflict in ACTION[5, "+"]`,
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
