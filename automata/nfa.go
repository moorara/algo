package automata

// NFABuilder implements the Builder design pattern for constructing NFA instances.
//
// The Builder pattern separates the construction of an NFA from its representation,
// This approach ensures the resulting NFA is immutable and optimized for simulation (running).
type NFABuilder struct {
	start State
	final States
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

// NFA represents a non-deterministic finite automaton.
// This NFA model is meant to be immutable once created.
type NFA struct {
	start State
	final States
}
