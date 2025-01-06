package dot

import (
	"bytes"
	"fmt"
)

// Node represents a graph node.
type Node struct {
	Name      string
	Group     string
	Label     string
	Color     Color
	Style     Style
	Shape     Shape
	FontColor Color
	FontName  string
}

// NewNode creates a new node.
func NewNode(name, group, label string, color Color, style Style, shape Shape, fontcolor Color, fontname string) Node {
	return Node{
		Name:      name,
		Group:     group,
		Label:     label,
		Color:     color,
		Style:     style,
		Shape:     shape,
		FontColor: fontcolor,
		FontName:  fontname,
	}
}

// DOT generates a DOT representation of the Node object.
func (n *Node) DOT() string {
	first := true
	buf := new(bytes.Buffer)

	buf.WriteString(n.Name + " [")
	first = addListAttr(buf, first, "group", n.Group)
	first = addListAttr(buf, first, "label", fmt.Sprintf("%q", n.Label))
	first = addListAttr(buf, first, "color", string(n.Color))
	first = addListAttr(buf, first, "style", string(n.Style))
	first = addListAttr(buf, first, "shape", string(n.Shape))
	first = addListAttr(buf, first, "fontcolor", string(n.FontColor))
	_ = addListAttr(buf, first, "fontname", `"`+n.FontName+`"`)
	buf.WriteString("];")

	return buf.String()
}
