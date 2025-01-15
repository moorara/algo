package grammar

import (
	"fmt"
	"strings"
)

// CNFError represents an error for a production rule in the form
// A → α that does not conform to Chomsky Normal Form (CNF).
type CNFError struct {
	P Production
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *CNFError) Error() string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "production %s is neither a binary rule, a terminal rule, nor S → ε", e.P)
	return b.String()
}

// LL1Error represents an error where two distinct production rules in the form
// A → α | β violate LL(1) parsing requirements for a context-free grammar.
type LL1Error struct {
	description    string
	A              NonTerminal
	Alpha, Beta    String[Symbol]
	FOLLOWA        *TerminalsAndEndmarker
	FIRSTα, FIRSTβ *TerminalsAndEmpty
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *LL1Error) Error() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "%s:\n", e.description)
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
