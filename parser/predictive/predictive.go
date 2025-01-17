// Package predictive provides data structures and algorithms for building predictive parsers.
// A predictive parser is a top-down recursive-descent parser without backtracking.
//
// Top-down parsing involves constructing a parse tree for the input string,
// starting from the root node (representing the start symbol) and expanding the nodes in preorder.
// Equivalently, top-down parsing can be viewed as finding a leftmost derivation for an input string.
//
// A recursive descent parser is a top-down parser constructed from mutually recursive procedures
// (or their non-recursive equivalents), where each procedure corresponds to a nonterminal in the grammar.
// This structure closely mirrors the grammar, making it intuitive and directly aligned with the rules it recognizes.
//
// Predictive parsers can be constructed for a class of grammars called LL(1).
// The first "L" in LL(1) stands for scanning the input from left to right,
// the second "L" for producing a leftmost derivation,
// and the "1" for using one input symbol of lookahead at each step.
// The class of LL(1) grammars is expressive enough to cover most programming constructs,
// such as arithmetic expressions and simple control structures.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package predictive

import (
	"errors"
	"fmt"
	"io"

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
	 * INPUT:  • A lexer for reading string w.
	 *         • A parsing table M for grammar G.
	 * OUTPUT: • If w ∈ L(G), a leftmost derivation of w; otherwise an error indication.
	 *
	 * METHOD: Initially, the parser is in a configuration with w$ in the input buffer
	 *         and the start symbol S of G on top of the stack, above $.
	 *
	 *         let a be the first symbol of w
	 *         let X be the top stack symbol
	 *         while (X != $) { // stack is not empty
	 *           if (X = a) {
	 *             pop the stack
	 *             let a be the next symbol of w
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

	M := buildParsingTable(p.G)
	if err := M.Error(); err != nil {
		return &ParseError{
			description: "failed to construct the parsing table",
			cause:       err,
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
			}
		}

		A := X.(grammar.NonTerminal)

		if M.IsEmpty(A, token.Terminal) {
			return &ParseError{
				description: fmt.Sprintf("unacceptable input <%s, %s> for non-terminal %s", token.Terminal, token.Lexeme, A),
				Pos:         token.Pos,
			}
		}

		// At this point, it is guaranteed that M[A,a] contains exactly one production.
		P, _ := M.GetProduction(A, token.Terminal)

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
