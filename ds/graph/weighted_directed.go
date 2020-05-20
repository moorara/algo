package graph

import (
	"fmt"

	"github.com/moorara/algo/ds/list"
	"github.com/moorara/algo/pkg/graphviz"
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
	rev := NewWeightedDirected(g.v)
	for v := 0; v < g.v; v++ {
		for _, e := range g.adj[v] {
			rev.AddEdge(DirectedEdge{e.To(), e.From(), e.Weight()})
		}
	}

	return rev
}

// DFS Traversal (Recursion)
func (g *WeightedDirected) _traverseDFS(visited []bool, v int, order TraversalOrder, vertexVisitor VertexVisitor, edgeVisitor WeightedEdgeVisitor) {
	visited[v] = true

	if order == PreOrder && vertexVisitor != nil {
		if !vertexVisitor.VisitVertex(v) {
			return
		}
	}

	for _, e := range g.adj[v] {
		w := e.To()
		if !visited[w] {
			if order == PreOrder && edgeVisitor != nil {
				if !edgeVisitor.VisitWeightedEdge(v, w, e.Weight()) {
					return
				}
			}

			g._traverseDFS(visited, w, order, vertexVisitor, edgeVisitor)
		}
	}

	if order == PostOrder && vertexVisitor != nil {
		if !vertexVisitor.VisitVertex(v) {
			return
		}
	}
}

// DFS Traversal (Driver)
func (g *WeightedDirected) traverseDFS(s int, order TraversalOrder, vertexVisitor VertexVisitor, edgeVisitor WeightedEdgeVisitor) {
	visited := make([]bool, g.V())
	g._traverseDFS(visited, s, order, vertexVisitor, edgeVisitor)
}

// Iterative DFS Traversal
func (g *WeightedDirected) traverseDFSi(s int, order TraversalOrder, vertexVisitor VertexVisitor, edgeVisitor WeightedEdgeVisitor) {
	visited := make([]bool, g.V())
	stack := list.NewStack(listNodeSize)

	visited[s] = true
	stack.Push(s)

	if order == PreOrder && vertexVisitor != nil {
		if !vertexVisitor.VisitVertex(s) {
			return
		}
	}

	for !stack.IsEmpty() {
		v := stack.Pop().(int)

		if order == PostOrder && vertexVisitor != nil {
			if !vertexVisitor.VisitVertex(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.To()
			if !visited[w] {
				visited[w] = true
				stack.Push(w)

				if order == PreOrder && vertexVisitor != nil {
					if !vertexVisitor.VisitVertex(w) {
						return
					}
				}

				if order == PreOrder && edgeVisitor != nil {
					if !edgeVisitor.VisitWeightedEdge(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// BFS Traversal
func (g *WeightedDirected) traverseBFS(s int, order TraversalOrder, vertexVisitor VertexVisitor, edgeVisitor WeightedEdgeVisitor) {
	visited := make([]bool, g.V())
	queue := list.NewQueue(listNodeSize)

	visited[s] = true
	queue.Enqueue(s)

	if order == PreOrder && vertexVisitor != nil {
		if !vertexVisitor.VisitVertex(s) {
			return
		}
	}

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)

		if order == PostOrder && vertexVisitor != nil {
			if !vertexVisitor.VisitVertex(v) {
				return
			}
		}

		for _, e := range g.adj[v] {
			w := e.To()
			if !visited[w] {
				visited[w] = true
				queue.Enqueue(w)

				if order == PreOrder && vertexVisitor != nil {
					if !vertexVisitor.VisitVertex(w) {
						return
					}
				}

				if order == PreOrder && edgeVisitor != nil {
					if !edgeVisitor.VisitWeightedEdge(v, w, e.Weight()) {
						return
					}
				}
			}
		}
	}
}

// TraverseVertices is used for visiting all vertices in graph.
func (g *WeightedDirected) TraverseVertices(s int, strategy TraversalStrategy, order TraversalOrder, visitor VertexVisitor) {
	if !g.isVertexValid(s) {
		return
	}

	if order != PreOrder && order != PostOrder {
		return
	}

	switch strategy {
	case DFS:
		g.traverseDFS(s, order, visitor, nil)
	case DFSi:
		g.traverseDFSi(s, order, visitor, nil)
	case BFS:
		g.traverseBFS(s, order, visitor, nil)
	}
}

// TraverseEdges is used for visiting all edges in graph.
func (g *WeightedDirected) TraverseEdges(s int, strategy TraversalStrategy, visitor WeightedEdgeVisitor) {
	if !g.isVertexValid(s) {
		return
	}

	switch strategy {
	case DFS:
		g.traverseDFS(s, PreOrder, nil, visitor)
	case DFSi:
		g.traverseDFSi(s, PreOrder, nil, visitor)
	case BFS:
		g.traverseBFS(s, PreOrder, nil, visitor)
	}
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *WeightedDirected) Graphviz() string {
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
