package automata

import "github.com/moorara/algo/generic"

var (
	cmpSymbol = generic.NewCompareFunc[Symbol]()
)

// Symbol represents an input symbol in a finite automaton.
type Symbol rune

// E is the empty string ε and is never a member of an input alphabet Σ.
const E Symbol = 0

// Symbols represents a collection of symbols in a finite automaton.
type Symbols []Symbol

// Contains determines whether or not a set of symbols contains a given symbol.
func (s Symbols) Contains(t Symbol) bool {
	for _, symbol := range s {
		if t == symbol {
			return true
		}
	}

	return false
}

// Equal determines whether or not two sets of symbols are equal.
func (s Symbols) Equal(rhs Symbols) bool {
	for _, symbol := range s {
		if !rhs.Contains(symbol) {
			return false
		}
	}

	for _, symbol := range rhs {
		if !s.Contains(symbol) {
			return false
		}
	}

	return true
}

// String represents an input string in a finite automaton.
type String []Symbol

func ToString(s string) String {
	res := make(String, len(s))
	for i, r := range s {
		res[i] = Symbol(r)
	}

	return res
}
