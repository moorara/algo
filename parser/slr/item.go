package slr

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

var (
	eqItem = func(lhs, rhs Item) bool {
		return lhs.Equals(rhs)
	}

	eqItemSet = func(lhs, rhs ItemSet) bool {
		return lhs.Equals(rhs)
	}
)

// ItemSet represents a set of items for a context-free grammar.
type ItemSet set.Set[Item]

// NewItemSet creates a new set of items.
func NewItemSet(items ...Item) ItemSet {
	return set.New(eqItem, items...)
}

// ItemSet represents a collection of item sets for a context-free grammar.
type ItemSetCollection set.Set[ItemSet]

// NewItemSetCollection creates a new collection of item sets.
func NewItemSetCollection(sets ...ItemSet) ItemSetCollection {
	return set.New(eqItemSet, sets...)
}

// Item is a production with a dot at some position of the body.
// For example, production A → XYZ yields the four items:
//
//	A → •XYZ
//	A → X•YZ
//	A → XY•Z
//	A → XYZ•
//
// The production A → ε generates only one item, A → •.
//
// An item tracks the progress of a production during parsing.
// For instance, A → X•YZ indicates that a string derivable from X has been seen,
// and the parser expects a string derivable from YZ next.
type Item struct {
	*grammar.Production

	// The position of dot in the body of a production.
	Dot int
}

// String returns a string representation of an item.
func (i Item) String() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "%s → ", i.Head)

	if α := i.Body[:i.Dot]; len(α) > 0 {
		b.WriteString(α.String())
	}

	fmt.Fprintf(b, "•")

	if β := i.Body[i.Dot:]; len(β) > 0 {
		b.WriteString(β.String())
	}

	return b.String()
}

// Equals determines whether or not two items are the same.
func (i Item) Equals(rhs Item) bool {
	return i.Production.Equals(*rhs.Production) && i.Dot == rhs.Dot
}

// IsKernel determines whether an item is a kernel item.
//
// Kernel items include:
//
//   - The initial item S′ → •S.
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are items where the dot is at the beginning, except for the initial item.
func (i Item) IsKernel(initial Item) bool {
	return i.Equals(initial) || i.Dot > 0
}

// IsComplete checks if the dot has reached the end of the body of the production.
func (i Item) IsComplete() bool {
	return i.Dot == len(i.Body)
}

// DotSymbol returns the grammar symbol at the dot position in the item's body.
// For example, in the item A → α•Bβ, it returns B.
// If the dot is at the end of the body, the function returns nil and false.
func (i Item) DotSymbol() (grammar.Symbol, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return i.Body[i.Dot], true
}

// NextItem generates a new item by advancing the dot one position forward in the item's body.
// For example, for the item A → α•Bβ, it returns A → αB•β.
// If the dot is at the end of the body, the function returns an empty item and false.
func (i Item) NextItem() (Item, bool) {
	if i.IsComplete() {
		return Item{}, false
	}

	return Item{
		Production: i.Production,
		Dot:        i.Dot + 1,
	}, true
}
