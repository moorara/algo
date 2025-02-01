package lr

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
)

var (
	primeSuffixes = []string{
		"′", // Prime (U+2032)
		"″", // Double Prime (U+2033)
		"‴", // Triple Prime (U+2034)
		"⁗", // Quadruple Prime (U+2057)
	}

	eqPrecedenceHandle = func(lhs, rhs *PrecedenceHandle) bool {
		return lhs.Equal(rhs)
	}
)

// augment augments a context-free grammar G by creating a new start symbol S′
// and adding a production "S′ → S", where S is the original start symbol of G.
// This transformation is used to prepare grammars for LR parsing.
// The function clones the input grammar G, ensuring the original grammar remains unmodified.
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

// NewGrammarWithLR0 creates a new augmented grammar based on the complete sets of LR(0) items.
func NewGrammarWithLR0(G *grammar.CFG, precedences PrecedenceLevels) *Grammar {
	G = augment(G)

	return &Grammar{
		CFG:              G,
		PrecedenceLevels: precedences,
		Automaton: &automaton{
			G: G,
			calculator: &calculator0{
				G: G,
			},
		},
	}
}

// NewGrammarWithLR1 creates a new augmented grammar based on the complete sets of LR(1) items.
func NewGrammarWithLR1(G *grammar.CFG, precedences PrecedenceLevels) *Grammar {
	G = augment(G)
	FIRST := G.ComputeFIRST()

	return &Grammar{
		CFG:              G,
		PrecedenceLevels: precedences,
		Automaton: &automaton{
			G: G,
			calculator: &calculator1{
				G:     G,
				FIRST: FIRST,
			},
		},
	}
}

// NewGrammarWithLR0Kernel creates a new augmented grammar based on the kernel sets of LR(0) items.
func NewGrammarWithLR0Kernel(G *grammar.CFG, precedences PrecedenceLevels) *Grammar {
	G = augment(G)

	return &Grammar{
		CFG:              G,
		PrecedenceLevels: precedences,
		Automaton: &kernelAutomaton{
			G: G,
			calculator: &calculator0{
				G: G,
			},
		},
	}
}

// NewGrammarWithLR1Kernel creates a new augmented grammar based on the kernel sets of LR(1) items.
func NewGrammarWithLR1Kernel(G *grammar.CFG, precedences PrecedenceLevels) *Grammar {
	G = augment(G)
	FIRST := G.ComputeFIRST()

	return &Grammar{
		CFG:              G,
		PrecedenceLevels: precedences,
		Automaton: &kernelAutomaton{
			G: G,
			calculator: &calculator1{
				G:     G,
				FIRST: FIRST,
			},
		},
	}
}

// Grammar represents a context-free grammar with additional features
// tailored for LR parsing and constructing an LR parser.
// The grammar is augmented with a new start symbol S′ and a new production "S′ → S"
// to facilitate the construction of the LR parser.
// It extends a regular context-free grammar by incorporating precedence and associativity information
// to handle ambiguities in the grammar and resolve conflicts.
// It also provides essential functionalities for building an LR automaton and the corresponding LR parsing table.
//
// Always use one of the provided functions to create a new instance of this type.
// Direct instantiation of the struct is discouraged.
type Grammar struct {
	*grammar.CFG
	PrecedenceLevels
	Automaton
}

// PrecedenceLevels represents an ordered list of precedence levels defined
// for specific terminals or production rules of a context-free grammar.
// The order of these levels is crucial for resolving conflicts.
type PrecedenceLevels []PrecedenceLevel

// Associativity represents the associativity property of a terminal or a production rule.
type Associativity int

const (
	NONE  Associativity = iota // Not associative
	LEFT                       // Left-associative
	RIGHT                      // Right-associative
)

// PrecedenceLevel defines a set of terminals and/or production rules (referred to as handles),
// each of which shares the same precedence level and associativity.
type PrecedenceLevel struct {
	Associativity
	PrecedenceHandles
}

// PrecedenceHandles represents a set of terminals and/or production rules (referred to as handles).
type PrecedenceHandles set.Set[*PrecedenceHandle]

// NewPrecedenceHandles creates a new set of terminals and/or production rules (referred to as handles).
func NewPrecedenceHandles(handles ...*PrecedenceHandle) PrecedenceHandles {
	return set.NewWithFormat(
		eqPrecedenceHandle,
		func(h []*PrecedenceHandle) string {
			if len(h) == 0 {
				return ""
			}

			sort.Insertion(h, cmpPrecedenceHandle)

			var b bytes.Buffer
			for i := range len(h) {
				fmt.Fprintf(&b, "%s, ", h[i])
			}
			b.Truncate(b.Len() - 2)

			return b.String()
		},
		handles...,
	)
}

// cmpPrecedenceHandles compares two sets of handles and establishes an order between them.
func cmpPrecedenceHandles(lhs, rhs PrecedenceHandles) int {
	if lhs.Size() < rhs.Size() {
		return -1
	} else if lhs.Size() > rhs.Size() {
		return 1
	}

	ls := generic.Collect1(lhs.All())
	sort.Quick(ls, cmpPrecedenceHandle)

	rs := generic.Collect1(rhs.All())
	sort.Quick(rs, cmpPrecedenceHandle)

	for i := range len(ls) {
		if cmp := cmpPrecedenceHandle(ls[i], rs[i]); cmp != 0 {
			return cmp
		}
	}

	return 0
}

// PrecedenceHandle represents either a terminal symbol or a production rule
// in the context of determining precedence for conflict resolution.
type PrecedenceHandle struct {
	*grammar.Terminal
	*grammar.Production
}

// IsTerminal returns true if the handle represents a terminal symbol.
func (h *PrecedenceHandle) IsTerminal() bool {
	return h.Terminal != nil && h.Production == nil
}

// IsProduction returns true if the handle represents a production rule.
func (h *PrecedenceHandle) IsProduction() bool {
	return h.Terminal == nil && h.Production != nil
}

// String returns a string representation of the handle.
func (h *PrecedenceHandle) String() string {
	if h.IsTerminal() {
		return h.Terminal.String()
	} else if h.IsProduction() {
		return fmt.Sprintf("%s = %s", h.Production.Head, h.Production.Body)
	}

	panic("PrecedenceHandle.String: invalid configuration")
}

// Equal determines whether or not two handles are the same.
func (h *PrecedenceHandle) Equal(rhs *PrecedenceHandle) bool {
	switch {
	case h.IsTerminal() && rhs.IsTerminal():
		return grammar.EqTerminal(*h.Terminal, *rhs.Terminal)
	case h.IsProduction() && rhs.IsProduction():
		return grammar.EqProduction(h.Production, rhs.Production)

	}

	return false
}

// cmpPrecedenceHandle compares two handles and establishes an order between them.
// Terminal handles come before Production handles.
func cmpPrecedenceHandle(lhs, rhs *PrecedenceHandle) int {
	switch {
	case lhs.IsTerminal() && rhs.IsProduction():
		return -1
	case lhs.IsProduction() && rhs.IsTerminal():
		return 1
	case lhs.IsTerminal() && rhs.IsTerminal():
		return grammar.CmpTerminal(*lhs.Terminal, *rhs.Terminal)
	case lhs.IsProduction() && rhs.IsProduction():
		return grammar.CmpProduction(lhs.Production, rhs.Production)
	}

	panic("cmpPrecedenceHandle: invalid configuration")
}

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
