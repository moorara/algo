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
	G     *grammar.CFG
	lexer lexer.Lexer
}

// New creates a new predictive parser for a given context-free grammar (CFG).
// It requires a lexer for lexical analysis, which reads the input tokens (terminal symbols).
func New(G *grammar.CFG, lexer lexer.Lexer) parser.Parser {
	return &predictiveParser{
		G:     G,
		lexer: lexer,
	}
}

// nextToken wraps the Lexer.NextToken method and ensures
// an Endmarker token is returned when the end of input is reached.
func (p *predictiveParser) nextToken() (lexer.Token, error) {
	token, err := p.lexer.NextToken()
	if err != nil && errors.Is(err, io.EOF) {
		token.Terminal, token.Lexeme = grammar.Endmarker, ""
		return token, nil
	}

	return token, err
}

// Parse analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// determining whether the input string belongs to the language defined by the grammar.
//
// The Parse method invokes the provided functions each time a token or a production rule is matched.
// This allows the caller to process or react to each step of the parsing process.
//
// An error is returned if the input fails to conform to the grammar rules, indicating a syntax issue,
// or if any of the provided functions return an error, indicating a semantic issue.
func (p *predictiveParser) Parse(tokenF parser.TokenFunc, prodF parser.ProductionFunc) error {
	/*
	 * INPUT:  • A lexer for reading input string w.
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

	M := BuildParsingTable(p.G)
	if err := M.Error(); err != nil {
		return &parser.ParseError{
			Description: "failed to construct the predictive parsing table",
			Cause:       err,
		}
	}

	stack := list.NewStack[grammar.Symbol](1024, grammar.EqSymbol)
	stack.Push(grammar.Endmarker)
	stack.Push(p.G.Start)

	// Read the first input token.
	token, err := p.nextToken()
	if err != nil {
		return &parser.ParseError{Cause: err}
	}

	for X, _ := stack.Peek(); !X.Equal(grammar.Endmarker); X, _ = stack.Peek() {
		if X.Equal(token.Terminal) {
			// Yield the token.
			if tokenF != nil {
				if err := tokenF(&token); err != nil {
					return &parser.ParseError{
						Cause: err,
						Pos:   token.Pos,
					}
				}
			}

			// Pop X from the stack.
			stack.Pop()

			// Read the next input token.
			token, err = p.nextToken()
			if err != nil {
				return &parser.ParseError{Cause: err}
			}

			continue
		}

		if X.IsTerminal() {
			return &parser.ParseError{
				Description: fmt.Sprintf("unexpected terminal %s on stack", X),
			}
		}

		A := X.(grammar.NonTerminal)

		if M.IsEmpty(A, token.Terminal) {
			return &parser.ParseError{
				Description: fmt.Sprintf("unacceptable input <%s, %s> for non-terminal %s", token.Terminal, token.Lexeme, A),
				Pos:         token.Pos,
			}
		}

		// At this point, it is guaranteed that M[A,a] contains exactly one production.
		prod, _ := M.GetProduction(A, token.Terminal)

		// Yield the production.
		if prodF != nil {
			if err := prodF(prod); err != nil {
				return &parser.ParseError{Cause: err}
			}
		}

		// Pop X from the stack.
		stack.Pop()

		// Pushes the symbols of the production body onto the stack in reverse order.
		for i := len(prod.Body) - 1; i >= 0; i-- {
			stack.Push(prod.Body[i])
		}
	}

	// Accept the input string.
	return nil
}

// ParseAndBuildAST analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// constructing an abstract syntax tree (AST) that reflects the structure of the input.
//
// If the input string is valid, the root node of the AST is returned,
// representing the syntactic structure of the input string.
//
// An error is returned if the input fails to conform to the grammar rules, indicating a syntax issue.
func (p *predictiveParser) ParseAndBuildAST() (parser.Node, error) {
	// Root of the abstract syntax tree.
	root := &parser.InternalNode{
		NonTerminal: p.G.Start,
	}

	// Stack for constructing the abstract syntax tree.
	nodes := list.NewStack[parser.Node](1024, parser.EqNode)
	nodes.Push(root)

	err := p.Parse(
		func(token *lexer.Token) error {
			// Complete the leaf node.
			n, _ := nodes.Pop()
			lf, _ := n.(*parser.LeafNode)
			lf.Lexeme = token.Lexeme
			lf.Position = token.Pos

			return nil
		},
		func(prod *grammar.Production) error {
			n, _ := nodes.Pop()
			in, _ := n.(*parser.InternalNode)
			in.Production = prod

			// Push production body nodes onto the stack in reverse order.
			for i := len(prod.Body) - 1; i >= 0; i-- {
				var child parser.Node

				switch Y := prod.Body[i].(type) {
				case grammar.NonTerminal:
					child = &parser.InternalNode{NonTerminal: Y}
				case grammar.Terminal:
					child = &parser.LeafNode{Terminal: Y}
				}

				// Prepend child nodes to maintain correct production body order.
				in.Children = append([]parser.Node{child}, in.Children...)

				nodes.Push(child)
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return root, nil
}
