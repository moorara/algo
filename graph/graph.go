// Package graph implements graph data structures and algorithms.
//
// There are four different type of graphs implementd:
//   Undirected Graph
//   Directed Graph
//   Weighted Undirected Graph
//   Weighted Directed Graph
package graph

import (
	"github.com/moorara/algo/list"
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

// DirectedCycle is used for determining if a directed graph has a cycle.
// A cycle is a path whose first and last vertices are the same.
type DirectedCycle struct {
	visited []bool
	edgeTo  []int
	onStack []bool
	cycle   list.Stack
}

func newDirectedCycle(g *Directed) *DirectedCycle {
	c := &DirectedCycle{
		visited: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
		onStack: make([]bool, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if !c.visited[v] && c.cycle == nil {
			c.dfs(g, v)
		}
	}

	return c
}

func (c *DirectedCycle) dfs(g *Directed, v int) {
	c.onStack[v] = true

	c.visited[v] = true
	for _, w := range g.adj[v] {
		if c.cycle != nil { // short circuit if a cycle already found
			return
		} else if !c.visited[w] {
			c.edgeTo[w] = v
			c.dfs(g, w)
		} else if c.onStack[w] { // cycle detected
			c.cycle = list.NewStack(listNodeSize)
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}

	c.onStack[v] = false
}

// Cycle returns a cyclic path.
// If no cycle exists, the second return value will be false.
func (c *DirectedCycle) Cycle() ([]int, bool) {
	if c.cycle == nil {
		return nil, false
	}

	cycle := make([]int, 0)
	for !c.cycle.IsEmpty() {
		cycle = append(cycle, c.cycle.Pop().(int))
	}

	return cycle, true
}

// Topological is used for determining the topological order of a directed acyclic graph (DAG).
// A directed graph has a topological order if and only if it is a DAG (no directed cycle exists).
type Topological struct {
	order []int // holds the topological order
	rank  []int // determines rank of a vertex in the topological order
}

// Order returns a topological order of the directed graph.
// If the directed graph does not have a topologial order (has a cycle), the second return value will be false.
func (t *Topological) Order() ([]int, bool) {
	if t.order == nil {
		return nil, false
	}

	return t.order, true
}

// Rank returns the rank of a vertex in the topological order.
// If the directed graph does not have a topologial order (has a cycle), the second return value will be false.
func (t *Topological) Rank(v int) (int, bool) {
	if t.rank == nil {
		return -1, false
	}

	return t.rank[v], true
}
