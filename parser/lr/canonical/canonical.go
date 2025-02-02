// Package canonical provides data structures and algorithms for building Canonical LR parsers.
// A canonical LR parser is a bottom-up parser for the class of LR(1) grammars.
//
// The canonical LR or just LR method makes full use of the lookahead symbols.
// This method uses a large set of items, called the LR(1) items.
// LR method handles a broader class of context-free grammars compared to SLR method.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package canonical

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/lr"
)

// New creates a new Canonical LR parser for a given context-free grammar (CFG).
// It requires a lexer for lexical analysis, which reads the input tokens (terminal symbols).
func New(L lexer.Lexer, G *grammar.CFG, precedences lr.PrecedenceLevels) (*lr.Parser, error) {
	T, err := BuildParsingTable(G, precedences)
	if err != nil {
		return nil, &parser.ParseError{
			Cause: err,
		}
	}

	return &lr.Parser{
		L: L,
		T: T,
	}, nil
}
