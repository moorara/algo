package slr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    grammar.CFG
		expectedTable        *lr.ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:          "1st",
			G:             grammars[0],
			expectedTable: pt[0],
		},
		{
			name: "2nd",
			G:    grammars[1],
			expectedErrorStrings: []string{
				`20 errors occurred:`,
				`shift/reduce conflict at ACTION[2, "("]`,
				`shift/reduce conflict at ACTION[2, "IDENT"]`,
				`shift/reduce conflict at ACTION[2, "STRING"]`,
				`shift/reduce conflict at ACTION[2, "TOKEN"]`,
				`shift/reduce conflict at ACTION[2, "["]`,
				`shift/reduce conflict at ACTION[2, "{"]`,
				`shift/reduce conflict at ACTION[2, "{{"]`,
				`shift/reduce conflict at ACTION[2, "|"]`,
				`shift/reduce conflict at ACTION[7, "IDENT"]`,
				`shift/reduce conflict at ACTION[7, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "("]`,
				`shift/reduce conflict at ACTION[14, "IDENT"]`,
				`shift/reduce conflict at ACTION[14, "STRING"]`,
				`shift/reduce conflict at ACTION[14, "TOKEN"]`,
				`shift/reduce conflict at ACTION[14, "["]`,
				`shift/reduce conflict at ACTION[14, "{"]`,
				`shift/reduce conflict at ACTION[14, "{{"]`,
				`shift/reduce conflict at ACTION[14, "|"]`,
				`shift/reduce conflict at ACTION[19, "IDENT"]`,
				`shift/reduce conflict at ACTION[19, "TOKEN"]`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table := BuildParsingTable(tc.G)
			err := table.Error()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equals(tc.expectedTable))
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
