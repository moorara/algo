package lr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableStringer(t *testing.T) {
	tests := []struct {
		name               string
		ts                 *tableStringer[int, string, string]
		expectedSubstrings []string
	}{
		{
			name: "OK",
			ts: &tableStringer[int, string, string]{
				K1Title:  "STATE",
				K1Values: []int{0, 1, 2, 3, 4},
				K2Title:  "ACTION",
				K2Values: []string{"+", "*", "(", ")", "id", "$"},
				K3Title:  "GOTO",
				K3Values: []string{"E", "T", "F"},
				GetK1K2: func(k1 int, k2 string) string {
					return fmt.Sprintf("ACTION(%d,%s)", k1, k2)
				},
				GetK1K3: func(k1 int, k2 string) string {
					return fmt.Sprintf("GOTO(%d,%s)", k1, k2)
				},
			},
			expectedSubstrings: []string{
				`┌───────┬────────────────────────────────────────────────────────────────────────────────────┬───────────────────────────────────┐`,
				`│       │                                       ACTION                                       │               GOTO                │`,
				`│ STATE ├─────────────┬─────────────┬─────────────┬─────────────┬──────────────┬─────────────┼───────────┬───────────┬───────────┤`,
				`│       │      +      │      *      │      (      │      )      │      id      │      $      │     E     │     T     │     F     │`,
				`├───────┼─────────────┼─────────────┼─────────────┼─────────────┼──────────────┼─────────────┼───────────┼───────────┼───────────┤`,
				`│   0   │ ACTION(0,+) │ ACTION(0,*) │ ACTION(0,() │ ACTION(0,)) │ ACTION(0,id) │ ACTION(0,$) │ GOTO(0,E) │ GOTO(0,T) │ GOTO(0,F) │`,
				`├───────┼─────────────┼─────────────┼─────────────┼─────────────┼──────────────┼─────────────┼───────────┼───────────┼───────────┤`,
				`│   1   │ ACTION(1,+) │ ACTION(1,*) │ ACTION(1,() │ ACTION(1,)) │ ACTION(1,id) │ ACTION(1,$) │ GOTO(1,E) │ GOTO(1,T) │ GOTO(1,F) │`,
				`├───────┼─────────────┼─────────────┼─────────────┼─────────────┼──────────────┼─────────────┼───────────┼───────────┼───────────┤`,
				`│   2   │ ACTION(2,+) │ ACTION(2,*) │ ACTION(2,() │ ACTION(2,)) │ ACTION(2,id) │ ACTION(2,$) │ GOTO(2,E) │ GOTO(2,T) │ GOTO(2,F) │`,
				`├───────┼─────────────┼─────────────┼─────────────┼─────────────┼──────────────┼─────────────┼───────────┼───────────┼───────────┤`,
				`│   3   │ ACTION(3,+) │ ACTION(3,*) │ ACTION(3,() │ ACTION(3,)) │ ACTION(3,id) │ ACTION(3,$) │ GOTO(3,E) │ GOTO(3,T) │ GOTO(3,F) │`,
				`├───────┼─────────────┼─────────────┼─────────────┼─────────────┼──────────────┼─────────────┼───────────┼───────────┼───────────┤`,
				`│   4   │ ACTION(4,+) │ ACTION(4,*) │ ACTION(4,() │ ACTION(4,)) │ ACTION(4,id) │ ACTION(4,$) │ GOTO(4,E) │ GOTO(4,T) │ GOTO(4,F) │`,
				`└───────┴─────────────┴─────────────┴─────────────┴─────────────┴──────────────┴─────────────┴───────────┴───────────┴───────────┘`,
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
