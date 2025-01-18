package slr

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
)

var (
	eqItem = func(lhs, rhs Item) bool {
		return lhs.Equals(rhs)
	}

	eqItemSet = func(lhs, rhs ItemSet) bool {
		return lhs.Equals(rhs)
	}
)

// ItemSet represents a collection of item sets for a context-free grammar.
type ItemSetCollection set.Set[ItemSet]

// NewItemSetCollection creates a new collection of item sets.
func NewItemSetCollection(sets ...ItemSet) ItemSetCollection {
	return set.New(eqItemSet, sets...)
}

// ItemSet represents a set of items for a context-free grammar.
type ItemSet set.Set[Item]

// NewItemSet creates a new set of items.
func NewItemSet(items ...Item) ItemSet {
	return set.New(eqItem, items...)
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
	Initial bool // True if this is the initial S′ → S production in the augmented grammar.
	Dot     int  // Position of the dot in the production body.
}

// String returns a string representation of an item.
func (i Item) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%s → ", i.Head)

	if α := i.Body[:i.Dot]; len(α) > 0 {
		b.WriteString(α.String())
	}

	b.WriteRune('•')

	if β := i.Body[i.Dot:]; len(β) > 0 {
		b.WriteString(β.String())
	}

	return b.String()
}

// Equals determines whether or not two items are the same.
func (i Item) Equals(rhs Item) bool {
	return i.Production.Equals(*rhs.Production) &&
		i.Initial == rhs.Initial &&
		i.Dot == rhs.Dot
}

// IsKernel determines whether an item is a kernel item.
//
// Kernel items include:
//
//   - The initial item S′ → •S.
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are items where the dot is at the beginning, except for the initial item.
func (i Item) IsKernel() bool {
	return i.Initial || i.Dot > 0
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
		Initial:    i.Initial,
		Dot:        i.Dot + 1,
	}, true
}

// cmpItem compares two items to establish a consistent and deterministic order.
// The comparison follows these rules:
//
//  1. Kernel items are prioritized over non-kernel items.
//  2. If both items are kernel or both are non-kernel, they are compared by their dot positions.
//  3. If the dot positions are identical, the items are compared based on their productions.
//
// This function is useful for sorting items, ensuring stability and predictability in their ordering.
func cmpItem(lhs, rhs Item) int {
	if lhs.IsKernel() && !rhs.IsKernel() {
		return -1
	} else if !lhs.IsKernel() && rhs.IsKernel() {
		return 1
	}

	if lhs.Dot < rhs.Dot {
		return -1
	} else if lhs.Dot > rhs.Dot {
		return 1
	}

	return grammar.CmpProduction(*lhs.Production, *rhs.Production)
}

// cmpItemSet compares two collections of item sets to establish a consistent and deterministic order.
// The comparison process is as follows:
//
//  1. Each item set is sorted using cmpItem first.
//  2. The items in the two lists are compared one by one in order.
//  3. If one list is shorter and all compared items are equal, the shorter list comes first.
//
// This function is useful for sorting collections of item sets,
// ensuring stability and predictability in their ordering.
func cmpItemSet(lhs, rhs ItemSet) int {
	ls := generic.Collect1(lhs.All())
	sort.Quick(ls, cmpItem)

	rs := generic.Collect1(rhs.All())
	sort.Quick(rs, cmpItem)

	for i := 0; i < len(ls) && i < len(rs); i++ {
		if cmp := cmpItem(ls[i], rs[i]); cmp != 0 {
			return cmp
		}
	}

	return len(ls) - len(rs)
}
