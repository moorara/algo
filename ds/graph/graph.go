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

// TraversalOrder specifies the order for traversing vertices in a graph.
type TraversalOrder int

const (
	// PreOrder is pre-order traversal.
	PreOrder TraversalOrder = iota
	// PostOrder is post-order traversal.
	PostOrder
)

type (
	// VertexVisitor provides methods for visiting vertices when traversing a graph.
	//   VisitVertex is called when visiting a vertex in a graph.
	VertexVisitor interface {
		VisitVertex(int) bool
	}

	// EdgeVisitor provides methods for visiting edges when traversing a graph.
	//   VisitEdge is called when visiting an edge in a graph.
	EdgeVisitor interface {
		VisitEdge(int, int) bool
	}

	// WeightedEdgeVisitor provides methods for visiting edges when traversing a graph with weighted edges.
	//   VisitWeightedEdge is called when visiting an edge in a graph with weighted edges.
	WeightedEdgeVisitor interface {
		VisitWeightedEdge(int, int, float64) bool
	}
)
