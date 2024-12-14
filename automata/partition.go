package automata

import (
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/symboltable"
)

// group represents a subset of states within a partition.
// Each group is uniquely identified by a representative state.
type group struct {
	rep    State
	states States
}

// Equals determines whether or not two partition groups are equal.
func (g group) Equals(rhs group) bool {
	return g.states.Equals(rhs.states)
}

// groups represents a list of partition groups.
type groups []group

// Contains determines whether or not a list of partition groups contains a given group.
func (g groups) Contains(h group) bool {
	for _, group := range g {
		if group.Equals(h) {
			return true
		}
	}

	return false
}

// Equals determines whether or not two lists of partition groups are equal.
func (g groups) Equals(rhs groups) bool {
	for _, group := range g {
		if !rhs.Contains(group) {
			return false
		}
	}

	for _, group := range rhs {
		if !g.Contains(group) {
			return false
		}
	}

	return true
}

// partition represents a partition of the states in an automaton.
// Each partition divides the states into disjoint groups.
type partition struct {
	groups
	nextRep State
}

// newPartition creates a new partition set.
func newPartition() *partition {
	return &partition{
		groups:  groups{},
		nextRep: 0,
	}
}

// Equals determines whether or not two partitions are equal.
func (p *partition) Equals(rhs *partition) bool {
	return p.groups.Equals(rhs.groups)
}

// Add adds new groups into the partition set.
// Each group is a set of states and assigned a unique representative state.
func (p *partition) Add(groups ...States) {
	for _, groupStates := range groups {
		p.groups = append(p.groups, group{
			rep:    p.nextRep,
			states: groupStates,
		})

		p.nextRep++
	}
}

// Rep finds the group containing the given state and returns its representative state.
// If the state is not found in any group, it returns -1 and false.
func (p *partition) Rep(s State) (State, bool) {
	for _, G := range p.groups {
		if G.states.Contains(s) {
			return G.rep, true
		}
	}

	return -1, false
}

// BuildGroupTrans creates a map of state-symbol pairs to the group representatives in the current partition.
// Instead of mapping to next states, the map associates each symbol with the representative of the group containing the next state.
func (p *partition) BuildGroupTrans(dfa *DFA, G group) doubleKeyMap[State, Symbol, State] {
	Gtrans := symboltable.NewRedBlack[State, symboltable.SymbolTable[Symbol, State]](cmpState, eqSymbolState)

	// For every state in the current group
	for _, s := range G.states {
		strans := symboltable.NewRedBlack[Symbol, State](cmpSymbol, eqState)

		// Create a map of symbols to the current partition's group representatives (instead of next states)
		if v, ok := dfa.trans.Get(s); ok {
			for a, next := range v.All() {
				if rep, ok := p.Rep(next); ok {
					strans.Put(a, rep)
				}
			}
		}

		Gtrans.Put(s, strans)
	}

	return Gtrans
}

// PartitionAndAddGroups partitions each group into subgroups and adds them to the partition set.
//
// This method partitions the group G into subgroups such that two states s and t are in the same subgroup
// if and only if, for all input symbols a, the transitions of s and t on a lead to states in the same group.
// If no such grouping is possible, a state will be placed in a subgroup by itself.
func (p *partition) PartitionAndAddGroups(Gtrans doubleKeyMap[State, Symbol, State]) {
	pairs := Collect(Gtrans.All())

	for i := 0; i < len(pairs); i++ {
		s, strans := pairs[i].Key, pairs[i].Val

		// If s is not already added to the new partition
		if _, ok := p.Rep(s); !ok {
			// Create a new group in the new partition
			H := States{}
			H = append(H, s)

			// Add all other state with same symbol->rep map to the new group
			for j := 1; j < len(pairs); j++ {
				t, ttrans := pairs[j].Key, pairs[j].Val

				if strans.Equals(ttrans) && !H.Contains(t) {
					H = append(H, t)
				}
			}

			p.Add(H)
		}
	}
}
