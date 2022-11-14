package graphviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	tests := []struct {
		name            string
		nodeName        string
		group           string
		label           string
		color           Color
		style           Style
		shape           Shape
		fontcolor       Color
		fontname        string
		expectedDotCode string
	}{
		{
			name:            "SimpleNode",
			nodeName:        "root",
			group:           "",
			label:           "",
			color:           "",
			style:           "",
			shape:           "",
			fontcolor:       "",
			fontname:        "",
			expectedDotCode: `root [];`,
		},
		{
			name:            "NodeWithLabel",
			nodeName:        "root",
			group:           "",
			label:           "root",
			color:           "",
			style:           "",
			shape:           "",
			fontcolor:       "",
			fontname:        "",
			expectedDotCode: `root [label="root"];`,
		},
		{
			name:            "NodeWithGroup",
			nodeName:        "struct0",
			group:           "group0",
			label:           "<f0> left|<f1> middle|<f2> right",
			color:           ColorBlue,
			style:           StyleBold,
			shape:           ShapeBox,
			fontcolor:       ColorGray,
			fontname:        "",
			expectedDotCode: `struct0 [group=group0, label="<f0> left|<f1> middle|<f2> right", color=blue, style=bold, shape=box, fontcolor=gray];`,
		},
		{
			name:            "ComplexNode",
			nodeName:        "struct1",
			group:           "group1",
			label:           "a | { b | { c | <here> d | e } | f } | g | h",
			color:           ColorNavy,
			style:           StyleDashed,
			shape:           ShapeOval,
			fontcolor:       ColorBlack,
			fontname:        "Arial",
			expectedDotCode: `struct1 [group=group1, label="a | { b | { c | <here> d | e } | f } | g | h", color=navy, style=dashed, shape=oval, fontcolor=black, fontname="Arial"];`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			n := NewNode(tc.nodeName, tc.group, tc.label, tc.color, tc.style, tc.shape, tc.fontcolor, tc.fontname)
			assert.Equal(t, tc.expectedDotCode, n.DotCode())
		})
	}
}
