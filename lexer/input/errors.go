package input

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/lexer"
)

// InputError represents an error encountered when reading from an input source.
type InputError struct {
	description string
	Pos         lexer.Position
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *InputError) Error() string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "%s: %s", e.Pos, e.description)
	return b.String()
}
