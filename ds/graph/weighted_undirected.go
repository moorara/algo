package graph

import (
	"fmt"

	"github.com/moorara/algo/ds/list"
	"github.com/moorara/algo/pkg/graphviz"
)

// UndirectedEdge represents a weighted undirected edge data type.
type UndirectedEdge struct {
	v, w   int
	weight float64
}

// Either returns either of this edge's vertices.
func (e UndirectedEdge) Either() int {
	return e.v
}

// Other returns the other vertext of this edge.
func (e UndirectedEdge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	return e.v
}

// Weight returns the weight of this edge.
func (e UndirectedEdge) Weight() float64 {
	return e.weight
}

// WeightedUndirected represents a weighted undirected graph data type.
type WeightedUndirected struct {
	v, e int
	adj  [][]UndirectedEdge
}

// NewWeightedUndirected creates a new weighted undirected graph.
func NewWeightedUndirected(V int, edges ...UndirectedEdge) *WeightedUndirected {
	adj := make([][]UndirectedEdge, V)
	for i := range adj {
		adj[i] = make([]UndirectedEdge, 0)
	}

	g := &WeightedUndirected{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		adj: adj,
	}

	for _, e := range edges {
		g.AddEdge(e)
	}

	return g
}

// V returns the number of vertices.
func (g *WeightedUndirected) V() int {
	return g.v
}

// E returns the number of edges.
func (g *WeightedUndirected) E() int {
	return g.e
}

func (g *WeightedUndirected) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// Degree returns the degree of a vertext.
func (g *WeightedUndirected) Degree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}
	return len(g.adj[v])
}

// Adj returns the vertices adjacent from vertex.
func (g *WeightedUndirected) Adj(v int) []UndirectedEdge {
	if !g.isVertexValid(v) {
		return nil
	}
	return g.adj[v]
}

// AddEdge adds a new edge to the graph.
func (g *WeightedUndirected) AddEdge(e UndirectedEdge) {
	v := e.Either()
	w := e.Other(v)

	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], e)
		g.adj[w] = append(g.adj[w], e)
	}
}

// Edges returns all edges in the graph.
func (g *WeightedUndirected) Edges() []UndirectedEdge {
	edges := make([]UndirectedEdge, 0)
	for v := range g.adj {
		for _, e := range g.adj[v] {
			if e.Other(v) > v { // Consider every edge only once
				edges = append(edges, e)
			}
		}
	}

	return edges
}

// DFS Traversal (Recursion)
func (g *WeightedUndirected) _traverseDFS(visited []bool, v int, visitors *Visitors) {
	visited[v] = true

	if visitors != nil {
		if !visitors.VertexPreOrder(v) {
			return
		}
	}

	for _, e := range g.adj[v] {
		w := e.Other(v)
		if !visited[w] {
			if visitors != nil {
				if !visitors.EdgePreOrder(v, w, e.Weight()) {
					return
				}
			}

			g._traverseDFS(visited, w, visitors)
		}
	}

	if visitors != nil {
		if !visitors.VertexPostOrder(v) {
			return
		}
	}
}

// DFS Traversal (Driver)
func (g *WeightedUndirected) traverseDFS(s int, visitors *Visitors) {
	visited := make([]bool, g.V())
	g._traverseDFS(visited, s, visitors)
}

// Iterative DFS Traversal
func (g *WeightedUndirected) traverseDFSi(s int, visitors *Visitors) {
	visited := make([]bool, g.V())
	stack := list.NewStack(listNodeSize)

	visited[s] = true
	stack.Push(s)

	if visitors != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !stack.IsEmpty() {
		v := stack.Pop().(int)

		if visitors != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.Other(v)
			if !visited[w] {
				visited[w] = true
				stack.Push(w)

				if visitors != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
					if !visitors.EdgePreOrder(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// BFS Traversal
func (g *WeightedUndirected) traverseBFS(s int, visitors *Visitors) {
	visited := make([]bool, g.V())
	queue := list.NewQueue(listNodeSize)

	visited[s] = true
	queue.Enqueue(s)

	if visitors != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)

		if visitors != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.Other(v)
			if !visited[w] {
				visited[w] = true
				queue.Enqueue(w)

				if visitors != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
					if !visitors.EdgePreOrder(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// Traverse is used for visiting all vertices and edges in graph.
func (g *WeightedUndirected) Traverse(s int, strategy TraversalStrategy, visitors *Visitors) {
	if !g.isVertexValid(s) {
		return
	}

	switch strategy {
	case DFS:
		g.traverseDFS(s, visitors)
	case DFSi:
		g.traverseDFSi(s, visitors)
	case BFS:
		g.traverseBFS(s, visitors)
	}
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *WeightedUndirected) Graphviz() string {
	graph := graphviz.NewGraph(true, false, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode("", "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, e := range g.adj[v] {
			if e.Other(v) > v {
				from := fmt.Sprintf("%d", v)
				to := fmt.Sprintf("%d", e.Other(v))
				graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeUndirected, "", "", "", "", ""))
			}
		}
	}

	return graph.DotCode()
}
