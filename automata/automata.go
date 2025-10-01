// Package automata provides data structures and algorithms for working with automata.
//
// In language theory, automata refer to abstract computational models used to define and study formal languages.
// Automata are mathematical structures that process input strings and determine whether they belong to a specific language.
package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

var (
	EqState   = generic.NewEqualFunc[State]()
	CmpState  = generic.NewCompareFunc[State]()
	HashState = hash.HashFuncForInt[State](nil)

	EqSymbol   = generic.NewEqualFunc[Symbol]()
	CmpSymbol  = generic.NewCompareFunc[Symbol]()
	HashSymbol = hash.HashFuncForInt32[Symbol](nil)

	eqStates = func(a, b States) bool {
		return a.Equal(b)
	}

	eqSymbolState = func(a, b symboltable.SymbolTable[Symbol, State]) bool {
		return a.Equal(b)
	}

	eqSymbolStates = func(a, b symboltable.SymbolTable[Symbol, States]) bool {
		return a.Equal(b)
	}
)

// State represents a state in an automaton, identified by an integer.
type State int

// States represents a set of states in an automaton.
type States set.Set[State]

// NewStates creates a new set of states, initialized with the given states.
func NewStates(s ...State) States {
	return set.NewSorted(CmpState, s...)
}

// Symbol represents an input symbol in an automaton, identified by a rune.
type Symbol rune

// E is the empty string ε and is never a member of an input alphabet Σ.
const E Symbol = 0

// Symbols represents a set of symbols in an automaton.
type Symbols set.Set[Symbol]

// NewSymbols creates a new set of symbols, initialized with the given symbols.
func NewSymbols(a ...Symbol) set.Set[Symbol] {
	return set.NewSorted(CmpSymbol, a...)
}

// String represents a sequence of symbols in an automaton.
type String []Symbol

// SymbolRange represents a range of symbols inclusive.
type SymbolRange struct {
	Start Symbol
	End   Symbol
}
