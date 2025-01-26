package lookahead

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// BuildParsingTable constructs a parsing table for an LALR parser.
//
// This method constructs an LALR(1) parsing table for any context-free grammar.
// To identify errors in the table, use the Error method.
func BuildParsingTable(G *grammar.CFG) (lr.ParsingTable, error) {
	// TODO
	return nil, nil
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

	auto0 := lr.NewLR0KernelAutomaton(G)
	auto1 := lr.NewLR1KernelAutomaton(G)

	K0 := auto0.Canonical()    // Construct the kernels of the sets of LR(0) items for G′.
	S0 := lr.BuildStateMap(K0) // Map Kernel sets of LR(0) items to state numbers.

	propagations := NewPropagationTable() // This table memoize which items propagate their lookaheads to which other items.
	lookaheads := NewLookaheadTable()     // This table is used for computing lookaheads for all states.

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
			J := auto1.CLOSURE(
				lr.NewItemSet(
					item.(*lr.Item0).Item1(grammar.Endmarker),
				),
			)

			for j := range J.All() {
				if X, ok := j.DotSymbol(); ok {
					nextI := auto0.GOTO(I, X) // GOTO(I,X)
					nextj, _ := j.Next()      // B → γX.δ

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
						updated = lookaheads.Add(to, fromL...)
					}
				}
			}
		}
	}

	// Build the kernels of the LALR(1) collection of item sets.
	K1 := lr.NewItemSetCollection()

	for s, items := range S0 {
		J := lr.NewItemSet()

		for i := range items {
			item := &scopedItem{ItemSet: lr.State(s), Item: i}
			for lookahead := range lookaheads.Get(item).All() {
				J.Add(items[i].(*lr.Item0).Item1(lookahead))
			}
		}

		K1.Add(J)
	}

	return K1
}
