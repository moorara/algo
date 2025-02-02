// Package simple provides data structures and algorithms for building Simple LR (SLR) parsers.
// An SLR parser is a bottom-up parser for the class of LR(1) grammars.
//
// An SLR parser uses the canonical LR(0) items to construct the state machine (DFA).
// An SLR parser is less powerful than a canonical LR(1) parser.
// SLR simplifies the construction process but sacrifices some parsing power compared to canonical LR(1).
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package simple

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/lr"
)

// New creates a new SLR parser for a given context-free grammar (CFG).
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
