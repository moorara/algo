package grammar

import (
	"fmt"
	"strings"
)

// CNFError represents an error for a production rule in the form
// A → α that does not conform to Chomsky Normal Form (CNF).
type CNFError struct {
	// The production rule not in Chomsky Normal Form (CNF).
	P Production
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *CNFError) Error() string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "Production %s is neither a binary rule, a terminal rule, nor S → ε", e.P)
	return b.String()
}

// LL1Error represents an error where two distinct production rules in the form
// A → α | β violate LL(1) parsing requirements for a context-free grammar.
type LL1Error struct {
	// A brief description of the error, explaining the issue.
	Description string
	// The head of both productions.
	A NonTerminal
	// The bodies of productions.
	Alpha, Beta String[Symbol]
	// The FOLLOW set for the head of productions.
	FOLLOWA *TerminalsAndEndmarker
	// The FIRST sets for the bodies of production.
	FIRSTα, FIRSTβ *TerminalsAndEmpty
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *LL1Error) Error() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "%s:\n", e.Description)
	fmt.Fprintf(b, "  %s → %s | %s\n", e.A, e.Alpha, e.Beta)

	if e.FOLLOWA != nil {
		fmt.Fprintf(b, "    FOLLOW(%s): %s\n", e.A, e.FOLLOWA)
	}

	if e.FIRSTα != nil {
		fmt.Fprintf(b, "    FIRST(%s): %s\n", e.Alpha, e.FIRSTα)
	}

	if e.FIRSTβ != nil {
		fmt.Fprintf(b, "    FIRST(%s): %s\n", e.Beta, e.FIRSTβ)
	}

	return b.String()
}
