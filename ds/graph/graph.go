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

// Visitor provides methods for visiting vertices and edges when traversing a graph.
//   VisitVertex is called when visiting a vertex in a graph (undirected, directed, weighted-undirected, and weighted-directed).
//   VisitEdge is called when visiting an edge in a graph (undirected and directed).
//   VisitWeightedEdge is called when visiting an edge in a graph (weighted-undirected and weighted-directed).
type Visitor struct {
	VisitVertex       func(int)
	VisitEdge         func(int, int)
	VisitWeightedEdge func(int, int, float64)
}
