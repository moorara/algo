package slr

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

var primeSuffixes = []string{
	"′", // Prime (U+2032)
	"″", // Double Prime (U+2033)
	"‴", // Triple Prime (U+2034)
	"⁗", // Quadruple Prime (U+2057)
}

// AugmentedCFG is a context-free grammar with a new start symbol S′ and production S′ → S.
// An augmented grammar is used in LR parsing to signal when
// the parser should stop and confirm that the input has been successfully parsed.
type AugmentedCFG struct {
	grammar.CFG
}

// AugmentCFG creates and returns an augmented grammar G′ from a give grammar G.
func AugmentCFG(G grammar.CFG) AugmentedCFG {
	start := G.Start
	newG := G.Clone()

	newStart := newG.AddNewNonTerminal(start, primeSuffixes...)
	newG.Start = newStart
	newG.Productions.Add(grammar.Production{
		Head: newStart,
		Body: grammar.String[grammar.Symbol]{start}},
	)

	return AugmentedCFG{
		CFG: newG,
	}
}

// Equals determines whether or not two augmented context-free grammars are the same.
func (g AugmentedCFG) Equals(rhs AugmentedCFG) bool {
	return g.CFG.Equals(rhs.CFG)
}

// Closure computes the closure of an item set.
//
// Intuitively, A → α•Bβ in CLOSURE(I) indicates that at some point in the parsing process,
// we anticipate seeing a substring derivable from Bβ as input.
// The substring derivable from Bβ will have a prefix derivable from B.
// Thus, for every production B → γ, we include the item B → •γ in CLOSURE(I).
func (g AugmentedCFG) CLOSURE(I set.Set[Item]) set.Set[Item] {
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

	for newItems := []Item{}; newItems != nil; {
		newItems = nil

		// For each item A → α•Bβ in J
		for i := range J.All() {
			if X, ok := i.DotSymbol(); ok {
				if B, ok := X.(grammar.NonTerminal); ok {
					// For each production B → γ of G
					for BProd := range g.Productions.Get(B).All() {
						// If B → •γ is not in J
						if j := (Item{&BProd, 0}); !J.Contains(j) {
							newItems = append(newItems, j)
						}
					}
				}
			}
		}

		J.Add(newItems...)
	}

	return J
}

// GOTO(I, X) computes the closure of the set of all items A → αX•β such that
// A → α•Xβ is in the set of items I, where I is a set of items and X is a grammar symbol.
//
// Intuitively, the GOTO function defines transitions in the LR(0) automaton for a grammar.
// The states of the automaton correspond to sets of items, and GOTO(I, X) specifies
// the transition from the state represented by I when the grammar symbol X is encountered.
func (g AugmentedCFG) GOTO(I set.Set[Item], X grammar.Symbol) set.Set[Item] {
	J := NewItemSet()

	for i := range I.All() {
		if Y, ok := i.DotSymbol(); ok {
			if Y.Equals(X) {
				if nextI, ok := i.NextItem(); ok {
					J = J.Union(g.CLOSURE(NewItemSet(nextI)))
				}
			}
		}
	}

	return J
}
