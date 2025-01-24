package lr

import "github.com/moorara/algo/grammar"

var (
	primeSuffixes = []string{
		"′", // Prime (U+2032)
		"″", // Double Prime (U+2033)
		"‴", // Triple Prime (U+2034)
		"⁗", // Quadruple Prime (U+2057)
	}
)

// calculator defines the interface required for an LR parser.
type calculator interface {
	// G returns the augmented context-free grammar.
	G() *grammar.CFG

	// Initial returns the initial item of an augmented grammar.
	// For LR(0), the initial item is "S′ → •S",
	// and for LR(1), the initial item is "S′ → •S, $".
	Initial() Item

	// Closure computes the closure of a given item set.
	// It may compute the closure for either an LR(0) item set or an LR(1) item set.
	CLOSURE(ItemSet) ItemSet
}

// calculator0 implemented the Calculator interface for LR(0) items.
type calculator0 struct {
	augG *grammar.CFG
}

// G returns the augmented context-free grammar.
func (c *calculator0) G() *grammar.CFG {
	return c.augG
}

// Initial returns the initial LR(0) item "S′ → •S" for an augmented grammar.
func (c *calculator0) Initial() Item {
	for p := range c.augG.Productions.Get(c.augG.Start).All() {
		return &Item0{
			Production: p,
			Start:      c.augG.Start,
			Dot:        0,
		}
	}

	// This will never be the case.
	return nil
}

// CLOSURE computes the closure of a given LR(0) item set.
func (c *calculator0) CLOSURE(I ItemSet) ItemSet {
	J := I.Clone()

	for newItems := []Item{}; newItems != nil; {
		newItems = nil

		// For each item A → α•Bβ in J
		for i := range J.All() {
			if i, ok := i.(*Item0); ok {
				if X, ok := i.DotSymbol(); ok {
					if B, ok := X.(grammar.NonTerminal); ok {
						// For each production B → γ of G′
						for BProd := range c.augG.Productions.Get(B).All() {
							j := &Item0{
								Production: BProd,
								Start:      c.augG.Start,
								Dot:        0,
							}

							// If B → •γ is not in J
							if !J.Contains(j) {
								newItems = append(newItems, j)
							}
						}
					}
				}
			}
		}

		J.Add(newItems...)
	}

	return J
}

// calculator1 implemented the Calculator interface for LR(1) items.
type calculator1 struct {
	augG  *grammar.CFG
	FIRST grammar.FIRST
}

// G returns the augmented context-free grammar.
func (c *calculator1) G() *grammar.CFG {
	return c.augG
}

// Initial returns the initial LR(1) item "S′ → •S, $" for an augmented grammar.
func (c *calculator1) Initial() Item {
	for p := range c.augG.Productions.Get(c.augG.Start).All() {
		return &Item1{
			Production: p,
			Start:      c.augG.Start,
			Dot:        0,
			Lookahead:  grammar.Endmarker,
		}
	}

	// This will never be the case.
	return nil
}

// CLOSURE computes the closure of a given LR(1) item set.
func (c *calculator1) CLOSURE(I ItemSet) ItemSet {
	J := I.Clone()

	for newItems := []Item{}; newItems != nil; {
		newItems = nil

		// For each item [A → α•Bβ, a] in J
		for i := range J.All() {
			if i, ok := i.(*Item1); ok {
				a := i.Lookahead
				if X, ok := i.DotSymbol(); ok {
					if B, ok := X.(grammar.NonTerminal); ok {
						// For each production B → γ of G′
						for BProd := range c.augG.Productions.Get(B).All() {
							β := i.GetSuffix()[1:]
							βa := β.Append(a)

							// For each terminal b in FIRST(βa)
							for b := range c.FIRST(βa).Terminals.All() {
								j := &Item1{
									Production: BProd,
									Start:      c.augG.Start,
									Dot:        0,
									Lookahead:  b,
								}

								// If [B → •γ, b] is not in J
								if !J.Contains(j) {
									newItems = append(newItems, j)
								}
							}
						}
					}
				}
			}
		}

		J.Add(newItems...)
	}

	return J
}
