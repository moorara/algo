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
	// The CompareFunc type is a function for comparing two values of the same type.
	CompareFunc func(interface{}, interface{}) int

	// The VisitFunc type is a function for visiting a key-value pair.
	VisitFunc func(interface{}, interface{}) bool

	// KeyValue represents a key-value pair.
	KeyValue struct {
		key   interface{}
		value interface{}
	}

	// SymbolTable represents an unordered symbol table abstract data type.
	SymbolTable interface {
		verify() bool
		Size() int
		Height() int
		IsEmpty() bool
		Put(interface{}, interface{})
		Get(interface{}) (interface{}, bool)
		Delete(interface{}) (interface{}, bool)
		KeyValues() []KeyValue
	}

	// OrderedSymbolTable represents an ordered symbol table abstract data type.
	OrderedSymbolTable interface {
		SymbolTable
		Min() (interface{}, interface{})
		Max() (interface{}, interface{})
		Floor(interface{}) (interface{}, interface{})
		Ceiling(interface{}) (interface{}, interface{})
		Rank(interface{}) int
		Select(int) (interface{}, interface{})
		DeleteMin() (interface{}, interface{})
		DeleteMax() (interface{}, interface{})
		RangeSize(interface{}, interface{}) int
		Range(interface{}, interface{}) []KeyValue
		Traverse(TraversalOrder, VisitFunc)
		Graphviz() string
	}
)
