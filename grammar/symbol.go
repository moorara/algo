package grammar

import (
	"fmt"

	. "github.com/moorara/algo/generic"
	. "github.com/moorara/algo/hash"
)

var (
	eqTerminal  = NewEqualFunc[Terminal]()
	cmpTerminal = NewCompareFunc[Terminal]()

	eqNonTerminal   = NewEqualFunc[NonTerminal]()
	cmpNonTerminal  = NewCompareFunc[NonTerminal]()
	hashNonTerminal = HashFuncForString[NonTerminal](nil)
)

// Symbol represents a grammar symbol (terminal or non-terminal).
type Symbol interface {
	fmt.Stringer

	Name() string
	Equals(Symbol) bool
	IsTerminal() bool
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
