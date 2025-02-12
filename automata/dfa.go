package automata

import (
	"github.com/moorara/algo/symboltable"
)

// DFA implements a deterministic finite automaton.
type DFA struct {
	Start State
	Final States
	trans doubleKeyMap[State, Symbol, State]
}

// NewDFA creates a new deterministic finite automaton.
// Finite automata are recognizers; they simply say yes or no for each possible input string.
func NewDFA(start State, final []State) *DFA {
	return &DFA{
		Start: start,
		Final: NewStates(final...),
		trans: symboltable.NewRedBlack(cmpState, eqSymbolState),
	}
}

// newDFA creates a new deterministic finite automaton with a set of final states.
// This function is intended for internal use within this package only.
func newDFA(start State, final States) *DFA {
	return &DFA{
		Start: start,
		Final: final.Clone(),
		trans: symboltable.NewRedBlack(cmpState, eqSymbolState),
	}
}

// Equal determines whether or not two DFAs are identical in structure and labeling.
// Two DFAs are considered equal if they have the same start state, final states, and transitions.
//
// For isomorphic equality, structural equivalence with potentially different state names, use the Isomorphic method.
func (d *DFA) Equal(rhs *DFA) bool {
	return d.Start == rhs.Start &&
		d.Final.Equal(rhs.Final) &&
		d.trans.Equal(rhs.trans)
}

// Add inserts a new transition into the DFA.
func (d *DFA) Add(s State, a Symbol, next State) {
	strans, ok := d.trans.Get(s)
	if !ok {
		strans = symboltable.NewRedBlack(cmpSymbol, eqState)
		d.trans.Put(s, strans)
	}

	strans.Put(a, next)
}

// Next returns the next state based on the current state s and the input symbol a.
// If no valid transition exists, it returns an invalid state (-1).
func (d *DFA) Next(s State, a Symbol) State {
	if strans, ok := d.trans.Get(s); ok {
		if next, ok := strans.Get(a); ok {
			return next
		}
	}

	return State(-1)
}

// Accept determines whether or not an input string is recognized (accepted) by the DFA.
func (d *DFA) Accept(s String) bool {
	var curr State
	for curr = d.Start; len(s) > 0; s = s[1:] {
		curr = d.Next(curr, s[0])
	}

	return d.Final.Contains(curr)
}

// States returns the set of all states of the DFA.
func (d *DFA) States() []State {
	return sortStates(d.states())
}

func (d *DFA) states() States {
	states := NewStates(d.Start)
	states = states.Union(d.Final)

	for s, strans := range d.trans.All() {
		for _, t := range strans.All() {
			states.Add(s, t)
		}
	}

	return states
}

// Symbols returns the set of all input symbols of the DFA.
func (d *DFA) Symbols() []Symbol {
	return sortSymbols(d.symbols())
}

func (d *DFA) symbols() Symbols {
	symbols := NewSymbols()

	for _, trans := range d.trans.All() {
		for a := range trans.All() {
			if a != E {
				symbols.Add(a)
			}
		}
	}

	return symbols
}
