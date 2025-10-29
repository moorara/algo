package automata

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

var εRange = disc.Range[Symbol]{Lo: E, Hi: E}

/* ------------------------------------------------------------------------------------------------------------------------ */

// nfaTransitionEnds represents the states involved in a NFA transition.
type nfaTransitionEnds struct {
	State State
	Next  States
}

var cmpNFATransitionEnds = func(lhs, rhs nfaTransitionEnds) int {
	if c := CmpState(lhs.State, rhs.State); c != 0 {
		return c
	}

	// Next States are sorted set and can be compared element-wise.

	s1 := generic.Collect1(lhs.Next.All())
	s2 := generic.Collect1(rhs.Next.All())

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if c := CmpState(s1[i], s2[i]); c != 0 {
			return c
		}
	}

	return len(s1) - len(s2)
}

// nfaTransition  represents a transition from one state to a set of states on a range of input symbols.
type nfaTransition struct {
	Ends  nfaTransitionEnds
	Range disc.Range[Symbol]
}

// nfaBound represents a boundary (start or end) of a range associated with a NFA transition.
type nfaBound struct {
	Value      Symbol
	Start      bool
	Transition nfaTransition
}

// nfaTransitionVector represents a NFA transition vector.
type nfaTransitionVector set.Set[nfaTransitionEnds]

func newNFATransitionVector() nfaTransitionVector {
	return set.NewSorted(cmpNFATransitionEnds)
}

var cmpNFATransitionVector = func(lhs, rhs nfaTransitionVector) int {
	v1 := generic.Collect1(lhs.All())
	v2 := generic.Collect1(rhs.All())

	for i := 0; i < len(v1) && i < len(v2); i++ {
		if c := cmpNFATransitionEnds(v1[i], v2[i]); c != 0 {
			return c
		}
	}

	return len(v1) - len(v2)
}

/* ------------------------------------------------------------------------------------------------------------------------ */

var eqClassIDStatesTable = func(a, b symboltable.SymbolTable[classID, States]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Equal(b)
}

// nfaTransitionTable represents the transition table of a NFA.
type nfaTransitionTable struct {
	// Use an ordered symbol table so iterations over states and classes are deterministic and
	// the resulting computation and textual output are reproducible.
	symboltable.SymbolTable[State, symboltable.SymbolTable[classID, States]]
}

// newNFATransitionTable creates a new instance of nfaTransitionTable.
func newNFATransitionTable() *nfaTransitionTable {
	return &nfaTransitionTable{
		symboltable.NewRedBlack(CmpState, eqClassIDStatesTable),
	}
}

// Clone implements the generic.Cloner interface.
func (t *nfaTransitionTable) Clone() *nfaTransitionTable {
	clone := newNFATransitionTable()

	for s, stab := range t.All() {
		stabClone := symboltable.NewRedBlack(cmpClassID, EqStates)
		for cid, next := range stab.All() {
			stabClone.Put(cid, next)
		}
		clone.Put(s, stabClone)
	}

	return clone
}

// Equal implements the generic.Equaler interface.
func (t *nfaTransitionTable) Equal(rhs *nfaTransitionTable) bool {
	return t.SymbolTable.Equal(rhs.SymbolTable)
}

// Add inserts a new transition into the NFA transition table.
func (t *nfaTransitionTable) Add(s State, cid classID, next States) *nfaTransitionTable {
	stab, ok := t.Get(s)
	if !ok {
		stab = symboltable.NewRedBlack(cmpClassID, EqStates)
		t.Put(s, stab)
	}

	stab.Put(cid, next)

	return t
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// NFABuilder implements the Builder design pattern for constructing NFA instances.
//
// The Builder pattern separates the construction of an NFA from its representation,
// This approach ensures the resulting NFA is immutable and optimized for simulation (running).
type NFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, disc.RangeMap[Symbol, States]]
}

// NewNFABuilder creates a new NFA builder instance.
func NewNFABuilder() *NFABuilder {
	return &NFABuilder{
		trans: symboltable.NewRedBlack[State, disc.RangeMap[Symbol, States]](CmpState, nil),
	}
}

// SetStart sets the start state of the NFA.
func (b *NFABuilder) SetStart(s State) *NFABuilder {
	b.start = s
	return b
}

// SetFinal sets the final (accepting) states of the NFA.
func (b *NFABuilder) SetFinal(f []State) *NFABuilder {
	b.final = NewStates(f...)
	return b
}

// AddTransition adds transitions from state s to states next on all input symbols in the range [start, end].
func (b *NFABuilder) AddTransition(s State, start, end Symbol, next []State) *NFABuilder {
	ranges, ok := b.trans.Get(s)
	if !ok {
		opts := &disc.RangeMapOpts[Symbol, States]{Resolve: unionStates}
		ranges = disc.NewRangeMap(EqStates, opts)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		NewStates(next...),
	)

	return b
}

// Build constructs the NFA based on the configurations provided to the builder.
//
// This method reduces the size of the alphabet by computing the equivalence classes of input symbols based on the transition function.
// The resulting NFA will recognize the same language, but with a minimized alphabet optimized for faster simulation and smaller memory footprint.
//
// Formally, given a NFA = (Q, Σ, δ, q₀, F), we want compute a partition of Σ into k equivalence classes such that
//
//	P = {C₁, C₂, ..., Cₖ} where each Cᵢ ⊆ Σ and ∪ Cᵢ = Σ
//	∀ a, b ∈ Cᵢ, ∀ q ∈ Q: δ(q, a) = δ(q, b)
//
// Informally, two symbols are equivalent if they lead to the same next state from any given state (every state treats them the same).
func (b *NFABuilder) Build() *NFA {
	// Collect boundaries for all ranges.
	bounds := make([]nfaBound, 0, b.trans.Size()*2*2) // Approximation
	for s, ranges := range b.trans.All() {
		for r, next := range ranges.All() {
			bounds = append(bounds,
				nfaBound{r.Lo, true, nfaTransition{nfaTransitionEnds{s, next}, r}},
				nfaBound{r.Hi + 1, false, nfaTransition{nfaTransitionEnds{s, next}, r}},
			)
		}
	}

	// Sort boundaries in increasing order.
	slices.SortFunc(bounds, func(lhs, rhs nfaBound) int {
		return int(lhs.Value - rhs.Value)
	})

	active := make(map[nfaTransitionEnds]disc.Range[Symbol])
	partition := make([]generic.KeyValue[disc.Range[Symbol], nfaTransitionVector], 0, len(bounds)/2) // Approximation

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
			vector := newNFATransitionVector()
			for ends := range active {
				vector.Add(ends)
			}

			// Skip empty classes corresponding to ranges between the end of one boundary and start of the next one.
			if vector.Size() > 0 {
				partition = append(partition, generic.KeyValue[disc.Range[Symbol], nfaTransitionVector]{
					Key: disc.Range[Symbol]{Lo: lo, Hi: hi},
					Val: vector,
				})
			}
		}
	}

	nextCID := classID(0)
	transitionVectors := symboltable.NewRedBlack(cmpNFATransitionVector, eqClassID)
	ranges := newRangeMapping(nil)
	transitions := newNFATransitionTable()

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

	return &NFA{
		start:  b.start,
		final:  b.final,
		ranges: ranges,
		trans:  transitions,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// NFA represents a non-deterministic finite automaton.
//
// A NFA is defined by a 5-tuple (Q, Σ, δ, q₀, F) where:
//
//   - Q is a finite set of states.
//   - Σ is a finite set of input symbols (alphabet).
//   - δ: Q × Σ → P(Q) is the transition function.
//   - q₀ ∈ Q is the initial (start) state.
//   - F ⊆ Q is the set of accepting (final) states.
//
// This model is meant to be an immutable representation of non-deterministic finite automata.
// Algorithms that transform or optimize a NFA must construct and return a new NFA.
type NFA struct {
	start  State
	final  States
	ranges rangeMapping
	trans  *nfaTransitionTable

	// Derived values (computed lazily)
	_states  []State
	_symbols []disc.Range[Symbol]
	_classes classMapping
}

// εClosure returns the set of NFA states reachable from some NFA state s in set T on ε-transitions alone.
// εClosure(T) = Union(εClosure(s)) for all s ∈ T.
func (n *NFA) εClosure(T States) States {
	closure := T.Clone()

	// Look up the class ID for ε
	_, eid, hasε := n.ranges.Find(E)
	if !hasε {
		return closure
	}

	stack := list.NewStack[State](64, nil) // Approximation
	for s := range T.All() {
		stack.Push(s)
	}

	for !stack.IsEmpty() {
		t, _ := stack.Pop()

		if next := n.next(t, eid); next != nil {
			for u := range next.All() {
				if !closure.Contains(u) {
					closure.Add(u)
					stack.Push(u)
				}
			}
		}
	}

	return closure
}

// move returns the set of NFA states to which there is a transition on the given input from some state s in T.
func (n *NFA) move(T States, cid classID) States {
	states := NewStates()

	for s := range T.All() {
		if next := n.next(s, cid); next != nil {
			states = states.Union(next)
		}
	}

	return states
}

// next returns the set of next states from state s on the given input.
func (n *NFA) next(s State, cid classID) States {
	if stab, ok := n.trans.Get(s); ok {
		if next, ok := stab.Get(cid); ok {
			return next
		}
	}

	return nil
}

// String implements the fmt.Stringer interface.
func (n *NFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", n.start)
	fmt.Fprintf(&b, "Final states: ")

	for s := range n.final.All() {
		fmt.Fprintf(&b, "%d, ", s)
	}

	if b.Len() >= 2 {
		b.Truncate(b.Len() - 2)
	}

	trans := make([]string, 0, n.trans.Size()*2) // Approximation
	for s, stab := range n.trans.All() {
		for cid, next := range stab.All() {
			if ranges, ok := n.classes().Get(cid); ok {
				trans = append(trans, fmt.Sprintf("  %d -- %s --> %s", s, ranges, next))
			}
		}
	}

	fmt.Fprintf(&b, "\nTransitions:\n%s\n", strings.Join(trans, "\n"))

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (n *NFA) Clone() *NFA {
	nn := &NFA{
		start:  n.start,
		final:  n.final.Clone(),
		ranges: n.ranges.Clone(),
		trans:  n.trans.Clone(),
	}

	return nn
}

// Equal implements the generic.Equaler interface.
func (n *NFA) Equal(rhs *NFA) bool {
	if rhs == nil {
		return false
	}

	return n.start == rhs.start &&
		n.final.Equal(rhs.final) &&
		n.ranges.Equal(rhs.ranges) &&
		n.trans.Equal(rhs.trans)
}

// Isomorphic determines whether or not two NFAs are isomorphically the same.
//
// Two NFAs N₁ and N₂ are said to be isomorphic if there exists a bijection f: S(N₁) → S(N₂) between their state sets such that,
// for every input symbol a, there is a transition from state s to state t on input a in N₁
// if and only if there is a transition from state f(s) to state f(t) on input a in N₂.
//
// In simpler terms, the two NFAs have the same structure:
// one can be transformed into the other by renaming its states and preserving the transitions.
//
// This is a very expensive operation as graph isomorphism problem is an NP (non-deterministic polynomial time) problem.
func (n *NFA) Isomorphic(rhs *NFA) bool {
	// N₁ and N₂ must have the same number of final states.
	if n.final.Size() != rhs.final.Size() {
		return false
	}

	// N₁ and N₂ must have the same number of states.
	Q1, Q2 := n.States(), rhs.States()
	if len(Q1) != len(Q2) {
		return false
	}

	// N₁ and N₂ must have the same input alphabet.
	Σ1, Σ2 := n.Symbols(), rhs.Symbols()

	if len(Σ1) != len(Σ2) {
		return false
	}

	for i := range Σ1 {
		if Σ1[i] != Σ2[i] {
			return false
		}
	}

	// N₁ and N₂ must have the same sorted degree sequence.
	// len(degs1) == len(degs2) since N₁ and N₂ have the same number of states.
	degs1, degs2 := n.getSortedDegreeSequence(), rhs.getSortedDegreeSequence()
	for i := range degs1 {
		if degs1[i] != degs2[i] {
			return false
		}
	}

	// Since generateStatePermutations uses backtracking and modifies the slice in-place, we need a copy.
	states := make([]State, len(Q1))
	copy(states, Q1)

	// Methodically checking if any permutation of N₁ states is equal to N₂.
	return !generateStatePermutations(states, 0, len(states)-1, func(permutation []State) bool {
		// Create a bijection between the states of N₁ and the current permutation of N₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(Q1))
		for i, s := range Q1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[n.start]

		permutedFinal := make([]State, 0, n.final.Size())
		for f := range n.final.All() {
			permutedFinal = append(permutedFinal, bijection[f])
		}

		b := NewNFABuilder().SetStart(permutedStart).SetFinal(permutedFinal)

		for s, stab := range n.trans.All() {
			for cid, ts := range stab.All() {
				ss := bijection[s]

				tts := make([]State, 0, ts.Size())
				for t := range ts.All() {
					tts = append(tts, bijection[t])
				}

				if ranges, ok := n.classes().Get(cid); ok {
					for r := range ranges.All() {
						b.AddTransition(ss, r.Lo, r.Hi, tts)
					}
				}
			}
		}

		// If the current permutation of N₁ is equal to N₂, we stop checking more permutations by returning false.
		// If the current permutation of N₁ is not equal to N₂, we continue with checking more permutations by returning true.
		return !b.Build().Equal(rhs)
	})
}

// getSortedDegreeSequence calculates the total degree (sum of in-degrees and out-degrees)
// for each state in the NFA and returns the degree sequence sorted in ascending order.
func (n *NFA) getSortedDegreeSequence() []int {
	totalDegrees := map[State]int{}
	for s, strans := range n.trans.All() {
		for _, states := range strans.All() {
			for t := range states.All() {
				totalDegrees[s]++
				totalDegrees[t]++
			}
		}
	}

	sortedDegrees := make([]int, len(totalDegrees))
	for i, degree := range totalDegrees {
		sortedDegrees[i] = degree
	}

	sort.Quick3Way[int](sortedDegrees, generic.NewCompareFunc[int]())

	return sortedDegrees
}

// Start returns the start state of the NFA.
func (n *NFA) Start() State {
	return n.start
}

// Final returns the final (accepting) states of the NFA.
func (n *NFA) Final() []State {
	return generic.Collect1(n.final.All())
}

// States returns all states in the NFA.
func (n *NFA) States() []State {
	// Lazy initialization
	if n._states == nil {
		states := NewStates(n.start).Union(n.final)
		for s, stab := range n.trans.All() {
			states.Add(s)
			for _, next := range stab.All() {
				states = states.Union(next)
			}
		}

		n._states = generic.Collect1(states.All())
	}

	return n._states
}

// Symbols returns all symbol ranges in the NFA.
func (n *NFA) Symbols() []disc.Range[Symbol] {
	// Lazy initialization
	if n._symbols == nil {
		n._symbols = make([]disc.Range[Symbol], 0, n.ranges.Size())
		for r := range n.ranges.All() {
			if r != εRange {
				n._symbols = append(n._symbols, r)
			}
		}
	}

	return n._symbols
}

// classes populates the equivalence classes of the input symbols lazily.
// It builds a classID-to-ranges mapping from the range-to-classID mapping.
func (n *NFA) classes() classMapping {
	// Lazy initialization
	if n._classes == nil {
		n._classes = newClassMapping(nil)

		for r, cid := range n.ranges.All() {
			ranges, ok := n._classes.Get(cid)
			if !ok {
				ranges = newRangeSet()
				n._classes.Put(cid, ranges)
			}

			ranges.Add(r)
		}
	}

	return n._classes
}

// Transitions returns all transitions in the NFA.
func (n *NFA) Transitions() iter.Seq2[State, iter.Seq2[[]disc.Range[Symbol], []State]] {
	return func(yield func(State, iter.Seq2[[]disc.Range[Symbol], []State]) bool) {
		for s := range n.trans.All() {
			if !yield(s, n.TransitionsFrom(s)) {
				return
			}
		}
	}
}

// TransitionsFrom returns all transitions from the given state in the NFA.
func (n *NFA) TransitionsFrom(s State) iter.Seq2[[]disc.Range[Symbol], []State] {
	return func(yield func([]disc.Range[Symbol], []State) bool) {
		if stab, ok := n.trans.Get(s); ok {
			for cid, next := range stab.All() {
				if ranges, ok := n.classes().Get(cid); ok {
					k := generic.Collect1(ranges.All())
					v := generic.Collect1(next.All())

					if !yield(k, v) {
						return
					}
				}
			}
		}
	}
}

// Star constructs a new NFA that accepts the Kleene star closure of the language accepted by the NFA.
func (n *NFA) Star() *NFA {
	start, final := State(0), State(1)
	sm := newStateManager(final)

	b := NewNFABuilder().SetStart(start).SetFinal([]State{final})

	for s, seq := range n.Transitions() {
		ss := sm.GetOrCreateState(0, s)

		for ranges, states := range seq {
			next := make([]State, 0, len(states))
			for _, t := range states {
				tt := sm.GetOrCreateState(0, t)
				next = append(next, tt)
			}

			for _, r := range ranges {
				b.AddTransition(ss, r.Lo, r.Hi, next)
			}
		}
	}

	ss := sm.GetOrCreateState(0, n.start)
	b.AddTransition(start, E, E, []State{ss})
	b.AddTransition(start, E, E, []State{final})

	for f := range n.final.All() {
		ff := sm.GetOrCreateState(0, f)
		b.AddTransition(ff, E, E, []State{ss})
		b.AddTransition(ff, E, E, []State{final})
	}

	return b.Build()
}

// Union constructs a new NFA that accepts the union of languages accepted by each individual NFA.
func (n *NFA) Union(ns ...*NFA) *NFA {
	all := append([]*NFA{n}, ns...)
	return UnionNFA(all...)
}

// UnionNFA constructs a new NFA that accepts the union of languages accepted by each individual NFA.
func UnionNFA(ns ...*NFA) *NFA {
	start, final := State(0), State(1)
	sm := newStateManager(final)

	b := NewNFABuilder().SetStart(start).SetFinal([]State{final})

	for id, nfa := range ns {
		for s, seq := range nfa.Transitions() {
			ss := sm.GetOrCreateState(id, s)

			for ranges, states := range seq {
				next := make([]State, 0, len(states))
				for _, t := range states {
					tt := sm.GetOrCreateState(id, t)
					next = append(next, tt)
				}

				for _, r := range ranges {
					b.AddTransition(ss, r.Lo, r.Hi, next)
				}
			}
		}

		ss := sm.GetOrCreateState(id, nfa.start)
		b.AddTransition(start, E, E, []State{ss})

		for f := range nfa.final.All() {
			ff := sm.GetOrCreateState(id, f)
			b.AddTransition(ff, E, E, []State{final})
		}
	}

	return b.Build()
}

// Concat constructs a new NFA that accepts the concatenation of languages accepted by each individual NFA.
func (n *NFA) Concat(ns ...*NFA) *NFA {
	all := append([]*NFA{n}, ns...)
	return ConcatNFA(all...)
}

// ConcatNFA constructs a new NFA that accepts the concatenation of languages accepted by each individual NFA.
func ConcatNFA(ns ...*NFA) *NFA {
	start, final := State(0), []State{0}
	sm := newStateManager(0)

	b := NewNFABuilder().SetStart(start).SetFinal(final)

	for id, nfa := range ns {
		for s, seq := range nfa.Transitions() {
			// If s is the start state of the current NFA,
			// we need to map it to the previous NFA final states.
			var sp []State
			if s == nfa.start {
				sp = final
			} else {
				ss := sm.GetOrCreateState(id, s)
				sp = []State{ss}
			}

			for ranges, states := range seq {
				// If any of the next state is the start state of the current NFA,
				// we need to map it to the previous NFA final states.
				var nextp []State
				for _, t := range states {
					if t == nfa.start {
						nextp = append(nextp, final...)
					} else {
						tt := sm.GetOrCreateState(id, t)
						nextp = append(nextp, tt)
					}
				}

				// Add new transitions.
				for _, s := range sp {
					for _, r := range ranges {
						b.AddTransition(s, r.Lo, r.Hi, nextp)
					}
				}
			}
		}

		// Update the current final states.
		final = make([]State, 0, nfa.final.Size())
		for f := range nfa.final.All() {
			ff := sm.GetOrCreateState(id, f)
			final = append(final, ff)
		}
	}

	b.SetFinal(final)

	return b.Build()
}

// ToDFA constructs a new DFA accepting the same language as the NFA.
// It implements the subset construction algorithm.
func (n *NFA) ToDFA() *DFA {
	// Look up the class ID for ε
	_, eid, hasε := n.ranges.Find(E)

	b := NewDFABuilder().SetStart(0)

	// Initially, ε-closure(s₀) is the only state in Dstates
	S0 := NewStates(n.start)
	Dstates := list.NewSoftQueue(EqStates)
	Dstates.Enqueue(n.εClosure(S0))

	for T, i := Dstates.Dequeue(); i >= 0; T, i = Dstates.Dequeue() {
		// For each input symbol c (or equivalency for each equivalence class of the input symbols)
		for cid, ranges := range n.classes().All() {
			if !hasε || cid != eid {
				U := n.εClosure(n.move(T, cid))

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
		for f := range n.final.All() {
			if S.Contains(f) {
				final.Add(State(i))
				break // The accepting states of D are all those sets of N's states that include at least one accepting state of N
			}
		}
	}

	b.final = final

	return b.Build()
}

// DOT generates a DOT representation of the NFA transition graph for visualization.
func (n *NFA) DOT() string {
	graph := dot.NewGraph(false, true, false, "NFA", dot.RankDirLR, "", "", dot.ShapeCircle)

	for _, s := range n.States() {
		name := fmt.Sprintf("%d", s)
		label := fmt.Sprintf("%d", s)

		if s == n.start {
			graph.AddNode(dot.NewNode("start", "", "", "", dot.StyleInvis, "", "", ""))
			graph.AddEdge(dot.NewEdge("start", name, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		var shape dot.Shape
		if n.final.Contains(s) {
			shape = dot.ShapeDoubleCircle
		}

		graph.AddNode(dot.NewNode(name, "", label, "", "", shape, "", ""))
	}

	// Group all transitions with the same states and merge their ranges into one label.
	edges := symboltable.NewRedBlack[State, symboltable.SymbolTable[State, rangeList]](CmpState, nil)

	for from, seq := range n.Transitions() {
		row, ok := edges.Get(from)
		if !ok {
			row = symboltable.NewRedBlack[State, rangeList](CmpState, nil)
			edges.Put(from, row)
		}

		for rs, states := range seq {
			for _, to := range states {
				ranges, ok := row.Get(to)
				if !ok {
					ranges = newRangeList()
					row.Put(to, ranges)
				}

				ranges.Add(rs...)
			}
		}
	}

	for from, row := range edges.All() {
		for to, ranges := range row.All() {
			from := fmt.Sprintf("%d", from)
			to := fmt.Sprintf("%d", to)

			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", ranges.String(), "", "", "", ""))
		}
	}

	return graph.DOT() + "\n"
}

// Runner constructs a new NFARunner for simulating (running) the NFA on input symbols.
func (n *NFA) Runner() *NFARunner {
	trans := symboltable.NewQuadraticHashTable(HashState, EqState, eqClassIDStatesTable, symboltable.HashOpts{})

	for s, stab := range n.trans.All() {
		stabClone := symboltable.NewQuadraticHashTable(hashClassID, eqClassID, EqStates, symboltable.HashOpts{})
		for cid, next := range stab.All() {
			stabClone.Put(cid, next)
		}

		trans.Put(s, stabClone)
	}

	return &NFARunner{
		start:  n.start,
		final:  n.final.Clone(),
		ranges: n.ranges.Clone(),
		trans:  trans,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// NFARunner is used for simulating (running) a NFA on input symbols.
// It is immutable and optimized for fast execution.
type NFARunner struct {
	start  State
	final  States
	ranges rangeMapping
	trans  symboltable.SymbolTable[State, symboltable.SymbolTable[classID, States]]
}

// εClosure returns the set of NFA states reachable from some NFA state s in set T on ε-transitions alone.
// εClosure(T) = Union(εClosure(s)) for all s ∈ T.
func (r *NFARunner) εClosure(T States) States {
	closure := T.Clone()

	stack := list.NewStack[State](64, nil) // Approximation
	for s := range T.All() {
		stack.Push(s)
	}

	for !stack.IsEmpty() {
		t, _ := stack.Pop()

		if next := r.next(t, E); next != nil {
			for u := range next.All() {
				if !closure.Contains(u) {
					closure.Add(u)
					stack.Push(u)
				}
			}
		}
	}

	return closure
}

// move returns the set of NFA states to which there is a transition on the given input from some state s in T.
func (r *NFARunner) move(T States, a Symbol) States {
	states := NewStates()

	for s := range T.All() {
		if next := r.next(s, a); next != nil {
			states = states.Union(next)
		}
	}

	return states
}

// next returns the set of next states from state s on the given input.
func (r *NFARunner) next(s State, a Symbol) States {
	if stab, ok := r.trans.Get(s); ok {
		if _, cid, ok := r.ranges.Find(a); ok {
			if next, ok := stab.Get(cid); ok {
				return next
			}
		}
	}

	return nil
}

// Next returns the next states from state s on input symbol a.
func (r *NFARunner) Next(s State, a Symbol) []State {
	if next := r.next(s, a); next != nil {
		return generic.Collect1(next.All())
	}

	return nil
}

// Accept determines whether an input string is recognized (accepted) by the NFA.
func (r *NFARunner) Accept(s String) bool {
	S := NewStates(r.start)

	for S = r.εClosure(S); len(s) > 0; s = s[1:] {
		S = r.εClosure(r.move(S, s[0]))
	}

	for s := range S.All() {
		if r.final.Contains(s) {
			return true
		}
	}

	return false

}

/* ------------------------------------------------------------------------------------------------------------------------ */
