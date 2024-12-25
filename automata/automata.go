// Package automata provides data structures and algorithms for working with finite automata,
// also known as finite state machines.
package automata

import (
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/symboltable"
)

var (
	cmpState  = NewCompareFunc[State]()
	cmpSymbol = NewCompareFunc[Symbol]()

	eqState  = NewEqualFunc[State]()
	eqStates = func(a, b States) bool {
		return a.Equals(b)
	}

	eqSymbolState = func(a, b symboltable.SymbolTable[Symbol, State]) bool {
		return a.Equals(b)
	}

	eqSymbolStates = func(a, b symboltable.SymbolTable[Symbol, States]) bool {
		return a.Equals(b)
	}
)

// doubleKeyMap is a map (symbol table) data structure with two keys.
type doubleKeyMap[K1, K2, V any] symboltable.SymbolTable[K1, symboltable.SymbolTable[K2, V]]

// State represents a state in a finite automaton.
type State int

// States represents a collection of states in a finite automaton.
type States []State

// Contains determines whether or not a set of states contains a given state.
func (s States) Contains(t State) bool {
	for _, state := range s {
		if state == t {
			return true
		}
	}

	return false
}

// Equals determines whether or not two sets of states are equal.
func (s States) Equals(rhs States) bool {
	for _, state := range s {
		if !rhs.Contains(state) {
			return false
		}
	}

	for _, state := range rhs {
		if !s.Contains(state) {
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

// Equals determines whether or not two sets of symbols are equal.
func (s Symbols) Equals(rhs Symbols) bool {
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

// generatePermutations generates all permutations of a sequence of states using recursion and backtracking.
// Each permutation is passed to the provided yield function.
func generatePermutations(states States, start, end int, yield func(States) bool) bool {
	if start == end {
		return yield(states)
	}

	for i := start; i <= end; i++ {
		states[start], states[i] = states[i], states[start]
		cont := generatePermutations(states, start+1, end, yield)
		states[start], states[i] = states[i], states[start]

		if !cont {
			return false
		}
	}

	return true
}
