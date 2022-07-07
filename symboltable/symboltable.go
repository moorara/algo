// Package symboltable implements symbol table data structures.
//
// Symbol tables are also known as maps, dictionaries, etc.
// Symbol tables can be ordered or unordered.
package symboltable

// TraversalOrder is the order for traversing nodes in a tree.
type TraversalOrder int

const (
	// PreOrder is pre-order traversal.
	PreOrder TraversalOrder = iota
	// InOrder is in-order traversal.
	InOrder
	// PostOrder is post-order traversal.
	PostOrder
)

type (
	// The VisitFunc type is a function for visiting a key-value pair.
	VisitFunc[K, V any] func(K, V) bool

	// KeyValue represents a key-value pair.
	KeyValue[K, V any] struct {
		key K
		val V
	}
)

// SymbolTable represents an unordered symbol table abstract data type.
type SymbolTable[K, V any] interface {
	verify() bool
	Size() int
	Height() int
	IsEmpty() bool
	Put(K, V)
	Get(K) (V, bool)
	Delete(K) (V, bool)
	KeyValues() []KeyValue[K, V]
}

// OrderedSymbolTable represents an ordered symbol table abstract data type.
type OrderedSymbolTable[K, V any] interface {
	SymbolTable[K, V]
	Min() (K, V, bool)
	Max() (K, V, bool)
	Floor(K) (K, V, bool)
	Ceiling(K) (K, V, bool)
	DeleteMin() (K, V, bool)
	DeleteMax() (K, V, bool)
	Select(int) (K, V, bool)
	Rank(K) int
	RangeSize(K, K) int
	Range(K, K) []KeyValue[K, V]
	Traverse(TraversalOrder, VisitFunc[K, V])
	Graphviz() string
}
