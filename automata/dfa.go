package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/symboltable"
)

// DFABuilder implements the Builder design pattern for constructing DFA instances.
//
// The Builder pattern separates the construction of a DFA from its representation.
// This approach ensures the resulting DFA is immutable and optimized for simulation (running).
type DFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, *disc.RangeMap[Symbol, State]]
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
		b.trans = symboltable.NewRedBlack[State, *disc.RangeMap[Symbol, State]](CmpState, nil)
	}

	ranges, ok := b.trans.Get(s)
	if !ok {
		ranges = disc.NewRangeMap[Symbol, State](EqState, nil)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		next,
	)

	return b
}

// Build creates the DFA based on the configurations provided to the builder.
func (b *DFABuilder) Build() *DFA {
	return &DFA{
		start: b.start,
		final: b.final,
	}
}

// DFA represents a deterministic finite automaton.
// This DFA model is meant to be immutable once created.
type DFA struct {
	start State
	final States
	// trans symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, State]]
}

// Equal implements the generic.Equaler interface.
func (d *DFA) Equal(rhs *DFA) bool {
	return d.start == rhs.start &&
		d.final.Equal(rhs.final) /* &&
		d.trans.Equal(rhs.trans) */
}

// Start returns the start state of the DFA.
func (d *DFA) Start() State {
	return d.start
}

// Final returns the final (accepting) states of the DFA.
func (d *DFA) Final() []State {
	return generic.Collect1(d.final.All())
}
