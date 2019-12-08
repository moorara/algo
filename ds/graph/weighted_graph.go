package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// Edge represents a weighted undirected edge abstract data type.
type Edge interface {
	Either() int
	Other(int) int
	Weight() float64
}

// edge implements a weighted undirected edge data type.
type edge struct {
	v, w   int
	weight float64
}

// NewEdge creates a new weighted undirected edge.
func NewEdge(v, w int, weight float64) Edge {
	return &edge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

// Either returns either of this edge's vertices.
func (e *edge) Either() int {
	return e.v
}

// Other returns the other vertext of this edge.
func (e *edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	return e.v
}

// Weight returns the weight of this edge.
func (e *edge) Weight() float64 {
	return e.weight
}

// WeightedGraph represents a weighted undirected graph abstract data type.
type WeightedGraph interface {
	V() int
	E() int
	Degree(int) int
	AddEdge(Edge)
	Adj(int) []Edge
	Edges() []Edge
	Graphviz() string
}

// weightedGraph
type weightedGraph struct {
	v, e int
	adj  [][]Edge
}

// NewWeightedGraph creates a new weighted undirected graph.
func NewWeightedGraph(V int, edges ...Edge) WeightedGraph {
	adj := make([][]Edge, V)
	for i := range adj {
		adj[i] = make([]Edge, 0)
	}

	g := &weightedGraph{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		adj: adj,
	}

	for _, edge := range edges {
		g.AddEdge(edge)
	}

	return g
}

func (g *weightedGraph) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *weightedGraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *weightedGraph) E() int {
	return g.e
}

// Degree returns the degree of a vertext.
func (g *weightedGraph) Degree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *weightedGraph) AddEdge(e Edge) {
	v := e.Either()
	w := e.Other(v)

	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], e)
		g.adj[w] = append(g.adj[w], e)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *weightedGraph) Adj(v int) []Edge {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Edges returns all edges in the graph.
func (g *weightedGraph) Edges() []Edge {
	edges := make([]Edge, 0)
	for v := range g.adj {
		for _, e := range g.adj[v] {
			if e.Other(v) > v {
				edges = append(edges, e)
			}
		}
	}

	return edges
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *weightedGraph) Graphviz() string {
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
