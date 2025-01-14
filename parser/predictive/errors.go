package predictive

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/set"
)

// ParsingTableError represents an error encountered when constructing a predictive parsing table.
// This error occurs due to the presence of left recursion or ambiguity in the grammar.
type ParsingTableError struct {
	NonTerminal grammar.NonTerminal
	Terminal    grammar.Terminal
	Productions set.Set[grammar.Production]
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *ParsingTableError) Error() string {
	b := new(strings.Builder)

	fmt.Fprintf(b, "multiple productions in parsing table at M[%s, %s]:\n", e.NonTerminal, e.Terminal)
	for _, p := range grammar.OrderProductionSet(e.Productions) {
		fmt.Fprintf(b, "  %s\n", p)
	}

	return b.String()
}
