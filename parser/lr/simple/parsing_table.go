package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// BuildParsingTable constructs a parsing table for an SLR parser.
func BuildParsingTable(G *grammar.CFG, precedences lr.PrecedenceLevels) (*lr.ParsingTable, error) {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The SLR parsing table functions ACTION and GOTO for G′.
	 */

	G0 := lr.NewGrammarWithLR0(G)
	FIRST := G0.ComputeFIRST()
	FOLLOW := G0.ComputeFOLLOW(FIRST)

	C := G0.Canonical()      // 1. Construct C = {I₀, I₁, ..., Iₙ}, the collection of sets of LR(0) items for G′.
	S := lr.BuildStateMap(C) // Map sets of LR(0) items to state numbers.

	terminals := G0.OrderTerminals()
	_, _, nonTerminals := G0.OrderNonTerminals()
	table := lr.NewParsingTable(S.States(), terminals, nonTerminals, precedences)

	// 2. State i is constructed from I.
	for i, I := range S.All() {
		// The parsing action for state i is determined as follows:

		for item := range I.All() {
			item := item.(*lr.Item0)

			// If "A → α•aβ" is in Iᵢ and GOTO(Iᵢ,a) = Iⱼ (a must be a terminal)
			if X, ok := item.DotSymbol(); ok {
				if a, ok := X.(grammar.Terminal); ok {
					J := G0.GOTO(I, a)
					j := S.FindItemSet(J)

					// Set ACTION[i,a] to SHIFT j
					table.AddACTION(i, a, &lr.Action{
						Type:  lr.SHIFT,
						State: j,
					})
				}
			}

			// If "A → α•" is in Iᵢ (A ≠ S′)
			if item.IsComplete() && !item.IsFinal() {
				FOLLOWA := FOLLOW(item.Head)

				// For all a in FOLLOW(A)
				for a := range FOLLOWA.Terminals.All() {
					// Set ACTION[i,a] to REDUCE A → α
					table.AddACTION(i, a, &lr.Action{
						Type:       lr.REDUCE,
						Production: item.Production,
					})
				}

				if FOLLOWA.IncludesEndmarker {
					// Set ACTION[i,$] to REDUCE A → α
					table.AddACTION(i, grammar.Endmarker, &lr.Action{
						Type:       lr.REDUCE,
						Production: item.Production,
					})
				}
			}

			// If "S′ → S•" is in Iᵢ
			if item.IsFinal() {
				// Set ACTION[i,$] to ACCEPT
				table.AddACTION(i, grammar.Endmarker, &lr.Action{
					Type: lr.ACCEPT,
				})
			}

			// If any conflicting actions result from the above rules, the grammar is not SLR(1).
			// The table.Error() method will list all conflicts, if any exist.
		}

		// 3. The goto transitions for state i are constructed for all non-terminals A using the rule:
		// If GOTO(Iᵢ,A) = Iⱼ
		for A := range G0.NonTerminals.All() {
			if !A.Equal(G0.Start) {
				J := G0.GOTO(I, A)
				j := S.FindItemSet(J)

				// Set GOTO[i,A] = j
				table.SetGOTO(i, A, j)
			}
		}

		// 4. All entries not defined by rules (2) and (3) are made ERROR.
	}

	// 5. The initial state of the parser is the one constructed from the set of items containing "S′ → •S".

	// Try resolving any conflicts in the ACTION parsing table.
	if err := table.ResolveConflicts(); err != nil {
		return table, err
	}

	return table, nil
}
