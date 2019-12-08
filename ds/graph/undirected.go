// Package graph implements graph data structures and algorithms.
package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// Undirected represents an undirected graph data type.
type Undirected struct {
	v, e int
	adj  [][]int
}

// NewUndirected creates a new undirected graph.
func NewUndirected(V int, edges ...[2]int) *Undirected {
	adj := make([][]int, V)
	for i := range adj {
		adj[i] = make([]int, 0)
	}

	g := &Undirected{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		adj: adj,
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *Undirected) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *Undirected) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Undirected) E() int {
	return g.e
}

// Degree returns the degree of a vertext.
func (g *Undirected) Degree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *Undirected) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], w)
		g.adj[w] = append(g.adj[w], v)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *Undirected) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *Undirected) Graphviz() string {
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
