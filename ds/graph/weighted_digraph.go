package graph

import (
	"fmt"

	"github.com/moorara/algo/pkg/graphviz"
)

// DirectedEdge represents a weighted directed edge abstract data type.
type DirectedEdge interface {
	From() int
	To() int
	Weight() float64
}

// directedEdge implements a weighted directed edge data type.
type directedEdge struct {
	from, to int
	weight   float64
}

// NewDirectedEdge creates a new weighted directed edge.
func NewDirectedEdge(from, to int, weight float64) DirectedEdge {
	return &directedEdge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

// From returns the vertex this edge points from.
func (e *directedEdge) From() int {
	return e.from
}

// To returns the vertex this edge points to.
func (e *directedEdge) To() int {
	return e.to
}

// Weight returns the weight of this edge.
func (e *directedEdge) Weight() float64 {
	return e.weight
}

// WeightedDigraph represents a weighted directed graph abstract data type.
type WeightedDigraph interface {
	V() int
	E() int
	InDegree(int) int
	OutDegree(int) int
	AddEdge(DirectedEdge)
	Adj(int) []DirectedEdge
	Edges() []DirectedEdge
	Reverse() WeightedDigraph
	Graphviz() string
}

// weightedDigraph
type weightedDigraph struct {
	v, e int
	ins  []int
	adj  [][]DirectedEdge
}

// NewWeightedDigraph creates a new weighted directed graph.
func NewWeightedDigraph(V int, edges ...DirectedEdge) WeightedDigraph {
	adj := make([][]DirectedEdge, V)
	for i := range adj {
		adj[i] = make([]DirectedEdge, 0)
	}

	g := &weightedDigraph{
		v:   V, // no. of vertices
		e:   0, // no. of edges
		ins: make([]int, V),
		adj: adj,
	}

	for _, edge := range edges {
		g.AddEdge(edge)
	}

	return g
}

func (g *weightedDigraph) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// V returns the number of vertices.
func (g *weightedDigraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *weightedDigraph) E() int {
	return g.e
}

// InDegree returns the number of directed edges incident to a vertex.
func (g *weightedDigraph) InDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return g.ins[v]
}

// OutDegree returns the number of directed edges incident from a vertex.
func (g *weightedDigraph) OutDegree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}

	return len(g.adj[v])
}

// AddEdge adds a new edge to the graph.
func (g *weightedDigraph) AddEdge(e DirectedEdge) {
	v := e.From()
	w := e.To()

	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.ins[w]++
		g.adj[v] = append(g.adj[v], e)
	}
}

// Adj returns the vertices adjacent from vertex.
func (g *weightedDigraph) Adj(v int) []DirectedEdge {
	if !g.isVertexValid(v) {
		return nil
	}

	return g.adj[v]
}

// Edges returns all directed edges in the graph.
func (g *weightedDigraph) Edges() []DirectedEdge {
	edges := make([]DirectedEdge, 0)
	for _, adjEdges := range g.adj {
		edges = append(edges, adjEdges...)
	}

	return edges
}

// Reverse returns the reverse of the directed graph.
func (g *weightedDigraph) Reverse() WeightedDigraph {
	rev := NewWeightedDigraph(g.v)
	for v := 0; v < g.v; v++ {
		for _, e := range g.adj[v] {
			edge := NewDirectedEdge(e.To(), e.From(), e.Weight())
			rev.AddEdge(edge)
		}
	}

	return rev
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *weightedDigraph) Graphviz() string {
	graph := graphviz.NewGraph(true, true, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode("", "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, e := range g.adj[v] {
			from := fmt.Sprintf("%d", e.From())
			to := fmt.Sprintf("%d", e.To())
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
