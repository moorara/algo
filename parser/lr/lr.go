// Package lr provides common data structures and algorithms for building LR parsers.
// LR parsers are bottom-up parsers that analyse deterministic context-free languages in linear time.
//
// Bottom-up parsing constructs a parse tree for an input string
// starting at the leaves (bottom) and working towards the root (top).
// This process involves reducing a string w to the start symbol of the grammar.
// At each reduction step, a specific substring matching the body of a production
// is replaced by the non-terminal at the head of that production.
//
// Bottom-up parsing during a left-to-right scan of the input constructs a rightmost derivation in reverse:
//
//	S = γ₀ ⇒ᵣₘ γ₁ ⇒ᵣₘ γ₂ ⇒ᵣₘ ... ⇒ᵣₘ γₙ₋₁ ⇒ᵣₘ γₙ = w
//
// At each step, the handle βₙ in γₙ is replaced by the head of the production Aₙ → βₙ
// to obtain the previous right-sentential form γₙ₋₁.
// If the process produces the start symbol S as the only sentential form, parsing is complete.
// If a grammar is unambiguous, then every right-sentential form of the grammar has exactly one handle.
//
// The most common type of bottom-up parser is LR(k) parsing.
// The L is for left-to-right scanning of the input, the R for constructing a rightmost derivation in reverse,
// and the k for the number of input symbols of lookahead that are used in making parsing decisions.
//
// Advantages of LR parsing:
//
//   - Can recognize nearly all programming language constructs defined by context-free grammars.
//   - Detects syntax errors at the earliest possible point during a left-to-right scan.
//   - The class of grammars that can be parsed using LR methods is a proper superset of
//     the class of grammars that can be parsed with predictive or LL methods.
//     For a grammar to be LR(k), we must be able to recognize the occurrence of the right side of
//     a production in a right-sentential form, with k input symbols of lookahead.
//     This requirement is far less stringent than that for LL(k) grammars where we must be able
//     to recognize the use of a production seeing only the first k symbols of what its right side derives.
//
// In LR(k) parsing, the cases k = 0 or k = 1 are most commonly used in practical applications.
// LR parsing methods use pushdown automata (PDA) to parse an input string.
// A pushdown automaton is a type of automaton used for Type 2 languages (context-free languages) in the Chomsky hierarchy.
// A PDA uses a state machine with a stack.
// The next state is determined by the current state, the next input, and the top of the stack.
// LR(0) parsers do not rely on any lookahead to make parsing decisions.
// An LR(0) parser bases its decisions entirely on the current state and the parsing stack.
// LR(1) parsers determine the next state based on the current state, one lookahead symbol, and the top of the stack.
//
// Shift-reduce parsing is a bottom-up parsing technique that uses
// a stack for grammar symbols and an input buffer for the remaining string.
// The parser alternates between shifting symbols from the input to the stack
// and reducing the top of the stack based on grammar rules.
// This process continues until the stack contains only the start symbol and the input is empty, or an error occurs.
//
// Certain context-free grammars cannot be parsed using shift-reduce parsers
// because they may encounter shift/reduce conflicts (indecision between shifting or reducing)
// or reduce/reduce conflicts (indecision between multiple reductions).
// Technically speaking, these grammars are not in the LR(k) class.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package lr

import (
	"errors"
	"fmt"
	"io"

	"github.com/moorara/algo/grammar"
	"github.com/moorara/algo/lexer"
	"github.com/moorara/algo/list"
	"github.com/moorara/algo/parser"
)

// Parser is a general LR parser for LR(1) grammars.
// It implements the parser.Parser interface.
type Parser struct {
	L lexer.Lexer
	T *ParsingTable
}

// nextToken wraps the Lexer.NextToken method and ensures
// an Endmarker token is returned when the end of input is reached.
func (p *Parser) nextToken() (lexer.Token, error) {
	token, err := p.L.NextToken()
	if err != nil && errors.Is(err, io.EOF) {
		token.Terminal, token.Lexeme = grammar.Endmarker, ""
		return token, nil
	}

	return token, err
}

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

// Parse implements the LR parsing algorithm.
// It analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// determining whether the input string belongs to the language defined by the grammar.
//
// This method requires a parsing table, which must be generated from a grammar
// by an LR parser (e.g., Simple LR, Canonical LR, or LALR).
//
// The Parse method invokes the provided functions each time a token or a production rule is matched.
// This allows the caller to process or react to each step of the parsing process.
//
// An error is returned if the input fails to conform to the grammar rules, indicating a syntax issue,
// or if any of the provided functions return an error, indicating a semantic issue.
func (p *Parser) Parse(tokenF parser.TokenFunc, prodF parser.ProductionFunc) error {
	stack := list.NewStack[State](1024, EqState)
	stack.Push(State(0)) // BuildStateMap ensures state 0 always includes the initial item "S′ → •S"

	// Read the first input token.
	token, err := p.nextToken()
	if err != nil {
		return &parser.ParseError{Cause: err}
	}

	for {
		s, _ := stack.Peek()
		a := token.Terminal

		action, err := p.T.ACTION(s, a)
		if err != nil {
			return &parser.ParseError{
				Description: fmt.Sprintf("unexpected string %q", token.Lexeme),
				Cause:       err,
				Pos:         token.Pos,
			}
		}

		switch action.Type {
		case SHIFT:
			stack.Push(action.State)

			// Yield the token.
			if tokenF != nil {
				if err := tokenF(&token); err != nil {
					return &parser.ParseError{
						Cause: err,
						Pos:   token.Pos,
					}
				}
			}

			// Read the next input token.
			token, err = p.nextToken()
			if err != nil {
				return &parser.ParseError{Cause: err}
			}

		case REDUCE:
			A, β := action.Production.Head, action.Production.Body

			for range len(β) {
				stack.Pop()
			}

			// An LR parser detects an error when it consults the ACTION table.
			// Errors are never identified by consulting the GOTO table.
			// If ACTION(s, a) is not an error entry, GOTO(t, A) will also not be an error entry.

			t, _ := stack.Peek()
			next, _ := p.T.GOTO(t, A)
			stack.Push(next)

			// Yield the production.
			if prodF != nil {
				if err := prodF(action.Production); err != nil {
					return &parser.ParseError{Cause: err}
				}
			}

		case ACCEPT:
			// Accept the input string.
			return nil

		case ERROR:
			// TODO: This is unreachable currently, since T.ACTION handles the error.
		}
	}
}

// ParseAndBuildAST implements the LR parsing algorithm.
// It analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// constructing an abstract syntax tree (AST) that reflects the structure of the input.
//
// If the input string is valid, the root node of the AST is returned,
// representing the syntactic structure of the input string.
//
// An error is returned if the input fails to conform to the grammar rules, indicating a syntax issue.
func (p *Parser) ParseAndBuildAST() (parser.Node, error) {
	// Stack for constructing the abstract syntax tree.
	nodes := list.NewStack[parser.Node](1024, parser.EqNode)

	err := p.Parse(
		func(token *lexer.Token) error {
			nodes.Push(&parser.LeafNode{
				Terminal: token.Terminal,
				Lexeme:   token.Lexeme,
				Position: token.Pos,
			})

			return nil
		},
		func(prod *grammar.Production) error {
			in := &parser.InternalNode{
				NonTerminal: prod.Head,
				Production:  prod,
			}

			for range len(prod.Body) {
				child, _ := nodes.Pop()
				in.Children = append([]parser.Node{child}, in.Children...) // Maintain correct production body order
			}

			nodes.Push(in)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	// The nodes stack only contains the root of AST at this point.
	root, _ := nodes.Pop()

	return root, nil
}

// ParseAndEvaluate implements the LR parsing algorithm.
// It analyzes a sequence of input tokens (terminal symbols) provided by a lexical analyzer.
// It attempts to parse the input according to the production rules of a context-free grammar,
// evaluating the input in a bottom-up fashion using a rightmost derivation.
//
// During the parsing process, the provided EvaluateFunc is invoked each time a production rule is matched.
// The function is called with values corresponding to the symbols in the body of the production,
// enabling the caller to process and evaluate the input incrementally.
//
// An error is returned if the input fails to conform to the grammar rules, indicating a syntax issue,
// or if the evaluation function returns an error, indicating a semantic issue.
func (p *Parser) ParseAndEvaluate(eval EvaluateFunc) (*Value, error) {
	// Stack for constructing the abstract syntax tree.
	nodes := list.NewStack[*Value](1024, nil)

	err := p.Parse(
		func(token *lexer.Token) error {
			copy := token.Pos
			nodes.Push(&Value{
				Val: token.Lexeme,
				Pos: &copy,
			})

			return nil
		},
		func(prod *grammar.Production) error {
			l := len(prod.Body)
			rhs := make([]*Value, l)

			// Maintain correct production body order
			for i := l - 1; i >= 0; i-- {
				v, _ := nodes.Pop()
				rhs[i] = v
			}

			lhs, err := eval(prod, rhs)
			if err != nil {
				return err
			}

			v := &Value{Val: lhs}
			if l > 0 {
				v.Pos = rhs[0].Pos
			}

			nodes.Push(v)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	// The nodes stack only contains the root of AST at this point.
	root, _ := nodes.Pop()

	return root, nil
}

// EvaluateFunc is a function invoked every time a production rule
// is matched or applied during the parsing of an input string.
//
// It receives a list of values corresponding to the right-hand side of the matched production
// and expects a value to be returned representing the left-hand side of the production.
//
// The returned value will be subsequently used as an input in the evaluation of other production rules.
// Both the input and output values are of the generic type any.
//
// The caller is responsible for ensuring that each value is converted to the appropriate type based on
// the production rule and the position of the symbol corresponding to the value in the production's right-hand side.
// The input values must retain the same type they were originally evaluated as when returned.
//
// The function may return an error if there are issues with the input values,
// such as mismatched types or unexpected inputs.
type EvaluateFunc func(*grammar.Production, []*Value) (any, error)

// Value represents a value used during the evaluation process,
// along with its corresponding positional information in the input.
type Value struct {
	Val any
	Pos *lexer.Position
}

// String returns a string representation of a value.
func (v *Value) String() string {
	if v.Pos == nil || v.Pos.IsZero() {
		return fmt.Sprintf("%v", v.Val)
	}

	return fmt.Sprintf("%v <%s>", v.Val, v.Pos)
}
