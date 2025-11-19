package automata

import "github.com/moorara/algo/set"

// group represents a subset of states within a partition.
// Each group is uniquely identified by a representative state.
type group struct {
	States
	Rep State
}

func eqGroup(a, b group) bool {
	return a.States.Equal(b.States)
}

func cmpGroup(a, b group) int {
	return CmpStates(a, b)
}

func hashGroup(g group) uint64 {
	return HashStates(g.States)
}

// groups represents a set of groups.
type groups set.Set[group]

func newGroups(gs ...group) groups {
	return set.NewSortedSet(cmpGroup, gs...)
}

// partition groups DFA states into disjoint sets, each tracked by a representative state.
type partition struct {
	groups
	nextRep State
}

// newPartition creates a new partition.
func newPartition() *partition {
	return &partition{
		groups:  newGroups(),
		nextRep: 0,
	}
}

// Equal implements the generic.Equaler interface.
func (p *partition) Equal(rhs *partition) bool {
	return p.groups.Equal(rhs.groups) && p.nextRep == rhs.nextRep
}

// Add inserts new groups into the partition set.
// Each group is a set of states and assigned a unique representative state.
func (p *partition) Add(groups ...States) {
	for _, states := range groups {
		p.groups.Add(group{
			States: states,
			Rep:    p.nextRep,
		})

		p.nextRep++
	}
}

// FindRep finds the group containing the given state and returns its representative state.
// If the state is not found in any group, it returns -1.
func (p *partition) FindRep(s State) State {
	for G := range p.groups.All() {
		if G.Contains(s) {
			return G.Rep
		}
	}

	return -1
}
