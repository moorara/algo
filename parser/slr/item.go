package slr

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/parser/lr"
)

// LR0Item represents an LR(0) item for a context-free grammar.
// It implements the Item interace.
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
type LR0Item struct {
	*grammar.Production
	Start *grammar.NonTerminal // The start symbol S′ in the augmented grammar.
	Dot   int                  // Position of the dot in the production body.
}

// String returns a string representation of an LR(0) item.
func (i LR0Item) String() string {
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

// Equals determines whether or not two LR(0) items are the same.
func (i LR0Item) Equals(rhs lr.Item) bool {
	ii, ok := rhs.(LR0Item)
	return ok &&
		i.Production.Equals(*ii.Production) &&
		i.Start.Equals(*ii.Start) &&
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
func (i LR0Item) Compare(rhs lr.Item) int {
	ii, ok := rhs.(LR0Item)
	if !ok {
		panic("Compare: rhs must be of type LR0Item")
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

	if i.Production.Head.Equals(*i.Start) && !ii.Production.Head.Equals(*ii.Start) {
		return -1
	} else if !i.Production.Head.Equals(*i.Start) && ii.Production.Head.Equals(*ii.Start) {
		return 1
	}

	if i.Dot > ii.Dot {
		return -1
	} else if i.Dot < ii.Dot {
		return 1
	}

	return grammar.CmpProduction(*i.Production, *ii.Production)
}

// IsInitial checks if an LR(0) item is the initial item "S′ → •S" in the augmented grammar.
func (i LR0Item) IsInitial() bool {
	return i.Production.Head.Equals(*i.Start) && i.Dot == 0
}

// IsKernel determines whether an LR(0) item is a kernel item.
//
// Kernel items include:
//
//   - The initial item, "S′ → •S".
//   - All items where the dot is not at the beginning (left end) of the item's body.
//
// Non-kernel items are those where the dot is at the beginning, except for the initial item.
func (i LR0Item) IsKernel() bool {
	return i.IsInitial() || i.Dot > 0
}

// IsComplete checks if the dot has reached the end of the item's body.
func (i LR0Item) IsComplete() bool {
	return i.Dot == len(i.Body)
}

// IsInitial checks if an LR(0) item is the final item "S′ → S•" in the augmented grammar.
func (i LR0Item) IsFinal() bool {
	return i.Production.Head.Equals(*i.Start) && i.IsComplete()
}

// Dot returns the grammar symbol at the dot position in the item's body.
// If the dot is at the end of the body, it returns nil and false.
// For example, in the LR(0) item A → α•Bβ, it returns B.
func (i LR0Item) DotSymbol() (grammar.Symbol, bool) {
	if i.IsComplete() {
		return nil, false
	}

	return i.Body[i.Dot], true
}

// NextItem generates a new LR(0) item by advancing the dot one position forward in the item's body.
// If the dot is at the end of the body, it returns an empty LR(0) item and false.
// For example, for the LR(0) item A → α•Bβ, it returns A → αB•β.
func (i LR0Item) Next() (lr.Item, bool) {
	if i.IsComplete() {
		return LR0Item{}, false
	}

	return LR0Item{
		Production: i.Production,
		Start:      i.Start,
		Dot:        i.Dot + 1,
	}, true
}
