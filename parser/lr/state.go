package lr

import (
	"iter"

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

// StateMap represents a two-leveled indexed map where the first index (state) maps to an item set,
// and within each item set, the second index maps to individual items.
type StateMap [][]Item

// BuildStateMap constructs a deterministic mapping of item sets and individual items.
// It creates a StateMap where the first index (state) corresponds to an item set,
// and the second index maps to individual items within each set, sorted in a consistent order.
func BuildStateMap(C ItemSetCollection) StateMap {
	sets := generic.Collect1(C.All())
	sort.Quick(sets, cmpItemSet)

	m := make([][]Item, len(sets))

	for i, set := range sets {
		m[i] = generic.Collect1(set.All())
		sort.Quick(m[i], CmpItem)
	}

	return m
}

// Item returns the item at the specified index i within the item set indexed by state s.
func (m StateMap) Item(s State, i int) Item {
	return m[s][i]
}

// ItemSet returns the item set associated with the specified state s.
// It creates a new ItemSet from the items corresponding to the given state.
func (m StateMap) ItemSet(s State) ItemSet {
	return NewItemSet(m[s]...)
}

// FindItemSet finds the state corresponding to a given item set.
// It returns the state if found, or ErrState if no match exists.
func (m StateMap) FindItem(s State, item Item) int {
	for i, it := range m[s] {
		if it.Equal(item) {
			return i
		}
	}

	return -1
}

// FindItemSet finds the state corresponding to a given item set.
// It returns the state if found, or ErrState if no match exists.
func (m StateMap) FindItemSet(I ItemSet) State {
	for i := range m {
		if s := State(i); m.equalToItemSet(s, I) {
			return s
		}
	}

	return ErrState
}

func (m StateMap) equalToItemSet(s State, I ItemSet) bool {
	for _, it := range m[s] {
		if !I.Contains(it) {
			return false
		}
	}

	for it := range I.All() {
		if m.FindItem(s, it) == -1 {
			return false
		}
	}

	return true
}

// States returns all states in the map.
func (m StateMap) States() []State {
	states := make([]State, len(m))
	for i := range m {
		states[i] = State(i)
	}

	return states
}

// All returns an iterator sequence that contains all state-ItemSet pairs in the StateMap.
func (m StateMap) All() iter.Seq2[State, ItemSet] {
	return func(yield func(State, ItemSet) bool) {
		for i := range m {
			if s := State(i); !yield(s, m.ItemSet(s)) {
				return
			}
		}
	}
}

// String returns a string representation of all states in the map.
func (m StateMap) String() string {
	cs := &itemSetCollectionStringer{
		sets: make([]ItemSet, len(m)),
	}

	for i, items := range m {
		cs.sets[i] = NewItemSet(items...)
	}

	return cs.String()
}
