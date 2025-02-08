package dot

import (
	"bytes"
	"fmt"
)

// Graph represents a graph.
type Graph struct {
	Strict      bool
	Digraph     bool
	Concentrate bool
	Name        string
	RankDir     RankDir
	NodeColor   Color
	NodeStyle   Style
	NodeShape   Shape
	Nodes       []Node
	Edges       []Edge
	Subgraphs   []Subgraph
}

// NewGraph creates a new graph.
func NewGraph(strict, digraph, concentrate bool, name string, rankDir RankDir, nodeColor Color, nodeStyle Style, nodeShape Shape) Graph {
	if name != "" {
		name = fmt.Sprintf("%q", name)
	}

	return Graph{
		Strict:      strict,
		Digraph:     digraph,
		Concentrate: concentrate,
		Name:        name,
		RankDir:     rankDir,
		NodeColor:   nodeColor,
		NodeStyle:   nodeStyle,
		NodeShape:   nodeShape,
		Nodes:       make([]Node, 0),
		Edges:       make([]Edge, 0),
		Subgraphs:   make([]Subgraph, 0),
	}
}

// AddNode adds a new node to this graph.
func (g *Graph) AddNode(nodes ...Node) {
	g.Nodes = append(g.Nodes, nodes...)
}

// AddEdge adds a new edge to this graph.
func (g *Graph) AddEdge(edges ...Edge) {
	g.Edges = append(g.Edges, edges...)
}

// AddSubgraph adds a new subgraph to this graph.
func (g *Graph) AddSubgraph(subgraphs ...Subgraph) {
	g.Subgraphs = append(g.Subgraphs, subgraphs...)
}

// DOT generates a DOT representation of the Graph object.
func (g *Graph) DOT() string {
	first := true
	var b bytes.Buffer

	if g.Strict {
		b.WriteString("strict ")
	}

	if g.Digraph {
		b.WriteString("digraph ")
	} else {
		b.WriteString("graph ")
	}

	if g.Name != "" {
		b.WriteString(g.Name)
		b.WriteString(" ")
	}

	b.WriteString("{\n")

	first = addAttr(&b, first, 2, "rankdir", string(g.RankDir))
	_ = addAttr(&b, first, 2, "concentrate", fmt.Sprintf("%t", g.Concentrate))

	first = true
	addIndent(&b, 2)
	b.WriteString("node [")
	first = addListAttr(&b, first, "color", string(g.NodeColor))
	first = addListAttr(&b, first, "style", string(g.NodeStyle))
	_ = addListAttr(&b, first, "shape", string(g.NodeShape))
	b.WriteString("];\n")
	first = false

	first = addSubgraphs(&b, first, 2, g.Subgraphs)
	first = addNodes(&b, first, 2, g.Nodes)
	addEdges(&b, first, 2, g.Edges)
	b.WriteString("}")

	return b.String()
}
