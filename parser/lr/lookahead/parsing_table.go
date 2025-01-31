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

// BuildParsingTable constructs a parsing table for an LALR parser.
//
// This method constructs an LALR(1) parsing table for any context-free grammar.
// To identify errors in the table, use the Error method.
func BuildParsingTable(G *grammar.CFG) (*lr.ParsingTable, error) {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The LALR parsing table functions ACTION and GOTO for G′.
	 */

	H := lr.NewGrammarWithLR1Kernel(G, nil)

	K := ComputeLALR1Kernels(G) // 1. Construct the kernels of the LALR(1) collection of sets of items for G′.
	S := lr.BuildStateMap(K)    // Map sets of LALR(1) items to state numbers.

	terminals := H.OrderTerminals()
	_, _, nonTerminals := H.OrderNonTerminals()
	table := lr.NewParsingTable(S.States(), terminals, nonTerminals)

	// 2. State i is constructed from I.
	for i, I := range S.All() {
		// The parsing action for state i is determined as follows:

		for item := range H.CLOSURE(I).All() {
			item := item.(*lr.Item1)

			// If "A → α•aβ, b" is in Iᵢ and GOTO(Iᵢ,a) = Iⱼ (a must be a terminal)
			if X, ok := item.DotSymbol(); ok {
				if a, ok := X.(grammar.Terminal); ok {
					// J is the union of multiple sets of LR(1) items (J = I₁ ∪ I₂ ∪ ... ∪ Iₖ).
					// By definition, I₁, I₂, ..., Iₖ all share the same core, which consists of the production and dot positions, excluding the lookahead.
					// The cores of GOTO(I₁, X), GOTO(I₂, X), ..., GOTO(Iₖ, X) are also the same because the sets I₁, I₂, ..., Iₖ have identical cores.
					// Let K be the union of all item sets that have the same core as GOTO(I₁, X). Then, GOTO(J, X) = K.
					// This means that the regular GOTO function returns an item set whose core matches one of the LALR(1) item sets, but it may have missing lookaheads.
					// Therefore, to find K, we need to identify an item set whose core is a superset of the set returned by the regular GOTO function.
					J := H.GOTO(I, a)
					j := findSuperset(S, J)

					// Set ACTION[i,a] to SHIFT j
					table.AddACTION(i, a, &lr.Action{
						Type:  lr.SHIFT,
						State: j,
					})
				}
			}

			// If "A → α•, a" is in Iᵢ (A ≠ S′)
			if item.IsComplete() && !item.IsFinal() {
				a := item.Lookahead

				// Set ACTION[i,a] to REDUCE A → α
				table.AddACTION(i, a, &lr.Action{
					Type:       lr.REDUCE,
					Production: item.Production,
				})
			}

			// If "S′ → S•, $" is in Iᵢ
			if item.IsFinal() {
				// Set ACTION[i,$] to ACCEPT
				table.AddACTION(i, grammar.Endmarker, &lr.Action{
					Type: lr.ACCEPT,
				})
			}

			// If any conflicting actions result from the above rules, the grammar is not LALR(1).
			// The table.Error() method will list all conflicts, if any exist.
		}

		// 3. The goto transitions for state i are constructed for all non-terminals A using the rule:
		// If GOTO(Iᵢ,A) = Iⱼ
		for A := range H.NonTerminals.All() {
			if !A.Equal(H.Start) {
				// J is the union of multiple sets of LR(1) items (J = I₁ ∪ I₂ ∪ ... ∪ Iₖ).
				// By definition, I₁, I₂, ..., Iₖ all share the same core, which consists of the production and dot positions, excluding the lookahead.
				// The cores of GOTO(I₁, X), GOTO(I₂, X), ..., GOTO(Iₖ, X) are also the same because the sets I₁, I₂, ..., Iₖ have identical cores.
				// Let K be the union of all item sets that have the same core as GOTO(I₁, X). Then, GOTO(J, X) = K.
				// This means that the regular GOTO function returns an item set whose core matches one of the LALR(1) item sets, but it may have missing lookaheads.
				// Therefore, to find K, we need to identify an item set whose core is a superset of the set returned by the regular GOTO function.
				J := H.GOTO(I, A)
				j := findSuperset(S, J)

				// Set GOTO[i,A] = j
				table.SetGOTO(i, A, j)
			}
		}

		// 4. All entries not defined by rules (2) and (3) are made ERROR.
	}

	// 5. The initial state of the parser is the one constructed from the set of items containing "S′ → •S, $".

	if err := table.Error(); err != nil {
		return nil, err
	}

	return table, nil
}

// ComputeLALR1Kernels computes and returns the kernels of the LALR(1) collection of sets of items for a context-free grammar.
func ComputeLALR1Kernels(G *grammar.CFG) lr.ItemSetCollection {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The kernels of the LALR(1) collection of sets of items for G′.
	 * METHOD:
	 *
	 *   1. Construct the kernels of the sets of LR(0) items for G′.
	 *      If space is not at a premium, the simplest way is to
	 *      construct the LR(0) sets of items normally, and then remove the non-kernel items.
	 *      If space is severely constrained, we may wish instead to store only the kernel items for each set,
	 *      and compute GOTO for a set of items I by first computing the closure of I.
	 *
	 *   2. Apply determining-lookaheads to the kernel of each set of LR(0) items and grammar symbol X
	 *      to determine which lookaheads are spontaneously generated for kernel items in GOTO(I,X),
	 *      and from which items in I lookaheads are propagated to kernel items in GOTO(I,X).
	 *
	 *        For a kernel K of a set of LR(0) items I and a grammar symbol X:
	 *          for ( each item A → α.β in K ) {
	 *            J = CLOSURE({[A → α.β, $]})
	 *            if ( [B → γ.Xδ, a] is in J, and a is not $ )
	 *              conclude that lookahead a is generated spontaneously for item B → γX.δ in GOTO(I, X);
	 *            if ( [B → γ.Xδ, $] is in J )
	 *              conclude that lookaheads propagate from A → α.β in I to B → γX.δ in GOTO(I, X);
	 *
	 *   3. Initialize a table that gives, for each kernel item in each set of items, the associated lookaheads.
	 *      Initially, each item has associated with it only those lookaheads
	 *      that we determined in step 2 were generated spontaneously.
	 *
	 *   4. Make repeated passes over the kernel items in all sets.
	 *      When we visit an item i, we look up the kernel items to which i propagates its lookaheads,
	 *      using information tabulated in step 2.
	 *      The current set of lookaheads for i is added to those already associated
	 *      with each of the items to which i propagates its lookaheads.
	 *      We continue making passes over the kernel items until no more new lookaheads are propagated.
	 */

	H0 := lr.NewGrammarWithLR0Kernel(G, nil)
	H1 := lr.NewGrammarWithLR1Kernel(G, nil)

	K0 := H0.Canonical()       // Construct the kernels of the sets of LR(0) items for G′.
	S0 := lr.BuildStateMap(K0) // Map Kernel sets of LR(0) items to state numbers.

	propagations := newPropagationTable(S0) // This table memoize which items propagate their lookaheads to which other items.
	lookaheads := newLookaheadTable(S0)     // This table is used for computing lookaheads for all states.

	// Initialize the table with the spontaneous lookahead $ for the initial item "S′ → •S".
	init := &scopedItem{ItemSet: 0, Item: 0}
	lookaheads.Add(init, grammar.Endmarker)

	// For each state, kernel set of LR(0) items, determine:
	//
	//   1. Which lookheads are generated spontaneously.
	//   2. Which lookheads are propagated.
	//
	// There are two goals:
	//
	//   • Determine which lookaheads are generated spontaneously for which states.
	//   • Build a table to memoize which states propagate their lookaheads to which other states.
	for s, items := range S0 {
		s := lr.State(s)
		I := S0.ItemSet(s)

		for i, item := range items {
			from := &scopedItem{ItemSet: s, Item: i}

			// J = CLOSURE({[A → α.β, $]})
			J := H1.CLOSURE(
				lr.NewItemSet(
					item.(*lr.Item0).Item1(grammar.Endmarker),
				),
			)

			for j := range J.All() {
				if X, ok := j.DotSymbol(); ok {
					nextI := H0.GOTO(I, X) // GOTO(I,X)
					nextj, _ := j.Next()   // B → γX.δ

					// Find B → γX.δ in GOTO(I,X)
					to := new(scopedItem)
					to.ItemSet = S0.FindItemSet(nextI)
					to.Item = S0.FindItem(to.ItemSet, nextj.(*lr.Item1).Item0())

					if lookahead := j.(*lr.Item1).Lookahead; lookahead.Equal(grammar.Endmarker) {
						// If [B → γ.Xδ, $] is in J,
						// conclude that lookaheads propagate from A → α.β in I to B → γX.δ in GOTO(I,X).
						propagations.Add(from, to)
					} else {
						// If [B → γ.Xδ, a] is in J, and a is not $,
						// conclude that lookahead is generated spontaneously for item B → γX.δ in GOTO(I,X).
						lookaheads.Add(to, lookahead)
					}
				}
			}
		}
	}

	// Propagate lookaheads between states until they stabilize,
	// meaning no further changes occur to the lookaheads for any state.
	for updated := true; updated; {
		updated = false

		for from := range lookaheads.All() {
			if fromL := generic.Collect1(lookaheads.Get(from).All()); len(fromL) > 0 {
				if set := propagations.Get(from); set != nil {
					for to := range set.All() {
						updated = updated || lookaheads.Add(to, fromL...)
					}
				}
			}
		}
	}

	K1 := lr.NewItemSetCollection()

	// Build the kernels of the LALR(1) collection of item sets.
	for s, items := range S0 {
		s := lr.State(s)
		J := lr.NewItemSet()

		for i := range items {
			item := &scopedItem{ItemSet: s, Item: i}
			for lookahead := range lookaheads.Get(item).All() {
				J.Add(items[i].(*lr.Item0).Item1(lookahead))
			}
		}

		K1.Add(J)
	}

	return K1
}

// findSuperset searches the state map for a state whose item set is a superset of the given item set.
// A superset means that all items in the given item set are contained within the state's item set.
// If such a state is found, its index is returned.
// If no such state exists, or if the provided item set is empty, ErrState is returned.
func findSuperset(S lr.StateMap, I lr.ItemSet) lr.State {
	if I.Size() == 0 {
		return lr.ErrState
	}

	for i := range S {
		if s := lr.State(i); S.ItemSet(s).IsSuperset(I) {
			return s
		}
	}

	return lr.ErrState
}

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

func newPropagationTable(S lr.StateMap) *propagationTable {
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

// lookaheadTable keeps track of lookaheads for each scoped item.
type lookaheadTable struct {
	S     lr.StateMap
	table symboltable.SymbolTable[*scopedItem, set.Set[grammar.Terminal]]
}

func newLookaheadTable(S lr.StateMap) *lookaheadTable {
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
