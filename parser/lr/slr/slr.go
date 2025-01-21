// Package slr provides data structures and algorithms for building Simple LR (SLR) parsers.
// An SLR parser is a bottom-up parser for the class of LR(1) grammars.
//
// An SLR parser uses the canonical LR(0) items to construct the state machine (DFA).
// An SLR parser is less powerful than a canonical LR(1) parser.
// SLR simplifies the construction process but sacrifices some parsing power compared to canonical LR(1).
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package slr

import (
	"errors"
	"fmt"
	"io"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/parser"
	"github.com/moorara/algo/parser/lr"
)

// slrParser is an SLR parser for LR(0) grammars.
// It implements the parser.Parser interface.
type slrParser struct {
	G     grammar.CFG
	lexer lexer.Lexer
}

// New creates a new SLR parser for a given context-free grammar (CFG).
// It requires a lexer for lexical analysis, which reads the input tokens (terminal symbols).
func New(G grammar.CFG, lexer lexer.Lexer) parser.Parser {
	return &slrParser{
		G:     G,
		lexer: lexer,
	}
}

// nextToken wraps the Lexer.NextToken method and ensures
// an Endmarker token is returned when the end of input is reached.
func (p *slrParser) nextToken() (lexer.Token, error) {
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
// The Parse method invokes the provided function each time a production rule is successfully matched.
// This allows the caller to process or react to each step of the parsing process.
//
// It returns an error if the input fails to conform to the grammar rules, indicating a syntax error.
func (p *slrParser) Parse(process parser.ProcessFunc) error {

	/*
	 * INPUT:  • A lexer for reading input string w.
	 *         • An LR parsing table with functions ACTION and GOTO for a grammar G.
	 * OUTPUT: • If w ∈ L(G), the reduction steps of a bottom-up parse for w; otherwise, an error indication.
	 *
	 * METHOD: Initially, the parser has s₀ on its stack,
	 *         where s₀ is the initial state, and w$ in the input buffer.
	 *
	 *         let a be the first symbol of w$;
	 *         while (true) {
	 *           let s be the state on top of the stack;
	 *           if (ACTION[s,a] = shift t) {
	 *             push t onto the stack;
	 *             let a be the next input symbol;
	 *           } else if (ACTION[s,a] = reduce A → β) {
	 *             pop |β| symbols off the stack;
	 *             let state t now be on top of the stack;
	 *             push GOTO[t,A] onto the stack;
	 *             output the production A → β;
	 *           } else if (ACTION[s,a] = accept) {
	 *             break;
	 *           } else {
	 *             call error-recovery routine;
	 *           }
	 *         }
	 */

	T := BuildParsingTable(p.G)
	if err := T.Error(); err != nil {
		return &parser.ParseError{
			Description: "failed to construct the SLR parsing table",
			Cause:       err,
		}
	}

	stack := list.NewStack[lr.State](1024, lr.EqState)

	// The BuildStateMap function ensures that state 0 always includes the initial item "S′ → •S".
	stack.Push(lr.State(0))

	// Read the first input token.
	token, err := p.lexer.NextToken()
	if err != nil {
		return &parser.ParseError{Cause: err}
	}

	for {
		s, _ := stack.Peek()
		a := token.Terminal

		action, err := T.ACTION(s, a)
		if err != nil {
			return &parser.ParseError{Cause: err}
		}

		if action.Type == lr.SHIFT {
			stack.Push(action.State)

			// Read the next input token.
			token, err = p.lexer.NextToken()
			if err != nil {
				return &parser.ParseError{Cause: err}
			}
		} else if action.Type == lr.REDUCE {
			A, β := action.Production.Head, action.Production.Body

			for range len(β) {
				stack.Pop()
			}

			t, _ := stack.Peek()

			next, err := T.GOTO(t, A)
			if err != nil {
				return &parser.ParseError{Cause: err}
			}

			stack.Push(next)

			process(*action.Production)
		} else if action.Type == lr.ACCEPT {
			break
		} else {
			// TODO:
			fmt.Println("ERROR RECOVERY!")
		}
	}

	// Accept the input string.
	return nil
}

// ParseAST analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// constructing an abstract syntax tree (AST) that reflects the structure of the input.
//
// If the input string is valid, the root node of the AST is returned,
// representing the syntactic structure of the input string.
//
// It returns an error if the input fails to conform to the grammar rules, indicating a syntax error.
func (p *slrParser) ParseAST() (parser.Node, error) {
	return nil, nil
}
