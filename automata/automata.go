// Package automata provides data structures and algorithms for working with automata.
//
// In language theory, automata refer to abstract computational models used to define and study formal languages.
// Automata are mathematical structures that process input strings and determine whether they belong to a specific language.
package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/set"
)

var (
	EqState   = generic.NewEqualFunc[State]()
	CmpState  = generic.NewCompareFunc[State]()
	HashState = hash.HashFuncForInt[State](nil)

	EqSymbol   = generic.NewEqualFunc[Symbol]()
	CmpSymbol  = generic.NewCompareFunc[Symbol]()
	HashSymbol = hash.HashFuncForInt32[Symbol](nil)

	eqClassID   = generic.NewEqualFunc[classID]()
	cmpClassID  = generic.NewCompareFunc[classID]()
	hashClassID = hash.HashFuncForInt[classID](nil)

	EqStates = func(a, b States) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		return a.Equal(b)
	}

	CmpStates = func(a, b States) int {
		if a == nil && b == nil {
			return 0
		} else if a == nil {
			return -1
		} else if b == nil {
			return 1
		}

		// Assume a and b are sorted.
		lhs := generic.Collect1(a.All())
		rhs := generic.Collect1(b.All())

		for i := 0; i < len(lhs) && i < len(rhs); i++ {
			if c := CmpState(lhs[i], rhs[i]); c != 0 {
				return c
			}
		}

		return len(lhs) - len(rhs)
	}

	HashStates = func(ss States) uint64 {
		// Use XOR which is commutative and associative, so insertion order does not matter.
		var h uint64
		for s := range ss.All() {
			h ^= HashState(s)
		}

		return h
	}

	EqSymbols = func(a, b Symbols) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		return a.Equal(b)
	}

	CmpSymbols = func(a, b Symbols) int {
		if a == nil && b == nil {
			return 0
		} else if a == nil {
			return -1
		} else if b == nil {
			return 1
		}

		// Assume a and b are sorted.
		lhs := generic.Collect1(a.All())
		rhs := generic.Collect1(b.All())

		for i := 0; i < len(lhs) && i < len(rhs); i++ {
			if c := CmpSymbol(lhs[i], rhs[i]); c != 0 {
				return c
			}
		}

		return len(lhs) - len(rhs)
	}

	HashSymbols = func(ss Symbols) uint64 {
		// Use XOR which is commutative and associative, so insertion order does not matter.
		var h uint64
		for s := range ss.All() {
			h ^= HashSymbol(s)
		}

		return h
	}

	unionStates = func(a, b States) States {
		if a == nil && b == nil {
			return nil
		} else if a == nil {
			return b
		} else if b == nil {
			return a
		}

		return a.Union(b)
	}
)

// State represents a state in an automaton, identified by an integer.
type State int

// States represents a set of states in an automaton.
type States set.Set[State]

// NewStates creates a new set of states, initialized with the given states.
func NewStates(s ...State) States {
	return set.NewSortedSet(CmpState, s...)
}

// Symbol represents an input symbol in an automaton, identified by a rune.
type Symbol rune

// E is the empty string ε and is never a member of an input alphabet Σ.
// It is intentionally chosen outside the Unicode range to avoid conflicts with valid symbols.
const E Symbol = -1

// Symbols represents a set of symbols in an automaton.
type Symbols set.Set[Symbol]

// NewSymbols creates a new set of symbols, initialized with the given symbols.
func NewSymbols(a ...Symbol) set.Set[Symbol] {
	return set.NewSortedSet(CmpSymbol, a...)
}

// String represents a sequence of symbols in an automaton.
type String []Symbol

// classID is used to identify equivalence classes of input symbols.
type classID int
