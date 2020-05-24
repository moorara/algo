// Package graph implements graph data structures and algorithms.
//
// There are four different type of graphs implementd:
//   Undirected Graph
//   Directed Graph
//   Weighted Undirected Graph
//   Weighted Directed Graph
package graph

import (
	"github.com/moorara/algo/ds/list"
)

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

func isStrategyValid(strategy TraversalStrategy) bool {
	return strategy == DFS || strategy == DFSi || strategy == BFS
}

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

// Paths is used for finding all paths from a source vertex to every other vertex in a graph.
// A path is a sequence of vertices connected by edges.
type Paths struct {
	s       int
	visited []bool
	edgeTo  []int
}

// To returns a path between the source vertex (s) and a vertex (v).
// If no such path exists, the second return value will be false.
func (p *Paths) To(v int) ([]int, bool) {
	if !p.visited[v] {
		return nil, false
	}

	stack := list.NewStack(listNodeSize)
	for x := v; x != p.s; x = p.edgeTo[x] {
		stack.Push(x)
	}
	stack.Push(p.s)

	path := make([]int, 0)
	for !stack.IsEmpty() {
		path = append(path, stack.Pop().(int))
	}

	return path, true
}

// Orders is used for determining ordering of vertices in a graph.
type Orders struct {
	preRank, postRank   []int
	preOrder, postOrder []int
}

// PreRank returns the rank of a vertex in pre ordering.
func (o *Orders) PreRank(v int) int {
	return o.preRank[v]
}

// PostRank returns the rank of a vertex in post ordering.
func (o *Orders) PostRank(v int) int {
	return o.postRank[v]
}

// PreOrder returns the pre ordering of vertices.
func (o *Orders) PreOrder() []int {
	return o.preOrder
}

// PostOrder returns the post ordering of vertices.
func (o *Orders) PostOrder() []int {
	return o.postOrder
}

// ReversePostOrder returns the reverse post ordering of vertices.
func (o *Orders) ReversePostOrder() []int {
	l := len(o.postOrder)
	revOrder := make([]int, l)
	for i, v := range o.postOrder {
		revOrder[l-1-i] = v
	}

	return revOrder
}

// ConnectedComponents is used for determining all the connected components in a an undirected graph.
// A connected component is a maximal set of connected vertices
// (every two vertices are connected with a path between them).
type ConnectedComponents struct {
	count int
	id    []int
}

// ID returns the component id of a vertex.
func (c *ConnectedComponents) ID(v int) int {
	return c.id[v]
}

// IsConnected determines if two vertices v and w are connected.
func (c *ConnectedComponents) IsConnected(v, w int) bool {
	return c.id[v] == c.id[w]
}

// Components returns the vertices partitioned in the connected components.
func (c *ConnectedComponents) Components() [][]int {
	comps := make([][]int, c.count)
	for i := range comps {
		comps[i] = make([]int, 0)
	}

	for v, id := range c.id {
		comps[id] = append(comps[id], v)
	}

	return comps
}

// StronglyConnectedComponents is used for determining all the strongly connected components in a directed graph.
// A strongly connected component is a maximal set of strongly connected vertices
// (every two vertices are strongly connected with paths in both directions between them).
type StronglyConnectedComponents struct {
	count int
	id    []int
}

// ID returns the component id of a vertex.
func (c *StronglyConnectedComponents) ID(v int) int {
	return c.id[v]
}

// IsStronglyConnected determines if two vertices v and w are strongly connected.
func (c *StronglyConnectedComponents) IsStronglyConnected(v, w int) bool {
	return c.id[v] == c.id[w]
}

// Components returns the vertices partitioned in the connected components.
func (c *StronglyConnectedComponents) Components() [][]int {
	comps := make([][]int, c.count)
	for i := range comps {
		comps[i] = make([]int, 0)
	}

	for v, id := range c.id {
		comps[id] = append(comps[id], v)
	}

	return comps
}
