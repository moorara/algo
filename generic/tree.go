package generic

// TraverseOrder represents the order in which nodes are traversed in a tree.
type TraverseOrder int

const (
	// VLR is a pre-order traversal from left to right.
	VLR TraverseOrder = iota
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
	// Ascending is an ascending traversal.
	Ascending
	// Descending is a descending traversal.
	Descending
)

type (
	// VisitFunc1 is a generic function type used during tree traversal
	//   for processing nodes with a single value.
	VisitFunc1[T any] func(T) bool

	// VisitFunc2 is a generic function type used during tree traversal
	//   for processing nodes with key-value pairs.
	VisitFunc2[K, V any] func(K, V) bool
)

type (
	// Tree1 represents a generic tree structure where nodes contain a single value.
	Tree1[T any] interface {
		// Traverse performs a traversal of the tree using the specified traversal order
		//   and yields the value of each node to the provided VisitFunc1 function.
		// If the function returns false, the traversal is halted.
		Traverse(TraverseOrder, VisitFunc1[T])

		// Graphviz generates and returns a string representation of the tree in DOT format.
		// This format is commonly used for visualizing graphs with Graphviz tools.
		Graphviz() string
	}

	// Tree2 represents a generic tree structure where nodes contain key-value pairs.
	Tree2[K, V any] interface {
		// Traverse performs a traversal of the tree using the specified traversal order
		//   and yields the key-value pair of each node to the provided VisitFunc2 function.
		// If the function returns false, the traversal is halted.
		Traverse(TraverseOrder, VisitFunc2[K, V])

		// Graphviz generates and returns a string representation of the tree in DOT format.
		// This format is commonly used for visualizing graphs with Graphviz tools.
		Graphviz() string
	}
)
