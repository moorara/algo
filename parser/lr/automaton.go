package lr

import "github.com/moorara/algo/grammar"

// Automaton represents the core of an LR automaton used in LR parsing.
// It provides essential functions for constructing LR parsing tables,
// including state transitions (GOTO) and generating item sets (Canonical).
//
// An LR automaton has two orthogonal dimensions:
//
//  1. The item dimension: it can be based on either LR(0) or LR(1) items.
//  2. The item set dimension: it can operate on either complete item sets or only kernel item sets.
type Automaton interface {
	// Initial returns the initial item of an augmented grammar.
	// For LR(0), the initial item is "S′ → •S",
	// and for LR(1), the initial item is "S′ → •S, $".
	Initial() Item

	// Closure computes the closure of a given item set.
	// It may compute the closure for either an LR(0) item set or an LR(1) item set.
	CLOSURE(ItemSet) ItemSet

	// GOTO(I, X) computes the closure of the set of all items "A → αX•β",
	// where "A → α•Xβ" is in the set of items I and X is a grammar symbol.
	//
	// The GOTO function defines transitions in the automaton for the grammar.
	// Each state of the automaton corresponds to a set of items, and
	// GOTO(I, X) specifies the transition from the state I on grammar symbol X.
	GOTO(ItemSet, grammar.Symbol) ItemSet

	// Canonical constructs the canonical collection of item sets for the augmented grammar G′.
	//
	// The canonical collection forms the basis for constructing an automaton, used for making parsing decisions.
	// Each state of the automaton corresponds to an item set in the canonical collection.
	Canonical() ItemSetCollection
}

// automaton implements the Automaton interface and considers both kernel and non-kernel items.
type automaton struct {
	G *grammar.CFG
	calculator
}

func (a *automaton) GOTO(I ItemSet, X grammar.Symbol) ItemSet {
	// Initialize J to be the empty set.
	J := NewItemSet()

	// For each item "A → αX•β" in I
	for i := range I.All() {
		if Y, ok := i.DotSymbol(); ok && Y.Equal(X) {
			// Add item "A → α•Xβ" to set J
			if next, ok := i.Next(); ok {
				J.Add(next)
			}
		}
	}

	// Compute CLOSURE(J)
	J = a.CLOSURE(J)

	return J
}

func (a *automaton) Canonical() ItemSetCollection {
	// Initialize C to { initial item set }
	C := NewItemSetCollection(
		a.CLOSURE(
			NewItemSet(a.Initial()),
		),
	)

	symbols := a.G.Symbols()

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

// kernelAutomaton implements the Automaton interface but considers only kernel items
type kernelAutomaton struct {
	G *grammar.CFG
	calculator
}

func (a *kernelAutomaton) GOTO(I ItemSet, X grammar.Symbol) ItemSet {
	// Initialize J to be the empty set.
	J := NewItemSet()

	// For each item "A → αX•β" in CLOSURE(I)
	for i := range a.CLOSURE(I).All() {
		if Y, ok := i.DotSymbol(); ok && Y.Equal(X) {
			// Add item "A → α•Xβ" to set J
			if next, ok := i.Next(); ok {
				J.Add(next)
			}
		}
	}

	return J
}

func (a *kernelAutomaton) Canonical() ItemSetCollection {
	// Initialize C to { initial item set }
	C := NewItemSetCollection(
		NewItemSet(a.Initial()),
	)

	symbols := a.G.Symbols()

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

// calculator defines the interface that specifies key functions for either LR(0) or LR(1) items.
type calculator interface {
	Initial() Item
	CLOSURE(ItemSet) ItemSet
}

// calculator0 implemented the Calculator interface for LR(0) items.
type calculator0 struct {
	G *grammar.CFG
}

// Initial returns the initial LR(0) item "S′ → •S" for an augmented grammar.
func (c *calculator0) Initial() Item {
	p, _ := c.G.Productions.Get(c.G.Start).FirstMatch(func(*grammar.Production) bool {
		return true
	})

	return &Item0{
		Production: p,
		Start:      c.G.Start,
		Dot:        0,
	}
}

// CLOSURE computes the closure of a given LR(0) item set.
func (c *calculator0) CLOSURE(I ItemSet) ItemSet {
	J := I.Clone()

	for newItems := []Item{}; newItems != nil; {
		newItems = nil

		// For each item A → α•Bβ in J
		for i := range J.All() {
			i := i.(*Item0)

			if X, ok := i.DotSymbol(); ok {
				if B, ok := X.(grammar.NonTerminal); ok {
					// For each production B → γ of G′
					for BProd := range c.G.Productions.Get(B).All() {
						j := &Item0{
							Production: BProd,
							Start:      c.G.Start,
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

		J.Add(newItems...)
	}

	return J
}

// calculator1 implemented the Calculator interface for LR(1) items.
type calculator1 struct {
	G     *grammar.CFG
	FIRST grammar.FIRST
}

// Initial returns the initial LR(1) item "S′ → •S, $" for an augmented grammar.
func (c *calculator1) Initial() Item {
	p, _ := c.G.Productions.Get(c.G.Start).FirstMatch(func(*grammar.Production) bool {
		return true
	})

	return &Item1{
		Production: p,
		Start:      c.G.Start,
		Dot:        0,
		Lookahead:  grammar.Endmarker,
	}
}

// CLOSURE computes the closure of a given LR(1) item set.
func (c *calculator1) CLOSURE(I ItemSet) ItemSet {
	J := I.Clone()

	for newItems := []Item{}; newItems != nil; {
		newItems = nil

		// For each item [A → α•Bβ, a] in J
		for i := range J.All() {
			i := i.(*Item1)
			a := i.Lookahead

			if X, ok := i.DotSymbol(); ok {
				if B, ok := X.(grammar.NonTerminal); ok {
					// For each production B → γ of G′
					for BProd := range c.G.Productions.Get(B).All() {
						β := i.GetSuffix()[1:]
						βa := β.Append(a)

						// For each terminal b in FIRST(βa)
						for b := range c.FIRST(βa).Terminals.All() {
							j := &Item1{
								Production: BProd,
								Start:      c.G.Start,
								Dot:        0,
								Lookahead:  b,
							}

							// If [B → •γ, b] is not in J
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
