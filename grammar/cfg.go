package grammar

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

var (
	primeSuffixes = []string{
		"′", // Prime U+2032
		"″", // Double Prime (U+2033)
		"‴", // Triple Prime (U+2034)
		"⁗", // Quadruple Prime (U+2057)
	}

	alphabeticSuffixes = []string{
		"ₙ", // Subscript Small N (U+2099)
		"ⁿ", // Superscript Small N (U+207F)
		"ᴺ", // Modifier Capital N (U+1D3A)
	}

	numericSuffixes = []string{
		// Subscript Zero (U+2080)
		"₁", // Subscript One (U+2081)
		"₂", // Subscript Two (U+2082)
		"₃", // Subscript Three (U+2083)
		"₄", // Subscript Four (U+2084)
		"₅", // Subscript Five (U+2085)
		"₆", // Subscript Six (U+2086)
		"₇", // Subscript Seven (U+2087)
		"₈", // Subscript Eight (U+2088)
		"₉", // Subscript Nine (U+2089)
		"₁₀", "₁₁", "₁₂", "₁₃", "₁₄", "₁₅", "₁₆", "₁₇", "₁₈", "₁₉",
		"₂₀", "₂₁", "₂₂", "₂₃", "₂₄", "₂₅", "₂₆", "₂₇", "₂₈", "₂₉",
		"₃₀", "₃₁", "₃₂", "₃₃", "₃₄", "₃₅", "₃₆", "₃₇", "₃₈", "₃₉",
		"₄₀", "₄₁", "₄₂", "₄₃", "₄₄", "₄₅", "₄₆", "₄₇", "₄₈", "₄₉",
		"₅₀", "₅₁", "₅₂", "₅₃", "₅₄", "₅₅", "₅₆", "₅₇", "₅₈", "₅₉",
		"₆₀", "₆₁", "₆₂", "₆₃", "₆₄", "₆₅", "₆₆", "₆₇", "₆₈", "₆₉",
		"₇₀", "₇₁", "₇₂", "₇₃", "₇₄", "₇₅", "₇₆", "₇₇", "₇₈", "₇₉",
		"₈₀", "₈₁", "₈₂", "₈₃", "₈₄", "₈₅", "₈₆", "₈₇", "₈₈", "₈₉",
		"₉₀", "₉₁", "₉₂", "₉₃", "₉₄", "₉₅", "₉₆", "₉₇", "₉₈", "₉₉",
	}
)

// CFG represents a context-free grammar in formal language theory.
//
// Context-free grammars can express a wide range of programming language constructs
// while remaining computationally efficient to parse.
// They are used in computer science and linguistics to describe the syntax of languages.
//
// A context-free grammar G = (V, Σ, R, S) is defined by four sets:
//
//  1. V is a set of terminal symbols from which strings are formed.
//     Terminal symbols are also referred to as tokens.
//
//  2. Σ is a set of non-terminals symbols that denote sets of strings.
//     Non-terminal symbols are sometimes called syntactic variables.
//     Non-terminals impose a hierarchical structure on the language.
//
//  3. R = V × (V ∪ Σ)* is a set of productions, where each production consists of
//     a non-terminal (head), an arrow, and a sequence of terminals and/or non-terminals (body).
//
//  4. S ∈ V is one of the non-terminal symbols designated as the start symbol.
//     The set of strings denoted by the start symbol is the language generated by the grammar.
//
// Context-free languages are a superset of regular languages and they are more expressive.
type CFG struct {
	Terminals    set.Set[Terminal]
	NonTerminals set.Set[NonTerminal]
	Productions  CFProductions
	Start        NonTerminal
}

// NewCFG creates a new context-free grammar.
func NewCFG(terms []Terminal, nonTerms []NonTerminal, prods []CFProduction, start NonTerminal) CFG {
	g := CFG{
		Terminals:    set.New(eqTerminal),
		NonTerminals: set.New(eqNonTerminal),
		Productions:  NewCFProductions(),
		Start:        start,
	}

	g.Terminals.Add(terms...)
	g.NonTerminals.Add(nonTerms...)
	g.Productions.Add(prods...)

	return g
}

// verify takes a context-free grammar and determines whether or not it is valid.
// If the given grammar is invalid, an error with a descriptive message will be returned.
func (g CFG) Verify() error {
	var err error

	getPredicate := func(n NonTerminal) Predicate1[CFProduction] {
		return func(p CFProduction) bool {
			return p.Head.Equals(n)
		}
	}

	// Check if the start symbol is in the set of non-terminal symbols.
	if !g.NonTerminals.Contains(g.Start) {
		err = errors.Join(err, fmt.Errorf("start symbol %s not in the set of non-terminal symbols", g.Start))
	}

	// Check if there is at least one production rule for the start symbol.
	if !g.Productions.AnyMatch(getPredicate(g.Start)) {
		err = errors.Join(err, fmt.Errorf("no production rule for start symbol %s", g.Start))
	}

	// Check if there is at least one prodcution rule for every non-terminal symbol.
	for n := range g.NonTerminals.All() {
		if !g.Productions.AnyMatch(getPredicate(n)) {
			err = errors.Join(err, fmt.Errorf("no production rule for non-terminal symbol %s", n))
		}
	}

	for p := range g.Productions.All() {
		// Check if the head of production rule is in the set of non-terminal symbols.
		if !g.NonTerminals.Contains(p.Head) {
			err = errors.Join(err, fmt.Errorf("production head %s not in the set of non-terminal symbols", p.Head))
		}

		// Check if every symbol in the body of production rule is either in the set of terminal or non-terminal symbols.
		for _, s := range p.Body {
			if v, ok := s.(Terminal); ok && !g.Terminals.Contains(v) {
				err = errors.Join(err, fmt.Errorf("terminal symbol %s not in the set of terminal symbols", v))
			}

			if v, ok := s.(NonTerminal); ok && !g.NonTerminals.Contains(v) {
				err = errors.Join(err, fmt.Errorf("non-terminal symbol %s not in the set of non-terminal symbols", v))
			}
		}
	}

	return err
}

// Equals determines whether or not two context-free grammars are the same.
func (g CFG) Equals(rhs CFG) bool {
	return g.Terminals.Equals(rhs.Terminals) &&
		g.NonTerminals.Equals(rhs.NonTerminals) &&
		g.Productions.Equals(rhs.Productions) &&
		g.Start.Equals(rhs.Start)
}

// Clone returns a deep copy of a context-free grammar, ensuring the clone is independent of the original.
func (g CFG) Clone() CFG {
	return CFG{
		Terminals:    g.Terminals.Clone(),
		NonTerminals: g.NonTerminals.Clone(),
		Productions:  g.Productions.Clone(),
		Start:        g.Start,
	}
}

// NullableNonTerminals finds all non-terminal symbols in a context-free grammar
// that can derive the empty string ε in one or more steps (A ⇒* ε for some non-terminal A).
func (g CFG) NullableNonTerminals() set.Set[NonTerminal] {
	// Define a set for all non-terminals that can derive the empty string ε
	nullable := set.New(eqNonTerminal)

	for updated := true; updated; {
		updated = false

		// Iterate through each production rule of the form A → α,
		// where A is a non-terminal symbol and α is a string of terminals and non-terminals.
		for head, list := range g.Productions.AllByHead() {
			// Skip the production rule if A is already in the nullable set.
			if nullable.Contains(head) {
				continue
			}

			for p := range list.All() {
				if p.IsEmpty() {
					// α is the empty string ε, add A to the nullable set.
					nullable.Add(p.Head)
					updated = true
				} else if n := p.Body.NonTerminals(); len(n) == len(p.Body) && nullable.Contains(n...) {
					// α consists of only non-terminal symbols already in the nullable set, add A to the nullable set.
					nullable.Add(p.Head)
					updated = true
				}
			}
		}
	}

	return nullable
}

// EliminateEmptyProductions converts a context-free grammar into an equivalent ε-free grammar.
//
// An empty production (ε-production) is any production of the form A → ε.
func (g CFG) EliminateEmptyProductions() CFG {
	nullable := g.NullableNonTerminals()

	newG := CFG{
		Terminals:    g.Terminals.Clone(),
		NonTerminals: g.NonTerminals.Clone(),
		Productions:  NewCFProductions(),
		Start:        g.Start,
	}

	// Iterate through each production rule in the input grammar.
	// For each production rule of the form A → α,
	//   generate all possible combinations of α by including and excluding nullable non-terminals.
	for p := range g.Productions.All() {
		// Ignore ε-production rules (A → ε)
		// Only consider the production rules of the form A → α
		if p.IsEmpty() {
			continue
		}

		// bodies holds all possible combinations of the right-hand side of a production rule.
		bodies, aux := []String[Symbol]{ε}, []String[Symbol]{}

		// Every nullable non-terminal symbol creates two possibilities, once by including and once by excluding it.
		for _, sym := range p.Body {
			v, ok := sym.(NonTerminal)
			nonTermNullable := ok && nullable.Contains(v)

			for _, β := range bodies {
				if nonTermNullable {
					aux = append(aux, β)
				}
				aux = append(aux, append(β, sym))
			}

			bodies, aux = aux, nil
		}

		for _, β := range bodies {
			// Skip ε-production rules (A → ε)
			if len(β) > 0 {
				newG.Productions.Add(CFProduction{p.Head, β})
			}
		}
	}

	// The set data structure automatically prevents duplicate items from being added.
	// Therefore, we don't need to worry about deduplicating the new production rules at this stage.

	// If the start symbol of the grammer is nullable (S ⇒* ε),
	//   a new start symbol with an ε-production rule must be introduced (S′ → S | ε).
	// This guarantees that the resulting grammar generates the same language as the original grammar.
	if start := newG.Start; nullable.Contains(start) {
		newStart := newG.addNewNonTerminal(start, primeSuffixes...)
		newG.Start = newStart
		newG.Productions.Add(CFProduction{newStart, String[Symbol]{start}}) // S′ → S
		newG.Productions.Add(CFProduction{newStart, ε})                     // S′ → ε
	}

	return newG
}

// EliminateSingleProductions converts a context-free grammar into an equivalent single-production-free grammar.
//
// A single production a.k.a. unit production is a production rule whose body is a single non-terminal symbol (A → B).
func (g CFG) EliminateSingleProductions() CFG {
	// Identify all single productions.
	singleProds := map[NonTerminal][]NonTerminal{}
	for p := range g.Productions.All() {
		if p.IsSingle() {
			singleProds[p.Head] = append(singleProds[p.Head], p.Body[0].(NonTerminal))
		}
	}

	// Compute the transitive closure for all non-terminal symbols.
	// The transitive closure of a non-terminal A is the the set of all non-terminals B
	//   such that there exists a sequence of single productions starting from A and reaching B (i.e., A → B₁ → B₂ → ... → B).

	closure := make(map[NonTerminal]map[NonTerminal]bool, g.NonTerminals.Size())

	// Initially, each non-terminal symbol is reachable from itself.
	for A := range g.NonTerminals.All() {
		closure[A] = map[NonTerminal]bool{A: true}
	}

	// Next, add directly reachable non-terminal symbols from single productions.
	for A, nonTerms := range singleProds {
		for _, B := range nonTerms {
			closure[A][B] = true
		}
	}

	// Repeat until no new non-terminal symbols can be added to the closure set.
	for updated := true; updated; {
		updated = false

		for A, closureA := range closure {
			for B := range closureA {
				for next := range closure[B] {
					if !closureA[next] {
						closure[A][next] = true
						updated = true
					}
				}
			}
		}
	}

	newG := CFG{
		Terminals:    g.Terminals.Clone(),
		NonTerminals: g.NonTerminals.Clone(),
		Productions:  NewCFProductions(),
		Start:        g.Start,
	}

	// For each production rule p of the form B → α, add a new production rule A → α
	//   if p is not a single production and B is in the transitive closure set of A.
	for A, closureA := range closure {
		for B := range closureA {
			for p := range g.Productions.Get(B).All() {
				// Skip single productions
				if !p.IsSingle() {
					newG.Productions.Add(CFProduction{A, p.Body})
				}
			}
		}
	}

	return newG
}

// EliminateUnreachableProductions converts a context-free grammar into an equivalent grammar
// with all unreachable productions and their associated non-terminal symbols removed.
//
// An unreachable production refers to a production rule in a grammar
// that cannot be used to derive any string starting from the start symbol.
//
// The function also removes unreachable terminals,
// which are terminals that do not appear in any reachable production.
func (g CFG) EliminateUnreachableProductions() CFG {
	reachableT := set.New(eqTerminal)
	reachableN := set.New(eqNonTerminal, g.Start)
	reachableP := NewCFProductions()

	// Reppeat until no new non-terminal is added to reachable non-terminals:
	//   For each production rule of the form A → α:
	//     If A is in reachable non-terminals, add all non-terminal in α to reachable non-terminals too.
	for updated := true; updated; {
		updated = false

		for p := range g.Productions.All() {
			if reachableN.Contains(p.Head) {
				for _, n := range p.Body.NonTerminals() {
					if !reachableN.Contains(n) {
						reachableN.Add(n)
						updated = true
					}
				}
			}
		}
	}

	// Gather reachable productions.
	for p := range g.Productions.All() {
		if reachableN.Contains(p.Head) {
			reachableP.Add(p)
		}
	}

	// Gather reachable terminals.
	for t := range g.Terminals.All() {
		if reachableP.AnyMatch(func(p CFProduction) bool {
			return p.Body.ContainsSymbol(t)
		}) {
			reachableT.Add(t)
		}
	}

	return CFG{
		Terminals:    reachableT,
		NonTerminals: reachableN,
		Productions:  reachableP,
		Start:        g.Start,
	}
}

// EliminateCycles converts a context-free grammar into an equivalent cycle-free grammar.
//
// A grammar is cyclic if it has derivations of one or more steps in which A ⇒* A for some non-terminal A.
func (g CFG) EliminateCycles() CFG {
	// Single productions (unit productions) can create cycles in a grammar.
	// Eliminating empty productions (ε-productions) may introduce additional single productions,
	// so it is necessary to eliminate empty productions first, followed by single productions.
	// After removing single productions, some productions may become unreachable.
	// These unreachable productions should then be removed from the grammar.
	return g.EliminateEmptyProductions().EliminateSingleProductions().EliminateUnreachableProductions()
}

// EliminateLeftRecursion converts a context-free grammar into an equivalent grammar with no left recursion.
//
// A grammar is left-recursive if it has a non-terminal A such that there is a derivation A ⇒+ Aα for some string.
// For top-down parsers, left recursion causes the parser to loop forever.
// Many bottom-up parsers also will not accept left-recursive grammars.
//
// Note that the resulting non-left-recursive grammar may have ε-productions.
func (g CFG) EliminateLeftRecursion() CFG {
	// Define predicates for identifying left-recursive and non-left-recursive productions
	isLeftRecursivePredicate := func(p CFProduction) bool { return p.IsLeftRecursive() }
	isNotLeftRecursivePredicate := func(p CFProduction) bool { return !p.IsLeftRecursive() }

	// The algorithm implemented here is guaranteed to work if the grammar has no cycles or ε-productions.
	newG := g.EliminateCycles()

	// Arrange the non-terminals in some order.
	// The exact order does not affect the eliminition of left recursions (immediate or indirect),
	//   but the resulting grammar can depend on the order in which non-terminals are processed.
	_, _, nonTerms := newG.orderNonTerminals()

	for i := 0; i < len(nonTerms); i++ {
		for j := 0; j < i-1; j++ {
			/*
			 * Replace each production of the form Aᵢ → Aⱼγ by the productions Aᵢ → δ₁γ | δ₂γ | ... | δₖγ,
			 * where Aⱼ → δ₁ | δ₂ | ... | δₖ are all current Aⱼ-productions.
			 */

			Ai, Aj := nonTerms[i], nonTerms[j]
			AiProds, AjProds := newG.Productions.Get(Ai), newG.Productions.Get(Aj)

			AiAjProds := AiProds.SelectMatch(func(p CFProduction) bool {
				return len(p.Body) > 0 && p.Body[0].Equals(Aj)
			})

			for AiAjProd := range AiAjProds.All() {
				newG.Productions.Remove(AiAjProd)
				for AjProd := range AjProds.All() {
					p := CFProduction{Ai, AjProd.Body.Concat(AiAjProd.Body[1:])}
					newG.Productions.Add(p)
				}
			}
		}

		/*
		 * Immediate left recursion can be eliminated by the following technique,
		 * which works for any number of A-productions.
		 *
		 * First, group the productions as
		 *
		 *   A → Aα₁ | Aα₂ | ... | Aαₘ | β₁ | β₂ | ... | βₙ
		 *
		 * where no αᵢ is ε and no βᵢ begins with an A. Then replace A-productions by
		 *
		 *    A → β₁A′ | β₂A′ | ... | βₙA′
		 *    A′ → α₁A′ | α₂A′ | ... | αₘA′ | ε
		 */

		A := nonTerms[i]
		AProds := newG.Productions.Get(A)
		hasLR := AProds.AnyMatch(isLeftRecursivePredicate)

		if hasLR {
			Anew := newG.addNewNonTerminal(A, primeSuffixes...)

			LRProds := AProds.SelectMatch(isLeftRecursivePredicate)       // Immediately Left-Recursive A-productions
			nonLRProds := AProds.SelectMatch(isNotLeftRecursivePredicate) // Not Immediately Left-Recursive A-productions

			// Remove A → Aα₁ | Aα₂ | ... | Aαₘ | β₁ | β₂ | ... | βₙ
			newG.Productions.RemoveAll(A)

			// Add A → β₁A′ | β₂A′ | ... | βₙA′
			for nonLRProd := range nonLRProds.All() {
				newG.Productions.Add(CFProduction{A, nonLRProd.Body.Append(Anew)})
			}

			// Single productions of the form A → A, where α = ε, are already eliminated.
			// Add A′ → α₁A′ | α₂A′ | ... | αₘA′ | ε
			for LRProd := range LRProds.All() {
				newG.Productions.Add(CFProduction{Anew, LRProd.Body[1:].Append(Anew)})
			}

			// Add A′ → ε
			newG.Productions.Add(CFProduction{Anew, ε})
		}
	}

	return newG
}

// LeftFactor converts a context-free grammar into an equivalent left-factored grammar.
//
// Left factoring is a grammar transformation for producing a grammar suitable predictive for top-down parsing.
// When the choice between two alternative A-productions is not clear,
// we may be able to rewrite the productions to defer the decision
// until enough of the input has been seen that we can make the right choice.
//
// For example, if we have the two productions
//
//	𝑠𝑡𝑚𝑡 → 𝐢𝐟 𝑒𝑥𝑝𝑟 𝐭𝐡𝐞𝐧 𝑠𝑡𝑚𝑡 𝐞𝐥𝐬𝐞 𝑠𝑡𝑚𝑡
//	    | 𝐢𝐟 𝑒𝑥𝑝𝑟 𝐭𝐡𝐞𝐧 𝑠𝑡𝑚𝑡
//
// on seeing the input 𝐢𝐟, we cannot immediately tell which productions to choose to expand 𝑠𝑡𝑚𝑡.
//
// Note that the resulting  left-factored grammar may have ε-productions and/or single productions.
func (g CFG) LeftFactor() CFG {
	/*
	 * For each non-terminal A, find the longest prefix α common to two or more A-productions.
	 * If α ≠ ε, there is a non-trivial common prefix, replace all of the A-productions
	 *
	 *   A → αβ₁ | αβ₂ | ... | αβₙ | γ
	 *
	 * where γ represents all the alternative productions that do not being with α, by
	 *
	 *   A → αA′ | γ
	 *   A′ → β₁ | β₂ | ... | βₙ
	 *
	 * We repeatedly apply this transformation until
	 * no two alternative productions for a non-terminal have a common prefix.
	 */

	newG := g.Clone()

	for updated := true; updated; {
		updated = false

		for A, AProds := range newG.Productions.AllByHead() {
			// Group production bodies by their common prefixes.
			groups := groupByCommonPrefix(AProds)

			// Select groups with two or more suffixes.
			// These correspond to A-productions A → αβ₁ | αβ₂ | ... | αβₙ
			prefixGroups := groups.SelectMatch(func(prefix String[Symbol], suffixes set.Set[String[Symbol]]) bool {
				return suffixes.Size() >= 2
			})

			// Select groups with exactly one suffix.
			// These correspond to alternative A-productions A → γ
			altGroups := groups.SelectMatch(func(prefix String[Symbol], suffixes set.Set[String[Symbol]]) bool {
				return suffixes.Size() == 1
			})

			if prefixGroups.Size() > 0 && altGroups.Size() > 0 {
				// Remove all A-productions A → αβ₁ | αβ₂ | ... | αβₙ | γ
				AProds.RemoveAll()

				for prefix, suffixes := range prefixGroups.All() {
					Anew := newG.addNewNonTerminal(A, primeSuffixes...)

					// Add A-production A → αA′
					newG.Productions.Add(CFProduction{A, prefix.Append(Anew)})

					// Add A′-productions A′ → β₁ | β₂ | ... | βₙ
					for suffix := range suffixes.All() {
						newG.Productions.Add(CFProduction{Anew, suffix})
					}
				}

				// Add alternative A-productions A → γ
				for prefix, suffixes := range altGroups.All() {
					for suffix := range suffixes.All() {
						newG.Productions.Add(CFProduction{A, prefix.Concat(suffix)})
					}
				}
			}
		}
	}

	return newG
}

// groupByCommonPrefix groups production bodies by their common prefixes.
// It prioritizes shorter prefixes that encompass more suffixes and production bodies
// over longer prefixes that encompass fewer suffixes or production bodies.
func groupByCommonPrefix(prods set.Set[CFProduction]) symboltable.SymbolTable[String[Symbol], set.Set[String[Symbol]]] {
	// Define a map of prefixes to their corresponding suffixes.
	groups := symboltable.NewQuadraticHashTable[String[Symbol], set.Set[String[Symbol]]](
		HashFuncForSymbolString(nil),
		eqString,
		eqStringSet,
		symboltable.HashOpts{},
	)

	for prod := range prods.All() {
		prefixFound := false

		// Attempt to find an existing prefix for the current production body.
		for prefix := range groups.All() {
			// Compute the longest common prefix between the current production body and an existing prefix in the groups.
			commonPrefix := String[Symbol]{}
			for i := 0; i < len(prefix) && i < len(prod.Body) && prefix[i].Equals(prod.Body[i]); i++ {
				commonPrefix = commonPrefix.Append(prefix[i])
			}

			// If a common prefix is found,
			// add the remaining part of the current production body as a suffix to the prefix group.
			if len(commonPrefix) > 0 {
				suffix := prod.Body[len(commonPrefix):]
				suffixes, _ := groups.Get(commonPrefix)
				suffixes.Add(suffix)
				prefixFound = true
				break
			}
		}

		// If no matching prefix is found,
		// initialize a new prefix with the first symbol of the production body and store the remaining part as the suffix.
		if !prefixFound {
			var prefix, suffix String[Symbol]
			if prod.IsEmpty() {
				prefix, suffix = ε, ε
			} else {
				prefix, suffix = prod.Body[:1], prod.Body[1:]
			}

			suffixes := set.New[String[Symbol]](eqString, suffix)
			groups.Put(prefix, suffixes)
		}
	}

	return groups
}

// ChomskyNormalForm converts a context-free grammar into an equivalent grammar in Chomsky Normal Form.
//
// A context-free grammar G is in Chomsky Normal Form (CNF) if all of its production rules are of the form
//
//   - A → BC, or
//   - A → a, or
//   - S → ε
//
// where A, B, and C are non-terminal symbols, a is a terminal symbol, and S is the start symbol.
// Also, neither B nor C may be the start symbol, and the third production rule can only appear if ε is in L(G).
func (g CFG) ChomskyNormalForm() CFG {
	/*
	 * The order in which these transformations are applied is critical.
	 * Some transformations may undo the effects of other ones.
	 *
	 * The blow-up in grammar size depends on the order between DEL and BIN.
	 * It may be exponential when DEL is applied first, but is linear otherwise.
	 * UNIT can incur a quadratic blow-up in the size of the grammar.
	 *
	 * The orderings START,TERM,BIN,DEL,UNIT and START,BIN,DEL,UNIT,TERM lead to
	 *   the least blow-up in grammar size while preserving the results of previous transformations.
	 */

	return g.eliminateStartSymbolFromRight(). // START
							eliminateNonSolitaryTerminals(). // TERM
							eliminateNonBinaryProductions(). // BIN
							EliminateEmptyProductions().     // DEL
							EliminateSingleProductions().    // UNIT
							EliminateUnreachableProductions()
}

// eliminateStartSymbolFromRight eliminate the start symbol from the right-hand side of any production rules
// by introducing a new start symbol S′ and add the production rule S′ → S.
func (g CFG) eliminateStartSymbolFromRight() CFG {
	newG := g.Clone()

	if S := newG.Start; newG.Productions.AnyMatch(func(p CFProduction) bool {
		return p.Body.ContainsSymbol(S)
	}) {
		Snew := newG.addNewNonTerminal(S, primeSuffixes...)
		newG.Start = Snew
		newG.Productions.Add(CFProduction{Snew, String[Symbol]{S}})
	}

	return newG
}

// eliminateNonSolitaryTerminals eliminate production rules with non-solitary terminals.
// It replaces terminals mixed with non-terminals in productions by introducing intermediate non-terminals.
func (g CFG) eliminateNonSolitaryTerminals() CFG {
	store := map[Terminal]NonTerminal{}

	newG := CFG{
		Terminals:    g.Terminals.Clone(),
		NonTerminals: g.NonTerminals.Clone(),
		Productions:  NewCFProductions(),
		Start:        g.Start,
	}

	for p := range g.Productions.All() {
		// Skip productions of the form A → a.
		if _, isTerminal := p.IsCNF(); isTerminal {
			newG.Productions.Add(p)
			continue
		}

		var newBody String[Symbol]
		for _, sym := range p.Body {
			if t, ok := sym.(Terminal); ok {
				newN, exist := store[t]
				if !exist {
					newN = newG.addNewNonTerminal(NonTerminal(t), alphabeticSuffixes...)
					store[t] = newN
				}

				newBody = append(newBody, newN)
				newG.Productions.Add(CFProduction{newN, String[Symbol]{sym}})
			} else {
				newBody = append(newBody, sym)
			}
		}

		newG.Productions.Add(CFProduction{p.Head, newBody})
	}

	return newG
}

// eliminateNonBinaryProductions eliminates production rules with more than two non-terminals in their right-hand sides.
// It replaces each production with a right-hand side longer than two non-terminals with intermediate non-terminals.
func (g CFG) eliminateNonBinaryProductions() CFG {
	newG := CFG{
		Terminals:    g.Terminals.Clone(),
		NonTerminals: g.NonTerminals.Clone(),
		Productions:  NewCFProductions(),
		Start:        g.Start,
	}

	for A := range g.Productions.AllByHead() {
		for _, p := range g.Productions.Order(A) {
			// Skip ε-production, single productions and productions already in CNF (A → BC or A → a).
			if isBinary, isTerminal := p.IsCNF(); isTerminal || isBinary || p.IsEmpty() || p.IsSingle() {
				newG.Productions.Add(p)
				continue
			}

			/*
			 * Replace each production
			 *
			 *   A → X₁ X₂ ... Xₙ
			 *
			 * with more than two non-terminals X₁, X₂, ..., Xₙ by rules
			 *
			 *   A → X₁ A₁
			 *   A₁ → X₂ A₂
			 *   ...
			 *   Aₙ₋₂ → Xₙ₋₁ Xₙ
			 */

			for head, i := A, 0; i <= len(p.Body)-2; i++ {
				if i == len(p.Body)-2 {
					newG.Productions.Add(CFProduction{head, p.Body[i:]})
				} else {
					headN := newG.addNewNonTerminal(A, numericSuffixes...)
					newG.Productions.Add(CFProduction{head, p.Body[i : i+1].Append(headN)})
					head = headN
				}
			}
		}
	}

	return newG
}

// String returns a string representation of a context-free grammar.
func (g CFG) String() string {
	var b bytes.Buffer

	terms := g.orderTerminals()
	visited, unvisited, nonTerms := g.orderNonTerminals()

	fmt.Fprintf(&b, "Terminal Symbols: %s\n", terms)
	fmt.Fprintf(&b, "Non-Terminal Symbols: %s\n", nonTerms)
	fmt.Fprintf(&b, "Start Symbol: %s\n", g.Start)
	fmt.Fprintln(&b, "Production Rules:")

	for _, head := range visited {
		fmt.Fprintf(&b, "  %s → ", head)
		for _, p := range g.Productions.Order(head) {
			fmt.Fprintf(&b, "%s | ", p.Body.String())
		}
		b.Truncate(b.Len() - 3)
		fmt.Fprintln(&b)
	}

	for _, head := range unvisited {
		fmt.Fprintf(&b, "  %s → ", head)
		for _, p := range g.Productions.Order(head) {
			fmt.Fprintf(&b, "%s | ", p.Body.String())
		}
		b.Truncate(b.Len() - 3)
		fmt.Fprintln(&b)
	}

	return b.String()
}

// addNewNonTerminal generates and adds a new non-terminal symbol to the grammar.
// It does so by appending each of the provided suffixes to the given prefix, in order,
// until it finds a non-terminal that does not already exist in the set of non-terminals.
//
// If all generated non-terminals already exist, the function panics.
func (g CFG) addNewNonTerminal(prefix NonTerminal, suffixes ...string) NonTerminal {
	// Use the base prefix without any previosuly applied suffix.
	for _, suffix := range suffixes {
		prefix = NonTerminal(strings.TrimSuffix(string(prefix), suffix))
	}

	for _, suffix := range suffixes {
		nonTerm := NonTerminal(string(prefix) + suffix)
		if !g.NonTerminals.Contains(nonTerm) {
			g.NonTerminals.Add(nonTerm)
			return nonTerm
		}
	}

	panic(fmt.Sprintf("Failed to generate a new non-terminal for %s", prefix))
}

// orderTerminals orders the unordered set of grammar terminals in a deterministic way.
//
// The goal of this function is to ensure a consistent and deterministic order for any given set of terminals.
func (g CFG) orderTerminals() String[Terminal] {
	terms := make(String[Terminal], 0)
	for t := range g.Terminals.All() {
		terms = append(terms, t)
	}

	// Sort terminals alphabetically based on the string representation of them.
	sort.Quick[Terminal](terms, cmpTerminal)

	return terms
}

// orderTerminals orders the unordered set of grammar non-terminals in a deterministic way.
//
// The goal of this function is to ensure a consistent and deterministic order for any given set of non-terminals.
func (g CFG) orderNonTerminals() (String[NonTerminal], String[NonTerminal], String[NonTerminal]) {
	visited := make(String[NonTerminal], 0)
	isVisited := func(n NonTerminal) bool {
		for _, v := range visited {
			if v == n {
				return true
			}
		}
		return false
	}

	visited = append(visited, g.Start)

	// Reppeat until no new non-terminal is added to visited:
	//   For each production rule of the form A → α:
	//     If A is in visited, add all non-terminal in α to visited.
	for updated := true; updated; {
		updated = false
		for head := range g.Productions.AllByHead() {
			for _, p := range g.Productions.Order(head) {
				if isVisited(p.Head) {
					for _, n := range p.Body.NonTerminals() {
						if !isVisited(n) {
							visited = append(visited, n)
							updated = true
						}
					}
				}
			}
		}
	}

	// Identify any unvisited non-terminals in the grammar.
	unvisited := make(String[NonTerminal], 0)
	for n := range g.NonTerminals.All() {
		if !isVisited(n) {
			unvisited = append(unvisited, n)
		}
	}

	// Sort unvisited non-terminals alphabetically based on the string representation of them.
	sort.Quick[NonTerminal](unvisited, cmpNonTerminal)

	allNonTerms := make(String[NonTerminal], 0)
	allNonTerms = append(allNonTerms, visited...)
	allNonTerms = append(allNonTerms, unvisited...)

	return visited, unvisited, allNonTerms
}
