package grammar

import (
	"github.com/moorara/algo/set"
	. "github.com/moorara/algo/symboltable"
)

var eqTerminalsAndEmpty = func(lhs, rhs *TerminalsAndEmpty) bool {
	return lhs.Terminals.Equals(rhs.Terminals) && lhs.IncludesEmpty == rhs.IncludesEmpty
}

// FIRST is the FIRST function associated with a context-free grammar.
//
// FIRST(α), where α is any string of grammar symbols (terminals and non-terminals),
// is the set of terminals that begin strings derived from α.
// If α ⇒* ε, then ε is also in FIRST(α).
type FIRST func(String[Symbol]) TerminalsAndEmpty

// TerminalsAndEmpty is the return type for the FIRST function.
//
// It contains:
//
//   - A set of terminals that may appear at the beginning of strings derived from α.
//   - A flag indicating whether the empty string ε is included in the FIRST set..
type TerminalsAndEmpty struct {
	Terminals     set.Set[Terminal]
	IncludesEmpty bool
}

func newTerminalsAndEmpty(terms ...Terminal) *TerminalsAndEmpty {
	return &TerminalsAndEmpty{
		Terminals:     set.New(eqTerminal, terms...),
		IncludesEmpty: false,
	}
}

// firstBySymbolTable is the type for a table that stores the FIRST set for each grammar symbol.
type firstBySymbolTable SymbolTable[Symbol, *TerminalsAndEmpty]

func newFirstBySymbolTable() firstBySymbolTable {
	return NewQuadraticHashTable(hashSymbol, eqSymbol, eqTerminalsAndEmpty, HashOpts{})
}

// firstByStringTable is the type for a table that stores the FIRST set for strings of grammar symbols.
type firstByStringTable SymbolTable[String[Symbol], *TerminalsAndEmpty]

func newFirstByStringTable() firstByStringTable {
	return NewQuadraticHashTable(hashString, eqString, eqTerminalsAndEmpty, HashOpts{})
}
