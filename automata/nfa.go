package automata

import (
	"github.com/moorara/algo/list"
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
