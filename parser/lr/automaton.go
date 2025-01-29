package lr

import "github.com/moorara/algo/grammar"

var primeSuffixes = []string{
	"′", // Prime (U+2032)
	"″", // Double Prime (U+2033)
	"‴", // Triple Prime (U+2034)
	"⁗", // Quadruple Prime (U+2057)
}

// augment augments a context-free grammar G by creating a new start symbol S′
// and adding a production "S′ → S", where S is the original start symbol of G.
// This transformation is used to prepare grammars for LR parsing.
// The function clones the input grammar G, ensuring that the original grammar remains unmodified.
func augment(G *grammar.CFG) *grammar.CFG {
	augG := G.Clone()

	// A special symbol used to indicate the end of a string.
	augG.Terminals.Add(grammar.Endmarker)

	newS := augG.AddNewNonTerminal(G.Start, primeSuffixes...)
	augG.Start = newS
	augG.Productions.Add(&grammar.Production{
		Head: newS,
		Body: grammar.String[grammar.Symbol]{G.Start},
	})

	return augG
}

// calculator defines the interface required for an LR parser.
type calculator interface {
	// G returns the augmented context-free grammar.
	G() *grammar.CFG

	// Initial returns the initial item of an augmented grammar.
	// For LR(0), the initial item is "S′ → •S",
	// and for LR(1), the initial item is "S′ → •S, $".
	Initial() Item

	// Closure computes the closure of a given item set.
	// It may compute the closure for either an LR(0) item set or an LR(1) item set.
	CLOSURE(ItemSet) ItemSet
}

// calculator0 implemented the Calculator interface for LR(0) items.
type calculator0 struct {
	augG *grammar.CFG
}

// G returns the augmented context-free grammar.
func (c *calculator0) G() *grammar.CFG {
	return c.augG
}

// Initial returns the initial LR(0) item "S′ → •S" for an augmented grammar.
func (c *calculator0) Initial() Item {
	for p := range c.augG.Productions.Get(c.augG.Start).All() {
		return &Item0{
			Production: p,
			Start:      c.augG.Start,
			Dot:        0,
		}
	}

	// This will never be the case.
	return nil
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
					for BProd := range c.augG.Productions.Get(B).All() {
						j := &Item0{
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

		J.Add(newItems...)
	}

	return J
}

// calculator1 implemented the Calculator interface for LR(1) items.
type calculator1 struct {
	augG  *grammar.CFG
	FIRST grammar.FIRST
}

// G returns the augmented context-free grammar.
func (c *calculator1) G() *grammar.CFG {
	return c.augG
}

// Initial returns the initial LR(1) item "S′ → •S, $" for an augmented grammar.
func (c *calculator1) Initial() Item {
	for p := range c.augG.Productions.Get(c.augG.Start).All() {
		return &Item1{
			Production: p,
			Start:      c.augG.Start,
			Dot:        0,
			Lookahead:  grammar.Endmarker,
		}
	}

	// This will never be the case.
	return nil
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
					for BProd := range c.augG.Productions.Get(B).All() {
						β := i.GetSuffix()[1:]
						βa := β.Append(a)

						// For each terminal b in FIRST(βa)
						for b := range c.FIRST(βa).Terminals.All() {
							j := &Item1{
								Production: BProd,
								Start:      c.augG.Start,
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

// Automaton represents the core of an LR automaton used in LR parsing.
// It provides key functions for constructing LR parsing tables,
// including state transitions (GOTO) and generating item sets (Canonical).
type Automaton interface {
	calculator

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

// NewLR0Automaton creates an Automaton for LR(0) items that considers both kernel and non-kernel items.
// It provides an implementation of the CLOSURE function based on the LR(0) items of the augmented grammar.
func NewLR0Automaton(G *grammar.CFG) Automaton {
	augG := augment(G)

	return &automaton{
		calculator: &calculator0{
			augG: augG,
		},
	}
}

// NewLR1Automaton creates an Automaton for LR(1) items that considers both kernel and non-kernel items.
// It provides an implementation of the CLOSURE function based on the LR(1) items of the augmented grammar.
func NewLR1Automaton(G *grammar.CFG) Automaton {
	augG := augment(G)
	FIRST := augG.ComputeFIRST()

	return &automaton{
		calculator: &calculator1{
			augG:  augG,
			FIRST: FIRST,
		},
	}
}

// kernelAutomaton implements the Automaton interface but considers only kernel items
type kernelAutomaton struct {
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

// NewLR0Automaton creates an Automaton for LR(0) items that considers only kernel items.
// It provides an implementation of the CLOSURE function based on the LR(0) items of the augmented grammar.
func NewLR0KernelAutomaton(G *grammar.CFG) Automaton {
	augG := augment(G)

	return &kernelAutomaton{
		calculator: &calculator0{
			augG: augG,
		},
	}
}

// NewLR1Automaton creates an Automaton for LR(1) items that considers only kernel items.
// It provides an implementation of the CLOSURE function based on the LR(1) items of the augmented grammar.
func NewLR1KernelAutomaton(G *grammar.CFG) Automaton {
	augG := augment(G)
	FIRST := augG.ComputeFIRST()

	return &kernelAutomaton{
		calculator: &calculator1{
			augG:  augG,
			FIRST: FIRST,
		},
	}
}
