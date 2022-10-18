// Package symboltable implements symbol table data structures.
//
// Symbol tables are also known as maps, dictionaries, etc.
// Symbol tables can be ordered or unordered.
package symboltable

// TraversalOrder is the order for traversing nodes in a tree.
type TraversalOrder int

const (
	// VLR is a pre-order traversal from left to right.
	VLR TraversalOrder = iota
	// VLR is a pre-order traversal from right to left.
	VRL
	// VLR is an in-order traversal from left to right.
	LVR
	// RVL is an in-order traversal from right to left.
	RVL
	// LRV is a post-order traversal from left to right.
	LRV
	// LRV is a post-order traversal from right to left.
	RLV
	// Ascending is key-ascending traversal.
	Ascending
	// Descending is key-descending traversal.
	Descending
)

type (
	// The VisitFunc type is a function for visiting a key-value pair.
	VisitFunc[K, V any] func(K, V) bool

	// KeyValue represents a key-value pair.
	KeyValue[K, V any] struct {
		Key K
		Val V
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
	Equals(SymbolTable[K, V]) bool
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
