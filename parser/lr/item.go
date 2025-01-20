package lr

import (
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

	cmpItem = func(lhs, rhs Item) int {
		return lhs.Compare(rhs)
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

// cmpItemSet compares two collections of item sets to establish a consistent and deterministic order.
// The comparison process is as follows:
//
//  1. Each item set is sorted using cmpItem first.
//  2. The items in the two lists are compared one by one in order using cmpItem.
//  3. If one list is longer and all compared items are equal, the longer list comes first.
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

	//  If all compared items are equal, the longer list comes first.
	return len(rs) - len(ls)
}

// Item represents an LR item for a context-free grammar.
//
// This interface defines the methods required for a general LR parser.
type Item interface {
	fmt.Stringer
	generic.Equaler[Item]

	// Compare compares two items to establish a consistent and deterministic order.
	// This function is useful for sorting items, ensuring stability and predictability in their ordering.
	Compare(Item) int

	// IsInitial checks if an item is the initial item in the augmented grammar.
	// For LR(0), the initial item is "S′ → •S",
	// and for LR(1), the initial item is "S′ → •S, $".
	IsInitial() bool

	// IsKernel determines whether an item is a kernel item.
	//
	// Kernel items include:
	//
	//   - The initial item.
	//   - All items where the dot is not at the beginning (left end) of the item's body.
	//
	// Non-kernel items are those where the dot is at the beginning, except for the initial item.
	IsKernel() bool

	// IsComplete checks if the dot has reached the end of the item's body.
	IsComplete() bool

	// IsFinal checks if an item is the final item in the augmented grammar.
	// For LR(0), the final item is "S′ → S•",
	// and for LR(1), the initial item is "S′ → S•, $".
	IsFinal() bool

	// Dot returns the grammar symbol at the dot position in the item's body.
	// If the dot is at the end of the body, it returns nil and false.
	DotSymbol() (grammar.Symbol, bool)

	// NextItem generates a new item by advancing the dot one position forward in the item's body.
	// If the dot is at the end of the body, it returns an empty item and false.
	Next() (Item, bool)
}
