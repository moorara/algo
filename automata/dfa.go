package automata

import (
	"slices"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

/* ------------------------------------------------------------------------------------------------------------------------ */

// dfaTransitionEnds represents the states involved in a DFA transition.
type dfaTransitionEnds struct {
	State State
	Next  State
}

var cmpDFATransitionEnds = func(lhs, rhs dfaTransitionEnds) int {
	if c := CmpState(lhs.State, rhs.State); c != 0 {
		return c
	}

	return CmpState(lhs.Next, rhs.Next)
}

// dfaTransition represents a transition from one state to another state on a range of input symbols.
type dfaTransition struct {
	Ends  dfaTransitionEnds
	Range disc.Range[Symbol]
}

// dfaBound represents a boundary (start or end) of a range associated with a DFA transition.
type dfaBound struct {
	Value      Symbol
	Start      bool
	Transition dfaTransition
}

// dfaTransitionVector represents a DFA transition vector.
type dfaTransitionVector set.Set[dfaTransitionEnds]

func newDFATransitionVector() dfaTransitionVector {
	return set.NewSorted(cmpDFATransitionEnds)
}

var cmpDFATransitionVector = func(lhs, rhs dfaTransitionVector) int {
	v1 := generic.Collect1(lhs.All())
	v2 := generic.Collect1(rhs.All())

	for i := 0; i < len(v1) && i < len(v2); i++ {
		if c := cmpDFATransitionEnds(v1[i], v2[i]); c != 0 {
			return c
		}
	}

	return len(v1) - len(v2)
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// DFABuilder implements the Builder design pattern for constructing DFA instances.
//
// The Builder pattern separates the construction of a DFA from its representation.
// This approach ensures the resulting DFA is immutable and optimized for simulation (running).
type DFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, disc.RangeMap[Symbol, State]]
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
		b.trans = symboltable.NewRedBlack[State, disc.RangeMap[Symbol, State]](CmpState, nil)
	}

	ranges, ok := b.trans.Get(s)
	if !ok {
		ranges = disc.NewRangeMap[Symbol, State](EqState, nil, nil)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		next,
	)

	return b
}

// Build constructs the DFA based on the configurations provided to the builder.
//
// This method reduces the size of the alphabet by computing the equivalence classes of input symbols based on the transition function.
// The resulting DFA will recognize the same language, but with a minimized alphabet optimized for faster simulation and smaller memory footprint.
//
// Formally, given a DFA = (Q, Σ, δ, q₀, F), we want compute a partition of Σ into k equivalence classes such that
//
//	P = {C₁, C₂, ..., Cₖ} where each Cᵢ ⊆ Σ and ∪ Cᵢ = Σ
//	∀ a, b ∈ Cᵢ, ∀ q ∈ Q: δ(q, a) = δ(q, b)
//
// Informally, two symbols are equivalent if they lead to the same next state from any given state (every state treats them the same).
func (b *DFABuilder) Build() *DFA {
	// Collect boundaries for all ranges.
	bounds := make([]dfaBound, 0, b.trans.Size()*2*2) // Approximation
	for s, ranges := range b.trans.All() {
		for r, next := range ranges.All() {
			bounds = append(bounds,
				dfaBound{r.Lo, true, dfaTransition{dfaTransitionEnds{s, next}, r}},
				dfaBound{r.Hi + 1, false, dfaTransition{dfaTransitionEnds{s, next}, r}},
			)
		}
	}

	// Sort boundaries in increasing order.
	slices.SortFunc(bounds, func(lhs, rhs dfaBound) int {
		return int(lhs.Value - rhs.Value)
	})

	active := make(map[dfaTransitionEnds]disc.Range[Symbol])
	partition := make([]generic.KeyValue[disc.Range[Symbol], dfaTransitionVector], 0, len(bounds)/2) // Approximation

	for i := 0; i < len(bounds)-1; i++ {
		// Maintain a list of transitions seen between the beginning and end of the current range.
		if b := bounds[i]; b.Start {
			active[b.Transition.Ends] = b.Transition.Range
		} else {
			delete(active, b.Transition.Ends)
		}

		lo := bounds[i].Value
		hi := bounds[i+1].Value - 1

		// This effectively dedups the boundaries.
		if lo <= hi {
			// Collect all transitions seen between the beginning and end of the current range.
			vector := newDFATransitionVector()
			for ends := range active {
				vector.Add(ends)
			}

			// Skip empty classes corresponding to ranges between the end of one boundary and start of the next one.
			if vector.Size() > 0 {
				partition = append(partition, generic.KeyValue[disc.Range[Symbol], dfaTransitionVector]{
					Key: disc.Range[Symbol]{Lo: lo, Hi: hi},
					Val: vector,
				})
			}
		}
	}

	nextCID := classID(0)
	transitionVectors := symboltable.NewRedBlack(cmpDFATransitionVector, eqClassID)
	equivalenceClasses := disc.NewRangeMap[Symbol, classID](eqClassID, nil, nil)

	// Group ranges by their transition vectors to form equivalence classes.
	for _, sub := range partition {
		cid, ok := transitionVectors.Get(sub.Val)
		if !ok {
			cid = nextCID
			nextCID++
			transitionVectors.Put(sub.Val, cid)
		}

		equivalenceClasses.Add(sub.Key, cid)
	}

	return &DFA{
		start:   b.start,
		final:   b.final,
		classes: equivalenceClasses,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// DFA represents a deterministic finite automaton.
//
// A DFA is defined by a 5-tuple (Q, Σ, δ, q₀, F) where:
//
//   - Q is a finite set of states.
//   - Σ is a finite set of input symbols (alphabet).
//   - δ: Q × Σ → Q is the transition function.
//   - q₀ ∈ Q is the initial (start) state.
//   - F ⊆ Q is the set of accepting (final) states.
//
// This DFA model is meant to be immutable once created.
type DFA struct {
	start   State
	final   States
	classes disc.RangeMap[Symbol, classID]
	trans   symboltable.SymbolTable[State, symboltable.SymbolTable[classID, State]]
}

// Equal implements the generic.Equaler interface.
func (d *DFA) Equal(rhs *DFA) bool {
	return d.start == rhs.start &&
		d.final.Equal(rhs.final) &&
		d.classes.Equal(rhs.classes) /* &&
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
