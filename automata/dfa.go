package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
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
	cmpKey := generic.NewCompareFunc[State]()

	return &DFA{
		Start: start,
		Final: final,
		trans: symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, State]](cmpKey),
	}
}

// Add adds a new transition to the DFA.
func (d *DFA) Add(s State, a Symbol, next State) {
	if v, ok := d.trans.Get(s); ok {
		v.Put(a, next)
	} else {
		cmpKey := generic.NewCompareFunc[Symbol]()
		v = symboltable.NewRedBlack[Symbol, State](cmpKey)
		v.Put(a, next)
		d.trans.Put(s, v)
	}
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

// Join merges another DFA with the current one and returns the set of new merged states.
func (d *DFA) Join(dfa *DFA) States {
	// Find the maximum state number
	base := State(0)
	for _, s := range d.States() {
		if s > base {
			base = s
		}
	}

	// Use the maximum state number in the current DFA as the offset for the new states
	base += 1

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

	return states
}

// Accept determines whether or not an input string is recognized (accepted) by the DFA.
func (d *DFA) Accept(s String) bool {
	var curr State
	for curr = d.Start; len(s) > 0; s = s[1:] {
		curr = d.Next(curr, s[0])
	}

	return d.Final.Contains(curr)
}

// Graphviz returns the transition graph of the DFA in DOT Language format.
func (d *DFA) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "DFA", graphviz.RankDirLR, "", "", "")

	for _, state := range d.States() {
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
