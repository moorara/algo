package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// Graph represents an undirected graph abstract data type.
type Graph interface {
	V() int
	E() int
	Degree(int) int
	AddEdge(int, int)
	Adj(int) []int
	Graphviz() string
}

// graph implements an undirected graph data type.
type graph struct {
	v, e int
	adj  [][]int
}

// NewGraph creates a new undirected graph.
func NewGraph(V int, edges ...[2]int) Graph {
	adj := make([][]int, V)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	g := &graph{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		adj: adj,
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *graph) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *graph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *graph) E() int {
	return g.e
}

// Degree returns the degree of a vertext.
func (g *graph) Degree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *graph) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], w)
		g.adj[w] = append(g.adj[w], v)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *graph) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *graph) Graphviz() string {
	graph := graphviz.NewGraph(true, false, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode(name, "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, w := range g.adj[v] {
			from := fmt.Sprintf("%d", v)
			to := fmt.Sprintf("%d", w)
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeUndirected, "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
