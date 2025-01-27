package lookahead

import (
	"bytes"
	"fmt"
	"iter"
	"strings"
	"unicode/utf8"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// scopedItem represents an individual item within an item set.
// The item set is represented by its state,
// and the individual item is identified by its index within the sorted item set.
type scopedItem struct {
	ItemSet lr.State
	Item    int
}

func (i *scopedItem) String(S lr.StateMap) string {
	return fmt.Sprintf("[%d] %s", i.ItemSet, S[i.ItemSet][i.Item])
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
	S     lr.StateMap
	table symboltable.SymbolTable[*scopedItem, set.Set[*scopedItem]]
}

func NewPropagationTable(S lr.StateMap) *propagationTable {
	return &propagationTable{
		S: S,
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
		t.table.Put(from, set.New(eqScopedItem))
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

func (t *propagationTable) String() string {
	var b bytes.Buffer
	title0, title1 := "FROM", "TO"

	// Calculate the width of the first column.
	w0 := len(title0)
	for item := range t.table.All() {
		if l := utf8.RuneCountInString(item.String(t.S)); l > w0 {
			w0 = l
		}
	}
	w0 += 2

	// Calculate the width of the second column.
	w1 := len(title1)
	for _, items := range t.table.All() {
		for item := range items.All() {
			if l := utf8.RuneCountInString(item.String(t.S)); l > w1 {
				w1 = l
			}
		}
	}
	w1 += 2

	printTopLine(&b, w0, w1) // ┌──────────┬──────────┐
	printFirst(&b, w0, title0)
	printSecond(&b, w1, title1)

	for from, set := range t.All() {
		printMiddleLine(&b, w0, w1) // ├──────────┼──────────┤

		first := from.String(t.S)
		items := generic.Collect1(set.All())
		sort.Quick(items, cmpScopedItem)

		for _, to := range items {
			printFirst(&b, w0, first)
			first = ""
			printSecond(&b, w1, to.String(t.S))
		}
	}

	printBottomLine(&b, w0, w1) // └──────────┴──────────┘

	return b.String()
}

type lookaheadTable struct {
	S     lr.StateMap
	table symboltable.SymbolTable[*scopedItem, set.Set[grammar.Terminal]]
}

func NewLookaheadTable(S lr.StateMap) *lookaheadTable {
	return &lookaheadTable{
		S: S,
		table: symboltable.NewRedBlack(
			cmpScopedItem,
			func(lhs, rhs set.Set[grammar.Terminal]) bool {
				return lhs.Equal(rhs)
			},
		),
	}
}

func formatLookaheads(lookaheads []grammar.Terminal) string {
	vals := make([]string, len(lookaheads))
	for i, l := range lookaheads {
		vals[i] = l.String()
	}

	return strings.Join(vals, ", ")
}

func (t *lookaheadTable) Add(item *scopedItem, lookahead ...grammar.Terminal) bool {
	if _, ok := t.table.Get(item); !ok {
		t.table.Put(item, set.NewWithFormat(grammar.EqTerminal, formatLookaheads))
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

func (t *lookaheadTable) String() string {
	var b bytes.Buffer
	title0, title1 := "ITEM", "LOOKAHEADS"

	// Calculate the width of the first column.
	w0 := len(title0)
	for item := range t.table.All() {
		if l := utf8.RuneCountInString(item.String(t.S)); l > w0 {
			w0 = l
		}
	}
	w0 += 2

	// Calculate the width of the second column.
	w1 := len(title1)
	for _, lookaheads := range t.table.All() {
		if l := utf8.RuneCountInString(lookaheads.String()); l > w1 {
			w1 = l
		}
	}
	w1 += 2

	printTopLine(&b, w0, w1) // ┌──────────┬──────────┐
	printFirst(&b, w0, title0)
	printSecond(&b, w1, title1)

	for item, lookaheads := range t.All() {
		printMiddleLine(&b, w0, w1) // ├──────────┼──────────┤
		printFirst(&b, w0, item.String(t.S))
		printSecond(&b, w1, lookaheads.String())
	}

	printBottomLine(&b, w0, w1) // └──────────┴──────────┘

	return b.String()
}

func printTopLine(b *bytes.Buffer, w0, w1 int) {
	b.WriteRune('┌')
	b.WriteString(strings.Repeat("─", w0))
	b.WriteRune('┬')
	b.WriteString(strings.Repeat("─", w1))
	b.WriteRune('┐')
	b.WriteRune('\n')
}

func printMiddleLine(b *bytes.Buffer, w0, w1 int) {
	b.WriteRune('├')
	b.WriteString(strings.Repeat("─", w0))
	b.WriteRune('┼')
	b.WriteString(strings.Repeat("─", w1))
	b.WriteRune('┤')
	b.WriteRune('\n')
}

func printBottomLine(b *bytes.Buffer, w0, w1 int) {
	b.WriteRune('└')
	b.WriteString(strings.Repeat("─", w0))
	b.WriteRune('┴')
	b.WriteString(strings.Repeat("─", w1))
	b.WriteRune('┘')
	b.WriteRune('\n')
}

func printFirst(b *bytes.Buffer, w0 int, s0 string) {
	pad0 := w0 - utf8.RuneCountInString(s0)
	lpad0 := 1
	rpad0 := pad0 - lpad0

	b.WriteRune('│')
	b.WriteString(strings.Repeat(" ", lpad0))
	b.WriteString(s0)
	b.WriteString(strings.Repeat(" ", rpad0))
	b.WriteRune('│')
}

func printSecond(b *bytes.Buffer, w1 int, s1 string) {
	pad1 := w1 - utf8.RuneCountInString(s1)
	lpad1 := 1
	rpad1 := pad1 - lpad1

	b.WriteString(strings.Repeat(" ", lpad1))
	b.WriteString(s1)
	b.WriteString(strings.Repeat(" ", rpad1))
	b.WriteRune('│')
	b.WriteRune('\n')
}
