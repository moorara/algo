package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/internal/parsertest"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser/lr"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		L             lexer.Lexer
		G             *grammar.CFG
		precedences   lr.PrecedenceLevels
		expectedError string
	}{
		{
			name:          "E→E+T",
			L:             nil,
			G:             parsertest.Grammars[3],
			precedences:   lr.PrecedenceLevels{},
			expectedError: ``,
		},
		{
			name:        "EBNF",
			L:           nil,
			G:           parsertest.Grammars[5],
			precedences: lr.PrecedenceLevels{},
			expectedError: `Error:      Ambiguous Grammar
Cause:      Multiple conflicts in the parsing table:
              1. Shift/Reduce conflict in ACTION[3, "("]
              2. Shift/Reduce conflict in ACTION[3, "IDENT"]
              3. Shift/Reduce conflict in ACTION[3, "STRING"]
              4. Shift/Reduce conflict in ACTION[3, "TOKEN"]
              5. Shift/Reduce conflict in ACTION[3, "["]
              6. Shift/Reduce conflict in ACTION[3, "{"]
              7. Shift/Reduce conflict in ACTION[3, "{{"]
              8. Shift/Reduce conflict in ACTION[3, "|"]
              9. Shift/Reduce conflict in ACTION[17, "TOKEN"]
              10. Shift/Reduce conflict in ACTION[18, "TOKEN"]
              11. Shift/Reduce conflict in ACTION[19, "TOKEN"]
              12. Shift/Reduce conflict in ACTION[24, "("]
              13. Shift/Reduce conflict in ACTION[24, "IDENT"]
              14. Shift/Reduce conflict in ACTION[24, "STRING"]
              15. Shift/Reduce conflict in ACTION[24, "TOKEN"]
              16. Shift/Reduce conflict in ACTION[24, "["]
              17. Shift/Reduce conflict in ACTION[24, "{"]
              18. Shift/Reduce conflict in ACTION[24, "{{"]
              19. Shift/Reduce conflict in ACTION[25, "("]
              20. Shift/Reduce conflict in ACTION[25, "IDENT"]
              21. Shift/Reduce conflict in ACTION[25, "STRING"]
              22. Shift/Reduce conflict in ACTION[25, "TOKEN"]
              23. Shift/Reduce conflict in ACTION[25, "["]
              24. Shift/Reduce conflict in ACTION[25, "{"]
              25. Shift/Reduce conflict in ACTION[25, "{{"]
              26. Shift/Reduce conflict in ACTION[25, "|"]
Resolution: Specify associativity and precedence for these Terminals/Productions:
              • "@left" vs. "TOKEN"
              • "@none" vs. "TOKEN"
              • "@right" vs. "TOKEN"
              • "|" vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{"
              • "|" vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"
              • rhs = rhs rhs vs. "(", "IDENT", "STRING", "TOKEN", "[", "{", "{{", "|"
            Terminals/Productions listed earlier will have higher precedence.
            Terminals/Productions in the same line will have the same precedence.
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			p, err := New(tc.L, tc.G, tc.precedences)

			if len(tc.expectedError) == 0 {
				assert.NotNil(t, p)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, p)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}
