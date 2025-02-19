package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

// group represents a subset of states within a partition.
// Each group is uniquely identified by a representative state.
type group struct {
	States
	rep State
}

func eqGroup(lhs, rhs group) bool {
	return lhs.Equal(rhs)
}

// groups represents a set of groups.
type groups set.Set[group]

// newGroups creates a new set of groups.
func newGroups(g ...group) groups {
	return set.NewStable(eqGroup, g...)
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
		groups:  newGroups(),
		nextRep: 0,
	}
}

// Equal determines whether or not two partitions are the same.
func (p *partition) Equal(rhs *partition) bool {
	return p.groups.Equal(rhs.groups) && p.nextRep == rhs.nextRep
}

// Add adds new groups into the partition set.
// Each group is a set of states and assigned a unique representative state.
func (p *partition) Add(groups ...States) {
	for _, states := range groups {
		p.groups.Add(group{
			States: states,
			rep:    p.nextRep,
		})

		p.nextRep++
	}
}

// Rep finds the group containing the given state and returns its representative state.
// If the state is not found in any group, it returns -1.
func (p *partition) Rep(s State) State {
	for G := range p.groups.All() {
		if G.Contains(s) {
			return G.rep
		}
	}

	return State(-1)
}

// BuildGroupTrans creates a map of state-symbol pairs to the group representatives in the current partition.
// Instead of mapping to next states, the map associates each symbol with the representative of the group containing the next state.
func (p *partition) BuildGroupTrans(dfa *DFA, G group) symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, State]] {
	Gtrans := symboltable.NewRedBlack(cmpState, eqSymbolState)

	// For every state in the current group
	for s := range G.States.All() {
		Gstrans := symboltable.NewRedBlack(cmpSymbol, eqState)

		// Create a map of symbols to the current partition's group representatives (instead of next states)
		if strans, ok := dfa.Trans.Get(s); ok {
			for a, next := range strans.All() {
				if rep := p.Rep(next); rep != -1 {
					Gstrans.Put(a, rep)
				}
			}
		}

		Gtrans.Put(s, Gstrans)
	}

	return Gtrans
}

// PartitionAndAddGroups partitions each group into subgroups and adds them to the partition set.
//
// This method partitions the group G into subgroups such that two states s and t are in the same subgroup
// if and only if, for all input symbols a, the transitions of s and t on a lead to states in the same group.
// If no such grouping is possible, a state will be placed in a subgroup by itself.
func (p *partition) PartitionAndAddGroups(Gtrans symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, State]]) {
	pairs := generic.Collect2(Gtrans.All())

	for i := 0; i < len(pairs); i++ {
		s, strans := pairs[i].Key, pairs[i].Val

		// If s is not already added to the new partition
		if p.Rep(s) == -1 {
			// Create a new group in the new partition
			H := NewStates(s)

			// Add all other state with same symbol->rep map to the new group
			for j := 1; j < len(pairs); j++ {
				t, ttrans := pairs[j].Key, pairs[j].Val

				if strans.Equal(ttrans) && !H.Contains(t) {
					H.Add(t)
				}
			}

			p.Add(H)
		}
	}
}
