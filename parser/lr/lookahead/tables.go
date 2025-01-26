package lookahead

import (
	"iter"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

// scopedItem represents an individual item within an item set.
// The item set is represented by its state,
// and the individual item is identified by its index within the sorted item set.
type scopedItem struct {
	ItemSet lr.State
	Item    int
}

func eqScopedItem(lhs, rhs *scopedItem) bool {
	return lhs.ItemSet == rhs.ItemSet && lhs.Item == rhs.Item
}

func cmpScopedItem(lhs, rhs *scopedItem) int {
	if lhs.ItemSet < rhs.ItemSet {
		return -1
	} else if lhs.ItemSet > rhs.ItemSet {
		return 1
	}

	return lhs.Item - rhs.Item
}

// propagationTable keeps track of which scoped items propagate their lookaheads to which other scoped items.
type propagationTable struct {
	table symboltable.SymbolTable[*scopedItem, set.Set[*scopedItem]]
}

func NewPropagationTable() *propagationTable {
	return &propagationTable{
		table: symboltable.NewRedBlack(
			cmpScopedItem,
			func(lhs, rhs set.Set[*scopedItem]) bool {
				return lhs.Equal(rhs)
			},
		),
	}
}

func (t *propagationTable) Add(from *scopedItem, to ...*scopedItem) bool {
	if _, ok := t.table.Get(from); !ok {
		t.table.Put(from, set.New[*scopedItem](eqScopedItem))
	}

	set, _ := t.table.Get(from)
	size := set.Size()
	set.Add(to...)

	return set.Size() > size
}

func (t *propagationTable) Get(from *scopedItem) set.Set[*scopedItem] {
	if set, ok := t.table.Get(from); ok {
		return set
	}

	return nil
}

func (t *propagationTable) All() iter.Seq2[*scopedItem, set.Set[*scopedItem]] {
	return t.table.All()
}

type lookaheadTable struct {
	table symboltable.SymbolTable[*scopedItem, set.Set[grammar.Terminal]]
}

func NewLookaheadTable() *lookaheadTable {
	return &lookaheadTable{
		table: symboltable.NewRedBlack(
			cmpScopedItem,
			func(lhs, rhs set.Set[grammar.Terminal]) bool {
				return lhs.Equal(rhs)
			},
		),
	}
}

func (t *lookaheadTable) Add(item *scopedItem, lookahead ...grammar.Terminal) bool {
	if _, ok := t.table.Get(item); !ok {
		t.table.Put(item, set.New[grammar.Terminal](grammar.EqTerminal))
	}

	set, _ := t.table.Get(item)
	size := set.Size()
	set.Add(lookahead...)

	return set.Size() > size
}

func (t *lookaheadTable) Get(item *scopedItem) set.Set[grammar.Terminal] {
	if set, ok := t.table.Get(item); ok {
		return set
	}

	return nil
}

func (t *lookaheadTable) All() iter.Seq2[*scopedItem, set.Set[grammar.Terminal]] {
	return t.table.All()
}
