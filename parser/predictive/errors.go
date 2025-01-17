package predictive

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/set"
)

// ParseError represents an error encountered when parsing an input string.
type ParseError struct {
	description string
	cause       error
	Pos         lexer.Position
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *ParseError) Error() string {
	b := new(strings.Builder)

	if !e.Pos.IsZero() {
		fmt.Fprintf(b, "%s", e.Pos)
	}

	if len(e.description) != 0 {
		if b.Len() > 0 {
			fmt.Fprint(b, ": ")
		}
		fmt.Fprintf(b, "%s", e.description)
	}

	if e.cause != nil {
		if b.Len() > 0 {
			fmt.Fprint(b, ": ")
		}
		fmt.Fprintf(b, "%s", e.cause)
	}

	return b.String()
}

// Error implements the unwrap interface.
func (e *ParseError) Unwrap() error {
	return e.cause
}

// parsingTableError represents an error encountered when constructing a predictive parsing table.
// This error occurs due to the presence of left recursion or ambiguity in the grammar.
type parsingTableError struct {
	NonTerminal grammar.NonTerminal
	Terminal    grammar.Terminal
	Productions set.Set[grammar.Production]
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *parsingTableError) Error() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "multiple productions in parsing table at M[%s, %s]:\n", e.NonTerminal, e.Terminal)
	for _, p := range grammar.OrderProductionSet(e.Productions) {
		fmt.Fprintf(b, "  %s\n", p)
	}

	return b.String()
}
