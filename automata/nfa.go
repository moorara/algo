package automata

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

// NFA implements a non-deterministic finite automaton.
type NFA struct {
	Start State
	Final States
	trans symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, States]]
}

// NewNFA creates a new non-deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewNFA(start State, final []State) *NFA {
	return &NFA{
		Start: start,
		Final: NewStates(final...),
		trans: symboltable.NewRedBlack(cmpState, eqSymbolStates),
	}
}

// newNFA creates a new non-deterministic finite automaton with a set of final states.
// This function is intended for internal use within this package only.
func newNFA(start State, final States) *NFA {
	return &NFA{
		Start: start,
		Final: final.Clone(),
		trans: symboltable.NewRedBlack(cmpState, eqSymbolStates),
	}
}

// εClosure returns the set of NFA states reachable from some NFA state s in set T on ε-transitions alone.
// εClosure(T) = Union(εClosure(s)) for all s ∈ T.
func (n *NFA) εClosure(T States) States {
	closure := T.Clone()

	stack := list.NewStack[State](128, nil)
	for s := range T.All() {
		stack.Push(s)
	}

	for !stack.IsEmpty() {
		t, _ := stack.Pop()

		if next := n.next(t, E); next != nil {
			for u := range next.All() {
				if !closure.Contains(u) {
					closure.Add(u)
					stack.Push(u)
				}
			}
		}
	}

	return closure
}

// move returns the set of NFA states to which there is a transition on input symbol a from some state s in T.
func (n *NFA) move(T States, a Symbol) States {
	states := NewStates()
	for s := range T.All() {
		if next := n.next(s, a); next != nil {
			states = states.Union(next)
		}
	}

	return states
}

// String returns a string representation of the NFA.
func (n *NFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", n.Start)
	fmt.Fprintf(&b, "Final states: ")

	for _, s := range sortStates(n.Final) {
		fmt.Fprintf(&b, "%d, ", s)
	}

	b.Truncate(b.Len() - 2)
	b.WriteString("\n")

	b.WriteString("Transitions:\n")
	for _, s := range n.States() {
		if next := n.next(s, E); next != nil {
			fmt.Fprintf(&b, "  (%d, ε) --> %s\n", s, next)
		}

		for _, a := range n.Symbols() {
			if next := n.next(s, a); next != nil {
				fmt.Fprintf(&b, "  (%d, %c) --> %s\n", s, a, next)
			}
		}
	}

	return b.String()
}

// Equal determines whether or not two NFAs are identical in structure and labeling.
// Two NFAs are considered equal if they have the same start state, final states, and transitions.
//
// For isomorphic equality, structural equivalence with potentially different state names, use the Isomorphic method.
func (n *NFA) Equal(rhs *NFA) bool {
	return n.Start == rhs.Start &&
		n.Final.Equal(rhs.Final) &&
		n.trans.Equal(rhs.trans)
}

// Add inserts a new transition into the NFA.
func (n *NFA) Add(s State, a Symbol, next []State) {
	strans, ok := n.trans.Get(s)
	if !ok {
		strans = symboltable.NewRedBlack(cmpSymbol, eqStateSet)
		n.trans.Put(s, strans)
	}

	states, ok := strans.Get(a)
	if !ok {
		states = NewStates()
		strans.Put(a, states)
	}

	states.Add(next...)
}

// Next returns the set of next states based on the current state s and the input symbol a.
// If no valid transition exists, it returns nil
func (n *NFA) Next(s State, a Symbol) []State {
	if next := n.next(s, a); next != nil {
		return sortStates(next)
	}

	return nil
}

func (n *NFA) next(s State, a Symbol) States {
	if strans, ok := n.trans.Get(s); ok {
		if next, ok := strans.Get(a); ok {
			return next
		}
	}

	return nil
}

// Accept determines whether or not an input string is recognized (accepted) by the NFA.
func (n *NFA) Accept(s String) bool {
	S := NewStates(n.Start)
	for S = n.εClosure(S); len(s) > 0; s = s[1:] {
		S = n.εClosure(n.move(S, s[0]))
	}

	for s := range S.All() {
		if n.Final.Contains(s) {
			return true
		}
	}

	return false
}

// States returns the set of all states of the NFA.
func (n *NFA) States() []State {
	return sortStates(n.states())
}

func (n *NFA) states() States {
	states := NewStates(n.Start)
	states = states.Union(n.Final)

	for s, strans := range n.trans.All() {
		for _, next := range strans.All() {
			states.Add(s)
			states = states.Union(next)
		}
	}

	return states
}

// Symbols returns the set of all input symbols of the NFA.
func (n *NFA) Symbols() []Symbol {
	return sortSymbols(n.symbols())
}

func (n *NFA) symbols() Symbols {
	symbols := NewSymbols()

	for _, trans := range n.trans.All() {
		for a := range trans.All() {
			if a != E {
				symbols.Add(a)
			}
		}
	}

	return symbols
}

// ToDFA constructs a new DFA accepting the same language as the NFA.
// It implements the subset construction algorithm.
//
// For more information and details, see "Compilers: Principles, Techniques, and Tools (2nd Edition)".
func (n *NFA) ToDFA() *DFA {
	symbols := n.Symbols()

	dfa := NewDFA(0, nil)
	Dstates := list.NewSoftQueue[States](func(s, t States) bool {
		return s.Equal(t)
	})

	// Initially, ε-closure(s0) is the only state in Dstates
	S0 := NewStates(n.Start)
	Dstates.Enqueue(n.εClosure(S0))

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
	dfa.Final = NewStates()

	for i, S := range Dstates.Values() {
		for f := range n.Final.All() {
			if S.Contains(f) {
				dfa.Final.Add(State(i))
				break // The accepting states of D are all those sets of N's states that include at least one accepting state of N
			}
		}
	}

	return dfa
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
	if n.Final.Size() != rhs.Final.Size() {
		return false
	}

	// N₁ and N₂ must have the same number of states.
	states1, states2 := n.States(), rhs.States()
	if len(states1) != len(states2) {
		return false
	}

	// N₁ and N₂ must have the same input alphabet.
	symbols1, symbols2 := n.symbols(), rhs.symbols()
	if !symbols1.Equal(symbols2) {
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
	states := make([]State, len(states1))
	copy(states, states1)

	// Methodically checking if any permutation of N₁ states is equal to N₂.
	return !generatePermutations(states, 0, len(states)-1, func(permutation []State) bool {
		// Create a bijection between the states of N₁ and the current permutation of N₁.
		// A bijection or bijective function is a type of function that creates a one-to-one correspondence between two sets (states1 ↔ permutation).
		bijection := make(map[State]State, len(states1))
		for i, s := range states1 {
			bijection[s] = permutation[i]
		}

		permutedStart := bijection[n.Start]

		permutedFinal := make([]State, 0, n.Final.Size())
		for f := range n.Final.All() {
			permutedFinal = append(permutedFinal, bijection[f])
		}

		permutedNFA := NewNFA(permutedStart, permutedFinal)

		for s, strans := range n.trans.All() {
			for a, ts := range strans.All() {
				ss := bijection[s]

				tts := make([]State, 0, ts.Size())
				for t := range ts.All() {
					tts = append(tts, bijection[t])
				}

				permutedNFA.Add(ss, a, tts)
			}
		}

		// If the current permutation of N₁ is equal to N₂, we stop checking more permutations by returning false.
		// If the current permutation of N₁ is not equal to N₂, we continue with checking more permutations by returning true.
		return !permutedNFA.Equal(rhs)
	})
}

// getSortedDegreeSequence calculates the total degree (sum of in-degrees and out-degrees)
// for each state in the NFA and returns the degree sequence sorted in ascending order.
func (n *NFA) getSortedDegreeSequence() []int {
	totalDegrees := map[State]int{}
	for s, strans := range n.trans.All() {
		for _, states := range strans.All() {
			for t := range states.All() {
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

// DOT generates a DOT representation of the transition graph of the NFA.
func (n *NFA) DOT() string {
	graph := dot.NewGraph(false, true, false, "NFA", dot.RankDirLR, "", "", dot.ShapeCircle)

	for _, s := range n.States() {
		name := fmt.Sprintf("%d", s)
		label := fmt.Sprintf("%d", s)

		if s == n.Start {
			graph.AddNode(dot.NewNode("start", "", "", "", dot.StyleInvis, "", "", ""))
			graph.AddEdge(dot.NewEdge("start", name, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		var shape dot.Shape
		if n.Final.Contains(s) {
			shape = dot.ShapeDoubleCircle
		}

		graph.AddNode(dot.NewNode(name, "", label, "", "", shape, "", ""))
	}

	/* Group all the transitions with the same states and combine their symbols into one label */

	edges := symboltable.NewRedBlack[State, symboltable.SymbolTable[State, []string]](cmpState, nil)

	for from, ftrans := range n.trans.All() {
		row, ok := edges.Get(from)
		if !ok {
			row = symboltable.NewRedBlack[State, []string](cmpState, nil)
			edges.Put(from, row)
		}

		for sym, states := range ftrans.All() {
			var label string
			if sym == E {
				label = "ε"
			} else {
				label = string(sym)
			}

			for to := range states.All() {
				labels, _ := row.Get(to)
				labels = append(labels, label)
				row.Put(to, labels)
			}
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
