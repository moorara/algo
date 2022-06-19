package graphviz

import "bytes"

// Edge represents an edge.
type Edge struct {
	From      string
	To        string
	EdgeType  EdgeType
	EdgeDir   EdgeDir
	Label     string
	Color     Color
	Style     Style
	Arrowhead Arrowhead
}

// NewEdge creates a new edge.
func NewEdge(from, to string, edgeType EdgeType, edgeDir EdgeDir, label string, color Color, style Style, arrowhead Arrowhead) Edge {
	return Edge{
		From:      from,
		To:        to,
		EdgeType:  edgeType,
		EdgeDir:   edgeDir,
		Label:     label,
		Color:     color,
		Style:     style,
		Arrowhead: arrowhead,
	}
}

// DotCode generates Graph dot language code for this edge.
func (e *Edge) DotCode() string {
	first := true
	buf := new(bytes.Buffer)

	buf.WriteString(e.From + " " + string(e.EdgeType) + " " + e.To + " [")
	first = addListAttr(buf, first, "dirType", string(e.EdgeDir))
	first = addListAttr(buf, first, "label", `"`+e.Label+`"`)
	first = addListAttr(buf, first, "color", string(e.Color))
	first = addListAttr(buf, first, "style", string(e.Style))
	first = addListAttr(buf, first, "arrowhead", string(e.Arrowhead))
	buf.WriteString("];")

	return buf.String()
}
