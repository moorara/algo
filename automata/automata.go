// Package automata provides data structures and algorithms for working with finite automata, a.k.a. finite state machines.
package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/symboltable"
)

// State represents a state in a finite automaton.
type State int

// States represents a collection of states in a finite automaton.
type States []State

// Contains determines whether or not a set of states contains a given state.
func (s States) Contains(t State) bool {
	for _, state := range s {
		if t == state {
			return true
		}
	}

	return false
}

// Equals determines whether or not two sets of states are equal.
func (s States) Equals(t States) bool {
	for _, u := range s {
		if !t.Contains(u) {
			return false
		}
	}

	for _, u := range t {
		if !s.Contains(u) {
			return false
		}
	}

	return true
}

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

// String represents an input string in a finite automaton.
type String []Symbol

func ToString(s string) String {
	res := make(String, len(s))
	for i, r := range s {
		res[i] = Symbol(r)
	}

	return res
}

var (
	cmpState  = generic.NewCompareFunc[State]()
	cmpSymbol = generic.NewCompareFunc[Symbol]()

	eqState  = generic.NewEqualFunc[State]()
	eqStates = func(a, b States) bool {
		return a.Equals(b)
	}

	eqSymbolState = func(a, b symboltable.OrderedSymbolTable[Symbol, State]) bool {
		return a.Equals(b)
	}

	eqSymbolStates = func(a, b symboltable.OrderedSymbolTable[Symbol, States]) bool {
		return a.Equals(b)
	}
)
