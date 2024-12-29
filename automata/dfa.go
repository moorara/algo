package automata

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// DFA implements a deterministic finite automaton.
type DFA struct {
	Start State
	Final States
	trans doubleKeyMap[State, Symbol, State]
}

// NewDFA creates a new deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewDFA(start State, final States) *DFA {
	return &DFA{
		Start: start,
		Final: final,
		trans: symboltable.NewRedBlack[State, symboltable.SymbolTable[Symbol, State]](cmpState, eqSymbolState),
	}
}

// Add adds a new transition to the DFA.
func (d *DFA) Add(s State, a Symbol, next State) {
	tab, exist := d.trans.Get(s)
	if !exist {
		tab = symboltable.NewRedBlack[Symbol, State](cmpSymbol, eqState)
		d.trans.Put(s, tab)
	}

	tab.Put(a, next)
}

// Next returns the next state from a given state and for a given symbol.
func (d *DFA) Next(s State, a Symbol) State {
	if v, ok := d.trans.Get(s); ok {
		if next, ok := v.Get(a); ok {
			return next
		}
	}

	return State(-1)
}

// States returns the set of all states of the DFA.
func (d *DFA) States() States {
	states := States{}

	states = append(states, d.Start)
	states = append(states, d.Final...)

	for s := range d.trans.All() {
		if !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, v := range d.trans.All() {
		for _, s := range v.All() {
			if !states.Contains(s) {
				states = append(states, s)
			}
		}
	}

	return states
}

// Symbols returns the set of all input symbols of the DFA
func (d *DFA) Symbols() Symbols {
	symbols := Symbols{}

	for _, v := range d.trans.All() {
		for a := range v.All() {
			if a != E && !symbols.Contains(a) {
				symbols = append(symbols, a)
			}
		}
	}

	return symbols
}

// Accept determines whether or not an input string is recognized (accepted) by the DFA.
func (d *DFA) Accept(s String) bool {
	var curr State
	for curr = d.Start; len(s) > 0; s = s[1:] {
		curr = d.Next(curr, s[0])
	}

	return d.Final.Contains(curr)
}

// ToNFA constructs a new NFA accepting the same language as the DFA (every DFA is an NFA).
func (d *DFA) ToNFA() *NFA {
	nfa := NewNFA(d.Start, d.Final)
	for key, val := range d.trans.All() {
		s := key
		for key, val := range val.All() {
			a, next := key, val
			nfa.Add(s, a, States{next})
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

	// F
	F := States{}
	F = append(F, d.Final...)

	// S - F
	NF := States{}
	for _, s := range d.States() {
		if !d.Final.Contains(s) {
			NF = append(NF, s)
		}
	}

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
		for _, G := range Π.groups {
			Gtrans := Π.BuildGroupTrans(d, G)
			Πnew.PartitionAndAddGroups(Gtrans)
		}

		if Πnew.Equals(Π) {
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

	start, _ := Π.Rep(d.Start)

	final := States{}
	for _, f := range d.Final {
		g, _ := Π.Rep(f)
		final = append(final, g)
	}

	dfa := NewDFA(start, final)

	for _, G := range Π.groups {
		// Get any (the first) state in the group
		s := G.states[0]

		if v, ok := d.trans.Get(s); ok {
			for a, next := range v.All() {
				rep, _ := Π.Rep(next)
				dfa.Add(G.rep, a, rep)
			}
		}
	}

	return dfa
}

// WithoutDeadStates eliminates the dead states and all transitions to them.
// The subset construction and minimization algorithms sometimes produce a DFA with a single dead state.
//
// Strictly speaking, a DFA must have a transition from every state on every input symbol in its input alphabet.
// When we construct a DFA to be used in a lexical analyzer, we need to treat the dead state differently.
// We must know when there is no longer any possibility of recognizing a longer lexeme.
// Thus, we should always omit transitions to the dead state and eliminate the dead state itself.
func (d *DFA) WithoutDeadStates() *DFA {

	// 1. Construct a directed graph from the DFA with all the transitions reversed.
	adj := map[State]States{}
	for key, val := range d.trans.All() {
		s := key
		for _, val := range val.All() {
			t := val
			adj[t] = append(adj[t], s)
		}
	}

	// 2. Add a new state with transitions to all other states representing the final states of the DFA.
	u := State(-1)
	adj[u] = append(adj[u], d.Final...)

	// 3. Finally, we find all states reachable from this new state using a depth-first search (DFS).
	//    All other states not connected to this new state will be identified as dead states.
	visited := map[State]bool{}
	for s := range adj {
		visited[s] = false
	}

	dfs(adj, visited, u)

	deads := States{}
	for s, visited := range visited {
		if !visited {
			deads = append(deads, s)
		}
	}

	dfa := NewDFA(d.Start, d.Final)
	for key, val := range d.trans.All() {
		s := key
		for key, val := range val.All() {
			a, next := key, val
			if !deads.Contains(s) && !deads.Contains(next) {
				dfa.Add(s, a, next)
			}
		}
	}

	return dfa
}

func dfs(adj map[State]States, visited map[State]bool, s State) {
	visited[s] = true
	for _, t := range adj[s] {
		if !visited[t] {
			dfs(adj, visited, t)
		}
	}
}

// Equals determines whether or not two DFAs are identical in structure and labeling.
// Two DFAs are considered equal if they have the same start state, final states, and transitions.
//
// For isomorphic equality, structural equivalence with potentially different state names, use the Isomorphic method.
func (d *DFA) Equals(rhs *DFA) bool {
	return d.Start == rhs.Start &&
		d.Final.Equals(rhs.Final) &&
		d.trans.Equals(rhs.trans)
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
	if len(d.Final) != len(rhs.Final) {
		return false
	}

	// D₁ and D₂ must have the same number of states.
	states1, states2 := d.States(), rhs.States()
	if len(states1) != len(states2) {
		return false
	}

	// D₁ and D₂ must have the same input alphabet.
	symbols1, symbols2 := d.Symbols(), rhs.Symbols()
	if !symbols1.Equals(symbols2) {
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
	clone := make(States, len(states1))
	copy(clone, states1)

	// Methodically checking if any permutation of D₁ states is equal to D₂.
	return !generatePermutations(clone, 0, len(clone)-1, func(permutation States) bool {
		// Create a bijection between the states of D₁ and the current permutation of D₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(states1))
		for i, s := range states1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[d.Start]

		permutedFinal := make(States, len(d.Final))
		for i, f := range d.Final {
			permutedFinal[i] = bijection[f]
		}

		permutedDFA := NewDFA(permutedStart, permutedFinal)

		for s, table := range d.trans.All() {
			for a, t := range table.All() {
				ss, tt := bijection[s], bijection[t]
				permutedDFA.Add(ss, a, tt)
			}
		}

		// If the current permutation of D₁ is equal to D₂, we stop checking more permutations by returning false.
		// If the current permutation of D₁ is not equal to D₂, we continue with checking more permutations by returning true.
		return !permutedDFA.Equals(rhs)
	})
}

// getSortedDegreeSequence calculates the total degree (sum of in-degrees and out-degrees)
// for each state in the DFA and returns the degree sequence sorted in ascending order.
func (d *DFA) getSortedDegreeSequence() []int {
	totalDegrees := map[State]int{}
	for s, table := range d.trans.All() {
		for _, t := range table.All() {
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

// Graphviz returns the transition graph of the DFA in DOT Language format.
func (d *DFA) Graphviz() string {
	graph := graphviz.NewGraph(false, true, false, "DFA", graphviz.RankDirLR, "", "", "")

	states := d.States()
	sort.Quick(states, NewCompareFunc[State]())

	for _, state := range states {
		name := fmt.Sprintf("%d", state)
		label := fmt.Sprintf("%d", state)

		var shape graphviz.Shape
		if d.Final.Contains(state) {
			shape = graphviz.ShapeDoubleCircle
		} else {
			shape = graphviz.ShapeCircle
		}

		if state == d.Start {
			graph.AddNode(graphviz.NewNode("start", "", "", "", graphviz.StyleInvis, "", "", ""))
			graph.AddEdge(graphviz.NewEdge("start", name, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		graph.AddNode(graphviz.NewNode(name, "", label, "", "", shape, "", ""))
	}

	// Group all the transitions with the same states and combine their symbols into one label

	edges := symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[State, []string]](cmpState, nil)

	for key, val := range d.trans.All() {
		from := key
		tab, exist := edges.Get(from)
		if !exist {
			tab = symboltable.NewRedBlack[State, []string](cmpState, nil)
			edges.Put(from, tab)
		}

		for key, val := range val.All() {
			symbol, to := string(key), val
			vals, _ := tab.Get(to)
			vals = append(vals, symbol)
			tab.Put(to, vals)
		}
	}

	for key, val := range edges.All() {
		from := key
		for key, val := range val.All() {
			from := fmt.Sprintf("%d", from)
			to := fmt.Sprintf("%d", key)
			symbols := val

			sort.Quick(symbols, NewCompareFunc[string]())
			label := strings.Join(symbols, ",")

			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DotCode()
}
