// Package lookahead provides data structures and algorithms for building Look-Ahead LR (LALR) parsers.
// An LALR parser is a bottom-up parser for the class of LR(1) grammars.
//
// An LALR parser, similar to SLR, uses the canonical LR(0) items to construct the state machine (DFA),
// but it refines the states by incorporating lookahead symbols explicitly.
// LALR merges states with identical core LR(0) items but handles lookahead symbols for each merged state separately,
// making it more precise than SLR and avoids many conflicts that SLR might encounter.
// LALR is more powerful than SLR as it can handle a wider range of grammars, including most programming languages.
// However, it is less powerful than canonical LR because state merging can lose distinctions in lookahead contexts,
// potentially leading to conflicts for some grammars.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package lookahead

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/lr"
)

// New creates a new LALR parser for a given context-free grammar (CFG).
// It requires a lexer for lexical analysis, which reads the input tokens (terminal symbols).
func New(L lexer.Lexer, G *grammar.CFG) (*lr.Parser, error) {
	T, err := BuildParsingTable(G)
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
