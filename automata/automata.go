// Package automata provides data structures and algorithms for working with automata.
//
// In language theory, automata refer to abstract computational models used to define and study formal languages.
// Automata are mathematical structures that process input strings and determine whether they belong to a specific language.
package automata

import "github.com/moorara/algo/symboltable"

// doubleKeyMap is a map (symbol table) data structure with two keys.
type doubleKeyMap[K1, K2, V any] symboltable.SymbolTable[K1, symboltable.SymbolTable[K2, V]]
