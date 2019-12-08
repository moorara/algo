package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// Digraph represents a directed graph abstract data type.
type Digraph interface {
	V() int
	E() int
	InDegree(int) int
	OutDegree(int) int
	AddEdge(int, int)
	Adj(int) []int
	Reverse() Digraph
	Graphviz() string
}

// digraph implements a directed graph data type.
type digraph struct {
	v, e int
	ins  []int
	adj  [][]int
}

// NewDigraph creates a new directed graph.
func NewDigraph(V int, edges ...[2]int) Digraph {
	adj := make([][]int, V)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	g := &digraph{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		ins: make([]int, V),
		adj: adj,
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *digraph) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *digraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *digraph) E() int {
	return g.e
}

// InDegree returns the number of directed edges incident to a vertex.
func (g *digraph) InDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return g.ins[v]
}

// OutDegree returns the number of directed edges incident from a vertex.
func (g *digraph) OutDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *digraph) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.ins[w]++
		g.adj[v] = append(g.adj[v], w)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *digraph) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Reverse returns the reverse of the directed graph.
func (g *digraph) Reverse() Digraph {
	rev := NewDigraph(g.v)
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			rev.AddEdge(w, v)
		}
	}

	return rev
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *digraph) Graphviz() string {
	graph := graphviz.NewGraph(true, true, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode(name, "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, w := range g.adj[v] {
			from := fmt.Sprintf("%d", v)
			to := fmt.Sprintf("%d", w)
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
