package slr

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// Initial returns the initial LR(0) item "S′ → •S" for an augmented grammar.
func Initial(augG grammar.CFG) lr.Item {
	for p := range augG.Productions.Get(augG.Start).All() {
		return LR0Item{
			Production: &p,
			Start:      &augG.Start,
			Dot:        0,
		}
	}

	// This will never be the case.
	return LR0Item{}
}

// LR0Closure returns a function that computes closure of LR(0) item sets for an augmented grammar.
func LR0Closure(augG grammar.CFG) lr.ClosureFunc {
	// CLOSURE computes the closure of a given LR(0) item set.
	//
	// Intuitively, A → α•Bβ in CLOSURE(I) indicates that at some point in the parsing process,
	// we anticipate seeing a substring derivable from Bβ as input.
	// The substring derivable from Bβ will have a prefix derivable from B.
	// Thus, for every production B → γ, we include the LR(0) item B → •γ in CLOSURE(I).
	return func(I lr.ItemSet) lr.ItemSet {
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
				if i, ok := i.(LR0Item); ok {
					if X, ok := i.DotSymbol(); ok {
						if B, ok := X.(grammar.NonTerminal); ok {
							// For each production B → γ of G
							for BProd := range augG.Productions.Get(B).All() {
								j := LR0Item{
									Production: &BProd,
									Start:      &augG.Start,
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
}
