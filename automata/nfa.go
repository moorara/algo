package automata

import (
	"fmt"
	"strings"

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
	trans doubleKeyMap[State, Symbol, States]
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

// Star constructs a new NFA that accepts the Kleene closure of the language accepted by the NFA.
func (n *NFA) Star() *NFA {
	star := NewNFA(n.Start, n.Final)
	for _, kv := range n.trans.KeyValues() {
		s := kv.Key
		for _, kv := range kv.Val.KeyValues() {
			a, next := kv.Key, kv.Val
			star.Add(s, a, next)
		}
	}

	star.Add(star.Start, E, star.Final)
	for _, f := range star.Final {
		star.Add(f, E, States{star.Start})
	}

	return star
}

// Union constructs a new NFA that accepts the union of languages accepted by each individual NFA.
func (n *NFA) Union(ns ...*NFA) *NFA {
	start := State(0)
	union := NewNFA(start, States{})
	factory := newStateFactory()

	nfas := append([]*NFA{n}, ns...)
	for id, nfa := range nfas {
		for _, kv := range nfa.trans.KeyValues() {
			s := kv.Key

			// If s is the start state of the current NFA,
			// we need to map it to the start state of the union NFA.
			var sp State
			if s == nfa.Start {
				sp = start
			} else {
				sp = factory.StateFor(id, s)
			}

			for _, kv := range kv.Val.KeyValues() {
				a, next := kv.Key, kv.Val

				// If any of the next state is the start state of the current NFA,
				// we need to map it to the start state of the union NFA.
				var nextp States
				for _, s := range next {
					if s == nfa.Start {
						nextp = append(nextp, start)
					} else {
						nextp = append(nextp, factory.StateFor(id, s))
					}
				}

				// Add new transition
				union.Add(sp, a, nextp)
			}
		}

		// Update the final states of the union NFA
		for _, f := range nfa.Final {
			union.Final = append(union.Final, factory.StateFor(id, f))
		}
	}

	return union
}

// Concat constructs a new NFA that accepts the concatenation of languages accepted by each individual NFA.
func (n *NFA) Concat(ns ...*NFA) *NFA {
	final := States{0}
	concat := NewNFA(0, final)
	factory := newStateFactory()

	nfas := append([]*NFA{n}, ns...)
	for id, nfa := range nfas {
		for _, kv := range nfa.trans.KeyValues() {
			s := kv.Key

			// If s is the start state of the current NFA,
			// we need to map it to the previous NFA final states.
			var sp States
			if s == nfa.Start {
				sp = final
			} else {
				sp = States{factory.StateFor(id, s)}
			}

			for _, kv := range kv.Val.KeyValues() {
				a, next := kv.Key, kv.Val

				// If any of the next state is the start state of the current NFA,
				// we need to map it to the previous NFA final states.
				var nextp States
				for _, s := range next {
					if s == nfa.Start {
						nextp = append(nextp, final...)
					} else {
						nextp = append(nextp, factory.StateFor(id, s))
					}
				}

				// Add new transitions
				for _, s := range sp {
					concat.Add(s, a, nextp)
				}
			}
		}

		// Update the current final states
		final = States{}
		for _, f := range nfa.Final {
			final = append(final, factory.StateFor(id, f))
		}
	}

	concat.Final = final

	return concat
}

// ToDFA constructs a new DFA accepting the same language as the NFA.
// It implements the subset construction algorithm.
//
// For more details, see Compilers: Principles, Techniques, and Tools (2nd Edition).
func (n *NFA) ToDFA() *DFA {
	symbols := n.Symbols()

	dfa := NewDFA(0, nil)
	Dstates := list.NewSoftQueue[States](func(s, t States) bool {
		return s.Equals(t)
	})

	// Initially, ε-closure(s0) is the only state in Dstates
	Dstates.Enqueue(n.εClosure(States{n.Start}))

	for T, i := Dstates.Dequeue(); i >= 0; T, i = Dstates.Dequeue() {
		for _, a := range symbols { // for each input symbol c
			U := n.εClosure(n.move(T, a))

			// If U is not in Dstates, add U to Dstates
			j := Dstates.Contains(U)
			if j == -1 {
				j = Dstates.Enqueue(U)
			}

			dfa.Add(State(i), a, State(j))
		}
	}

	dfa.Start = State(0)
	dfa.Final = States{}

	for i, S := range Dstates.Values() {
		for _, f := range n.Final {
			if S.Contains(f) {
				dfa.Final = append(dfa.Final, State(i))
				break // The accepting states of D are all those sets of N's states that include at least one accepting state of N
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
	graph := graphviz.NewGraph(false, true, false, "NFA", graphviz.RankDirLR, "", "", graphviz.ShapeCircle)

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

	// Group all the transitions with the same states and combine their symbols into one label

	var edges doubleKeyMap[State, State, []string]
	edges = symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[State, []string]](cmpState, nil)

	for _, kv := range n.trans.KeyValues() {
		from := kv.Key
		tab, exist := edges.Get(from)
		if !exist {
			tab = symboltable.NewRedBlack[State, []string](cmpState, nil)
			edges.Put(from, tab)
		}

		for _, kv := range kv.Val.KeyValues() {
			var symbol string
			if kv.Key == E {
				symbol = "ε"
			} else {
				symbol = string(kv.Key)
			}

			for _, to := range kv.Val {
				vals, _ := tab.Get(to)
				vals = append(vals, symbol)
				tab.Put(to, vals)
			}
		}
	}

	for _, kv := range edges.KeyValues() {
		from := kv.Key
		for _, kv := range kv.Val.KeyValues() {
			from := fmt.Sprintf("%d", from)
			to := fmt.Sprintf("%d", kv.Key)
			symbols := kv.Val

			sort.Quick(symbols, generic.NewCompareFunc[string]())
			label := strings.Join(symbols, ",")

			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DotCode()
}
