// Package grammar provides data types and algorithms for building parsers.
//
// A parser generally relies on a lexer (also known as a lexical analyzer or scanner)
// to process a stream of tokens.
// Lexical analysis (scanning) is distinct from syntax analysis (parsing):
// lexical analysis deals with regular languages and grammars (Type 3),
// while syntax analysis deals with context-free languages and grammars (Type 2).
package parser

import (
	"fmt"
	"strings"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
)

// Action is a function that gets called whenever a production
// rule is selected from the parsing table for an input token.
// It performs the necessary actions associated with the production rule
// and the corresponding lexical token during the predictive top-down parsing.
type Action func(grammar.Production, lexer.Token)

// Parser defines the interface for a syntax analyzer.
type Parser interface {
	// Parse analyzes input tokens (terminal symbols) provided by a lexical analyzer
	// and attempts to construct a syntactic representation (parse tree).
	//
	// The Parse method invokes the given function for each production and token during parsing.
	// It returns an error if the input fails to conform to the grammar rules.
	Parse(Action) error
}

// ParseError represents an error encountered when parsing an input string.
type ParseError struct {
	Description string
	Cause       error
	Pos         lexer.Position
}

// Error implements the error interface.
// It returns a formatted string describing the error in detail.
func (e *ParseError) Error() string {
	b := new(strings.Builder)

	if !e.Pos.IsZero() {
		fmt.Fprintf(b, "%s", e.Pos)
	}

	if len(e.Description) != 0 {
		if b.Len() > 0 {
			fmt.Fprint(b, ": ")
		}
		fmt.Fprintf(b, "%s", e.Description)
	}

	if e.Cause != nil {
		if b.Len() > 0 {
			fmt.Fprint(b, ": ")
		}
		fmt.Fprintf(b, "%s", e.Cause)
	}

	return b.String()
}

// Error implements the unwrap interface.
func (e *ParseError) Unwrap() error {
	return e.Cause
}
