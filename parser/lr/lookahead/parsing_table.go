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
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The LALR parsing table functions ACTION and GOTO for G′.
	 */

	auto1 := lr.NewLR1KernelAutomaton(G)

	K := ComputeLALR1Kernels(G) // 1. Construct the kernels of the LALR(1) collection of sets of items for G′.
	S := lr.BuildStateMap(K)    // Map sets of LALR(1) items to state numbers.

	terminals := auto1.G().OrderTerminals()
	_, _, nonTerminals := auto1.G().OrderNonTerminals()
	table := lr.NewParsingTable(S.States(), terminals, nonTerminals)

	// 2. State i is constructed from I.
	for i, I := range S.All() {
		// The parsing action for state i is determined as follows:

		for item := range auto1.CLOSURE(I).All() {
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
					J := auto1.GOTO(I, a)
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
		for A := range auto1.G().NonTerminals.All() {
			if !A.Equal(auto1.G().Start) {
				// J is the union of multiple sets of LR(1) items (J = I₁ ∪ I₂ ∪ ... ∪ Iₖ).
				// By definition, I₁, I₂, ..., Iₖ all share the same core, which consists of the production and dot positions, excluding the lookahead.
				// The cores of GOTO(I₁, X), GOTO(I₂, X), ..., GOTO(Iₖ, X) are also the same because the sets I₁, I₂, ..., Iₖ have identical cores.
				// Let K be the union of all item sets that have the same core as GOTO(I₁, X). Then, GOTO(J, X) = K.
				// This means that the regular GOTO function returns an item set whose core matches one of the LALR(1) item sets, but it may have missing lookaheads.
				// Therefore, to find K, we need to identify an item set whose core is a superset of the set returned by the regular GOTO function.
				J := auto1.GOTO(I, A)
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

	auto0 := lr.NewLR0KernelAutomaton(G)
	auto1 := lr.NewLR1KernelAutomaton(G)

	K0 := auto0.Canonical()    // Construct the kernels of the sets of LR(0) items for G′.
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
