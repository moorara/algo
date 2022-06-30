package graphviz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	tests := []struct {
		name            string
		from            string
		to              string
		edgeType        EdgeType
		edgeDir         EdgeDir
		label           string
		color           Color
		style           Style
		arrowHead       ArrowType
		arrowTail       ArrowType
		expectedDotCode string
	}{
		{
			"SimpleEdge",
			"root", "left", EdgeTypeDirected, "",
			"", "", "", "", "",
			`root -> left [];`,
		},
		{
			"EdgeWithType",
			"root", "right", EdgeTypeUndirected, "",
			"normal", "", "", "", "",
			`root -- right [label="normal"];`,
		},
		{
			"EdgeWithProperties",
			"parent", "child", EdgeTypeDirected, EdgeDirNone,
			"red", ColorGold, StyleDashed, ArrowTypeDot, ArrowTypeODot,
			`parent -> child [dirType=none, label="red", color=gold, style=dashed, arrowhead=dot, arrowtail=odot];`,
		},
		{
			"EdgeWithProperties",
			"parent", "child", EdgeTypeUndirected, EdgeDirBoth,
			"black", ColorOrchid, StyleDotted, ArrowTypeBox, ArrowTypeOBox,
			`parent -- child [dirType=both, label="black", color=orchid, style=dotted, arrowhead=box, arrowtail=obox];`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := NewEdge(tc.from, tc.to, tc.edgeType, tc.edgeDir, tc.label, tc.color, tc.style, tc.arrowHead, tc.arrowTail)
			assert.Equal(t, tc.expectedDotCode, e.DotCode())
		})
	}
}
