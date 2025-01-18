// Package lalr provides data structures and algorithms for building Look-Ahead LR (LALR) parsers.
// An LALR parser is a bottom-up parser for the class of LR(1) grammars.
//
// Bottom-up parsing constructs a parse tree for an input string
// starting at the leaves (bottom) and working towards the root (top).
// This process involves reducing a string w to the start symbol of the grammar.
// At each reduction step, a specific substring matching the body of a production
// is replaced by the non-terminal at the head of that production.
//
// Bottom-up parsing during a left-to-right scan of the inputconstructs a rightmost derivation in reverse:
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
//
// LR(1) parsers determine the next state based on the current state, one lookahead symbol, and the top of the stack.
// An LALR parser, similar to SLR, uses the canonical LR(0) items to construct the state machine (DFA),
// but refines the states by incorporating lookahead symbols explicitly.
// LALR merges states with identical core LR(0) items but handles lookahead symbols for each merged state separately,
// making it more precise than SLR and avoids many conflicts that SLR might encounter.
// LALR is more powerful than SLR as it can handle a wider range of grammars, including most programming languages.
// However, it is less powerful than canonical LR(1) because state merging
// can lose distinctions in lookahead contexts, potentially leading to conflicts for some grammars.
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
package lalr
