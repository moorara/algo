package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/set"
)

func TestConflictError(t *testing.T) {
	tests := []struct {
		name                   string
		e                      *ConflictError
		expectedIsShiftReduce  bool
		expectedIsReduceReduce bool
		expectedErrorStrings   []string
	}{
		{
			name: "ShiftReduce",
			e: &ConflictError{
				State:    2,
				Terminal: "(",
				Actions:  set.New(eqAction, actions[1][3], actions[1][5]),
			},
			expectedIsShiftReduce:  true,
			expectedIsReduceReduce: false,
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Shift/Reduce conflict in ACTION[2, "("]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Shift the terminal "(", or`,
				`              2. Reduce by production rhs → rhs "|" rhs`,
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "|" vs. "("`,
				`            Terminals/Productions listed earlier will have higher precedence.`,
				`            Terminals/Productions in the same line will have the same precedence.`,
			},
		},
		{
			name: "ReduceReduce",
			e: &ConflictError{
				State:    40,
				Terminal: ";",
				Actions:  set.New(eqAction, actions[1][6], actions[1][7]),
			},
			expectedIsShiftReduce:  false,
			expectedIsReduceReduce: true,
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Reduce/Reduce conflict in ACTION[40, ";"]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Reduce by production nonterm → "IDENT", or`,
				`              2. Reduce by production rhs → "IDENT"`,
				`Resolution: Specify associativity for "IDENT".`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsShiftReduce, tc.e.IsShiftReduce())
			assert.Equal(t, tc.expectedIsReduceReduce, tc.e.IsReduceReduce())

			s := tc.e.Error()
			for _, expectedErrorString := range tc.expectedErrorStrings {
				assert.Contains(t, s, expectedErrorString)
			}
		})
	}
}

func TestAggregatedConflictError_ErrorOrNil(t *testing.T) {
	tests := []struct {
		name          string
		e             AggregatedConflictError
		expectedError error
	}{
		{
			name:          "Nil",
			e:             nil,
			expectedError: nil,
		},
		{
			name:          "Zero",
			e:             AggregatedConflictError{},
			expectedError: nil,
		},
		{
			name: "One",
			e: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
			},
			expectedError: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
			},
		},
		{
			name: "Multiple",
			e: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
				&ConflictError{State: 4, Terminal: "id"},
			},
			expectedError: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
				&ConflictError{State: 4, Terminal: "id"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, tc.e.ErrorOrNil())
		})
	}
}

func TestAggregatedConflictError_Error(t *testing.T) {
	tests := []struct {
		name                 string
		e                    AggregatedConflictError
		expectedErrorStrings []string
	}{
		{
			name:                 "Nil",
			e:                    nil,
			expectedErrorStrings: nil,
		},
		{
			name:                 "Zero",
			e:                    AggregatedConflictError{},
			expectedErrorStrings: nil,
		},
		{
			name: "One",
			e: AggregatedConflictError{
				&ConflictError{
					State:    2,
					Terminal: "(",
					Actions:  set.New(eqAction, actions[1][3], actions[1][5]),
				},
			},
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Shift/Reduce conflict in ACTION[2, "("]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Shift the terminal "(", or`,
				`              2. Reduce by production rhs → rhs "|" rhs`,
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "|" vs. "("`,
				`            Terminals/Productions listed earlier will have higher precedence.`,
				`            Terminals/Productions in the same line will have the same precedence.`,
			},
		},
		{
			name: "Multiple",
			e: AggregatedConflictError{
				&ConflictError{
					State:    2,
					Terminal: "(",
					Actions:  set.New(eqAction, actions[1][3], actions[1][5]),
				},
				&ConflictError{
					State:    19,
					Terminal: "(",
					Actions:  set.New(eqAction, actions[1][3], actions[1][4]),
				},
			},
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Multiple conflicts in the parsing table:`,
				`              1. Shift/Reduce conflict in ACTION[2, "("]`,
				`              2. Shift/Reduce conflict in ACTION[19, "("]`,
				`Resolution: Specify associativity and precedence for these Terminals/Productions:`,
				`              • "|" vs. "("`,
				`              • rhs = rhs rhs vs. "("`,
				`            Terminals/Productions listed earlier will have higher precedence.`,
				`            Terminals/Productions in the same line will have the same precedence.`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.e.Error()
			for _, expectedErrorString := range tc.expectedErrorStrings {
				assert.Contains(t, s, expectedErrorString)
			}
		})
	}
}

func TestAggregatedConflictError_Unwrap(t *testing.T) {
	tests := []struct {
		name           string
		e              AggregatedConflictError
		expectedErrors []error
	}{
		{
			name:           "Nil",
			e:              AggregatedConflictError{},
			expectedErrors: nil,
		},
		{
			name:           "Zero",
			e:              AggregatedConflictError{},
			expectedErrors: nil,
		},
		{
			name: "One",
			e: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
			},
			expectedErrors: []error{
				&ConflictError{State: 2, Terminal: "*"},
			},
		},
		{
			name: "Multiple",
			e: AggregatedConflictError{
				&ConflictError{State: 2, Terminal: "*"},
				&ConflictError{State: 4, Terminal: "id"},
			},
			expectedErrors: []error{
				&ConflictError{State: 2, Terminal: "*"},
				&ConflictError{State: 4, Terminal: "id"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := tc.e.Unwrap()
			assert.Equal(t, tc.expectedErrors, errs)
		})
	}
}

func TestPrecedenceHandleGroup(t *testing.T) {
	tests := []struct {
		name           string
		g              *precedenceHandleGroup
		expectedUnion  PrecedenceHandles
		expectedString string
	}{
		{
			name: "ShiftReduce",
			g: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			expectedUnion: NewPrecedenceHandles(
				handles[1][1],
				handles[1][2],
				handles[1][3],
				handles[1][4],
				handles[1][5],
			),
			expectedString: `"|" vs. "(", "[", "{", "{{"`,
		},
		{
			name: "ReduceReduce",
			g: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][9],
					handles[1][10],
				),
				shifts: NewPrecedenceHandles(),
			},
			expectedUnion: NewPrecedenceHandles(
				handles[1][9],
				handles[1][10],
			),
			expectedString: `rhs = rhs "|" rhs vs. rhs = rhs rhs`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, tc.g.Union().Equal(tc.expectedUnion))
			assert.Equal(t, tc.expectedString, tc.g.String())
		})
	}
}

func TestEqPrecedenceHandleGroup(t *testing.T) {
	tests := []struct {
		name          string
		lhs           *precedenceHandleGroup
		rhs           *precedenceHandleGroup
		expectedEqual bool
	}{
		{
			name: "Equal",
			lhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			rhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			expectedEqual: true,
		},
		{
			name: "NotEqual",
			lhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			rhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][9],
					handles[1][10],
				),
				shifts: NewPrecedenceHandles(),
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, eqPrecedenceHandleGroup(tc.lhs, tc.rhs))
		})
	}
}

func TestCmpPrecedenceHandleGroup(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *precedenceHandleGroup
		rhs             *precedenceHandleGroup
		expectedCompare int
	}{
		{
			name: "ByReduces",
			lhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][9],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			rhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][10],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			expectedCompare: 1,
		},
		{
			name: "ByShifts",
			lhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			rhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][6],
					handles[1][7],
					handles[1][8],
				),
			},
			expectedCompare: 1,
		},
		{
			name: "Equal",
			lhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			rhs: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			expectedCompare: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedCompare, cmpPrecedenceHandleGroup(tc.lhs, tc.rhs))
		})
	}
}

func TestHashPrecedenceHandleGroup(t *testing.T) {
	tests := []struct {
		name         string
		g            *precedenceHandleGroup
		expectedHash uint64
	}{
		{
			name: "Accept",
			g: &precedenceHandleGroup{
				reduces: NewPrecedenceHandles(
					handles[1][1],
				),
				shifts: NewPrecedenceHandles(
					handles[1][2],
					handles[1][3],
					handles[1][4],
					handles[1][5],
				),
			},
			expectedHash: 0x5333a4b38e5fa7eb,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hash := hashPrecedenceHandleGroup(tc.g)
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}
