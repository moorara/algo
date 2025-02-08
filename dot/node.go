package dot

import (
	"bytes"
	"fmt"
	"strings"
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
	var b bytes.Buffer

	label := n.Label
	label = strings.ReplaceAll(label, `"`, `\"`)
	label = fmt.Sprintf(`"%s"`, label)

	b.WriteString(n.Name + " [")
	first = addListAttr(&b, first, "group", n.Group)
	first = addListAttr(&b, first, "label", label)
	first = addListAttr(&b, first, "color", string(n.Color))
	first = addListAttr(&b, first, "style", string(n.Style))
	first = addListAttr(&b, first, "shape", string(n.Shape))
	first = addListAttr(&b, first, "fontcolor", string(n.FontColor))
	_ = addListAttr(&b, first, "fontname", `"`+n.FontName+`"`)
	b.WriteString("];")

	return b.String()
}
