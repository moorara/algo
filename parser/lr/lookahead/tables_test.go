package lookahead

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/set"
	"github.com/stretchr/testify/assert"
)

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
