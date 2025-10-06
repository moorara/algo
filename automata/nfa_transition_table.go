package automata

import (
	"bytes"
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/symboltable"
)

// rangeStates represents a pair of a symbol range and a set of states.
// It is used in NFA transitions to group transitions that share the same range of input symbols.
type rangeStates struct {
	SymbolRange
	States
}

// String implements the fmt.Stringer interface.
func (rs rangeStates) String() string {
	return fmt.Sprintf("%s → %s", rs.SymbolRange, rs.States)
}

// Equal implements the generic.Equaler interface.
func (rs rangeStates) Equal(rhs rangeStates) bool {
	return rs.SymbolRange.Equal(rhs.SymbolRange) && rs.States.Equal(rhs.States)
}

// nfaTransitionTable implements a transition table for non-deterministic finite automata (NFA).
// It is used in NFA to manage transitions from one state to a set of states over ranges of input symbols.
type nfaTransitionTable struct {
	table symboltable.SymbolTable[State, []rangeStates]
}

// newNFATransitionTable creates a new instance of nfaTransitionTable.
func newNFATransitionTable(trans map[State][]rangeStates) *nfaTransitionTable {
	table := symboltable.NewAVL[State, []rangeStates](CmpState, nil)

	for s, pairs := range trans {
		for _, pair := range pairs {
			pair.Validate()
		}

		slices.SortFunc(pairs, func(lhs, rhs rangeStates) int {
			return int(lhs.Start) - int(rhs.Start)
		})

		table.Put(s, mergeRangeStatesSortedList(pairs))
	}

	return &nfaTransitionTable{
		table: table,
	}
}

// String implements the fmt.Stringer interface.
func (t *nfaTransitionTable) String() string {
	var b bytes.Buffer

	b.WriteString("Transitions:\n")

	for s, pairs := range t.table.All() {
		for _, pair := range pairs {
			fmt.Fprintf(&b, "  %d --%s--> %s\n", s, pair.SymbolRange, pair.States)
		}
	}

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (t *nfaTransitionTable) Clone() *nfaTransitionTable {
	tt := &nfaTransitionTable{
		table: symboltable.NewAVL[State, []rangeStates](CmpState, nil),
	}

	for s, pairs := range t.table.All() {
		pp := make([]rangeStates, len(pairs))
		for i, pair := range pairs {
			pp[i] = rangeStates{
				SymbolRange: pair.SymbolRange,
				States:      pair.States.Clone(),
			}
		}
		tt.table.Put(s, pp)
	}

	return tt
}

// Equal implements the generic.Equaler interface.
func (t *nfaTransitionTable) Equal(rhs *nfaTransitionTable) bool {
	if t.table.Size() != rhs.table.Size() {
		return false
	}

	for s, pairs := range t.table.All() {
		rhsPairs, ok := rhs.table.Get(s)
		if !ok || len(pairs) != len(rhsPairs) {
			return false
		}

		for i, pair := range pairs {
			if !pair.Equal(rhsPairs[i]) {
				return false
			}
		}
	}

	return true
}

// All returns all transitions in the table.
func (t *nfaTransitionTable) All() iter.Seq2[State, iter.Seq2[SymbolRange, []State]] {
	return func(yield func(State, iter.Seq2[SymbolRange, []State]) bool) {
		for s := range t.table.All() {
			if !yield(s, t.From(s)) {
				return
			}
		}
	}
}

// From returns all transitions from the given state in the table.
func (t *nfaTransitionTable) From(s State) iter.Seq2[SymbolRange, []State] {
	return func(yield func(SymbolRange, []State) bool) {
		if pairs, ok := t.table.Get(s); ok {
			for _, pair := range pairs {
				states := generic.Collect1(pair.States.All())
				if !yield(pair.SymbolRange, states) {
					return
				}
			}
		}
	}
}

// Next returns the next state for the given state and input symbol.
func (t *nfaTransitionTable) Next(s State, a Symbol) ([]State, bool) {
	if pairs, ok := t.table.Get(s); ok {
		if i, ok := searchRangeStatesSortedList(pairs, a); ok {
			states := generic.Collect1(pairs[i].States.All())
			return states, true
		}
	}

	return nil, false
}

// Add inserts a new transition to the NFA transition table.
// It will merge any overlapping or adjacent ranges as necessary.
// The states associated with any overlapping range will be overridden by the new states given in the new range.
func (t *nfaTransitionTable) Add(s State, start, end Symbol, next []State) {
	new := rangeStates{
		SymbolRange{Start: start, End: end},
		NewStates(next...),
	}

	new.Validate()

	pairs, ok := t.table.Get(s)
	if !ok {
		pairs = make([]rangeStates, 0, 1)
	}

	// Find the insertion point
	i, ok := searchRangeStatesSortedList(pairs, new.Start)
	if ok {
		i++
	}

	// Insert the new entry at position i
	pairs = append(pairs, rangeStates{})
	copy(pairs[i+1:], pairs[i:])
	pairs[i] = new

	// Merge overlapping or adjacent ranges
	t.table.Put(s, mergeRangeStatesSortedList(pairs))
}

// searchRangeStatesSortedList performs a binary search to find the index of the range that contains the given symbol.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func searchRangeStatesSortedList(pairs []rangeStates, a Symbol) (int, bool) {
	lo, hi := 0, len(pairs)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if a < pairs[mid].Start {
			hi = mid - 1
		} else if pairs[mid].End < a {
			lo = mid + 1
		} else {
			return mid, true
		}
	}

	return lo, false
}

// mergeRangeStatesSortedList merges overlapping or adjacent ranges in a sorted list of range-states pairs.
func mergeRangeStatesSortedList(pairs []rangeStates) []rangeStates {
	merged := make([]rangeStates, 0, len(pairs))

	for _, curr := range pairs {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if curr.Start <= last.End {
			if curr.End < last.End {
				if last.States.Equal(curr.States) {
					// Case curr.Start < last.End && curr.End < last.End && last.States == curr.States:
					//
					//   last:  |_____|_____|_____|  States: {1,2}    ---->    |_________________|  States: {1,2}
					//   curr:        |_____|        States: {1,2}    ---->
					//
					// Impossible case of curr.Start == last.End && curr.End < last.End
					//
				} else {
					// Case curr.Start < last.End && curr.End < last.End && last.States != curr.States:
					//
					//   last:  |_____|_____|_____|  States: {1,2}    ---->    |____||     ||    |  States: {1,2}
					//   curr:        |_____|        States: {3,4}    ---->          |_____||    |  States: {3,4}
					//                                                ---->                 |____|  States: {1,2}
					//
					// Impossible case of curr.Start == last.End && curr.End < last.End
					//

					lastEnd := last.End
					last.End = curr.Start - 1
					merged = append(merged, curr)
					merged = append(merged, rangeStates{
						SymbolRange{Start: curr.End + 1, End: lastEnd},
						last.States.Clone(),
					})
				}
			} else if curr.End == last.End {
				if last.States.Equal(curr.States) {
					// Case curr.Start < last.End && curr.End == last.End && last.States == curr.States:
					//
					//   last:  |_____|___________|  States: {1,2}    ---->    |_________________|  States: {1,2}
					//   curr:        |___________|  States: {1,2}    ---->
					//
					// Case curr.Start == last.End && curr.End == last.End && last.States == curr.States:
					//
					//   last:  |_________________|  States: {1,2}    ---->    |_________________|  States: {1,2}
					//   curr:                    |  States: {1,2}    ---->
					//
				} else {
					// Case curr.Start < last.End && curr.End == last.End && last.States != curr.States:
					//
					//   last:  |_____|___________|  States: {1,2}    ---->    |____||           |  States: {1,2}
					//   curr:        |___________|  States: {3,4}    ---->          |___________|  States: {3,4}
					//
					// Case curr.Start == last.End && curr.End == last.End && last.States != curr.States:
					//
					//   last:  |_________________|  States: {1,2}    ---->    |________________||  States: {1,2}
					//   curr:                    |  States: {3,4}    ---->                      |  States: {3,4}
					//

					last.End = curr.Start - 1
					merged = append(merged, curr)
				}
			} else /* if curr.End > last.End */ {
				if last.States.Equal(curr.States) {
					// Case curr.Start < last.End && curr.End > last.End && last.States == curr.States:
					//
					//   last:  |_____|_____|     |  States: {1,2}    ---->    |_________________|  States: {1,2}
					//   curr:        |___________|  States: {1,2}    ---->
					//
					// Case curr.Start == last.End && curr.End > last.End && last.States == curr.States:
					//
					//   last:  |___________|     |  States: {1,2}    ---->    |_________________|  States: {1,2}
					//   curr:              |_____|  States: {1,2}    ---->
					//

					last.End = curr.End
				} else {
					// Case curr.Start < last.End && curr.End > last.End && last.States != curr.States:
					//
					//   last:  |_____|_____|     |  States: {1,2}    ---->    |____||           |  States: {1,2}
					//   curr:        |___________|  States: {3,4}    ---->          |___________|  States: {3,4}
					//
					// Case curr.Start == last.End && curr.End > last.End && last.States != curr.States:
					//
					//   last:  |___________|     |  States: {1,2}    ---->    |__________||     |  States: {1,2}
					//   curr:              |_____|  States: {3,4}    ---->                |_____|  States: {3,4}
					//

					last.End = curr.Start - 1
					merged = append(merged, curr)
				}
			}
		} else if curr.Start == last.End+1 && last.States.Equal(curr.States) {
			// Case curr.Start is adjacent to last.End && last.States != curr.States:
			//
			//   last:  |__________||     |  States: {1,2}    ---->    |_________________|  States: {1,2}
			//   curr:              |_____|  States: {1,2}    ---->
			//

			last.End = curr.End
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
