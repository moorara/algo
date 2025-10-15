package automata

// DFABuilder implements the Builder design pattern for constructing DFA instances.
//
// The Builder pattern separates the construction of a DFA from its representation.
// This approach ensures the resulting DFA is immutable and optimized for simulation (running).
type DFABuilder struct {
	start State
	final States
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

// DFA represents a deterministic finite automaton.
// This DFA model is meant to be immutable once created.
type DFA struct {
	start State
	final States
}
