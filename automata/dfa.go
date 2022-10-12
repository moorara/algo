package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/symboltable"
)

// DTrans is the type for a deterministic finite automaton transition function (table).
type DTrans struct {
	tab symboltable.OrderedSymbolTable[State, symboltable.OrderedSymbolTable[Symbol, State]]
}

// NewNFATrans creates a new deterministic finite automaton transition function (table).
func NewDTrans() *DTrans {
	cmpKey := generic.NewCompareFunc[State]()
	return &DTrans{
		tab: symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, State]](cmpKey),
	}
}

// Add adds a new transition for a deterministic finite automaton transition function (table).
func (t *DTrans) Add(s State, a Symbol, next State) {
	if v, ok := t.tab.Get(s); ok {
		v.Put(a, next)
	} else {
		cmpKey := generic.NewCompareFunc[Symbol]()
		v = symboltable.NewRedBlack[Symbol, State](cmpKey)
		v.Put(a, next)
		t.tab.Put(s, v)
	}
}

// Next returns the next state from a given state and for a given symbol.
func (t *DTrans) Next(s State, a Symbol) State {
	if v, ok := t.tab.Get(s); ok {
		if next, ok := v.Get(a); ok {
			return next
		}
	}

	return State(-1)
}

// Symbols returns the set of DFA states
func (t *DTrans) States() States {
	states := States{}

	for _, kv := range t.tab.KeyValues() {
		if s := kv.Key; !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, kv := range t.tab.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			if s := kv.Val; !states.Contains(s) {
				states = append(states, s)
			}
		}
	}

	return states
}

// Symbols returns the set of DFA input symbols.
func (t *DTrans) Symbols() Symbols {
	symbols := Symbols{}

	for _, kv := range t.tab.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			if a := kv.Key; a != E && !symbols.Contains(a) {
				symbols = append(symbols, a)
			}
		}
	}

	return symbols
}

// DFA implements a deterministic finite automaton.
type DFA struct {
	trans *DTrans
	start State
	final States
}

// NewDFA creates a new deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewDFA(trans *DTrans, start State, final States) *DFA {
	return &DFA{
		trans: trans,
		start: start,
		final: final,
	}
}

// Accept determines whether or not an input string is recognized (accepted) by the DFA.
func (d *DFA) Accept(s String) bool {
	var curr State
	for curr = d.start; len(s) > 0; s = s[1:] {
		curr = d.trans.Next(curr, s[0])
	}

	return d.final.Contains(curr)
}

// Graphviz returns the transition graph of the DFA in DOT Language format.
func (d *DFA) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "DFA", graphviz.RankDirLR, "", "", "")

	for _, state := range d.trans.States() {
		name := fmt.Sprintf("%d", state)
		label := fmt.Sprintf("%d", state)

		var shape graphviz.Shape
		if d.final.Contains(state) {
			shape = graphviz.ShapeDoubleCircle
		} else {
			shape = graphviz.ShapeCircle
		}

		if state == d.start {
			graph.AddNode(graphviz.NewNode("start", "", "", "", graphviz.StyleInvis, "", "", ""))
			graph.AddEdge(graphviz.NewEdge("start", name, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		graph.AddNode(graphviz.NewNode(name, "", label, "", "", shape, "", ""))
	}

	for _, kv := range d.trans.tab.KeyValues() {
		from := fmt.Sprintf("%d", kv.Key)

		for _, kv := range kv.Val.KeyValues() {
			label := string(kv.Key)
			to := fmt.Sprintf("%d", kv.Val)

			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DotCode()
}
