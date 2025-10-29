package automata

import (
	"bytes"
	"fmt"
	"iter"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/range/disc"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/symboltable"
)

var (
	eqClassID   = generic.NewEqualFunc[classID]()
	cmpClassID  = generic.NewCompareFunc[classID]()
	hashClassID = hash.HashFuncForInt[classID](nil)
)

// rangeSet represents a set of symbol ranges.
type rangeSet set.Set[disc.Range[Symbol]]

// newRangeSet creates a new set of symbol ranges.
func newRangeSet(rs ...disc.Range[Symbol]) rangeSet {
	return set.NewStableWithFormat(
		func(a, b disc.Range[Symbol]) bool {
			return a.Lo == b.Lo && a.Hi == b.Hi
		},
		func(all []disc.Range[Symbol]) string {
			vals := make([]string, len(all))
			for i, r := range all {
				vals[i] = fmtRange(r)
			}

			return strings.Join(vals, ", ")
		},
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
					vals = append(vals, fmtRange(r))
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
				fmt.Fprintf(&b, "  %s: %d\n", fmtRange(r), cid)
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

// fmtRange formats a symbol range as a string.
func fmtRange(r disc.Range[Symbol]) string {
	var lo, hi Symbol

	if r.Lo == E {
		lo = 'ε'
	} else {
		lo = r.Lo
	}

	if r.Hi == E {
		hi = 'ε'
	} else {
		hi = r.Hi
	}

	return fmt.Sprintf("[%c..%c]", lo, hi)
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
