package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// DFA implements a deterministic finite automaton.
type DFA struct {
	Start State
	Final States
	trans symboltable.OrderedSymbolTable[State, symboltable.OrderedSymbolTable[Symbol, State]]
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
	graph := graphviz.NewGraph(true, true, false, "DFA", graphviz.RankDirLR, "", "", "")

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
