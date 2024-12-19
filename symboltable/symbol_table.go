// Package symboltable implements symbol table data structures.
//
// Symbol tables are also known as maps, dictionaries, etc.
// Symbol tables can be ordered or unordered.
package symboltable

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/moorara/algo/generic"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// SymbolTable represents an unordered symbol table abstract data type.
type SymbolTable[K, V any] interface {
	fmt.Stringer
	Equaler[SymbolTable[K, V]]
	Collection2[K, V]

	verify() bool
}

// OrderedSymbolTable represents an ordered symbol table abstract data type.
type OrderedSymbolTable[K, V any] interface {
	SymbolTable[K, V]
	Tree2[K, V]

	Height() int
	Min() (K, V, bool)
	Max() (K, V, bool)
	Floor(K) (K, V, bool)
	Ceiling(K) (K, V, bool)
	DeleteMin() (K, V, bool)
	DeleteMax() (K, V, bool)
	Select(int) (K, V, bool)
	Rank(K) int
	Range(K, K) []KeyValue[K, V]
	RangeSize(K, K) int
}
