package graph

import (
	"fmt"

	"github.com/moorara/algo/internal/graphviz"
	"github.com/moorara/algo/list"
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

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	return g
}

// V returns the number of vertices.
func (g *Undirected) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Undirected) E() int {
	return g.e
}

func (g *Undirected) isVertexValid(v int) bool {
	return v >= 0 && v < g.v
}

// Degree returns the degree of a vertex.
func (g *Undirected) Degree(v int) int {
	if !g.isVertexValid(v) {
		return -1
	}
	return len(g.adj[v])
}

// Adj returns the vertices adjacent from vertex.
func (g *Undirected) Adj(v int) []int {
	if !g.isVertexValid(v) {
		return nil
	}
	return g.adj[v]
}

// AddEdge adds a new edge to the graph.
func (g *Undirected) AddEdge(v, w int) {
	if g.isVertexValid(v) && g.isVertexValid(w) {
		g.e++
		g.adj[v] = append(g.adj[v], w)
		g.adj[w] = append(g.adj[w], v)
	}
}

// DFS Traversal (Recursion)
func (g *Undirected) traverseDFS(v int, visited []bool, visitors *Visitors) {
	visited[v] = true

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(v) {
			return
		}
	}

	for _, w := range g.adj[v] {
		if !visited[w] {
			if visitors != nil && visitors.EdgePreOrder != nil {
				if !visitors.EdgePreOrder(v, w, 0) {
					return
				}
			}

			g.traverseDFS(w, visited, visitors)
		}
	}

	if visitors != nil && visitors.VertexPostOrder != nil {
		if !visitors.VertexPostOrder(v) {
			return
		}
	}
}

// Iterative DFS Traversal
func (g *Undirected) traverseDFSi(s int, visited []bool, visitors *Visitors) {
	stack := list.NewStack[int](listNodeSize, nil)

	visited[s] = true
	stack.Push(s)

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !stack.IsEmpty() {
		v, _ := stack.Pop()

		if visitors != nil && visitors.VertexPostOrder != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, w := range g.adj[v] {
			if !visited[w] {
				visited[w] = true
				stack.Push(w)

				if visitors != nil && visitors.VertexPreOrder != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
				}

				if visitors != nil && visitors.EdgePreOrder != nil {
					if !visitors.EdgePreOrder(v, w, 0) {
						return
					}
				}
			}
		}
	}
}

// BFS Traversal
func (g *Undirected) traverseBFS(s int, visited []bool, visitors *Visitors) {
	queue := list.NewQueue[int](listNodeSize, nil)

	visited[s] = true
	queue.Enqueue(s)

	if visitors != nil && visitors.VertexPreOrder != nil {
		if !visitors.VertexPreOrder(s) {
			return
		}
	}

	for !queue.IsEmpty() {
		v, _ := queue.Dequeue()

		if visitors != nil && visitors.VertexPostOrder != nil {
			if !visitors.VertexPostOrder(v) {
				return
			}
		}

		for _, w := range g.adj[v] {
			if !visited[w] {
				visited[w] = true
				queue.Enqueue(w)

				if visitors != nil && visitors.VertexPreOrder != nil {
					if !visitors.VertexPreOrder(w) {
						return
					}
				}

				if visitors != nil && visitors.EdgePreOrder != nil {
					if !visitors.EdgePreOrder(v, w, 0) {
						return
					}
				}
			}
		}
	}
}

// Traverse is used for visiting all vertices and edges in graph.
func (g *Undirected) Traverse(s int, strategy TraversalStrategy, visitors *Visitors) {
	if !g.isVertexValid(s) {
		return
	}

	visited := make([]bool, g.V())

	switch strategy {
	case DFS:
		g.traverseDFS(s, visited, visitors)
	case DFSi:
		g.traverseDFSi(s, visited, visitors)
	case BFS:
		g.traverseBFS(s, visited, visitors)
	}
}

// Paths finds all paths from a source vertex to every other vertex.
func (g *Undirected) Paths(s int, strategy TraversalStrategy) *Paths {
	p := &Paths{
		s:       s,
		visited: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
	}

	if g.isVertexValid(s) && isStrategyValid(strategy) {
		visitors := &Visitors{
			EdgePreOrder: func(v, w int, _ float64) bool {
				p.edgeTo[w] = v
				return true
			},
		}

		switch strategy {
		case DFS:
			g.traverseDFS(s, p.visited, visitors)
		case DFSi:
			g.traverseDFSi(s, p.visited, visitors)
		case BFS:
			g.traverseBFS(s, p.visited, visitors)
		}
	}

	return p
}

// Orders determines ordering of vertices in the graph.
func (g *Undirected) Orders(strategy TraversalStrategy) *Orders {
	o := &Orders{
		preRank:   make([]int, g.V()),
		postRank:  make([]int, g.V()),
		preOrder:  make([]int, 0),
		postOrder: make([]int, 0),
	}

	if isStrategyValid(strategy) {
		var preCounter, postCounter int
		visited := make([]bool, g.V())
		visitors := &Visitors{
			VertexPreOrder: func(v int) bool {
				o.preRank[v] = preCounter
				preCounter++
				o.preOrder = append(o.preOrder, v)
				return true
			},
			VertexPostOrder: func(v int) bool {
				o.postRank[v] = postCounter
				postCounter++
				o.postOrder = append(o.postOrder, v)
				return true
			},
		}

		for v := 0; v < g.V(); v++ {
			if !visited[v] {
				switch strategy {
				case DFS:
					g.traverseDFS(v, visited, visitors)
				case DFSi:
					g.traverseDFSi(v, visited, visitors)
				case BFS:
					g.traverseBFS(v, visited, visitors)
				}
			}
		}
	}

	return o
}

// ConnectedComponents determines all the connected components in the graph.
func (g *Undirected) ConnectedComponents() *ConnectedComponents {
	cc := &ConnectedComponents{
		count: 0,
		id:    make([]int, g.V()),
	}

	visited := make([]bool, g.V())
	visitors := &Visitors{
		VertexPreOrder: func(v int) bool {
			cc.id[v] = cc.count
			return true
		},
	}

	for v := 0; v < g.V(); v++ {
		if !visited[v] {
			g.traverseDFS(v, visited, visitors)
			cc.count++
		}
	}

	return cc
}

// Graphviz returns a visualization of the graph in Graphviz format.
func (g *Undirected) Graphviz() string {
	graph := graphviz.NewGraph(true, false, false, "", "", "", graphviz.StyleSolid, graphviz.ShapeCircle)

	for i := 0; i < g.v; i++ {
		name := fmt.Sprintf("%d", i)
		graph.AddNode(graphviz.NewNode(name, "", name, "", "", "", "", ""))
	}

	for v := range g.adj {
		for _, w := range g.adj[v] {
			from := fmt.Sprintf("%d", v)
			to := fmt.Sprintf("%d", w)
			graph.AddEdge(graphviz.NewEdge(from, to, graphviz.EdgeTypeUndirected, "", "", "", "", "", ""))
		}
	}

	return graph.DotCode()
}
