package automata

import (
	"bytes"
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

/* ------------------------------------------------------------------------------------------------------------------------ */

// dfaTransitionEnds represents the states involved in a DFA transition.
type dfaTransitionEnds struct {
	State State
	Next  State
}

var cmpDFATransitionEnds = func(lhs, rhs dfaTransitionEnds) int {
	if c := CmpState(lhs.State, rhs.State); c != 0 {
		return c
	}

	return CmpState(lhs.Next, rhs.Next)
}

// dfaTransition represents a transition from one state to another state on a range of input symbols.
type dfaTransition struct {
	Ends  dfaTransitionEnds
	Range disc.Range[Symbol]
}

// dfaBound represents a boundary (start or end) of a range associated with a DFA transition.
type dfaBound struct {
	Value      Symbol
	Start      bool
	Transition dfaTransition
}

// dfaTransitionVector represents a DFA transition vector.
type dfaTransitionVector set.Set[dfaTransitionEnds]

func newDFATransitionVector() dfaTransitionVector {
	return set.NewSorted(cmpDFATransitionEnds)
}

var cmpDFATransitionVector = func(lhs, rhs dfaTransitionVector) int {
	v1 := generic.Collect1(lhs.All())
	v2 := generic.Collect1(rhs.All())

	for i := 0; i < len(v1) && i < len(v2); i++ {
		if c := cmpDFATransitionEnds(v1[i], v2[i]); c != 0 {
			return c
		}
	}

	return len(v1) - len(v2)
}

/* ------------------------------------------------------------------------------------------------------------------------ */

var eqClassIDStateTable = func(a, b symboltable.SymbolTable[classID, State]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Equal(b)
}

// dfaTransitionTable represents the transition table of a DFA.
type dfaTransitionTable struct {
	// Use an ordered symbol table so iterations over states and classes are deterministic and
	// the resulting computation and textual output are reproducible.
	symboltable.SymbolTable[State, symboltable.SymbolTable[classID, State]]
}

// newDFATransitionTable creates a new instance of dfaTransitionTable.
func newDFATransitionTable() *dfaTransitionTable {
	return &dfaTransitionTable{
		symboltable.NewRedBlack(CmpState, eqClassIDStateTable),
	}
}

// Clone implements the generic.Cloner interface.
func (t *dfaTransitionTable) Clone() *dfaTransitionTable {
	clone := newDFATransitionTable()

	for s, stab := range t.All() {
		stabClone := symboltable.NewRedBlack(cmpClassID, EqState)
		for cid, next := range stab.All() {
			stabClone.Put(cid, next)
		}
		clone.Put(s, stabClone)
	}

	return clone
}

// Equal implements the generic.Equaler interface.
func (t *dfaTransitionTable) Equal(rhs *dfaTransitionTable) bool {
	return t.SymbolTable.Equal(rhs.SymbolTable)
}

// Add inserts a new transition into the DFA transition table.
func (t *dfaTransitionTable) Add(s State, cid classID, next State) *dfaTransitionTable {
	stab, ok := t.Get(s)
	if !ok {
		stab = symboltable.NewRedBlack(cmpClassID, EqState)
		t.Put(s, stab)
	}

	stab.Put(cid, next)

	return t
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// DFABuilder implements the Builder design pattern for constructing DFA instances.
//
// The Builder pattern separates the construction of a DFA from its representation.
// This approach ensures the resulting DFA is immutable and optimized for simulation (running).
type DFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, disc.RangeMap[Symbol, State]]
}

// NewDFABuilder creates a new DFA builder instance.
func NewDFABuilder() *DFABuilder {
	return &DFABuilder{
		trans: symboltable.NewRedBlack[State, disc.RangeMap[Symbol, State]](CmpState, nil),
	}
}

// SetStart sets the start state of the DFA.
func (b *DFABuilder) SetStart(s State) *DFABuilder {
	b.start = s
	return b
}

// SetFinal sets the final (accepting) states of the DFA.
func (b *DFABuilder) SetFinal(f []State) *DFABuilder {
	b.final = NewStates(f...)
	return b
}

// AddTransition adds transitions from state s to state next on all input symbols in the range [start, end].
func (b *DFABuilder) AddTransition(s State, start, end Symbol, next State) *DFABuilder {
	ranges, ok := b.trans.Get(s)
	if !ok {
		opts := &disc.RangeMapOpts[Symbol, State]{}
		ranges = disc.NewRangeMap(EqState, opts)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		next,
	)

	return b
}

// Build constructs the DFA based on the configurations provided to the builder.
//
// This method reduces the size of the alphabet by computing the equivalence classes of input symbols based on the transition function.
// The resulting DFA will recognize the same language, but with a minimized alphabet optimized for faster simulation and smaller memory footprint.
//
// Formally, given a DFA = (Q, Σ, δ, q₀, F), we want compute a partition of Σ into k equivalence classes such that
//
//	P = {C₁, C₂, ..., Cₖ} where each Cᵢ ⊆ Σ and ∪ Cᵢ = Σ
//	∀ a, b ∈ Cᵢ, ∀ q ∈ Q: δ(q, a) = δ(q, b)
//
// Informally, two symbols are equivalent if they lead to the same next state from any given state (every state treats them the same).
func (b *DFABuilder) Build() *DFA {
	// Collect boundaries for all ranges.
	bounds := make([]dfaBound, 0, b.trans.Size()*2*2) // Approximation
	for s, ranges := range b.trans.All() {
		for r, next := range ranges.All() {
			bounds = append(bounds,
				dfaBound{r.Lo, true, dfaTransition{dfaTransitionEnds{s, next}, r}},
				dfaBound{r.Hi + 1, false, dfaTransition{dfaTransitionEnds{s, next}, r}},
			)
		}
	}

	// Sort boundaries in increasing order.
	slices.SortFunc(bounds, func(lhs, rhs dfaBound) int {
		return int(lhs.Value - rhs.Value)
	})

	active := make(map[dfaTransitionEnds]disc.Range[Symbol])
	partition := make([]generic.KeyValue[disc.Range[Symbol], dfaTransitionVector], 0, len(bounds)/2) // Approximation

	for i := 0; i < len(bounds)-1; i++ {
		// Maintain a list of transitions seen between the beginning and end of the current range.
		if b := bounds[i]; b.Start {
			active[b.Transition.Ends] = b.Transition.Range
		} else {
			delete(active, b.Transition.Ends)
		}

		lo := bounds[i].Value
		hi := bounds[i+1].Value - 1

		// This effectively dedups the boundaries.
		if lo <= hi {
			// Collect all transitions seen between the beginning and end of the current range.
			vector := newDFATransitionVector()
			for ends := range active {
				vector.Add(ends)
			}

			// Skip empty classes corresponding to ranges between the end of one boundary and start of the next one.
			if vector.Size() > 0 {
				partition = append(partition, generic.KeyValue[disc.Range[Symbol], dfaTransitionVector]{
					Key: disc.Range[Symbol]{Lo: lo, Hi: hi},
					Val: vector,
				})
			}
		}
	}

	nextCID := classID(0)
	transitionVectors := symboltable.NewRedBlack(cmpDFATransitionVector, eqClassID)
	ranges := newRangeMapping(nil)
	transitions := newDFATransitionTable()

	// Group ranges by their transition vectors to form equivalence classes.
	for _, sub := range partition {
		cid, ok := transitionVectors.Get(sub.Val)
		if !ok {
			cid = nextCID
			nextCID++
			transitionVectors.Put(sub.Val, cid)
		}

		ranges.Add(sub.Key, cid)

		// Build class-based transitions for the current range and its transitions.
		for ends := range sub.Val.All() {
			transitions.Add(ends.State, cid, ends.Next)
		}
	}

	return &DFA{
		start:  b.start,
		final:  b.final,
		ranges: ranges,
		trans:  transitions,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// DFA represents a deterministic finite automaton.
//
// A DFA is defined by a 5-tuple (Q, Σ, δ, q₀, F) where:
//
//   - Q is a finite set of states.
//   - Σ is a finite set of input symbols (alphabet).
//   - δ: Q × Σ → Q is the transition function.
//   - q₀ ∈ Q is the initial (start) state.
//   - F ⊆ Q is the set of accepting (final) states.
//
// This model is meant to be an immutable representation of deterministic finite automata.
// Algorithms that transform or optimize a DFA must construct and return a new DFA.
type DFA struct {
	start  State
	final  States
	ranges rangeMapping
	trans  *dfaTransitionTable

	// Derived values (computed lazily)
	_states  []State
	_symbols []disc.Range[Symbol]
	_classes classMapping
}

// String implements the fmt.Stringer interface.
func (d *DFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", d.start)
	fmt.Fprintf(&b, "Final states: ")

	for s := range d.final.All() {
		fmt.Fprintf(&b, "%d, ", s)
	}

	if b.Len() >= 2 {
		b.Truncate(b.Len() - 2)
	}

	b.WriteString("\nTransitions:\n")

	for s, seq := range d.Transitions() {
		for ranges, next := range seq {
			fmt.Fprintf(&b, "  %d -- %s --> %d\n", s, formatRangeSlice(ranges), next)
		}
	}

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (d *DFA) Clone() *DFA {
	dd := &DFA{
		start:  d.start,
		final:  d.final.Clone(),
		ranges: d.ranges.Clone(),
		trans:  d.trans.Clone(),
	}

	return dd
}

// Equal implements the generic.Equaler interface.
func (d *DFA) Equal(rhs *DFA) bool {
	if rhs == nil {
		return false
	}

	return d.start == rhs.start &&
		d.final.Equal(rhs.final) &&
		d.ranges.Equal(rhs.ranges) &&
		d.trans.Equal(rhs.trans)
}

// Isomorphic determines whether or not two DFAs are isomorphically the same.
//
// Two DFAs D₁ and D₂ are said to be isomorphic if there exists a bijection f: S(D₁) → S(D₂) between their state sets such that,
// for every input symbol a, there is a transition from state s to state t on input a in D₁
// if and only if there is a transition from state f(s) to state f(t) on input a in D₂.
//
// In simpler terms, the two DFAs have the same structure:
// one can be transformed into the other by renaming its states and preserving the transitions.
//
// This is a very expensive operation as graph isomorphism problem is an NP (non-deterministic polynomial time) problem.
func (d *DFA) Isomorphic(rhs *DFA) bool {
	// D₁ and D₂ must have the same number of final states.
	if d.final.Size() != rhs.final.Size() {
		return false
	}

	// D₁ and D₂ must have the same number of states.
	Q1, Q2 := d.States(), rhs.States()
	if len(Q1) != len(Q2) {
		return false
	}

	// D₁ and D₂ must have the same input alphabet.
	Σ1, Σ2 := d.Symbols(), rhs.Symbols()

	if len(Σ1) != len(Σ2) {
		return false
	}

	for i := range Σ1 {
		if Σ1[i] != Σ2[i] {
			return false
		}
	}

	// D₁ and D₂ must have the same sorted degree sequence.
	// len(degs1) == len(degs2) since D₁ and D₂ have the same number of states.
	degs1, degs2 := d.getSortedDegreeSequence(), rhs.getSortedDegreeSequence()
	for i := range degs1 {
		if degs1[i] != degs2[i] {
			return false
		}
	}

	// Since generateStatePermutations uses backtracking and modifies the slice in-place, we need a copy.
	states := make([]State, len(Q1))
	copy(states, Q1)

	// Methodically checking if any permutation of D₁ states is equal to D₂.
	return !generateStatePermutations(states, 0, len(states)-1, func(permutation []State) bool {
		// Create a bijection between the states of D₁ and the current permutation of D₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(Q1))
		for i, s := range Q1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[d.start]

		permutedFinal := make([]State, 0, d.final.Size())
		for f := range d.final.All() {
			permutedFinal = append(permutedFinal, bijection[f])
		}

		b := NewDFABuilder().SetStart(permutedStart).SetFinal(permutedFinal)

		for s, stab := range d.trans.All() {
			for cid, t := range stab.All() {
				ss, tt := bijection[s], bijection[t]

				if ranges, ok := d.classes().Get(cid); ok {
					for r := range ranges.All() {
						b.AddTransition(ss, r.Lo, r.Hi, tt)
					}
				}
			}
		}

		// If the current permutation of D₁ is equal to D₂, we stop checking more permutations by returning false.
		// If the current permutation of D₁ is not equal to D₂, we continue with checking more permutations by returning true.
		return !b.Build().Equal(rhs)
	})
}

// getSortedDegreeSequence calculates the total degree (sum of in-degrees and out-degrees)
// for each state in the DFA and returns the degree sequence sorted in ascending order.
func (d *DFA) getSortedDegreeSequence() []int {
	totalDegrees := map[State]int{}
	for s, strans := range d.trans.All() {
		for _, t := range strans.All() {
			totalDegrees[s]++
			totalDegrees[t]++
		}
	}

	sortedDegrees := make([]int, len(totalDegrees))
	for i, degree := range totalDegrees {
		sortedDegrees[i] = degree
	}

	sort.Quick3Way[int](sortedDegrees, generic.NewCompareFunc[int]())

	return sortedDegrees
}

// Start returns the start state of the DFA.
func (d *DFA) Start() State {
	return d.start
}

// Final returns the final (accepting) states of the DFA.
func (d *DFA) Final() []State {
	return generic.Collect1(d.final.All())
}

// States returns all states in the DFA.
func (d *DFA) States() []State {
	// Lazy initialization
	if d._states == nil {
		states := NewStates(d.start).Union(d.final)
		for s, stab := range d.trans.All() {
			states.Add(s)
			for _, next := range stab.All() {
				states.Add(next)
			}
		}

		d._states = generic.Collect1(states.All())
	}

	return d._states
}

// Symbols returns all symbol ranges in the DFA.
func (d *DFA) Symbols() []disc.Range[Symbol] {
	// Lazy initialization
	if d._symbols == nil {
		d._symbols = make([]disc.Range[Symbol], 0, d.ranges.Size())
		for r := range d.ranges.All() {
			d._symbols = append(d._symbols, r)
		}
	}

	return d._symbols
}

// classes populates the equivalence classes of the input symbols lazily.
// It builds a classID-to-ranges mapping from the range-to-classID mapping.
func (d *DFA) classes() classMapping {
	// Lazy initialization
	if d._classes == nil {
		d._classes = newClassMapping(nil)

		for r, cid := range d.ranges.All() {
			ranges, ok := d._classes.Get(cid)
			if !ok {
				ranges = newRangeSet()
				d._classes.Put(cid, ranges)
			}

			ranges.Add(r)
		}
	}

	return d._classes
}

// Transitions returns all transitions in the DFA.
func (d *DFA) Transitions() iter.Seq2[State, iter.Seq2[[]disc.Range[Symbol], State]] {
	return func(yield func(State, iter.Seq2[[]disc.Range[Symbol], State]) bool) {
		for s := range d.trans.All() {
			if !yield(s, d.TransitionsFrom(s)) {
				return
			}
		}
	}
}

// TransitionsFrom returns all transitions from the given state in the DFA.
func (d *DFA) TransitionsFrom(s State) iter.Seq2[[]disc.Range[Symbol], State] {
	// Aggregate ranges leading to the same next state.
	agg := symboltable.NewRedBlack[State, rangeList](CmpState, nil)

	if stab, ok := d.trans.Get(s); ok {
		for cid, next := range stab.All() {
			// Ensure there is a range list for the current next state.
			list, ok := agg.Get(next)
			if !ok {
				list = newRangeList()
				agg.Put(next, list)
			}

			// Convert from class ID to ranges and combine all ranges.
			if ranges, ok := d.classes().Get(cid); ok {
				for r := range ranges.All() {
					list.Add(r)
				}
			}
		}
	}

	// Yield all aggregated ranges and their corresponding next states.
	return func(yield func([]disc.Range[Symbol], State) bool) {
		for next, list := range agg.All() {
			ranges := generic.Collect1(list.All())

			if !yield(ranges, next) {
				return
			}
		}
	}
}

// Minimize creates a unique DFA with the minimum number of states.
//
// The minimization algorithm sometimes produces a DFA with one dead state.
// This state is not accepting and transfers to itself on each input symbol.
//
// We often want to know when there is no longer any possibility of acceptance.
// If so, we may want to eliminate the dead state and use an automaton that is missing some transitions.
// This automaton has one fewer state than the minimum-state DFA.
// Strictly speaking, such an automaton is not a DFA, because of the missing transitions to the dead state.
//
// For more information and details, see "Compilers: Principles, Techniques, and Tools (2nd Edition)".
func (d *DFA) Minimize() *DFA {
	/*
	 * 1. Start with an initial partition P with two groups,
	 *    F and S - F, the accepting and non-accepting states.
	 */

	// Gather the set of all states
	states := NewStates(d.States()...)

	F := d.final.Clone()       // F
	NF := states.Difference(F) // S - F

	Π := newPartition()
	Π.Add(NF, F)

	/*
	 * 2. Initially, let Πnew = Π.
	 *    For (each group G of Π) {
	 *      Partition G into subgroups such that two states s and t are in the same subgroup
	 *      if and only if for all input symbols a, states s and t have transitions on a to states in the same group of Π
	 *      (at worst, a state will be in a subgroup by itself).
	 *
	 *      Replace G in Pnew by the set of all subgroups formed.
	 *    }
	 *
	 * 3. If Πnew = Π, let Πfinal = Π and continue with step (4).
	 *    Otherwise, repeat step (2) with Πnew in place of Π.
	 */

	for {
		Πnew := newPartition()

		// For every group in the current partition
		for G := range Π.groups.All() {
			Gtrans := buildGroupTransitions(d, Π, G)
			partitionGroup(Πnew, Gtrans)
		}

		if Πnew.Equal(Π) {
			break
		}

		Π = Πnew
	}

	/*
	 * 4. Choose one state in each group of Πfinal as the representative for that group.
	 *    The representatives will be the states of the minimum-state DFA D′.
	 *    The other components of D′ are constructed as follows:
	 *
	 *    (a) The start state of D′ is the representative of the group containing the start state of D.
	 *    (b) The accepting states of D′ are the representatives of those groups that contain an accepting state of D
	 *        (each group contains either only accepting states, or only non-accepting states).
	 *    (c) Let s be the representative of some group G of Πfinal, and let the transition of D from s on input a be to state t.
	 *        Let r be the representative of t's group H. Then in D′, there is a transition from s to r on input a.
	 */

	start := Π.FindRep(d.start)

	final := NewStates()
	for f := range d.final.All() {
		g := Π.FindRep(f)
		final.Add(g)
	}

	b := NewDFABuilder().SetStart(start)
	b.final = final

	for G := range Π.groups.All() {
		// Get any state in the group
		s, _ := G.States.FirstMatch(func(State) bool {
			return true
		})

		if stab, ok := d.trans.Get(s); ok {
			for cid, next := range stab.All() {
				rep := Π.FindRep(next)

				if ranges, ok := d.classes().Get(cid); ok {
					for r := range ranges.All() {
						b.AddTransition(G.Rep, r.Lo, r.Hi, rep)
					}
				}
			}
		}
	}

	return b.Build()
}

// buildGroupTransitions constructs a transition table for the states in group G using the current partition and the DFA.
//
// For every DFA transition s --classID--> next where s is a member of G, the table records
// a mapping from the pair (s, classID) to the representative state of the partition group that contains next.
//
// The resulting table maps (state, classID) -> representative state and
// captures how each state in G behaves with respect to the current partition.
//
// The partitioning algorithm uses this table to further split G into smaller subgroups.
func buildGroupTransitions(d *DFA, P *partition, G group) *dfaTransitionTable {
	Gtrans := newDFATransitionTable()

	for s := range G.All() {
		if stab, ok := d.trans.Get(s); ok {
			for cid, next := range stab.All() {
				if rep := P.FindRep(next); rep != -1 {
					Gtrans.Add(s, cid, rep)
				}
			}
		}
	}

	return Gtrans
}

// partitionGroup splits the states described by Gtrans into subgroups and adds those subgroups to P.
//
// Gtrans maps each state to its transition profile:
// for every classID, it records the group representative of the next state.
//
// This method partition the group G into subgroups such that two states s and t are in the same subgroup
// if and only if, for all input symbols a (or the equivalance classes of the input symbols),
// the transitions of s and t on all inputs lead to states in the same group.
//
// If no such grouping is possible, a state will be placed in a subgroup by itself.
func partitionGroup(P *partition, Gtrans *dfaTransitionTable) {
	all := generic.Collect2(Gtrans.All())

	for i := 0; i < len(all); i++ {
		s, stab := all[i].Key, all[i].Val

		// If s is not already added to the new partition
		if P.FindRep(s) == -1 {
			// Create a new group in the new partition
			H := NewStates(s)

			// Add all other states with same classID --> rep mapping to the new group
			for j := i + 1; j < len(all); j++ {
				t, ttab := all[j].Key, all[j].Val

				if stab.Equal(ttab) && !H.Contains(t) {
					H.Add(t)
				}
			}

			P.Add(H)
		}
	}
}

// EliminateDeadStates eliminates the dead states and all transitions to them.
// The subset construction and minimization algorithms sometimes produce a DFA with a single dead state.
//
// Strictly speaking, a DFA must have a transition from every state on every input symbol in its input alphabet.
// When we construct a DFA to be used in a lexical analyzer, we need to treat the dead state differently.
// We must know when there is no longer any possibility of recognizing a longer lexeme.
// Thus, we should always omit transitions to the dead state and eliminate the dead state itself.
func (d *DFA) EliminateDeadStates() *DFA {
	// 1. Construct a directed graph from the DFA with all the transitions reversed.
	adj := map[State]States{}
	for s, strans := range d.trans.All() {
		for _, t := range strans.All() {
			if adj[t] == nil {
				adj[t] = NewStates()
			}
			adj[t].Add(s)
		}
	}

	// 2. Add a new state that transitions to all final states of the DFA.
	u := State(-1)
	adj[u] = d.final.Clone()

	// 3. Finally, we find all states reachable from this new state using a depth-first search (DFS).
	//    All other states not connected to this new state will be identified as dead states.
	visited := map[State]bool{}
	for s := range adj {
		visited[s] = false
	}

	markReachable(adj, visited, u)

	deads := NewStates()
	for s, visited := range visited {
		if !visited {
			deads.Add(s)
		}
	}

	b := NewDFABuilder().SetStart(d.start)
	b.final = d.final.Clone()

	for s, stab := range d.trans.All() {
		for cid, t := range stab.All() {
			if !deads.Contains(s) && !deads.Contains(t) {
				if ranges, ok := d.classes().Get(cid); ok {
					for r := range ranges.All() {
						b.AddTransition(s, r.Lo, r.Hi, t)
					}
				}
			}
		}
	}

	return b.Build()
}

// markReachable is a depth-first search on the DFA graph.
func markReachable(adj map[State]States, visited map[State]bool, s State) {
	visited[s] = true

	if adj[s] != nil {
		for t := range adj[s].All() {
			if !visited[t] {
				markReachable(adj, visited, t)
			}
		}
	}
}

// ReindexStates reassigns indices to states based on a
// breadth-first traversal of the DFA, starting from the initial state.
// This method is typically called after removing unreachable or dead states from the DFA.
func (d *DFA) ReindexStates() *DFA {
	sm := newStateManager(-1)

	visited := map[State]bool{}
	queue := list.NewQueue[State](64, nil)

	visited[d.start] = true
	queue.Enqueue(d.start)
	sm.GetOrCreateState(0, d.start)

	for !queue.IsEmpty() {
		s, _ := queue.Dequeue()
		if adj, ok := d.trans.Get(s); ok {
			for _, t := range adj.All() {
				if !visited[t] {
					visited[t] = true
					queue.Enqueue(t)
					sm.GetOrCreateState(0, t)
				}
			}
		}
	}

	start := sm.GetOrCreateState(0, d.start)

	b := NewDFABuilder().SetStart(start)

	final := NewStates()
	for f := range d.final.All() {
		ff := sm.GetOrCreateState(0, f)
		final.Add(ff)
	}

	b.final = final

	for s, stab := range d.trans.All() {
		ss := sm.GetOrCreateState(0, s)

		for cid, t := range stab.All() {
			tt := sm.GetOrCreateState(0, t)

			if ranges, ok := d.classes().Get(cid); ok {
				for r := range ranges.All() {
					b.AddTransition(ss, r.Lo, r.Hi, tt)
				}
			}
		}
	}

	return b.Build()
}

// Union constructs a DFA that recognizes the union of the languages accepted by the provided DFAs.
//
// The process involves:
//
//  1. Converting each DFA to an NFA.
//  2. Building a single NFA that accepts the union of all input NFAs.
//  3. Converting the unified NFA to a DFA.
//  4. Removing dead states and transitions to them.
//  5. Reindexing states to maintain a compact representation.
//
// The returned DFA accepts any string accepted by at least one input DFA.
// The second return value maps each input DFA to its corresponding final states in the resulting DFA.
//
// Note: The resulting DFA is not minimized, so final states remain distinguishable for each input DFA.
// This is useful for applications like lexical analysis, where tracking the origin of acceptance matters.
func (d *DFA) Union(ds ...*DFA) (*DFA, [][]State) {
	all := append([]*DFA{d}, ds...)
	return UnionDFA(all...)
}

// UnionDFA constructs a DFA that recognizes the union of the languages accepted by the provided DFAs.
//
// The process involves:
//
//  1. Converting each DFA to an NFA.
//  2. Building a single NFA that accepts the union of all input NFAs.
//  3. Converting the unified NFA to a DFA.
//  4. Removing dead states and transitions to them.
//  5. Reindexing states to maintain a compact representation.
//
// The returned DFA accepts any string accepted by at least one input DFA.
// The second return value maps each input DFA to its corresponding final states in the resulting DFA.
//
// Note: The resulting DFA is not minimized, so final states remain distinguishable for each input DFA.
// This is useful for applications like lexical analysis, where tracking the origin of acceptance matters.
func UnionDFA(ds ...*DFA) (*DFA, [][]State) {
	var finalMap [][]State

	// 1. Convert all DFAs to NFAs.
	ns := make([]*NFA, len(ds))
	for i, d := range ds {
		ns[i] = d.ToNFA()
	}

	// 2. Construct a new NFA that accepts the union of the languages accepted by each NFA.
	var N *NFA

	{
		start, final := State(0), State(1)
		finalMap = make([][]State, len(ns))

		b := NewNFABuilder().SetStart(start).SetFinal([]State{final})
		sm := newStateManager(final)

		for id, n := range ns {
			for s, stab := range n.trans.All() {
				ss := sm.GetOrCreateState(id, s)

				for cid, states := range stab.All() {
					next := make([]State, 0, states.Size())
					for t := range states.All() {
						tt := sm.GetOrCreateState(id, t)
						next = append(next, tt)
					}

					if ranges, ok := n.classes().Get(cid); ok {
						for r := range ranges.All() {
							b.AddTransition(ss, r.Lo, r.Hi, next)
						}
					}
				}
			}

			ss := sm.GetOrCreateState(id, n.start)
			b.AddTransition(start, E, E, []State{ss})

			for f := range n.final.All() {
				ff := sm.GetOrCreateState(id, f)
				b.AddTransition(ff, E, E, []State{final})
				finalMap[id] = append(finalMap[id], ff)
			}
		}

		N = b.Build()
	}

	// 3. Convert the NFA into a DFA.
	var D *DFA

	{
		// Look up the class ID for ε
		_, eid, hasε := N.ranges.Find(E)

		b := NewDFABuilder().SetStart(0)

		// Initially, ε-closure(s₀) is the only state in Dstates
		S0 := NewStates(N.start)
		Dstates := list.NewSoftQueue(EqStates)
		Dstates.Enqueue(N.εClosure(S0))

		for T, i := Dstates.Dequeue(); i >= 0; T, i = Dstates.Dequeue() {
			// For each input symbol c (or equivalency for each equivalence class of the input symbols)
			for cid, ranges := range N.classes().All() {
				if !hasε || cid != eid {
					U := N.εClosure(N.move(T, cid))

					// If U is not in Dstates, add U to Dstates
					j := Dstates.Contains(U)
					if j == -1 {
						j = Dstates.Enqueue(U)
					}

					for r := range ranges.All() {
						b.AddTransition(State(i), r.Lo, r.Hi, State(j))
					}
				}
			}
		}

		final := NewStates()

		for i, S := range Dstates.Values() {
			for f := range N.final.All() {
				if S.Contains(f) {
					final.Add(State(i))
					break // The accepting states of D are all those sets of N's states that include at least one accepting state of N
				}
			}
		}

		b.final = final
		D = b.Build()

		// Remap the final states from the union NFA to combined DFA.
		for id, states := range finalMap {
			mapped := NewStates()
			for _, f := range states {
				for i, S := range Dstates.Values() {
					if S.Contains(f) {
						mapped.Add(State(i))
					}
				}
			}
			finalMap[id] = generic.Collect1(mapped.All())
		}
	}

	// 4. Remove dead states and their transitions.
	D = D.EliminateDeadStates()

	// 5. Reassign state indices.
	{
		sm := newStateManager(-1)

		visited := map[State]bool{}
		queue := list.NewQueue[State](64, nil)

		visited[D.start] = true
		queue.Enqueue(D.start)
		sm.GetOrCreateState(0, D.start)

		for !queue.IsEmpty() {
			s, _ := queue.Dequeue()
			if stab, ok := D.trans.Get(s); ok {
				for _, t := range stab.All() {
					if !visited[t] {
						visited[t] = true
						queue.Enqueue(t)
						sm.GetOrCreateState(0, t)
					}
				}
			}
		}

		start := sm.GetOrCreateState(0, D.start)
		b := NewDFABuilder().SetStart(start)

		final := NewStates()
		for f := range D.final.All() {
			ff := sm.GetOrCreateState(0, f)
			final.Add(ff)
		}

		b.final = final

		for s, stab := range D.trans.All() {
			ss := sm.GetOrCreateState(0, s)

			for cid, t := range stab.All() {
				tt := sm.GetOrCreateState(0, t)

				if ranges, ok := D.classes().Get(cid); ok {
					for r := range ranges.All() {
						b.AddTransition(ss, r.Lo, r.Hi, tt)
					}
				}
			}
		}

		D = b.Build()

		// Remap the final states from the old indices to new indices.
		for id, states := range finalMap {
			mapped := NewStates()
			for _, f := range states {
				ff := sm.GetOrCreateState(0, f)
				mapped.Add(ff)
			}
			finalMap[id] = generic.Collect1(mapped.All())
		}
	}

	return D, finalMap
}

// ToNFA constructs a new NFA accepting the same language as the DFA (every DFA is an NFA).
func (d *DFA) ToNFA() *NFA {
	b := NewNFABuilder().SetStart(d.start)
	b.final = d.final.Clone()

	for s, seq := range d.Transitions() {
		for ranges, next := range seq {
			for _, r := range ranges {
				b.AddTransition(s, r.Lo, r.Hi, []State{next})
			}
		}
	}

	return b.Build()
}

// DOT generates a DOT representation of the DFA transition graph for visualization.
func (d *DFA) DOT() string {
	graph := dot.NewGraph(false, true, false, "DFA", dot.RankDirLR, "", "", dot.ShapeCircle)

	for _, s := range d.States() {
		name := fmt.Sprintf("%d", s)
		label := fmt.Sprintf("%d", s)

		if s == d.start {
			graph.AddNode(dot.NewNode("start", "", "", "", dot.StyleInvis, "", "", ""))
			graph.AddEdge(dot.NewEdge("start", name, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		var shape dot.Shape
		if d.final.Contains(s) {
			shape = dot.ShapeDoubleCircle
		}

		graph.AddNode(dot.NewNode(name, "", label, "", "", shape, "", ""))
	}

	for s, seq := range d.Transitions() {
		for ranges, next := range seq {
			from := fmt.Sprintf("%d", s)
			to := fmt.Sprintf("%d", next)
			label := formatRangeSlice(ranges)

			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DOT() + "\n"
}

// Runner constructs a new DFARunner for simulating (running) the DFA on input symbols.
func (d *DFA) Runner() *DFARunner {
	trans := symboltable.NewQuadraticHashTable(HashState, EqState, eqClassIDStateTable, symboltable.HashOpts{})

	for s, stab := range d.trans.All() {
		stabClone := symboltable.NewQuadraticHashTable(hashClassID, eqClassID, EqState, symboltable.HashOpts{})
		for cid, next := range stab.All() {
			stabClone.Put(cid, next)
		}

		trans.Put(s, stabClone)
	}

	return &DFARunner{
		start:  d.start,
		final:  d.final.Clone(),
		ranges: d.ranges.Clone(),
		trans:  trans,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// DFARunner is used for simulating (running) a DFA on input symbols.
// It is immutable and optimized for fast execution.
type DFARunner struct {
	start  State
	final  States
	ranges rangeMapping
	trans  symboltable.SymbolTable[State, symboltable.SymbolTable[classID, State]]
}

// Next returns the next state from state s on input symbol a.
func (r *DFARunner) Next(s State, a Symbol) State {
	if stab, ok := r.trans.Get(s); ok {
		if _, cid, ok := r.ranges.Find(a); ok {
			if next, ok := stab.Get(cid); ok {
				return next
			}
		}
	}

	return -1
}

// Accept determines whether an input string is recognized (accepted) by the DFA.
func (r *DFARunner) Accept(s String) bool {
	var curr State
	for curr = r.start; len(s) > 0; s = s[1:] {
		curr = r.Next(curr, s[0])
	}

	return r.final.Contains(curr)
}

/* ------------------------------------------------------------------------------------------------------------------------ */
