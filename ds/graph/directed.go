package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// Directed represents a directed graph data type.
type Directed struct {
	v, e int
	ins  []int
	adj  [][]int
}

// NewDirected creates a new directed graph.
func NewDirected(V int, edges ...[2]int) *Directed {
	adj := make([][]int, V)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	g := &Directed{
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

func (g *Directed) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *Directed) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Directed) E() int {
	return g.e
}

// InDegree returns the number of directed edges incident to a vertex.
func (g *Directed) InDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return g.ins[v]
}

// OutDegree returns the number of directed edges incident from a vertex.
func (g *Directed) OutDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *Directed) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.ins[w]++
		g.adj[v] = append(g.adj[v], w)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *Directed) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Reverse returns the reverse of the directed graph.
func (g *Directed) Reverse() *Directed {
	rev := NewDirected(g.v)
	for v := 0; v < g.v; v++ {
		for _, w := range g.adj[v] {
			rev.AddEdge(w, v)
		}
	}

	return rev
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *Directed) Graphviz() string {
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
