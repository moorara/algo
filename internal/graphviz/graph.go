package graphviz

import "bytes"

// Graph represents a graph.
type Graph struct {
	Strict    bool
	Digraph   bool
	Name      string
	RankDir   RankDir
	NodeColor Color
	NodeStyle Style
	NodeShape Shape
	Nodes     []Node
	Edges     []Edge
	Subgraphs []Subgraph
}

// NewGraph creates a new graph.
func NewGraph(strict, digraph bool, name string, rankDir RankDir, nodeColor Color, nodeStyle Style, nodeShape Shape) Graph {
	return Graph{
		Strict:    strict,
		Digraph:   digraph,
		Name:      name,
		RankDir:   rankDir,
		NodeColor: nodeColor,
		NodeStyle: nodeStyle,
		NodeShape: nodeShape,
		Nodes:     make([]Node, 0),
		Edges:     make([]Edge, 0),
		Subgraphs: make([]Subgraph, 0),
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

// DotCode generates Graph dot language code for this graph.
func (g *Graph) DotCode() string {
	first := true
	buf := new(bytes.Buffer)

	if g.Strict {
		buf.WriteString("strict ")
	}

	if g.Digraph {
		buf.WriteString("digraph ")
	} else {
		buf.WriteString("graph ")
	}

	if g.Name != "" {
		buf.WriteString(g.Name)
		buf.WriteString(" ")
	}

	buf.WriteString("{\n")

	first = addAttr(buf, first, 2, "rankdir", string(g.RankDir))

	first = true
	addIndent(buf, 2)
	buf.WriteString("node [")
	first = addListAttr(buf, first, "color", string(g.NodeColor))
	first = addListAttr(buf, first, "style", string(g.NodeStyle))
	first = addListAttr(buf, first, "shape", string(g.NodeShape))
	buf.WriteString("];\n")
	first = false

	first = addSubgraphs(buf, first, 2, g.Subgraphs)
	first = addNodes(buf, first, 2, g.Nodes)
	addEdges(buf, first, 2, g.Edges)
	buf.WriteString("}")

	return buf.String()
}
