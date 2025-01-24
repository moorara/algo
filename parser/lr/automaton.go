package lr

import "github.com/moorara/algo/grammar"

// augment augments a context-free grammar G by creating a new start symbol S′
// and adding a production "S′ → S", where S is the original start symbol of G.
// This transformation is used to prepare grammars for LR parsing.
// The function clones the input grammar G, ensuring that the original grammar remains unmodified.
func augment(G *grammar.CFG) *grammar.CFG {
	augG := G.Clone()

	newS := augG.AddNewNonTerminal(G.Start, primeSuffixes...)
	augG.Start = newS
	augG.Productions.Add(&grammar.Production{
		Head: newS,
		Body: grammar.String[grammar.Symbol]{G.Start},
	})

	return augG
}

// Automaton represents the core of an LR automaton used in LR parsing.
// It provides key functions for constructing LR parsing tables,
// including state transitions (GOTO) and generating item sets (Canonical).
type Automaton interface {
	Calculator

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

// automaton implements the Automaton interface.
type automaton struct {
	Calculator
}

func (a *automaton) GOTO(I ItemSet, X grammar.Symbol) ItemSet {
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
	J = a.CLOSURE(J)

	return J
}

func (a *automaton) Canonical() ItemSetCollection {
	I0 := NewItemSet(a.Initial())
	I0 = a.CLOSURE(I0)

	// Initialize C to { initial item set }
	C := NewItemSetCollection(I0)

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

// NewLR0Automaton returns an Automaton for LR(0) items.
// It provides an implementation of the CLOSURE function
// based on the LR(0) items of the augmented grammar.
func NewLR0Automaton(G *grammar.CFG) Automaton {
	augG := augment(G)

	return &automaton{
		Calculator: &calculator0{
			augG: augG,
		},
	}
}

// NewLR1Automaton returns an Automaton for LR(1) items.
// It provides an implementation of the CLOSURE function
// based on the LR(1) items of the augmented grammar.
func NewLR1Automaton(G *grammar.CFG) Automaton {
	augG := augment(G)
	FIRST := augG.ComputeFIRST()

	return &automaton{
		Calculator: &calculator1{
			augG:  augG,
			FIRST: FIRST,
		},
	}
}
