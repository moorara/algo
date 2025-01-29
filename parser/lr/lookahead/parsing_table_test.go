package lookahead

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/set"
)

var prods = [][]*grammar.Production{
	{
		{Head: "S′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("S")}},                                                 // S′ → S
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L"), grammar.Terminal("="), grammar.NonTerminal("R")}}, // S → L = R
		{Head: "S", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("R")}},                                                  // S → R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("*"), grammar.NonTerminal("R")}},                           // L → *R
		{Head: "L", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // L → id
		{Head: "R", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("L")}},                                                  // R → L
	},
	{
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("T")}}, // E → E + T
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T")}},                                                  // E → T
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("T"), grammar.Terminal("*"), grammar.NonTerminal("F")}}, // T → T * F
		{Head: "T", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("F")}},                                                  // T → F
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // F → ( E )
		{Head: "F", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // F → id
	},
	{
		{Head: "E′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E")}},                                                 // E′ → E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("+"), grammar.NonTerminal("E")}}, // E → E + E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("E"), grammar.Terminal("*"), grammar.NonTerminal("E")}}, // E → E * E
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("E"), grammar.Terminal(")")}},    // E → ( E )
		{Head: "E", Body: grammar.String[grammar.Symbol]{grammar.Terminal("id")}},                                                    // E → id
	},
	{
		{Head: "grammar′", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("grammar")}},                           // grammar′ → grammar
		{Head: "grammar", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("name"), grammar.NonTerminal("decls")}}, // grammar → name decls
		{Head: "name", Body: grammar.String[grammar.Symbol]{grammar.Terminal("grammar"), grammar.Terminal("IDENT")}},       // name → "grammar" IDENT
		{Head: "decls", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("decls"), grammar.NonTerminal("decl")}},   // decls → decls decl
		{Head: "decls", Body: grammar.E}, // decls → ε
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("token")}},                                                  // decl → token
		{Head: "decl", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rule")}},                                                   // decl → rule
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("STRING")}}, // token → TOKEN "=" STRING
		{Head: "token", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN"), grammar.Terminal("="), grammar.Terminal("REGEX")}},  // token → TOKEN "=" REGEX
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("="), grammar.NonTerminal("rhs")}}, // rule → lhs "=" rhs
		{Head: "rule", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("lhs"), grammar.Terminal("=")}},                             // rule → lhs "="
		{Head: "lhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // lhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.NonTerminal("rhs")}},                         // rhs → rhs rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("rhs"), grammar.Terminal("|"), grammar.NonTerminal("rhs")}},  // rhs → rhs "|" rhs
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("("), grammar.NonTerminal("rhs"), grammar.Terminal(")")}},       // rhs → "(" rhs ")"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("["), grammar.NonTerminal("rhs"), grammar.Terminal("]")}},       // rhs → "[" rhs "]"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{"), grammar.NonTerminal("rhs"), grammar.Terminal("}")}},       // rhs → "{" rhs "}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.Terminal("{{"), grammar.NonTerminal("rhs"), grammar.Terminal("}}")}},     // rhs → "{{" rhs "}}"
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("nonterm")}},                                                 // rhs → nonterm
		{Head: "rhs", Body: grammar.String[grammar.Symbol]{grammar.NonTerminal("term")}},                                                    // rhs → term
		{Head: "nonterm", Body: grammar.String[grammar.Symbol]{grammar.Terminal("IDENT")}},                                                  // nonterm → IDENT
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("TOKEN")}},                                                     // term → TOKEN
		{Head: "term", Body: grammar.String[grammar.Symbol]{grammar.Terminal("STRING")}},                                                    // term → STRING
	},
}

var grammars = []*grammar.CFG{
	grammar.NewCFG(
		[]grammar.Terminal{"=", "*", "id"},
		[]grammar.NonTerminal{"S", "L", "R"},
		prods[0][1:],
		"S",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E", "T", "F"},
		prods[1][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"+", "*", "(", ")", "id"},
		[]grammar.NonTerminal{"E"},
		prods[2][1:],
		"E",
	),
	grammar.NewCFG(
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX"},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
		prods[3][1:],
		"grammar",
	),
}

var sm = lr.StateMap{
	{
		&lr.Item0{Production: prods[0][0], Start: `S′`, Dot: 0}, // S′ → •S
	},
	{
		&lr.Item0{Production: prods[0][0], Start: `S′`, Dot: 1}, // S′ → S•
	},
	{
		&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 3}, // S → L "=" R•
	},
	{
		&lr.Item0{Production: prods[0][3], Start: `S′`, Dot: 2}, // L → "*" R•
	},
	{
		&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 2}, // S → L "="•R
	},
	{
		&lr.Item0{Production: prods[0][3], Start: `S′`, Dot: 1}, // L → "*"•R
	},
	{
		&lr.Item0{Production: prods[0][4], Start: `S′`, Dot: 1}, // L → "id"•
	},
	{
		&lr.Item0{Production: prods[0][5], Start: `S′`, Dot: 1}, // R → L•
		&lr.Item0{Production: prods[0][1], Start: `S′`, Dot: 1}, // S → L•"=" R
	},
	{
		&lr.Item0{Production: prods[0][5], Start: `S′`, Dot: 1}, // R → L•
	},
	{
		&lr.Item0{Production: prods[0][2], Start: `S′`, Dot: 1}, // S → R•
	},
}

func getTestParsingTables() []*lr.ParsingTable {
	pt0 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]grammar.Terminal{"=", "*", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"S", "L", "R"},
	)

	pt0.AddACTION(0, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt0.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][1]})
	pt0.AddACTION(3, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][3]})
	pt0.AddACTION(4, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(4, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(5, "*", &lr.Action{Type: lr.SHIFT, State: 5})
	pt0.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 6})
	pt0.AddACTION(6, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(6, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][4]})
	pt0.AddACTION(7, "=", &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(7, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(8, "=", &lr.Action{Type: lr.SHIFT, State: 4})
	pt0.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][5]})
	pt0.AddACTION(9, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[0][2]})

	pt0.SetGOTO(0, "S", 1)
	pt0.SetGOTO(0, "L", 8)
	pt0.SetGOTO(0, "R", 9)
	pt0.SetGOTO(4, "L", 7)
	pt0.SetGOTO(4, "R", 2)
	pt0.SetGOTO(5, "L", 7)
	pt0.SetGOTO(5, "R", 3)

	pt1 := lr.NewParsingTable(
		[]lr.State{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]grammar.Terminal{"+", "*", "(", ")", "id", grammar.Endmarker},
		[]grammar.NonTerminal{"E", "T", "F"},
	)

	pt1.AddACTION(0, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(0, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(1, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(1, grammar.Endmarker, &lr.Action{Type: lr.ACCEPT})
	pt1.AddACTION(2, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(2, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(2, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(2, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][1]})
	pt1.AddACTION(3, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(3, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][5]})
	pt1.AddACTION(4, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(4, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][3]})
	pt1.AddACTION(5, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(5, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(6, ")", &lr.Action{Type: lr.SHIFT, State: 3})
	pt1.AddACTION(6, "+", &lr.Action{Type: lr.SHIFT, State: 5})
	pt1.AddACTION(7, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(7, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(8, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(8, "*", &lr.Action{Type: lr.SHIFT, State: 7})
	pt1.AddACTION(8, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(8, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][2]})
	pt1.AddACTION(9, "(", &lr.Action{Type: lr.SHIFT, State: 9})
	pt1.AddACTION(9, "id", &lr.Action{Type: lr.SHIFT, State: 10})
	pt1.AddACTION(10, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(10, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][6]})
	pt1.AddACTION(11, ")", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, "*", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, "+", &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})
	pt1.AddACTION(11, grammar.Endmarker, &lr.Action{Type: lr.REDUCE, Production: prods[1][4]})

	pt1.SetGOTO(0, "E", 1)
	pt1.SetGOTO(0, "T", 8)
	pt1.SetGOTO(0, "F", 11)
	pt1.SetGOTO(5, "T", 2)
	pt1.SetGOTO(5, "F", 11)
	pt1.SetGOTO(7, "F", 4)
	pt1.SetGOTO(9, "E", 6)
	pt1.SetGOTO(9, "T", 8)
	pt1.SetGOTO(9, "F", 11)

	pt2 := lr.NewParsingTable(
		[]lr.State{},
		[]grammar.Terminal{"=", "|", "(", ")", "[", "]", "{", "}", "{{", "}}", "grammar", "IDENT", "TOKEN", "STRING", "REGEX", grammar.Endmarker},
		[]grammar.NonTerminal{"grammar", "name", "decls", "decl", "token", "rule", "lhs", "rhs", "nonterm", "term"},
	)

	return []*lr.ParsingTable{pt0, pt1, pt2}
}

func TestBuildParsingTable(t *testing.T) {
	pt := getTestParsingTables()

	tests := []struct {
		name                 string
		G                    *grammar.CFG
		expectedTable        *lr.ParsingTable
		expectedErrorStrings []string
	}{
		{
			name:          "1st",
			G:             grammars[0],
			expectedTable: pt[0],
		},
		{
			name:          "2nd",
			G:             grammars[1],
			expectedTable: pt[1],
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.NoError(t, tc.G.Verify())
			table, err := BuildParsingTable(tc.G)

			if len(tc.expectedErrorStrings) == 0 {
				assert.NoError(t, err)
				assert.True(t, table.Equal(tc.expectedTable))
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

func TestScopedItem(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *scopedItem
		rhs             *scopedItem
		expectedEqual   bool
		expectedCompare int
	}{
		{
			name:            "Equal",
			lhs:             &scopedItem{ItemSet: 2, Item: 4},
			rhs:             &scopedItem{ItemSet: 2, Item: 4},
			expectedEqual:   true,
			expectedCompare: 0,
		},
		{
			name:            "FirstStateSmaller",
			lhs:             &scopedItem{ItemSet: 1, Item: 4},
			rhs:             &scopedItem{ItemSet: 2, Item: 4},
			expectedEqual:   false,
			expectedCompare: -1,
		},
		{
			name:            "FirstStateLarger",
			lhs:             &scopedItem{ItemSet: 3, Item: 4},
			rhs:             &scopedItem{ItemSet: 2, Item: 4},
			expectedEqual:   false,
			expectedCompare: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedEqual, eqScopedItem(tc.lhs, tc.rhs))
			assert.Equal(t, tc.expectedCompare, cmpScopedItem(tc.lhs, tc.rhs))
		})
	}
}

func TestNewPropagationTable(t *testing.T) {
	tests := []struct {
		name string
		S    lr.StateMap
	}{
		{
			name: "OK",
			S:    lr.StateMap{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pt := newPropagationTable(tc.S)

			assert.NotNil(t, pt)
			assert.NotNil(t, pt.table)
		})
	}
}

func TestPropagationTable_Add(t *testing.T) {
	pt := newPropagationTable(nil)

	tests := []struct {
		name       string
		pt         *propagationTable
		from       *scopedItem
		to         []*scopedItem
		expectedOK bool
	}{
		{
			name: "Added",
			pt:   pt,
			from: &scopedItem{ItemSet: 2, Item: 4},
			to: []*scopedItem{
				{ItemSet: 6, Item: 1},
			},
			expectedOK: true,
		},
		{
			name: "NotAdded",
			pt:   pt,
			from: &scopedItem{ItemSet: 2, Item: 4},
			to: []*scopedItem{
				{ItemSet: 6, Item: 1},
			},
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.pt.Add(tc.from, tc.to...)

			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestPropagationTable_Get(t *testing.T) {
	pt := newPropagationTable(nil)
	pt.Add(
		&scopedItem{ItemSet: 2, Item: 4},
		&scopedItem{ItemSet: 6, Item: 1},
	)

	tests := []struct {
		name        string
		pt          *propagationTable
		from        *scopedItem
		expectedSet set.Set[*scopedItem]
	}{
		{
			name:        "Exist",
			pt:          pt,
			from:        &scopedItem{ItemSet: 2, Item: 4},
			expectedSet: set.New(eqScopedItem, &scopedItem{ItemSet: 6, Item: 1}),
		},
		{
			name:        "NotExist",
			pt:          pt,
			from:        &scopedItem{ItemSet: 4, Item: 2},
			expectedSet: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.pt.Get(tc.from)

			if tc.expectedSet == nil {
				assert.Nil(t, set)
			} else {
				assert.True(t, set.Equal(tc.expectedSet))
			}
		})
	}
}

func TestPropagationTable_All(t *testing.T) {
	pt := newPropagationTable(nil)
	pt.Add(
		&scopedItem{ItemSet: 2, Item: 4},
		&scopedItem{ItemSet: 6, Item: 1},
	)
	pt.Add(
		&scopedItem{ItemSet: 4, Item: 2},
		&scopedItem{ItemSet: 7, Item: 1},
		&scopedItem{ItemSet: 8, Item: 1},
	)

	tests := []struct {
		name string
		pt   *propagationTable
	}{
		{
			name: "OK",
			pt:   pt,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for item, set := range tc.pt.All() {
				assert.NotNil(t, item)
				assert.NotNil(t, set)
			}
		})
	}
}

func TestPropagationTable_String(t *testing.T) {
	pt := newPropagationTable(sm)
	pt.Add(&scopedItem{ItemSet: 0, Item: 0},
		&scopedItem{ItemSet: 1, Item: 0},
		&scopedItem{ItemSet: 5, Item: 0},
		&scopedItem{ItemSet: 6, Item: 0},
		&scopedItem{ItemSet: 7, Item: 0},
		&scopedItem{ItemSet: 7, Item: 1},
		&scopedItem{ItemSet: 9, Item: 0},
	)
	pt.Add(&scopedItem{ItemSet: 4, Item: 0},
		&scopedItem{ItemSet: 2, Item: 0},
		&scopedItem{ItemSet: 5, Item: 0},
		&scopedItem{ItemSet: 6, Item: 0},
		&scopedItem{ItemSet: 8, Item: 0},
	)
	pt.Add(&scopedItem{ItemSet: 5, Item: 0},
		&scopedItem{ItemSet: 3, Item: 0},
		&scopedItem{ItemSet: 5, Item: 0},
		&scopedItem{ItemSet: 6, Item: 0},
		&scopedItem{ItemSet: 8, Item: 0},
	)
	pt.Add(&scopedItem{ItemSet: 7, Item: 1},
		&scopedItem{ItemSet: 4, Item: 0},
	)

	tests := []struct {
		name               string
		pt                 *propagationTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			pt:   pt,
			expectedSubstrings: []string{
				`┌─────────────────┬──────────────────┐`,
				`│ FROM            │ TO               │`,
				`├─────────────────┼──────────────────┤`,
				`│ [0] S′ → •S     │ [1] S′ → S•      │`,
				`│                 │ [5] L → "*"•R    │`,
				`│                 │ [6] L → "id"•    │`,
				`│                 │ [7] R → L•       │`,
				`│                 │ [7] S → L•"=" R  │`,
				`│                 │ [9] S → R•       │`,
				`├─────────────────┼──────────────────┤`,
				`│ [4] S → L "="•R │ [2] S → L "=" R• │`,
				`│                 │ [5] L → "*"•R    │`,
				`│                 │ [6] L → "id"•    │`,
				`│                 │ [8] R → L•       │`,
				`├─────────────────┼──────────────────┤`,
				`│ [5] L → "*"•R   │ [3] L → "*" R•   │`,
				`│                 │ [5] L → "*"•R    │`,
				`│                 │ [6] L → "id"•    │`,
				`│                 │ [8] R → L•       │`,
				`├─────────────────┼──────────────────┤`,
				`│ [7] S → L•"=" R │ [4] S → L "="•R  │`,
				`└─────────────────┴──────────────────┘`,
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

func TestNewLookaheadTable(t *testing.T) {
	tests := []struct {
		name string
		S    lr.StateMap
	}{
		{
			name: "OK",
			S:    lr.StateMap{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lt := newLookaheadTable(tc.S)

			assert.NotNil(t, lt)
			assert.NotNil(t, lt.table)
		})
	}
}

func TestLookaheadTable_Add(t *testing.T) {
	lt := newLookaheadTable(nil)

	tests := []struct {
		name       string
		lt         *lookaheadTable
		item       *scopedItem
		lookahead  []grammar.Terminal
		expectedOK bool
	}{
		{
			name:       "Added",
			lt:         lt,
			item:       &scopedItem{ItemSet: 2, Item: 4},
			lookahead:  []grammar.Terminal{"$"},
			expectedOK: true,
		},
		{
			name:       "NotAdded",
			lt:         lt,
			item:       &scopedItem{ItemSet: 2, Item: 4},
			lookahead:  []grammar.Terminal{"$"},
			expectedOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok := tc.lt.Add(tc.item, tc.lookahead...)

			assert.Equal(t, tc.expectedOK, ok)
		})
	}
}

func TestLookaheadTable_Get(t *testing.T) {
	lt := newLookaheadTable(nil)
	lt.Add(&scopedItem{ItemSet: 2, Item: 4}, "$")

	tests := []struct {
		name        string
		lt          *lookaheadTable
		item        *scopedItem
		expectedSet set.Set[grammar.Terminal]
	}{
		{
			name:        "Exist",
			lt:          lt,
			item:        &scopedItem{ItemSet: 2, Item: 4},
			expectedSet: set.New(grammar.EqTerminal, "$"),
		},
		{
			name:        "NotExist",
			lt:          lt,
			item:        &scopedItem{ItemSet: 4, Item: 2},
			expectedSet: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.lt.Get(tc.item)

			if tc.expectedSet == nil {
				assert.Nil(t, set)
			} else {
				assert.True(t, set.Equal(tc.expectedSet))
			}
		})
	}
}

func TestLookaheadTable_All(t *testing.T) {
	lt := newLookaheadTable(nil)
	lt.Add(&scopedItem{ItemSet: 2, Item: 4}, "$")
	lt.Add(&scopedItem{ItemSet: 4, Item: 2}, "$", "=")

	tests := []struct {
		name string
		lt   *lookaheadTable
	}{
		{
			name: "OK",
			lt:   lt,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for item, set := range tc.lt.All() {
				assert.NotNil(t, item)
				assert.NotNil(t, set)
			}
		})
	}
}

func TestLookaheadTable_String(t *testing.T) {
	lt := newLookaheadTable(sm)
	lt.Add(&scopedItem{ItemSet: 0, Item: 0}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 1, Item: 0}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 2, Item: 0}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 3, Item: 0}, grammar.Endmarker, "=")
	lt.Add(&scopedItem{ItemSet: 4, Item: 0}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 5, Item: 0}, grammar.Endmarker, "=")
	lt.Add(&scopedItem{ItemSet: 6, Item: 0}, grammar.Endmarker, "=")
	lt.Add(&scopedItem{ItemSet: 7, Item: 0}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 7, Item: 1}, grammar.Endmarker)
	lt.Add(&scopedItem{ItemSet: 8, Item: 0}, grammar.Endmarker, "=")
	lt.Add(&scopedItem{ItemSet: 9, Item: 0}, grammar.Endmarker)

	tests := []struct {
		name               string
		lt                 *lookaheadTable
		expectedSubstrings []string
	}{
		{
			name: "OK",
			lt:   lt,
			expectedSubstrings: []string{
				`┌──────────────────┬────────────┐`,
				`│ ITEM             │ LOOKAHEADS │`,
				`├──────────────────┼────────────┤`,
				`│ [0] S′ → •S      │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [1] S′ → S•      │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [2] S → L "=" R• │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [3] L → "*" R•   │ $, "="     │`,
				`├──────────────────┼────────────┤`,
				`│ [4] S → L "="•R  │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [5] L → "*"•R    │ $, "="     │`,
				`├──────────────────┼────────────┤`,
				`│ [6] L → "id"•    │ $, "="     │`,
				`├──────────────────┼────────────┤`,
				`│ [7] R → L•       │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [7] S → L•"=" R  │ $          │`,
				`├──────────────────┼────────────┤`,
				`│ [8] R → L•       │ $, "="     │`,
				`├──────────────────┼────────────┤`,
				`│ [9] S → R•       │ $          │`,
				`└──────────────────┴────────────┘`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.lt.String()

			for _, expectedSubstring := range tc.expectedSubstrings {
				assert.Contains(t, s, expectedSubstring)
			}
		})
	}
}
