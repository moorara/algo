package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/lexer"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name                 string
		L                    lexer.Lexer
		G                    *grammar.CFG
		expectedErrorStrings []string
	}{
		{
			name:                 "Success",
			L:                    nil,
			G:                    parsertest.Grammars[3],
			expectedErrorStrings: nil,
		},
		{
			name: "None_SLR(1)_Grammar",
			L:    nil,
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
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "=" vs. "IDENT", "TOKEN"`,
				`              • "|" vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"`,
				`              • rhs = rhs rhs vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"`,
				`            Terminals/Productions listed earlier will have higher precedence.`,
				`            Terminals/Productions in the same line will have the same precedence.`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			p, err := New(tc.L, tc.G)

			if len(tc.expectedErrorStrings) == 0 {
				assert.NotNil(t, p)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, p)
				assert.Error(t, err)
				s := err.Error()
				for _, expectedErrorString := range tc.expectedErrorStrings {
					assert.Contains(t, s, expectedErrorString)
				}
			}
		})
	}
}
