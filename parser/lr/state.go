package lr

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

var (
	ErrState  = State(-1)
	EqState   = generic.NewEqualFunc[State]()
	HashState = hash.HashFuncForInt[State](nil)
)

// State represents a state in the LR parsing table or automaton.
type State int

// StateMap is a generic map that associates states (indexes) with their corresponding representations.
type StateMap[T generic.Equaler[T]] []T

// For finds the state corresponding to the given value v.
// Returns the state index if found, or ErrState if no match exists.
func (m StateMap[T]) For(v T) State {
	for i := range m {
		if m[i].Equals(v) {
			return State(i)
		}
	}

	return ErrState
}

// All returns all states in the map.
func (m StateMap[T]) All() []State {
	states := make([]State, len(m))
	for i := range m {
		states[i] = State(i)
	}

	return states
}
