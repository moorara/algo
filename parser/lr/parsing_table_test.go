package lr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

func getTestParsingTables() []*ParsingTable {
	pt0 := NewParsingTable(
		[]State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt0.AddACTION(0, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(0, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(1, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &Action{Type: ACCEPT})
	pt0.AddACTION(2, "+", &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(2, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(2, ")", &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(2, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][2]})
	pt0.AddACTION(3, "+", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, "*", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, ")", &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(3, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][4]})
	pt0.AddACTION(4, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(4, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(5, "+", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, "*", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, ")", &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(5, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][6]})
	pt0.AddACTION(6, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(6, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(7, "id", &Action{Type: SHIFT, State: 5})
	pt0.AddACTION(7, "(", &Action{Type: SHIFT, State: 4})
	pt0.AddACTION(8, "+", &Action{Type: SHIFT, State: 6})
	pt0.AddACTION(8, ")", &Action{Type: SHIFT, State: 11})
	pt0.AddACTION(9, "+", &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(9, "*", &Action{Type: SHIFT, State: 7})
	pt0.AddACTION(9, ")", &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(9, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][1]})
	pt0.AddACTION(10, "+", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, "*", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, ")", &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(10, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][3]})
	pt0.AddACTION(11, "+", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, "*", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, ")", &Action{Type: REDUCE, Production: prods[2][5]})
	pt0.AddACTION(11, grammar.Endmarker, &Action{Type: REDUCE, Production: prods[2][5]})

	pt0.SetGOTO(0, "E", 1)
	pt0.SetGOTO(0, "T", 2)
	pt0.SetGOTO(0, "F", 3)
	pt0.SetGOTO(4, "E", 8)
	pt0.SetGOTO(4, "T", 2)
	pt0.SetGOTO(4, "F", 3)
	pt0.SetGOTO(6, "T", 9)
	pt0.SetGOTO(6, "F", 3)
	pt0.SetGOTO(7, "F", 10)

	pt1 := NewParsingTable(
		[]State{0, 1, 2, 3, 4, 5, 6},
		[]grammar.Terminal{"a", "b", "c", "d", grammar.Endmarker},
		[]grammar.NonTerminal{"A", "B", "C", "D"},
	)

	pt1.AddACTION(0, "a", &Action{
		Type:  SHIFT,
		State: 5,
	})

	pt1.AddACTION(0, "a", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "A",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("a"), grammar.NonTerminal("A")},
		},
	})

	pt1.AddACTION(1, "b", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "B",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("b"), grammar.NonTerminal("B")},
		},
	})

	pt1.AddACTION(1, "b", &Action{
		Type: REDUCE,
		Production: &grammar.Production{
			Head: "C",
			Body: grammar.String[grammar.Symbol]{grammar.Terminal("c"), grammar.NonTerminal("C")},
		},
	})

	return []*ParsingTable{pt0, pt1}
}

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

func TestParsingTable_Error(t *testing.T) {
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
				`2 errors occurred:`,
				`shift/reduce conflict at ACTION[0, "a"]`,
				`SHIFT 5`,
				`REDUCE A → "a" A`,
				`reduce/reduce conflict at ACTION[1, "b"]`,
				`REDUCE B → "b" B`,
				`REDUCE C → "c" C`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.pt.Error()

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
		name           string
		pt             *ParsingTable
		s              State
		a              grammar.Terminal
		expectedAction *Action
		expectedError  string
	}{
		{
			name:           "NoAction",
			pt:             pt[0],
			s:              State(4),
			a:              grammar.Terminal("+"),
			expectedAction: &Action{Type: ERROR},
			expectedError:  "no action for ACTION[4, \"+\"]",
		},
		{
			name:           "Conflict",
			pt:             pt[1],
			s:              State(0),
			a:              grammar.Terminal("a"),
			expectedAction: &Action{Type: ERROR},
			expectedError:  "shift/reduce conflict at ACTION[0, \"a\"]\n  SHIFT 5\n  REDUCE A → \"a\" A\n",
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
			expectedError: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			action, err := tc.pt.ACTION(tc.s, tc.a)

			if len(tc.expectedError) == 0 {
				assert.True(t, action.Equal(tc.expectedAction))
				assert.NoError(t, err)
			} else {
				assert.True(t, action.Equal(tc.expectedAction))
				assert.EqualError(t, err, tc.expectedError)
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
			expectedError: "no state for GOTO[5, E]",
		},
		{
			name:          "NoNonTerminal",
			pt:            pt[0],
			s:             State(7),
			A:             grammar.NonTerminal("T"),
			expectedState: ErrState,
			expectedError: "no state for GOTO[7, T]",
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
				Type:   NO_ACTION,
				State:  State(7),
				Symbol: grammar.Terminal("+"),
			},
			expectedError: "no action for ACTION[7, \"+\"]",
		},
		{
			name: "NoState",
			e: &ParsingTableError{
				Type:   NO_GOTO,
				State:  State(5),
				Symbol: grammar.NonTerminal("T"),
			},
			expectedError: "no state for GOTO[5, T]",
		},
		{
			name: "Shift_Reduce_Conflict",
			e: &ParsingTableError{
				Type:    CONFLICT,
				State:   State(2),
				Symbol:  grammar.Terminal("*"),
				Actions: set.New(eqAction, actions[0], actions[2]),
			},
			expectedError: "shift/reduce conflict at ACTION[2, \"*\"]\n  SHIFT 5\n  REDUCE E → T\n",
		},
		{
			name: "Reduce_Reduce_Conflict",
			e: &ParsingTableError{
				Type:    CONFLICT,
				State:   State(4),
				Symbol:  grammar.Terminal("id"),
				Actions: set.New(eqAction, actions[2], actions[3]),
			},
			expectedError: "reduce/reduce conflict at ACTION[4, \"id\"]\n  REDUCE E → T\n  REDUCE F → \"id\"\n",
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
