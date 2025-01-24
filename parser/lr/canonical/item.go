package canonical

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// LR1Item represents an LR(1) item for a context-free grammar.
// The "1" indicates the length of the lookahead of the item.
// LR1Item implements the Item interace.
//
// An LR(1) item is a production with a dot at some position of the body,
// followed by either a terminal symbol or the right endmarker $.
//
//	A → α.β, a
//
// Informally, for an item of the form "A → α., a",
// a reduction by A → α is performed only if the next input symbol matches a.
// The lookahead does not impact an item of the form "A → α.β, a" when β is not ε.
type LR1Item struct {
	*grammar.Production
	Start     grammar.NonTerminal // The start symbol S′ in the augmented grammar.
	Dot       int                 // Position of the dot in the production body.
	Lookahead grammar.Terminal
}

// String returns a string representation of an LR(1) item.
func (i *LR1Item) String() string {
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

// Equals determines whether or not two LR(1) items are the same.
func (i *LR1Item) Equals(rhs lr.Item) bool {
	ii, ok := rhs.(*LR1Item)
	return ok &&
		i.Production.Equals(ii.Production) &&
		i.Start.Equals(ii.Start) &&
		i.Dot == ii.Dot &&
		i.Lookahead.Equals(ii.Lookahead)
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
func (i *LR1Item) Compare(rhs lr.Item) int {
	ii, ok := rhs.(*LR1Item)
	if !ok {
		panic("Compare: rhs must be of type LR1Item")
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

	if i.Production.Head.Equals(i.Start) && !ii.Production.Head.Equals(ii.Start) {
		return -1
	} else if !i.Production.Head.Equals(i.Start) && ii.Production.Head.Equals(ii.Start) {
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
func (i *LR1Item) IsInitial() bool {
	return i.Production.Head.Equals(i.Start) &&
		i.Dot == 0 &&
		i.Lookahead.Equals(grammar.Endmarker)
}

// IsKernel determines whether an LR(1) item is a kernel item.
//
// Kernel items include:
//
//   - The initial item, "S′ → •S, $".
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are those where the dot is at the beginning, except for the initial item.
func (i *LR1Item) IsKernel() bool {
	return i.IsInitial() || i.Dot > 0
}

// IsComplete checks if the dot has reached the end of the item's body.
func (i *LR1Item) IsComplete() bool {
	return i.Dot == len(i.Body)
}

// IsInitial checks if an LR(1) item is the final item "S′ → S•, $" in the augmented grammar.
func (i *LR1Item) IsFinal() bool {
	return i.Production.Head.Equals(i.Start) &&
		i.IsComplete() &&
		i.Lookahead.Equals(grammar.Endmarker)
}

// Dot returns the grammar symbol at the dot position in the item's body.
// If the dot is at the end of the body, it returns nil and false.
// For example, in the LR(1) item A → α•Bβ, it returns B.
func (i *LR1Item) DotSymbol() (grammar.Symbol, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return i.Body[i.Dot], true
}

// NextItem generates a new LR(1) item by advancing the dot one position forward in the item's body.
// If the dot is at the end of the body, it returns an empty LR(1) item and false.
// For example, for the LR(1) item A → α•Bβ, it returns A → αB•β.
func (i *LR1Item) Next() (lr.Item, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return &LR1Item{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot + 1,
		Lookahead:  i.Lookahead,
	}, true
}

// GetαPrefix returns the α (prefix) part of an LR(1) item of the form "A → α•Bβ, a".
// α represents the portion of the production that has already been parsed.
// α can be the empty string (ε) if nothing has been parsed yet.
func (i *LR1Item) GetαPrefix() grammar.String[grammar.Symbol] {
	return i.Body[:i.Dot]
}

// GetβSuffix returns the β (suffix) part of an LR(1) item of the form "A → α•Bβ, a".
// β represents the portion of the production that has not yet been parsed.
// β can be the empty string (ε) if there is no remaining unparsed portion.
func (i *LR1Item) GetβSuffix() grammar.String[grammar.Symbol] {
	if i.IsComplete() {
		return grammar.E
	}

	return i.Body[i.Dot+1:]
}
