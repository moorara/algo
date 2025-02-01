package lr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

func TestNewParsingTable(t *testing.T) {
	tests := []struct {
		name         string
		states       []State
		terminals    []grammar.Terminal
		nonTerminals []grammar.NonTerminal
	}{
		{
			name:         "OK",
			states:       []State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			terminals:    []grammar.Terminal{"+", "-", "*", "/", "(", ")", "id"},
			nonTerminals: []grammar.NonTerminal{"E", "T", "F"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pt := NewParsingTable(tc.states, tc.terminals, tc.nonTerminals)

			assert.NotNil(t, pt)
			assert.NotNil(t, pt.actions)
			assert.NotNil(t, pt.gotos)
		})
	}
}

func TestParsingTable_String(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name               string
		pt                 *ParsingTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			pt:   pt[0],
			expectedSubstrings: []string{
				`┌───────┬───────────────────────────────────────────────────────────────────────────────────────────────────────────────┬────────────┐`,
				`│       │                                                    ACTION                                                     │    GOTO    │`,
				`│ STATE ├──────────────────────┬──────────────────────┬─────────┬──────────────────────┬─────────┬──────────────────────┼───┬───┬────┤`,
				`│       │         "+"          │         "*"          │   "("   │         ")"          │  "id"   │          $           │ E │ T │ F  │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   0   │                      │                      │ SHIFT 4 │                      │ SHIFT 5 │                      │ 1 │ 2 │ 3  │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   1   │       SHIFT 6        │                      │         │                      │         │        ACCEPT        │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   2   │     REDUCE E → T     │       SHIFT 7        │         │     REDUCE E → T     │         │     REDUCE E → T     │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   3   │     REDUCE T → F     │     REDUCE T → F     │         │     REDUCE T → F     │         │     REDUCE T → F     │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   4   │                      │                      │ SHIFT 4 │                      │ SHIFT 5 │                      │ 8 │ 2 │ 3  │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   5   │   REDUCE F → "id"    │   REDUCE F → "id"    │         │   REDUCE F → "id"    │         │   REDUCE F → "id"    │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   6   │                      │                      │ SHIFT 4 │                      │ SHIFT 5 │                      │   │ 9 │ 3  │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   7   │                      │                      │ SHIFT 4 │                      │ SHIFT 5 │                      │   │   │ 10 │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   8   │       SHIFT 6        │                      │         │       SHIFT 11       │         │                      │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│   9   │  REDUCE E → E "+" T  │       SHIFT 7        │         │  REDUCE E → E "+" T  │         │  REDUCE E → E "+" T  │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│  10   │  REDUCE T → T "*" F  │  REDUCE T → T "*" F  │         │  REDUCE T → T "*" F  │         │  REDUCE T → T "*" F  │   │   │    │`,
				`├───────┼──────────────────────┼──────────────────────┼─────────┼──────────────────────┼─────────┼──────────────────────┼───┼───┼────┤`,
				`│  11   │ REDUCE F → "(" E ")" │ REDUCE F → "(" E ")" │         │ REDUCE F → "(" E ")" │         │ REDUCE F → "(" E ")" │   │   │    │`,
				`└───────┴──────────────────────┴──────────────────────┴─────────┴──────────────────────┴─────────┴──────────────────────┴───┴───┴────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.pt.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}

func TestParsingTable_Equal(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name          string
		pt            *ParsingTable
		rhs           *ParsingTable
		expectedEqual bool
	}{
		{
			name:          "Equal",
			pt:            pt[0],
			rhs:           pt[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			pt:            pt[0],
			rhs:           pt[1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.pt.Equal(tc.rhs))
		})
	}
}

func TestParsingTable_AddACTION(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name       string
		pt         *ParsingTable
		s          State
		a          grammar.Terminal
		action     *Action
		expectedOK bool
	}{
		{
			name: "OK",
			pt:   pt[1],
			s:    State(2),
			a:    grammar.Terminal("c"),
			action: &Action{
				Type:  SHIFT,
				State: 4,
			},
			expectedOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.pt.AddACTION(tc.s, tc.a, tc.action)
			assert.Equal(t, tc.expectedOK, ok)

			if tc.expectedOK {
				actions, ok := tc.pt.getActions(tc.s, tc.a)
				assert.True(t, ok)
				assert.True(t, actions.Contains(tc.action))
			}
		})
	}
}

func TestParsingTable_SetGOTO(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name string
		pt   *ParsingTable
		s    State
		A    grammar.NonTerminal
		next State
	}{
		{
			name: "OK",
			pt:   pt[1],
			s:    State(3),
			A:    grammar.NonTerminal("D"),
			next: State(6),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.pt.SetGOTO(tc.s, tc.A, tc.next)

			row, ok := tc.pt.gotos.Get(tc.s)
			assert.True(t, ok)
			state, ok := row.Get(tc.A)
			assert.True(t, ok)
			assert.Equal(t, tc.next, state)
		})
	}
}

func TestParsingTable_Conflicts(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		pt                   *ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:                 "NoError",
			pt:                   pt[0],
			expectedErrorStrings: nil,
		},
		{
			name: "Error",
			pt:   pt[1],
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Multiple conflicts in the parsing table:`,
				`              1. Shift/Reduce conflict in ACTION[2, "+"]`,
				`              2. Shift/Reduce conflict in ACTION[2, "*"]`,
				`              3. Shift/Reduce conflict in ACTION[3, "+"]`,
				`              4. Shift/Reduce conflict in ACTION[3, "*"]`,
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
			err := tc.pt.Conflicts()

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
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

func TestParsingTable_ACTION(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		pt                   *ParsingTable
		s                    State
		a                    grammar.Terminal
		expectedAction       *Action
		expectedErrorStrings []string
	}{
		{
			name:           "NoAction",
			pt:             pt[0],
			s:              State(4),
			a:              grammar.Terminal("+"),
			expectedAction: &Action{Type: ERROR},
			expectedErrorStrings: []string{
				`no action exists in the parsing table for ACTION[4, "+"]`,
			},
		},
		{
			name:           "Conflict",
			pt:             pt[1],
			s:              State(2),
			a:              grammar.Terminal("+"),
			expectedAction: &Action{Type: ERROR},
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Shift/Reduce conflict in ACTION[2, "+"]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Shift the terminal "+", or`,
				`              2. Reduce by production E → E "*" E`,
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "*" vs. "+"`,
				`            Terminals/Productions listed earlier will have higher precedence.`,
				`            Terminals/Productions in the same line will have the same precedence.`,
			},
		},
		{
			name: "Success",
			pt:   pt[0],
			s:    State(4),
			a:    grammar.Terminal("id"),
			expectedAction: &Action{
				Type:  SHIFT,
				State: 5,
			},
			expectedErrorStrings: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			action, err := tc.pt.ACTION(tc.s, tc.a)
			assert.True(t, action.Equal(tc.expectedAction))

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
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

func TestParsingTable_GOTO(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name          string
		pt            *ParsingTable
		s             State
		A             grammar.NonTerminal
		expectedState State
		expectedError string
	}{
		{
			name:          "NoState",
			pt:            pt[0],
			s:             State(5),
			A:             grammar.NonTerminal("E"),
			expectedState: ErrState,
			expectedError: "no state exists in the parsing table for GOTO[5, E]",
		},
		{
			name:          "NoNonTerminal",
			pt:            pt[0],
			s:             State(7),
			A:             grammar.NonTerminal("T"),
			expectedState: ErrState,
			expectedError: "no state exists in the parsing table for GOTO[7, T]",
		},
		{
			name:          "Success",
			pt:            pt[0],
			s:             State(7),
			A:             grammar.NonTerminal("F"),
			expectedState: State(10),
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			state, err := tc.pt.GOTO(tc.s, tc.A)

			if len(tc.expectedError) == 0 {
				assert.Equal(t, tc.expectedState, state)
				assert.NoError(t, err)
			} else {
				assert.Equal(t, ErrState, state)
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestParsingTableError(t *testing.T) {
	tests := []struct {
		name          string
		e             *ParsingTableError
		expectedError string
	}{
		{
			name: "NoAction",
			e: &ParsingTableError{
				Type:   MISSING_ACTION,
				State:  State(7),
				Symbol: grammar.Terminal("+"),
			},
			expectedError: "no action exists in the parsing table for ACTION[7, \"+\"]",
		},
		{
			name: "NoState",
			e: &ParsingTableError{
				Type:   MISSING_GOTO,
				State:  State(5),
				Symbol: grammar.NonTerminal("T"),
			},
			expectedError: "no state exists in the parsing table for GOTO[5, T]",
		},
		{
			name: "Invalid",
			e: &ParsingTableError{
				Type: ParsingTableErrorType(0),
			},
			expectedError: "invalid error: 0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.EqualError(t, tc.e, tc.expectedError)
		})
	}
}

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
