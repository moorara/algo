package lr

import "github.com/moorara/algo/grammar"

var (
	primeSuffixes = []string{
		"′", // Prime (U+2032)
		"″", // Double Prime (U+2033)
		"‴", // Triple Prime (U+2034)
		"⁗", // Quadruple Prime (U+2057)
	}
)

// Augment augments a context-free grammar G by creating a new start symbol S′
// and adding a production "S′ → S", where S is the original start symbol of G.
// This transformation is used to prepare grammars for LR parsing.
// The function clones the input grammar G, ensuring that the original grammar remains unmodified.
func Augment(G grammar.CFG) grammar.CFG {
	augG := G.Clone()

	newS := augG.AddNewNonTerminal(G.Start, primeSuffixes...)
	augG.Start = newS
	augG.Productions.Add(grammar.Production{
		Head: newS,
		Body: grammar.String[grammar.Symbol]{G.Start},
	})

	return augG
}

// Calculator defines the interface required for a general LR parser.
// Packages implementing an LR parser (e.g., Simple LR, Canonical LR, or LALR)
// must provide an implementation of this interface.
type Calculator interface {
	// G returns the augmented context-free grammar.
	G() grammar.CFG

	// Initial returns the initial item of an augmented grammar.
	// For LR(0), the initial item is "S′ → •S",
	// and for LR(1), the initial item is "S′ → •S, $".
	Initial() Item

	// Closure computes the closure of a given item set.
	// It may compute the closure for either an LR(0) item set or an LR(1) item set.
	CLOSURE(ItemSet) ItemSet
}

// AutomatonCalculator is used for constructing an LR automaton.
// It provides implementations of the GOTO and CLOSURE functions while
// delegating the computation of item set closures to the Calculator interface.
//
// Packages implementing an LR parser (e.g., Simple LR, Canonical LR, or LALR)
// must provide an implementation of the Calculator interface.
type AutomatonCalculator struct {
	Calculator
}

// GOTO(I, X) computes the closure of the set of all items "A → αX•β",
// where "A → α•Xβ" is in the set of items I and X is a grammar symbol.
//
// The GOTO function defines transitions in the automaton for the grammar.
// Each state of the automaton corresponds to a set of items, and
// GOTO(I, X) specifies the transition from the state I on grammar symbol X.
func (a *AutomatonCalculator) GOTO(I ItemSet, X grammar.Symbol) ItemSet {
	// Initialize J to be the empty set.
	J := NewItemSet()

	// For each item "A → αX•β" in I
	for i := range I.All() {
		if Y, ok := i.DotSymbol(); ok && Y.Equals(X) {
			// Add item "A → α•Xβ" to set J
			if next, ok := i.Next(); ok {
				J.Add(next)
			}
		}
	}

	// Compute CLOSURE(J)
	return a.CLOSURE(J)
}

// Canonical constructs the canonical collection of item sets for the augmented grammar G′.
//
// The canonical collection forms the basis for constructing an automaton, used for making parsing decisions.
// Each state of the automaton corresponds to an item set in the canonical collection.
func (a *AutomatonCalculator) Canonical() ItemSetCollection {
	// Initialize C to { CLOSURE(initial item) }.
	C := NewItemSetCollection(
		a.CLOSURE(NewItemSet(a.Initial())),
	)

	symbols := a.G().Symbols()

	for newItemSets := []ItemSet{}; newItemSets != nil; {
		newItemSets = nil

		// For each set of items I in C
		for I := range C.All() {
			// For each grammar symbol X
			for X := range symbols.All() {
				// If GOTO(I,X) is not empty and not in C, add GOTO(I,X) to C
				if J := a.GOTO(I, X); !J.IsEmpty() && !C.Contains(J) {
					newItemSets = append(newItemSets, J)
				}
			}
		}

		C.Add(newItemSets...)
	}

	return C
}
