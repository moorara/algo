package automata

import (
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// NFA implements a non-deterministic finite automaton.
type NFA struct {
	Start State
	Final States
	trans symboltable.OrderedSymbolTable[State, symboltable.OrderedSymbolTable[Symbol, States]]
}

// NewNFA creates a new non-deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewNFA(start State, final States) *NFA {
	return &NFA{
		Start: start,
		Final: final,
		trans: symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[Symbol, States]](cmpState, eqSymbolStates),
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
		for _, u := range n.Next(t, E) {
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
		U := n.Next(s, a)
		states = append(states, U...)
	}

	return states
}

// Add adds a new transition to the NFA.
func (n *NFA) Add(s State, a Symbol, next States) {
	tab, exist := n.trans.Get(s)
	if !exist {
		tab = symboltable.NewRedBlack[Symbol, States](cmpSymbol, eqStates)
		n.trans.Put(s, tab)
	}

	states, _ := tab.Get(a)
	states = append(states, next...)
	tab.Put(a, states)
}

// Next returns the next set of states from a given state and for a given symbol.
func (n *NFA) Next(s State, a Symbol) States {
	if v, ok := n.trans.Get(s); ok {
		if next, ok := v.Get(a); ok {
			return next
		}
	}

	return States{}
}

// Symbols returns the set of all states of the NFA.
func (n *NFA) States() States {
	states := States{}

	states = append(states, n.Start)
	for _, s := range n.Final {
		states = append(states, s)
	}

	for _, kv := range n.trans.KeyValues() {
		if s := kv.Key; !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, kv := range n.trans.KeyValues() {
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

// LastState returns the state with the maximum number.
// This information can be used for adding new states to the NFA.
func (n *NFA) LastState() State {
	max := State(0)
	for _, s := range n.States() {
		if s > max {
			max = s
		}
	}

	return max
}

// Symbols returns the set of all input symbols of the NFA.
func (n *NFA) Symbols() Symbols {
	symbols := Symbols{}

	for _, kv := range n.trans.KeyValues() {
		for _, kv := range kv.Val.KeyValues() {
			if a := kv.Key; a != E && !symbols.Contains(a) {
				symbols = append(symbols, a)
			}
		}
	}

	return symbols
}

// Join merges another NFA with the current one.
//
// The first return value is the set of all states of the merged NFA after merging.
// The second return value is the start (initial) state of the merged NFA after merging.
// The third return value is the set of final states of the merged NFA after merging.
func (n *NFA) Join(nfa *NFA) (States, State, States) {
	// Use the maximum state number plus one as the offset for the new states
	base := n.LastState() + 1

	for _, kv := range nfa.trans.KeyValues() {
		s := base + kv.Key
		for _, kv := range kv.Val.KeyValues() {
			a := kv.Key

			next := make(States, len(kv.Val))
			for i, n := range kv.Val {
				next[i] = base + n
			}

			n.Add(s, a, next)
		}
	}

	states := States{}
	for _, s := range nfa.States() {
		states = append(states, base+s)
	}

	start := base + nfa.Start

	final := States{}
	for _, s := range nfa.Final {
		final = append(final, base+s)
	}

	return states, start, final
}

// Accept determines whether or not an input string is recognized (accepted) by the NFA.
func (n *NFA) Accept(s String) bool {
	var S States
	for S = n.εClosure(States{n.Start}); len(s) > 0; s = s[1:] {
		S = n.εClosure(n.move(S, s[0]))
	}

	for _, s := range S {
		if n.Final.Contains(s) {
			return true
		}
	}

	return false
}

// ToDFA constructs a new DFA accepting the same language as the NFA.
// It implements the subset construction algorithm.
func (n *NFA) ToDFA() *DFA {
	symbols := n.Symbols()

	dfa := NewDFA(0, nil)
	Dstates := newMarkList[States](func(s, t States) bool {
		return s.Equals(t)
	})

	Dstates.AddUnmarked(n.εClosure(States{n.Start}))

	for T, i := Dstates.GetUnmarked(); i >= 0; T, i = Dstates.GetUnmarked() {
		Dstates.MarkByIndex(i)

		for _, a := range symbols {
			U := n.εClosure(n.move(T, a))

			j := Dstates.Contains(U)
			if j == -1 {
				j = Dstates.AddUnmarked(U)
			}

			dfa.Add(State(i), a, State(j))
		}
	}

	dfa.Start = State(0)
	dfa.Final = States{}

	for i, s := range Dstates.Values() {
		for _, f := range n.Final {
			if s.Contains(f) {
				dfa.Final = append(dfa.Final, State(i))
				break
			}
		}
	}

	return dfa
}

// Equals determines whether or not two NFAs are the same.
//
// TODO: Implement isomorphic equality.
func (n *NFA) Equals(nfa *NFA) bool {
	return n.Start == nfa.Start &&
		n.Final.Equals(nfa.Final) &&
		n.trans.Equals(nfa.trans)
}

// Graphviz returns the transition graph of the NFA in DOT Language format.
func (n *NFA) Graphviz() string {
	graph := graphviz.NewGraph(true, true, false, "NFA", graphviz.RankDirLR, "", "", graphviz.ShapeCircle)

	states := n.States()
	sort.Quick(states, generic.NewCompareFunc[State]())

	for _, state := range states {
		name := fmt.Sprintf("%d", state)
		label := fmt.Sprintf("%d", state)

		var shape graphviz.Shape
		if n.Final.Contains(state) {
			shape = graphviz.ShapeDoubleCircle
		} else {
			shape = graphviz.ShapeCircle
		}

		if state == n.Start {
			graph.AddNode(graphviz.NewNode("start", "", "", "", graphviz.StyleInvis, "", "", ""))
			graph.AddEdge(graphviz.NewEdge("start", name, graphviz.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		graph.AddNode(graphviz.NewNode(name, "", label, "", "", shape, "", ""))
	}

	for _, kv := range n.trans.KeyValues() {
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
