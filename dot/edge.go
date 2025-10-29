package dot

import (
	"bytes"
	"strconv"
)

// Edge represents an edge.
type Edge struct {
	From      string
	To        string
	EdgeType  EdgeType
	EdgeDir   EdgeDir
	Label     string
	Color     Color
	Style     Style
	ArrowHead ArrowType
	ArrowTail ArrowType
}

// NewEdge creates a new edge.
func NewEdge(from, to string, edgeType EdgeType, edgeDir EdgeDir, label string, color Color, style Style, arrowHead, arrowTail ArrowType) Edge {
	return Edge{
		From:      from,
		To:        to,
		EdgeType:  edgeType,
		EdgeDir:   edgeDir,
		Label:     label,
		Color:     color,
		Style:     style,
		ArrowHead: arrowHead,
		ArrowTail: arrowTail,
	}
}

// DOT generates a DOT representation of the Edge object.
func (e *Edge) DOT() string {
	first := true
	var b bytes.Buffer

	label := e.Label
	label = strconv.Quote(label)
	label = labelReplacer.Replace(label)

	b.WriteString(e.From + " " + string(e.EdgeType) + " " + e.To + " [")
	first = addListAttr(&b, first, "dirType", string(e.EdgeDir))
	first = addListAttr(&b, first, "label", label)
	first = addListAttr(&b, first, "color", string(e.Color))
	first = addListAttr(&b, first, "style", string(e.Style))
	first = addListAttr(&b, first, "arrowhead", string(e.ArrowHead))
	_ = addListAttr(&b, first, "arrowtail", string(e.ArrowTail))
	b.WriteString("];")

	return b.String()
}
