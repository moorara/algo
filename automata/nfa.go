package automata

import (
	"bytes"
	"fmt"
	"iter"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/generic"
)

// NFABuilder implements the Builder design pattern for constructing NFA instances.
//
// The Builder pattern separates the construction of a complex object from its representation,
// allowing step-by-step configuration and assembly.
// This approach ensures the resulting NFA is constructed in an optimal and efficient way for simulation (running).
type NFABuilder struct {
	start State
	final States
	trans *nfaTransitionTable
}

// SetStart sets the start state of the NFA.
func (b *NFABuilder) SetStart(s State) *NFABuilder {
	b.start = s
	return b
}

// SetFinal sets the final (accepting) states of the NFA.
func (b *NFABuilder) SetFinal(ss ...State) *NFABuilder {
	b.final = NewStates(ss...)
	return b
}

// AddTransition adds transitions from state s to states next on all input symbols in the range [start, end].
func (b *NFABuilder) AddTransition(s State, start, end Symbol, next []State) *NFABuilder {
	if b.trans == nil {
		b.trans = newNFATransitionTable(nil)
	}

	b.trans.Add(s, start, end, next)

	return b
}

// Build creates the NFA based on the configurations provided to the builder.
func (b *NFABuilder) Build() *NFA {
	return &NFA{
		start: b.start,
		final: b.final,
		trans: b.trans,
	}
}

// NFA represents a non-deterministic finite automaton.
// This NFA model is meant to be immutable once created.
type NFA struct {
	start State
	final States
	trans *nfaTransitionTable

	// Derived values calculated lazily
	states  []State
	symbols []SymbolRange
}

// String implements the fmt.Stringer interface.
func (n *NFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", n.start)
	fmt.Fprintf(&b, "Final states: ")

	for s := range n.final.All() {
		fmt.Fprintf(&b, "%d, ", s)
	}

	b.Truncate(b.Len() - 2)
	b.WriteString("\n")

	b.WriteString(n.trans.String())

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (n *NFA) Clone() *NFA {
	nn := &NFA{
		start: n.start,
		final: n.final.Clone(),
		trans: n.trans.Clone(),
	}

	if n.states != nil {
		nn.states = make([]State, len(n.states))
		copy(nn.states, n.states)
	}

	if n.symbols != nil {
		nn.symbols = make([]SymbolRange, len(n.symbols))
		copy(nn.symbols, n.symbols)
	}

	return nn
}

// Equal implements the generic.Equaler interface.
func (n *NFA) Equal(rhs *NFA) bool {
	return n.start == rhs.start &&
		n.final.Equal(rhs.final) &&
		n.trans.Equal(rhs.trans)
}

// Start returns the start state of the NFA.
func (n *NFA) Start() State {
	return n.start
}

// Final returns the final (accepting) states of the NFA.
func (n *NFA) Final() []State {
	return generic.Collect1(n.final.All())
}

// States returns all states in the NFA.
func (n *NFA) States() []State {
	// Lazy initialization
	if n.states == nil {
		states := NewStates(n.start).Union(n.final)
		for s, pairs := range n.trans.All() {
			states.Add(s)
			for _, next := range pairs {
				states.Add(next...)
			}
		}

		n.states = generic.Collect1(states.All())
	}

	return n.states
}

// Symbols returns all symbol ranges in the NFA.
func (n *NFA) Symbols() []SymbolRange {
	// Lazy initialization
	if n.symbols == nil {
		n.symbols = n.trans.SymbolRanges()
	}

	return n.symbols
}

// Transitions returns all transitions in the NFA.
func (n *NFA) Transitions() iter.Seq2[State, iter.Seq2[SymbolRange, []State]] {
	return n.trans.All()
}

// TransitionsFrom returns all transitions from the given state in the NFA.
func (n *NFA) TransitionsFrom(s State) iter.Seq2[SymbolRange, []State] {
	return n.trans.From(s)
}

// DOT generates a DOT representation of the NFA transition graph for visualization.
func (n *NFA) DOT() string {
	graph := dot.NewGraph(false, true, false, "NFA", dot.RankDirLR, "", "", dot.ShapeCircle)

	for _, s := range n.States() {
		name := fmt.Sprintf("%d", s)
		label := fmt.Sprintf("%d", s)

		if s == n.start {
			graph.AddNode(dot.NewNode("start", "", "", "", dot.StyleInvis, "", "", ""))
			graph.AddEdge(dot.NewEdge("start", name, dot.EdgeTypeDirected, "", "", "", "", "", ""))
		}

		var shape dot.Shape
		if n.final.Contains(s) {
			shape = dot.ShapeDoubleCircle
		}

		graph.AddNode(dot.NewNode(name, "", label, "", "", shape, "", ""))
	}

	for s, pairs := range n.trans.All() {
		for r, next := range pairs {
			for _, t := range next {
				from := fmt.Sprintf("%d", s)
				to := fmt.Sprintf("%d", t)

				graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", r.String(), "", "", "", ""))
			}
		}
	}

	return graph.DOT()
}
