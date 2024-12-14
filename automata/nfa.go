package automata

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
	. "github.com/moorara/algo/generic"
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
		trans: symboltable.NewRedBlack[State, symboltable.SymbolTable[Symbol, States]](cmpState, eqSymbolStates),
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

// States returns the set of all states of the NFA.
func (n *NFA) States() States {
	states := States{}

	states = append(states, n.Start)
	states = append(states, n.Final...)

	for s := range n.trans.All() {
		if !states.Contains(s) {
			states = append(states, s)
		}
	}

	for _, v := range n.trans.All() {
		for _, states := range v.All() {
			for _, s := range states {
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

	for _, v := range n.trans.All() {
		for a := range v.All() {
			if a != E && !symbols.Contains(a) {
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
	start, final := State(0), State(1)
	star := NewNFA(start, States{final})
	factory := newStateFactory(final)

	for key, val := range n.trans.All() {
		s := factory.StateFor(0, key)
		for key, val := range val.All() {
			a := key

			next := make(States, len(val))
			for i, t := range val {
				next[i] = factory.StateFor(0, t)
			}

			star.Add(s, a, next)
		}
	}

	ss := factory.StateFor(0, n.Start)

	star.Add(start, E, States{ss})
	star.Add(start, E, States{final})

	for _, f := range n.Final {
		ff := factory.StateFor(0, f)
		star.Add(ff, E, States{ss})
		star.Add(ff, E, States{final})
	}

	return star
}

// Union constructs a new NFA that accepts the union of languages accepted by each individual NFA.
func (n *NFA) Union(ns ...*NFA) *NFA {
	start, final := State(0), State(1)
	union := NewNFA(start, States{final})
	factory := newStateFactory(final)

	nfas := append([]*NFA{n}, ns...)
	for id, nfa := range nfas {
		for key, val := range nfa.trans.All() {
			s := factory.StateFor(id, key)
			for key, val := range val.All() {
				a := key

				next := make(States, len(val))
				for i, t := range val {
					next[i] = factory.StateFor(id, t)
				}

				union.Add(s, a, next)
			}
		}

		ss := factory.StateFor(id, nfa.Start)
		union.Add(start, E, States{ss})

		for _, f := range nfa.Final {
			ff := factory.StateFor(id, f)
			union.Add(ff, E, States{final})
		}
	}

	return union
}

// Concat constructs a new NFA that accepts the concatenation of languages accepted by each individual NFA.
func (n *NFA) Concat(ns ...*NFA) *NFA {
	final := States{0}
	concat := NewNFA(0, final)
	factory := newStateFactory(0)

	nfas := append([]*NFA{n}, ns...)
	for id, nfa := range nfas {
		for key, val := range nfa.trans.All() {
			s := key

			// If s is the start state of the current NFA,
			// we need to map it to the previous NFA final states.
			var sp States
			if s == nfa.Start {
				sp = final
			} else {
				sp = States{factory.StateFor(id, s)}
			}

			for key, val := range val.All() {
				a, next := key, val

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

// Equals determines whether or not two NFAs are identical in structure and labeling.
// Two NFAs are considered equal if they have the same start state, final states, and transitions.
//
// For isomorphic equality, structural equivalence with potentially different state names, use the Isomorphic method.
func (n *NFA) Equals(rhs *NFA) bool {
	return n.Start == rhs.Start &&
		n.Final.Equals(rhs.Final) &&
		n.trans.Equals(rhs.trans)
}

// Isomorphic determines whether or not two NFAs are isomorphically the same.
//
// Two NFAs N₁ and N₂ are said to be isomorphic if there exists a bijection f: S(N₁) → S(N₂) between their state sets such that,
// for every input symbol a, there is a transition from state s to state t on input a in N₁
// if and only if there is a transition from state f(s) to state f(t) on input a in N₂.
//
// In simpler terms, the two NFAs have the same structure:
// one can be transformed into the other by renaming its states and preserving the transitions.
func (n *NFA) Isomorphic(rhs *NFA) bool {
	// N₁ and N₂ must have the same number of final states.
	if len(n.Final) != len(rhs.Final) {
		return false
	}

	// N₁ and N₂ must have the same number of states.
	states1, states2 := n.States(), rhs.States()
	if len(states1) != len(states2) {
		return false
	}

	// N₁ and N₂ must have the same input alphabet.
	symbols1, symbols2 := n.Symbols(), rhs.Symbols()
	if !symbols1.Equals(symbols2) {
		return false
	}

	// N₁ and N₂ must have the same sorted degree sequence.
	// len(degrees1) == len(degrees2) since N₁ and N₂ have the same number of states.
	degrees1, degrees2 := n.getSortedDegreeSequence(), rhs.getSortedDegreeSequence()
	for i := range degrees1 {
		if degrees1[i] != degrees2[i] {
			return false
		}
	}

	// Since generatePermutations uses backtracking and modifies the slice in-place, we need a copy.
	clone := make(States, len(states1))
	copy(clone, states1)

	// Methodically checking if any permutation of N₁ states is equal to N₂.
	return !generatePermutations(clone, 0, len(clone)-1, func(permutation States) bool {
		// Create a bijection between the states of N₁ and the current permutation of N₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(states1))
		for i, s := range states1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[n.Start]

		permutedFinal := make(States, len(n.Final))
		for i, f := range n.Final {
			permutedFinal[i] = bijection[f]
		}

		permutedNFA := NewNFA(permutedStart, permutedFinal)

		for s, table := range n.trans.All() {
			for a, ts := range table.All() {
				ss := bijection[s]

				tts := make(States, len(ts))
				for i, t := range ts {
					tts[i] = bijection[t]
				}

				permutedNFA.Add(ss, a, tts)
			}
		}

		// If the current permutation of N₁ is equal to N₂, we stop checking more permutations by returning false.
		// If the current permutation of N₁ is not equal to N₂, we continue with checking more permutations by returning true.
		return !permutedNFA.Equals(rhs)
	})
}

// getSortedDegreeSequence calculates the total degree (sum of in-degrees and out-degrees)
// for each state in the NFA and returns the degree sequence sorted in ascending order.
func (n *NFA) getSortedDegreeSequence() []int {
	totalDegrees := map[State]int{}
	for s, table := range n.trans.All() {
		for _, states := range table.All() {
			for _, t := range states {
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

// Graphviz returns the transition graph of the NFA in DOT Language format.
func (n *NFA) Graphviz() string {
	graph := graphviz.NewGraph(false, true, false, "NFA", graphviz.RankDirLR, "", "", graphviz.ShapeCircle)

	states := n.States()
	sort.Quick(states, NewCompareFunc[State]())

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

	edges := symboltable.NewRedBlack[State, symboltable.OrderedSymbolTable[State, []string]](cmpState, nil)

	for key, val := range n.trans.All() {
		from := key
		tab, exist := edges.Get(from)
		if !exist {
			tab = symboltable.NewRedBlack[State, []string](cmpState, nil)
			edges.Put(from, tab)
		}

		for key, val := range val.All() {
			var symbol string
			if key == E {
				symbol = "ε"
			} else {
				symbol = string(key)
			}

			for _, to := range val {
				vals, _ := tab.Get(to)
				vals = append(vals, symbol)
				tab.Put(to, vals)
			}
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
