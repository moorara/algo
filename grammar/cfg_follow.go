package grammar

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/set"
	"github.com/moorara/algo/sort"
	"github.com/moorara/algo/symboltable"
)

var EqTerminalsAndEndmarker = func(lhs, rhs *TerminalsAndEndmarker) bool {
	return lhs.Terminals.Equal(rhs.Terminals) && lhs.IncludesEndmarker == rhs.IncludesEndmarker
}

// FOLLOW is the FIRST function associated with a context-free grammar.
//
// FOLLOW(A), for non-terminal A, is the set of terminals 𝑎
// that can appear immediately to the right of A in some sentential form;
// that is; the set of terminals 𝑎 such that there exists a derivation of the form S ⇒* αAaβ
// for some α and β strings of grammar symbols (terminals and non-terminals).
type FOLLOW func(NonTerminal) TerminalsAndEndmarker

// TerminalsAndEndmarker is the return type for the FOLLOW function.
//
// It contains:
//
//   - A set of terminals that can appear immediately after the given non-terminal.
//   - A flag indicating whether the special endmarker symbol is included in the FOLLOW set.
type TerminalsAndEndmarker struct {
	Terminals         set.Set[Terminal]
	IncludesEndmarker bool
}

// newTerminalsAndEndmarker creates a new TerminalsAndEndmarker instance with the given set of terminals.
func newTerminalsAndEndmarker(terms ...Terminal) *TerminalsAndEndmarker {
	return &TerminalsAndEndmarker{
		Terminals:         set.New(EqTerminal, terms...),
		IncludesEndmarker: false,
	}
}

// String returns a string representation of the FOLLOW set.
func (s *TerminalsAndEndmarker) String() string {
	members := []string{}

	for term := range s.Terminals.All() {
		members = append(members, term.String())
	}

	if s.IncludesEndmarker {
		members = append(members, Endmarker.String())
	}

	sort.Quick(members, generic.NewCompareFunc[string]())

	return fmt.Sprintf("{%s}", strings.Join(members, ", "))
}

// followTable is the type for a table that stores the FOLLOW set for each non-terminal.
type followTable symboltable.SymbolTable[NonTerminal, *TerminalsAndEndmarker]

func newFollowTable() followTable {
	return symboltable.NewQuadraticHashTable(
		HashNonTerminal,
		EqNonTerminal,
		EqTerminalsAndEndmarker,
		symboltable.HashOpts{},
	)
}
