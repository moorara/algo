// Package combinator provides data types and primitive constructs for building parser combinators.
// A parser combinator is a higher-order function that takes
// one or more parsers as input and produces a new composite parser as output.
//
// A parser itself is a function that processes an input stream of characters and returns an output structure,
// such as an abstract syntax tree (AST), a finite automaton, or another representation of the parsed data.
//
// Parser combinators enable a modular, top-down recursive descent parsing strategy.
// They allow complex parsers to be constructed from smaller, reusable components,
// making the parsing process easier to build, maintain, and test.
//
// Top-down parsing involves constructing a parse tree for the input string,
// starting from the root node (representing the start symbol) and expanding the nodes in preorder.
// Equivalently, top-down parsing can be viewed as finding a leftmost derivation for an input string.
//
// A recursive descent parser is a top-down parser constructed from mutually recursive procedures
// (or their non-recursive equivalents), where each procedure corresponds to a nonterminal in the grammar.
// This structure closely mirrors the grammar, making it intuitive and directly aligned with the rules it recognizes.
//
// For more details on parsing theory,
// refer to "Compilers: Principles, Techniques, and Tools (2nd Edition)".
package combinator

import "slices"

// Parser is the type for a function that receives a parsing input and returns a parsing output.
// The second return value determines whether or not the parsing was successful and the output is valid.
type Parser func(Input) (Output, bool)

// Input is the input to a parser function.
type Input interface {
	// Current returns the current rune from input along with its position in the input.
	Current() (rune, int)
	// Remaining returns the remaining of input. If no input left, it returns nil.
	Remaining() Input
}

// Output is the output of a parser function.
type Output struct {
	Result    Result
	Remaining Input
}

// List is the type for the result of concatenation or repetition.
type List []Result

// Result is the result of parsing a production rule.
// It represents a production rule result.
type Result struct {
	// Val is the actual result of a parser function.
	// It can be an abstract syntax tree, a finite automata, or any other data structure.
	Val any
	// Pos is the first position in the source corresponding to the parsing result.
	Pos int
	// Bag is an optional collection of key-value pairs holding extra information and metadata about the parsing result.
	// You should always check this field to be not nil before using it.
	Bag Bag
}

type (
	// Bag is the type for a collection of key-value pairs.
	Bag map[BagKey]BagVal

	// BagKey is the type for the keys in Bag type.
	BagKey string

	// BagVal is the type for the values in Bag type.
	BagVal any
)

// Get returns the parsing result of a symbol from the right-side of a production rule.
// If the position i is out of bounds, the second return value is false.
//
// Example:
//
//	// Production Rule: range → "{" num ( "," num? )? "}"
//	r = {2,4}
//	r.Get(1) = 2
//	r.Get(3) = 4
func (r *Result) Get(i int) (Result, bool) {
	if l, ok := r.Val.(List); ok {
		if 0 <= i && i < len(l) {
			return l[i], true
		}
	}

	return Result{}, false
}

// Empty represents the empty string ε.
type Empty struct{}

// E is the empty parser for consuming the empty string ε.
// It always succeeds without consuming any input.
var E Parser = func(in Input) (Output, bool) {
	_, pos := in.Current()

	return Output{
		Result: Result{
			Val: Empty{},
			Pos: pos,
		},
		Remaining: in,
	}, true
}

// ExpectRune creates a parser that returns a successful result only if the input starts with the given rune.
func ExpectRune(r rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil {
			return Output{}, false
		}

		if curr, pos := in.Current(); curr == r {
			return Output{
				Result:    Result{curr, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// NotExpectRune creates a parser that returns a successful result only if the input does not start with the given rune.
func NotExpectRune(r rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil {
			return Output{}, false
		}

		if curr, pos := in.Current(); curr != r {
			return Output{
				Result:    Result{curr, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// ExpectRuneIn creates a parser that returns a successful result only if the input starts with any of the given runes.
func ExpectRuneIn(runes ...rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil {
			return Output{}, false
		}

		if r, pos := in.Current(); slices.Contains(runes, r) {
			return Output{
				Result:    Result{r, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// NotExpectRuneIn creates a parser that returns a successful result only if the input does not start with any of the given runes.
func NotExpectRuneIn(runes ...rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil {
			return Output{}, false
		}

		if r, pos := in.Current(); !slices.Contains(runes, r) {
			return Output{
				Result:    Result{r, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// ExpectRuneInRange creates a parser that returns a successful result only if the input starts with a rune in the given range.
func ExpectRuneInRange(lo, hi rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil || lo > hi {
			return Output{}, false
		}

		if r, pos := in.Current(); lo <= r && r <= hi {
			return Output{
				Result:    Result{r, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// NotExpectRuneInRange creates a parser that returns a successful result only if the input does not start with a rune in the given range.
func NotExpectRuneInRange(lo, hi rune) Parser {
	return func(in Input) (Output, bool) {
		if in == nil || lo > hi {
			return Output{}, false
		}

		if r, pos := in.Current(); r < lo || hi < r {
			return Output{
				Result:    Result{r, pos, nil},
				Remaining: in.Remaining(),
			}, true
		}

		return Output{}, false
	}
}

// ExpectRunes creates a parser that returns a successful result only if the input starts with the given runes in the given order.
func ExpectRunes(runes ...rune) Parser {
	return func(in Input) (Output, bool) {
		var pos int

		for i, r := range runes {
			if in == nil {
				return Output{}, false
			}

			curr, p := in.Current()
			if curr != r {
				return Output{}, false
			}

			// Save only the first position
			if i == 0 {
				pos = p
			}

			in = in.Remaining()
		}

		return Output{
			Result:    Result{runes, pos, nil},
			Remaining: in,
		}, true
	}
}

// NotExpectRunes creates a parser that returns a successful result only if the input does not start with the given runes in the given order.
func NotExpectRunes(runes ...rune) Parser {
	return func(in Input) (Output, bool) {
		var pos int
		val := make([]rune, len(runes))

		for i, r := range runes {
			if in == nil {
				return Output{}, false
			}

			curr, p := in.Current()
			if curr == r {
				return Output{}, false
			}

			// Accumulate the parsed runes
			val[i] = curr

			// Save only the first position
			if i == 0 {
				pos = p
			}

			in = in.Remaining()
		}

		return Output{
			Result:    Result{val, pos, nil},
			Remaining: in,
		}, true
	}
}

// ExpectString creates a parser that returns a successful result only if the input starts with the given string.
func ExpectString(s string) Parser {
	return func(in Input) (Output, bool) {
		if out, ok := ExpectRunes([]rune(s)...)(in); ok {
			out.Result.Val = s
			return out, true
		}

		return Output{}, false
	}
}

// NotExpectString creates a parser that returns a successful result only if the input does not start with the given string.
func NotExpectString(s string) Parser {
	return func(in Input) (Output, bool) {
		if out, ok := NotExpectRunes([]rune(s)...)(in); ok {
			out.Result.Val = string(out.Result.Val.([]rune))
			return out, true
		}

		return Output{}, false
	}
}

// ALT composes a parser that alternates a sequence of parsers.
// It applies the first parser to the input and if it does not succeed,
// it applies the next parser to the same input, and continues parsing to the last parser.
// It stops at the first successful parsing and returns its result.
//
//   - EBNF Operator: Alternation
//   - EBNF Notation: p | q
func ALT(p ...Parser) Parser {
	return func(in Input) (Output, bool) {
		for _, parse := range p {
			if out, ok := parse(in); ok {
				return out, true
			}
		}

		return Output{}, false
	}
}

// CONCAT composes a parser that concats a sequence of parsers.
// It applies the first parser to the input, then applies the next parser to the remaining of the input,
// and continues parsing to the last parser.
//
//   - EBNF Operator: Concatenation
//   - EBNF Notation: p q
func CONCAT(p ...Parser) Parser {
	return func(in Input) (Output, bool) {
		var l List

		for _, parse := range p {
			out, ok := parse(in)
			if !ok {
				return Output{}, false
			}

			l = append(l, out.Result)
			in = out.Remaining
		}

		return Output{
			Result:    Result{l, l[0].Pos, nil},
			Remaining: in,
		}, true
	}
}

// OPT composes a parser that applies parser p zero or one time to the input.
// If the parser does not succeed, it will return an empty result.
//
//   - EBNF Operator: Optional
//   - EBNF Notation: [ p ] or p?
func OPT(p Parser) Parser {
	return func(in Input) (Output, bool) {
		if out, ok := p(in); ok {
			return out, true
		}

		return Output{
			Result: Result{
				Val: Empty{},
			},
			Remaining: in,
		}, true
	}
}

// REP composes a parser that applies parser p zero or more times to the input and accumulates the results.
// If the parser does not succeed, it will return an empty result.
//
//   - EBNF Operator: Repetition (Kleene Star)
//   - EBNF Notation: { p } or p*
func REP(p Parser) Parser {
	return func(in Input) (Output, bool) {
		var l List

		for i := 0; in != nil; i++ {
			out, ok := p(in)
			if !ok {
				break
			}

			l = append(l, out.Result)
			in = out.Remaining
		}

		out := Output{
			Remaining: in,
		}

		if len(l) == 0 {
			out.Result = Result{
				Val: Empty{},
			}
		} else {
			out.Result = Result{l, l[0].Pos, nil}
		}

		return out, true
	}
}

// REP1 composes a parser that applies parser p one or more times to the input and accumulates the results.
// This does not allow parsing zero times (empty result).
//
//   - EBNF Operator: Kleene Plus
//   - EBNF Notation: p+
func REP1(p Parser) Parser {
	return func(in Input) (Output, bool) {
		if out, ok := p.REP()(in); ok {
			if res, ok := out.Result.Val.(List); ok && len(res) > 0 {
				return out, true
			}
		}

		return Output{}, false
	}
}

// ALT composes a parser that alternates parser p with a sequence of parsers.
// It applies parser p to the input and if it does not succeed,
// it applies the next parser to the same input, and continues parsing to the last parser.
// It stops at the first successful parsing and returns its result.
//
//   - EBNF Operator: Alternation
//   - EBNF Notation: p | q
func (p Parser) ALT(q ...Parser) Parser {
	all := append([]Parser{p}, q...)
	return ALT(all...)
}

// CONCAT composes a parser that concats parser p to a sequence of parsers.
// It applies parser p to the input, then applies the next parser to the remaining of the input,
// and continues parsing to the last parser.
//
//   - EBNF Operator: Concatenation
//   - EBNF Notation: p q
func (p Parser) CONCAT(q ...Parser) Parser {
	all := append([]Parser{p}, q...)
	return CONCAT(all...)
}

// OPT composes a parser that applies parser p zero or one time to the input.
// If the parser does not succeed, it will return an empty result.
//
//   - EBNF Operator: Optional
//   - EBNF Notation: [ p ] or p?
func (p Parser) OPT() Parser {
	return OPT(p)
}

// REP composes a parser that applies parser p zero or more times to the input and accumulates the results.
// If the parser does not succeed, it will return an empty result.
//
//   - EBNF Operator: Repetition (Kleene Star)
//   - EBNF Notation: { p } or p*
func (p Parser) REP() Parser {
	return REP(p)
}

// REP1 composes a parser that applies parser p one or more times to the input and accumulates the results.
// This does not allow parsing zero times (empty result).
//
//   - EBNF Operator: Kleene Plus
//   - EBNF Notation: p+
func (p Parser) REP1() Parser {
	return REP1(p)
}

// Flatten composes a parser that applies parser p to the input and flattens all results into a single list.
// This can be used for accessing the values of symbols in the right-side of a production rule more intuitively.
func (p Parser) Flatten() Parser {
	return func(in Input) (Output, bool) {
		if out, ok := p(in); ok {
			out.Result.Val = flatten(out.Result)
			return out, true
		}

		return Output{}, false
	}
}

func flatten(r Result) List {
	switch v := r.Val.(type) {
	case Empty:
		return List{}

	case List:
		var l List
		for _, w := range v {
			l = append(l, flatten(w)...)
		}
		return l

	default:
		return List{r}
	}
}

// Select composes a parser that applies parser p to the input and returns a list of symbols from the right-side of the production rule.
//
// This will not have any effect if the result of parsing is not a list.
// If indices are invalid, you will get the empty string ε.
func (p Parser) Select(i ...int) Parser {
	return func(in Input) (Output, bool) {
		out, ok := p(in)
		if !ok {
			return Output{}, false
		}

		l, ok := out.Result.Val.(List)
		if !ok {
			return out, true
		}

		sub := make(List, 0, len(i))
		for _, j := range i {
			if 0 <= j && j < len(l) {
				sub = append(sub, l[j])
			}
		}

		var res Result
		if len(sub) > 0 {
			res = Result{sub, sub[0].Pos, nil}
		} else {
			res = Result{
				Val: Empty{},
			}
		}

		return Output{
			Result:    res,
			Remaining: out.Remaining,
		}, true
	}
}

// Get composes a parser that applies parser p to the input and returns the value of a symbol from the right-side of the production rule.
// This can be used after CONCAT, REP, REP1, Flatten, and/or Select.
//
// It will not have any effect if used after other operators and the result of parsing is not a list.
// If index is invalid, you will get the empty string ε.
func (p Parser) Get(i int) Parser {
	return func(in Input) (Output, bool) {
		out, ok := p(in)
		if !ok {
			return Output{}, false
		}

		l, ok := out.Result.Val.(List)
		if !ok {
			return out, true
		}

		var res Result
		if 0 <= i && i < len(l) {
			res = l[i]
		} else {
			res = Result{
				Val: Empty{},
			}
		}

		return Output{
			Result:    res,
			Remaining: out.Remaining,
		}, ok
	}
}

// Map composes a parser that uses parser p to parse the input and applies a map function to the result.
// If the parser does not succeed, the map function will not be applied.
//
// Use Map to transform, annotate or convert parsing results into another form
// (for example: convert a matched rune into an int, build AST nodes, attach metadata, etc.).
//
// The remaining input returned by p is preserved when the mapping succeeds.
func (p Parser) Map(f MapFunc) Parser {
	return func(in Input) (Output, bool) {
		if out, ok := p(in); ok {
			if res, ok := f(out.Result); ok {
				out.Result = res
				return out, true
			}
		}

		return Output{}, false
	}
}

// MapFunc is a function that receives a parsing result and returns a new result.
// The second return value determines whether or not the mapping was successful.
type MapFunc func(Result) (Result, bool)

// Bind composes a parser that uses parser p to parse the input and produces a new parser from the result.
// It then applies the new parser to the remaining input from the first parser.
//
// Bind lets later parsing decisions depend on values parsed earlier.
// This is useful for context-sensitive checks that cannot be expressed by purely context-free grammar rules.
//
// Quick refresher:
//
//   - Context-free grammars define productions where
//     each rule's right-hand side does not depend on previously parsed values.
//     They are sufficient for many language constructs.
//   - Context-sensitive constructs require the parser to adapt based on previously parsed data
//     (for example, requiring a sequence to repeat N times where N was parsed earlier).
//
// Bind is a convenient way to add lightweight, context-sensitive constraints
// (sometimes called syntax annotations) on top of a context-free parser.
// The bound function inspects the parsing result and returns a new parser that enforces the constraint.
//
// Consider the following context-free production, where a number is followed by zero or more identifiers:
//
//	stmt → num id*
//
// Context-sensitive constraint: the number indicates how many identifiers must follow.
//
// Using Bind, you parse the number first, then return a parser that consumes exactly that many id occurrences.
// This keeps the grammar mostly context-free while enforcing a small context-sensitive constraint using a Bind-produced parser.
func (p Parser) Bind(f BindFunc) Parser {
	return func(in Input) (Output, bool) {
		out, ok := p(in)
		if !ok {
			return Output{}, false
		}

		return f(out.Result)(out.Remaining)
	}
}

// BindFunc is a function that receives a parsing result and returns a new parser.
// It is more powerful than MapFunc as it can produce a new parser based on the value of the parsing result.
type BindFunc func(Result) Parser
