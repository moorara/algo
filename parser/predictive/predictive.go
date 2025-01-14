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
