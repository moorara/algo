// Package predictive provides data structures and algorithms for building predictive parsers.
//
// A predictive parser is a recursive-descent parser without backtracking.
// It is a top-down parser, meaning it constructs the parse tree
// starting from the start symbol and works down to the input symbols.
// Predictive parsers can be constructed for a class of grammars called LL(1).
// The first "L" in LL(1) stands for scanning the input from left to right,
// the second "L" for producing a leftmost derivation,
// and the "1" for using one input symbol of lookahead at each step.
// The class of LL(1) grammars is expressive enough to cover most programming constructs.
package predictive

import (
	"errors"
	"fmt"
	"io"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/parser"
)

// predictiveParser is a predictive parser for LL(1) grammars.
// It implements the parser.Parser interface.
type predictiveParser struct {
	G     grammar.CFG
	lexer lexer.Lexer
}

// New creates a new predictive parser for a given context-free grammar (CFG).
// It requires a lexer for lexical analysis, which reads the input tokens (terminal symbols).
func New(G grammar.CFG, lexer lexer.Lexer) parser.Parser {
	return &predictiveParser{
		G:     G,
		lexer: lexer,
	}
}

// Parse analyzes input tokens (terminal symbols) provided by the lexical analyzer
// and attempts to construct a syntactic representation (parse tree).
//
// The Parse method invokes the given function for each production and token during parsing.
// It returns an error if the input fails to conform to the grammar rules.
func (p *predictiveParser) Parse(yield parser.Action) error {
	/*
	 * INPUT:  • A lexer for reading string ω.
	 *         • A parsing table M for grammar G.
	 * OUTPUT: • If ω ∈ L(G), a leftmost derivation of ω; otherwise an error indication.
	 *
	 * METHOD: Initially, the parser is in a configuration with ω$ in the input buffer
	 *         and the start symbol S of G on top of the stack, above $.
	 *
	 *         let a be the first symbol of ω
	 *         let X be the top stack symbol
	 *         while (X != $) { // stack is not empty
	 *           if (X = a) {
	 *             pop the stack
	 *             let a be the next symbol of ω
	 *           } else if (X is a terminal) {
	 *             error()
	 *           } else if (M[X,a] is an error entry) {
	 *             error()
	 *           } else if (M[X,a] = X → Y₁Y₂...Yₖ) {
	 *             output the production X → Y₁Y₂...Yₖ
	 *             pop the stack
	 *             push Yₖ, Yₖ₋₁, ..., Y₁ onto the stack, with Y₁ on top
	 *           }
	 *           let X be the top stack symbol
	 *         }
	 */

	stack := list.NewStack[grammar.Symbol](1024, grammar.EqSymbol)

	M := BuildParsingTable(p.G)
	if err := M.Error(); err != nil {
		return &ParseError{
			description: "failed to construct the parsing table",
			cause:       err,
			Table:       M,
		}
	}

	stack.Push(grammar.Endmarker)
	stack.Push(p.G.Start)

	// Read the first input token.
	token, err := p.lexer.NextToken()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return &ParseError{cause: err}
	}

	for X, _ := stack.Peek(); !X.Equals(grammar.Endmarker); X, _ = stack.Peek() {
		if X.Equals(token.Terminal) {
			// Pop X from the stack.
			stack.Pop()

			// Read the next input token.
			token, err = p.lexer.NextToken()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return &ParseError{cause: err}
			}

			continue
		}

		if X.IsTerminal() {
			return &ParseError{
				description: fmt.Sprintf("unexpected terminal %s on stack", X),
				Table:       M,
			}
		}

		A := X.(grammar.NonTerminal)
		prods := M.Get(A, token.Terminal)

		if prods.Size() == 0 {
			return &ParseError{
				description: fmt.Sprintf("unacceptable input <%s, %s> for non-terminal %s", token.Terminal, token.Lexeme, A),
				Pos:         token.Pos,
				Table:       M,
			}
		}

		P := generic.Collect1(prods.All())[0]

		// Pop X from the stack.
		stack.Pop()

		// Pushes the symbols of the production body onto the stack in reverse order.
		for i := len(P.Body) - 1; i >= 0; i-- {
			stack.Push(P.Body[i])
		}

		yield(P, token)
	}

	return nil
}
