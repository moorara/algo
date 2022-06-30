package graph

import (
	"fmt"
	"math"

	"github.com/moorara/algo/internal/graphviz"
)

// FlowEdge represents a capacitated edge data type.
type FlowEdge struct {
	from, to       int
	capacity, flow float64
}

// From returns the tail vertex of the edge.
func (e *FlowEdge) From() int {
	return e.from
}

// To returns the head vertex of the edge.
func (e *FlowEdge) To() int {
	return e.to
}

// Other returns the other vertext of this edge.
func (e *FlowEdge) Other(v int) int {
	switch v {
	case e.from:
		return e.to
	case e.to:
		return e.from
	default:
		return -1
	}
}

// Capacity returns the capacity of the edge.
func (e *FlowEdge) Capacity() float64 {
	return e.capacity
}

// Flow returns the flow on the edge.
func (e *FlowEdge) Flow() float64 {
	return e.flow
}

// ResidualCapacityTo returns the residual capacity of the edge in the direction to the given vertex.
func (e *FlowEdge) ResidualCapacityTo(v int) float64 {
	switch v {
	// backward edge
	case e.from:
		return e.flow
	// forward edge
	case e.to:
		return e.capacity - e.flow
	// invalid
	default:
		return -1
	}
}

// AddResidualFlowTo changes the flow on the edge in the direction to the given vertex.
// If vertex is the tail vertex (backward edge), this decreases the flow on the edge by delta.
// If vertex is the head vertex (forward edge), this increases the flow on the edge by delta.
// If delta is valid, edge will be modified and true will be returned.
// If delta is not valid, edge will not be modified and false will be returned.
func (e *FlowEdge) AddResidualFlowTo(v int, delta float64) bool {
	if delta < 0 {
		return false
	}

	var flow float64

	switch v {
	// backward edge
	case e.from:
		flow = e.flow - delta
	// forward edge
	case e.to:
		flow = e.flow + delta
	// invalid
	default:
		return false
	}

	// Round flow to 0 or capacity if within floating point precision
	if math.Abs(flow) <= float64Epsilon {
		flow = 0
	}
	if math.Abs(flow-e.capacity) <= float64Epsilon {
		flow = e.capacity
	}

	if flow < 0 || flow > e.capacity {
		return false
	}

	e.flow = flow

	return true
}

// FlowNetwork represents a capacitated network graph data type.
// Each directed edge has a real number capacity and flow.
type FlowNetwork struct {
	v, e int
	adj  [][]FlowEdge
}

// NewFlowNetwork creates a new capacitated network graph.
func NewFlowNetwork(V int, edges ...FlowEdge) *FlowNetwork {
	adj := make([][]FlowEdge, V)
	for i := range adj {
		adj[i] = make([]FlowEdge, 0)
	}

	g := &FlowNetwork{
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
func (g *FlowNetwork) V() int {
	return g.v
}

// E returns the number of edges.
func (g *FlowNetwork) E() int {
	return g.e
}

func (g *FlowNetwork) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// Adj returns the edges both pointing to and from a vertex.
func (g *FlowNetwork) Adj(v int) []FlowEdge {
	if !g.isVertexValid(v) {
		return nil
	}
	return g.adj[v]
}

// AddEdge adds a new edge to the network.
func (g *FlowNetwork) AddEdge(e FlowEdge) {
	v := e.From()
	w := e.To()

	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], e)
		g.adj[w] = append(g.adj[w], e)
	}
}

// Edges returns all directed edges in the graph.
func (g *FlowNetwork) Edges() []FlowEdge {
	edges := make([]FlowEdge, 0)
	for v, adjEdges := range g.adj {
		for _, e := range adjEdges {
			if e.To() != v {
				edges = append(edges, e)
			}
		}
	}

	return edges
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *FlowNetwork) Graphviz() string {
	graph := graphviz.NewGraph(true, true, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode("", "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, e := range g.adj[v] {
			from := fmt.Sprintf("%d", e.From())
			to := fmt.Sprintf("%d", e.To())
			label := fmt.Sprintf("%f/%f", e.Flow(), e.Capacity())
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeDirected, "", label, "", "", "", ""))
		}
	}

	return graph.DotCode()
}
