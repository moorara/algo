package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// calculator implemented the lr.Calculator interface for LR(0) items.
type calculator struct {
	augG *grammar.CFG
}

// NewCalculator returns an lr.Calculator for LR(0) items.
// It provides an implementation of the CLOSURE function
// based on the LR(0) items of the augmented grammar.
func NewCalculator(G *grammar.CFG) *lr.AutomatonCalculator {
	augG := lr.Augment(G)

	return &lr.AutomatonCalculator{
		Calculator: &calculator{
			augG: augG,
		},
	}
}

// G returns the augmented context-free grammar.
func (c *calculator) G() *grammar.CFG {
	return c.augG
}

// Initial returns the initial LR(0) item "S′ → •S" for an augmented grammar.
func (c *calculator) Initial() lr.Item {
	for p := range c.augG.Productions.Get(c.augG.Start).All() {
		return &LR0Item{
			Production: p,
			Start:      c.augG.Start,
			Dot:        0,
		}
	}

	// This will never be the case.
	return nil
}

// CLOSURE computes the closure of a given LR(0) item set.
func (c *calculator) CLOSURE(I lr.ItemSet) lr.ItemSet {
	/*
	 * If I is a set of items for a grammar G,
	 * then CLOSURE(I) is the set of items constructed from I by the two rules:
	 *
	 *   1. Initially, add every item in I to CLOSURE(I).
	 *   2. If A → α•Bβ is in CLOSURE(I) and B → γ is a production,
	 *      then add the item B → •γ to CLOSURE(I), if it is not already there.
	 *      Apply this rule until no more new items can be added to CLOSURE(I).
	 */

	J := I.Clone()

	for newItems := []lr.Item{}; newItems != nil; {
		newItems = nil

		// For each item A → α•Bβ in J
		for i := range J.All() {
			if i, ok := i.(*LR0Item); ok {
				if X, ok := i.DotSymbol(); ok {
					if B, ok := X.(grammar.NonTerminal); ok {
						// For each production B → γ of G′
						for BProd := range c.augG.Productions.Get(B).All() {
							j := &LR0Item{
								Production: BProd,
								Start:      c.augG.Start,
								Dot:        0,
							}

							// If B → •γ is not in J
							if !J.Contains(j) {
								newItems = append(newItems, j)
							}
						}
					}
				}
			}
		}

		J.Add(newItems...)
	}

	return J
}

// BuildParsingTable constructs a parsing table for an SLR parser.
//
// This method constructs an LR(0) parsing table for any context-free grammar.
// To identify errors in the table, use the Error method.
func BuildParsingTable(G *grammar.CFG) (*lr.ParsingTable, error) {
	/*
	 * INPUT:  An augmented grammar G′.
	 * OUTPUT: The SLR-parsing table functions ACTION and GOTO for G′.
	 */

	calc := NewCalculator(G)
	FIRST := calc.G().ComputeFIRST()
	FOLLOW := calc.G().ComputeFOLLOW(FIRST)

	// 1. Construct C = {I₀, I₁, ..., Iₙ}, the collection of sets of LR(0) items for G′.
	C := calc.Canonical()

	states := lr.BuildStateMap(C)
	terminals := G.OrderTerminals()
	_, _, nonTerminals := G.OrderNonTerminals()
	table := lr.NewParsingTable(states.All(), terminals, nonTerminals)

	// 2. State i is constructed from I.
	for i, I := range states {
		// The parsing action for state i is determined as follows:

		for item := range I.All() {
			if item, ok := item.(*LR0Item); ok {
				// If "A → α•aβ" is in Iᵢ and GOTO(Iᵢ,a) = Iⱼ (a must be a terminal)
				if X, ok := item.DotSymbol(); ok {
					if a, ok := X.(grammar.Terminal); ok {
						J := calc.GOTO(I, a)
						j := states.For(J)

						// Set ACTION[i,a] to SHIFT j
						table.AddACTION(lr.State(i), a, &lr.Action{
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
						table.AddACTION(lr.State(i), a, &lr.Action{
							Type:       lr.REDUCE,
							Production: item.Production,
						})
					}

					if FOLLOWA.IncludesEndmarker {
						// Set ACTION[i,$] to REDUCE A → α
						table.AddACTION(lr.State(i), grammar.Endmarker, &lr.Action{
							Type:       lr.REDUCE,
							Production: item.Production,
						})
					}
				}

				// If "S′ → S•" is in Iᵢ
				if item.IsFinal() {
					// Set ACTION[i,$] to ACCEPT
					table.AddACTION(lr.State(i), grammar.Endmarker, &lr.Action{
						Type: lr.ACCEPT,
					})
				}

				// If any conflicting actions result from the above rules, the grammar is not SLR(1).
				// The table.Error() method will list all conflicts, if any exist.
			}
		}

		// 3. The goto transitions for state i are constructed for all non-terminals A using the rule:
		// If GOTO(Iᵢ,A) = Iⱼ
		for A := range calc.G().NonTerminals.All() {
			if !A.Equals(calc.G().Start) {
				J := calc.GOTO(I, A)
				j := states.For(J)

				// Set GOTO[i,A] = j
				table.SetGOTO(lr.State(i), A, j)
			}
		}

		// 4. All entries not defined by rules (2) and (3) are made ERROR.
	}

	// 5. The initial state of the parser is the one constructed from the set of items containing "S′ → •S".

	if err := table.Error(); err != nil {
		return nil, err
	}

	return table, nil
}
