package automata

import (
	"bytes"
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

func formatRangeSlice(rs []disc.Range[Symbol]) string {
	var b bytes.Buffer

	for _, r := range rs {
		fmt.Fprintf(&b, "%s, ", formatRange(r))
	}

	if b.Len() >= 2 {
		b.Truncate(b.Len() - 2)
	}

	return b.String()
}

func formatRange(r disc.Range[Symbol]) string {
	return fmt.Sprintf("[%s..%s]", formatRangeBound(r.Lo), formatRangeBound(r.Hi))
}

func formatRangeBound(a Symbol) string {
	switch a {
	case E:
		return "ε"
	case 0:
		return "NUL"
	case '\t':
		return "\\t"
	case '\n':
		return "\\n"
	case '\v':
		return "\\v"
	case '\f':
		return "\\f"
	case '\r':
		return "\\r"
	case ' ':
		return "SP"
	default:
		return fmt.Sprintf("%c", a)
	}
}

// rangeSet represents a set of symbol ranges.
type rangeSet set.Set[disc.Range[Symbol]]

// newRangeSet creates a new set of symbol ranges.
func newRangeSet(rs ...disc.Range[Symbol]) rangeSet {
	return set.NewStableSetWithFormat(
		func(a, b disc.Range[Symbol]) bool {
			return a.Lo == b.Lo && a.Hi == b.Hi
		},
		formatRangeSlice,
		rs...,
	)
}

// rangeList represents a list of symbol ranges.
type rangeList disc.RangeList[Symbol]

// newRangeList creates a new list of symbol ranges.
func newRangeList(rs ...disc.Range[Symbol]) rangeList {
	return disc.NewRangeList(
		&disc.RangeListOpts[Symbol]{
			Format: func(all iter.Seq[disc.Range[Symbol]]) string {
				vals := make([]string, 0)
				for r := range all {
					vals = append(vals, formatRange(r))
				}
				return strings.Join(vals, ", ")
			},
		},
		rs...,
	)
}

// rangeMapping represents the equivalence classes of the input symbols.
// It is mapping from the a range of input symbols to the class ID they belong to.
type rangeMapping disc.RangeMap[Symbol, classID]

func newRangeMapping(pairs []disc.RangeValue[Symbol, classID]) rangeMapping {
	opts := &disc.RangeMapOpts[Symbol, classID]{
		Format: func(all iter.Seq2[disc.Range[Symbol], classID]) string {
			var b bytes.Buffer

			b.WriteString("Ranges:\n")
			for r, cid := range all {
				fmt.Fprintf(&b, "  %s: %d\n", formatRange(r), cid)
			}

			return b.String()
		},
	}

	return disc.NewRangeMap(eqClassID, opts, pairs...)
}

// classMapping represents the equivalence classes of the input symbols.
// It is mapping from the classs ID to the set of ranges of symbols belonging to that class.
type classMapping symboltable.SymbolTable[classID, rangeSet]

func newClassMapping(pairs []generic.KeyValue[classID, rangeSet]) classMapping {
	// Use an ordered symbol table so iterations over classes are deterministic and
	// the resulting computation and textual output are reproducible.
	tab := symboltable.NewRedBlack(
		cmpClassID,
		func(a, b rangeSet) bool {
			if a == nil && b == nil {
				return true
			}

			if a == nil || b == nil {
				return false
			}

			return a.Equal(b)
		},
	)

	for _, p := range pairs {
		tab.Put(p.Key, p.Val)
	}

	return tab
}

// stateManager is used for keeping track of states when combining multiple automata.
type stateManager struct {
	last   State
	states map[int]map[State]State
}

func newStateManager(last State) *stateManager {
	return &stateManager{
		last:   last,
		states: map[int]map[State]State{},
	}
}

func (m *stateManager) GetOrCreateState(id int, s State) State {
	if _, ok := m.states[id]; !ok {
		m.states[id] = make(map[State]State)
	}

	if t, ok := m.states[id][s]; ok {
		return t
	}

	m.last++
	m.states[id][s] = m.last
	return m.last
}

// generateStatePermutations generates all permutations of a sequence of states using recursion and backtracking.
// Each permutation is passed to the provided yield function.
func generateStatePermutations(states []State, start, end int, yield func([]State) bool) bool {
	if start == end {
		return yield(states)
	}

	for i := start; i <= end; i++ {
		states[start], states[i] = states[i], states[start]
		cont := generateStatePermutations(states, start+1, end, yield)
		states[start], states[i] = states[i], states[start]

		if !cont {
			return false
		}
	}

	return true
}
