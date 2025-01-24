package lr

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
		return lhs.Equal(rhs)
	}

	cmpItem = func(lhs, rhs Item) int {
		return lhs.Compare(rhs)
	}

	eqItemSet = func(lhs, rhs ItemSet) bool {
		return lhs.Equal(rhs)
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
// This interface defines the methods required by an LR parser.
type Item interface {
	fmt.Stringer
	generic.Equaler[Item]
	generic.Comparer[Item]

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

// Item0 represents an LR(0) item for a context-free grammar.
// It implements the Item interface.
//
// An LR(0) item is a production with a dot at some position of the body.
// For example, production A → XYZ yields the four LR(0) items:
//
//	A → •XYZ
//	A → X•YZ
//	A → XY•Z
//	A → XYZ•
//
// The production A → ε generates only one LR(0) item, A → •.
//
// Intuitively, an LR(0) item tracks the progress of a production during parsing.
// For instance, A → X•YZ indicates that a string derivable from X has been seen,
// and the parser expects a string derivable from YZ next.
type Item0 struct {
	*grammar.Production
	Start grammar.NonTerminal // The start symbol S′ in the augmented grammar.
	Dot   int                 // Position of the dot in the production body.
}

// String returns a string representation of an LR(0) item.
func (i *Item0) String() string {
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

// Equal determines whether or not two LR(0) items are the same.
func (i *Item0) Equal(rhs Item) bool {
	ii, ok := rhs.(*Item0)
	return ok &&
		i.Production.Equal(ii.Production) &&
		i.Start.Equal(ii.Start) &&
		i.Dot == ii.Dot
}

// Compare compares two LR(0) items to establish a consistent and deterministic order.
// The comparison follows these rules:
//
//  1. Initial items precede non-initial items.
//  2. Kernel items precede non-kernel items.
//  3. If both items are kernel or non-kernel,
//     items with production heads equal to the augmented start symbol S′ come first.
//  4. Next, items with further dot positions come first.
//  5. If dot positions are identical, the order is determined by comparing productions.
//
// This function is useful for sorting LR(0) items, ensuring stability and predictability in their ordering.
func (i *Item0) Compare(rhs Item) int {
	ii, ok := rhs.(*Item0)
	if !ok {
		panic("Compare: rhs must be of type Item0")
	}

	if i.IsInitial() && !ii.IsInitial() {
		return -1
	} else if !i.IsInitial() && ii.IsInitial() {
		return 1
	}

	if i.IsKernel() && !ii.IsKernel() {
		return -1
	} else if !i.IsKernel() && ii.IsKernel() {
		return 1
	}

	if i.Production.Head.Equal(i.Start) && !ii.Production.Head.Equal(ii.Start) {
		return -1
	} else if !i.Production.Head.Equal(i.Start) && ii.Production.Head.Equal(ii.Start) {
		return 1
	}

	if i.Dot > ii.Dot {
		return -1
	} else if i.Dot < ii.Dot {
		return 1
	}

	return grammar.CmpProduction(i.Production, ii.Production)
}

// IsInitial checks if an LR(0) item is the initial item "S′ → •S" in the augmented grammar.
func (i *Item0) IsInitial() bool {
	return i.Production.Head.Equal(i.Start) && i.Dot == 0
}

// IsKernel determines whether an LR(0) item is a kernel item.
//
// Kernel items include:
//
//   - The initial item, "S′ → •S".
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are those where the dot is at the beginning, except for the initial item.
func (i *Item0) IsKernel() bool {
	return i.IsInitial() || i.Dot > 0
}

// IsComplete checks if the dot has reached the end of the item's body.
func (i *Item0) IsComplete() bool {
	return i.Dot == len(i.Body)
}

// IsInitial checks if an LR(0) item is the final item "S′ → S•" in the augmented grammar.
func (i *Item0) IsFinal() bool {
	return i.Production.Head.Equal(i.Start) && i.IsComplete()
}

// Dot returns the grammar symbol at the dot position in the item's body.
// If the dot is at the end of the body, it returns nil and false.
// For example, in the LR(0) item A → α•Bβ, it returns B.
func (i *Item0) DotSymbol() (grammar.Symbol, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return i.Body[i.Dot], true
}

// NextItem generates a new LR(0) item by advancing the dot one position forward in the item's body.
// If the dot is at the end of the body, it returns an empty LR(0) item and false.
// For example, for the LR(0) item A → α•Bβ, it returns A → αB•β.
func (i *Item0) Next() (Item, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return &Item0{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot + 1,
	}, true
}

// Item1 converts an LR(0) item into its equivalent LR(1) item by adding a lookahead.
func (i *Item0) Item1(lookahead grammar.Terminal) *Item1 {
	return &Item1{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot,
		Lookahead:  lookahead,
	}
}

// Item1 represents an LR(1) item for a context-free grammar.
// The "1" indicates the length of the lookahead of the item.
// Item1 implements the Item interace.
//
// An LR(1) item is a production with a dot at some position of the body,
// followed by either a terminal symbol or the right endmarker $.
//
//	A → α.β, a
//
// Informally, for an item of the form "A → α., a",
// a reduction by A → α is performed only if the next input symbol matches a.
// The lookahead does not impact an item of the form "A → α.β, a" when β is not ε.
type Item1 struct {
	*grammar.Production
	Start     grammar.NonTerminal // The start symbol S′ in the augmented grammar.
	Dot       int                 // Position of the dot in the production body.
	Lookahead grammar.Terminal
}

// String returns a string representation of an LR(1) item.
func (i *Item1) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%s → ", i.Head)

	if α := i.Body[:i.Dot]; len(α) > 0 {
		b.WriteString(α.String())
	}

	b.WriteRune('•')

	if β := i.Body[i.Dot:]; len(β) > 0 {
		b.WriteString(β.String())
	}

	fmt.Fprintf(&b, ", %s", i.Lookahead)

	return b.String()
}

// Equal determines whether or not two LR(1) items are the same.
func (i *Item1) Equal(rhs Item) bool {
	ii, ok := rhs.(*Item1)
	return ok &&
		i.Production.Equal(ii.Production) &&
		i.Start.Equal(ii.Start) &&
		i.Dot == ii.Dot &&
		i.Lookahead.Equal(ii.Lookahead)
}

// Compare compares two LR(1) items to establish a consistent and deterministic order.
// The comparison follows these rules:
//
//  1. Initial items precede non-initial items.
//  2. Kernel items precede non-kernel items.
//  3. If both items are kernel or non-kernel,
//     items with production heads equal to the augmented start symbol S′ come first.
//  4. Next, items with further dot positions come first.
//  5. If dot positions are identical, the order is determined by comparing productions.
//  6. If both dot positions and productions are identical,
//     the order is determined by comparing lookahead symbols.
//
// This function is useful for sorting LR(1) items, ensuring stability and predictability in their ordering.
func (i *Item1) Compare(rhs Item) int {
	ii, ok := rhs.(*Item1)
	if !ok {
		panic("Compare: rhs must be of type Item1")
	}

	if i.IsInitial() && !ii.IsInitial() {
		return -1
	} else if !i.IsInitial() && ii.IsInitial() {
		return 1
	}

	if i.IsKernel() && !ii.IsKernel() {
		return -1
	} else if !i.IsKernel() && ii.IsKernel() {
		return 1
	}

	if i.Production.Head.Equal(i.Start) && !ii.Production.Head.Equal(ii.Start) {
		return -1
	} else if !i.Production.Head.Equal(i.Start) && ii.Production.Head.Equal(ii.Start) {
		return 1
	}

	if i.Dot > ii.Dot {
		return -1
	} else if i.Dot < ii.Dot {
		return 1
	}

	if cmp := grammar.CmpProduction(i.Production, ii.Production); cmp < 0 {
		return -1
	} else if cmp > 0 {
		return 1
	}

	return grammar.CmpTerminal(i.Lookahead, ii.Lookahead)
}

// IsInitial checks if an LR(1) item is the initial item "S′ → •S, $" in the augmented grammar.
func (i *Item1) IsInitial() bool {
	return i.Production.Head.Equal(i.Start) && i.Dot == 0 &&
		i.Lookahead.Equal(grammar.Endmarker)
}

// IsKernel determines whether an LR(1) item is a kernel item.
//
// Kernel items include:
//
//   - The initial item, "S′ → •S, $".
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are those where the dot is at the beginning, except for the initial item.
func (i *Item1) IsKernel() bool {
	return i.IsInitial() || i.Dot > 0
}

// IsComplete checks if the dot has reached the end of the item's body.
func (i *Item1) IsComplete() bool {
	return i.Dot == len(i.Body)
}

// IsInitial checks if an LR(1) item is the final item "S′ → S•, $" in the augmented grammar.
func (i *Item1) IsFinal() bool {
	return i.Production.Head.Equal(i.Start) && i.IsComplete() &&
		i.Lookahead.Equal(grammar.Endmarker)
}

// Dot returns the grammar symbol at the dot position in the item's body.
// If the dot is at the end of the body, it returns nil and false.
// For example, in the LR(1) item A → α•Bβ, it returns B.
func (i *Item1) DotSymbol() (grammar.Symbol, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return i.Body[i.Dot], true
}

// NextItem generates a new LR(1) item by advancing the dot one position forward in the item's body.
// If the dot is at the end of the body, it returns an empty LR(1) item and false.
// For example, for the LR(1) item A → α•Bβ, it returns A → αB•β.
func (i *Item1) Next() (Item, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return &Item1{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot + 1,
		Lookahead:  i.Lookahead,
	}, true
}

// GetPrefix returns the prefix α of an LR(1) item of the form "A → α•β, a.
// α represents the portion of the production that has already been parsed.
// α can be the empty string ε if nothing has been parsed yet.
func (i *Item1) GetPrefix() grammar.String[grammar.Symbol] {
	return i.Body[:i.Dot]
}

// GetSuffix returns the suffix β of an LR(1) item of the form "A → α•β, a".
// β represents the portion of the production that has not yet been parsed.
// β can be the empty string ε if there is no remaining unparsed portion.
func (i *Item1) GetSuffix() grammar.String[grammar.Symbol] {
	return i.Body[i.Dot:]
}

// Item0 converts an LR(1) item into its equivalent LR(0) item by dropping the lookahead.
func (i *Item1) Item0() *Item0 {
	return &Item0{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot,
	}
}
