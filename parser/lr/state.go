package lr

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/sort"
)

const ErrState = State(-1)

var (
	EqState   = generic.NewEqualFunc[State]()
	HashState = hash.HashFuncForInt[State](nil)

	CmpState = func(lhs, rhs State) int {
		return int(lhs) - int(rhs)
	}
)

// State represents a state in the LR parsing table or automaton.
type State int

// StateMap is a generic map that associates an state (index) with an item set.
type StateMap []ItemSet

// BuildStateMap constructs a deterministic mapping of states to item sets.
// It creates a StateMap that associates a state (index) with each item set in the collection.
func BuildStateMap(C ItemSetCollection) StateMap {
	states := generic.Collect1(C.All())
	sort.Quick(states, cmpItemSet)

	return states
}

// Find finds the state corresponding to a given item set.
// Returns the state if found, or ErrState if no match exists.
func (m StateMap) Find(I ItemSet) State {
	for s := range m {
		if m[s].Equal(I) {
			return State(s)
		}
	}

	return ErrState
}

// States returns all states in the map.
func (m StateMap) States() []State {
	states := make([]State, len(m))
	for i := range m {
		states[i] = State(i)
	}

	return states
}

// String returns a string representation of all states in the map.
func (m StateMap) String() string {
	cs := &itemSetCollectionStringer{
		sets: m,
	}

	return cs.String()
}
