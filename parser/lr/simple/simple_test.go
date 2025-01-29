package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
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
			G:                    grammars[0],
			expectedErrorStrings: nil,
		},
		{
			name: "None_SLR(1)_Grammar",
			L:    nil,
			G:    grammars[1],
			expectedErrorStrings: []string{
				`failed to construct SLR parsing table: 20 errors occurred:`,
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
