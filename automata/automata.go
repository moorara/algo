// Package automata provides data structures and algorithms for working with automata.
//
// In language theory, automata refer to abstract computational models used to define and study formal languages.
// Automata are mathematical structures that process input strings and determine whether they belong to a specific language.
package automata

import (
	"fmt"

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
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		return a.Equal(b)
	}

	eqSymbols = func(a, b Symbols) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

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

// SymbolRange represents an inclusive range of input symbols.
type SymbolRange struct {
	Start Symbol
	End   Symbol
}

func (r SymbolRange) Validate() {
	if r.End < r.Start {
		panic(fmt.Sprintf("invalid symbol range [%c-%c]", r.Start, r.End))
	}
}

// String implements the fmt.Stringer interface.
func (r SymbolRange) String() string {
	var start, end string

	if r.Start == E {
		start = "ε"
	} else {
		start = fmt.Sprintf("%c", r.Start)
	}

	if r.End == E {
		end = "ε"
	} else {
		end = fmt.Sprintf("%c", r.End)
	}

	if start == end {
		return fmt.Sprintf("[%s]", start)
	}

	return fmt.Sprintf("[%s..%s]", start, end)
}

// Equal implements the generic.Equaler interface.
func (r SymbolRange) Equal(rhs SymbolRange) bool {
	return r.Start == rhs.Start && r.End == rhs.End
}
