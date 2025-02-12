// Package automata provides data structures and algorithms for working with automata.
//
// In language theory, automata refer to abstract computational models used to define and study formal languages.
// Automata are mathematical structures that process input strings and determine whether they belong to a specific language.
package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

var (
	eqState   = generic.NewEqualFunc[State]()
	cmpState  = generic.NewCompareFunc[State]()
	hashState = hash.HashFuncForInt[State](nil)

	eqSymbol   = generic.NewEqualFunc[Symbol]()
	cmpSymbol  = generic.NewCompareFunc[Symbol]()
	hashSymbol = hash.HashFuncForInt32[Symbol](nil)

	eqStateSet = func(a, b States) bool {
		return a.Equal(b)
	}

	eqSymbolState = func(a, b symboltable.SymbolTable[Symbol, State]) bool {
		return a.Equal(b)
	}

	eqSymbolStates = func(a, b symboltable.SymbolTable[Symbol, States]) bool {
		return a.Equal(b)
	}
)

// doubleKeyMap is a map (symbol table) data structure with two keys.
type doubleKeyMap[K1, K2, V any] symboltable.SymbolTable[K1, symboltable.SymbolTable[K2, V]]

// State represents a state in an automaton, identified by an integer.
type State int

// States represents a set of states in an automaton.
type States set.Set[State]

// NewStates creates a new set of states, initialized with the given states.
func NewStates(s ...State) States {
	return set.New(eqState, s...)
}

// sortStates sorts a set of states into a canonical order, ensuring a consistent representation.
func sortStates(states States) []State {
	sorted := generic.Collect1(states.All())
	sort.Quick(sorted, cmpState)

	return sorted
}

// Symbol represents an input symbol in an automaton, identified by a rune.
type Symbol rune

// E is the empty string ε and is never a member of an input alphabet Σ.
const E Symbol = 0

// Symbols represents a set of symbols in an automaton.
type Symbols set.Set[Symbol]

// NewSymbols creates a new set of symbols, initialized with the given symbols.
func NewSymbols(a ...Symbol) set.Set[Symbol] {
	return set.New(eqSymbol, a...)
}

// sortSymbols sorts a set of symbols into a canonical order, ensuring a consistent representation.
func sortSymbols(symbols Symbols) []Symbol {
	sorted := generic.Collect1(symbols.All())
	sort.Insertion(sorted, cmpSymbol)

	return sorted
}

// String represents a sequence of symbols in an automaton.
type String []Symbol

// toString creates a string of symbols from a string.
func toString(s string) String {
	res := make(String, len(s))
	for i, r := range s {
		res[i] = Symbol(r)
	}

	return res
}

// stateFactory is used for keeping track of states when combining multiple automata.
type stateFactory struct {
	last   State
	states map[int]map[State]State
}

func newStateFactory(last State) *stateFactory {
	return &stateFactory{
		last:   last,
		states: map[int]map[State]State{},
	}
}

func (f *stateFactory) StateFor(id int, s State) State {
	m, ok := f.states[id]
	if !ok {
		m = map[State]State{}
		f.states[id] = m
	}

	t, ok := m[s]
	if !ok {
		f.last++
		t = f.last
		m[s] = t
	}

	return t
}

// generatePermutations generates all permutations of a sequence of states using recursion and backtracking.
// Each permutation is passed to the provided yield function.
func generatePermutations(states []State, start, end int, yield func([]State) bool) bool {
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
