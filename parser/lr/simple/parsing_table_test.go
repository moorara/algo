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
				`Error:      Ambiguous Grammar`,
				`Cause:      Multiple conflicts in the parsing table:`,
				`              1. Shift/Reduce conflict in ACTION[2, "("]`,
				`              2. Shift/Reduce conflict in ACTION[2, "IDENT"]`,
				`              3. Shift/Reduce conflict in ACTION[2, "STRING"]`,
				`              4. Shift/Reduce conflict in ACTION[2, "TOKEN"]`,
				`              5. Shift/Reduce conflict in ACTION[2, "["]`,
				`              6. Shift/Reduce conflict in ACTION[2, "{"]`,
				`              7. Shift/Reduce conflict in ACTION[2, "{{"]`,
				`              8. Shift/Reduce conflict in ACTION[2, "|"]`,
				`              9. Shift/Reduce conflict in ACTION[7, "IDENT"]`,
				`              10. Shift/Reduce conflict in ACTION[7, "TOKEN"]`,
				`              11. Shift/Reduce conflict in ACTION[14, "("]`,
				`              12. Shift/Reduce conflict in ACTION[14, "IDENT"]`,
				`              13. Shift/Reduce conflict in ACTION[14, "STRING"]`,
				`              14. Shift/Reduce conflict in ACTION[14, "TOKEN"]`,
				`              15. Shift/Reduce conflict in ACTION[14, "["]`,
				`              16. Shift/Reduce conflict in ACTION[14, "{"]`,
				`              17. Shift/Reduce conflict in ACTION[14, "{{"]`,
				`              18. Shift/Reduce conflict in ACTION[14, "|"]`,
				`              19. Shift/Reduce conflict in ACTION[19, "IDENT"]`,
				`              20. Shift/Reduce conflict in ACTION[19, "TOKEN"]`,
				`Resolution: Specify precedence for the following in the grammar directives:`,
				`              • "("`,
				`              • "="`,
				`              • "IDENT"`,
				`              • "STRING"`,
				`              • "TOKEN"`,
				`              • "["`,
				`              • "{"`,
				`              • "{{"`,
				`              • "|"`,
				`              • rhs = rhs rhs`,
				`            Terminals or Productions listed earlier in the directives will have higher precedence.`,
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
