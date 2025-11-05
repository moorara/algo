package parser

import (
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
)

// TokenFunc is a function that is invoked each time a token
// is matched and removed from an input string during parsing.
//
// It executes the actions associated with the matched token,
// such as semantic processing, constructing abstract syntax trees (AST),
// or performing other custom logic required for the parsing process.
//
// The function may return an error, indicating an issue during token processing.
// The parser may stop immediately or continue parsing and accumulate more errors.
type TokenFunc func(*lexer.Token) error

// ProductionFunc is a function that is invoked each time a production rule
// is matched or applied during the parsing process of an input string.
//
// It executes the actions associated with the matched production rule,
// such as semantic processing, constructing abstract syntax trees (AST),
// or performing other custom logic required for the parsing process.
//
// The function may return an error, indicating an issue during production rule processing.
// The parser may stop immediately or continue parsing and accumulate more errors.
type ProductionFunc func(*grammar.Production) error
