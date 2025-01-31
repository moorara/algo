package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

func TestConflictError(t *testing.T) {
	star := grammar.Terminal("*")
	id := grammar.Terminal("id")

	tests := []struct {
		name                   string
		e                      *ConflictError
		expectedIsShiftReduce  bool
		expectedIsReduceReduce bool
		expectedHandles        set.Set[*PrecedenceHandle]
		expectedErrorStrings   []string
	}{
		{
			name: "ShiftReduce",
			e: &ConflictError{
				State:    2,
				Terminal: "*",
				Actions:  set.New(eqAction, actions[0], actions[2]),
			},
			expectedIsShiftReduce:  true,
			expectedIsReduceReduce: false,
			expectedHandles: set.New(eqPrecedenceHandle,
				&PrecedenceHandle{
					Terminal: &star,
				},
				&PrecedenceHandle{
					Production: &grammar.Production{
						Head: "E",
						Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")},
					},
				},
			),
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Shift/Reduce conflict in ACTION[2, "*"]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Shift the terminal "*", or`,
				`              2. Reduce by production E → T`,
				`Resolution: Specify precedence for the following in the grammar directives:`,
				`              • "*"`,
				`              • E = T`,
				`            Terminals or Productions listed earlier in the directives will have higher precedence.`,
			},
		},
		{
			name: "ReduceReduce",
			e: &ConflictError{
				State:    4,
				Terminal: "id",
				Actions:  set.New(eqAction, actions[2], actions[3]),
			},
			expectedIsShiftReduce:  false,
			expectedIsReduceReduce: true,
			expectedHandles: set.New(eqPrecedenceHandle,
				&PrecedenceHandle{
					Terminal: &id,
				},
				&PrecedenceHandle{
					Production: &grammar.Production{
						Head: "E",
						Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")},
					},
				},
			),
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Reduce/Reduce conflict in ACTION[4, "id"]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Reduce by production E → T, or`,
				`              2. Reduce by production F → "id"`,
				`Resolution: Specify precedence for the following in the grammar directives:`,
				`              • "id"`,
				`              • E = T`,
				`            Terminals or Productions listed earlier in the directives will have higher precedence.`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsShiftReduce, tc.e.IsShiftReduce())
			assert.Equal(t, tc.expectedIsReduceReduce, tc.e.IsReduceReduce())

			handles := tc.e.Handles()
			assert.True(t, handles.Equal(tc.expectedHandles))

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
					Terminal: "*",
					Actions:  set.New(eqAction, actions[0], actions[2]),
				},
			},
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Shift/Reduce conflict in ACTION[2, "*"]`,
				`Context:    The parser cannot decide whether to`,
				`              1. Shift the terminal "*", or`,
				`              2. Reduce by production E → T`,
				`Resolution: Specify precedence for the following in the grammar directives:`,
				`              • "*"`,
				`              • E = T`,
				`            Terminals or Productions listed earlier in the directives will have higher precedence.`,
			},
		},
		{
			name: "Multiple",
			e: AggregatedConflictError{
				&ConflictError{
					State:    2,
					Terminal: "*",
					Actions:  set.New(eqAction, actions[0], actions[2]),
				},
				&ConflictError{
					State:    4,
					Terminal: "id",
					Actions:  set.New(eqAction, actions[2], actions[3]),
				},
			},
			expectedErrorStrings: []string{
				`Error:      Ambiguous Grammar`,
				`Cause:      Multiple conflicts in the parsing table`,
				`              1. Shift/Reduce conflict in ACTION[2, "*"]`,
				`              2. Reduce/Reduce conflict in ACTION[4, "id"]`,
				`Resolution: Specify precedence for the following in the grammar directives:`,
				`              • "*"`,
				`              • "id"`,
				`              • E = T`,
				`            Terminals or Productions listed earlier in the directives will have higher precedence.`,
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
