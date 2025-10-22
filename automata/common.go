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

// classID is used to identify equivalence classes of input symbols.
type classID int

// classMapping represents the equivalence classes of the input symbols.
// It is mapping from the classs ID to the set of ranges of symbols belonging to that class.
type classMapping symboltable.SymbolTable[classID, rangeSet]

func newClassMapping(pairs []generic.KeyValue[classID, rangeSet]) classMapping {
	tab := symboltable.NewQuadraticHashTable(
		hashClassID,
		eqClassID,
		func(a, b rangeSet) bool {
			if a == nil && b == nil {
				return true
			}

			if a == nil || b == nil {
				return false
			}

			return a.Equal(b)
		},
		symboltable.HashOpts{},
	)

	for _, p := range pairs {
		tab.Put(p.Key, p.Val)
	}

	return tab
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
