package grammar

import (
	"fmt"
	"hash/fnv"
	"io"

	"github.com/moorara/algo/generic"
	"github.com/moorara/algo/hash"
)

// The endmarker is a special symbol that is used to indicate the end of a string.
// This special symbol assumed not to be a symbol of any grammar and
// it is taken from a Private Use Area (PUA) in Unicode.
//
// The endmarker is not a formal part of the grammar itself but is introduced during parsing
// to simplify the handling of end-of-input scenarios, especially in parsing algorithms like LL(1) or LR(1).
//
// For more information and details, see "Compilers: Principles, Techniques, and Tools (2nd Edition)".
// const endmarker rune = 0xEEEE

var (
	hashSymbol = hashFuncForSymbol()
	eqSymbol   = func(lhs, rhs Symbol) bool {
		return lhs.Equals(rhs)
	}

	eqTerminal  = generic.NewEqualFunc[Terminal]()
	cmpTerminal = generic.NewCompareFunc[Terminal]()

	eqNonTerminal   = generic.NewEqualFunc[NonTerminal]()
	cmpNonTerminal  = generic.NewCompareFunc[NonTerminal]()
	hashNonTerminal = hash.HashFuncForString[NonTerminal](nil)
)

// Symbol represents a grammar symbol (terminal or non-terminal).
type Symbol interface {
	fmt.Stringer

	Name() string
	Equals(Symbol) bool
	IsTerminal() bool
}

// WriteSymbol writes the string representation of a symbol to the provided io.Writer.
// It returns the number of bytes written and any error encountered.
func WriteSymbol(w io.Writer, s Symbol) (n int, err error) {
	return w.Write([]byte(s.String()))
}

// hashFuncForSymbol creates a HashFunc for hashing symbols.
func hashFuncForSymbol() hash.HashFunc[Symbol] {
	h := fnv.New64()

	return func(s Symbol) uint64 {
		h.Reset()
		_, _ = WriteSymbol(h, s) // Hash.Write never returns an error
		return h.Sum64()
	}
}

// Terminal represents a terminal symbol.
// Terminals are the basic symbols from which strings of a language are formed.
// Token name or token for short are equivalent to terminal.
type Terminal string

// String returns a string representation of a terminal symbol.
func (t Terminal) String() string {
	return fmt.Sprintf("%q", t.Name())
}

// Name returns the name of terminal symbol.
func (t Terminal) Name() string {
	return string(t)
}

// Equals determines whether or not two terminal symbols are the same.
func (t Terminal) Equals(rhs Symbol) bool {
	if v, ok := rhs.(Terminal); ok {
		return t == v
	}
	return false
}

// IsTerminal always returns true for terminal symbols.
func (t Terminal) IsTerminal() bool {
	return true
}

// NonTerminal represents a non-terminal symbol.
// Non-terminals are syntaxtic variables that denote sets of strings.
// Non-terminals impose a hierarchical structure on a language.
type NonTerminal string

// String returns a string representation of a non-terminal symbol.
func (n NonTerminal) String() string {
	return n.Name()
}

// Name returns the name of non-terminal symbol.
func (n NonTerminal) Name() string {
	return string(n)
}

// Equals determines whether or not two non-terminal symbols are the same.
func (n NonTerminal) Equals(rhs Symbol) bool {
	if v, ok := rhs.(NonTerminal); ok {
		return n == v
	}
	return false
}

// IsTerminal always returns false for non-terminal symbols.
func (n NonTerminal) IsTerminal() bool {
	return false
}
