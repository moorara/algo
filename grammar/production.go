package grammar

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"iter"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

var (
	CmpProduction  = cmpProduction
	HashProduction = hashFuncForProduction()

	EqProduction = func(lhs, rhs *Production) bool {
		return lhs.Equal(rhs)
	}

	EqProductionSet = func(lhs, rhs set.Set[*Production]) bool {
		return lhs.Equal(rhs)
	}
)

// Production represents a context-free production rule.
// The productions of a context-free grammar determine how the terminals and non-terminals can be combined to form strings.
type Production struct {
	// Head or left side defines some of the strings denoted by the non-terminal symbol.
	Head NonTerminal
	// Body or right side describes one way in which strings of the non-terminal at the head can be constructed.
	Body String[Symbol]
}

// String returns a string representation of a production rule.
func (p *Production) String() string {
	return fmt.Sprintf("%s → %s", p.Head, p.Body)
}

// Equal determines whether or not two production rules are the same.
func (p *Production) Equal(rhs *Production) bool {
	return p.Head.Equal(rhs.Head) && p.Body.Equal(rhs.Body)
}

// IsEmpty determines whether or not a production rule is an empty production (ε-production).
//
// An empty production (ε-production) is any production of the form A → ε.
func (p *Production) IsEmpty() bool {
	return len(p.Body) == 0
}

// IsSingle determines whether or not a production rule is a single production (unit production).
//
// A single production (unit production) is a production whose body is a single non-terminal (A → B).
func (p *Production) IsSingle() bool {
	return len(p.Body) == 1 && !p.Body[0].IsTerminal()
}

// IsLeftRecursive determines whether or not a production rule is left recursive (immediate left recursive).
//
// A left recursive production is a production rule of the form of A → Aα
func (p *Production) IsLeftRecursive() bool {
	return len(p.Body) > 0 && p.Body[0].Equal(p.Head)
}

// IsCNF checks if a production rule is in Chomsky Normal Form (CNF).
//
// A production is in CNF if it is either:
//
//  1. A → BC: where A, B, and C are non-terminal symbols.
//  2. A → a: where A is a non-terminal symbol and a is a terminal symbol.
//
// The function returns two boolean values:
//
//   - The first value indicates if the rule is of the form A → BC.
//   - The second value indicates if the rule is of the form A → a.
func (p *Production) IsCNF() (bool, bool) {
	return len(p.Body) == 2 && !p.Body[0].IsTerminal() && !p.Body[1].IsTerminal(),
		len(p.Body) == 1 && p.Body[0].IsTerminal()
}

// Productions represents a set of production rules for a context-free grammar.
type Productions struct {
	table symboltable.SymbolTable[NonTerminal, set.Set[*Production]]
}

// NewProductions creates a new instance of the Productions.
func NewProductions() *Productions {
	return &Productions{
		table: symboltable.NewQuadraticHashTable(
			HashNonTerminal,
			EqNonTerminal,
			EqProductionSet,
			symboltable.HashOpts{},
		),
	}
}

// String returns a string representation of production rules.
func (p *Productions) String() string {
	var b bytes.Buffer

	for head, prods := range p.table.All() {
		fmt.Fprintf(&b, "%s → ", head)

		for _, q := range OrderProductionSet(prods) {
			fmt.Fprintf(&b, "%s | ", q.Body.String())
		}

		// Remove the last " | "
		if b.Len() >= 3 {
			b.Truncate(b.Len() - 3)
		}

		fmt.Fprintln(&b)
	}

	return b.String()
}

// Equal determines whether or not two sets of production rules are the same.
func (p *Productions) Equal(rhs *Productions) bool {
	return p.table.Equal(rhs.table)
}

// Clone returns a deep copy of the production rules, ensuring the clone is independent of the original.
func (p *Productions) Clone() *Productions {
	newP := NewProductions()
	for q := range p.All() {
		newP.Add(q)
	}

	return newP
}

// Size returns the number of production rules.
func (p *Productions) Size() int {
	size := 0
	for _, prods := range p.table.All() {
		size += prods.Size()
	}

	return size
}

// Add adds a new production rule.
func (p *Productions) Add(ps ...*Production) {
	for _, q := range ps {
		if _, ok := p.table.Get(q.Head); !ok {
			p.table.Put(q.Head, set.New(EqProduction))
		}

		list, _ := p.table.Get(q.Head)
		list.Add(q)
	}
}

// Remove removes a production rule.
func (p *Productions) Remove(ps ...*Production) {
	for _, q := range ps {
		if list, ok := p.table.Get(q.Head); ok {
			list.Remove(q)
			if list.IsEmpty() {
				p.table.Delete(q.Head)
			}
		}
	}
}

// RemoveAll removes all production rules with the specified head non-terminal.
func (p *Productions) RemoveAll(heads ...NonTerminal) {
	for _, head := range heads {
		p.table.Delete(head)
	}
}

// Get finds and returns a production rule by its head non-terminal symbol.
// It returns nil if no production rules are found for the specified head.
func (p *Productions) Get(head NonTerminal) set.Set[*Production] {
	list, ok := p.table.Get(head)
	if !ok {
		return nil
	}

	return list
}

// All returns an iterator sequence containing all production rules.
func (p *Productions) All() iter.Seq[*Production] {
	return func(yield func(*Production) bool) {
		for _, prods := range p.table.All() {
			for prod := range prods.All() {
				if !yield(prod) {
					return
				}
			}
		}
	}
}

// AllByHead returns an iterator sequence sequence of pairs,
// where each pair consists of a head non-terminal and its associated set of production rules.
func (p *Productions) AllByHead() iter.Seq2[NonTerminal, set.Set[*Production]] {
	return p.table.All()
}

// AnyMatch returns true if at least one production rule satisfies the provided predicate.
func (p *Productions) AnyMatch(pred generic.Predicate1[*Production]) bool {
	for q := range p.All() {
		if pred(q) {
			return true
		}
	}

	return false
}

// AllMatch returns true if all production rules satisfy the provided predicate.
// If the set of production rules is empty, it returns true.
func (p *Productions) AllMatch(pred generic.Predicate1[*Production]) bool {
	for q := range p.All() {
		if !pred(q) {
			return false
		}
	}

	return true
}

// SelectMatch selects a subset of production rules that satisfy the given predicate.
// It returns a new set of production rules containing the matching productions, of the same type as the original set of production rules.
func (p *Productions) SelectMatch(pred generic.Predicate1[*Production]) *Productions {
	newP := NewProductions()

	for q := range p.All() {
		if pred(q) {
			newP.Add(q)
		}
	}

	return newP
}

// OrderProductionSet orders an unordered set of production rules in a deterministic way.
func OrderProductionSet(set set.Set[*Production]) []*Production {
	prods := generic.Collect1(set.All())
	orderProductionSlice(prods)
	return prods
}

// orderProductionSlice orders a slice of production rules in a deterministic way.
func orderProductionSlice(prods []*Production) {
	// Sort the productions using a custom comparison function.
	sort.Quick(prods, cmpProduction)
}

// cmpProduction is a CompareFunc for Production type.
//
// The comparing criteria are as follows:
//
//  1. Production heads are compared alphabetically.
//  2. If two productions have the same heads, productions bodies are compared.
//
// This function can be used for sorting productions
// to ensure a consistent and deterministic order for any given set of production rules.
func cmpProduction(lhs, rhs *Production) int {
	// First, compare the heads of productions.
	if cmp := CmpNonTerminal(lhs.Head, rhs.Head); cmp < 0 {
		return -1
	} else if cmp > 0 {
		return 1
	}

	// Next, compare production bodies.
	return CmpString(lhs.Body, rhs.Body)
}

// hashFuncForProduction creates a HashFunc for hashing productions.
func hashFuncForProduction() hash.HashFunc[*Production] {
	h := fnv.New64()

	return func(p *Production) uint64 {
		h.Reset()
		_, _ = WriteSymbol(h, p.Head) // Hash.Write never returns an error
		_, _ = WriteString(h, p.Body) // Hash.Write never returns an error
		return h.Sum64()
	}
}
