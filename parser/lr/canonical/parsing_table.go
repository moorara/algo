package canonical

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// BuildParsingTable constructs a parsing table for an SLR parser.
func BuildParsingTable(G *grammar.CFG) (*lr.ParsingTable, error) {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The canonical LR parsing table functions ACTION and GOTO for G′.
	 */

	H := lr.NewGrammarWithLR1(G, nil)

	C := H.Canonical()       // 1. Construct C = {I₀, I₁, ..., Iₙ}, the collection of sets of LR(1) items for G′.
	S := lr.BuildStateMap(C) // Map sets of LR(1) items to state numbers.

	terminals := H.OrderTerminals()
	_, _, nonTerminals := H.OrderNonTerminals()
	table := lr.NewParsingTable(S.States(), terminals, nonTerminals)

	// 2. State i is constructed from I.
	for i, I := range S.All() {
		// The parsing action for state i is determined as follows:

		for item := range I.All() {
			item := item.(*lr.Item1)

			// If "A → α•aβ, b" is in Iᵢ and GOTO(Iᵢ,a) = Iⱼ (a must be a terminal)
			if X, ok := item.DotSymbol(); ok {
				if a, ok := X.(grammar.Terminal); ok {
					J := H.GOTO(I, a)
					j := S.FindItemSet(J)

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

			// If any conflicting actions result from the above rules, the grammar is not LR(1).
			// The table.Error() method will list all conflicts, if any exist.
		}

		// 3. The goto transitions for state i are constructed for all non-terminals A using the rule:
		// If GOTO(Iᵢ,A) = Iⱼ
		for A := range H.NonTerminals.All() {
			if !A.Equal(H.Start) {
				J := H.GOTO(I, A)
				j := S.FindItemSet(J)

				// Set GOTO[i,A] = j
				table.SetGOTO(i, A, j)
			}
		}

		// 4. All entries not defined by rules (2) and (3) are made ERROR.
	}

	// 5. The initial state of the parser is the one constructed from the set of items containing "S′ → •S, $".

	return table, table.Conflicts()
}
