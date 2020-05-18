package graph

import (
	"fmt"

	"github.com/moorara/algo/ds/list"
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

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	return g
}

// V returns the number of vertices.
func (g *Directed) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Directed) E() int {
	return g.e
}

func (g *Directed) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
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

// Adj returns the vertices adjacent from vertex.
func (g *Directed) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}
	return g.adj[v]
}

// AddEdge adds a new edge to the graph.
func (g *Directed) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.ins[w]++
		g.adj[v] = append(g.adj[v], w)
	}
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

// DFS Traversal (Recursion)
func (g *Directed) _traverseDFS(visited []bool, v int, order TraverseOrder, visitor *Visitor) {
	visited[v] = true

	if order == PreOrder && visitor != nil && visitor.VisitVertex != nil {
		visitor.VisitVertex(v)
	}

	for _, w := range g.adj[v] {
		if !visited[w] {
			g._traverseDFS(visited, w, order, visitor)
		}
	}

	if order == PostOrder && visitor != nil && visitor.VisitVertex != nil {
		visitor.VisitVertex(v)
	}
}

// DFS Traversal (Driver)
func (g *Directed) traverseDFS(s int, order TraverseOrder, visitor *Visitor) {
	visited := make([]bool, g.V())
	g._traverseDFS(visited, s, order, visitor)
}

// Iterative DFS Traversal
func (g *Directed) traverseDFSIterative(s int, order TraverseOrder, visitor *Visitor) {
	visited := make([]bool, g.V())
	stack := list.NewStack(listNodeSize)

	visited[s] = true
	stack.Push(s)
	if order == PreOrder && visitor != nil && visitor.VisitVertex != nil {
		visitor.VisitVertex(s)
	}

	for !stack.IsEmpty() {
		v := stack.Pop().(int)
		if order == PostOrder && visitor != nil && visitor.VisitVertex != nil {
			visitor.VisitVertex(v)
		}

		for _, w := range g.adj[v] {
			if !visited[w] {
				visited[w] = true
				stack.Push(w)
				if order == PreOrder && visitor != nil && visitor.VisitVertex != nil {
					visitor.VisitVertex(w)
				}
			}
		}
	}
}

// BFS Traversal
func (g *Directed) traverseBFS(s int, order TraverseOrder, visitor *Visitor) {
	visited := make([]bool, g.V())
	queue := list.NewQueue(listNodeSize)

	visited[s] = true
	queue.Enqueue(s)
	if order == PreOrder && visitor != nil && visitor.VisitVertex != nil {
		visitor.VisitVertex(s)
	}

	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)
		if order == PostOrder && visitor != nil && visitor.VisitVertex != nil {
			visitor.VisitVertex(v)
		}

		for _, w := range g.adj[v] {
			if !visited[w] {
				visited[w] = true
				queue.Enqueue(w)
				if order == PreOrder && visitor != nil && visitor.VisitVertex != nil {
					visitor.VisitVertex(w)
				}
			}
		}
	}
}

// Traverse is used for visiting all vertices in graph.
func (g *Directed) Traverse(s int, strategy TraverseStrategy, order TraverseOrder, visitor *Visitor) {
	if !g.isVertexValid(s) {
		return
	}

	if order != PreOrder && order != PostOrder {
		return
	}

	switch strategy {
	case DFS:
		g.traverseDFS(s, order, visitor)
	case DFSIterative:
		g.traverseDFSIterative(s, order, visitor)
	case BFS:
		g.traverseBFS(s, order, visitor)
	}
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
