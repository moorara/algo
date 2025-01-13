package grammar

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
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

// newTerminalsAndEmpty creates a new TerminalsAndEmpty instance with the given set of terminals.
func newTerminalsAndEmpty(terms ...Terminal) *TerminalsAndEmpty {
	return &TerminalsAndEmpty{
		Terminals:     set.New(eqTerminal, terms...),
		IncludesEmpty: false,
	}
}

// String returns a string representation of the FIRST set.
func (s TerminalsAndEmpty) String() string {
	members := []string{}

	for term := range s.Terminals.All() {
		members = append(members, term.String())
	}

	if s.IncludesEmpty {
		members = append(members, "ε")
	}

	sort.Quick(members, generic.NewCompareFunc[string]())

	return fmt.Sprintf("{%s}", strings.Join(members, ", "))
}

// firstBySymbolTable is the type for a table that stores the FIRST set for each grammar symbol.
type firstBySymbolTable symboltable.SymbolTable[Symbol, *TerminalsAndEmpty]

func newFirstBySymbolTable() firstBySymbolTable {
	return symboltable.NewQuadraticHashTable(hashSymbol, eqSymbol, eqTerminalsAndEmpty, symboltable.HashOpts{})
}

// firstByStringTable is the type for a table that stores the FIRST set for strings of grammar symbols.
type firstByStringTable symboltable.SymbolTable[String[Symbol], *TerminalsAndEmpty]

func newFirstByStringTable() firstByStringTable {
	return symboltable.NewQuadraticHashTable(hashString, eqString, eqTerminalsAndEmpty, symboltable.HashOpts{})
}
