package predictive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableStringer(t *testing.T) {
	tests := []struct {
		name               string
		ts                 *tableStringer[string, string]
		expectedSubstrings []string
	}{
		{
			name: "OK",
			ts: &tableStringer[string, string]{
				K1Title:  "None-Terminal",
				K1Values: []string{"A", "B", "C", "D"},
				K2Title:  "Input",
				K2Values: []string{"a", "b", "c", "d"},
				GetK1K2: func(k1 string, k2 string) string {
					return fmt.Sprintf("next(%s,%s)", k1, k2)
				},
			},
			expectedSubstrings: []string{
				`┌───────────────┬───────────────────────────────────────────────┐`,
				`│               │                     Input                     │`,
				`│ None-Terminal ├───────────┬───────────┬───────────┬───────────┤`,
				`│               │     a     │     b     │     c     │     d     │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       A       │ next(A,a) │ next(A,b) │ next(A,c) │ next(A,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       B       │ next(B,a) │ next(B,b) │ next(B,c) │ next(B,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       C       │ next(C,a) │ next(C,b) │ next(C,c) │ next(C,d) │`,
				`├───────────────┼───────────┼───────────┼───────────┼───────────┤`,
				`│       D       │ next(D,a) │ next(D,b) │ next(D,c) │ next(D,d) │`,
				`└───────────────┴───────────┴───────────┴───────────┴───────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.ts.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}
