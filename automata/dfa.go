package automata

import (
	"bytes"
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// DFA implements a deterministic finite automaton.
type DFA struct {
	Start State
	Final States
	trans symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, State]]
}

// NewDFA creates a new deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewDFA(start State, final []State) *DFA {
	return &DFA{
		Start: start,
		Final: NewStates(final...),
		trans: symboltable.NewRedBlack(CmpState, eqSymbolState),
	}
}

// newDFA creates a new deterministic finite automaton with a set of final states.
// This function is intended for internal use within this package only.
func newDFA(start State, final States) *DFA {
	return &DFA{
		Start: start,
		Final: final.Clone(),
		trans: symboltable.NewRedBlack(CmpState, eqSymbolState),
	}
}

// String returns a string representation of the DFA.
func (d *DFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", d.Start)
	fmt.Fprintf(&b, "Final states: ")

	for s := range d.Final.All() {
		fmt.Fprintf(&b, "%d, ", s)
	}

	b.Truncate(b.Len() - 2)
	b.WriteString("\n")

	b.WriteString("Transitions:\n")
	for _, s := range d.States() {
		for _, a := range d.Symbols() {
			if next := d.Next(s, a); next != -1 {
				fmt.Fprintf(&b, "  (%d, %c) --> %d\n", s, a, next)
			}
		}
	}

	return b.String()
}

// Equal determines whether or not two DFAs are identical in structure and labeling.
// Two DFAs are considered equal if they have the same start state, final states, and transitions.
//
// For isomorphic equality, structural equivalence with potentially different state names, use the Isomorphic method.
func (d *DFA) Equal(rhs *DFA) bool {
	return d.Start == rhs.Start &&
		d.Final.Equal(rhs.Final) &&
		d.trans.Equal(rhs.trans)
}

// Clone returns a deep copy of the DFA, ensuring the clone is independent of the original.
func (d *DFA) Clone() *DFA {
	dfa := newDFA(d.Start, d.Final)

	for s, strans := range d.trans.All() {
		for a, next := range strans.All() {
			dfa.Add(s, a, next)
		}
	}

	return dfa
}

// Add inserts a new transition into the DFA.
func (d *DFA) Add(s State, a Symbol, next State) {
	strans, ok := d.trans.Get(s)
	if !ok {
		strans = symboltable.NewRedBlack(CmpSymbol, EqState)
		d.trans.Put(s, strans)
	}

	strans.Put(a, next)
}

// Next returns the next state based on the current state s and the input symbol a.
// If no valid transition exists, it returns an invalid state (-1).
func (d *DFA) Next(s State, a Symbol) State {
	if strans, ok := d.trans.Get(s); ok {
		if next, ok := strans.Get(a); ok {
			return next
		}
	}

	return State(-1)
}

// Accept determines whether or not an input string is recognized (accepted) by the DFA.
func (d *DFA) Accept(s String) bool {
	var curr State
	for curr = d.Start; len(s) > 0; s = s[1:] {
		curr = d.Next(curr, s[0])
	}

	return d.Final.Contains(curr)
}

// States returns the set of all states of the DFA.
func (d *DFA) States() []State {
	return generic.Collect1(d.states().All())
}

func (d *DFA) states() States {
	states := NewStates(d.Start)
	states = states.Union(d.Final)

	for s, strans := range d.trans.All() {
		for _, t := range strans.All() {
			states.Add(s, t)
		}
	}

	return states
}

// Symbols returns the set of all input symbols of the DFA.
func (d *DFA) Symbols() []Symbol {
	return generic.Collect1(d.symbols().All())
}

func (d *DFA) symbols() Symbols {
	symbols := NewSymbols()

	for _, strans := range d.trans.All() {
		for a := range strans.All() {
			symbols.Add(a)
		}
	}

	return symbols
}

// Transitions returns an iterator sequence containing all transitions in the DFA.
func (d *DFA) Transitions() iter.Seq[*Transition[State]] {
	return func(yield func(*Transition[State]) bool) {
		for s, strans := range d.trans.All() {
			for a, next := range strans.All() {
				tr := &Transition[State]{
					State:  s,
					Symbol: a,
					Next:   next,
				}

				if !yield(tr) {
					return
				}
			}
		}
	}
}

// ToNFA constructs a new NFA accepting the same language as the DFA (every DFA is an NFA).
func (d *DFA) ToNFA() *NFA {
	nfa := newNFA(d.Start, d.Final)
	for s, strans := range d.trans.All() {
		for a, next := range strans.All() {
			nfa.Add(s, a, []State{next})
		}
	}

	return nfa
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

	F := d.Final.Clone()           // F
	NF := d.states().Difference(F) // S - F

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
			Gtrans := Π.BuildGroupTrans(d, G)
			Πnew.PartitionAndAddGroups(Gtrans)
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

	start := Π.Rep(d.Start)

	final := NewStates()
	for f := range d.Final.All() {
		g := Π.Rep(f)
		final.Add(g)
	}

	dfa := newDFA(start, final)

	for G := range Π.groups.All() {
		// Get any state in the group
		s, _ := G.States.FirstMatch(func(State) bool {
			return true
		})

		if v, ok := d.trans.Get(s); ok {
			for a, next := range v.All() {
				rep := Π.Rep(next)
				dfa.Add(G.rep, a, rep)
			}
		}
	}

	return dfa
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
	adj[u] = d.Final.Clone()

	// 3. Finally, we find all states reachable from this new state using a depth-first search (DFS).
	//    All other states not connected to this new state will be identified as dead states.
	visited := map[State]bool{}
	for s := range adj {
		visited[s] = false
	}

	dfs(adj, visited, u)

	deads := NewStates()
	for s, visited := range visited {
		if !visited {
			deads.Add(s)
		}
	}

	dfa := newDFA(d.Start, d.Final)
	for s, strans := range d.trans.All() {
		for a, t := range strans.All() {
			if !deads.Contains(s) && !deads.Contains(t) {
				dfa.Add(s, a, t)
			}
		}
	}

	return dfa
}

func dfs(adj map[State]States, visited map[State]bool, s State) {
	visited[s] = true

	if adj[s] != nil {
		for t := range adj[s].All() {
			if !visited[t] {
				dfs(adj, visited, t)
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

	visited[d.Start] = true
	queue.Enqueue(d.Start)
	sm.GetOrCreateState(0, d.Start)

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

	start := sm.GetOrCreateState(0, d.Start)
	dfa := NewDFA(start, nil)

	for f := range d.Final.All() {
		ff := sm.GetOrCreateState(0, f)
		dfa.Final.Add(ff)
	}

	for s, strans := range d.trans.All() {
		ss := sm.GetOrCreateState(0, s)
		for a, t := range strans.All() {
			tt := sm.GetOrCreateState(0, t)
			dfa.Add(ss, a, tt)
		}
	}

	return dfa
}

// Isomorphic determines whether or not two DFAs are isomorphically the same.
//
// Two DFAs D₁ and D₂ are said to be isomorphic if there exists a bijection f: S(D₁) → S(D₂) between their state sets such that,
// for every input symbol a, there is a transition from state s to state t on input a in D₁
// if and only if there is a transition from state f(s) to state f(t) on input a in D₂.
//
// In simpler terms, the two DFAs have the same structure:
// one can be transformed into the other by renaming its states and preserving the transitions.
func (d *DFA) Isomorphic(rhs *DFA) bool {
	// D₁ and D₂ must have the same number of final states.
	if d.Final.Size() != rhs.Final.Size() {
		return false
	}

	// D₁ and D₂ must have the same number of states.
	states1, states2 := d.States(), rhs.States()
	if len(states1) != len(states2) {
		return false
	}

	// D₁ and D₂ must have the same input alphabet.
	symbols1, symbols2 := d.symbols(), rhs.symbols()
	if !symbols1.Equal(symbols2) {
		return false
	}

	// D₁ and D₂ must have the same sorted degree sequence.
	// len(degrees1) == len(degrees2) since D₁ and D₂ have the same number of states.
	degrees1, degrees2 := d.getSortedDegreeSequence(), rhs.getSortedDegreeSequence()
	for i := range degrees1 {
		if degrees1[i] != degrees2[i] {
			return false
		}
	}

	// Since generatePermutations uses backtracking and modifies the slice in-place, we need a copy.
	states := make([]State, len(states1))
	copy(states, states1)

	// Methodically checking if any permutation of D₁ states is equal to D₂.
	return !generatePermutations(states, 0, len(states)-1, func(permutation []State) bool {
		// Create a bijection between the states of D₁ and the current permutation of D₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(states1))
		for i, s := range states1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[d.Start]

		permutedFinal := make([]State, 0, d.Final.Size())
		for f := range d.Final.All() {
			permutedFinal = append(permutedFinal, bijection[f])
		}

		permutedDFA := NewDFA(permutedStart, permutedFinal)

		for s, strans := range d.trans.All() {
			for a, t := range strans.All() {
				ss, tt := bijection[s], bijection[t]
				permutedDFA.Add(ss, a, tt)
			}
		}

		// If the current permutation of D₁ is equal to D₂, we stop checking more permutations by returning false.
		// If the current permutation of D₁ is not equal to D₂, we continue with checking more permutations by returning true.
		return !permutedDFA.Equal(rhs)
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

// DOT generates a DOT representation of the transition graph of the DFA.
func (d *DFA) DOT() string {
	graph := dot.NewGraph(false, true, false, "DFA", dot.RankDirLR, "", "", dot.ShapeCircle)

	for _, s := range d.States() {
		name := fmt.Sprintf("%d", s)
		label := fmt.Sprintf("%d", s)

		if s == d.Start {
			graph.AddNode(dot.NewNode("start", "", "", "", dot.StyleInvis, "", "", ""))
			graph.AddEdge(dot.NewEdge("start", name, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		var shape dot.Shape
		if d.Final.Contains(s) {
			shape = dot.ShapeDoubleCircle
		}

		graph.AddNode(dot.NewNode(name, "", label, "", "", shape, "", ""))
	}

	/* Group all the transitions with the same states and combine their symbols into one label */

	edges := symboltable.NewRedBlack[State, symboltable.SymbolTable[State, []string]](CmpState, nil)

	for from, ftrans := range d.trans.All() {
		row, ok := edges.Get(from)
		if !ok {
			row = symboltable.NewRedBlack[State, []string](CmpState, nil)
			edges.Put(from, row)
		}

		for sym, to := range ftrans.All() {
			labels, _ := row.Get(to)
			labels = append(labels, string(sym))
			row.Put(to, labels)
		}
	}

	for from, fedges := range edges.All() {
		for to, labels := range fedges.All() {
			from := fmt.Sprintf("%d", from)
			to := fmt.Sprintf("%d", to)

			sort.Quick(labels, generic.NewCompareFunc[string]())
			label := strings.Join(labels, ",")

			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DOT()
}

// CombineDFA merges multiple deterministic finite automata (DFAs) into a single DFA
// that recognizes the union of the languages accepted by each individual DFA.
//
// The process follows these steps:
//
//  1. Convert each DFA into a nondeterministic finite automaton (NFA).
//  2. Construct a single NFA that accepts the union of all individual NFAs.
//  3. Convert the unified NFA into a DFA.
//  4. Remove dead states and eliminate transitions leading to them.
//  5. Reassign state indices to maintain a compact representation.
//
// Throughout this process, the method tracks the final states of each original DFA,
// returning a slice where each index maps to the final states of the corresponding DFA.
//
// Note: This method does not minimize the resulting DFA, as minimization would merge final states,
// making it impossible to distinguish between the final states of the individual DFAs.
//
// This function is commonly used when constructing a DFA for lexical analysis.
func CombineDFA(ds ...*DFA) (*DFA, [][]State) {
	var finalMap [][]State

	// 1. Convert all DFAs to NFAs.
	ns := make([]*NFA, len(ds))
	for i, d := range ds {
		ns[i] = d.ToNFA()
	}

	// 2. Construct a new NFA that accepts the union of the languages accepted by each NFA.
	var union *NFA

	{
		start, final := State(0), State(1)
		union = NewNFA(start, []State{final})
		finalMap = make([][]State, len(ns))

		sm := newStateManager(final)

		for id, n := range ns {
			for s, strans := range n.trans.All() {
				ss := sm.GetOrCreateState(id, s)

				for a, states := range strans.All() {
					next := make([]State, 0, states.Size())
					for t := range states.All() {
						tt := sm.GetOrCreateState(id, t)
						next = append(next, tt)
					}

					union.Add(ss, a, next)
				}
			}

			ss := sm.GetOrCreateState(id, n.Start)
			union.Add(start, E, []State{ss})

			for f := range n.Final.All() {
				ff := sm.GetOrCreateState(id, f)
				union.Add(ff, E, []State{final})
				finalMap[id] = append(finalMap[id], ff)
			}
		}
	}

	// 3. Convert the NFA into a DFA.
	var combined *DFA

	{
		combined = NewDFA(0, nil)

		Dstates := list.NewSoftQueue(func(s, t States) bool {
			return s.Equal(t)
		})

		// Initially, ε-closure(s0) is the only state in Dstates
		S0 := NewStates(union.Start)
		Dstates.Enqueue(union.εClosure(S0))

		for T, i := Dstates.Dequeue(); i >= 0; T, i = Dstates.Dequeue() {
			for _, a := range union.Symbols() { // for each input symbol c
				U := union.εClosure(union.move(T, a))

				// If U is not in Dstates, add U to Dstates
				j := Dstates.Contains(U)
				if j == -1 {
					j = Dstates.Enqueue(U)
				}

				combined.Add(State(i), a, State(j))
			}
		}

		combined.Final = NewStates()

		for i, S := range Dstates.Values() {
			for f := range union.Final.All() {
				if S.Contains(f) {
					combined.Final.Add(State(i))
					break
				}
			}
		}

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
	combined = combined.EliminateDeadStates()

	// 5. Reassign state indices.
	var reindexed *DFA

	{
		sm := newStateManager(-1)

		visited := map[State]bool{}
		queue := list.NewQueue[State](64, nil)

		visited[combined.Start] = true
		queue.Enqueue(combined.Start)
		sm.GetOrCreateState(0, combined.Start)

		for !queue.IsEmpty() {
			s, _ := queue.Dequeue()
			if adj, ok := combined.trans.Get(s); ok {
				for _, t := range adj.All() {
					if !visited[t] {
						visited[t] = true
						queue.Enqueue(t)
						sm.GetOrCreateState(0, t)
					}
				}
			}
		}

		start := sm.GetOrCreateState(0, combined.Start)
		reindexed = NewDFA(start, nil)

		for f := range combined.Final.All() {
			ff := sm.GetOrCreateState(0, f)
			reindexed.Final.Add(ff)
		}

		for s, strans := range combined.trans.All() {
			ss := sm.GetOrCreateState(0, s)
			for a, t := range strans.All() {
				tt := sm.GetOrCreateState(0, t)
				reindexed.Add(ss, a, tt)
			}
		}

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

	return reindexed, finalMap
}
