package automata

import (
	"slices"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

/* ------------------------------------------------------------------------------------------------------------------------ */

// nfaTransitionEnds represents the states involved in a NFA transition.
type nfaTransitionEnds struct {
	State State
	Next  States
}

var cmpNFATransitionEnds = func(lhs, rhs nfaTransitionEnds) int {
	if c := CmpState(lhs.State, rhs.State); c != 0 {
		return c
	}

	// Next States are sorted set and can be compared element-wise.

	s1 := generic.Collect1(lhs.Next.All())
	s2 := generic.Collect1(rhs.Next.All())

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if c := CmpState(s1[i], s2[i]); c != 0 {
			return c
		}
	}

	return len(s1) - len(s2)
}

// nfaTransition  represents a transition from one state to a set of states on a range of input symbols.
type nfaTransition struct {
	Ends  nfaTransitionEnds
	Range disc.Range[Symbol]
}

// nfaBound represents a boundary (start or end) of a range associated with a NFA transition.
type nfaBound struct {
	Value      Symbol
	Start      bool
	Transition nfaTransition
}

// nfaTransitionVector represents a NFA transition vector.
type nfaTransitionVector set.Set[nfaTransitionEnds]

func newNFATransitionVector() nfaTransitionVector {
	return set.NewSorted(cmpNFATransitionEnds)
}

var cmpNFATransitionVector = func(lhs, rhs nfaTransitionVector) int {
	v1 := generic.Collect1(lhs.All())
	v2 := generic.Collect1(rhs.All())

	for i := 0; i < len(v1) && i < len(v2); i++ {
		if c := cmpNFATransitionEnds(v1[i], v2[i]); c != 0 {
			return c
		}
	}

	return len(v1) - len(v2)
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// NFABuilder implements the Builder design pattern for constructing NFA instances.
//
// The Builder pattern separates the construction of an NFA from its representation,
// This approach ensures the resulting NFA is immutable and optimized for simulation (running).
type NFABuilder struct {
	start State
	final States
	trans symboltable.SymbolTable[State, disc.RangeMap[Symbol, States]]
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
		b.trans = symboltable.NewRedBlack[State, disc.RangeMap[Symbol, States]](CmpState, nil)
	}

	ranges, ok := b.trans.Get(s)
	if !ok {
		opts := &disc.RangeMapOpts[Symbol, States]{Resolve: unionStates}
		ranges = disc.NewRangeMap[Symbol, States](EqStates, opts, nil)
		b.trans.Put(s, ranges)
	}

	ranges.Add(
		disc.Range[Symbol]{Lo: start, Hi: end},
		NewStates(next...),
	)

	return b
}

// Build constructs the NFA based on the configurations provided to the builder.
//
// This method reduces the size of the alphabet by computing the equivalence classes of input symbols based on the transition function.
// The resulting NFA will recognize the same language, but with a minimized alphabet optimized for faster simulation and smaller memory footprint.
//
// Formally, given a NFA = (Q, Σ, δ, q₀, F), we want compute a partition of Σ into k equivalence classes such that
//
//	P = {C₁, C₂, ..., Cₖ} where each Cᵢ ⊆ Σ and ∪ Cᵢ = Σ
//	∀ a, b ∈ Cᵢ, ∀ q ∈ Q: δ(q, a) = δ(q, b)
//
// Informally, two symbols are equivalent if they lead to the same next state from any given state (every state treats them the same).
func (b *NFABuilder) Build() *NFA {
	// Collect boundaries for all ranges.
	bounds := make([]nfaBound, 0, b.trans.Size()*2*2) // Approximation
	for s, ranges := range b.trans.All() {
		for r, next := range ranges.All() {
			bounds = append(bounds,
				nfaBound{r.Lo, true, nfaTransition{nfaTransitionEnds{s, next}, r}},
				nfaBound{r.Hi + 1, false, nfaTransition{nfaTransitionEnds{s, next}, r}},
			)
		}
	}

	// Sort boundaries in increasing order.
	slices.SortFunc(bounds, func(lhs, rhs nfaBound) int {
		return int(lhs.Value - rhs.Value)
	})

	active := make(map[nfaTransitionEnds]disc.Range[Symbol])
	partition := make([]generic.KeyValue[disc.Range[Symbol], nfaTransitionVector], 0, len(bounds)/2) // Approximation

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
			vector := newNFATransitionVector()
			for ends := range active {
				vector.Add(ends)
			}

			// Skip empty classes corresponding to ranges between the end of one boundary and start of the next one.
			if vector.Size() > 0 {
				partition = append(partition, generic.KeyValue[disc.Range[Symbol], nfaTransitionVector]{
					Key: disc.Range[Symbol]{Lo: lo, Hi: hi},
					Val: vector,
				})
			}
		}
	}

	nextCID := classID(0)
	transitionVectors := symboltable.NewRedBlack(cmpNFATransitionVector, eqClassID)
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

	return &NFA{
		start:   b.start,
		final:   b.final,
		classes: equivalenceClasses,
	}
}

/* ------------------------------------------------------------------------------------------------------------------------ */

// NFA represents a non-deterministic finite automaton.
//
// A NFA is defined by a 5-tuple (Q, Σ, δ, q₀, F) where:
//
//   - Q is a finite set of states.
//   - Σ is a finite set of input symbols (alphabet).
//   - δ: Q × Σ → P(Q) is the transition function.
//   - q₀ ∈ Q is the initial (start) state.
//   - F ⊆ Q is the set of accepting (final) states.
//
// This NFA model is meant to be immutable once created.
type NFA struct {
	start   State
	final   States
	classes disc.RangeMap[Symbol, classID]
	trans   symboltable.SymbolTable[State, symboltable.SymbolTable[classID, States]]
}

// Equal implements the generic.Equaler interface.
func (n *NFA) Equal(rhs *NFA) bool {
	return n.start == rhs.start &&
		n.final.Equal(rhs.final) &&
		n.classes.Equal(rhs.classes) /* &&
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
