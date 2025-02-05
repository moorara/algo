package dot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdge(t *testing.T) {
	tests := []struct {
		name        string
		from        string
		to          string
		edgeType    EdgeType
		edgeDir     EdgeDir
		label       string
		color       Color
		style       Style
		arrowHead   ArrowType
		arrowTail   ArrowType
		expectedDOT string
	}{
		{
			name:        "DirectedEdge",
			from:        "root",
			to:          "left",
			edgeType:    EdgeTypeDirected,
			edgeDir:     "",
			label:       "",
			color:       "",
			style:       "",
			arrowHead:   "",
			arrowTail:   "",
			expectedDOT: `root -> left [];`,
		},
		{
			name:        "UndirectedEdge",
			from:        "root",
			to:          "right",
			edgeType:    EdgeTypeUndirected,
			edgeDir:     "",
			label:       "",
			color:       "",
			style:       "",
			arrowHead:   "",
			arrowTail:   "",
			expectedDOT: `root -- right [];`,
		},
		{
			name:        "EdgeWithLabel",
			from:        "root",
			to:          "right",
			edgeType:    EdgeTypeUndirected,
			edgeDir:     "",
			label:       `"id"`,
			color:       "",
			style:       "",
			arrowHead:   "",
			arrowTail:   "",
			expectedDOT: `root -- right [label="\"id\""];`,
		},
		{
			name:        "DirectedEdgeWithProperties",
			from:        "parent",
			to:          "child",
			edgeType:    EdgeTypeDirected,
			edgeDir:     EdgeDirNone,
			label:       "red",
			color:       ColorGold,
			style:       StyleDashed,
			arrowHead:   ArrowTypeDot,
			arrowTail:   ArrowTypeODot,
			expectedDOT: `parent -> child [dirType=none, label="red", color=gold, style=dashed, arrowhead=dot, arrowtail=odot];`,
		},
		{
			name:        "UndirectedEdgeWithProperties",
			from:        "parent",
			to:          "child",
			edgeType:    EdgeTypeUndirected,
			edgeDir:     EdgeDirBoth,
			label:       "black",
			color:       ColorOrchid,
			style:       StyleDotted,
			arrowHead:   ArrowTypeBox,
			arrowTail:   ArrowTypeOBox,
			expectedDOT: `parent -- child [dirType=both, label="black", color=orchid, style=dotted, arrowhead=box, arrowtail=obox];`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := NewEdge(tc.from, tc.to, tc.edgeType, tc.edgeDir, tc.label, tc.color, tc.style, tc.arrowHead, tc.arrowTail)
			assert.Equal(t, tc.expectedDOT, e.DOT())
		})
	}
}
