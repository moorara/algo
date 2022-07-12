// Package trie implements prefix tree data structures.
package trie

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
	VisitFunc[V any] func(string, V) bool

	// KeyValue represents a key-value pair.
	KeyValue[V any] struct {
		key string
		val V
	}
)

// OrderedSymbolTable represents a trie (prefix tree) abstract data type.
type Trie[V any] interface {
	verify() bool
	Size() int
	Height() int
	IsEmpty() bool
	Put(string, V)
	Get(string) (V, bool)
	Delete(string) (V, bool)
	KeyValues() []KeyValue[V]

	Min() (string, V, bool)
	Max() (string, V, bool)
	Floor(string) (string, V, bool)
	Ceiling(string) (string, V, bool)
	DeleteMin() (string, V, bool)
	DeleteMax() (string, V, bool)
	Select(int) (string, V, bool)
	Rank(string) int
	RangeSize(string, string) int
	Range(string, string) []KeyValue[V]
	Traverse(TraversalOrder, VisitFunc[V])
	Graphviz() string

	Match(string) []KeyValue[V]
	WithPrefix(string) []KeyValue[V]
	LongestPrefix(string) (string, V, bool)
}
