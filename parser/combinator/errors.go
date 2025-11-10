package combinator

import (
	"errors"
	"fmt"
)

// errEOF is returned when the end of input is reached unexpectedly.
var errEOF = errors.New("end of input")

// syntaxError represents a syntactic error encountered during parsing.
type syntaxError struct {
	Pos  int
	Rune rune
}

// Error implements the error interface.
func (e *syntaxError) Error() string {
	return fmt.Sprintf("%d: unexpected rune %q", e.Pos, e.Rune)
}

// semanticError represents a semantic error encountered during parsing.
type semanticError struct {
	Pos int
	Err error
}

// Error implements the error interface.
func (e *semanticError) Error() string {
	return fmt.Sprintf("%d: %s", e.Pos, e.Err)
}

// Unwrap implements the unwrapper interface.
func (e *semanticError) Unwrap() error {
	return e.Err
}
