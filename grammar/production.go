package grammar

import (
	"bytes"
	"fmt"
	"iter"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	. "github.com/moorara/algo/symboltable"
)

var (
	eqProduction = func(lhs, rhs Production) bool {
		return lhs.Equals(rhs)
	}

	eqProductionSet = func(lhs, rhs set.Set[Production]) bool {
		return lhs.Equals(rhs)
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
func (p Production) String() string {
	return fmt.Sprintf("%s → %s", p.Head, p.Body)
}

// Equals determines whether or not two production rules are the same.
func (p Production) Equals(rhs Production) bool {
	return p.Head.Equals(rhs.Head) && p.Body.Equals(rhs.Body)
}

// IsEmpty determines whether or not a production rule is an empty production (ε-production).
//
// An empty production (ε-production) is any production of the form A → ε.
func (p Production) IsEmpty() bool {
	return len(p.Body) == 0
}

// IsSingle determines whether or not a production rule is a single production (unit production).
//
// A single production (unit production) is a production whose body is a single non-terminal (A → B).
func (p Production) IsSingle() bool {
	return len(p.Body) == 1 && !p.Body[0].IsTerminal()
}

// IsLeftRecursive determines whether or not a production rule is left recursive (immediate left recursive).
//
// A left recursive production is a production rule of the form of A → Aα
func (p Production) IsLeftRecursive() bool {
	return len(p.Body) > 0 && p.Body[0].Equals(p.Head)
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
func (p Production) IsCNF() (bool, bool) {
	return len(p.Body) == 2 && !p.Body[0].IsTerminal() && !p.Body[1].IsTerminal(),
		len(p.Body) == 1 && p.Body[0].IsTerminal()
}

// Productions is the interface for the set of production rules of a context-free grammar.
type Productions interface {
	fmt.Stringer
	Cloner[Productions]
	Equaler[Productions]

	Add(...Production)
	Remove(...Production)
	RemoveAll(...NonTerminal)
	Get(NonTerminal) set.Set[Production]
	Order(NonTerminal) []Production
	All() iter.Seq[Production]
	AllByHead() iter.Seq2[NonTerminal, set.Set[Production]]
	AnyMatch(Predicate1[Production]) bool
	AllMatch(Predicate1[Production]) bool
}

// productions implements the Productions interface.
type productions struct {
	table SymbolTable[NonTerminal, set.Set[Production]]
}

// NewProductions creates a new instance of the Productions.
func NewProductions() Productions {
	return &productions{
		table: NewQuadraticHashTable(hashNonTerminal, eqNonTerminal, eqProductionSet, HashOpts{}),
	}
}

// String returns a string representation of production rules.
func (p *productions) String() string {
	var b bytes.Buffer

	for head := range p.table.All() {
		fmt.Fprintf(&b, "%s → ", head)
		for _, q := range p.Order(head) {
			fmt.Fprintf(&b, "%s | ", q.Body.String())
		}
		b.Truncate(b.Len() - 3)
		fmt.Fprintln(&b)
	}

	return b.String()
}

// Clone returns a deep copy of the production rules, ensuring the clone is independent of the original.
func (p *productions) Clone() Productions {
	newP := NewProductions()
	for q := range p.All() {
		newP.Add(q)
	}

	return newP
}

// Equals determines whether or not two sets of production rules are the same.
func (p *productions) Equals(rhs Productions) bool {
	q, ok := rhs.(*productions)
	return ok && p.table.Equals(q.table)
}

// Add adds a new production rule.
func (p *productions) Add(ps ...Production) {
	for _, q := range ps {
		if _, ok := p.table.Get(q.Head); !ok {
			p.table.Put(q.Head, set.New(eqProduction))
		}

		list, _ := p.table.Get(q.Head)
		list.Add(q)
	}
}

// Remove removes a production rule.
func (p *productions) Remove(ps ...Production) {
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
func (p *productions) RemoveAll(heads ...NonTerminal) {
	for _, head := range heads {
		p.table.Delete(head)
	}
}

// Get finds and returns a production rule by its head non-terminal symbol.
// It returns nil if no production rules are found for the specified head.
func (p *productions) Get(head NonTerminal) set.Set[Production] {
	list, ok := p.table.Get(head)
	if !ok {
		return nil
	}

	return list
}

// Order orders an unordered set of production rules with the same head non-terminal in a deterministic way.
//
// The ordering criteria are as follows:
//
//  1. Productions whose bodies contain more non-terminal symbols are prioritized first.
//  2. If two productions have the same number of non-terminals, those with more terminal symbols in the body come first.
//  3. If two productions have the same number of non-terminals and terminals, they are ordered alphabetically based on the symbols in their bodies.
//
// The goal of this function is to ensure a consistent and deterministic order for any given set of production rules.
func (p *productions) Order(head NonTerminal) []Production {
	// Collect all production rules into a slice from the set iterator.
	prods := Collect1(p.Get(head).All())

	// Sort the productions using a custom comparison function.
	sort.Quick[Production](prods, func(lhs, rhs Production) int {
		// First, compare based on the number of non-terminal symbols in the body.
		lhsNonTermsLen, rhsNonTermsLen := len(lhs.Body.NonTerminals()), len(rhs.Body.NonTerminals())
		if lhsNonTermsLen > rhsNonTermsLen {
			return -1
		} else if rhsNonTermsLen > lhsNonTermsLen {
			return 1
		}

		// Next, if the number of non-terminals is the same,
		//   compare based on the number of terminal symbols.
		lhsTermsLen, rhsTermsLen := len(lhs.Body.Terminals()), len(rhs.Body.Terminals())
		if lhsTermsLen > rhsTermsLen {
			return -1
		} else if rhsTermsLen > lhsTermsLen {
			return 1
		}

		// Then, if the number of terminals is also the same,
		//   compare alphabetically based on the string representation of the bodies.
		lhsString, rhsString := lhs.String(), rhs.String()
		if lhsString < rhsString {
			return -1
		} else if rhsString < lhsString {
			return 1
		}

		return 0
	})

	return prods
}

// All returns an iterator sequence containing all production rules.
func (p *productions) All() iter.Seq[Production] {
	return func(yield func(Production) bool) {
		for _, list := range p.table.All() {
			for q := range list.All() {
				if !yield(q) {
					return
				}
			}
		}
	}
}

// AllByHead returns an iterator sequence sequence of pairs,
// where each pair consists of a head non-terminal and its associated set of production rules.
func (p *productions) AllByHead() iter.Seq2[NonTerminal, set.Set[Production]] {
	return p.table.All()
}

// AnyMatch returns true if at least one production rule satisfies the provided predicate.
func (p *productions) AnyMatch(pred Predicate1[Production]) bool {
	for q := range p.All() {
		if pred(q) {
			return true
		}
	}

	return false
}

// AllMatch returns true if all production rules satisfy the provided predicate.
// If the set of production rules is empty, it returns true.
func (p *productions) AllMatch(pred Predicate1[Production]) bool {
	for q := range p.All() {
		if !pred(q) {
			return false
		}
	}

	return true
}

// SelectMatch selects a subset of production rules that satisfy the given predicate.
// It returns a new set of production rules containing the matching productions, of the same type as the original set of production rules.
func (p *productions) SelectMatch(pred Predicate1[Production]) Productions {
	newP := NewProductions()

	for q := range p.All() {
		if pred(q) {
			newP.Add(q)
		}
	}

	return newP
}
