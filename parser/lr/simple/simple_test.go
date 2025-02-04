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
              1. Shift/Reduce conflict in ACTION[2, "("]
              2. Shift/Reduce conflict in ACTION[2, "IDENT"]
              3. Shift/Reduce conflict in ACTION[2, "STRING"]
              4. Shift/Reduce conflict in ACTION[2, "TOKEN"]
              5. Shift/Reduce conflict in ACTION[2, "["]
              6. Shift/Reduce conflict in ACTION[2, "{"]
              7. Shift/Reduce conflict in ACTION[2, "{{"]
              8. Shift/Reduce conflict in ACTION[2, "|"]
              9. Shift/Reduce conflict in ACTION[7, "IDENT"]
              10. Shift/Reduce conflict in ACTION[7, "STRING"]
              11. Shift/Reduce conflict in ACTION[7, "TOKEN"]
              12. Shift/Reduce conflict in ACTION[12, "IDENT"]
              13. Shift/Reduce conflict in ACTION[12, "TOKEN"]
              14. Shift/Reduce conflict in ACTION[13, "IDENT"]
              15. Shift/Reduce conflict in ACTION[13, "TOKEN"]
              16. Shift/Reduce conflict in ACTION[14, "IDENT"]
              17. Shift/Reduce conflict in ACTION[14, "TOKEN"]
              18. Shift/Reduce conflict in ACTION[20, "("]
              19. Shift/Reduce conflict in ACTION[20, "IDENT"]
              20. Shift/Reduce conflict in ACTION[20, "STRING"]
              21. Shift/Reduce conflict in ACTION[20, "TOKEN"]
              22. Shift/Reduce conflict in ACTION[20, "["]
              23. Shift/Reduce conflict in ACTION[20, "{"]
              24. Shift/Reduce conflict in ACTION[20, "{{"]
              25. Shift/Reduce conflict in ACTION[20, "|"]
              26. Shift/Reduce conflict in ACTION[25, "IDENT"]
              27. Shift/Reduce conflict in ACTION[25, "STRING"]
              28. Shift/Reduce conflict in ACTION[25, "TOKEN"]
Resolution: Specify associativity and precedence for these Terminals/Productions:
              • "=" vs. "IDENT", "STRING", "TOKEN"
              • "@left" vs. "IDENT", "TOKEN"
              • "@none" vs. "IDENT", "TOKEN"
              • "@right" vs. "IDENT", "TOKEN"
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
