package graph

import (
	"fmt"

	"github.com/moorara/algo/dot"
	"github.com/moorara/algo/list"
)

// DirectedEdge represents a weighted directed edge data type.
type DirectedEdge struct {
	from, to int
	weight   float64
}

// From returns the vertex this edge points from.
func (e DirectedEdge) From() int {
	return e.from
}

// To returns the vertex this edge points to.
func (e DirectedEdge) To() int {
	return e.to
}

// Weight returns the weight of this edge.
func (e DirectedEdge) Weight() float64 {
	return e.weight
}

// WeightedDirected represents a weighted directed graph data type.
type WeightedDirected struct {
	v, e int
	ins  []int
	adj  [][]DirectedEdge
}

// NewWeightedDirected creates a new weighted directed graph.
func NewWeightedDirected(V int, edges ...DirectedEdge) *WeightedDirected {
	adj := make([][]DirectedEdge, V)
	for i := range adj {
		adj[i] = make([]DirectedEdge, 0)
	}

	g := &WeightedDirected{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		ins: make([]int, V),
		adj: adj,
	}

	for _, e := range edges {
		g.AddEdge(e)
	}

	return g
}

// V returns the number of vertices.
func (g *WeightedDirected) V() int {
	return g.v
}

// E returns the number of edges.
func (g *WeightedDirected) E() int {
	return g.e
}

func (g *WeightedDirected) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// InDegree returns the number of directed edges incident to a vertex.
func (g *WeightedDirected) InDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}
	return g.ins[v]
}

// OutDegree returns the number of directed edges incident from a vertex.
func (g *WeightedDirected) OutDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}
	return len(g.adj[v])
}

// Adj returns the vertices adjacent from vertex.
func (g *WeightedDirected) Adj(v int) []DirectedEdge {
	if !g.isVertexValid(v) {
		return nil
	}
	return g.adj[v]
}

// AddEdge adds a new edge to the graph.
func (g *WeightedDirected) AddEdge(e DirectedEdge) {
	v := e.From()
	w := e.To()

	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.ins[w]++
		g.adj[v] = append(g.adj[v], e)
	}
}

// Edges returns all directed edges in the graph.
func (g *WeightedDirected) Edges() []DirectedEdge {
	edges := make([]DirectedEdge, 0)
	for _, adjEdges := range g.adj {
		edges = append(edges, adjEdges...)
	}

	return edges
}

// Reverse returns the reverse of the directed graph.
func (g *WeightedDirected) Reverse() *WeightedDirected {
	rev := NewWeightedDirected(g.V())
	for v := 0; v < g.V(); v++ {
		for _, e := range g.adj[v] {
			rev.AddEdge(DirectedEdge{e.To(), e.From(), e.Weight()})
		}
	}

	return rev
}

// DFS Traversal (Recursion)
func (g *WeightedDirected) traverseDFS(v int, visited []bool, visitors *Visitors) {
	visited[v] = true

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(v) {
			return
		}
	}

	for _, e := range g.adj[v] {
		w := e.To()
		if !visited[w] {
			if visitors != nil && visitors.EdgePreOrder != nil {
				if !visitors.EdgePreOrder(v, w, e.Weight()) {
					return
				}
			}

			g.traverseDFS(w, visited, visitors)
		}
	}

	if visitors != nil && visitors.VertexPostOrder != nil {
		if !visitors.VertexPostOrder(v) {
			return
		}
	}
}

// Iterative DFS Traversal
func (g *WeightedDirected) traverseDFSi(s int, visited []bool, visitors *Visitors) {
	stack := list.NewStack[int](listNodeSize, nil)

	visited[s] = true
	stack.Push(s)

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !stack.IsEmpty() {
		v, _ := stack.Pop()

		if visitors != nil && visitors.VertexPostOrder != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.To()
			if !visited[w] {
				visited[w] = true
				stack.Push(w)

				if visitors != nil && visitors.VertexPreOrder != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
				}

				if visitors != nil && visitors.EdgePreOrder != nil {
					if !visitors.EdgePreOrder(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// BFS Traversal
func (g *WeightedDirected) traverseBFS(s int, visited []bool, visitors *Visitors) {
	queue := list.NewQueue[int](listNodeSize, nil)

	visited[s] = true
	queue.Enqueue(s)

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !queue.IsEmpty() {
		v, _ := queue.Dequeue()

		if visitors != nil && visitors.VertexPostOrder != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.To()
			if !visited[w] {
				visited[w] = true
				queue.Enqueue(w)

				if visitors != nil && visitors.VertexPreOrder != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
				}

				if visitors != nil && visitors.EdgePreOrder != nil {
					if !visitors.EdgePreOrder(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// Traverse is used for visiting all vertices and edges in graph.
func (g *WeightedDirected) Traverse(s int, strategy TraversalStrategy, visitors *Visitors) {
	if !g.isVertexValid(s) {
		return
	}

	visited := make([]bool, g.V())

	switch strategy {
	case DFS:
		g.traverseDFS(s, visited, visitors)
	case DFSi:
		g.traverseDFSi(s, visited, visitors)
	case BFS:
		g.traverseBFS(s, visited, visitors)
	}
}

// Paths finds all paths from a source vertex to every other vertex.
func (g *WeightedDirected) Paths(s int, strategy TraversalStrategy) *Paths {
	p := &Paths{
		s:       s,
		visited: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
	}

	if g.isVertexValid(s) && isStrategyValid(strategy) {
		visitors := &Visitors{
			EdgePreOrder: func(v, w int, weight float64) bool {
				p.edgeTo[w] = v
				return true
			},
		}

		switch strategy {
		case DFS:
			g.traverseDFS(s, p.visited, visitors)
		case DFSi:
			g.traverseDFSi(s, p.visited, visitors)
		case BFS:
			g.traverseBFS(s, p.visited, visitors)
		}
	}

	return p
}

// Orders determines ordering of vertices in the graph.
func (g *WeightedDirected) Orders(strategy TraversalStrategy) *Orders {
	o := &Orders{
		preRank:   make([]int, g.V()),
		postRank:  make([]int, g.V()),
		preOrder:  make([]int, 0),
		postOrder: make([]int, 0),
	}

	if isStrategyValid(strategy) {
		var preCounter, postCounter int
		visited := make([]bool, g.V())
		visitors := &Visitors{
			VertexPreOrder: func(v int) bool {
				o.preRank[v] = preCounter
				preCounter++
				o.preOrder = append(o.preOrder, v)
				return true
			},
			VertexPostOrder: func(v int) bool {
				o.postRank[v] = postCounter
				postCounter++
				o.postOrder = append(o.postOrder, v)
				return true
			},
		}

		for v := 0; v < g.V(); v++ {
			if !visited[v] {
				switch strategy {
				case DFS:
					g.traverseDFS(v, visited, visitors)
				case DFSi:
					g.traverseDFSi(v, visited, visitors)
				case BFS:
					g.traverseBFS(v, visited, visitors)
				}
			}
		}
	}

	return o
}

// StronglyConnectedComponents determines all the connected components in the graph.
func (g *WeightedDirected) StronglyConnectedComponents() *StronglyConnectedComponents {
	scc := &StronglyConnectedComponents{
		count: 0,
		id:    make([]int, g.V()),
	}

	visited := make([]bool, g.V())
	visitors := &Visitors{
		VertexPreOrder: func(v int) bool {
			scc.id[v] = scc.count
			return true
		},
	}

	order := g.Reverse().Orders(DFS).ReversePostOrder()
	for _, v := range order {
		if !visited[v] {
			g.traverseDFS(v, visited, visitors)
			scc.count++
		}
	}

	return scc
}

// ShortestPathTree calculates the shortest path tree of the graph.
func (g *WeightedDirected) ShortestPathTree(s int) *ShortestPathTree {
	return newShortestPathTree(g, s)
}

// DOT generates a DOT representation of the graph.
func (g *WeightedDirected) DOT() string {
	graph := dot.NewGraph(true, true, false, "", "", "", dot.StyleSolid, dot.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(dot.NewNode("", "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, e := range g.adj[v] {
			from := fmt.Sprintf("%d", e.From())
			to := fmt.Sprintf("%d", e.To())
			weight := fmt.Sprintf("%f", e.Weight())
			graph.AddEdge(dot.NewEdge(from, to, dot.EdgeTypeDirected, "", weight, "", "", "", ""))
		}
	}

	return graph.DOT()
}
