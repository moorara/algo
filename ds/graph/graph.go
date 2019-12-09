// Package graph implements graph data structures and algorithms.
//
// There are four different type of graphs implementd:
//   - Undirected Graph
//   - Directed Graph
//   - Weighted Undirected Graph
//   - Weighted Directed Graph
package graph

const (
	listNodeSize = 1024
)

// TraverseStrategy specifies the strategy for traversing vertices in a graph.
type TraverseStrategy int

const (
	// DFS is recursive depth-first search traversal.
	DFS TraverseStrategy = iota
	// DFSIterative is iterative depth-first search traversal.
	DFSIterative
	// BFS is breadth-first search traversal.
	BFS
)

// TraverseOrder specifies the order for traversing vertices in a graph.
type TraverseOrder int

const (
	// PreOrder is pre-order traversal.
	PreOrder TraverseOrder = iota
	// PostOrder is post-order traversal.
	PostOrder
)

type (
	// The VisitFunc type is a function for visiting a graph vertex.
	VisitFunc func(int)
)
