package canonical

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
			G:                    parsertest.Grammars[1],
			expectedErrorStrings: nil,
		},
		{
			name: "None_LR(1)_Grammar",
			L:    nil,
			G:    parsertest.Grammars[4],
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Multiple conflicts in the parsing table:`,
				`              1. Shift/Reduce conflict in ACTION[2, "*"]`,
				`              2. Shift/Reduce conflict in ACTION[2, "+"]`,
				`              3. Shift/Reduce conflict in ACTION[3, "*"]`,
				`              4. Shift/Reduce conflict in ACTION[3, "+"]`,
				`              5. Shift/Reduce conflict in ACTION[4, "*"]`,
				`              6. Shift/Reduce conflict in ACTION[4, "+"]`,
				`              7. Shift/Reduce conflict in ACTION[5, "*"]`,
				`              8. Shift/Reduce conflict in ACTION[5, "+"]`,
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "*" vs. "*", "+"`,
				`              • "+" vs. "*", "+"`,
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
