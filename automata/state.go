package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/symboltable"
)

var (
	eqState  = generic.NewEqualFunc[State]()
	cmpState = generic.NewCompareFunc[State]()

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

// stateFactory is used for keeping track of states when combining multiple finite automata.
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
