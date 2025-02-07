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
			name:        "EBNF",
			G:           parsertest.Grammars[5],
			precedences: lr.PrecedenceLevels{},
			expectedError: `Error:      Ambiguous Grammar
Cause:      Multiple conflicts in the parsing table:
              1. Shift/Reduce conflict in ACTION[10, "("]
              2. Shift/Reduce conflict in ACTION[10, "IDENT"]
              3. Shift/Reduce conflict in ACTION[10, "STRING"]
              4. Shift/Reduce conflict in ACTION[10, "TOKEN"]
              5. Shift/Reduce conflict in ACTION[10, "["]
              6. Shift/Reduce conflict in ACTION[10, "{"]
              7. Shift/Reduce conflict in ACTION[10, "{{"]
              8. Shift/Reduce conflict in ACTION[10, "|"]
              9. Shift/Reduce conflict in ACTION[22, "TOKEN"]
              10. Shift/Reduce conflict in ACTION[23, "TOKEN"]
              11. Shift/Reduce conflict in ACTION[24, "TOKEN"]
              12. Shift/Reduce conflict in ACTION[29, "("]
              13. Shift/Reduce conflict in ACTION[29, "IDENT"]
              14. Shift/Reduce conflict in ACTION[29, "STRING"]
              15. Shift/Reduce conflict in ACTION[29, "TOKEN"]
              16. Shift/Reduce conflict in ACTION[29, "["]
              17. Shift/Reduce conflict in ACTION[29, "{"]
              18. Shift/Reduce conflict in ACTION[29, "{{"]
              19. Shift/Reduce conflict in ACTION[30, "("]
              20. Shift/Reduce conflict in ACTION[30, "IDENT"]
              21. Shift/Reduce conflict in ACTION[30, "STRING"]
              22. Shift/Reduce conflict in ACTION[30, "TOKEN"]
              23. Shift/Reduce conflict in ACTION[30, "["]
              24. Shift/Reduce conflict in ACTION[30, "{"]
              25. Shift/Reduce conflict in ACTION[30, "{{"]
              26. Shift/Reduce conflict in ACTION[30, "|"]
Resolution: Specify associativity and precedence for these Terminals/Productions:
              • "|" vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{"
              • "|" vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"
              • rhs = rhs rhs vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"
              • semi_opt = ε vs. "TOKEN"
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
