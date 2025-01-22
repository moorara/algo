package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// BuildParsingTable constructs a parsing table for an SLR parser.
//
// This method constructs an LR(0) parsing table for any context-free grammar.
// To identify errors in the table, use the Error method.
func BuildParsingTable(G grammar.CFG) *lr.ParsingTable {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The SLR-parsing table functions ACTION and GOTO for G′.
	 */

	augG := lr.Augment(G)
	first := augG.ComputeFIRST()
	follow := augG.ComputeFOLLOW(first)

	init := Initial(augG)
	CLOSURE := LR0Closure(augG)

	// 1. Construct C = {I₀, I₁, ..., Iₙ}, the collection of sets of LR(0) items for G′.
	C := lr.Canonical(augG, init, CLOSURE)

	states := lr.BuildStateMap(C)
	terminals := G.OrderTerminals()
	_, _, nonTerminals := G.OrderNonTerminals()

	table := lr.NewParsingTable(states.All(), terminals, nonTerminals)

	// 2. State i is constructed from I.
	for i, I := range states {
		// The parsing actions for state i are determined as follows:

		for item := range I.All() {
			if item, ok := item.(LR0Item); ok {
				// If A → α•aβ is in Iᵢ and GOTO(Iᵢ,a) = Iⱼ (a must be a terminal)
				if X, ok := item.DotSymbol(); ok {
					if a, ok := X.(grammar.Terminal); ok {
						J := lr.GOTO(CLOSURE, I, a)
						j := states.For(J)

						// Set ACTION[i,a] to SHIFT j
						table.AddACTION(lr.State(i), a, lr.Action{
							Type:  lr.SHIFT,
							State: j,
						})
					}
				}

				// If A → α• is in Iᵢ (A ≠ S′)
				if item.IsComplete() && !item.IsFinal() {
					followA := follow(item.Head)

					// For all a in FOLLOW(A)
					for a := range followA.Terminals.All() {
						// Set ACTION[i,a] to REDUCE A → α
						table.AddACTION(lr.State(i), a, lr.Action{
							Type:       lr.REDUCE,
							Production: item.Production,
						})
					}

					if followA.IncludesEndmarker {
						// Set ACTION[i,$] to REDUCE A → α
						table.AddACTION(lr.State(i), grammar.Endmarker, lr.Action{
							Type:       lr.REDUCE,
							Production: item.Production,
						})
					}
				}

				// If S′ → S• is in Iᵢ
				if item.IsFinal() {
					// Set ACTION[i,$] to ACCEPT
					table.AddACTION(lr.State(i), grammar.Endmarker, lr.Action{
						Type: lr.ACCEPT,
					})
				}

				// If any conflicting actions result from the above rules, the grammar is not SLR(1).
				// The table.Error() method will list all conflicts, if any exist.
			}
		}

		// 3. The goto transitions for state i are constructed for all non-terminals A using the rule:
		// If GOTO(Iᵢ,A) = Iⱼ
		for A := range augG.NonTerminals.All() {
			if !A.Equals(augG.Start) {
				J := lr.GOTO(CLOSURE, I, A)
				j := states.For(J)

				// Set GOTO[i,A] = j
				table.SetGOTO(lr.State(i), A, j)
			}
		}

		// 4. All entries not defined by rules (2) and (3) are made ERROR.
	}

	// 5. The initial state of the parser is the one constructed from the set of items containing "S′ → •S".

	return table
}
