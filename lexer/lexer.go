// Package lexer defines abstractions and data types for constructing lexers.
//
// A lexer, also known as a lexical analyzer or scanner, is responsible for tokenizing input source code.
// It processes a stream of characters and converts them into a stream of tokens,
// which represent meaningful units of the language.
// These tokens are subsequently passed to a parser for syntax analysis and the construction of parse trees.
//
// Lexical analysis (scanning) belongs to a different domain than syntax analysis (parsing).
// Lexical analysis deals with regular languages and grammars (Type 3),
// while syntax analysis deals with context-free languages and grammars (Type 2).
// A lexical analyzer is, in principle, a deterministic finite automaton (DFA)
// with additional functionality built on top of it.
// Lexers can be implemented either by hand or auto-generated.
package lexer

import (
	"bytes"
	"fmt"

	"github.com/moorara/algo/grammar"
)

// Lexer defines the interface for a lexical analyzer.
type Lexer interface {
	// NextToken reads characters from the input source and returns the next token.
	// It may also return an error if there is an issue during tokenization.
	NextToken() (Token, error)
}

// Token represents a unit of the input language.
//
// A token consists of a terminal symbol, along with additional information such as
// the lexeme (the actual value of the token in the input) and its position in the input stream.
//
// For example, identifiers in a programming language may have different names,
// but their token type (terminal symbol) is typically "ID".
// Similarly, the token "NUM" can have various lexeme values,
// representing different numerical values in the input.
type Token struct {
	grammar.Terminal
	Lexeme string
	Pos    Position
}

// String implements the fmt.Stringer interface.
//
// It returns a formatted string representation of the token.
func (t Token) String() string {
	return fmt.Sprintf("%s <%s, %s>", t.Terminal, t.Lexeme, t.Pos)
}

// Equals determines whether or not two tokens are the same.
func (t Token) Equals(rhs Token) bool {
	return t.Terminal.Equals(rhs.Terminal) &&
		t.Lexeme == rhs.Lexeme &&
		t.Pos.Equals(rhs.Pos)
}

// Position represents a specific location in an input source.
type Position struct {
	Filename string // The name of the input source file (optional).
	Offset   int    // The byte offset from the beginning of the file.
	Line     int    // The line number (1-based).
	Column   int    // The column number on the line (1-based).
}

// String implements the fmt.Stringer interface.
//
// It returns a formatted string representation of the position.
func (p Position) String() string {
	var b bytes.Buffer

	if len(p.Filename) > 0 {
		fmt.Fprintf(&b, "%s:", p.Filename)
	}

	if p.Line > 0 && p.Column > 0 {
		fmt.Fprintf(&b, "%d:%d", p.Line, p.Column)
	} else {
		fmt.Fprintf(&b, "%d", p.Offset)
	}

	return b.String()
}

// Equals determines whether or not two positions are the same.
func (p Position) Equals(rhs Position) bool {
	return p.Filename == rhs.Filename &&
		p.Offset == rhs.Offset &&
		p.Line == rhs.Line &&
		p.Column == rhs.Column
}

// IsZero checks if a position is a zero (empty) value.
func (p Position) IsZero() bool {
	var zero Position
	return p == zero
}
