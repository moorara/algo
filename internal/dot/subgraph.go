package dot

import (
	"bytes"
	"fmt"
)

// Subgraph represents a subgraph.
type Subgraph struct {
	Name      string
	Label     string
	Color     Color
	Style     Style
	Rank      Rank
	RankDir   RankDir
	NodeColor Color
	NodeStyle Style
	NodeShape Shape
	Nodes     []Node
	Edges     []Edge
	Subgraphs []Subgraph
}

// NewSubgraph creates a new subgraph.
func NewSubgraph(name, label string, color Color, style Style, rank Rank, rankDir RankDir, nodeColor Color, nodeStyle Style, nodeShape Shape) Subgraph {
	return Subgraph{
		Name:      name,
		Label:     label,
		Color:     color,
		Style:     style,
		Rank:      rank,
		RankDir:   rankDir,
		NodeColor: nodeColor,
		NodeStyle: nodeStyle,
		NodeShape: nodeShape,
		Nodes:     make([]Node, 0),
		Edges:     make([]Edge, 0),
		Subgraphs: make([]Subgraph, 0),
	}
}

// AddNode adds a new node to this subgraph.
func (s *Subgraph) AddNode(nodes ...Node) {
	s.Nodes = append(s.Nodes, nodes...)
}

// AddEdge adds a new edge to this subgraph.
func (s *Subgraph) AddEdge(edges ...Edge) {
	s.Edges = append(s.Edges, edges...)
}

// AddSubgraph adds a new subgraph to this subgraph.
func (s *Subgraph) AddSubgraph(subgraphs ...Subgraph) {
	s.Subgraphs = append(s.Subgraphs, subgraphs...)
}

// DOT generates a DOT representation of the Subgraph object.
func (s *Subgraph) DOT(indent int) string {
	first := true
	var b bytes.Buffer

	addIndent(&b, indent)
	b.WriteString("subgraph ")
	if s.Name != "" {
		b.WriteString(s.Name)
		b.WriteString(" ")
	}
	b.WriteString("{\n")

	first = addAttr(&b, first, indent+2, "label", fmt.Sprintf("%q", s.Label))
	first = addAttr(&b, first, indent+2, "color", string(s.Color))
	first = addAttr(&b, first, indent+2, "style", string(s.Style))
	first = addAttr(&b, first, indent+2, "rank", string(s.Rank))
	_ = addAttr(&b, first, indent+2, "rankdir", string(s.RankDir))

	first = true
	addIndent(&b, indent+2)
	b.WriteString("node [")
	first = addListAttr(&b, first, "color", string(s.NodeColor))
	first = addListAttr(&b, first, "style", string(s.NodeStyle))
	_ = addListAttr(&b, first, "shape", string(s.NodeShape))
	b.WriteString("];\n")
	first = false

	first = addSubgraphs(&b, first, indent+2, s.Subgraphs)
	first = addNodes(&b, first, indent+2, s.Nodes)
	_ = addEdges(&b, first, indent+2, s.Edges)

	addIndent(&b, indent)
	b.WriteString("}")

	return b.String()
}
