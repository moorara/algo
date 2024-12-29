package grammar

import (
	"github.com/moorara/algo/set"
	. "github.com/moorara/algo/symboltable"
)

var eqTerminalsAndEndmarker = func(lhs, rhs *TerminalsAndEndmarker) bool {
	return lhs.Terminals.Equals(rhs.Terminals) && lhs.IncludesEndmarker == rhs.IncludesEndmarker
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

func newTerminalsAndEndmarker(terms ...Terminal) *TerminalsAndEndmarker {
	return &TerminalsAndEndmarker{
		Terminals:         set.New(eqTerminal, terms...),
		IncludesEndmarker: false,
	}
}

// followTable is the type for a table that stores the FOLLOW set for each non-terminal.
type followTable SymbolTable[NonTerminal, *TerminalsAndEndmarker]

func newFollowTable() followTable {
	return NewQuadraticHashTable(hashNonTerminal, eqNonTerminal, eqTerminalsAndEndmarker, HashOpts{})
}
