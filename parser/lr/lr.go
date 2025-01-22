// Package lr provides common data structures and algorithms for building LR parsers.
// LR parsers are bottom-up parsers that analyse deterministic context-free languages in linear time.
//
// Bottom-up parsing constructs a parse tree for an input string
// starting at the leaves (bottom) and working towards the root (top).
// This process involves reducing a string w to the start symbol of the grammar.
// At each reduction step, a specific substring matching the body of a production
// is replaced by the non-terminal at the head of that production.
//
// Bottom-up parsing during a left-to-right scan of the inputconstructs a rightmost derivation in reverse:
//
//	S = γ₀ ⇒ᵣₘ γ₁ ⇒ᵣₘ γ₂ ⇒ᵣₘ ... ⇒ᵣₘ γₙ₋₁ ⇒ᵣₘ γₙ = w
//
// At each step, the handle βₙ in γₙ is replaced by the head of the production Aₙ → βₙ
// to obtain the previous right-sentential form γₙ₋₁.
// If the process produces the start symbol S as the only sentential form, parsing is complete.
// If a grammar is unambiguous, then every right-sentential form of the grammar has exactly one handle.
//
// The most common type of bottom-up parser is LR(k) parsing.
// The L is for left-to-right scanning of the input, the R for constructing a rightmost derivation in reverse,
// and the k for the number of input symbols of lookahead that are used in making parsing decisions.
//
// Advantages of LR parsing:
//
//   - Can recognize nearly all programming language constructs defined by context-free grammars.
//   - Detects syntax errors at the earliest possible point during a left-to-right scan.
//   - The class of grammars that can be parsed using LR methods is a proper superset of
//     the class of grammars that can be parsed with predictive or LL methods.
//     For a grammar to be LR(k), we must be able to recognize the occurrence of the right side of
//     a production in a right-sentential form, with k input symbols of lookahead.
//     This requirement is far less stringent than that for LL(k) grammars where we must be able
//     to recognize the use of a production seeing only the first k symbols of what its right side derives.
//
// In LR(k) parsing, the cases k = 0 or k = 1 are most commonly used in practical applications.
// LR parsing methods use pushdown automata (PDA) to parse an input string.
// A pushdown automaton is a type of automaton used for Type 2 languages (context-free languages) in the Chomsky hierarchy.
// A PDA uses a state machine with a stack.
// The next state is determined by the current state, the next input, and the top of the stack.
// LR(0) parsers do not rely on any lookahead to make parsing decisions.
// An LR(0) parser bases its decisions entirely on the current state and the parsing stack.
// LR(1) parsers determine the next state based on the current state, one lookahead symbol, and the top of the stack.
//
// Shift-reduce parsing is a bottom-up parsing technique that uses
// a stack for grammar symbols and an input buffer for the remaining string.
// The parser alternates between shifting symbols from the input to the stack
// and reducing the top of the stack based on grammar rules.
// This process continues until the stack contains only the start symbol and the input is empty, or an error occurs.
//
// Certain context-free grammars cannot be parsed using shift-reduce parsers
// because they may encounter shift/reduce conflicts (indecision between shifting or reducing)
// or reduce/reduce conflicts (indecision between multiple reductions).
// Technically speaking, these grammars are not in the LR(k) class.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
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
	S := G.Start
	augG := G.Clone()

	newS := augG.AddNewNonTerminal(S, primeSuffixes...)
	augG.Start = newS
	newP := grammar.Production{
		Head: newS,
		Body: grammar.String[grammar.Symbol]{S},
	}
	augG.Productions.Add(newP)

	return augG
}

// Closure computes the closure of a given item set.
// A clousre function may compute the closure for either an LR(0) item set or an LR(1) item set.
type ClosureFunc func(ItemSet) ItemSet

// GOTO(I, X) computes the closure of the set of all items "A → αX•β",
// where "A → α•Xβ" is in the set of items I and X is a grammar symbol.
//
// The GOTO function defines transitions in the automaton for the grammar.
// Each state of the automaton corresponds to a set of items, and
// GOTO(I, X) specifies the transition from the state I on grammar symbol X.
func GOTO(CLOSURE ClosureFunc, I ItemSet, X grammar.Symbol) ItemSet {
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
	return CLOSURE(J)
}

// Canonical constructs the canonical collection of item sets for the augmented grammar G′.
//
// The canonical collection forms the basis for constructing
// an automaton, used for making parsing decisions.
// Each state of the automaton corresponds to an item set in the canonical collection.
func Canonical(augG grammar.CFG, initial Item, CLOSURE ClosureFunc) ItemSetCollection {
	// Initialize C to { CLOSURE(initial item) }.
	C := NewItemSetCollection(
		CLOSURE(NewItemSet(initial)),
	)

	for newItemSets := []ItemSet{}; newItemSets != nil; {
		newItemSets = nil

		// For each set of items I in C
		for I := range C.All() {
			// For each grammar symbol X
			for X := range augG.Symbols().All() {
				// If GOTO(I,X) is not empty and not in C, add GOTO(I,X) to C
				if J := GOTO(CLOSURE, I, X); !J.IsEmpty() && !C.Contains(J) {
					newItemSets = append(newItemSets, J)
				}
			}
		}

		C.Add(newItemSets...)
	}

	return C
}
