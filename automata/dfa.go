package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

type dfaTrans symboltable.OrderedSymbolTable[State, symboltable.OrderedSymbolTable[Symbol, State]]

// DFA implements a deterministic finite automaton.
type DFA struct {
	Start State
	Final States
	trans dfaTrans
}

// NewDFA creates a new deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewDFA(start State, final States) *DFA {
	return &DFA{
		Start: start,
		Final: final,
		trans: symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, State]](cmpState, eqSymbolState),
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

// Symbols returns the set of all states of the DFA.
func (d *DFA) States() States {
	states := States{}

	states = append(states, d.Start)
	for _, s := range d.Final {
		states = append(states, s)
	}

	for _, kv := range d.trans.KeyValues() {
		if s := kv.Key; !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, kv := range d.trans.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			if s := kv.Val; !states.Contains(s) {
				states = append(states, s)
			}
		}
	}

	return states
}

// LastState returns the state with the maximum number.
// This information can be used for adding new states to the DFA.
func (d *DFA) LastState() State {
	max := State(0)
	for _, s := range d.States() {
		if s > max {
			max = s
		}
	}

	return max
}

// Symbols returns the set of all input symbols of the DFA
func (d *DFA) Symbols() Symbols {
	symbols := Symbols{}

	for _, kv := range d.trans.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			if a := kv.Key; a != E && !symbols.Contains(a) {
				symbols = append(symbols, a)
			}
		}
	}

	return symbols
}

// Join merges another DFA with the current one.
//
// The first return value is the set of all states of the merged DFA after merging.
// The second return value is the start (initial) state of the merged DFA after merging.
// The third return value is the set of final states of the merged DFA after merging.
func (d *DFA) Join(dfa *DFA) (States, State, States) {
	// Use the maximum state number plus one as the offset for the new states
	base := d.LastState() + 1

	for _, kv := range dfa.trans.KeyValues() {
		s := base + kv.Key
		for _, kv := range kv.Val.KeyValues() {
			a, next := kv.Key, base+kv.Val
			d.Add(s, a, next)
		}
	}

	states := States{}
	for _, s := range dfa.States() {
		states = append(states, base+s)
	}

	start := base + dfa.Start

	final := States{}
	for _, s := range dfa.Final {
		final = append(final, base+s)
	}

	return states, start, final
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
	for _, kv := range d.trans.KeyValues() {
		S := kv.Key
		for _, kv := range kv.Val.KeyValues() {
			a, T := kv.Key, kv.Val
			nfa.Add(S, a, States{T})
		}
	}

	return nfa
}

// Minimize creates a unique DFA with the minimum number of states.
//
// The minimization algorithm sometimes produces a DFA with one dead state.
// This state is not accepting and transfers to itself on each input symbol.
//
// For more details, see Compilers: Principles, Techniques, and Tools (2nd Edition).
func (d *DFA) Minimize() *DFA {
	eqFunc := func(a, b State) bool { return a == b }
	setEqFunc := func(a, b set.Set[State]) bool { return a.Equals(b) }

	// 1. Start with an initial partition P with two groups,
	//    F and S - F, the accepting and non-accepting states.

	S := set.New[State](eqFunc)
	S.Add(d.States()...)

	F := set.New[State](eqFunc)
	F.Add(d.Final...)

	NF := S.Difference(F)

	P := set.New[set.Set[State]](setEqFunc)
	P.Add(NF, F)

	// 2. Initially, let Pnew = P.
	//    For (each group G of P) {
	//      Partition G into subgroups such that two states s and t are in the same subgroup
	//      if and only if for all input symbols a, states s and t have transitions on a to states in the same group of P
	//      (at worst, a state will be in a subgroup by itself).
	//
	//      Replace G in Pnew by the set of all subgroups formed.
	//    }
	//
	// 3. If Pnew = P, let Pfinal = P and continue with step (4).
	//    Otherwise, repeat step (2) with Pnew in place of P.

	for {
		Pnew := set.New[set.Set[State]](setEqFunc)

		for _, G := range P.Members() { // For every group in the current partition
			gtrans := d.createGroupTrans(P, G)
			populateSubgroups(Pnew, gtrans)
		}

		if Pnew.Equals(P) {
			break
		}

		P = Pnew
	}

	// 4. Choose one state in each group of Pfinal as the representative for that group.
	//    The representatives will be the states of the minimum-state DFA D′.
	//    The other components of D′ are constructed as follows:
	//
	//    (a) The start state of D′ is the representative of the group containing the start state of D.
	//    (b) The accepting states of D′ are the representatives of those groups that contain an accepting state of D
	//        (each group contains either only accepting states, or only non-accepting states).
	//    (c) Let s be the representative of some group G of Pfinal, and let the transition of D from s on input a be to state t.
	//        Let r be the representative of t's group H. Then in D′, there is a transition from s to r on input a.

	start, _ := groupRep(P, d.Start)

	final := States{}
	for _, f := range d.Final {
		g, _ := groupRep(P, f)
		final = append(final, g)
	}

	dfa := NewDFA(start, final)

	for s, G := range P.Members() {
		g := G.Members()[0] // G is non-empty
		if tab, ok := d.trans.Get(g); ok {
			for _, kv := range tab.KeyValues() {
				a, next := kv.Key, kv.Val
				rep, _ := groupRep(P, next)
				dfa.Add(State(s), a, rep)
			}
		}
	}

	return dfa
}

// createGroupTrans create a map of states to symbols to the current partition's group representatives (instead of next states).
func (d *DFA) createGroupTrans(P set.Set[set.Set[State]], G set.Set[State]) dfaTrans {
	gtrans := symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, State]](cmpState, eqSymbolState)

	for _, s := range G.Members() { // For every state in the current group
		strans := symboltable.NewRedBlack[Symbol, State](cmpSymbol, eqState)

		// Create a map of symbols to the current partition's group representatives (instead of next states)
		if stab, ok := d.trans.Get(s); ok {
			for _, kv := range stab.KeyValues() {
				a, next := kv.Key, kv.Val
				if rep, ok := groupRep(P, next); ok {
					strans.Put(a, rep)
				}
			}
		}

		gtrans.Put(s, strans)
	}

	return gtrans
}

// populateSubgroups creates new subgroups based on the transition map of a group and add them to the new partition.
func populateSubgroups(Pnew set.Set[set.Set[State]], gtrans dfaTrans) {
	eqFunc := func(a, b State) bool { return a == b }

	kvs := gtrans.KeyValues()
	for i := 0; i < len(kvs); i++ {
		s, sreps := kvs[i].Key, kvs[i].Val

		if _, ok := groupRep(Pnew, s); !ok { // If s is not already added to the new partition
			// Create a new group in the new partition
			H := set.New[State](eqFunc)
			H.Add(s)

			// Add all other state with same symbol->rep map to the new group
			for j := 1; j < len(kvs); j++ {
				t, treps := kvs[j].Key, kvs[j].Val

				if sreps.Equals(treps) {
					H.Add(t)
				}
			}

			Pnew.Add(H)
		}
	}
}

// groupRep returns the group representaive for a state.
func groupRep(P set.Set[set.Set[State]], s State) (State, bool) {
	for i, G := range P.Members() {
		for _, state := range G.Members() {
			if state == s {
				return State(i), true
			}
		}
	}

	return -1, false
}

// Equals determines whether or not two DFAs are the same.
//
// TODO: Implement isomorphic equality.
func (d *DFA) Equals(dfa *DFA) bool {
	return d.Start == dfa.Start &&
		d.Final.Equals(dfa.Final) &&
		d.trans.Equals(dfa.trans)
}

// Graphviz returns the transition graph of the DFA in DOT Language format.
func (d *DFA) Graphviz() string {
	graph := graphviz.NewGraph(false, true, false, "DFA", graphviz.RankDirLR, "", "", "")

	states := d.States()
	sort.Quick(states, generic.NewCompareFunc[State]())

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

	for _, kv := range d.trans.KeyValues() {
		from := fmt.Sprintf("%d", kv.Key)

		for _, kv := range kv.Val.KeyValues() {
			label := string(kv.Key)
			to := fmt.Sprintf("%d", kv.Val)

			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DotCode()
}
