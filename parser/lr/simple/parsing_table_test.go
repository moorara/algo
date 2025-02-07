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
		name          string
		G             *grammar.CFG
		precedences   lr.PrecedenceLevels
		expectedTable *lr.ParsingTable
		expectedError string
	}{
		{
			name:          "E→E+T",
			G:             parsertest.Grammars[3],
			precedences:   lr.PrecedenceLevels{},
			expectedTable: pt[0],
		},
		{
			name:        "E→E+E",
			G:           parsertest.Grammars[4],
			precedences: lr.PrecedenceLevels{},
			expectedError: `Error:      Ambiguous Grammar
Cause:      Multiple conflicts in the parsing table:
              1. Shift/Reduce conflict in ACTION[2, "*"]
              2. Shift/Reduce conflict in ACTION[2, "+"]
              3. Shift/Reduce conflict in ACTION[3, "*"]
              4. Shift/Reduce conflict in ACTION[3, "+"]
Resolution: Specify associativity and precedence for these Terminals/Productions:
              • "*" vs. "*", "+"
              • "+" vs. "*", "+"
            Terminals/Productions listed earlier will have higher precedence.
            Terminals/Productions in the same line will have the same precedence.
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table, err := BuildParsingTable(tc.G, tc.precedences)

			if len(tc.expectedError) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equal(tc.expectedTable))
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
