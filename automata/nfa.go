package automata

import (
	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/symboltable"
)

// NFABuilder implements the Builder design pattern for constructing NFA instances.
//
// The Builder pattern separates the construction of an NFA from its representation,
// This approach ensures the resulting NFA is immutable and optimized for simulation (running).
type NFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, *disc.RangeMap[Symbol, States]]
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
		b.trans = symboltable.NewRedBlack[State, *disc.RangeMap[Symbol, States]](CmpState, nil)
	}

	ranges, ok := b.trans.Get(s)
	if !ok {
		ranges = disc.NewRangeMap[Symbol, States](EqStates, nil)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		NewStates(next...),
	)

	return b
}

// Build creates the NFA based on the configurations provided to the builder.
func (b *NFABuilder) Build() *NFA {
	return &NFA{
		start: b.start,
		final: b.final,
	}
}

// NFA represents a non-deterministic finite automaton.
// This NFA model is meant to be immutable once created.
type NFA struct {
	start State
	final States
	// trans symboltable.SymbolTable[State, symboltable.SymbolTable[Symbol, States]]
}

// Equal implements the generic.Equaler interface.
func (n *NFA) Equal(rhs *NFA) bool {
	return n.start == rhs.start &&
		n.final.Equal(rhs.final) /* &&
		n.trans.Equal(rhs.trans) */
}

// Start returns the start state of the NFA.
func (n *NFA) Start() State {
	return n.start
}

// Final returns the final (accepting) states of the NFA.
func (n *NFA) Final() []State {
	return generic.Collect1(n.final.All())
}
