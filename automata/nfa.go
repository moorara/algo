package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/symboltable"
)

// NTrans is the type for a non-deterministic finite automaton transition function (table).
type NTrans struct {
	tab symboltable.OrderedSymbolTable[State, symboltable.OrderedSymbolTable[Symbol, States]]
}

// NewNTrans creates a new non-deterministic finite automaton transition function (table).
func NewNTrans() *NTrans {
	cmpKey := generic.NewCompareFunc[State]()
	return &NTrans{
		tab: symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, States]](cmpKey),
	}
}

// Add adds a new transition for a non-deterministic finite automaton transition function (table).
func (t *NTrans) Add(s State, a Symbol, next States) {
	if v, ok := t.tab.Get(s); ok {
		v.Put(a, next)
	} else {
		cmpKey := generic.NewCompareFunc[Symbol]()
		v = symboltable.NewRedBlack[Symbol, States](cmpKey)
		v.Put(a, next)
		t.tab.Put(s, v)
	}
}

// Next returns the next set of states from a given state and for a given symbol.
func (t *NTrans) Next(s State, a Symbol) States {
	if v, ok := t.tab.Get(s); ok {
		if next, ok := v.Get(a); ok {
			return next
		}
	}

	return States{}
}

// Symbols returns the set of DFA states
func (t *NTrans) States() States {
	states := States{}

	for _, kv := range t.tab.KeyValues() {
		if s := kv.Key; !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, kv := range t.tab.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			for _, s := range kv.Val {
				if !states.Contains(s) {
					states = append(states, s)
				}
			}
		}
	}

	return states
}

// Symbols returns the set of NFA input symbols.
func (t *NTrans) Symbols() Symbols {
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

// NFA implements a non-deterministic finite automaton.
type NFA struct {
	trans *NTrans
	start State
	final States
}

// NewNFA creates a new non-deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewNFA(trans *NTrans, start State, final States) *NFA {
	return &NFA{
		trans: trans,
		start: start,
		final: final,
	}
}

// εClosure returns the set of NFA states reachable from some NFA state s in set T on ε-transitions alone.
// εClosure(T) = Union(εClosure(s)) for all s ∈ T.
func (n *NFA) εClosure(T States) States {
	closure := T

	stack := list.NewStack[State](1024, nil)
	for _, s := range T {
		stack.Push(s)
	}

	for !stack.IsEmpty() {
		t, _ := stack.Pop()
		for _, u := range n.trans.Next(t, E) {
			if !closure.Contains(u) {
				closure = append(closure, u)
				stack.Push(u)
			}
		}
	}

	return closure
}

// move returns the set of NFA states to which there is a transition on input symbol a from some state s in T.
func (n *NFA) move(T States, a Symbol) States {
	states := States{}
	for _, s := range T {
		U := n.trans.Next(s, a)
		states = append(states, U...)
	}

	return states
}

// UpdateFinal updates the final states of the current NFA.
// This is usually required after joining another NFA.
func (n *NFA) UpdateFinal(final States) {
	n.final = final
}

// Accept determines whether or not an input string is recognized (accepted) by the NFA.
func (n *NFA) Accept(s String) bool {
	var S States
	for S = n.εClosure(States{n.start}); len(s) > 0; s = s[1:] {
		S = n.εClosure(n.move(S, s[0]))
	}

	for _, s := range S {
		if n.final.Contains(s) {
			return true
		}
	}

	return false
}

// Join merges another NFA with the current one and returns the set of new merged states.
func (n *NFA) Join(nfa *NFA) States {
	// Find the maximum state number
	base := State(0)
	for _, s := range n.trans.States() {
		if s > base {
			base = s
		}
	}

	// Use the maximum state number in the current NFA as the offset for the new states
	base += 1

	for _, kv := range nfa.trans.tab.KeyValues() {
		s := base + kv.Key
		for _, kv := range kv.Val.KeyValues() {
			a := kv.Key

			next := make(States, len(kv.Val))
			for i, n := range kv.Val {
				next[i] = base + n
			}

			n.trans.Add(s, a, next)
		}
	}

	states := States{}
	for _, s := range nfa.trans.States() {
		states = append(states, base+s)
	}

	return states
}

// ToDFA constructs a new DFA accepting the same language as the NFA.
// It implements the subset construction algorithm.
func (n *NFA) ToDFA() *DFA {
	symbols := n.trans.Symbols()

	Dtrans := NewDTrans()
	Dstates := newMarkList[States](func(s, t States) bool {
		return s.Equals(t)
	})

	Dstates.AddUnmarked(n.εClosure(States{n.start}))

	for T, i := Dstates.GetUnmarked(); i >= 0; T, i = Dstates.GetUnmarked() {
		Dstates.MarkByIndex(i)

		for _, a := range symbols {
			U := n.εClosure(n.move(T, a))

			j := Dstates.Contains(U)
			if j == -1 {
				j = Dstates.AddUnmarked(U)
			}

			Dtrans.Add(State(i), a, State(j))
		}
	}

	Dstart := State(0)
	Dfinal := States{}

	for i, s := range Dstates.Values() {
		for _, f := range n.final {
			if s.Contains(f) {
				Dfinal = append(Dfinal, State(i))
				break
			}
		}
	}

	return &DFA{
		trans: Dtrans,
		start: Dstart,
		final: Dfinal,
	}
}

// Graphviz returns the transition graph of the NFA in DOT Language format.
func (n *NFA) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "NFA", graphviz.RankDirLR, "", "", graphviz.ShapeCircle)

	for _, state := range n.trans.States() {
		name := fmt.Sprintf("%d", state)
		label := fmt.Sprintf("%d", state)

		var shape graphviz.Shape
		if n.final.Contains(state) {
			shape = graphviz.ShapeDoubleCircle
		} else {
			shape = graphviz.ShapeCircle
		}

		if state == n.start {
			graph.AddNode(graphviz.NewNode("start", "", "", "", graphviz.StyleInvis, "", "", ""))
			graph.AddEdge(graphviz.NewEdge("start", name, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		graph.AddNode(graphviz.NewNode(name, "", label, "", "", shape, "", ""))
	}

	for _, kv := range n.trans.tab.KeyValues() {
		from := fmt.Sprintf("%d", kv.Key)

		for _, kv := range kv.Val.KeyValues() {
			var label string
			if symbol := kv.Key; symbol == E {
				label = "ε"
			} else {
				label = string(symbol)
			}

			for _, s := range kv.Val {
				to := fmt.Sprintf("%d", s)

				graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
			}
		}
	}

	return graph.DotCode()
}
