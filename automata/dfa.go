package automata

import (
	"bytes"
	"fmt"
	"iter"

	"github.com/moorara/algo/generic"
)

// DFABuilder implements the Builder design pattern for constructing DFA instances.
//
// The Builder pattern separates the construction of a complex object from its representation,
// allowing step-by-step configuration and assembly.
// This approach ensures the resulting DFA is constructed in an optimal and efficient way for simulation (running).
type DFABuilder struct {
	start State
	final States
	trans *dfaTransitionTable
}

// SetStart sets the start state of the DFA.
func (b *DFABuilder) SetStart(s State) *DFABuilder {
	b.start = s
	return b
}

// SetFinal sets the final (accepting) states of the DFA.
func (b *DFABuilder) SetFinal(ss ...State) *DFABuilder {
	b.final = NewStates(ss...)
	return b
}

// AddTransition adds transitions from state s to state next on all input symbols in the range [start, end].
func (b *DFABuilder) AddTransition(s State, start, end Symbol, next State) *DFABuilder {
	if b.trans == nil {
		b.trans = newDFATransitionTable(nil)
	}

	b.trans.Add(s, start, end, next)

	return b
}

// Build creates the DFA based on the configurations provided to the builder.
func (b *DFABuilder) Build() *DFA {
	return &DFA{
		start: b.start,
		final: b.final,
		trans: b.trans,
	}
}

// DFA represents a deterministic finite automaton.
// This DFA model is meant to be immutable once created.
type DFA struct {
	start State
	final States
	trans *dfaTransitionTable

	// Derived values calculated lazily
	states  States
	symbols Symbols
}

// String implements the fmt.Stringer interface.
func (d *DFA) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Start state: %d\n", d.start)
	fmt.Fprintf(&b, "Final states: ")

	for s := range d.final.All() {
		fmt.Fprintf(&b, "%d, ", s)
	}

	b.Truncate(b.Len() - 2)
	b.WriteString("\n")

	b.WriteString(d.trans.String())

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (d *DFA) Clone() *DFA {
	dd := &DFA{
		start: d.start,
		final: d.final.Clone(),
		trans: d.trans.Clone(),
	}

	if d.states != nil {
		dd.states = d.states.Clone()
	}

	if d.symbols != nil {
		dd.symbols = d.symbols.Clone()
	}

	return dd
}

// Equal implements the generic.Equaler interface.
func (d *DFA) Equal(rhs *DFA) bool {
	return d.start == rhs.start &&
		d.final.Equal(rhs.final) &&
		d.trans.Equal(rhs.trans)
}

// Start returns the start state of the DFA.
func (d *DFA) Start() State {
	return d.start
}

// Final returns the final (accepting) states of the DFA.
func (d *DFA) Final() []State {
	return generic.Collect1(d.final.All())
}

// States returns all states in the DFA.
func (d *DFA) States() []State {
	// Lazy initialization
	if d.states == nil {
		d.states = NewStates(d.start).Union(d.final)
		for s, pairs := range d.trans.All() {
			d.states.Add(s)
			for _, next := range pairs {
				d.states.Add(next)
			}
		}
	}

	return generic.Collect1(d.states.All())
}

// Symbols returns all symbols in the DFA.
func (d *DFA) Symbols() []Symbol {
	// Lazy initialization
	if d.symbols == nil {
		d.symbols = NewSymbols()
		for _, pairs := range d.trans.All() {
			for r := range pairs {
				for a := r.Start; a <= r.End; a++ {
					d.symbols.Add(a)
				}
			}
		}
	}

	return generic.Collect1(d.symbols.All())
}

// Transitions returns all transitions in the DFA.
func (d *DFA) Transitions() iter.Seq2[State, iter.Seq2[SymbolRange, State]] {
	return d.trans.All()
}

// TransitionsFrom returns all transitions from the given state in the DFA.
func (d *DFA) TransitionsFrom(s State) iter.Seq2[SymbolRange, State] {
	return d.trans.From(s)
}
