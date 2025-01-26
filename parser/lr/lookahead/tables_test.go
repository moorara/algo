package lookahead

import (
	"testing"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/set"
	"github.com/stretchr/testify/assert"
)

func TestScopedItem(t *testing.T) {
	tests := []struct {
		name            string
		lhs             *scopedItem
		rhs             *scopedItem
		expectedEqual   bool
		expectedCompare int
	}{
		{
			name: "Equal",
			lhs: &scopedItem{
				ItemSet: lr.State(2),
				Item:    4,
			},
			rhs: &scopedItem{
				ItemSet: lr.State(2),
				Item:    4,
			},
			expectedEqual:   true,
			expectedCompare: 0,
		},
		{
			name: "FirstStateSmaller",
			lhs: &scopedItem{
				ItemSet: lr.State(1),
				Item:    4,
			},
			rhs: &scopedItem{
				ItemSet: lr.State(2),
				Item:    4,
			},
			expectedEqual:   false,
			expectedCompare: -1,
		},
		{
			name: "FirstStateLarger",
			lhs: &scopedItem{
				ItemSet: lr.State(3),
				Item:    4,
			},
			rhs: &scopedItem{
				ItemSet: lr.State(2),
				Item:    4,
			},
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
	t.Run("OK", func(t *testing.T) {
		pt := NewPropagationTable()

		assert.NotNil(t, pt)
		assert.NotNil(t, pt.table)
	})
}

func TestPropagationTable_Add(t *testing.T) {
	pt := NewPropagationTable()

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
			from: &scopedItem{ItemSet: lr.State(2), Item: 4},
			to: []*scopedItem{
				{ItemSet: lr.State(6), Item: 1},
			},
			expectedOK: true,
		},
		{
			name: "NotAdded",
			pt:   pt,
			from: &scopedItem{ItemSet: lr.State(2), Item: 4},
			to: []*scopedItem{
				{ItemSet: lr.State(6), Item: 1},
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
	pt := NewPropagationTable()
	pt.Add(
		&scopedItem{ItemSet: lr.State(2), Item: 4},
		&scopedItem{ItemSet: lr.State(6), Item: 1},
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
			from:        &scopedItem{ItemSet: lr.State(2), Item: 4},
			expectedSet: set.New(eqScopedItem, &scopedItem{ItemSet: lr.State(6), Item: 1}),
		},
		{
			name:        "NotExist",
			pt:          pt,
			from:        &scopedItem{ItemSet: lr.State(4), Item: 2},
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
	pt := NewPropagationTable()
	pt.Add(
		&scopedItem{ItemSet: lr.State(2), Item: 4},
		&scopedItem{ItemSet: lr.State(6), Item: 1},
	)
	pt.Add(
		&scopedItem{ItemSet: lr.State(4), Item: 2},
		&scopedItem{ItemSet: lr.State(7), Item: 1},
		&scopedItem{ItemSet: lr.State(8), Item: 1},
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

func TestNewLookaheadTable(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		lt := NewLookaheadTable()

		assert.NotNil(t, lt)
		assert.NotNil(t, lt.table)
	})
}

func TestLookaheadTable_Add(t *testing.T) {
	lt := NewLookaheadTable()

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
			item:       &scopedItem{ItemSet: lr.State(2), Item: 4},
			lookahead:  []grammar.Terminal{"$"},
			expectedOK: true,
		},
		{
			name:       "NotAdded",
			lt:         lt,
			item:       &scopedItem{ItemSet: lr.State(2), Item: 4},
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
	lt := NewLookaheadTable()
	lt.Add(&scopedItem{ItemSet: lr.State(2), Item: 4}, "$")

	tests := []struct {
		name        string
		lt          *lookaheadTable
		item        *scopedItem
		expectedSet set.Set[grammar.Terminal]
	}{
		{
			name:        "Exist",
			lt:          lt,
			item:        &scopedItem{ItemSet: lr.State(2), Item: 4},
			expectedSet: set.New(grammar.EqTerminal, "$"),
		},
		{
			name:        "NotExist",
			lt:          lt,
			item:        &scopedItem{ItemSet: lr.State(4), Item: 2},
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
	lt := NewLookaheadTable()
	lt.Add(&scopedItem{ItemSet: lr.State(2), Item: 4}, "$")
	lt.Add(&scopedItem{ItemSet: lr.State(4), Item: 2}, "$", "=")

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
