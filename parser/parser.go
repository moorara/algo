// Package grammar provides data types and algorithms for building parsers.
//
// A parser generally relies on a lexer (also known as a lexical analyzer or scanner)
// to process a stream of tokens.
// Lexical analysis (scanning) is distinct from syntax analysis (parsing):
// lexical analysis deals with regular languages and grammars (Type 3),
// while syntax analysis deals with context-free languages and grammars (Type 2).
package parser

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/lexer"
)

// Parser defines the interface for a syntax analyzer that processes input tokens.
type Parser interface {
	// Parse analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
	// It attempts to parse the input according to the production rules of a context-free grammar,
	// determining whether the input string belongs to the language defined by the grammar.
	//
	// The Parse method invokes the provided functions each time a token or a production rule is successfully matched.
	// This allows the caller to process or react to each step of the parsing process.
	//
	// It returns an error if the input fails to conform to the grammar rules, indicating a syntax error.
	Parse(TokenFunc, ProductionFunc) error

	// ParseAST analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
	// It attempts to parse the input according to the production rules of a context-free grammar,
	// constructing an abstract syntax tree (AST) that reflects the structure of the input.
	//
	// If the input string is valid, the root node of the AST is returned,
	// representing the syntactic structure of the input string.
	//
	// It returns an error if the input fails to conform to the grammar rules, indicating a syntax error.
	ParseAST() (Node, error)
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
	var b bytes.Buffer

	if !e.Pos.IsZero() {
		fmt.Fprintf(&b, "%s", e.Pos)
	}

	if len(e.Description) != 0 {
		if b.Len() > 0 {
			fmt.Fprint(&b, ": ")
		}
		fmt.Fprintf(&b, "%s", e.Description)
	}

	if e.Cause != nil {
		if b.Len() > 0 {
			fmt.Fprint(&b, ": ")
		}
		fmt.Fprintf(&b, "%s", e.Cause)
	}

	return b.String()
}

// Error implements the unwrap interface.
func (e *ParseError) Unwrap() error {
	return e.Cause
}
