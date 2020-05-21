// Package graph implements graph data structures and algorithms.
//
// There are four different type of graphs implementd:
//   Undirected Graph
//   Directed Graph
//   Weighted Undirected Graph
//   Weighted Directed Graph
package graph

const (
	listNodeSize = 1024
)

// TraversalStrategy specifies the strategy for traversing vertices in a graph.
type TraversalStrategy int

const (
	// DFS is recursive depth-first search traversal.
	DFS TraversalStrategy = iota
	// DFSi is iterative depth-first search traversal.
	DFSi
	// BFS is breadth-first search traversal.
	BFS
)

// Visitors provides a method for visiting vertices and edges when traversing a graph.
// VertexPreOrder is called when visiting a vertex in a graph.
// VertexPostOrder is called when visiting a vertex in a graph.
// EdgePreOrder is called when visiting an edge in a graph.
// The graph traversal will immediately stop if the return value from any of these functions is false.
type Visitors struct {
	VertexPreOrder  func(int) bool
	VertexPostOrder func(int) bool
	EdgePreOrder    func(int, int, float64) bool
}
