package lookahead

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
			name:          "S→L=R",
			L:             nil,
			G:             parsertest.Grammars[2],
			precedences:   lr.PrecedenceLevels{},
			expectedError: ``,
		},
		{
			name:          "E→E+T",
			L:             nil,
			G:             parsertest.Grammars[3],
			precedences:   lr.PrecedenceLevels{},
			expectedError: ``,
		},
		{
			name:        "E→E+E",
			L:           nil,
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
