package canonical

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
			name: "None_LR(1)_Grammar",
			L:    nil,
			G:    grammars[1],
			expectedErrorStrings: []string{
				`failed to construct Canonical LR parsing table: 8 errors occurred:`,
				`shift/reduce conflict at ACTION[2, "*"]`,
				`shift/reduce conflict at ACTION[2, "+"]`,
				`shift/reduce conflict at ACTION[3, "*"]`,
				`shift/reduce conflict at ACTION[3, "+"]`,
				`shift/reduce conflict at ACTION[4, "*"]`,
				`shift/reduce conflict at ACTION[4, "+"]`,
				`shift/reduce conflict at ACTION[5, "*"]`,
				`shift/reduce conflict at ACTION[5, "+"]`,
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
