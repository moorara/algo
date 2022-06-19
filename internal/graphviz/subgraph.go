package graphviz

import "bytes"

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

// DotCode generates Graph dot language code for this subgraph.
func (s *Subgraph) DotCode(indent int) string {
	first := true
	buf := new(bytes.Buffer)

	addIndent(buf, indent)
	buf.WriteString("subgraph ")
	if s.Name != "" {
		buf.WriteString(s.Name)
		buf.WriteString(" ")
	}
	buf.WriteString("{\n")

	first = addAttr(buf, first, indent+2, "label", `"`+s.Label+`"`)
	first = addAttr(buf, first, indent+2, "color", string(s.Color))
	first = addAttr(buf, first, indent+2, "style", string(s.Style))
	first = addAttr(buf, first, indent+2, "rank", string(s.Rank))
	first = addAttr(buf, first, indent+2, "rankdir", string(s.RankDir))

	first = true
	addIndent(buf, indent+2)
	buf.WriteString("node [")
	first = addListAttr(buf, first, "color", string(s.NodeColor))
	first = addListAttr(buf, first, "style", string(s.NodeStyle))
	first = addListAttr(buf, first, "shape", string(s.NodeShape))
	buf.WriteString("];\n")
	first = false

	first = addSubgraphs(buf, first, indent+2, s.Subgraphs)
	first = addNodes(buf, first, indent+2, s.Nodes)
	first = addEdges(buf, first, indent+2, s.Edges)

	addIndent(buf, indent)
	buf.WriteString("}")

	return buf.String()
}
