// Package grammar provides data types and algorithms for building parsers.
//
// A parser generally relies on a lexer (also known as a lexical analyzer or scanner)
// to process a stream of tokens.
// Lexical analysis (scanning) is distinct from syntax analysis (parsing):
// lexical analysis deals with regular languages and grammars (Type 3),
// while syntax analysis deals with context-free languages and grammars (Type 2).
package parser

// Parser defines the interface for a syntax analyzer.
type Parser interface {
	// Parse analyzes input tokens (terminal symbols) provided by a lexical analyzer
	// and attempts to construct a syntactic representation (parse tree).
	// It returns an error if the input violates the grammar rules.
	Parse() error
}
