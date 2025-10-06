package automata

import (
	"bytes"
	"fmt"
	"iter"
	"slices"

	"github.com/moorara/algo/symboltable"
)

// rangeState represents a pair of a symbol range and a single state.
// It is used in DFA transitions to group transitions that share the same range of input symbols.
type rangeState struct {
	SymbolRange
	State
}

// String implements the fmt.Stringer interface.
func (rs rangeState) String() string {
	return fmt.Sprintf("%s → %d", rs.SymbolRange, rs.State)
}

// Equal implements the generic.Equaler interface.
func (rs rangeState) Equal(rhs rangeState) bool {
	return rs.SymbolRange.Equal(rhs.SymbolRange) && rs.State == rhs.State
}

// dfaTransitionTable implements a transition table for deterministic finite automata (DFA).
// It is used in DFA to manage transitions from one state to another over ranges of input symbols.
type dfaTransitionTable struct {
	table symboltable.SymbolTable[State, []rangeState]
}

// newDFATransitionTable creates a new instance of dfaTransitionTable.
func newDFATransitionTable(trans map[State][]rangeState) *dfaTransitionTable {
	table := symboltable.NewAVL[State, []rangeState](CmpState, nil)

	for s, pairs := range trans {
		for _, pair := range pairs {
			pair.Validate()
		}

		slices.SortFunc(pairs, func(lhs, rhs rangeState) int {
			return int(lhs.Start) - int(rhs.Start)
		})

		table.Put(s, mergeRangeStateSortedList(pairs))
	}

	return &dfaTransitionTable{
		table: table,
	}
}

// String implements the fmt.Stringer interface.
func (t *dfaTransitionTable) String() string {
	var b bytes.Buffer

	b.WriteString("Transitions:\n")

	for s, pairs := range t.table.All() {
		for _, pair := range pairs {
			fmt.Fprintf(&b, "  %d --%s--> %d\n", s, pair.SymbolRange, pair.State)
		}
	}

	return b.String()
}

// Clone implements the generic.Cloner interface.
func (t *dfaTransitionTable) Clone() *dfaTransitionTable {
	tt := &dfaTransitionTable{
		table: symboltable.NewAVL[State, []rangeState](CmpState, nil),
	}

	for s, pairs := range t.table.All() {
		pp := make([]rangeState, len(pairs))
		for i, pair := range pairs {
			pp[i] = rangeState{
				SymbolRange: pair.SymbolRange,
				State:       pair.State,
			}
		}
		tt.table.Put(s, pp)
	}

	return tt
}

// Equal implements the generic.Equaler interface.
func (t *dfaTransitionTable) Equal(rhs *dfaTransitionTable) bool {
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
func (t *dfaTransitionTable) All() iter.Seq2[State, iter.Seq2[SymbolRange, State]] {
	return func(yield func(State, iter.Seq2[SymbolRange, State]) bool) {
		for s := range t.table.All() {
			if !yield(s, t.From(s)) {
				return
			}
		}
	}
}

// From returns all transitions from the given state in the table.
func (t *dfaTransitionTable) From(s State) iter.Seq2[SymbolRange, State] {
	return func(yield func(SymbolRange, State) bool) {
		if pairs, ok := t.table.Get(s); ok {
			for _, pair := range pairs {
				if !yield(pair.SymbolRange, pair.State) {
					return
				}
			}
		}
	}
}

// Next returns the next state for the given state and input symbol.
func (t *dfaTransitionTable) Next(s State, a Symbol) (State, bool) {
	if pairs, ok := t.table.Get(s); ok {
		if i, ok := searchRangeStateSortedList(pairs, a); ok {
			return pairs[i].State, true
		}
	}

	return -1, false
}

// Add inserts a new transition to the DFA transition table.
// It will merge any overlapping or adjacent ranges as necessary.
// The state associated with any overlapping range will be overridden by the new state given in the new range.
func (t *dfaTransitionTable) Add(s State, start, end Symbol, next State) {
	new := rangeState{
		SymbolRange{Start: start, End: end},
		next,
	}

	new.Validate()

	pairs, ok := t.table.Get(s)
	if !ok {
		pairs = make([]rangeState, 0, 1)
	}

	// Find the insertion point
	i, ok := searchRangeStateSortedList(pairs, new.Start)
	if ok {
		i++
	}

	// Insert the new entry at position i
	pairs = append(pairs, rangeState{})
	copy(pairs[i+1:], pairs[i:])
	pairs[i] = new

	// Merge overlapping or adjacent ranges
	t.table.Put(s, mergeRangeStateSortedList(pairs))
}

// searchRangeStateSortedList performs a binary search to find the index of the range that contains the given symbol.
// If found, it returns the index and true; otherwise, it returns the insertion point and false.
func searchRangeStateSortedList(pairs []rangeState, a Symbol) (int, bool) {
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

// mergeRangeStateSortedList merges overlapping or adjacent ranges in a sorted list of range-state pairs.
func mergeRangeStateSortedList(pairs []rangeState) []rangeState {
	merged := make([]rangeState, 0, len(pairs))

	for _, curr := range pairs {
		if len(merged) == 0 {
			merged = append(merged, curr)
			continue
		}

		last := &merged[len(merged)-1]

		if curr.Start <= last.End {
			if curr.End < last.End {
				if last.State == curr.State {
					// Case curr.Start < last.End && curr.End < last.End && last.State == curr.State:
					//
					//   last:  |_____|_____|_____|  State: 1    ---->    |_________________|  State: 1
					//   curr:        |_____|        State: 1    ---->
					//
					// Impossible case of curr.Start == last.End && curr.End < last.End
					//
				} else {
					// Case curr.Start < last.End && curr.End < last.End && last.State != curr.State:
					//
					//   last:  |_____|_____|_____|  State: 1    ---->    |____||     ||    |  State: 1
					//   curr:        |_____|        State: 2    ---->          |_____||    |  State: 2
					//                                           ---->                 |____|  State: 1
					//
					// Impossible case of curr.Start == last.End && curr.End < last.End
					//

					lastEnd := last.End
					last.End = curr.Start - 1
					merged = append(merged, curr)
					merged = append(merged, rangeState{
						SymbolRange{Start: curr.End + 1, End: lastEnd},
						last.State,
					})
				}
			} else if curr.End == last.End {
				if last.State == curr.State {
					// Case curr.Start < last.End && curr.End == last.End && last.State == curr.State:
					//
					//   last:  |_____|___________|  State: 1    ---->    |_________________|  State: 1
					//   curr:        |___________|  State: 1    ---->
					//
					// Case curr.Start == last.End && curr.End == last.End && last.State == curr.State:
					//
					//   last:  |_________________|  State: 1    ---->    |_________________|  State: 1
					//   curr:                    |  State: 1    ---->
					//
				} else {
					// Case curr.Start < last.End && curr.End == last.End && last.State != curr.State:
					//
					//   last:  |_____|___________|  State: 1    ---->    |____||           |  State: 1
					//   curr:        |___________|  State: 2    ---->          |___________|  State: 2
					//
					// Case curr.Start == last.End && curr.End == last.End && last.State != curr.State:
					//
					//   last:  |_________________|  State: 1    ---->    |________________||  State: 1
					//   curr:                    |  State: 2    ---->                      |  State: 2
					//

					last.End = curr.Start - 1
					merged = append(merged, curr)
				}
			} else /* if curr.End > last.End */ {
				if last.State == curr.State {
					// Case curr.Start < last.End && curr.End > last.End && last.State == curr.State:
					//
					//   last:  |_____|_____|     |  State: 1    ---->    |_________________|  State: 1
					//   curr:        |___________|  State: 1    ---->
					//
					// Case curr.Start == last.End && curr.End > last.End && last.State == curr.State:
					//
					//   last:  |___________|     |  State: 1    ---->    |_________________|  State: 1
					//   curr:              |_____|  State: 1    ---->
					//

					last.End = curr.End
				} else {
					// Case curr.Start < last.End && curr.End > last.End && last.State != curr.State:
					//
					//   last:  |_____|_____|     |  State: 1    ---->    |____||           |  State: 1
					//   curr:        |___________|  State: 2    ---->          |___________|  State: 2
					//
					// Case curr.Start == last.End && curr.End > last.End && last.State != curr.State:
					//
					//   last:  |___________|     |  State: 1    ---->    |__________||     |  State: 1
					//   curr:              |_____|  State: 2    ---->                |_____|  State: 2
					//

					last.End = curr.Start - 1
					merged = append(merged, curr)
				}
			}
		} else if curr.Start == last.End+1 && last.State == curr.State {
			// Case curr.Start is adjacent to last.End && last.State != curr.State:
			//
			//   last:  |__________||     |  State: 1    ---->    |_________________|  State: 1
			//   curr:              |_____|  State: 1    ---->
			//

			last.End = curr.End
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
