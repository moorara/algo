// Package graph implements graph data structures and algorithms.
//
// There are four different type of graphs implementd:
//
//	Undirected Graph
//	Directed Graph
//	Weighted Undirected Graph
//	Weighted Directed Graph
package graph

import (
	"math"

	. "github.com/moorara/algo/generic"
	"github.com/moorara/algo/heap"
	"github.com/moorara/algo/list"
)

const (
	listNodeSize   = 1024
	float64Epsilon = 1e-9
)

// TraversalStrategy is the strategy for traversing vertices in a graph.
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

// OptimizationStrategy is the strategy for optimizing weighted graphs.
type OptimizationStrategy int

const (
	// None ignores edge weights.
	None OptimizationStrategy = iota
	// Minimize picks the edges with minimum weight.
	Minimize
	// Maximize picks the edges with maximum weight.
	Maximize
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

// Paths is used for finding all paths from a source vertex to every other vertex in a graph.
// A path is a sequence of vertices connected by edges.
type Paths struct {
	s       int
	visited []bool
	edgeTo  []int
}

// To returns a path between the source vertex (s) and vertex (v).
// If no such path exists, the second return value will be false.
func (p *Paths) To(v int) ([]int, bool) {
	if !p.visited[v] {
		return nil, false
	}

	stack := list.NewStack[int](listNodeSize, nil)
	for x := v; x != p.s; x = p.edgeTo[x] {
		stack.Push(x)
	}
	stack.Push(p.s)

	path := make([]int, 0)
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		path = append(path, v)
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
	cycle   list.Stack[int]
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
			c.cycle = list.NewStack[int](listNodeSize, nil)
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
		v, _ := c.cycle.Pop()
		cycle = append(cycle, v)
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

// MinimumSpanningTree is used for calculating the minimum spanning trees (forest) of a weighted undirected graph.
// Given an edge-weighted undirected graph G with positive edge weights, an MST of G is a sub-graph T that is:
//
//	Tree: connected and acyclic
//	Spanning: includes all of the vertices
//	Minimum: sum of the edge wights are minimum
type MinimumSpanningTree struct {
	visited []bool                         // visited[v] = true if v on tree, false otherwise
	edgeTo  []UndirectedEdge               // edgeTo[v] = shortest edge from tree vertex to non-tree vertex
	distTo  []float64                      // distTo[v] = weight of shortest such edge (edgeTo[v].weight())
	pq      heap.IndexedHeap[float64, any] // indexed priority queue of vertices connected by an edge to tree
}

func newMinimumSpanningTree(g *WeightedUndirected) *MinimumSpanningTree {
	mst := &MinimumSpanningTree{
		visited: make([]bool, g.V()),
		edgeTo:  make([]UndirectedEdge, g.V()),
		distTo:  make([]float64, g.V()),
		pq:      heap.NewIndexedBinary[float64, any](g.V(), NewCompareFunc[float64](), nil),
	}

	for v := 0; v < g.V(); v++ {
		mst.distTo[v] = math.MaxFloat64
	}

	// run from each vertex to find minimum spanning forest
	for v := 0; v < g.V(); v++ {
		if !mst.visited[v] {
			mst.prim(g, v)
		}
	}

	return mst
}

// Prim's algorithm (eager version) for calculating minimum spanning tree.
func (mst *MinimumSpanningTree) prim(g *WeightedUndirected, s int) {
	mst.distTo[s] = 0.0
	mst.pq.Insert(s, mst.distTo[s], nil)

	for !mst.pq.IsEmpty() {
		v, _, _, _ := mst.pq.Delete()
		mst.visited[v] = true

		for _, e := range g.Adj(v) {
			w := e.Other(v)
			if mst.visited[w] {
				continue
			}

			if e.Weight() < mst.distTo[w] {
				mst.edgeTo[w] = e
				mst.distTo[w] = e.Weight()

				if mst.pq.ContainsIndex(w) {
					mst.pq.ChangeKey(w, mst.distTo[w])
				} else {
					mst.pq.Insert(w, mst.distTo[w], nil)
				}
			}
		}
	}
}

// Edges returns the edges in a minimum spanning tree (or forest).
func (mst *MinimumSpanningTree) Edges() []UndirectedEdge {
	zero := UndirectedEdge{}
	edges := make([]UndirectedEdge, 0)
	for _, e := range mst.edgeTo {
		if e != zero {
			edges = append(edges, e)
		}
	}

	return edges
}

// Weight returns the sum of the edge weights in a minimum spanning tree (or forest).
func (mst *MinimumSpanningTree) Weight() float64 {
	var weight float64
	for _, e := range mst.Edges() {
		weight += e.Weight()
	}

	return weight
}

// ShortestPathTree is used for calculating the shortest path tree of a weighted directed graph.
// A shortest path from vertex s to vertex t in a weighted directed graph is a directed path from s to t such that no other path has a lower weight.
type ShortestPathTree struct {
	edgeTo []DirectedEdge                 // edgeTo[v] = last edge on shortest path s->v
	distTo []float64                      // distTo[v] = distance of shortest path s->v
	pq     heap.IndexedHeap[float64, any] // indexed priority queue of vertices
}

func newShortestPathTree(g *WeightedDirected, s int) *ShortestPathTree {
	spt := &ShortestPathTree{
		edgeTo: make([]DirectedEdge, g.V()),
		distTo: make([]float64, g.V()),
		pq:     heap.NewIndexedBinary[float64, any](g.V(), NewCompareFunc[float64](), nil),
	}

	for v := 0; v < g.V(); v++ {
		spt.distTo[v] = math.MaxFloat64
	}

	spt.dijkstra(g, s)

	return spt
}

// Dijkstra's algorithm (eager version) for calculating shortest path tree.
func (spt *ShortestPathTree) dijkstra(g *WeightedDirected, s int) {
	spt.distTo[s] = 0.0
	spt.pq.Insert(s, spt.distTo[s], nil)

	for !spt.pq.IsEmpty() {
		v, _, _, _ := spt.pq.Delete()

		// Relaxing edges
		for _, e := range g.Adj(v) {
			v, w := e.From(), e.To()
			if dist := spt.distTo[v] + e.Weight(); dist < spt.distTo[w] {
				spt.edgeTo[w] = e
				spt.distTo[w] = dist

				if spt.pq.ContainsIndex(w) {
					spt.pq.ChangeKey(w, spt.distTo[w])
				} else {
					spt.pq.Insert(w, spt.distTo[w], nil)
				}
			}
		}
	}
}

// PathTo returns shortest path from the source vertex (s) to vertex (v).
// The second return value is distance from the source vertex (s) to vertex (v).
// If no such path exists, the last return value will be false.
func (spt *ShortestPathTree) PathTo(v int) ([]DirectedEdge, float64, bool) {
	if spt.distTo[v] == math.MaxFloat64 {
		return nil, -1, false
	}

	zero := DirectedEdge{}
	stack := list.NewStack[DirectedEdge](listNodeSize, nil)
	for e := spt.edgeTo[v]; e != zero; e = spt.edgeTo[e.From()] {
		stack.Push(e)
	}

	path := make([]DirectedEdge, stack.Size())
	for i := range path {
		path[i], _ = stack.Pop()
	}

	return path, spt.distTo[v], true
}
