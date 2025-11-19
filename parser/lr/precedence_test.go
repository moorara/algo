package lr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
)

func TestAssociativity(t *testing.T) {
	tests := []struct {
		name           string
		a              Associativity
		expectedString string
	}{
		{name: "NONE", a: NONE, expectedString: "NONE"},
		{name: "NONE", a: LEFT, expectedString: "LEFT"},
		{name: "NONE", a: RIGHT, expectedString: "RIGHT"},
		{name: "Invalid", a: -1, expectedString: "Invalid Associativity(-1)"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.a.String())
		})
	}
}

func TestPrecedence(t *testing.T) {
	tests := []struct {
		name           string
		p              *Precedence
		expectedString string
		equal          *Precedence
		expectedEqual  bool
	}{
		{
			name: "OK",
			p: &Precedence{
				Order:         1,
				Associativity: LEFT,
			},
			expectedString: "1:LEFT",
			equal: &Precedence{
				Order:         1,
				Associativity: LEFT,
			},
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.equal))
		})
	}
}

func TestActionHandlePair(t *testing.T) {
	tests := []struct {
		name           string
		p              *ActionHandlePair
		expectedString string
		equal          *ActionHandlePair
		expectedEqual  bool
	}{
		{
			name: "OK",
			p: &ActionHandlePair{
				Action: actions[0][2], // SHIFT 5
				Handle: handles[0][2], // *
			},
			expectedString: `<SHIFT 5, "*">`,
			equal: &ActionHandlePair{
				Action: actions[1][5], // REDUCE rhs → rhs | rhs
				Handle: handles[1][1], // |
			},
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.equal))
		})
	}
}

func TestPrecedenceLevels_String(t *testing.T) {
	tests := []struct {
		name           string
		p              PrecedenceLevels
		expectedString string
	}{
		{
			name: "Equal",
			p:    precedences[0],
			expectedString: `LEFT "*", "/"
LEFT "+", "-"
NONE "<", ">"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
		})
	}
}

func TestPrecedenceLevels_Equal(t *testing.T) {
	tests := []struct {
		name          string
		p             PrecedenceLevels
		rhs           PrecedenceLevels
		expectedEqual bool
	}{
		{
			name:          "Equal",
			p:             precedences[0],
			rhs:           precedences[0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			p:             precedences[0],
			rhs:           precedences[1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestPrecedenceLevels_Verify(t *testing.T) {
	tests := []struct {
		name          string
		p             PrecedenceLevels
		expectedError string
	}{
		{
			name:          "Valid",
			p:             precedences[0],
			expectedError: "",
		},
		{
			name: "Invalid",
			p: PrecedenceLevels{
				{
					Associativity: RIGHT,
					Handles:       NewPrecedenceHandles(handles[0][1]),
				},
				{
					Associativity: LEFT,
					Handles:       NewPrecedenceHandles(handles[0][2], handles[0][3]),
				},
				{
					Associativity: LEFT,
					Handles:       NewPrecedenceHandles(handles[0][0], handles[0][1]),
				},
			},
			expectedError: "\"-\" appeared in more than one precedence level\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.p.Verify()

			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestPrecedenceLevels_Precedence(t *testing.T) {
	tests := []struct {
		name               string
		p                  PrecedenceLevels
		h                  *PrecedenceHandle
		expectedOK         bool
		expectedPrecedence *Precedence
	}{
		{
			name:       "Found",
			p:          precedences[0],
			h:          handles[0][0],
			expectedOK: true,
			expectedPrecedence: &Precedence{
				Order:         1,
				Associativity: LEFT,
			},
		},
		{
			name:               "NotFound",
			p:                  precedences[0],
			h:                  handles[1][0],
			expectedOK:         false,
			expectedPrecedence: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p, ok := tc.p.Precedence(tc.h)

			if tc.expectedOK {
				assert.True(t, ok)
				assert.True(t, p.Equal(tc.expectedPrecedence))
			} else {
				assert.False(t, ok)
				assert.Nil(t, p)
			}
		})
	}
}

func TestPrecedenceLevels_Compare(t *testing.T) {
	tests := []struct {
		name            string
		p               PrecedenceLevels
		lhs             *ActionHandlePair
		rhs             *ActionHandlePair
		expectedCompare int
		expectedError   string
	}{
		{
			name: "PairsEqual",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			rhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			expectedCompare: 0,
		},
		{
			name: "BothHandlesNotFound",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[1][3], // SHIFT 27
				Handle: handles[1][2], // (
			},
			rhs: &ActionHandlePair{
				Action: actions[1][5], // REDUCE rhs → rhs | rhs
				Handle: handles[1][1], // |
			},
			expectedError: `no associativity and precedence specified: "(", "|"`,
		},
		{
			name: "FirstHandleNotFound",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[1][3], // SHIFT 27
				Handle: handles[1][2], // (
			},
			rhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			expectedError: `no associativity and precedence specified: "("`,
		},
		{
			name: "SecondHandleNotFound",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][2], // SHIFT 5
				Handle: handles[0][2], // *
			},
			rhs: &ActionHandlePair{
				Action: actions[1][5], // REDUCE rhs → rhs | rhs
				Handle: handles[1][1], // |
			},
			expectedError: `no associativity and precedence specified: "|"`,
		},
		{
			name: "FirstHandlePrecedes",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][2], // SHIFT 5
				Handle: handles[0][2], // *
			},
			rhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			expectedCompare: 1,
		},
		{
			name: "SecondHandlePrecedes",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			rhs: &ActionHandlePair{
				Action: actions[0][2], // SHIFT 5
				Handle: handles[0][2], // *
			},
			expectedCompare: -1,
		},
		{
			name: "SameLevel_NoneAssociative_SameHandle",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: &Action{Type: SHIFT},             // SHIFT
				Handle: PrecedenceHandleForTerminal("<"), // <
			},
			rhs: &ActionHandlePair{
				Action: &Action{Type: REDUCE},            // REDUCE E → E < E
				Handle: PrecedenceHandleForTerminal("<"), // <
			},
			expectedError: `not associative: "<"`,
		},
		{
			name: "SameLevel_NoneAssociative_DistinctHandles",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: &Action{Type: SHIFT},             // SHIFT
				Handle: PrecedenceHandleForTerminal(">"), // >
			},
			rhs: &ActionHandlePair{
				Action: &Action{Type: REDUCE},            // REDUCE E → E < E
				Handle: PrecedenceHandleForTerminal("<"), // <
			},
			expectedError: `not associative: ">" and "<"`,
		},
		{
			name: "SameLevel_LeftAssociative_FirstPrecedes",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			rhs: &ActionHandlePair{
				Action: actions[0][3], // SHIFT 6
				Handle: handles[0][0], // +
			},
			expectedCompare: 1,
		},
		{
			name: "SameLevel_LeftAssociative_SecondPrecedes",
			p:    precedences[0],
			lhs: &ActionHandlePair{
				Action: actions[0][3], // SHIFT 6
				Handle: handles[0][0], // +
			},
			rhs: &ActionHandlePair{
				Action: actions[0][4], // REDUCE E → E + E
				Handle: handles[0][0], // +
			},
			expectedCompare: -1,
		},
		{
			name: "SameLevel_RightAssociative_FirstPrecedes",
			p:    precedences[1],
			lhs: &ActionHandlePair{
				Action: actions[1][2], // SHIFT 13
				Handle: handles[1][1], // |
			},
			rhs: &ActionHandlePair{
				Action: actions[1][5], // REDUCE rhs → rhs | rhs
				Handle: handles[1][1], // |
			},
			expectedCompare: 1,
		},
		{
			name: "SameLevel_RightAssociative_SecondPrecedes",
			p:    precedences[1],
			lhs: &ActionHandlePair{
				Action: actions[1][5], // REDUCE rhs → rhs | rhs
				Handle: handles[1][1], // |
			},
			rhs: &ActionHandlePair{
				Action: actions[1][2], // SHIFT 13
				Handle: handles[1][1], // |
			},
			expectedCompare: -1,
		},
		{
			name: "SameLevel_ReduceReduce_SameHandle",
			p:    precedences[1],
			lhs: &ActionHandlePair{
				Action: actions[1][6], // REDUCE rhs → IDENT
				Handle: handles[1][6], // IDENT
			},
			rhs: &ActionHandlePair{
				Action: actions[1][7], // REDUCE nonterm → IDENT
				Handle: handles[1][6], // IDENT
			},
			expectedError: `assign separate precedences: rhs → "IDENT" and nonterm → "IDENT"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmp, err := tc.p.Compare(tc.lhs, tc.rhs)

			if tc.expectedError == "" {
				assert.Equal(t, tc.expectedCompare, cmp)
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedError)
			}
		})
	}
}

func TestPrecedenceLevel_String(t *testing.T) {
	tests := []struct {
		name           string
		p              *PrecedenceLevel
		expectedString string
	}{
		{
			name:           "OK",
			p:              precedences[0][0],
			expectedString: `LEFT "*", "/"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedString, tc.p.String())
		})
	}
}

func TestPrecedenceLevel_Equal(t *testing.T) {
	tests := []struct {
		name          string
		p             *PrecedenceLevel
		rhs           *PrecedenceLevel
		expectedEqual bool
	}{
		{
			name:          "Equal",
			p:             precedences[0][0],
			rhs:           precedences[0][0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			p:             precedences[0][0],
			rhs:           precedences[0][1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.p.Equal(tc.rhs))
		})
	}
}

func TestNewPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name           string
		handles        []*PrecedenceHandle
		expectedString string
	}{
		{
			name:           "Zero",
			handles:        []*PrecedenceHandle{},
			expectedString: ``,
		},
		{
			name: "One",
			handles: []*PrecedenceHandle{
				handles[0][0],
			},
			expectedString: `"+"`,
		},
		{
			name: "Two",
			handles: []*PrecedenceHandle{
				handles[0][0],
				handles[0][1],
			},
			expectedString: `"+", "-"`,
		},
		{
			name: "Multiple",
			handles: []*PrecedenceHandle{
				handles[0][0],
				handles[0][1],
				handles[0][2],
				handles[0][3],
			},
			expectedString: `"*", "+", "-", "/"`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := NewPrecedenceHandles(tc.handles...)

			assert.Equal(t, tc.expectedString, h.String())
		})
	}
}

func TestEqPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name          string
		lhs           PrecedenceHandles
		rhs           PrecedenceHandles
		expectedEqual bool
	}{
		{
			name: "NotEqual",
			lhs: NewPrecedenceHandles(
				handles[0][0],
			),
			rhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			expectedEqual: false,
		},
		{
			name: "Equal",
			lhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			rhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			expectedEqual: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			eq := eqPrecedenceHandles(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expectedEqual, eq)
		})
	}
}

func TestCmpPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name            string
		lhs             PrecedenceHandles
		rhs             PrecedenceHandles
		expectedCompare int
	}{
		{
			name: "FirstShorter",
			lhs: NewPrecedenceHandles(
				handles[0][0],
			),
			rhs: NewPrecedenceHandles(
				handles[0][2],
				handles[0][3],
			),
			expectedCompare: -1,
		},
		{
			name: "FirstLonger",
			lhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			rhs: NewPrecedenceHandles(
				handles[0][2],
			),
			expectedCompare: 1,
		},
		{
			name: "EqualLength",
			lhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			rhs: NewPrecedenceHandles(
				handles[0][2],
				handles[0][3],
			),
			expectedCompare: 1,
		},
		{
			name: "Equal",
			lhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			rhs: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
			),
			expectedCompare: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmp := cmpPrecedenceHandles(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expectedCompare, cmp)
		})
	}
}

func TestHashPrecedenceHandles(t *testing.T) {
	tests := []struct {
		name         string
		h            PrecedenceHandles
		expectedHash uint64
	}{
		{
			name: "OK",
			h: NewPrecedenceHandles(
				handles[0][0],
				handles[0][1],
				handles[0][2],
				handles[0][3],
			),
			expectedHash: 0xbd8ef5321806dfcb,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hash := hashPrecedenceHandles(tc.h)
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}

func TestPrecedenceHandleForTerminal(t *testing.T) {
	tests := []struct {
		name string
		term grammar.Terminal
	}{
		{
			name: "OK",
			term: "IDENT",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := PrecedenceHandleForTerminal(tc.term)

			assert.NotNil(t, h)
			assert.Equal(t, tc.term, *h.Terminal)
		})
	}
}

func TestPrecedenceHandleForProduction(t *testing.T) {
	tests := []struct {
		name string
		prod *grammar.Production
	}{
		{
			name: "OK",
			prod: &grammar.Production{
				Head: "rhs",
				Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := PrecedenceHandleForProduction(tc.prod)

			assert.NotNil(t, h)
			assert.True(t, h.Production.Equal(tc.prod))
		})
	}
}

func TestPrecedenceHandle(t *testing.T) {
	tests := []struct {
		name                 string
		h                    *PrecedenceHandle
		expectedIsTerminal   bool
		expectedIsProduction bool
		expectedString       string
	}{
		{
			name:                 "Terminal",
			h:                    handles[1][1],
			expectedIsTerminal:   true,
			expectedIsProduction: false,
			expectedString:       `"|"`,
		},
		{
			name:                 "Production",
			h:                    handles[1][9],
			expectedIsTerminal:   false,
			expectedIsProduction: true,
			expectedString:       `rhs = rhs rhs`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedIsTerminal, tc.h.IsTerminal())
			assert.Equal(t, tc.expectedIsProduction, tc.h.IsProduction())
			assert.Equal(t, tc.expectedString, tc.h.String())
		})
	}
}

func TestPrecedenceHandle_Equal(t *testing.T) {
	tests := []struct {
		name          string
		h             *PrecedenceHandle
		rhs           *PrecedenceHandle
		expectedEqual bool
	}{
		{
			name:          "Terminal_Equal",
			h:             handles[0][0],
			rhs:           handles[0][0],
			expectedEqual: true,
		},
		{
			name:          "Terminal_NotEqual",
			h:             handles[0][0],
			rhs:           handles[0][1],
			expectedEqual: false,
		},
		{
			name:          "Production_Equal",
			h:             handles[0][4],
			rhs:           handles[0][4],
			expectedEqual: true,
		},
		{
			name:          "Production_NotEqual",
			h:             handles[0][4],
			rhs:           handles[0][5],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, tc.h.Equal(tc.rhs))
		})
	}
}

func TestEqPrecedenceHandle(t *testing.T) {
	tests := []struct {
		name          string
		lhs           *PrecedenceHandle
		rhs           *PrecedenceHandle
		expectedEqual bool
	}{
		{
			name:          "Equal",
			lhs:           handles[0][0],
			rhs:           handles[0][0],
			expectedEqual: true,
		},
		{
			name:          "NotEqual",
			lhs:           handles[0][0],
			rhs:           handles[0][1],
			expectedEqual: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			eq := eqPrecedenceHandle(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expectedEqual, eq)
		})
	}
}

func TestCmpPrecedenceHandle(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *PrecedenceHandle
		rhs             *PrecedenceHandle
		expectedCompare int
	}{
		{
			name:            "BothTerminal",
			lhs:             handles[0][0],
			rhs:             handles[0][1],
			expectedCompare: -1,
		},
		{
			name:            "BothProduction",
			lhs:             handles[0][4],
			rhs:             handles[0][5],
			expectedCompare: 1,
		},
		{
			name:            "TerminalAndProduction",
			lhs:             handles[0][2],
			rhs:             handles[0][4],
			expectedCompare: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmp := cmpPrecedenceHandle(tc.lhs, tc.rhs)
			assert.Equal(t, tc.expectedCompare, cmp)
		})
	}
}

func TestHashPrecedenceHandle(t *testing.T) {
	tests := []struct {
		name         string
		h            *PrecedenceHandle
		expectedHash uint64
	}{
		{
			name:         "Terminal",
			h:            handles[0][0],
			expectedHash: 0xaf63bd4c8601b7f4,
		},
		{
			name:         "Production",
			h:            handles[0][4],
			expectedHash: 0xb17a6a839462b089,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			hash := hashPrecedenceHandle(tc.h)
			assert.Equal(t, tc.expectedHash, hash)
		})
	}
}
